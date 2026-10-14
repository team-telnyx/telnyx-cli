package smtpverify_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/team-telnyx/telnyx-cli/internal/smtpverify"
	"github.com/team-telnyx/telnyx-cli/internal/smtpverify/smtpmock"
)

const testKey = "KEYtest_0123456789abcdef"

var fixedNow = func() time.Time { return time.Date(2026, 10, 14, 12, 0, 0, 0, time.UTC) }

func config(srv *smtpmock.Server) smtpverify.Config {
	return smtpverify.Config{
		Host:           srv.Host,
		Port:           srv.Port,
		APIKey:         testKey,
		From:           "verify@example.com",
		TLSConfig:      srv.ClientTLS,
		CommandTimeout: 3 * time.Second,
		DataTimeout:    3 * time.Second,
		Now:            fixedNow,
	}
}

func run(t *testing.T, cfg smtpverify.Config) (smtpverify.Result, string) {
	t.Helper()
	var out bytes.Buffer
	res, err := smtpverify.Run(context.Background(), cfg, &out)
	require.NoError(t, err)
	return res, out.String()
}

var (
	verbsUpToAuth = []string{"EHLO", "STARTTLS", "EHLO", "AUTH"}
	verbsUpToMail = append(append([]string{}, verbsUpToAuth...), "MAIL")
	verbsUpToRcpt = append(append([]string{}, verbsUpToMail...), "RCPT")
	verbsUpToData = append(append([]string{}, verbsUpToRcpt...), "DATA")
)

func with(verbs []string, more ...string) []string {
	return append(append([]string{}, verbs...), more...)
}

func TestRunSessions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		script      smtpmock.Script
		mutate      func(*smtpverify.Config)
		wantVerdict string // exact, unless wantPrefix
		wantPrefix  bool
		wantStage   smtpverify.Stage
		wantVerbs   []string // commands the server received, in order
		wantHint    string   // substring
		wantOutput  []string // transcript substrings
	}{
		{
			name:        "verified",
			wantVerdict: "VERIFIED: message UUID " + smtpmock.QueuedUUID,
			wantStage:   smtpverify.StageQuit,
			wantVerbs:   with(verbsUpToData, "QUIT"),
		},
		{
			name:        "banner refused",
			script:      smtpmock.Script{Banner: "554 5.3.2 Service not available"},
			wantVerdict: "FAILED at CONNECT: 554 5.3.2 Service not available",
			wantStage:   smtpverify.StageConnect,
			wantVerbs:   []string{"QUIT"},
		},
		{
			name:        "EHLO rejected",
			script:      smtpmock.Script{EHLOReply: "502 5.5.1 Unrecognized command"},
			wantVerdict: "FAILED at EHLO: 502 5.5.1 Unrecognized command",
			wantStage:   smtpverify.StageEHLO,
			wantVerbs:   []string{"EHLO", "QUIT"},
		},
		{
			name:        "STARTTLS not advertised: credentials never sent",
			script:      smtpmock.Script{NoSTARTTLS: true},
			wantVerdict: "FAILED at STARTTLS: server did not advertise STARTTLS; refusing to send credentials without TLS",
			wantStage:   smtpverify.StageSTARTTLS,
			wantVerbs:   []string{"EHLO", "QUIT"},
			wantHint:    "538 5.7.11 Encryption required",
		},
		{
			name:        "STARTTLS refused",
			script:      smtpmock.Script{STARTTLSReply: "454 4.7.0 TLS not available due to temporary reason"},
			wantVerdict: "FAILED at STARTTLS: 454 4.7.0 TLS not available due to temporary reason",
			wantStage:   smtpverify.StageSTARTTLS,
			wantVerbs:   []string{"EHLO", "STARTTLS", "QUIT"},
		},
		{
			name:        "untrusted certificate: credentials never sent",
			mutate:      func(c *smtpverify.Config) { c.TLSConfig = nil },
			wantVerdict: "FAILED at STARTTLS: TLS handshake failed: ",
			wantPrefix:  true,
			wantStage:   smtpverify.StageSTARTTLS,
			wantVerbs:   []string{"EHLO", "STARTTLS"},
		},
		{
			name:        "invalid API key",
			script:      smtpmock.Script{AuthReply: "535 5.7.8 Authentication credentials invalid"},
			wantVerdict: "FAILED at AUTH: 535 5.7.8 Authentication credentials invalid",
			wantStage:   smtpverify.StageAuth,
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantHint:    `literal "apikey"`,
		},
		{
			name:        "key lacks email send capability",
			script:      smtpmock.Script{AuthReply: "535 5.7.1 Authorization failed: email send capability required"},
			wantVerdict: "FAILED at AUTH: 535 5.7.1 Authorization failed: email send capability required",
			wantStage:   smtpverify.StageAuth,
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantHint:    "email send capability",
		},
		{
			name:        "AUTH throttled: 421 closes, no QUIT",
			script:      smtpmock.Script{AuthReply: "421 4.7.0 Temporary authentication failure, retry later"},
			wantVerdict: "FAILED at AUTH: 421 4.7.0 Temporary authentication failure, retry later",
			wantStage:   smtpverify.StageAuth,
			wantVerbs:   verbsUpToAuth,
		},
		{
			name:        "AUTH PLAIN not advertised is attempted anyway",
			script:      smtpmock.Script{NoAuthAdvert: true},
			wantVerdict: "VERIFIED: message UUID " + smtpmock.QueuedUUID,
			wantStage:   smtpverify.StageQuit,
			wantVerbs:   with(verbsUpToData, "QUIT"),
		},
		{
			name:        "sender domain not verified",
			script:      smtpmock.Script{MailReply: "550 5.7.1 Domain is not verified. Complete DNS setup before sending email."},
			wantVerdict: "FAILED at MAIL FROM: 550 5.7.1 Domain is not verified. Complete DNS setup before sending email.",
			wantStage:   smtpverify.StageMailFrom,
			wantVerbs:   with(verbsUpToMail, "QUIT"),
			wantHint:    "verified sending domain",
		},
		{
			name:        "sender domain not on account",
			script:      smtpmock.Script{MailReply: "550 5.7.1 Sending domain example.com is not available for this account"},
			wantVerdict: "FAILED at MAIL FROM: 550 5.7.1 Sending domain example.com is not available for this account",
			wantStage:   smtpverify.StageMailFrom,
			wantVerbs:   with(verbsUpToMail, "QUIT"),
		},
		{
			name:        "domain lookup outage",
			script:      smtpmock.Script{MailReply: "451 4.3.0 Temporary submission failure, try again"},
			wantVerdict: "FAILED at MAIL FROM: 451 4.3.0 Temporary submission failure, try again",
			wantStage:   smtpverify.StageMailFrom,
			wantVerbs:   with(verbsUpToMail, "QUIT"),
		},
		{
			name:        "session expired",
			script:      smtpmock.Script{MailReply: "421 4.4.2 Session expired, close and retry"},
			wantVerdict: "FAILED at MAIL FROM: 421 4.4.2 Session expired, close and retry",
			wantStage:   smtpverify.StageMailFrom,
			wantVerbs:   verbsUpToMail,
		},
		{
			name:        "recipient cap",
			script:      smtpmock.Script{RcptReply: "452 4.5.3 Too many recipients"},
			wantVerdict: "FAILED at RCPT TO: 452 4.5.3 Too many recipients",
			wantStage:   smtpverify.StageRcptTo,
			wantVerbs:   with(verbsUpToRcpt, "QUIT"),
		},
		{
			name:        "DATA command refused",
			script:      smtpmock.Script{DataReply: "503 Error: need RCPT command"},
			wantVerdict: "FAILED at DATA: 503 Error: need RCPT command",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
		},
		{
			name:        "unauthenticated at end of DATA",
			script:      smtpmock.Script{DataEndReply: "530 5.7.0 Authentication required"},
			wantVerdict: "FAILED at DATA: 530 5.7.0 Authentication required",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
		},
		{
			name:        "oversize refused by the server's SIZE guard",
			script:      smtpmock.Script{DataEndReply: "552 Message too large"},
			wantVerdict: "FAILED at DATA: 552 Message too large",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantHint:    "SIZE limit",
		},
		{
			name:        "idle timeout reply closes the session",
			script:      smtpmock.Script{MailReply: "421 Error: timeout exceeded"},
			wantVerdict: "FAILED at MAIL FROM: 421 Error: timeout exceeded",
			wantStage:   smtpverify.StageMailFrom,
			wantVerbs:   verbsUpToMail,
			wantHint:    "idle timeout",
		},
		{
			name:        "entitlement lost after AUTH",
			script:      smtpmock.Script{RcptReply: "421 4.7.0 Temporary authentication failure, retry later"},
			wantVerdict: "FAILED at RCPT TO: 421 4.7.0 Temporary authentication failure, retry later",
			wantStage:   smtpverify.StageRcptTo,
			wantVerbs:   verbsUpToRcpt,
			wantHint:    "revoked",
		},
		{
			name:        "AUTH continuation challenge is cancelled before QUIT",
			script:      smtpmock.Script{AuthReply: "334 VXNlcm5hbWU6"},
			wantVerdict: "FAILED at AUTH: 334 VXNlcm5hbWU6",
			wantStage:   smtpverify.StageAuth,
			wantVerbs:   with(verbsUpToAuth, "*", "QUIT"),
		},
		{
			name:        "wrong certificate name: credentials never sent",
			mutate:      func(c *smtpverify.Config) { wrongName(c) },
			wantVerdict: "FAILED at STARTTLS: TLS handshake failed: ",
			wantPrefix:  true,
			wantStage:   smtpverify.StageSTARTTLS,
			wantVerbs:   []string{"EHLO", "STARTTLS"},
			wantHint:    "certificate",
			wantOutput:  []string{"not wrong.example"},
		},
		{
			name:        "certificate names cannot inject terminal escapes",
			script:      smtpmock.Script{CertName: "evil\x1b[2K\rVERIFIED: message UUID " + smtpmock.QueuedUUID},
			mutate:      func(c *smtpverify.Config) { wrongName(c) },
			wantVerdict: "FAILED at STARTTLS: TLS handshake failed: ",
			wantPrefix:  true,
			wantStage:   smtpverify.StageSTARTTLS,
			wantVerbs:   []string{"EHLO", "STARTTLS"},
			wantOutput:  []string{`evil\x1b[2K\x0dVERIFIED`},
		},
		{
			name:        "421 before AUTH is not blamed on the key",
			script:      smtpmock.Script{Banner: "421 4.7.0 Too many connections"},
			wantVerdict: "FAILED at CONNECT: 421 4.7.0 Too many connections",
			wantStage:   smtpverify.StageConnect,
			wantVerbs:   nil,
			wantHint:    "temporarily refusing service",
		},
		{
			name:        "stalled TLS handshake is a timeout, not a certificate problem",
			script:      smtpmock.Script{StallTLS: true},
			mutate:      func(c *smtpverify.Config) { c.CommandTimeout = 150 * time.Millisecond },
			wantVerdict: "FAILED at STARTTLS: TLS handshake: timed out after 150ms waiting for the server",
			wantStage:   smtpverify.StageSTARTTLS,
			wantVerbs:   []string{"EHLO", "STARTTLS"},
			wantHint:    "--timeout",
		},
		{
			name:        "partial multi-line reply is still printed",
			script:      smtpmock.Script{EHLOReply: "250-email-submission.telnyx.com\n250-STARTTLS", CloseAfter: "EHLO"},
			wantVerdict: "FAILED at EHLO: connection closed by the server",
			wantStage:   smtpverify.StageEHLO,
			wantVerbs:   []string{"EHLO"},
			wantOutput:  []string{"[EHLO]      S: 250-email-submission.telnyx.com\n", "[EHLO]      S: 250-STARTTLS\n"},
		},
		{
			name:        "terminal escapes in a reply are neutralized",
			script:      smtpmock.Script{Banner: "554 x\x1b[2K\rVERIFIED: message UUID " + smtpmock.QueuedUUID + "\x1b[8m"},
			wantVerdict: `FAILED at CONNECT: 554 x\x1b[2K\x0dVERIFIED: message UUID ` + smtpmock.QueuedUUID + `\x1b[8m`,
			wantStage:   smtpverify.StageConnect,
			wantVerbs:   []string{"QUIT"},
		},
		{
			name:        "quota exhausted at DATA",
			script:      smtpmock.Script{DataEndReply: "452 4.5.3 Mail quota exceeded, try again later"},
			wantVerdict: "FAILED at DATA: 452 4.5.3 Mail quota exceeded, try again later",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantHint:    "quota",
		},
		{
			name:        "message too large",
			script:      smtpmock.Script{DataEndReply: "552 5.3.4 Message size exceeds fixed maximum message size"},
			wantVerdict: "FAILED at DATA: 552 5.3.4 Message size exceeds fixed maximum message size",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
		},
		{
			name:        "content rejected",
			script:      smtpmock.Script{DataEndReply: "554 5.6.0 Message content rejected"},
			wantVerdict: "FAILED at DATA: 554 5.6.0 Message content rejected",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
		},
		{
			name:        "From header domain suspended",
			script:      smtpmock.Script{DataEndReply: "550 5.7.1 Domain is suspended."},
			wantVerdict: "FAILED at DATA: 550 5.7.1 Domain is suspended.",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
		},
		{
			name:        "submission authorization failed at DATA",
			script:      smtpmock.Script{DataEndReply: "535 5.7.1 Authorization failed on message submission"},
			wantVerdict: "FAILED at DATA: 535 5.7.1 Authorization failed on message submission",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantHint:    "when the message was submitted",
		},
		{
			name:        "accepted without a message UUID",
			script:      smtpmock.Script{DataEndReply: "250 Ok"},
			wantVerdict: "FAILED at DATA: 250 Ok",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantHint:    "no message UUID",
		},
		{
			name:        "no reply to AUTH: stage timeout",
			script:      smtpmock.Script{HangOn: "AUTH"},
			mutate:      func(c *smtpverify.Config) { c.CommandTimeout = 150 * time.Millisecond },
			wantVerdict: "FAILED at AUTH: timed out after 150ms waiting for the server",
			wantStage:   smtpverify.StageAuth,
			wantVerbs:   verbsUpToAuth,
			wantHint:    "--timeout",
		},
		{
			name:        "no final DATA reply: data timeout",
			script:      smtpmock.Script{HangOn: "."},
			mutate:      func(c *smtpverify.Config) { c.DataTimeout = 150 * time.Millisecond },
			wantVerdict: "FAILED at DATA: timed out after 150ms waiting for the server",
			wantStage:   smtpverify.StageData,
			wantVerbs:   verbsUpToData,
		},
		{
			name:        "connection dropped at MAIL",
			script:      smtpmock.Script{CloseOn: "MAIL"},
			wantVerdict: "FAILED at MAIL FROM: connection closed by the server",
			wantStage:   smtpverify.StageMailFrom,
			wantVerbs:   verbsUpToMail,
		},
		{
			name:       "plaintext injected after STARTTLS: credentials never sent",
			script:     smtpmock.Script{GarbageAfterSTARTTLS: "250 injected"},
			wantPrefix: true,
			// Either the buffered-bytes check or the TLS handshake catches it.
			wantVerdict: "FAILED at STARTTLS: ",
			wantStage:   smtpverify.StageSTARTTLS,
			wantVerbs:   []string{"EHLO", "STARTTLS"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			srv := smtpmock.Start(t, tt.script)
			cfg := config(srv)
			if tt.mutate != nil {
				tt.mutate(&cfg)
			}
			res, out := run(t, cfg)

			if tt.wantPrefix {
				assert.True(t, strings.HasPrefix(res.Verdict(), tt.wantVerdict), "verdict %q", res.Verdict())
			} else {
				assert.Equal(t, tt.wantVerdict, res.Verdict())
			}
			assert.Equal(t, tt.wantStage, res.Stage)
			assert.Equal(t, strings.HasPrefix(tt.wantVerdict, "VERIFIED"), res.Verified)
			assert.Contains(t, out, "\n"+res.Verdict()+"\n", "verdict line printed")
			if tt.wantHint != "" {
				assert.Contains(t, res.Hint, tt.wantHint)
			}
			if !res.Verified {
				assert.NotEmpty(t, res.Hint, "every failure is actionable")
				assert.Contains(t, out, "hint: "+res.Hint)
			}
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assert.NotContains(t, out, "\x1b", "raw escape bytes reach the terminal")
			assert.NotContains(t, out, "\r", "raw CR bytes reach the terminal")

			// The server's view: the session stopped at the first terminal reply.
			srv.Wait()
			assert.Equal(t, tt.wantVerbs, srv.Verbs())
			// Credentials only ever cross TLS, and never reach the transcript.
			for _, c := range srv.Commands() {
				if strings.HasPrefix(c.Line, "AUTH") {
					assert.True(t, c.TLS, "AUTH sent without TLS")
				}
			}
			assert.NotContains(t, out, testKey)
			assert.NotContains(t, out, authPayload)
			// Nothing the client put on the wire, including after a failed
			// handshake, carries the key in the clear.
			raw := string(srv.Raw())
			assert.NotContains(t, raw, testKey)
			assert.NotContains(t, raw, authPayload)
		})
	}
}

var authPayload = base64.StdEncoding.EncodeToString([]byte("\x00apikey\x00" + testKey))

// wrongName keeps the trusted certificate but connects under a name the
// certificate doesn't cover.
func wrongName(c *smtpverify.Config) {
	addr := net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
	c.Host = "wrong.example"
	c.Dial = func(ctx context.Context, network, _ string) (net.Conn, error) {
		var d net.Dialer
		return d.DialContext(ctx, network, addr)
	}
}

func TestRunVerifiedTranscript(t *testing.T) {
	t.Parallel()
	srv := smtpmock.Start(t, smtpmock.Script{})
	cfg := config(srv)
	cfg.To = "Ops Team <ops@example.net>"
	res, out := run(t, cfg)
	require.True(t, res.Verified, out)
	assert.Equal(t, smtpmock.QueuedUUID, res.MessageUUID)

	// Every stage's server reply is printed verbatim, labelled by stage.
	for _, want := range []string{
		"[CONNECT]   S: " + smtpmock.Banner,
		"[EHLO]      C: EHLO localhost",
		"[EHLO]      S: 250-SIZE 25000000",
		"[EHLO]      S: 250 STARTTLS",
		"[STARTTLS]  S: 220 2.0.0 Ready to start TLS",
		"[STARTTLS]  -- TLS established: TLS 1.3",
		"[EHLO]      S: 250 AUTH PLAIN LOGIN",
		`[AUTH]      C: AUTH PLAIN <redacted: username "apikey", password = API key>`,
		"[AUTH]      S: " + smtpmock.ReplyAuthOK,
		"[MAIL FROM] C: MAIL FROM:<verify@example.com>",
		"[RCPT TO]   C: RCPT TO:<ops@example.net>",
		"[DATA]      S: 354 End data with <CR><LF>.<CR><LF>",
		"[DATA]      S: " + smtpmock.ReplyQueued,
		"[QUIT]      S: " + smtpmock.ReplyQuit,
		"VERIFIED: message UUID " + smtpmock.QueuedUUID,
	} {
		assert.Contains(t, out, want)
	}

	// AUTH PLAIN carries authzid "", authcid "apikey", password = the key.
	var auth string
	for _, c := range srv.Commands() {
		if strings.HasPrefix(c.Line, "AUTH ") {
			auth = c.Line
		}
	}
	fields := strings.Fields(auth)
	require.Len(t, fields, 3)
	assert.Equal(t, "PLAIN", fields[1])
	decoded, err := base64.StdEncoding.DecodeString(fields[2])
	require.NoError(t, err)
	assert.Equal(t, "\x00apikey\x00"+testKey, string(decoded))

	// The message: From/To match the envelope, CRLF line endings, a
	// Message-ID on the sender's domain.
	bodies := srv.Bodies()
	require.Len(t, bodies, 1)
	body := bodies[0]
	assert.Contains(t, body, "From: <verify@example.com>\r\n")
	assert.Contains(t, body, "To: <ops@example.net>\r\n")
	assert.Contains(t, body, "Date: Wed, 14 Oct 2026 12:00:00 +0000\r\n")
	assert.Regexp(t, `Message-ID: <telnyx-smtp-verify\.[0-9a-f]{16}@example\.com>\r\n`, body)
	assert.NotContains(t, strings.ReplaceAll(body, "\r\n", ""), "\n", "bare LF in message")
}

func TestRunToDefaultsToFrom(t *testing.T) {
	t.Parallel()
	srv := smtpmock.Start(t, smtpmock.Script{})
	res, _ := run(t, config(srv))
	require.True(t, res.Verified)
	assert.Contains(t, srv.Verbs(), "RCPT")
	for _, c := range srv.Commands() {
		if strings.HasPrefix(c.Line, "RCPT") {
			assert.Equal(t, "RCPT TO:<verify@example.com>", c.Line)
		}
	}
}

func TestRunDialFailures(t *testing.T) {
	t.Parallel()

	closed, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	closedPort := closed.Addr().(*net.TCPAddr).Port
	require.NoError(t, closed.Close())

	tests := []struct {
		name     string
		mutate   func(*smtpverify.Config)
		wantHint string
	}{
		{
			name: "host does not resolve",
			mutate: func(c *smtpverify.Config) {
				c.Host = "smtp.example.invalid"
				c.Dial = func(context.Context, string, string) (net.Conn, error) {
					return nil, &net.OpError{Op: "dial", Net: "tcp", Err: &net.DNSError{Err: "no such host", Name: "smtp.example.invalid", IsNotFound: true}}
				}
			},
			wantHint: "does not resolve",
		},
		{
			name:     "connection refused",
			mutate:   func(c *smtpverify.Config) { c.Host, c.Port = "127.0.0.1", closedPort },
			wantHint: "Could not connect",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cfg := smtpverify.Config{APIKey: testKey, From: "verify@example.com", CommandTimeout: 2 * time.Second}
			tt.mutate(&cfg)
			res, out := run(t, cfg)
			assert.False(t, res.Verified)
			assert.Equal(t, smtpverify.StageConnect, res.Stage)
			assert.Nil(t, res.Reply)
			assert.True(t, strings.HasPrefix(res.Verdict(), "FAILED at CONNECT: "), res.Verdict())
			assert.Contains(t, res.Hint, tt.wantHint)
			assert.Contains(t, out, res.Verdict())
		})
	}
}

func TestRunCanceled(t *testing.T) {
	t.Parallel()
	srv := smtpmock.Start(t, smtpmock.Script{HangOn: "AUTH"})
	cfg := config(srv)
	cfg.CommandTimeout = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	var out bytes.Buffer
	start := time.Now()
	res, err := smtpverify.Run(ctx, cfg, &out)
	require.NoError(t, err)
	assert.Less(t, time.Since(start), 5*time.Second, "cancellation interrupts a blocked read")
	assert.Equal(t, smtpverify.StageAuth, res.Stage)
	assert.True(t, strings.HasPrefix(res.Verdict(), "FAILED at AUTH: canceled"), res.Verdict())
}

func TestRunConfigValidation(t *testing.T) {
	t.Parallel()
	valid := smtpverify.Config{Host: "smtp.example.com", APIKey: testKey, From: "verify@example.com"}
	tests := []struct {
		name    string
		mutate  func(*smtpverify.Config)
		wantErr string
	}{
		{"missing host", func(c *smtpverify.Config) { c.Host = " " }, "host is required"},
		{"host with path", func(c *smtpverify.Config) { c.Host = "smtp.example.com/x" }, "invalid host"},
		{"port too large", func(c *smtpverify.Config) { c.Port = 70000 }, "invalid port"},
		{"negative port", func(c *smtpverify.Config) { c.Port = -1 }, "invalid port"},
		{"missing API key", func(c *smtpverify.Config) { c.APIKey = "" }, "API key is required"},
		{"API key with newline", func(c *smtpverify.Config) { c.APIKey = "KEY\r\nRCPT TO:<x@y>" }, "invalid characters"},
		{"missing from", func(c *smtpverify.Config) { c.From = "" }, "a from address is required"},
		{"from not an address", func(c *smtpverify.Config) { c.From = "not-an-address" }, "invalid from address"},
		{"from with CRLF injection", func(c *smtpverify.Config) { c.From = "a@example.com>\r\nRCPT TO:<b@example.com" }, "invalid from address"},
		{"to with CRLF injection", func(c *smtpverify.Config) { c.To = "a@example.com\r\nDATA" }, "invalid to address"},
		{"quoted local part injecting a MAIL parameter", func(c *smtpverify.Config) { c.From = `"a@b> SMTPUTF8"@example.com` }, "invalid from address"},
		{"quoted local part forging the null sender", func(c *smtpverify.Config) { c.From = `">"@example.com` }, "invalid from address"},
		{"tab in local part", func(c *smtpverify.Config) { c.To = "\"a\tb\"@example.com" }, "invalid to address"},
		{"non-ASCII address", func(c *smtpverify.Config) { c.To = "ü@exämple.com" }, "plain ASCII address"},
		{"leading dot", func(c *smtpverify.Config) { c.From = ".a@example.com" }, "invalid from address"},
		{"host with port", func(c *smtpverify.Config) { c.Host = "smtp.example.com:2525" }, "--port"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cfg := valid
			tt.mutate(&cfg)
			var out bytes.Buffer
			_, err := smtpverify.Run(context.Background(), cfg, &out)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
			assert.Empty(t, out.String(), "no session is attempted on invalid config")
		})
	}
}

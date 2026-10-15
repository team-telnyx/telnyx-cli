package smtpverify_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"net"
	"regexp"
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

// assertVerdictLast pins the output's ending: the hint (or, when verified,
// the portal note) and then the verdict, which is always the last line. No
// earlier line may read as a VERIFIED verdict.
func assertVerdictLast(t *testing.T, res smtpverify.Result, out string) {
	t.Helper()
	require.True(t, strings.HasSuffix(out, "\n"), "output ends with a newline")
	lines := strings.Split(strings.TrimSuffix(out, "\n"), "\n")
	require.GreaterOrEqual(t, len(lines), 2)
	assert.Equal(t, res.Verdict(), lines[len(lines)-1], "the verdict is the last line")
	if res.Verified {
		assert.Contains(t, lines[len(lines)-2], "Mission Control portal message log", "the portal note comes just before the verdict")
	} else {
		assert.Equal(t, "hint: "+res.Hint, lines[len(lines)-2], "the hint comes just before the verdict")
	}
	for _, l := range lines[:len(lines)-1] {
		assert.False(t, strings.HasPrefix(l, "VERIFIED"), "only the last line is a verdict: %q", l)
	}
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
			name:        "trace UUID in an acceptance is not the message UUID",
			script:      smtpmock.Script{DataEndReply: "250 2.0.0 accepted; trace " + smtpmock.QueuedUUID + "@mta"},
			wantVerdict: "FAILED at DATA: 250 2.0.0 accepted; trace " + smtpmock.QueuedUUID + "@mta",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantHint:    "no message UUID",
		},
		{
			name:        "acceptance phrase with a trailing trace UUID is not the contract shape",
			script:      smtpmock.Script{DataEndReply: smtpmock.ReplyQueued + " trace 11111111-1111-4111-8111-111111111111"},
			wantVerdict: "FAILED at DATA: " + smtpmock.ReplyQueued + " trace 11111111-1111-4111-8111-111111111111",
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantHint:    "no message UUID",
		},
		{
			name:        "multi-line acceptance is not the contract shape",
			script:      smtpmock.Script{DataEndReply: "250-accepted\n" + smtpmock.ReplyQueued},
			wantVerdict: "FAILED at DATA: 250-accepted | " + smtpmock.ReplyQueued,
			wantStage:   smtpverify.StageData,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantHint:    "no message UUID",
		},
		{
			name:        "acceptance cut off by a close is not VERIFIED",
			script:      smtpmock.Script{DataEndPartial: smtpmock.ReplyQueued},
			wantVerdict: `FAILED at DATA: incomplete reply "` + smtpmock.ReplyQueued + `" (no CRLF): connection closed by the server`,
			wantStage:   smtpverify.StageData,
			wantVerbs:   verbsUpToData,
			wantHint:    "cut off",
			wantOutput:  []string{"[DATA]      -- incomplete reply line (no CRLF): " + smtpmock.ReplyQueued + "\n"},
		},
		{
			name:        "acceptance cut off by a stall is not VERIFIED",
			script:      smtpmock.Script{DataEndPartial: smtpmock.ReplyQueued, HangOn: "."},
			mutate:      func(c *smtpverify.Config) { c.DataTimeout = 150 * time.Millisecond },
			wantVerdict: `FAILED at DATA: incomplete reply "` + smtpmock.ReplyQueued + `" (no CRLF): timed out after 150ms waiting for the server`,
			wantStage:   smtpverify.StageData,
			wantVerbs:   verbsUpToData,
			wantHint:    "cut off",
			wantOutput:  []string{"[DATA]      -- incomplete reply line (no CRLF): " + smtpmock.ReplyQueued + "\n"},
		},
		{
			name:        "AUTH payload echoed in a rejection is redacted",
			script:      smtpmock.Script{AuthReply: "535 5.7.8 Authentication credentials invalid: " + smtpmock.Echo},
			wantVerdict: "FAILED at AUTH: 535 5.7.8 Authentication credentials invalid: AUTH PLAIN <redacted: AUTH PLAIN credentials>",
			wantStage:   smtpverify.StageAuth,
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantHint:    `literal "apikey"`,
			wantOutput:  []string{"[AUTH]      S: 535 5.7.8 Authentication credentials invalid: AUTH PLAIN <redacted: AUTH PLAIN credentials>\n"},
		},
		{
			name:        "AUTH payload echoed as a malformed reply is redacted",
			script:      smtpmock.Script{AuthReply: smtpmock.Echo},
			wantVerdict: `FAILED at AUTH: malformed reply: "AUTH PLAIN <redacted: AUTH PLAIN credentials>"`,
			wantStage:   smtpverify.StageAuth,
			wantVerbs:   verbsUpToAuth,
		},
		{
			name:        "API key echoed in an acceptance is redacted",
			script:      smtpmock.Script{AuthReply: "235 2.7.0 Authentication successful for apikey:" + testKey},
			wantVerdict: "VERIFIED: message UUID " + smtpmock.QueuedUUID,
			wantStage:   smtpverify.StageQuit,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantOutput:  []string{"[AUTH]      S: 235 2.7.0 Authentication successful for apikey:<redacted: API key>\n"},
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
			assertVerdictLast(t, res, out)
			if tt.wantHint != "" {
				assert.Contains(t, res.Hint, tt.wantHint)
			}
			if !res.Verified {
				assert.NotEmpty(t, res.Hint, "every failure is actionable")
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
			// Nor does the Result, whatever the server echoed.
			assert.NotContains(t, res.Detail, testKey)
			assert.NotContains(t, res.Detail, authPayload)
			if res.Reply != nil {
				assert.NotContains(t, res.Reply.Text(), testKey)
				assert.NotContains(t, res.Reply.Text(), authPayload)
			}
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
	assertVerdictLast(t, res, out)

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

// A server can wrap an echoed credential across the lines of one reply, so
// that no line holds it whole ("535-<half1>" / "535 <half2>"). No piece may
// reach the transcript, the Result or the verdict, and joining the printed
// reply lines must not rebuild it.
func TestAstraSplitPayloadMustRedact(t *testing.T) {
	t.Parallel()
	const escKey = `KEY"esc\aped_0123456789`
	escaped := strconv.Quote(escKey)
	escaped = escaped[1 : len(escaped)-1] // KEY\"esc\\aped_0123456789
	split := func(s string, at ...int) []string {
		var pieces []string
		prev := 0
		for _, i := range append(at, len(s)) {
			pieces = append(pieces, s[prev:i])
			prev = i
		}
		return pieces
	}
	p, k, e := split(authPayload, 13), split(testKey, 9, 17), split(escaped, 4)
	tests := []struct {
		name        string
		key         string
		secret      string   // the form the server echoes
		pieces      []string // as the reply's lines carry it
		reply       string
		wantVerdict string
	}{
		{
			name:        "AUTH PLAIN payload split mid-token",
			key:         testKey,
			secret:      authPayload,
			pieces:      p,
			reply:       "535-5.7.8 Authentication credentials invalid: AUTH PLAIN " + p[0] + "\n535 " + p[1],
			wantVerdict: "FAILED at AUTH: 535-5.7.8 Authentication credentials invalid: AUTH PLAIN <redacted: AUTH PLAIN credentials>",
		},
		{
			name:        "raw key split across three lines",
			key:         testKey,
			secret:      testKey,
			pieces:      k,
			reply:       "535-5.7.8 Authentication credentials invalid for apikey:" + k[0] + "\n535-5.7.8 " + k[1] + "\n535 5.7.8 " + k[2] + " (check the key)",
			wantVerdict: "FAILED at AUTH: 535-5.7.8 Authentication credentials invalid for apikey:<redacted: API key> (check the key)",
		},
		{
			name:        "escaped key split inside an escape",
			key:         escKey,
			secret:      escaped,
			pieces:      e,
			reply:       `535-5.7.8 Authentication credentials invalid: {"password":"` + e[0] + "\n535 5.7.8 " + e[1] + `"}`,
			wantVerdict: `FAILED at AUTH: 535-5.7.8 Authentication credentials invalid: {"password":"<redacted: API key>"}`,
		},
	}
	framing := regexp.MustCompile(`^\d{3}[ -](?:\d\.\d{1,3}\.\d{1,3} )?`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.secret, strings.Join(tt.pieces, ""))
			srv := smtpmock.Start(t, smtpmock.Script{AuthReply: tt.reply})
			cfg := config(srv)
			cfg.APIKey = tt.key
			res, out := run(t, cfg)

			assert.Equal(t, tt.wantVerdict, res.Verdict())
			assert.Contains(t, out, "[AUTH]      S: "+strings.TrimPrefix(tt.wantVerdict, "FAILED at AUTH: ")+"\n")
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, with(verbsUpToAuth, "QUIT"), srv.Verbs())
			for _, c := range srv.Commands() {
				if strings.HasPrefix(c.Line, "AUTH") {
					assert.Equal(t, "AUTH PLAIN "+base64.StdEncoding.EncodeToString([]byte("\x00apikey\x00"+tt.key)), c.Line)
				}
			}

			require.NotNil(t, res.Reply)
			for name, s := range map[string]string{"output": out, "reply": res.Reply.Text(), "verdict": res.Verdict(), "detail": res.Detail} {
				for _, piece := range tt.pieces {
					assert.NotContains(t, s, piece, "a piece of the secret in the %s", name)
				}
			}
			var joined strings.Builder
			for _, l := range strings.Split(out, "\n") {
				if _, text, ok := strings.Cut(l, " S: "); ok {
					joined.WriteString(framing.ReplaceAllString(text, ""))
				}
			}
			assert.NotContains(t, joined.String(), tt.secret, "the printed reply lines rebuild the secret")
		})
	}
}

// The rest of a wrapped echo can also arrive in bytes that end the read: a
// line whose reply code changed mid-reply, a line with no reply code, or a
// line cut off before its CRLF, mid-secret or after it. Those bytes reach
// the transcript and the Result through the error, so they are masked
// together with the lines read before them. No piece may be printed or
// returned, and joining the printed reply lines with the bytes quoted by
// errors must not rebuild what the server sent of the secret.
func TestAstraSplitIntoErrorMustRedact(t *testing.T) {
	t.Parallel()
	k0, k1 := testKey[:9], testKey[9:] // k1 = "123456789abcdef": parses as reply code 123
	p0, p1 := authPayload[:13], authPayload[13:]
	tests := []struct {
		name        string
		script      smtpmock.Script
		sent        string   // what the server sent of the secret
		pieces      []string // as the lines and the bytes ending the read carry it
		wantVerdict string
		wantVerbs   []string
		wantOutput  []string
	}{
		{
			name:        "reply code changes mid-reply",
			script:      smtpmock.Script{AuthReply: "535-5.7.8 Authentication credentials invalid: AUTH PLAIN " + p0 + "\n550 " + p1},
			sent:        authPayload,
			pieces:      []string{p0, p1},
			wantVerdict: `FAILED at AUTH: malformed reply: code changed mid-reply in "535-5.7.8 Authentication credentials invalid: AUTH PLAIN <redacted: AUTH PLAIN credentials>"`,
			wantVerbs:   verbsUpToAuth,
		},
		{
			name:        "line with no reply code",
			script:      smtpmock.Script{AuthReply: "535-5.7.8 Authentication credentials invalid for apikey:" + k0 + "\n" + k1},
			sent:        testKey,
			pieces:      []string{k0, k1},
			wantVerdict: `FAILED at AUTH: malformed reply: reply code 123 out of range in "535-5.7.8 Authentication credentials invalid for apikey:<redacted: API key>"`,
			wantVerbs:   verbsUpToAuth,
		},
		{
			name:        "line cut off after the second half",
			script:      smtpmock.Script{AuthPartial: "535-5.7.8 Authentication credentials invalid for apikey:" + k0 + "\r\n535 5.7.8 " + k1 + " (check the k"},
			sent:        testKey,
			pieces:      []string{k0, k1},
			wantVerdict: `FAILED at AUTH: incomplete reply "535-5.7.8 Authentication credentials invalid for apikey:<redacted: API key> (check the k" (no CRLF): connection closed by the server`,
			wantVerbs:   verbsUpToAuth,
			wantOutput:  []string{"[AUTH]      -- incomplete reply line (no CRLF): 535-5.7.8 Authentication credentials invalid for apikey:<redacted: API key> (check the k\n"},
		},
		{
			name:        "line cut off inside the second half",
			script:      smtpmock.Script{AuthPartial: "535-5.7.8 Authentication credentials invalid: AUTH PLAIN " + p0 + "\r\n535 5.7.8 " + p1[:17]},
			sent:        p0 + p1[:17],
			pieces:      []string{p0, p1[:17]},
			wantVerdict: `FAILED at AUTH: incomplete reply "535-5.7.8 Authentication credentials invalid: AUTH PLAIN <redacted: AUTH PLAIN credentials>" (no CRLF): connection closed by the server`,
			wantVerbs:   verbsUpToAuth,
			wantOutput:  []string{"[AUTH]      -- incomplete reply line (no CRLF): 535-5.7.8 Authentication credentials invalid: AUTH PLAIN <redacted: AUTH PLAIN credentials>\n"},
		},
		{
			name: "QUIT reply code changes mid-reply",
			script: smtpmock.Script{
				AuthReply: "535 5.7.8 Authentication credentials invalid",
				QuitReply: "221-Bye apikey:" + k0 + "\n421 " + k1,
			},
			sent:        testKey,
			pieces:      []string{k0, k1},
			wantVerdict: "FAILED at AUTH: 535 5.7.8 Authentication credentials invalid",
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantOutput:  []string{`[QUIT]      -- no reply to QUIT: malformed reply: code changed mid-reply in "221-Bye apikey:<redacted: API key>"` + "\n"},
		},
	}
	framing := regexp.MustCompile(`^\d{3}[ -](?:\d\.\d{1,3}\.\d{1,3} )?`)
	quoted := regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.sent, strings.Join(tt.pieces, ""))
			srv := smtpmock.Start(t, tt.script)
			res, out := run(t, config(srv))

			assert.Equal(t, tt.wantVerdict, res.Verdict())
			assert.Equal(t, smtpverify.StageAuth, res.Stage)
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, tt.wantVerbs, srv.Verbs())

			texts := map[string]string{"output": out, "verdict": res.Verdict(), "detail": res.Detail}
			if res.Reply != nil {
				texts["reply"] = res.Reply.Text()
			}
			for name, s := range texts {
				for _, piece := range tt.pieces {
					assert.NotContains(t, s, piece, "a piece of the secret in the %s", name)
				}
			}
			// Every reply byte the output shows, in order: the reply lines,
			// the cut-off fragment, and the lines errors quote.
			var joined strings.Builder
			for _, l := range strings.Split(out, "\n") {
				if _, text, ok := strings.Cut(l, " S: "); ok {
					joined.WriteString(framing.ReplaceAllString(text, ""))
				}
				if _, text, ok := strings.Cut(l, " -- incomplete reply line (no CRLF): "); ok {
					joined.WriteString(framing.ReplaceAllString(text, ""))
				}
				for _, q := range quoted.FindAllString(l, -1) {
					text, err := strconv.Unquote(q)
					require.NoError(t, err)
					joined.WriteString(framing.ReplaceAllString(text, ""))
				}
			}
			assert.NotContains(t, joined.String(), tt.sent, "the printed reply lines and error texts rebuild the secret")
		})
	}
}

// longKey's AUTH PLAIN payload ends in "==": cut before the padding, it still
// decodes, unpadded, to every byte of the key.
const longKey = "KEY01892F3C4A5B6D7E8F9A0B1C2D3E4F_Vb7Nq2Wx9Rt4Ys6Uk8Jm1Lp3Hd5G"

func payloadOf(key string) string {
	return base64.StdEncoding.EncodeToString([]byte("\x00apikey\x00" + key))
}

var (
	replyFraming = regexp.MustCompile(`^\d{3}[ -](?:\d\.\d{1,3}\.\d{1,3} )?`)
	quotedText   = regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)
	base64Text   = regexp.MustCompile(`[A-Za-z0-9+/]+`)
)

// assertNothingRebuildsKey checks every surface a session shows: the output,
// the verdict, Detail, MessageUUID and Reply. None may hold a piece the
// server sent, nor smtpverify.MinFragment characters in a row of the key or
// its AUTH PLAIN payload, padded or not. Nor may the server text they print,
// joined in the order the output shows it (reply lines, cut-off fragments,
// bytes errors quote, the verdict's reply or message UUID), framing removed:
// not the secret the server sent, not MinFragment characters of a secret,
// and not MinFragment bytes of the key when its base64 runs, alone or all
// run together, are decoded at every alignment ("=" padding carries no
// bytes, so this covers the padded and unpadded payload alike).
func assertNothingRebuildsKey(t *testing.T, key, sent string, pieces []string, res smtpverify.Result, out string) {
	t.Helper()
	surfaces := map[string]string{"output": out, "verdict": res.Verdict(), "detail": res.Detail, "message UUID": res.MessageUUID}
	if res.Reply != nil {
		surfaces["reply"] = res.Reply.Text()
	}
	var joined strings.Builder
	add := func(text string) {
		joined.WriteString(replyFraming.ReplaceAllString(strings.ReplaceAll(text, `\x0d`, ""), ""))
	}
	for _, l := range strings.Split(out, "\n") {
		if _, text, ok := strings.Cut(l, " S: "); ok {
			add(text)
		}
		if _, text, ok := strings.Cut(l, " -- incomplete reply line (no CRLF): "); ok {
			add(text)
		}
		if text, ok := strings.CutPrefix(l, "FAILED at "); ok {
			_, text, _ = strings.Cut(text, ": ")
			for _, part := range strings.Split(text, " | ") {
				add(part)
			}
		}
		if text, ok := strings.CutPrefix(l, "VERIFIED: message UUID "); ok {
			add(text)
		}
		for _, q := range quotedText.FindAllString(l, -1) {
			if text, err := strconv.Unquote(q); err == nil {
				add(text)
			}
		}
	}
	surfaces["joined server text"] = joined.String()

	n := smtpverify.MinFragment
	payload := payloadOf(key)
	forms := []string{key, payload, strings.TrimRight(payload, "=")}
	for name, s := range surfaces {
		for _, piece := range pieces {
			assert.NotContains(t, s, piece, "a piece of the secret in the %s", name)
		}
		for _, form := range forms {
			for i := 0; i+n <= len(form); i++ {
				assert.NotContains(t, s, form[i:i+n], "%d characters of a secret in the %s", n, name)
			}
		}
	}
	assert.NotContains(t, joined.String(), sent, "the printed server text rebuilds the secret")
	runs := base64Text.FindAllString(joined.String(), -1)
	for _, run := range append(runs, strings.Join(runs, "")) {
		for off := 0; off < 4 && off < len(run); off++ {
			b64 := run[off:]
			if len(b64)%4 == 1 { // RawStdEncoding rejects a length of 4n+1
				b64 = b64[:len(b64)-1]
			}
			decoded, err := base64.RawStdEncoding.DecodeString(b64)
			require.NoError(t, err)
			for i := 0; i+n <= len(key); i++ {
				assert.NotContains(t, string(decoded), key[i:i+n], "decoding the printed base64 gives back the key")
			}
		}
	}
}

// Astra R5: redaction ran per reply, so a secret split across replies (or
// between a reply and the line ending the next read) printed whole in
// pieces. The session's server text is redacted together, the Result's
// repeat of a reply included, so no split leaves a piece to rejoin.
func TestIndependentAstraCrossReplyMustRedact(t *testing.T) {
	t.Parallel()
	p := authPayload
	k0, k1 := testKey[:12], testKey[12:] // k1 = "456789abcdef": parses as reply code 456
	l := []string{longKey[:13], longKey[13:26], longKey[26:39], longKey[39:51], longKey[51:]}
	tests := []struct {
		name        string
		key         string
		script      smtpmock.Script
		sent        string   // what the server sent of the secret
		pieces      []string // as the replies carry it
		wantVerdict string
		wantVerbs   []string
		wantOutput  []string
	}{
		{
			name:        "payload: AUTH half, QUIT half",
			key:         testKey,
			script:      smtpmock.Script{AuthReply: "535 " + p[:19], QuitReply: "221 " + p[19:]},
			sent:        p,
			pieces:      []string{p[:19], p[19:]},
			wantVerdict: "FAILED at AUTH: 535 <redacted: AUTH PLAIN credentials>",
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantOutput:  []string{"[QUIT]      S: 221 <redacted: AUTH PLAIN credentials>\n"},
		},
		{
			name:        "key: halves under 16 characters",
			key:         testKey,
			script:      smtpmock.Script{AuthReply: "535 5.7.8 Authentication credentials invalid for apikey:" + k0, QuitReply: "221 2.0.0 " + k1},
			sent:        testKey,
			pieces:      []string{k0, k1},
			wantVerdict: "FAILED at AUTH: 535 5.7.8 Authentication credentials invalid for apikey:<redacted: API key>",
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantOutput:  []string{"[QUIT]      S: 221 2.0.0 <redacted: API key>\n"},
		},
		{
			name:        "key: reply, then the line ending the next read",
			key:         testKey,
			script:      smtpmock.Script{AuthReply: "535 5.7.8 Authentication credentials invalid for apikey:" + k0, QuitReply: k1},
			sent:        testKey,
			pieces:      []string{k0, k1},
			wantVerdict: "FAILED at AUTH: 535 5.7.8 Authentication credentials invalid for apikey:<redacted: API key>",
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantOutput:  []string{`[QUIT]      -- no reply to QUIT: malformed reply: "<redacted: API key>"` + "\n"},
		},
		{
			name:        "key: second half first, the verdict repeating it after the first",
			key:         testKey,
			script:      smtpmock.Script{AuthReply: "535 5.7.8 " + k1, QuitReply: "221 " + k0},
			sent:        testKey,
			pieces:      []string{k0, k1},
			wantVerdict: "FAILED at AUTH: 535 5.7.8 <redacted: API key>",
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantOutput:  []string{"[AUTH]      S: 535 5.7.8 <redacted: API key>\n", "[QUIT]      S: 221 <redacted: API key>\n"},
		},
		{
			name: "key: across five replies, under 16 characters each",
			key:  longKey,
			script: smtpmock.Script{
				AuthReply:    "235 2.7.0 Authentication successful " + l[0],
				MailReply:    "250 2.1.0 " + l[1],
				RcptReply:    "250 2.1.5 " + l[2],
				DataReply:    "354 " + l[3],
				DataEndReply: "250 2.0.0 " + l[4],
			},
			sent:        longKey,
			pieces:      l,
			wantVerdict: "FAILED at DATA: 250 2.0.0 <redacted: API key>",
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantOutput: []string{
				"[AUTH]      S: 235 2.7.0 Authentication successful <redacted: API key>\n",
				"[MAIL FROM] S: 250 2.1.0 <redacted: API key>\n",
				"[RCPT TO]   S: 250 2.1.5 <redacted: API key>\n",
				"[DATA]      S: 354 <redacted: API key>\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.sent, strings.Join(tt.pieces, ""))
			srv := smtpmock.Start(t, tt.script)
			cfg := config(srv)
			cfg.APIKey = tt.key
			res, out := run(t, cfg)

			assert.Equal(t, tt.wantVerdict, res.Verdict())
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, tt.wantVerbs, srv.Verbs())
			assertNothingRebuildsKey(t, tt.key, tt.sent, tt.pieces, res, out)
		})
	}
}

// Astra R5: a cut-off line holding the AUTH PLAIN payload up to its "=="
// padding was left alone as a lone start of a secret, yet decodes unpadded to
// the whole key. Any 16 characters in a row of a secret are masked, on their
// own too.
func TestIndependentAstraUnpaddedCutMustRedact(t *testing.T) {
	t.Parallel()
	const prefix = "535 5.7.8 Authentication credentials invalid"
	unpadded := payloadOf(longKey)
	require.True(t, strings.HasSuffix(unpadded, "=="))
	unpadded = strings.TrimSuffix(unpadded, "==")
	tests := []struct {
		name   string
		key    string
		sent   string
		reply  string // cut off after sent
		masked string // what the output shows of reply
	}{
		{
			name:   "payload cut before its ==",
			key:    longKey,
			sent:   unpadded,
			reply:  prefix + ": AUTH PLAIN ",
			masked: prefix + ": AUTH PLAIN <redacted: AUTH PLAIN credentials>",
		},
		{
			name:   "payload cut before its =",
			key:    testKey,
			sent:   strings.TrimSuffix(authPayload, "="),
			reply:  prefix + ": AUTH PLAIN ",
			masked: prefix + ": AUTH PLAIN <redacted: AUTH PLAIN credentials>",
		},
		{
			name:   "key cut after 16 characters",
			key:    longKey,
			sent:   longKey[:16],
			reply:  prefix + " for apikey:",
			masked: prefix + " for apikey:<redacted: API key>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			srv := smtpmock.Start(t, smtpmock.Script{AuthPartial: tt.reply + tt.sent})
			cfg := config(srv)
			cfg.APIKey = tt.key
			res, out := run(t, cfg)

			assert.Equal(t, `FAILED at AUTH: incomplete reply "`+tt.masked+`" (no CRLF): connection closed by the server`, res.Verdict())
			assert.Contains(t, out, "[AUTH]      -- incomplete reply line (no CRLF): "+tt.masked+"\n")
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, verbsUpToAuth, srv.Verbs())
			assertNothingRebuildsKey(t, tt.key, tt.sent, []string{tt.sent}, res, out)
		})
	}
}

// Astra R5: a reply cut off right after a bare CR ("535-<half1>\r\n535
// <half2>\r") kept both halves, as the start of a secret had to run to the
// very end of the text and could not take in the CR. A CR (or blank) at the
// cut ends the start of a secret like the cut itself.
func TestIndependentAstraSplitCutAtCRMustRedact(t *testing.T) {
	t.Parallel()
	unpadded := strings.TrimSuffix(payloadOf(longKey), "==")
	u0, u1 := unpadded[:19], unpadded[19:]
	k0, k1 := testKey[:5], testKey[5:9]
	tests := []struct {
		name        string
		key         string
		partial     string
		sent        string
		pieces      []string
		wantVerdict string
		wantOutput  []string
	}{
		{
			name:        "unpadded payload in halves",
			key:         longKey,
			partial:     "535-" + u0 + "\r\n535 " + u1 + "\r",
			sent:        unpadded,
			pieces:      []string{u0, u1},
			wantVerdict: `FAILED at AUTH: incomplete reply "535 <redacted: AUTH PLAIN credentials>\r" (no CRLF): connection closed by the server`,
			wantOutput: []string{
				"[AUTH]      S: 535-<redacted: AUTH PLAIN credentials>\n",
				`[AUTH]      -- incomplete reply line (no CRLF): 535 <redacted: AUTH PLAIN credentials>\x0d` + "\n",
			},
		},
		{
			name:        "start of the key under 16 characters",
			key:         testKey,
			partial:     "535-5.7.8 Authentication credentials invalid for apikey:" + k0 + "\r\n535 5.7.8 " + k1 + "\r",
			sent:        k0 + k1,
			pieces:      []string{k0, k1},
			wantVerdict: `FAILED at AUTH: incomplete reply "535 5.7.8 <redacted: API key>\r" (no CRLF): connection closed by the server`,
			wantOutput: []string{
				"[AUTH]      S: 535-5.7.8 Authentication credentials invalid for apikey:<redacted: API key>\n",
				`[AUTH]      -- incomplete reply line (no CRLF): 535 5.7.8 <redacted: API key>\x0d` + "\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.sent, strings.Join(tt.pieces, ""))
			srv := smtpmock.Start(t, smtpmock.Script{AuthPartial: tt.partial})
			cfg := config(srv)
			cfg.APIKey = tt.key
			res, out := run(t, cfg)

			assert.Equal(t, tt.wantVerdict, res.Verdict())
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, verbsUpToAuth, srv.Verbs())
			assertNothingRebuildsKey(t, tt.key, tt.sent, tt.pieces, res, out)
		})
	}
}

// chunks splits s into pieces of n bytes, the last one perhaps shorter.
func chunks(s string, n int) []string {
	var pieces []string
	for i := 0; i < len(s); i += n {
		pieces = append(pieces, s[i:min(i+n, len(s))])
	}
	return pieces
}

// Astra R6: pieces under 16 characters, each with other text around it, gave
// back the whole key: halves around "221 Bye", halves and 12-character chunks
// in reverse order, five labeled replies, labeled 12-character chunks of the
// AUTH PLAIN payload. Any smtpverify.MinFragment characters in a row of a
// secret are masked wherever they are, whatever text is around them and in
// whatever order they come.
func TestAstraR6PiecesMustRedact(t *testing.T) {
	t.Parallel()
	k := testKey
	ck := longKey[:60]
	c := chunks(ck, 12)
	l := []string{longKey[:13], longKey[13:26], longKey[26:39], longKey[39:51], longKey[51:]}
	p := chunks(payloadOf(longKey), 12)
	nk := testKey[:3*smtpverify.MinFragment]
	n := chunks(nk, smtpverify.MinFragment)
	const verified = "VERIFIED: message UUID " + smtpmock.QueuedUUID
	tests := []struct {
		name        string
		key         string
		script      smtpmock.Script
		sent        string   // what the server sent of the secret, in order
		pieces      []string // as the replies carry it, in that order
		wantVerdict string
		wantVerbs   []string
		wantOutput  []string
	}{
		{
			name:        "key halves around other text",
			key:         testKey,
			script:      smtpmock.Script{AuthReply: "535 5.7.8 Authentication credentials invalid for apikey:" + k[:12], QuitReply: "221 Bye " + k[12:]},
			sent:        testKey,
			pieces:      []string{k[:12], k[12:]},
			wantVerdict: "FAILED at AUTH: 535 5.7.8 Authentication credentials invalid for apikey:<redacted: API key>",
			wantVerbs:   with(verbsUpToAuth, "QUIT"),
			wantOutput:  []string{"[QUIT]      S: 221 Bye <redacted: API key>\n"},
		},
		{
			name:        "key halves in reverse order",
			key:         testKey,
			script:      smtpmock.Script{AuthReply: smtpmock.ReplyAuthOK + " " + k[12:], MailReply: "250 2.1.0 Ok " + k[:12]},
			sent:        testKey,
			pieces:      []string{k[:12], k[12:]},
			wantVerdict: verified,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantOutput:  []string{"[AUTH]      S: " + smtpmock.ReplyAuthOK + " <redacted: API key>\n", "[MAIL FROM] S: 250 2.1.0 Ok <redacted: API key>\n"},
		},
		{
			name: "12-character chunks in reverse order",
			key:  ck,
			script: smtpmock.Script{
				AuthReply: smtpmock.ReplyAuthOK + " " + c[4],
				MailReply: "250 2.1.0 Ok " + c[3],
				RcptReply: "250 2.1.5 Ok " + c[2],
				DataReply: "354 go ahead " + c[1],
				QuitReply: "221 Bye " + c[0],
			},
			sent:        ck,
			pieces:      c,
			wantVerdict: verified,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantOutput: []string{
				"[AUTH]      S: " + smtpmock.ReplyAuthOK + " <redacted: API key>\n",
				"[DATA]      S: 354 go ahead <redacted: API key>\n",
				"[QUIT]      S: 221 Bye <redacted: API key>\n",
			},
		},
		{
			name: "key across five labeled replies",
			key:  longKey,
			script: smtpmock.Script{
				AuthReply:    smtpmock.ReplyAuthOK + " part 1/5: " + l[0],
				MailReply:    "250 2.1.0 Ok part 2/5: " + l[1],
				RcptReply:    "250 2.1.5 Ok part 3/5: " + l[2],
				DataReply:    "354 part 4/5: " + l[3],
				DataEndReply: "250 2.0.0 part 5/5: " + l[4],
			},
			sent:        longKey,
			pieces:      l,
			wantVerdict: "FAILED at DATA: 250 2.0.0 part 5/5: <redacted: API key>",
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantOutput:  []string{"[MAIL FROM] S: 250 2.1.0 Ok part 2/5: <redacted: API key>\n", "[DATA]      S: 354 part 4/5: <redacted: API key>\n"},
		},
		{
			name: "labeled 12-character chunks of the AUTH PLAIN payload",
			key:  longKey,
			script: smtpmock.Script{
				AuthReply: "235-2.7.0 chunk 1: " + p[0] + "\n235-2.7.0 chunk 2: " + p[1] + "\n" + smtpmock.ReplyAuthOK + " chunk 3: " + p[2],
				MailReply: "250-2.1.0 chunk 4: " + p[3] + "\n250 2.1.0 Ok chunk 5: " + p[4],
				RcptReply: "250-2.1.5 chunk 6: " + p[5] + "\n250 2.1.5 Ok chunk 7: " + p[6],
				DataReply: "354 chunk 8: " + p[7],
			},
			sent:        payloadOf(longKey),
			pieces:      p,
			wantVerdict: verified,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantOutput: []string{
				"[AUTH]      S: 235-2.7.0 chunk 1: <redacted: AUTH PLAIN credentials>\n",
				"[RCPT TO]   S: 250 2.1.5 Ok chunk 7: <redacted: AUTH PLAIN credentials>\n",
				"[DATA]      S: 354 chunk 8: <redacted: AUTH PLAIN credentials>\n",
			},
		},
		{
			name: "labeled chunks of MinFragment characters out of order",
			key:  testKey,
			script: smtpmock.Script{
				AuthReply: smtpmock.ReplyAuthOK + " [2] " + n[1],
				MailReply: "250 2.1.0 Ok [3] " + n[2],
				RcptReply: "250 2.1.5 Ok [1] " + n[0],
			},
			sent:        nk,
			pieces:      n,
			wantVerdict: verified,
			wantVerbs:   with(verbsUpToData, "QUIT"),
			wantOutput:  []string{"[RCPT TO]   S: 250 2.1.5 Ok [1] <redacted: API key>\n"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.sent, strings.Join(tt.pieces, ""))
			srv := smtpmock.Start(t, tt.script)
			cfg := config(srv)
			cfg.APIKey = tt.key
			res, out := run(t, cfg)

			assert.Equal(t, tt.wantVerdict, res.Verdict())
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, tt.wantVerbs, srv.Verbs())
			assertNothingRebuildsKey(t, tt.key, tt.sent, tt.pieces, res, out)
		})
	}
}

// r7Key is a key of the Telnyx format, 58 bytes: "KEY", 32 upper-case hex
// digits, "_", 22 letters and digits. Its hex opens with 8 digits.
const r7Key = "KEY123456789ABCDEF00B1C2D3E4F5A6B7C_Vb7Nq2Wx9Rt4Ys6Uk8Jm1L"

// pieces7 splits s into pieces of 7 bytes, the last one ending where s
// does: together they hold all of s.
func pieces7(s string) []string {
	var pieces []string
	for i := 0; i < len(s); i += 7 {
		pieces = append(pieces, s[min(i, len(s)-7):min(i+7, len(s))])
	}
	return pieces
}

// Astra R7: a key cut into labeled pieces of 7 characters, each in its own
// reply line, the last in the verdict and the Result's Reply, was printed
// whole in pieces, as were labeled 7-character pieces of its AUTH PLAIN
// payload: only 8 characters in a row of a secret were masked. Any 7 are.
func TestAstraR7SevenCharacterPiecesMustRedact(t *testing.T) {
	t.Parallel()
	k := pieces7(r7Key)
	p := pieces7(payloadOf(r7Key))
	require.Len(t, k, 9)
	require.Len(t, p, 13)
	tests := []struct {
		name        string
		script      smtpmock.Script
		sent        string
		pieces      []string
		wantVerdict string
		wantOutput  []string
	}{
		{
			name: "labeled 7-character pieces of the key",
			script: smtpmock.Script{
				AuthReply:    "235-2.7.0 part 1/9: " + k[0] + "\n235-2.7.0 part 2/9: " + k[1] + "\n235 2.7.0 part 3/9: " + k[2],
				MailReply:    "250-2.1.0 part 4/9: " + k[3] + "\n250 2.1.0 Ok part 5/9: " + k[4],
				RcptReply:    "250-2.1.5 part 6/9: " + k[5] + "\n250 2.1.5 Ok part 7/9: " + k[6],
				DataReply:    "354 part 8/9: " + k[7],
				DataEndReply: "250 2.0.0 part 9/9: " + k[8],
			},
			sent:        r7Key,
			pieces:      k,
			wantVerdict: "FAILED at DATA: 250 2.0.0 part 9/9: <redacted: API key>",
			wantOutput: []string{
				"[AUTH]      S: 235-2.7.0 part 1/9: <redacted: API key>\n",
				"[MAIL FROM] S: 250 2.1.0 Ok part 5/9: <redacted: API key>\n",
				"[DATA]      S: 354 part 8/9: <redacted: API key>\n",
			},
		},
		{
			name: "labeled 7-character pieces of the AUTH PLAIN payload",
			script: smtpmock.Script{
				AuthReply:    "235-2.7.0 chunk 1: " + p[0] + "\n235-2.7.0 chunk 2: " + p[1] + "\n235-2.7.0 chunk 3: " + p[2] + "\n235 2.7.0 chunk 4: " + p[3],
				MailReply:    "250-2.1.0 chunk 5: " + p[4] + "\n250-2.1.0 chunk 6: " + p[5] + "\n250 2.1.0 chunk 7: " + p[6],
				RcptReply:    "250-2.1.5 chunk 8: " + p[7] + "\n250-2.1.5 chunk 9: " + p[8] + "\n250 2.1.5 chunk 10: " + p[9],
				DataReply:    "354 chunk 11: " + p[10],
				DataEndReply: "250-2.0.0 chunk 12: " + p[11] + "\n250 2.0.0 chunk 13: " + p[12],
			},
			sent:        payloadOf(r7Key),
			pieces:      p,
			wantVerdict: "FAILED at DATA: 250-2.0.0 chunk 12: <redacted: AUTH PLAIN credentials> | 250 2.0.0 chunk 13: <redacted: AUTH PLAIN credentials>",
			wantOutput: []string{
				"[AUTH]      S: 235-2.7.0 chunk 1: <redacted: AUTH PLAIN credentials>\n",
				"[RCPT TO]   S: 250 2.1.5 chunk 10: <redacted: AUTH PLAIN credentials>\n",
				"[DATA]      S: 354 chunk 11: <redacted: AUTH PLAIN credentials>\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			srv := smtpmock.Start(t, tt.script)
			cfg := config(srv)
			cfg.APIKey = r7Key
			res, out := run(t, cfg)

			assert.Equal(t, tt.wantVerdict, res.Verdict())
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, with(verbsUpToData, "QUIT"), srv.Verbs())
			assertNothingRebuildsKey(t, r7Key, tt.sent, tt.pieces, res, out)
		})
	}
}

// Astra R7: the message UUID that VERIFIED prints came from the acceptance
// reply, but not through redaction, so a server could print key material in
// it: digits of the key, its hex letters lower-cased, or the rest of a run
// of the key that the reply before it ends in. The UUID is masked with the
// session's server text, as the server sent it and as VERIFIED prints it,
// lower-cased; a UUID that holds no key material is printed as is.
func TestAstraR7VerifiedUUIDMustRedact(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		script     smtpmock.Script
		pieces     []string
		wantUUID   string
		wantOutput []string
	}{
		{
			name:       "UUID opening with 8 digits of the key",
			script:     smtpmock.Script{DataEndReply: "250 2.0.0 Ok: queued as 12345678-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
			pieces:     []string{"12345678"},
			wantUUID:   "<redacted: API key>-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			wantOutput: []string{"[DATA]      S: 250 2.0.0 Ok: queued as <redacted: API key>-3a0f-4c3e-9d6a-6f1f2b8e4c21\n"},
		},
		{
			name:       "UUID of hex letters and digits of the key, lower-cased",
			script:     smtpmock.Script{DataEndReply: "250 2.0.0 Ok: queued as 89ABCDEF-3A0F-4C3E-9D6A-0B1C2D3E4F5A"},
			pieces:     []string{"89abcdef", "0b1c2d3e4f5a"},
			wantUUID:   "<redacted: API key>-3a0f-4c3e-9d6a-<redacted: API key>",
			wantOutput: []string{"[DATA]      S: 250 2.0.0 Ok: queued as <redacted: API key>-3A0F-4C3E-9D6A-<redacted: API key>\n"},
		},
		{
			name: "run of the key from the QUIT reply into the UUID",
			script: smtpmock.Script{
				DataEndReply: "250 2.0.0 Ok: queued as 45678900-3a0f-4c3e-9d6a-6f1f2b8e4c21",
				QuitReply:    "221 2.0.0 Bye " + r7Key[:6],
			},
			pieces:     []string{r7Key[:6]},
			wantUUID:   "<redacted: API key>00-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			wantOutput: []string{"[QUIT]      S: 221 2.0.0 Bye <redacted: API key>\n"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			srv := smtpmock.Start(t, tt.script)
			cfg := config(srv)
			cfg.APIKey = r7Key
			res, out := run(t, cfg)

			assert.True(t, res.Verified)
			assert.Equal(t, tt.wantUUID, res.MessageUUID)
			assert.Equal(t, "VERIFIED: message UUID "+tt.wantUUID, res.Verdict())
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, with(verbsUpToData, "QUIT"), srv.Verbs())
			assertNothingRebuildsKey(t, r7Key, r7Key, nil, res, out)
			// Nothing the server sent in the key's own case shows lower-cased.
			for name, s := range map[string]string{"message UUID": res.MessageUUID, "verdict": res.Verdict(), "output": out} {
				s = strings.ToLower(s)
				for _, piece := range tt.pieces {
					assert.NotContains(t, s, strings.ToLower(piece), "a piece of the key in the %s", name)
				}
				for i := 0; i+7 <= len(r7Key); i++ {
					assert.NotContains(t, s, strings.ToLower(r7Key[i:i+7]), "7 characters of the key, in either case, in the %s", name)
				}
			}
		})
	}
}

// Astra R7: pieces of a key too short to mask alone, split by five line
// breaks each opening a line with framing, printed 9 characters of the key:
// a run split by gaps was masked only across up to 4 line breaks. Up to 8
// are taken in now, framing and blanks too, still a constant bound.
func TestAstraR7GapOfFiveLineBreaksMustRedact(t *testing.T) {
	t.Parallel()
	framed := func(framing string, n int) string { return strings.Repeat("\n"+framing, n) }
	tests := []struct {
		name       string
		authReply  string
		pieces     []string
		wantOutput []string
	}{
		{
			name:      "3-character pieces across 5 line breaks of framing and enhanced codes",
			authReply: "235-2.7.0 user " + r7Key[:3] + framed("235-2.7.0 ", 4) + "\n235-2.7.0 " + r7Key[3:6] + framed("235-2.7.0 ", 4) + "\n235 2.7.0 " + r7Key[6:9] + " Authentication successful",
			pieces:    []string{r7Key[:3], r7Key[3:6], r7Key[6:9]},
			wantOutput: []string{
				"[AUTH]      S: 235-2.7.0 user <redacted: API key>\n",
				"[AUTH]      S: 235-2.7.0 <redacted: API key>\n",
				"[AUTH]      S: 235 2.7.0 <redacted: API key> Authentication successful\n",
			},
		},
		{
			name:      "6-character pieces across 5 line breaks of framing alone",
			authReply: "235-2.7.0 user " + r7Key[:6] + framed("235-", 4) + "\n235 " + r7Key[6:12] + " Authentication successful",
			pieces:    []string{r7Key[:6], r7Key[6:12]},
			wantOutput: []string{
				"[AUTH]      S: 235-2.7.0 user <redacted: API key>\n",
				"[AUTH]      S: 235 <redacted: API key> Authentication successful\n",
			},
		},
		{
			name:      "6-character pieces across 8 line breaks of framing and enhanced codes",
			authReply: "235-2.7.0 user " + r7Key[:6] + framed("235-2.7.0 ", 7) + "\n235 2.7.0 " + r7Key[6:12] + " Authentication successful",
			pieces:    []string{r7Key[:6], r7Key[6:12]},
			wantOutput: []string{
				"[AUTH]      S: 235-2.7.0 user <redacted: API key>\n",
				"[AUTH]      S: 235 2.7.0 <redacted: API key> Authentication successful\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			srv := smtpmock.Start(t, smtpmock.Script{AuthReply: tt.authReply})
			cfg := config(srv)
			cfg.APIKey = r7Key
			res, out := run(t, cfg)

			assert.True(t, res.Verified, res.Verdict())
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, with(verbsUpToData, "QUIT"), srv.Verbs())
			assertNothingRebuildsKey(t, r7Key, strings.Join(tt.pieces, ""), nil, res, out)
		})
	}
}

// r8Key is Astra R8's key, of the Telnyx format, 58 bytes: its last 22
// characters hold a run of lower-case hex letters, "abcdef".
const r8Key = "KEY0123456789ABCDEF0123456789ABCDEF_xyabcdefuvwxyz01234567"

// assertNoPairRebuildsKey checks that no reply line a session prints,
// framing removed, joined to a showing of the message UUID, before it or
// after it, holds smtpverify.MinFragment characters in a row of key. The
// UUID shows in the acceptance line as sent, and so lower-cased too, as
// VERIFIED shows it, and in VERIFIED.
func assertNoPairRebuildsKey(t *testing.T, key string, res smtpverify.Result, out string) {
	t.Helper()
	var texts []string
	uuids := []string{res.MessageUUID}
	for _, l := range strings.Split(out, "\n") {
		if _, text, ok := strings.Cut(l, " S: "); ok {
			texts = append(texts, replyFraming.ReplaceAllString(text, ""))
			if _, uuid, ok := strings.Cut(text, " queued as "); ok {
				uuids = append(uuids, uuid, strings.ToLower(uuid))
			}
		}
	}
	n := smtpverify.MinFragment
	for _, a := range texts {
		for _, u := range uuids {
			for k := 0; k+n <= len(key); k++ {
				assert.NotContains(t, a+u, key[k:k+n], "%d characters of the key in %q joined to %q", n, a, u)
				assert.NotContains(t, u+a, key[k:k+n], "%d characters of the key in %q joined to %q", n, u, a)
			}
		}
	}
}

// Astra R8: a run of the key from the QUIT reply into the message UUID was
// printed whole when the server sent the UUID in upper case: VERIFIED shows
// it lower-cased, so the QUIT reply's "xya" and the UUID's "bcdef" rejoined
// as "xyabcdef", but runs split by a gap were looked for only in the text as
// the server sent it, "xya" then "BCDEF". They are looked for in the text as
// it is shown too, the UUID lower-cased wherever it is, in VERIFIED and in
// the acceptance line, which comes before the QUIT reply, so that the UUID
// may end a run or start one, whichever case it shows in. The UUID is masked
// the same way wherever it shows: one showing left unmasked would give away
// what another masked.
func TestAstraR8LowerCasedUUIDMustRedact(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		uuid       string // as the server sends it
		quit       string
		wantUUID   string
		wantOutput []string
	}{
		{
			name:     "Astra's probe: QUIT reply ends xya, upper-case UUID opens BCDEF",
			uuid:     "BCDEF999-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			quit:     "221 Bye xya",
			wantUUID: "<redacted: API key>999-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as <redacted: API key>999-3a0f-4c3e-9d6a-6f1f2b8e4c21\n",
				"[QUIT]      S: 221 Bye <redacted: API key>\n",
			},
		},
		{
			name:     "QUIT reply ends xy, upper-case UUID opens ABCDEF",
			uuid:     "ABCDEF99-3A0F-4C3E-9D6A-6F1F2B8E4C21",
			quit:     "221 Bye xy",
			wantUUID: "<redacted: API key>99-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as <redacted: API key>99-3A0F-4C3E-9D6A-6F1F2B8E4C21\n",
				"[QUIT]      S: 221 Bye <redacted: API key>\n",
			},
		},
		{
			name:     "QUIT reply ends xyabcd, upper-case UUID opens EF",
			uuid:     "EF999999-3A0F-4C3E-9D6A-6F1F2B8E4C21",
			quit:     "221 Bye xyabcd",
			wantUUID: "<redacted: API key>999999-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as <redacted: API key>999999-3A0F-4C3E-9D6A-6F1F2B8E4C21\n",
				"[QUIT]      S: 221 Bye <redacted: API key>\n",
			},
		},
		{
			name:     "QUIT reply ends xya, mixed-case UUID opens bCdEf",
			uuid:     "bCdEf999-3A0F-4C3E-9D6A-6F1F2B8E4C21",
			quit:     "221 Bye xya",
			wantUUID: "<redacted: API key>999-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as <redacted: API key>999-3A0F-4C3E-9D6A-6F1F2B8E4C21\n",
				"[QUIT]      S: 221 Bye <redacted: API key>\n",
			},
		},
		{
			name:     "QUIT reply ends xya, lower-case UUID opens bcdef",
			uuid:     "bcdef999-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			quit:     "221 Bye xya",
			wantUUID: "<redacted: API key>999-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as <redacted: API key>999-3a0f-4c3e-9d6a-6f1f2b8e4c21\n",
				"[QUIT]      S: 221 Bye <redacted: API key>\n",
			},
		},
		{
			name:     "upper-case UUID ends ABC, QUIT reply opens defuvw",
			uuid:     "3A0F9999-3A0F-4C3E-9D6A-000000000ABC",
			quit:     "221 2.0.0 defuvw",
			wantUUID: "3a0f9999-3a0f-4c3e-9d6a-000000000<redacted: API key>",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as 3A0F9999-3A0F-4C3E-9D6A-000000000<redacted: API key>\n",
				"[QUIT]      S: 221 2.0.0 <redacted: API key>\n",
			},
		},
		{
			name:     "lower-case UUID ends abc, QUIT reply opens defuvw",
			uuid:     "3a0f9999-3a0f-4c3e-9d6a-000000000abc",
			quit:     "221 2.0.0 defuvw",
			wantUUID: "3a0f9999-3a0f-4c3e-9d6a-000000000<redacted: API key>",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as 3a0f9999-3a0f-4c3e-9d6a-000000000<redacted: API key>\n",
				"[QUIT]      S: 221 2.0.0 <redacted: API key>\n",
			},
		},
		{
			name:     "key's own upper case: QUIT reply ends 9AB, UUID opens CDEF01",
			uuid:     "CDEF0199-3A0F-4C3E-9D6A-6F1F2B8E4C21",
			quit:     "221 Bye 9AB",
			wantUUID: "<redacted: API key>99-3a0f-4c3e-9d6a-6f1f2b8e4c21",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as <redacted: API key>99-3A0F-4C3E-9D6A-6F1F2B8E4C21\n",
				"[QUIT]      S: 221 Bye <redacted: API key>\n",
			},
		},
		{
			name:     "key's own upper case: UUID ends 789ABC, QUIT reply opens DEF012",
			uuid:     "5A0F9999-3A0F-4C3E-9D6A-000000789ABC",
			quit:     "221 2.0.0 DEF012",
			wantUUID: "5a0f9999-3a0f-4c3e-9d6a-000000<redacted: API key>",
			wantOutput: []string{
				"[DATA]      S: 250 2.0.0 Ok: queued as 5A0F9999-3A0F-4C3E-9D6A-000000<redacted: API key>\n",
				"[QUIT]      S: 221 2.0.0 <redacted: API key>\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			srv := smtpmock.Start(t, smtpmock.Script{DataEndReply: "250 2.0.0 Ok: queued as " + tt.uuid, QuitReply: tt.quit})
			cfg := config(srv)
			cfg.APIKey = r8Key
			res, out := run(t, cfg)

			assert.True(t, res.Verified, res.Verdict())
			assert.Equal(t, tt.wantUUID, res.MessageUUID)
			assert.Equal(t, "VERIFIED: message UUID "+tt.wantUUID, res.Verdict())
			for _, want := range tt.wantOutput {
				assert.Contains(t, out, want)
			}
			assertVerdictLast(t, res, out)
			srv.Wait()
			assert.Equal(t, with(verbsUpToData, "QUIT"), srv.Verbs())
			assertNothingRebuildsKey(t, r8Key, r8Key, nil, res, out)
			assertNoPairRebuildsKey(t, r8Key, res, out)
		})
	}
}

// The false-positive control for the rules above: a cut-off reply of UUID
// tails, an address and base64 MIME text, ending in what starts the AUTH
// PLAIN payload ("A"), CR or not, has nothing in common with the session's
// secrets but chance single characters, and is printed as sent.
func TestCutOffNoiseIsNotRedacted(t *testing.T) {
	t.Parallel()
	const noise = "250 2.0.0 Ok: queued as 3A0F-4C3E-9D6A-6F1F2B8E4C21 for <ops@example.net>" +
		" part Q29udGVudC1UeXBlOiB0ZXh0L3BsYWluOyBjaGFyc2V0PXVzLWFzY2lp ref 0B6C5B5E-3A"
	for _, end := range []string{"", "\r"} {
		t.Run(strconv.Quote(end), func(t *testing.T) {
			t.Parallel()
			srv := smtpmock.Start(t, smtpmock.Script{
				AuthReply:      "235 2.7.0 Authentication successful (QWxhZGRpbjpvcGVuIHNlc2FtZQ==)",
				DataEndPartial: noise + end,
			})
			res, out := run(t, config(srv))

			assert.Equal(t, "FAILED at DATA: incomplete reply "+strconv.Quote(noise+end)+" (no CRLF): connection closed by the server", res.Verdict())
			assert.Contains(t, out, "[AUTH]      S: 235 2.7.0 Authentication successful (QWxhZGRpbjpvcGVuIHNlc2FtZQ==)\n")
			assert.Contains(t, out, "[DATA]      -- incomplete reply line (no CRLF): "+printableCR(noise+end)+"\n")
			assert.NotContains(t, out, "<redacted: API key>")
			assert.NotContains(t, out, "<redacted: AUTH PLAIN credentials>")
			assertVerdictLast(t, res, out)
		})
	}
}

func printableCR(s string) string { return strings.ReplaceAll(s, "\r", `\x0d`) }

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
			assertVerdictLast(t, res, out)
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
	assertVerdictLast(t, res, out.String())
}

// Canceling the context, as Ctrl-C does, ends the wait for the final DATA
// reply at once; the transcript so far is still written, masked like any
// other, then the hint and an "interrupted" verdict.
func TestRunInterruptedFlushesTranscript(t *testing.T) {
	t.Parallel()
	srv := smtpmock.Start(t, smtpmock.Script{
		AuthReply: smtpmock.ReplyAuthOK + " for apikey:" + testKey[:12],
		MailReply: "250 2.1.0 Ok " + testKey[12:],
		HangOn:    ".",
	})
	cfg := config(srv)
	cfg.DataTimeout = 30 * time.Second
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for len(srv.Bodies()) == 0 && ctx.Err() == nil {
			time.Sleep(5 * time.Millisecond)
		}
		cancel() // the message is sent: the session waits for the final reply
	}()
	var out bytes.Buffer
	start := time.Now()
	res, err := smtpverify.Run(ctx, cfg, &out)
	require.NoError(t, err)
	assert.Less(t, time.Since(start), 10*time.Second, "the interrupt ends the wait")

	assert.Equal(t, "FAILED at DATA: interrupted", res.Verdict())
	assert.Equal(t, "The command was interrupted.", res.Hint)
	for _, want := range []string{
		"[AUTH]      S: " + smtpmock.ReplyAuthOK + " for apikey:<redacted: API key>\n",
		"[MAIL FROM] S: 250 2.1.0 Ok <redacted: API key>\n",
		"[DATA]      C: .\n",
	} {
		assert.Contains(t, out.String(), want)
	}
	assertVerdictLast(t, res, out.String())
	assertNothingRebuildsKey(t, testKey, testKey, []string{testKey[:12], testKey[12:]}, res, out.String())
}

// Astra R6: a banner the parser accepts, of lines full of framing ("250\r"
// over and over), made redaction quadratic: 16 KB took 10 s and 7.7 GB, far
// past the command timeout. Redaction is linear now (see
// TestSessionRedactionIsLinear), so such a session ends as quickly as any.
func TestRunHostileBannerIsBounded(t *testing.T) {
	// Not parallel: timed.
	for name, body := range map[string]string{
		"framing over and over":        strings.Repeat("250\r", 510),
		"framing that opens each line": "\r" + strings.Repeat("250\r", 510),
	} {
		t.Run(name, func(t *testing.T) {
			var lines []string
			for size := 0; size < 16<<10; size += 4 + len(body) {
				lines = append(lines, "220-"+body)
			}
			lines[len(lines)-1] = "220 " + body
			srv := smtpmock.Start(t, smtpmock.Script{Banner: strings.Join(lines, "\n")})
			cfg := config(srv)
			cfg.CommandTimeout = time.Second
			start := time.Now()
			res, out := run(t, cfg)
			assert.Less(t, time.Since(start), 500*time.Millisecond)
			assert.True(t, res.Verified, res.Verdict())
			assert.Contains(t, out, "[CONNECT]   S: 220-"+printableCR(body)+"\n")
			assertVerdictLast(t, res, out)
		})
	}
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

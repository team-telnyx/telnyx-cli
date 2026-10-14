package smtpverify

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadReply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		wire         string
		wantCode     int
		wantEnhanced string
		wantLines    []string
		wantErr      error
	}{
		{
			name:         "single line with enhanced code",
			wire:         "535 5.7.8 Authentication credentials invalid\r\n",
			wantCode:     535,
			wantEnhanced: "5.7.8",
			wantLines:    []string{"535 5.7.8 Authentication credentials invalid"},
		},
		{
			name:         "two-digit enhanced detail",
			wire:         "538 5.7.11 Encryption required\r\n",
			wantCode:     538,
			wantEnhanced: "5.7.11",
			wantLines:    []string{"538 5.7.11 Encryption required"},
		},
		{
			name:      "no enhanced code",
			wire:      "221 Bye\r\n",
			wantCode:  221,
			wantLines: []string{"221 Bye"},
		},
		{
			name:      "bare code",
			wire:      "250\r\n",
			wantCode:  250,
			wantLines: []string{"250"},
		},
		{
			name:      "enhanced class must match reply class",
			wire:      "250 5.0.0 odd\r\n",
			wantCode:  250,
			wantLines: []string{"250 5.0.0 odd"},
		},
		{
			name:     "multi-line EHLO",
			wire:     "250-email-submission.telnyx.com\r\n250-SIZE 25000000\r\n250 STARTTLS\r\n",
			wantCode: 250,
			wantLines: []string{
				"250-email-submission.telnyx.com",
				"250-SIZE 25000000",
				"250 STARTTLS",
			},
		},
		{
			name:         "bare LF tolerated",
			wire:         "250 2.0.0 Ok: queued as 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21\n",
			wantCode:     250,
			wantEnhanced: "2.0.0",
			wantLines:    []string{"250 2.0.0 Ok: queued as 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
		},
		{name: "too short", wire: "25\r\n", wantErr: errMalformedReply},
		{name: "non-digit code", wire: "2x0 Ok\r\n", wantErr: errMalformedReply},
		{name: "code out of range", wire: "199 Ok\r\n", wantErr: errMalformedReply},
		{name: "bad separator", wire: "250_Ok\r\n", wantErr: errMalformedReply},
		{name: "code changes mid-reply", wire: "250-a\r\n251 b\r\n", wantErr: errMalformedReply},
		{name: "line too long", wire: "250 " + strings.Repeat("x", maxReplyLineBytes) + "\r\n", wantErr: errMalformedReply},
		{name: "EOF before reply", wire: "", wantErr: io.EOF},
		{name: "EOF mid multi-line reply", wire: "250-a\r\n", wantErr: io.ErrUnexpectedEOF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := readReply(bufio.NewReader(strings.NewReader(tt.wire)))
			if tt.wantErr != nil {
				require.Error(t, err)
				assert.True(t, errors.Is(err, tt.wantErr), "got %v, want %v", err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantCode, got.Code)
			assert.Equal(t, tt.wantEnhanced, got.Enhanced)
			assert.Equal(t, tt.wantLines, got.Lines)
		})
	}
}

func TestReadReplyConsumesExactlyOneReply(t *testing.T) {
	t.Parallel()
	r := bufio.NewReader(strings.NewReader("250-a\r\n250 b\r\n221 Bye\r\n"))
	first, err := readReply(r)
	require.NoError(t, err)
	assert.Equal(t, []string{"250-a", "250 b"}, first.Lines)
	second, err := readReply(r)
	require.NoError(t, err)
	assert.Equal(t, []string{"221 Bye"}, second.Lines)
}

func TestMessageUUID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lines []string
		want  string
	}{
		{"contract shape", []string{"250 2.0.0 Ok: queued as 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"}, "0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
		{"upper case normalized", []string{"250 2.0.0 Ok: queued as 0B6C5B5E-3A0F-4C3E-9D6A-6F1F2B8E4C21"}, "0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
		{"no uuid", []string{"250 Ok"}, ""},
		{"not a uuid", []string{"250 2.0.0 Ok: queued as 1234ABCD"}, ""},
		{"uuid on a later line", []string{"250-accepted", "250 id 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"}, "0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
		{"queued-as wins over an earlier uuid", []string{"250 trace 11111111-1111-4111-8111-111111111111 queued as 22222222-2222-4222-8222-222222222222"}, "22222222-2222-4222-8222-222222222222"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, messageUUID(Reply{Code: 250, Lines: tt.lines}))
		})
	}
}

func TestPrintable(t *testing.T) {
	t.Parallel()
	tests := []struct{ in, want string }{
		{"250 2.0.0 Ok", "250 2.0.0 Ok"},
		{"220 x\x1b[2K\rfake", `220 x\x1b[2K\x0dfake`},
		{"tab\tkept", "tab\tkept"},
		{"del\x7f", `del\x7f`},
		{"c1 \u009b", `c1 \u009b`},
		{"bad utf8 \xff", `bad utf8 \xff`},
		{"unicode ok ü", "unicode ok ü"},
		{"bidi \u202eoverride", `bidi \u202eoverride`},
		{"zero\u200bwidth", `zero\u200bwidth`},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, printable(tt.in), "%q", tt.in)
	}
}

func TestCapabilities(t *testing.T) {
	t.Parallel()
	caps := capabilities(Reply{Code: 250, Lines: []string{
		"250-email-submission.telnyx.com",
		"250-SIZE 25000000",
		"250-starttls",
		"250 AUTH PLAIN LOGIN",
	}})
	assert.Equal(t, "25000000", caps["SIZE"])
	_, hasTLS := caps["STARTTLS"]
	assert.True(t, hasTLS, "keywords are case-insensitive")
	_, hasGreeting := caps["EMAIL-SUBMISSION.TELNYX.COM"]
	assert.False(t, hasGreeting, "the greeting line is not a capability")
	assert.True(t, hasAuthMechanism(caps, "PLAIN"))
	assert.True(t, hasAuthMechanism(caps, "login"))
	assert.False(t, hasAuthMechanism(caps, "CRAM-MD5"))
	assert.False(t, hasAuthMechanism(map[string]string{}, "PLAIN"))
}

func TestHintForReply(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		stage Stage
		reply string
		want  string // substring
	}{
		{"bad key", StageAuth, "535 5.7.8 Authentication credentials invalid", `literal "apikey"`},
		{"missing capability at AUTH", StageAuth, "535 5.7.1 Authorization failed: email send capability required", "lacks the email send capability"},
		{"intake authorization at DATA", StageData, "535 5.7.1 Authorization failed on message submission", "when the message was submitted"},
		{"pre-TLS AUTH", StageAuth, "538 5.7.11 Encryption required", "STARTTLS"},
		{"throttled", StageAuth, "421 4.7.0 Temporary authentication failure, retry later", "throttled"},
		{"session expired", StageMailFrom, "421 4.4.2 Session expired, close and retry", "maximum lifetime"},
		{"domain refused", StageMailFrom, "550 5.7.1 Domain is not verified. Complete DNS setup before sending email.", "verified sending domain"},
		{"recipient cap", StageRcptTo, "452 4.5.3 Too many recipients", "Too many recipients"},
		{"quota at DATA", StageData, "452 4.5.3 Mail quota exceeded, try again later", "quota"},
		{"size", StageData, "552 5.3.4 Message size exceeds fixed maximum message size", "SIZE"},
		{"unknown 4xx", StageRcptTo, "450 4.2.0 Mailbox busy", "Temporary failure"},
		{"unknown 5xx", StageRcptTo, "550 5.1.1 No such user", "Permanent failure"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r, err := readReply(bufio.NewReader(strings.NewReader(tt.reply + "\r\n")))
			require.NoError(t, err)
			assert.Contains(t, hintForReply(tt.stage, r), tt.want)
		})
	}
}

func TestVerdict(t *testing.T) {
	t.Parallel()
	reply := Reply{Code: 535, Enhanced: "5.7.8", Lines: []string{"535 5.7.8 Authentication credentials invalid"}}
	multi := Reply{Code: 554, Lines: []string{"554-first", "554 second"}}
	tests := []struct {
		name string
		res  Result
		want string
	}{
		{"verified", Result{Verified: true, MessageUUID: "0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"}, "VERIFIED: message UUID 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
		{"server reply", Result{Stage: StageAuth, Reply: &reply}, "FAILED at AUTH: 535 5.7.8 Authentication credentials invalid"},
		{"multi-line reply kept whole", Result{Stage: StageData, Reply: &multi}, "FAILED at DATA: 554-first | 554 second"},
		{"local failure", Result{Stage: StageSTARTTLS, Detail: "server did not advertise STARTTLS"}, "FAILED at STARTTLS: server did not advertise STARTTLS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.res.Verdict())
		})
	}
}

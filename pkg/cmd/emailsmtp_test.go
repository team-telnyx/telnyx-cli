package cmd

import (
	"bytes"
	"context"
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"

	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-cli/internal/smtpverify/smtpmock"
)

// newEmailTestRoot builds a fresh root (flag state must not leak between
// runs) with the same global --api-key flag as the real one.
func newEmailTestRoot(out io.Writer) *cli.Command {
	return &cli.Command{
		Name:      "telnyx",
		Writer:    out,
		ErrWriter: io.Discard,
		Flags: []cli.Flag{
			&requestflag.Flag[string]{Name: "api-key", Sources: cli.EnvVars("TELNYX_API_KEY")},
		},
		Commands: []*cli.Command{newEmailCommand()},
	}
}

func TestEmailSMTPVerifyRegistered(t *testing.T) {
	email := Command.Command("email")
	require.NotNil(t, email, "telnyx email")
	smtp := email.Command("smtp")
	require.NotNil(t, smtp, "telnyx email smtp")
	require.NotNil(t, smtp.Command("verify"), "telnyx email smtp verify")
}

func TestEmailSMTPVerifyCommand(t *testing.T) {
	// Not parallel: subtests use t.Setenv and the package-level TLS hook.
	tlsSrv := smtpmock.Start(t, smtpmock.Script{})
	authFail := smtpmock.Start(t, smtpmock.Script{AuthReply: "535 5.7.8 Authentication credentials invalid"})
	t.Cleanup(func() { emailSMTPVerifyTLSConfig = nil })

	hostPort := func(s *smtpmock.Server) []string {
		return []string{"--host", s.Host, "--port", strings.TrimPrefix(s.Addr, s.Host+":")}
	}
	tests := []struct {
		name       string
		server     *smtpmock.Server // trusted for STARTTLS
		env        map[string]string
		args       []string
		wantErr    string // "" = success
		wantOutput []string
		wantNoOut  bool
		check      func(t *testing.T)
	}{
		{
			name:      "host is required (no default host ships)",
			args:      []string{"--api-key", "KEYx", "--from", "a@example.com"},
			wantErr:   "--host is required",
			wantNoOut: true,
		},
		{
			name:      "API key is required",
			args:      []string{"--host", "smtp.example.com", "--from", "a@example.com"},
			wantErr:   "an API key is required",
			wantNoOut: true,
		},
		{
			name:      "from is required",
			args:      []string{"--host", "smtp.example.com", "--api-key", "KEYx"},
			wantErr:   "--from is required",
			wantNoOut: true,
		},
		{
			name:      "invalid from is rejected before connecting",
			args:      []string{"--host", "smtp.example.com", "--api-key", "KEYx", "--from", "not-an-address"},
			wantErr:   "invalid from address",
			wantNoOut: true,
		},
		{
			name:      "extra arguments rejected",
			args:      []string{"--host", "smtp.example.com", "--api-key", "KEYx", "--from", "a@example.com", "stray"},
			wantErr:   "Unexpected extra arguments",
			wantNoOut: true,
		},
		{
			name:   "verified; --api-key after the subcommand; --to defaults to --from",
			server: tlsSrv,
			args:   append(hostPort(tlsSrv), "--api-key", "KEYx", "--from", "verify@example.com"),
			wantOutput: []string{
				"[AUTH]      S: " + smtpmock.ReplyAuthOK,
				"[RCPT TO]   C: RCPT TO:<verify@example.com>",
				"VERIFIED: message UUID " + smtpmock.QueuedUUID,
			},
		},
		{
			name:   "host and API key from the environment",
			server: tlsSrv,
			env:    map[string]string{"TELNYX_SMTP_HOST": tlsSrv.Host, "TELNYX_API_KEY": "KEYenv"},
			args:   []string{"--port", strings.TrimPrefix(tlsSrv.Addr, tlsSrv.Host+":"), "--from", "verify@example.com", "--to", "ops@example.net"},
			wantOutput: []string{
				"[RCPT TO]   C: RCPT TO:<ops@example.net>",
				"VERIFIED: message UUID " + smtpmock.QueuedUUID,
			},
		},
		{
			name:    "failure exits non-zero with the transcript and verdict printed",
			server:  authFail,
			args:    append(hostPort(authFail), "--api-key", "KEYbad", "--from", "verify@example.com"),
			wantErr: "SMTP verification failed at AUTH",
			wantOutput: []string{
				"FAILED at AUTH: 535 5.7.8 Authentication credentials invalid",
				`hint: The API key was rejected.`,
			},
			check: func(t *testing.T) {
				assert.False(t, slices.Contains(authFail.Verbs(), "MAIL"), "session stops at the first failing reply")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("TELNYX_API_KEY", "")
			t.Setenv("TELNYX_SMTP_HOST", "")
			for k, v := range tt.env {
				t.Setenv(k, v)
			}
			emailSMTPVerifyTLSConfig = nil
			if tt.server != nil {
				emailSMTPVerifyTLSConfig = tt.server.ClientTLS
			}
			var out bytes.Buffer
			args := append([]string{"telnyx", "email", "smtp", "verify"}, tt.args...)
			err := newEmailTestRoot(&out).Run(context.Background(), args)
			if tt.wantErr == "" {
				require.NoError(t, err, out.String())
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			}
			for _, want := range tt.wantOutput {
				assert.Contains(t, out.String(), want)
			}
			if tt.wantNoOut {
				assert.Empty(t, out.String())
			}
			assert.NotContains(t, out.String(), "KEYx")
			assert.NotContains(t, out.String(), "KEYenv")
			assert.NotContains(t, out.String(), "KEYbad")
			if tt.check != nil {
				tt.check(t)
			}
		})
	}
}

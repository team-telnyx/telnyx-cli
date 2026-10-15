//go:build unix

package cmd

import (
	"bytes"
	"context"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"

	"github.com/team-telnyx/telnyx-cli/internal/smtpverify/smtpmock"
)

// Ctrl-C while the command waits for the final DATA reply ends the session,
// not the process: the transcript so far is printed, redacted, then the hint
// and an "interrupted" verdict, last.
func TestEmailSMTPVerifyInterruptPrintsTranscript(t *testing.T) {
	// Not parallel: it signals the test process, and uses the package-level
	// TLS hook and urfave/cli's process-wide exit hook.
	const key = "KEYsigint_0123456789abcdef"
	srv := smtpmock.Start(t, smtpmock.Script{AuthReply: smtpmock.ReplyAuthOK + " for apikey:" + key, HangOn: "."})
	emailSMTPVerifyTLSConfig = srv.ClientTLS
	oldExiter := cli.OsExiter
	cli.OsExiter = func(int) {}
	t.Cleanup(func() { emailSMTPVerifyTLSConfig, cli.OsExiter = nil, oldExiter })
	t.Setenv("TELNYX_API_KEY", "")
	t.Setenv("TELNYX_SMTP_HOST", "")
	// Should the command not catch SIGINT, this keeps it from ending the test
	// binary: the command then waits out --data-timeout instead.
	caught := make(chan os.Signal, 1)
	signal.Notify(caught, os.Interrupt)
	t.Cleanup(func() { signal.Stop(caught) })

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for len(srv.Bodies()) == 0 && ctx.Err() == nil {
			time.Sleep(5 * time.Millisecond)
		}
		if ctx.Err() == nil { // the message is sent: the command waits for the final reply
			_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}
	}()
	var out bytes.Buffer
	start := time.Now()
	err := newEmailTestRoot(&out).Run(ctx, []string{"telnyx", "email", "smtp", "verify",
		"--host", srv.Host, "--port", strconv.Itoa(srv.Port), "--api-key", key,
		"--from", "verify@example.com", "--data-timeout", "10s"})
	assert.Less(t, time.Since(start), 5*time.Second, "the interrupt ends the wait")
	var exitErr cli.ExitCoder
	require.ErrorAs(t, err, &exitErr)
	assert.Equal(t, 1, exitErr.ExitCode())

	lines := strings.Split(strings.TrimSuffix(out.String(), "\n"), "\n")
	require.GreaterOrEqual(t, len(lines), 2, out.String())
	assert.Equal(t, "FAILED at DATA: interrupted", lines[len(lines)-1], "the verdict is the last line printed")
	assert.Equal(t, "hint: The command was interrupted.", lines[len(lines)-2])
	assert.Contains(t, out.String(), "[AUTH]      S: "+smtpmock.ReplyAuthOK+" for apikey:<redacted: API key>\n")
	assert.Contains(t, out.String(), "[DATA]      C: .\n")
	for i := 0; i+8 <= len(key); i++ {
		assert.NotContains(t, out.String(), key[i:i+8])
	}
}

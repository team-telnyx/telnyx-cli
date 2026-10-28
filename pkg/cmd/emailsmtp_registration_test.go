package cmd

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

// Exercise the real generated root, not a test-only command assembled with
// newEmailCommand: SMTP must remain reachable after regeneration/promotion.
func TestEmailSMTPRegisteredOnRoot(t *testing.T) {
	var out bytes.Buffer
	oldWriter, oldErrWriter := Command.Writer, Command.ErrWriter
	Command.Writer, Command.ErrWriter = &out, &out
	t.Cleanup(func() {
		Command.Writer, Command.ErrWriter = oldWriter, oldErrWriter
	})
	if err := Command.Run(context.Background(), []string{"telnyx", "email", "smtp", "verify", "--help"}); err != nil {
		t.Fatalf("SMTP command not reachable from public root: %v\n%s", err, out.String())
	}
	for _, want := range []string{"STARTTLS", "--host", "--from", "--data-timeout"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("SMTP help missing %q:\n%s", want, out.String())
		}
	}
	counts := map[string]int{}
	for _, command := range Command.Commands {
		counts[command.Name]++
	}
	for _, name := range []string{"email", "email-blocks", "email-domains", "email-messages"} {
		if counts[name] != 1 {
			t.Errorf("root must retain exactly one %q command; got %d", name, counts[name])
		}
	}
}

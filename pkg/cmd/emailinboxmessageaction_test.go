// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailInboxesMessagesActionsForward(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:messages:actions", "forward",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--to", "new@example.com",
			"--bcc", "[blind@example.com]",
			"--cc", "[{email: copy@example.com, name: name}]",
			"--html", "html",
			"--text", "FYI",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"to: new@example.com\n" +
			"bcc:\n" +
			"  - blind@example.com\n" +
			"cc:\n" +
			"  - email: copy@example.com\n" +
			"    name: name\n" +
			"html: html\n" +
			"text: FYI\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:messages:actions", "forward",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesMessagesActionsReply(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:messages:actions", "reply",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--html", "P",
			"--text", "Thanks for the update.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"html: P\n" +
			"text: Thanks for the update.\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:messages:actions", "reply",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesMessagesActionsReplyAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:messages:actions", "reply-all",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--html", "P",
			"--text", "Everyone, please review.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"html: P\n" +
			"text: Everyone, please review.\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:messages:actions", "reply-all",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

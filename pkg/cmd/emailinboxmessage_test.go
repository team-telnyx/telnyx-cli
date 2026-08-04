// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailInboxesMessagesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:messages", "update",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--read-at=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("read_at: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:messages", "update",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesMessagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:messages", "list",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-from", "filter[from]",
			"--filter-label", "filter[label]",
			"--filter-read=true",
			"--filter-received-after", "'2019-12-27T18:11:19.117Z'",
			"--filter-received-before", "'2019-12-27T18:11:19.117Z'",
			"--filter-search", "filter[search]",
			"--filter-subject", "filter[subject]",
			"--filter-unread=true",
			"--page-after", "page[after]",
			"--page-size", "1",
		)
	})
}

func TestEmailInboxesMessagesDrafts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:messages", "drafts",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--attachment", "{}",
			"--bcc", "string",
			"--cc", "string",
			"--from-email", "from_email",
			"--from-name", "from_name",
			"--headers", "{foo: string}",
			"--html", "html",
			"--html-body", "html_body",
			"--label", "string",
			"--metadata", "{}",
			"--reply-to", "reply_to",
			"--subject", "subject",
			"--tag", "string",
			"--text", "text",
			"--text-body", "Thanks for the update — I will review today.",
			"--to", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"attachments:\n" +
			"  - {}\n" +
			"bcc:\n" +
			"  - string\n" +
			"cc:\n" +
			"  - string\n" +
			"from_email: from_email\n" +
			"from_name: from_name\n" +
			"headers:\n" +
			"  foo: string\n" +
			"html: html\n" +
			"html_body: html_body\n" +
			"labels:\n" +
			"  - string\n" +
			"metadata: {}\n" +
			"reply_to: reply_to\n" +
			"subject: subject\n" +
			"tags:\n" +
			"  - string\n" +
			"text: text\n" +
			"text_body: Thanks for the update — I will review today.\n" +
			"to:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:messages", "drafts",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

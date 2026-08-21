// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailInboxesDraftsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:drafts", "create",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--attachment", "{foo: bar}",
			"--bcc", "string",
			"--cc", "string",
			"--from-email", "from_email",
			"--from-name", "from_name",
			"--headers", "{foo: string}",
			"--html", "html",
			"--html-body", "html_body",
			"--label", "important",
			"--metadata", "{foo: bar}",
			"--reply-to", "reply_to",
			"--subject", "Quarterly update",
			"--tag", "string",
			"--text", "text",
			"--text-body", "Here is the update.",
			"--to", "{email: recipient@example.com, name: Recipient}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"attachments:\n" +
			"  - foo: bar\n" +
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
			"  - important\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"reply_to: reply_to\n" +
			"subject: Quarterly update\n" +
			"tags:\n" +
			"  - string\n" +
			"text: text\n" +
			"text_body: Here is the update.\n" +
			"to:\n" +
			"  - email: recipient@example.com\n" +
			"    name: Recipient\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:drafts", "create",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesDraftsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:drafts", "retrieve",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--draft-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesDraftsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:drafts", "update",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--draft-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--attachment", "{foo: bar}",
			"--bcc", "string",
			"--cc", "string",
			"--from-email", "from_email",
			"--from-name", "from_name",
			"--headers", "{foo: string}",
			"--html", "html",
			"--html-body", "html_body",
			"--label", "string",
			"--metadata", "{foo: bar}",
			"--reply-to", "reply_to",
			"--subject", "Quarterly update (revised)",
			"--tag", "string",
			"--text", "text",
			"--text-body", "Updated body.",
			"--to", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"attachments:\n" +
			"  - foo: bar\n" +
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
			"metadata:\n" +
			"  foo: bar\n" +
			"reply_to: reply_to\n" +
			"subject: Quarterly update (revised)\n" +
			"tags:\n" +
			"  - string\n" +
			"text: text\n" +
			"text_body: Updated body.\n" +
			"to:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:drafts", "update",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--draft-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesDraftsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:drafts", "list",
			"--max-items", "10",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-status", "draft",
			"--page-after", "page[after]",
			"--page-size", "1",
		)
	})
}

func TestEmailInboxesDraftsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:drafts", "delete",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--draft-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesDraftsPatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:drafts", "patch",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--draft-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--attachment", "{foo: bar}",
			"--bcc", "string",
			"--cc", "string",
			"--from-email", "from_email",
			"--from-name", "from_name",
			"--headers", "{foo: string}",
			"--html", "html",
			"--html-body", "html_body",
			"--label", "string",
			"--metadata", "{foo: bar}",
			"--reply-to", "reply_to",
			"--subject", "Quarterly update (revised)",
			"--tag", "string",
			"--text", "text",
			"--text-body", "Updated body.",
			"--to", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"attachments:\n" +
			"  - foo: bar\n" +
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
			"metadata:\n" +
			"  foo: bar\n" +
			"reply_to: reply_to\n" +
			"subject: Quarterly update (revised)\n" +
			"tags:\n" +
			"  - string\n" +
			"text: text\n" +
			"text_body: Updated body.\n" +
			"to:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:drafts", "patch",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--draft-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesDraftsSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:drafts", "send",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--draft-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

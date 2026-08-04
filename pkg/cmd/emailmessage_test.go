// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestEmailMessagesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "create",
			"--from", "sender@example.com",
			"--to", "recipient@example.com",
			"--attachment", "{content: content, content_id: content_id, content_type: content_type, disposition: disposition, filename: filename}",
			"--bcc", "string",
			"--cc", "string",
			"--forward-of-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--from-name", "from_name",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--headers", "{foo: string}",
			"--html-body", "html_body",
			"--ignore-suppression=true",
			"--in-reply-to-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--inline-css=true",
			"--metadata", "{foo: bar}",
			"--reply-to", "string",
			"--reply-to-all=true",
			"--sandbox-mode=true",
			"--scheduled-at", "'2019-12-27T18:11:19.117Z'",
			"--send-at", "'2019-12-27T18:11:19.117Z'",
			"--subject", "Hello from Telnyx",
			"--tag", "string",
			"--template-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--template-variables", "{foo: bar}",
			"--text-body", "This is a test email.",
			"--tracking-settings", "{click_tracking: true, open_tracking: true}",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailMessagesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "create",
			"--from", "sender@example.com",
			"--to", "recipient@example.com",
			"--attachment.content", "content",
			"--attachment.content-id", "content_id",
			"--attachment.content-type", "content_type",
			"--attachment.disposition", "disposition",
			"--attachment.filename", "filename",
			"--bcc", "string",
			"--cc", "string",
			"--forward-of-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--from-name", "from_name",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--headers", "{foo: string}",
			"--html-body", "html_body",
			"--ignore-suppression=true",
			"--in-reply-to-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--inline-css=true",
			"--metadata", "{foo: bar}",
			"--reply-to", "string",
			"--reply-to-all=true",
			"--sandbox-mode=true",
			"--scheduled-at", "'2019-12-27T18:11:19.117Z'",
			"--send-at", "'2019-12-27T18:11:19.117Z'",
			"--subject", "Hello from Telnyx",
			"--tag", "string",
			"--template-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--template-variables", "{foo: bar}",
			"--text-body", "This is a test email.",
			"--tracking-settings.click-tracking=true",
			"--tracking-settings.open-tracking=true",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"from: sender@example.com\n" +
			"to:\n" +
			"  - recipient@example.com\n" +
			"attachments:\n" +
			"  - content: content\n" +
			"    content_id: content_id\n" +
			"    content_type: content_type\n" +
			"    disposition: disposition\n" +
			"    filename: filename\n" +
			"bcc:\n" +
			"  - string\n" +
			"cc:\n" +
			"  - string\n" +
			"forward_of_message_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"from_name: from_name\n" +
			"group_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"headers:\n" +
			"  foo: string\n" +
			"html_body: html_body\n" +
			"ignore_suppression: true\n" +
			"in_reply_to_message_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"inline_css: true\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"reply_to: string\n" +
			"reply_to_all: true\n" +
			"sandbox_mode: true\n" +
			"scheduled_at: '2019-12-27T18:11:19.117Z'\n" +
			"send_at: '2019-12-27T18:11:19.117Z'\n" +
			"subject: Hello from Telnyx\n" +
			"tags:\n" +
			"  - string\n" +
			"template_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"template_variables:\n" +
			"  foo: bar\n" +
			"text_body: This is a test email.\n" +
			"tracking_settings:\n" +
			"  click_tracking: true\n" +
			"  open_tracking: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-messages", "create",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})
}

func TestEmailMessagesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailMessagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "list",
			"--page-cursor", "page_cursor",
			"--page-size", "1",
		)
	})
}

func TestEmailMessagesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailMessagesBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "batch",
			"--message", "{from: sender@example.com, to: [recipient1@example.com], attachments: [{content: content, content_id: content_id, content_type: content_type, disposition: disposition, filename: filename}], bcc: [string], cc: [string], from_name: from_name, group_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, headers: {foo: string}, html_body: html_body, ignore_suppression: true, inline_css: true, metadata: {foo: bar}, reply_to: string, sandbox_mode: true, scheduled_at: '2019-12-27T18:11:19.117Z', send_at: '2019-12-27T18:11:19.117Z', subject: Hello 1, tags: [string], template_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, template_variables: {foo: bar}, text_body: Message 1, tracking_settings: {click_tracking: true, open_tracking: true}}",
			"--message", "{from: sender@example.com, to: [recipient2@example.com], attachments: [{content: content, content_id: content_id, content_type: content_type, disposition: disposition, filename: filename}], bcc: [string], cc: [string], from_name: from_name, group_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, headers: {foo: string}, html_body: html_body, ignore_suppression: true, inline_css: true, metadata: {foo: bar}, reply_to: string, sandbox_mode: true, scheduled_at: '2019-12-27T18:11:19.117Z', send_at: '2019-12-27T18:11:19.117Z', subject: Hello 2, tags: [string], template_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, template_variables: {foo: bar}, text_body: Message 2, tracking_settings: {click_tracking: true, open_tracking: true}}",
			"--sandbox-mode=false",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailMessagesBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "batch",
			"--message.from", "sender@example.com",
			"--message.to", "[recipient1@example.com]",
			"--message.attachments", "[{content: content, content_id: content_id, content_type: content_type, disposition: disposition, filename: filename}]",
			"--message.bcc", "[string]",
			"--message.cc", "[string]",
			"--message.from-name", "from_name",
			"--message.group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.headers", "{foo: string}",
			"--message.html-body", "html_body",
			"--message.ignore-suppression=true",
			"--message.inline-css=true",
			"--message.metadata", "{foo: bar}",
			"--message.reply-to", "string",
			"--message.sandbox-mode=true",
			"--message.scheduled-at", "2019-12-27T18:11:19.117Z",
			"--message.send-at", "2019-12-27T18:11:19.117Z",
			"--message.subject", "Hello 1",
			"--message.tags", "[string]",
			"--message.template-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.template-variables", "{foo: bar}",
			"--message.text-body", "Message 1",
			"--message.tracking-settings", "{click_tracking: true, open_tracking: true}",
			"--message.from", "sender@example.com",
			"--message.to", "[recipient2@example.com]",
			"--message.attachments", "[{content: content, content_id: content_id, content_type: content_type, disposition: disposition, filename: filename}]",
			"--message.bcc", "[string]",
			"--message.cc", "[string]",
			"--message.from-name", "from_name",
			"--message.group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.headers", "{foo: string}",
			"--message.html-body", "html_body",
			"--message.ignore-suppression=true",
			"--message.inline-css=true",
			"--message.metadata", "{foo: bar}",
			"--message.reply-to", "string",
			"--message.sandbox-mode=true",
			"--message.scheduled-at", "2019-12-27T18:11:19.117Z",
			"--message.send-at", "2019-12-27T18:11:19.117Z",
			"--message.subject", "Hello 2",
			"--message.tags", "[string]",
			"--message.template-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.template-variables", "{foo: bar}",
			"--message.text-body", "Message 2",
			"--message.tracking-settings", "{click_tracking: true, open_tracking: true}",
			"--sandbox-mode=false",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"messages:\n" +
			"  - from: sender@example.com\n" +
			"    to:\n" +
			"      - recipient1@example.com\n" +
			"    attachments:\n" +
			"      - content: content\n" +
			"        content_id: content_id\n" +
			"        content_type: content_type\n" +
			"        disposition: disposition\n" +
			"        filename: filename\n" +
			"    bcc:\n" +
			"      - string\n" +
			"    cc:\n" +
			"      - string\n" +
			"    from_name: from_name\n" +
			"    group_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    headers:\n" +
			"      foo: string\n" +
			"    html_body: html_body\n" +
			"    ignore_suppression: true\n" +
			"    inline_css: true\n" +
			"    metadata:\n" +
			"      foo: bar\n" +
			"    reply_to: string\n" +
			"    sandbox_mode: true\n" +
			"    scheduled_at: '2019-12-27T18:11:19.117Z'\n" +
			"    send_at: '2019-12-27T18:11:19.117Z'\n" +
			"    subject: Hello 1\n" +
			"    tags:\n" +
			"      - string\n" +
			"    template_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    template_variables:\n" +
			"      foo: bar\n" +
			"    text_body: Message 1\n" +
			"    tracking_settings:\n" +
			"      click_tracking: true\n" +
			"      open_tracking: true\n" +
			"  - from: sender@example.com\n" +
			"    to:\n" +
			"      - recipient2@example.com\n" +
			"    attachments:\n" +
			"      - content: content\n" +
			"        content_id: content_id\n" +
			"        content_type: content_type\n" +
			"        disposition: disposition\n" +
			"        filename: filename\n" +
			"    bcc:\n" +
			"      - string\n" +
			"    cc:\n" +
			"      - string\n" +
			"    from_name: from_name\n" +
			"    group_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    headers:\n" +
			"      foo: string\n" +
			"    html_body: html_body\n" +
			"    ignore_suppression: true\n" +
			"    inline_css: true\n" +
			"    metadata:\n" +
			"      foo: bar\n" +
			"    reply_to: string\n" +
			"    sandbox_mode: true\n" +
			"    scheduled_at: '2019-12-27T18:11:19.117Z'\n" +
			"    send_at: '2019-12-27T18:11:19.117Z'\n" +
			"    subject: Hello 2\n" +
			"    tags:\n" +
			"      - string\n" +
			"    template_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    template_variables:\n" +
			"      foo: bar\n" +
			"    text_body: Message 2\n" +
			"    tracking_settings:\n" +
			"      click_tracking: true\n" +
			"      open_tracking: true\n" +
			"sandbox_mode: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-messages", "batch",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})
}

func TestEmailMessagesDeleteAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "delete-all",
			"--address", "dev@stainless.com",
		)
	})
}

func TestEmailMessagesDeleteSchedule(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "delete-schedule",
			"--email-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailMessagesRetrieveEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages", "retrieve-events",
			"--email-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-cursor", "page_cursor",
			"--page-size", "1",
		)
	})
}

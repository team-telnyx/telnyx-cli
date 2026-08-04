// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailDomainsWebhooksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains:webhooks", "create",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--event", "email.sent",
			"--event", "email.delivered",
			"--event", "email.bounced",
			"--url", "https://example.com/webhooks/email",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"events:\n" +
			"  - email.sent\n" +
			"  - email.delivered\n" +
			"  - email.bounced\n" +
			"url: https://example.com/webhooks/email\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-domains:webhooks", "create",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDomainsWebhooksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains:webhooks", "retrieve",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDomainsWebhooksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains:webhooks", "update",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--event", "email.sent",
			"--event", "email.delivered",
			"--event", "email.opened",
			"--url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"events:\n" +
			"  - email.sent\n" +
			"  - email.delivered\n" +
			"  - email.opened\n" +
			"url: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-domains:webhooks", "update",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDomainsWebhooksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains:webhooks", "list",
			"--max-items", "10",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "1",
			"--page-size", "1",
			"--sort", "created_at",
		)
	})
}

func TestEmailDomainsWebhooksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains:webhooks", "delete",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

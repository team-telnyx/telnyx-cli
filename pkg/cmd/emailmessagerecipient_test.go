// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailMessagesRecipientsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages:recipients", "retrieve",
			"--email-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailMessagesRecipientsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-messages:recipients", "list",
			"--max-items", "10",
			"--email-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--kind", "to",
			"--page-cursor", "page_cursor",
			"--page-size", "1",
			"--status", "queued",
		)
	})
}

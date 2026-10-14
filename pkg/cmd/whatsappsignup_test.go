// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappSignupList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:signup", "list",
			"--max-items", "10",
			"--filter-created-at-gte", "'2019-12-27T18:11:19.117Z'",
			"--filter-created-at-lte", "'2019-12-27T18:11:19.117Z'",
			"--filter-customer-id", "filter[customer_id]",
			"--filter-waba-id", "filter[waba_id]",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestWhatsappSignupRetrieveStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:signup", "retrieve-status",
			"--id", "id",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestActionsPurchaseCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"actions:purchase", "create",
			"--amount", "10",
			"--product", "whitelabel",
			"--sim-card-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--status", "standby",
			"--tag", "personal",
			"--tag", "customers",
			"--tag", "active-customers",
			"--whitelabel-name", "Custom SPN",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 10\n" +
			"product: whitelabel\n" +
			"sim_card_group_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"status: standby\n" +
			"tags:\n" +
			"  - personal\n" +
			"  - customers\n" +
			"  - active-customers\n" +
			"whitelabel_name: Custom SPN\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"actions:purchase", "create",
		)
	})
}

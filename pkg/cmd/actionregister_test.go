// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestActionsRegisterCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "actions:register", "create",
			"--api-key", "string",
			"--registration-code", "0000000001",
			"--registration-code", "0000000002",
			"--registration-code", "0000000003",
			"--sim-card-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--status", "standby",
			"--tag", "personal",
			"--tag", "customers",
			"--tag", "active-customers",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"registration_codes:\n" +
			"  - '0000000001'\n" +
			"  - '0000000002'\n" +
			"  - '0000000003'\n" +
			"sim_card_group_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"status: standby\n" +
			"tags:\n" +
			"  - personal\n" +
			"  - customers\n" +
			"  - active-customers\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "actions:register", "create",
			"--api-key", "string",
		)
	})
}

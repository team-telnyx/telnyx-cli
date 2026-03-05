// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestManagedAccountsActionsDisable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts:actions", "disable",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestManagedAccountsActionsEnable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts:actions", "enable",
			"--api-key", "string",
			"--id", "id",
			"--reenable-all-connections=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reenable_all_connections: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "managed-accounts:actions", "enable",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

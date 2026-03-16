// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessagingNumbersBulkUpdatesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-numbers-bulk-updates", "create",
			"--messaging-profile-id", "00000000-0000-0000-0000-000000000000",
			"--number", "+18880000000",
			"--number", "+18880000001",
			"--number", "+18880000002",
			"--assign-only=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"messaging_profile_id: 00000000-0000-0000-0000-000000000000\n" +
			"numbers:\n" +
			"  - '+18880000000'\n" +
			"  - '+18880000001'\n" +
			"  - '+18880000002'\n" +
			"assign_only: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-numbers-bulk-updates", "create",
		)
	})
}

func TestMessagingNumbersBulkUpdatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-numbers-bulk-updates", "retrieve",
			"--order-id", "order_id",
		)
	})
}

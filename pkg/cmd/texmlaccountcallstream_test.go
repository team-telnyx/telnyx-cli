// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsCallsStreamsStreamingSidJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls:streams", "streaming-sid-json",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
			"--streaming-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--status", "stopped",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("Status: stopped")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml:accounts:calls:streams", "streaming-sid-json",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
			"--streaming-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

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
			t,
			"--api-key", "string",
			"texml:accounts:calls:streams", "streaming-sid-json",
			"--account-sid", "f5586561-8ff0-4291-a51b-c7dfe4139ff9",
			"--call-sid", "call_sid",
			"--streaming-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--status", "stopped",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("Status: stopped")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:calls:streams", "streaming-sid-json",
			"--account-sid", "f5586561-8ff0-4291-a51b-c7dfe4139ff9",
			"--call-sid", "call_sid",
			"--streaming-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

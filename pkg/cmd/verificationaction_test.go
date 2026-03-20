// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVerificationsActionsVerify(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verifications:actions", "verify",
			"--verification-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--code", "17686",
			"--status", "accepted",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"code: '17686'\n" +
			"status: accepted\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"verifications:actions", "verify",
			"--verification-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}

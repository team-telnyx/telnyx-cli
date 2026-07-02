// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVerifiedNumbersActionsSubmitVerificationCode(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verified-numbers:actions", "submit-verification-code",
			"--phone-number", "+15551234567",
			"--verification-code", "123456",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("verification_code: '123456'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"verified-numbers:actions", "submit-verification-code",
			"--phone-number", "+15551234567",
		)
	})
}

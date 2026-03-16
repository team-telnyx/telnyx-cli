// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVerificationsByPhoneNumberActionsVerify(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verifications:by-phone-number:actions", "verify",
			"--phone-number", "+13035551234",
			"--code", "17686",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"code: '17686'\n" +
			"verify_profile_id: 12ade33a-21c0-473b-b055-b3c836e1c292\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"verifications:by-phone-number:actions", "verify",
			"--phone-number", "+13035551234",
		)
	})
}

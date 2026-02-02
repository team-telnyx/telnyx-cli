// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestVerificationsByPhoneNumberActionsVerify(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verifications:by-phone-number:actions", "verify",
		"--phone-number", "+13035551234",
		"--code", "17686",
		"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestVerifiedNumbersActionsSubmitVerificationCode(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verified-numbers:actions", "submit-verification-code",
		"--phone-number", "+15551234567",
		"--verification-code", "123456",
	)
}

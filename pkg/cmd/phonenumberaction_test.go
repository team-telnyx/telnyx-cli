// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPhoneNumbersActionsChangeBundleStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:actions", "change-bundle-status",
		"--api-key", "string",
		"--id", "1293384261075731499",
		"--bundle-id", "5194d8fc-87e6-4188-baa9-1c434bbe861b",
	)
}

func TestPhoneNumbersActionsEnableEmergency(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:actions", "enable-emergency",
		"--api-key", "string",
		"--id", "1293384261075731499",
		"--emergency-address-id", "53829456729313",
		"--emergency-enabled=true",
	)
}

func TestPhoneNumbersActionsVerifyOwnership(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:actions", "verify-ownership",
		"--api-key", "string",
		"--phone-number", "+15551234567",
	)
}

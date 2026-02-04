// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPhoneNumbersRegulatoryRequirementsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers-regulatory-requirements", "retrieve",
		"--filter", "{phone_number: '+41215470622,+41215470633'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumbersRegulatoryRequirementsRetrieve)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers-regulatory-requirements", "retrieve",
		"--filter.phone-number", "+41215470622,+41215470633",
	)
}

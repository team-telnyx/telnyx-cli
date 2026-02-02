// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestRegulatoryRequirementsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"regulatory-requirements", "retrieve",
		"--filter", "{action: ordering, country_code: DE, phone_number: '+41215470622', phone_number_type: local, requirement_group_id: d4c0b4a6-7bd2-40c5-a3b9-2acd99e212b2}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(regulatoryRequirementsRetrieve)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"regulatory-requirements", "retrieve",
		"--filter.action", "ordering",
		"--filter.country-code", "DE",
		"--filter.phone-number", "+41215470622",
		"--filter.phone-number-type", "local",
		"--filter.requirement-group-id", "d4c0b4a6-7bd2-40c5-a3b9-2acd99e212b2",
	)
}

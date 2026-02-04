// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestDynamicEmergencyAddressesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-addresses", "create",
		"--administrative-area", "TX",
		"--country-code", "US",
		"--house-number", "600",
		"--locality", "Austin",
		"--postal-code", "78701",
		"--street-name", "Congress",
		"--extended-address", "extended_address",
		"--house-suffix", "house_suffix",
		"--street-post-directional", "street_post_directional",
		"--street-pre-directional", "street_pre_directional",
		"--street-suffix", "St",
	)
}

func TestDynamicEmergencyAddressesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-addresses", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestDynamicEmergencyAddressesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-addresses", "list",
		"--filter", "{country_code: country_code, status: pending}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(dynamicEmergencyAddressesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-addresses", "list",
		"--filter.country-code", "country_code",
		"--filter.status", "pending",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestDynamicEmergencyAddressesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-addresses", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

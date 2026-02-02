// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestInventoryCoverageList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"inventory-coverage", "list",
		"--filter", "{administrative_area: administrative_area, count: true, country_code: AT, features: [voice, sms], groupBy: locality, npa: 0, nxx: 0, phone_number_type: local}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(inventoryCoverageList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"inventory-coverage", "list",
		"--filter.administrative-area", "administrative_area",
		"--filter.count=true",
		"--filter.country-code", "AT",
		"--filter.features", "[voice, sms]",
		"--filter.group-by", "locality",
		"--filter.npa", "0",
		"--filter.nxx", "0",
		"--filter.phone-number-type", "local",
	)
}

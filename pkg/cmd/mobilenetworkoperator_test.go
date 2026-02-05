// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMobileNetworkOperatorsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-network-operators", "list",
		"--filter", "{country_code: US, mcc: '310', mnc: '410', name: {contains: T&T, ends_with: T, starts_with: AT}, network_preferences_enabled: true, tadig: USACG}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(mobileNetworkOperatorsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-network-operators", "list",
		"--filter.country-code", "US",
		"--filter.mcc", "310",
		"--filter.mnc", "410",
		"--filter.name", "{contains: T&T, ends_with: T, starts_with: AT}",
		"--filter.network-preferences-enabled=true",
		"--filter.tadig", "USACG",
		"--page-number", "0",
		"--page-size", "0",
	)
}

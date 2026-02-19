// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNetworkCoverageList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"network-coverage", "list",
		"--filter", "{location.code: silicon_valley-ca, location.pop: SV1, location.region: AMER, location.site: SJC}",
		"--filters", "{available_services: cloud_vpn}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(networkCoverageList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"network-coverage", "list",
		"--filter.location-code", "silicon_valley-ca",
		"--filter.location-pop", "SV1",
		"--filter.location-region", "AMER",
		"--filter.location-site", "SJC",
		"--filters.available-services", "cloud_vpn",
		"--page-number", "0",
		"--page-size", "0",
	)
}

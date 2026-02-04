// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestVirtualCrossConnectsCoverageList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"virtual-cross-connects-coverage", "list",
		"--filter", "{cloud_provider: aws, cloud_provider_region: us-east-1, location.code: silicon_valley-ca, location.pop: SV1, location.region: AMER, location.site: SJC}",
		"--filters", "{available_bandwidth: 0}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(virtualCrossConnectsCoverageList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"virtual-cross-connects-coverage", "list",
		"--filter.cloud-provider", "aws",
		"--filter.cloud-provider-region", "us-east-1",
		"--filter.location-code", "silicon_valley-ca",
		"--filter.location-pop", "SV1",
		"--filter.location-region", "AMER",
		"--filter.location-site", "SJC",
		"--filters.available-bandwidth", "0",
		"--page-number", "0",
		"--page-size", "0",
	)
}

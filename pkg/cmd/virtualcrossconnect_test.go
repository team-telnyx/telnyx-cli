// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestVirtualCrossConnectsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"virtual-cross-connects", "create",
		"--region-code", "ashburn-va",
		"--bandwidth-mbps", "50",
		"--bgp-asn", "1234",
		"--cloud-provider", "aws",
		"--cloud-provider-region", "us-east-1",
		"--name", "test interface",
		"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--primary-bgp-key", "yFV4wEPtPVPfDUGLWiyQzwga",
		"--primary-cloud-account-id", "123456789012",
		"--primary-cloud-ip", "169.254.0.2",
		"--primary-telnyx-ip", "169.254.0.1",
		"--secondary-bgp-key", "ge1lONeK9RcA83uuWaw9DvZy",
		"--secondary-cloud-account-id", "",
		"--secondary-cloud-ip", "169.254.0.4",
		"--secondary-telnyx-ip", "169.254.0.3",
	)
}

func TestVirtualCrossConnectsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"virtual-cross-connects", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestVirtualCrossConnectsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"virtual-cross-connects", "update",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--primary-cloud-ip", "169.254.0.2",
		"--primary-enabled=true",
		"--primary-routing-announcement=false",
		"--secondary-cloud-ip", "169.254.0.4",
		"--secondary-enabled=true",
		"--secondary-routing-announcement=false",
	)
}

func TestVirtualCrossConnectsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"virtual-cross-connects", "list",
		"--filter", "{network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(virtualCrossConnectsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"virtual-cross-connects", "list",
		"--filter.network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestVirtualCrossConnectsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"virtual-cross-connects", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

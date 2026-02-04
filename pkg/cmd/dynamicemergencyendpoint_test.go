// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestDynamicEmergencyEndpointsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-endpoints", "create",
		"--callback-number", "+13125550000",
		"--caller-name", "Jane Doe Desk Phone",
		"--dynamic-emergency-address-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
	)
}

func TestDynamicEmergencyEndpointsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-endpoints", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestDynamicEmergencyEndpointsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-endpoints", "list",
		"--filter", "{country_code: country_code, status: pending}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(dynamicEmergencyEndpointsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-endpoints", "list",
		"--filter.country-code", "country_code",
		"--filter.status", "pending",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestDynamicEmergencyEndpointsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dynamic-emergency-endpoints", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

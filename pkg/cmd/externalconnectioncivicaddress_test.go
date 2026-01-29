// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestExternalConnectionsCivicAddressesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:civic-addresses", "retrieve",
		"--id", "id",
		"--address-id", "318fb664-d341-44d2-8405-e6bfb9ced6d9",
	)
}

func TestExternalConnectionsCivicAddressesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:civic-addresses", "list",
		"--id", "id",
		"--filter", "{country: [US, CA, MX, BR]}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(externalConnectionsCivicAddressesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:civic-addresses", "list",
		"--id", "id",
		"--filter.country", "[US, CA, MX, BR]",
	)
}

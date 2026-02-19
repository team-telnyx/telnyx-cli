// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAddressesActionsAcceptSuggestions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses:actions", "accept-suggestions",
		"--address-uuid", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--id", "id",
	)
}

func TestAddressesActionsValidate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses:actions", "validate",
		"--country-code", "US",
		"--postal-code", "78701",
		"--street-address", "600 Congress Avenue",
		"--administrative-area", "TX",
		"--extended-address", "14th Floor",
		"--locality", "Austin",
	)
}

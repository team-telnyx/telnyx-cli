// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestRequirementsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirements", "retrieve",
		"--id", "a9dad8d5-fdbd-49d7-aa23-39bb08a5ebaa",
	)
}

func TestRequirementsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirements", "list",
		"--filter", "{action: porting, country_code: US, phone_number_type: local}",
		"--page", "{number: 1, size: 1}",
		"--sort", "country_code",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(requirementsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirements", "list",
		"--filter.action", "porting",
		"--filter.country-code", "US",
		"--filter.phone-number-type", "local",
		"--page.number", "1",
		"--page.size", "1",
		"--sort", "country_code",
	)
}

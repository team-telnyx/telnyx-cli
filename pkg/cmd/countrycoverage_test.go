// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestCountryCoverageRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"country-coverage", "retrieve",
		"--api-key", "string",
	)
}

func TestCountryCoverageRetrieveCountry(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"country-coverage", "retrieve-country",
		"--api-key", "string",
		"--country-code", "US",
	)
}

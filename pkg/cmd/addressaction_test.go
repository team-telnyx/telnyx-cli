// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAddressesActionsAcceptSuggestions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"addresses:actions", "accept-suggestions",
			"--address-uuid", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("id: id")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"addresses:actions", "accept-suggestions",
			"--address-uuid", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAddressesActionsValidate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"addresses:actions", "validate",
			"--country-code", "US",
			"--postal-code", "78701",
			"--street-address", "600 Congress Avenue",
			"--administrative-area", "TX",
			"--extended-address", "14th Floor",
			"--locality", "Austin",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"country_code: US\n" +
			"postal_code: '78701'\n" +
			"street_address: 600 Congress Avenue\n" +
			"administrative_area: TX\n" +
			"extended_address: 14th Floor\n" +
			"locality: Austin\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"addresses:actions", "validate",
		)
	})
}

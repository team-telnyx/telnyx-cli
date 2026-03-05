// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAvailablePhoneNumberBlocksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "available-phone-number-blocks", "list",
			"--api-key", "string",
			"--filter", "{country_code: country_code, locality: locality, national_destination_code: national_destination_code, phone_number_type: local}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(availablePhoneNumberBlocksList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "available-phone-number-blocks", "list",
			"--api-key", "string",
			"--filter.country-code", "country_code",
			"--filter.locality", "locality",
			"--filter.national-destination-code", "national_destination_code",
			"--filter.phone-number-type", "local",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestAvailablePhoneNumbersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"available-phone-numbers", "list",
		"--filter", "{administrative_area: administrative_area, best_effort: true, country_code: country_code, exclude_held_numbers: true, features: [sms], limit: 0, locality: locality, national_destination_code: national_destination_code, phone_number: {contains: contains, ends_with: ends_with, starts_with: starts_with}, phone_number_type: local, quickship: true, rate_center: rate_center, reservable: true}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(availablePhoneNumbersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"available-phone-numbers", "list",
		"--filter.administrative-area", "administrative_area",
		"--filter.best-effort=true",
		"--filter.country-code", "country_code",
		"--filter.exclude-held-numbers=true",
		"--filter.features", "[sms]",
		"--filter.limit", "0",
		"--filter.locality", "locality",
		"--filter.national-destination-code", "national_destination_code",
		"--filter.phone-number", "{contains: contains, ends_with: ends_with, starts_with: starts_with}",
		"--filter.phone-number-type", "local",
		"--filter.quickship=true",
		"--filter.rate-center", "rate_center",
		"--filter.reservable=true",
	)
}

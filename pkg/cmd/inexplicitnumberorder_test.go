// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestInexplicitNumberOrdersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"inexplicit-number-orders", "create",
		"--ordering-group", "{count_requested: count_requested, country_iso: US, phone_number_type: phone_number_type, administrative_area: administrative_area, exclude_held_numbers: true, features: [string], locality: locality, national_destination_code: national_destination_code, phone_number: {contains: contains, ends_with: ends_with, starts_with: starts_with}, quickship: true, strategy: always}",
		"--billing-group-id", "billing_group_id",
		"--connection-id", "connection_id",
		"--customer-reference", "customer_reference",
		"--messaging-profile-id", "messaging_profile_id",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(inexplicitNumberOrdersCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"inexplicit-number-orders", "create",
		"--ordering-group.count-requested", "count_requested",
		"--ordering-group.country-iso", "US",
		"--ordering-group.phone-number-type", "phone_number_type",
		"--ordering-group.administrative-area", "administrative_area",
		"--ordering-group.exclude-held-numbers=true",
		"--ordering-group.features", "[string]",
		"--ordering-group.locality", "locality",
		"--ordering-group.national-destination-code", "national_destination_code",
		"--ordering-group.phone-number", "{contains: contains, ends_with: ends_with, starts_with: starts_with}",
		"--ordering-group.quickship=true",
		"--ordering-group.strategy", "always",
		"--billing-group-id", "billing_group_id",
		"--connection-id", "connection_id",
		"--customer-reference", "customer_reference",
		"--messaging-profile-id", "messaging_profile_id",
	)
}

func TestInexplicitNumberOrdersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"inexplicit-number-orders", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestInexplicitNumberOrdersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"inexplicit-number-orders", "list",
		"--page-number", "1",
		"--page-size", "1",
	)
}

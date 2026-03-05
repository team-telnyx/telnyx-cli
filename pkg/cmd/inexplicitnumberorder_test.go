// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestInexplicitNumberOrdersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "inexplicit-number-orders", "create",
			"--api-key", "string",
			"--ordering-group", "{count_requested: count_requested, country_iso: US, phone_number_type: phone_number_type, administrative_area: administrative_area, exclude_held_numbers: true, features: [string], locality: locality, national_destination_code: national_destination_code, phone_number: {contains: contains, ends_with: ends_with, starts_with: starts_with}, quickship: true, strategy: always}",
			"--billing-group-id", "billing_group_id",
			"--connection-id", "connection_id",
			"--customer-reference", "customer_reference",
			"--messaging-profile-id", "messaging_profile_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(inexplicitNumberOrdersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "inexplicit-number-orders", "create",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ordering_groups:\n" +
			"  - count_requested: count_requested\n" +
			"    country_iso: US\n" +
			"    phone_number_type: phone_number_type\n" +
			"    administrative_area: administrative_area\n" +
			"    exclude_held_numbers: true\n" +
			"    features:\n" +
			"      - string\n" +
			"    locality: locality\n" +
			"    national_destination_code: national_destination_code\n" +
			"    phone_number:\n" +
			"      contains: contains\n" +
			"      ends_with: ends_with\n" +
			"      starts_with: starts_with\n" +
			"    quickship: true\n" +
			"    strategy: always\n" +
			"billing_group_id: billing_group_id\n" +
			"connection_id: connection_id\n" +
			"customer_reference: customer_reference\n" +
			"messaging_profile_id: messaging_profile_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "inexplicit-number-orders", "create",
			"--api-key", "string",
		)
	})
}

func TestInexplicitNumberOrdersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "inexplicit-number-orders", "retrieve",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestInexplicitNumberOrdersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "inexplicit-number-orders", "list",
			"--api-key", "string",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

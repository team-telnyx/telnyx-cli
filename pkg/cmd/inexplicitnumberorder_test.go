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
			t,
			"--api-key", "string",
			"inexplicit-number-orders", "create",
			"--ordering-group", "{count_requested: '5', country_iso: US, phone_number_type: local, administrative_area: CA, exclude_held_numbers: true, features: [voice], locality: locality, national_destination_code: national_destination_code, phone_number: {contains: contains, ends_with: ends_with, starts_with: starts_with}, quickship: true, strategy: always}",
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
			t,
			"--api-key", "string",
			"inexplicit-number-orders", "create",
			"--ordering-group.count-requested", "5",
			"--ordering-group.country-iso", "US",
			"--ordering-group.phone-number-type", "local",
			"--ordering-group.administrative-area", "CA",
			"--ordering-group.exclude-held-numbers=true",
			"--ordering-group.features", "[voice]",
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
			"  - count_requested: '5'\n" +
			"    country_iso: US\n" +
			"    phone_number_type: local\n" +
			"    administrative_area: CA\n" +
			"    exclude_held_numbers: true\n" +
			"    features:\n" +
			"      - voice\n" +
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
			t, pipeData,
			"--api-key", "string",
			"inexplicit-number-orders", "create",
		)
	})
}

func TestInexplicitNumberOrdersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inexplicit-number-orders", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestInexplicitNumberOrdersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inexplicit-number-orders", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

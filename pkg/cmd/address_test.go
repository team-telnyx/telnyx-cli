// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAddressesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "create",
			"--api-key", "string",
			"--business-name", "Toy-O'Kon",
			"--country-code", "US",
			"--first-name", "Alfred",
			"--last-name", "Foster",
			"--locality", "Austin",
			"--street-address", "600 Congress Avenue",
			"--address-book=false",
			"--administrative-area", "TX",
			"--borough", "Guadalajara",
			"--customer-reference", "MY REF 001",
			"--extended-address", "14th Floor",
			"--neighborhood", "Ciudad de los deportes",
			"--phone-number", "+12125559000",
			"--postal-code", "78701",
			"--validate-address=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"business_name: Toy-O'Kon\n" +
			"country_code: US\n" +
			"first_name: Alfred\n" +
			"last_name: Foster\n" +
			"locality: Austin\n" +
			"street_address: 600 Congress Avenue\n" +
			"address_book: false\n" +
			"administrative_area: TX\n" +
			"borough: Guadalajara\n" +
			"customer_reference: MY REF 001\n" +
			"extended_address: 14th Floor\n" +
			"neighborhood: Ciudad de los deportes\n" +
			"phone_number: '+12125559000'\n" +
			"postal_code: '78701'\n" +
			"validate_address: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "addresses", "create",
			"--api-key", "string",
		)
	})
}

func TestAddressesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAddressesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{address_book: {eq: eq}, customer_reference: string, street_address: {contains: contains}, used_as_emergency: used_as_emergency}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "street_address",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(addressesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.address-book", "{eq: eq}",
			"--filter.customer-reference", "string",
			"--filter.street-address", "{contains: contains}",
			"--filter.used-as-emergency", "used_as_emergency",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "street_address",
		)
	})
}

func TestAddressesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "addresses", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

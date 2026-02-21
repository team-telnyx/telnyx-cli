// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAddressesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "create",
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
}

func TestAddressesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "retrieve",
		"--id", "id",
	)
}

func TestAddressesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "list",
		"--filter", "{address_book: {eq: eq}, customer_reference: string, street_address: {contains: contains}, used_as_emergency: used_as_emergency}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "street_address",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(addressesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "list",
		"--filter.address-book", "{eq: eq}",
		"--filter.customer-reference", "string",
		"--filter.street-address", "{contains: contains}",
		"--filter.used-as-emergency", "used_as_emergency",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "street_address",
	)
}

func TestAddressesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"addresses", "delete",
		"--id", "id",
	)
}

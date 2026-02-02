// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestUserAddressesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"user-addresses", "create",
		"--business-name", "Toy-O'Kon",
		"--country-code", "US",
		"--first-name", "Alfred",
		"--last-name", "Foster",
		"--locality", "Austin",
		"--street-address", "600 Congress Avenue",
		"--administrative-area", "TX",
		"--borough", "Guadalajara",
		"--customer-reference", "MY REF 001",
		"--extended-address", "14th Floor",
		"--neighborhood", "Ciudad de los deportes",
		"--phone-number", "+12125559000",
		"--postal-code", "78701",
		"--skip-address-verification=true",
	)
}

func TestUserAddressesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"user-addresses", "retrieve",
		"--id", "id",
	)
}

func TestUserAddressesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"user-addresses", "list",
		"--filter", "{customer_reference: {contains: contains, eq: eq}, street_address: {contains: contains}}",
		"--page", "{number: 1, size: 1}",
		"--sort", "street_address",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(userAddressesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"user-addresses", "list",
		"--filter.customer-reference", "{contains: contains, eq: eq}",
		"--filter.street-address", "{contains: contains}",
		"--page.number", "1",
		"--page.size", "1",
		"--sort", "street_address",
	)
}

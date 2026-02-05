// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSimCardOrdersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-orders", "create",
		"--address-id", "1293384261075731499",
		"--quantity", "23",
	)
}

func TestSimCardOrdersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-orders", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardOrdersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-orders", "list",
		"--filter", "{address.administrative_area: TX, address.country_code: US, address.extended_address: 14th Floor, address.id: '1293384261075731499', address.locality: Austin, address.postal_code: '78701', address.street_address: 600 Congress Avenue, cost.amount: '2.53', cost.currency: USD, created_at: '2018-02-02T22:25:27.521Z', quantity: 21, updated_at: '2018-02-02T22:25:27.521Z'}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(simCardOrdersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-orders", "list",
		"--filter.address-administrative-area", "TX",
		"--filter.address-country-code", "US",
		"--filter.address-extended-address", "14th Floor",
		"--filter.address-id", "1293384261075731499",
		"--filter.address-locality", "Austin",
		"--filter.address-postal-code", "78701",
		"--filter.address-street-address", "600 Congress Avenue",
		"--filter.cost-amount", "2.53",
		"--filter.cost-currency", "USD",
		"--filter.created-at", "2018-02-02T22:25:27.521Z",
		"--filter.quantity", "21",
		"--filter.updated-at", "2018-02-02T22:25:27.521Z",
		"--page-number", "0",
		"--page-size", "0",
	)
}

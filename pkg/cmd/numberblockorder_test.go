// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNumberBlockOrdersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-block-orders", "create",
		"--range", "10",
		"--starting-number", "+19705555000",
		"--connection-id", "346789098765567",
		"--customer-reference", "MY REF 001",
		"--messaging-profile-id", "abc85f64-5717-4562-b3fc-2c9600",
	)
}

func TestNumberBlockOrdersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-block-orders", "retrieve",
		"--number-block-order-id", "number_block_order_id",
	)
}

func TestNumberBlockOrdersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-block-orders", "list",
		"--filter", "{created_at: {gt: '2018-01-01T00:00:00.000000Z', lt: '2018-01-01T00:00:00.000000Z'}, phone_numbers.starting_number: '+19705555000', status: pending}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(numberBlockOrdersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-block-orders", "list",
		"--filter.created-at", "{gt: '2018-01-01T00:00:00.000000Z', lt: '2018-01-01T00:00:00.000000Z'}",
		"--filter.phone-numbers-starting-number", "+19705555000",
		"--filter.status", "pending",
		"--page-number", "0",
		"--page-size", "0",
	)
}

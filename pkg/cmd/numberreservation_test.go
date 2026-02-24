// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNumberReservationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-reservations", "create",
		"--customer-reference", "MY REF 001",
		"--phone-number", "{phone_number: '+19705555098'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(numberReservationsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-reservations", "create",
		"--customer-reference", "MY REF 001",
		"--phone-number.phone-number", "+19705555098",
	)
}

func TestNumberReservationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-reservations", "retrieve",
		"--number-reservation-id", "number_reservation_id",
	)
}

func TestNumberReservationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-reservations", "list",
		"--filter", "{created_at: {gt: gt, lt: lt}, customer_reference: customer_reference, phone_numbers.phone_number: phone_numbers.phone_number, status: status}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(numberReservationsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-reservations", "list",
		"--filter.created-at", "{gt: gt, lt: lt}",
		"--filter.customer-reference", "customer_reference",
		"--filter.phone-numbers-phone-number", "phone_numbers.phone_number",
		"--filter.status", "status",
		"--page-number", "0",
		"--page-size", "0",
	)
}

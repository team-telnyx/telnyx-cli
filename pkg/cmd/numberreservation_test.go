// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNumberReservationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-reservations", "create",
			"--customer-reference", "MY REF 001",
			"--phone-number", "{phone_number: '+19705555098'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(numberReservationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-reservations", "create",
			"--customer-reference", "MY REF 001",
			"--phone-number.phone-number", "+19705555098",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customer_reference: MY REF 001\n" +
			"phone_numbers:\n" +
			"  - phone_number: '+19705555098'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"number-reservations", "create",
		)
	})
}

func TestNumberReservationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-reservations", "retrieve",
			"--number-reservation-id", "number_reservation_id",
		)
	})
}

func TestNumberReservationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-reservations", "list",
			"--max-items", "10",
			"--filter", "{created_at: {gt: gt, lt: lt}, customer_reference: customer_reference, phone_numbers.phone_number: phone_numbers.phone_number, status: status}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(numberReservationsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"number-reservations", "list",
			"--max-items", "10",
			"--filter.created-at", "{gt: gt, lt: lt}",
			"--filter.customer-reference", "customer_reference",
			"--filter.phone-numbers-phone-number", "phone_numbers.phone_number",
			"--filter.status", "status",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

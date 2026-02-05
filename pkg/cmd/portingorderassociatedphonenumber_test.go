// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersAssociatedPhoneNumbersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:associated-phone-numbers", "create",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--action", "keep",
		"--phone-number-range", "{end_at: '+441234567899', start_at: '+441234567890'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersAssociatedPhoneNumbersCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:associated-phone-numbers", "create",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--action", "keep",
		"--phone-number-range.end-at", "+441234567899",
		"--phone-number-range.start-at", "+441234567890",
	)
}

func TestPortingOrdersAssociatedPhoneNumbersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:associated-phone-numbers", "list",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter", "{action: keep, phone_number: '+441234567890'}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "{value: '-created_at'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersAssociatedPhoneNumbersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:associated-phone-numbers", "list",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.action", "keep",
		"--filter.phone-number", "+441234567890",
		"--page-number", "0",
		"--page-size", "0",
		"--sort.value", "-created_at",
	)
}

func TestPortingOrdersAssociatedPhoneNumbersDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:associated-phone-numbers", "delete",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersPhoneNumberExtensionsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-extensions", "create",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--activation-range", "{end_at: 10, start_at: 1}",
		"--extension-range", "{end_at: 10, start_at: 1}",
		"--porting-phone-number-id", "f24151b6-3389-41d3-8747-7dd8c681e5e2",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersPhoneNumberExtensionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-extensions", "create",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--activation-range.end-at", "10",
		"--activation-range.start-at", "1",
		"--extension-range.end-at", "10",
		"--extension-range.start-at", "1",
		"--porting-phone-number-id", "f24151b6-3389-41d3-8747-7dd8c681e5e2",
	)
}

func TestPortingOrdersPhoneNumberExtensionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-extensions", "list",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter", "{porting_phone_number_id: 04f8f1b9-310c-4a3c-963e-7dfc54765140}",
		"--page", "{number: 1, size: 1}",
		"--sort", "{value: created_at}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersPhoneNumberExtensionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-extensions", "list",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.porting-phone-number-id", "04f8f1b9-310c-4a3c-963e-7dfc54765140",
		"--page.number", "1",
		"--page.size", "1",
		"--sort.value", "created_at",
	)
}

func TestPortingOrdersPhoneNumberExtensionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-extensions", "delete",
		"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

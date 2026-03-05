// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersPhoneNumberBlocksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-blocks", "create",
			"--api-key", "string",
			"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--activation-range", "{end_at: '+4930244999910', start_at: '+4930244999901'}",
			"--phone-number-range", "{end_at: '+4930244999910', start_at: '+4930244999901'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingOrdersPhoneNumberBlocksCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-blocks", "create",
			"--api-key", "string",
			"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--activation-range.end-at", "+4930244999910",
			"--activation-range.start-at", "+4930244999901",
			"--phone-number-range.end-at", "+4930244999910",
			"--phone-number-range.start-at", "+4930244999901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"activation_ranges:\n" +
			"  - end_at: '+4930244999910'\n" +
			"    start_at: '+4930244999901'\n" +
			"phone_number_range:\n" +
			"  end_at: '+4930244999910'\n" +
			"  start_at: '+4930244999901'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting-orders:phone-number-blocks", "create",
			"--api-key", "string",
			"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingOrdersPhoneNumberBlocksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-blocks", "list",
			"--api-key", "string",
			"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter", "{activation_status: Active, phone_number: ['+12003151212'], portability_status: confirmed, porting_order_id: [f3575e15-32ce-400e-a4c0-dd78800c20b0], status: in-process, support_key: sr_a12345}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "{value: created_at}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingOrdersPhoneNumberBlocksList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-blocks", "list",
			"--api-key", "string",
			"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter.activation-status", "Active",
			"--filter.phone-number", "['+12003151212']",
			"--filter.portability-status", "confirmed",
			"--filter.porting-order-id", "[f3575e15-32ce-400e-a4c0-dd78800c20b0]",
			"--filter.status", "in-process",
			"--filter.support-key", "sr_a12345",
			"--page-number", "0",
			"--page-size", "0",
			"--sort.value", "created_at",
		)
	})
}

func TestPortingOrdersPhoneNumberBlocksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-blocks", "delete",
			"--api-key", "string",
			"--porting-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

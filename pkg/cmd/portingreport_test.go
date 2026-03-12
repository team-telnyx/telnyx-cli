// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingReportsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:reports", "create",
			"--api-key", "string",
			"--params", "{filters: {created_at__gt: '2019-12-27T18:11:19.117Z', created_at__lt: '2019-12-27T18:11:19.117Z', customer_reference__in: [my-customer-reference], status__in: [draft]}}",
			"--report-type", "export_porting_orders_csv",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingReportsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting:reports", "create",
			"--api-key", "string",
			"--params.filters", "{created_at__gt: '2019-12-27T18:11:19.117Z', created_at__lt: '2019-12-27T18:11:19.117Z', customer_reference__in: [my-customer-reference], status__in: [draft]}",
			"--report-type", "export_porting_orders_csv",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"params:\n" +
			"  filters:\n" +
			"    created_at__gt: '2019-12-27T18:11:19.117Z'\n" +
			"    created_at__lt: '2019-12-27T18:11:19.117Z'\n" +
			"    customer_reference__in:\n" +
			"      - my-customer-reference\n" +
			"    status__in:\n" +
			"      - draft\n" +
			"report_type: export_porting_orders_csv\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting:reports", "create",
			"--api-key", "string",
		)
	})
}

func TestPortingReportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:reports", "retrieve",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingReportsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:reports", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{report_type: export_porting_orders_csv, status: completed}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingReportsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting:reports", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.report-type", "export_porting_orders_csv",
			"--filter.status", "completed",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

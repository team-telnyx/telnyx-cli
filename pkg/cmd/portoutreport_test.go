// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortoutsReportsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:reports", "create",
		"--params", "{filters: {created_at__gt: '2019-12-27T18:11:19.117Z', created_at__lt: '2019-12-27T18:11:19.117Z', customer_reference__in: [my-customer-reference], end_user_name: McPortersen, phone_numbers__overlaps: ['+1234567890'], status__in: [pending]}}",
		"--report-type", "export_portouts_csv",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portoutsReportsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:reports", "create",
		"--params.filters", "{created_at__gt: '2019-12-27T18:11:19.117Z', created_at__lt: '2019-12-27T18:11:19.117Z', customer_reference__in: [my-customer-reference], end_user_name: McPortersen, phone_numbers__overlaps: ['+1234567890'], status__in: [pending]}",
		"--report-type", "export_portouts_csv",
	)
}

func TestPortoutsReportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:reports", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestPortoutsReportsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:reports", "list",
		"--filter", "{report_type: export_portouts_csv, status: completed}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portoutsReportsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:reports", "list",
		"--filter.report-type", "export_portouts_csv",
		"--filter.status", "completed",
		"--page-number", "0",
		"--page-size", "0",
	)
}

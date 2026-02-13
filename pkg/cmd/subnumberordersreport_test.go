// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSubNumberOrdersReportCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders-report", "create",
		"--country-code", "US",
		"--created-at-gt", "2023-04-05T10:22:08.230549Z",
		"--created-at-lt", "2025-06-05T10:22:08.230549Z",
		"--customer-reference", "STRING",
		"--order-request-id", "12ade33a-21c0-473b-b055-b3c836e1c293",
		"--status", "success",
	)
}

func TestSubNumberOrdersReportRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders-report", "retrieve",
		"--report-id", "12ade33a-21c0-473b-b055-b3c836e1c293",
	)
}

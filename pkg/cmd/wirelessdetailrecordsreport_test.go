// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestWirelessDetailRecordsReportsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless:detail-records-reports", "create",
		"--end-time", "2018-02-02T22:25:27.521Z",
		"--start-time", "2018-02-02T22:25:27.521Z",
	)
}

func TestWirelessDetailRecordsReportsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless:detail-records-reports", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestWirelessDetailRecordsReportsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless:detail-records-reports", "list",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestWirelessDetailRecordsReportsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless:detail-records-reports", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

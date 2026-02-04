// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestReportsMdrUsageReportsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports:mdr-usage-reports", "create",
		"--aggregation-type", "NO_AGGREGATION",
		"--end-date", "2020-07-01T00:00:00-06:00",
		"--start-date", "2020-07-01T00:00:00-06:00",
		"--profiles", "My profile",
	)
}

func TestReportsMdrUsageReportsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports:mdr-usage-reports", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestReportsMdrUsageReportsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports:mdr-usage-reports", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestReportsMdrUsageReportsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports:mdr-usage-reports", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestReportsMdrUsageReportsFetchSync(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports:mdr-usage-reports", "fetch-sync",
		"--aggregation-type", "PROFILE",
		"--end-date", "2020-07-01T00:00:00-06:00",
		"--profile", "My profile",
		"--start-date", "2020-07-01T00:00:00-06:00",
	)
}

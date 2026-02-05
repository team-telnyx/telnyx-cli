// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestReportsCdrUsageReportsFetchSync(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports:cdr-usage-reports", "fetch-sync",
		"--aggregation-type", "NO_AGGREGATION",
		"--product-breakdown", "NO_BREAKDOWN",
		"--connection", "1234567890123",
		"--end-date", "2020-07-01T00:00:00-06:00",
		"--start-date", "2020-07-01T00:00:00-06:00",
	)
}

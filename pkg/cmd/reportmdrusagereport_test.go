// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestReportsMdrUsageReportsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "reports:mdr-usage-reports", "create",
			"--api-key", "string",
			"--aggregation-type", "NO_AGGREGATION",
			"--end-date", "'2020-07-01T00:00:00-06:00'",
			"--start-date", "'2020-07-01T00:00:00-06:00'",
			"--profiles", "My profile",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aggregation_type: NO_AGGREGATION\n" +
			"end_date: '2020-07-01T00:00:00-06:00'\n" +
			"start_date: '2020-07-01T00:00:00-06:00'\n" +
			"profiles: My profile\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "reports:mdr-usage-reports", "create",
			"--api-key", "string",
		)
	})
}

func TestReportsMdrUsageReportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "reports:mdr-usage-reports", "retrieve",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestReportsMdrUsageReportsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "reports:mdr-usage-reports", "list",
			"--api-key", "string",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestReportsMdrUsageReportsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "reports:mdr-usage-reports", "delete",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestReportsMdrUsageReportsFetchSync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "reports:mdr-usage-reports", "fetch-sync",
			"--api-key", "string",
			"--aggregation-type", "PROFILE",
			"--end-date", "'2020-07-01T00:00:00-06:00'",
			"--profile", "My profile",
			"--start-date", "'2020-07-01T00:00:00-06:00'",
		)
	})
}

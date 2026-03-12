// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWirelessDetailRecordsReportsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless:detail-records-reports", "create",
			"--api-key", "string",
			"--end-time", "2018-02-02T22:25:27.521Z",
			"--start-time", "2018-02-02T22:25:27.521Z",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"end_time: '2018-02-02T22:25:27.521Z'\n" +
			"start_time: '2018-02-02T22:25:27.521Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "wireless:detail-records-reports", "create",
			"--api-key", "string",
		)
	})
}

func TestWirelessDetailRecordsReportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless:detail-records-reports", "retrieve",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestWirelessDetailRecordsReportsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless:detail-records-reports", "list",
			"--api-key", "string",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestWirelessDetailRecordsReportsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless:detail-records-reports", "delete",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

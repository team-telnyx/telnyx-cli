// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingUsageReportsMessagingCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:messaging", "create",
			"--api-key", "string",
			"--aggregation-type", "0",
			"--end-time", "'2020-01-02T00:00:00Z'",
			"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			"--profile", "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			"--profile", "7d4e3f8a-9b2c-4e1d-8f5a-1a2b3c4d5e6f",
			"--select-all-managed-accounts=true",
			"--start-time", "'2020-01-01T00:00:00Z'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aggregation_type: 0\n" +
			"end_time: '2020-01-02T00:00:00Z'\n" +
			"managed_accounts:\n" +
			"  - f47ac10b-58cc-4372-a567-0e02b2c3d479\n" +
			"  - 6ba7b810-9dad-11d1-80b4-00c04fd430c8\n" +
			"profiles:\n" +
			"  - 3fa85f64-5717-4562-b3fc-2c963f66afa6\n" +
			"  - 7d4e3f8a-9b2c-4e1d-8f5a-1a2b3c4d5e6f\n" +
			"select_all_managed_accounts: true\n" +
			"start_time: '2020-01-01T00:00:00Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legacy:reporting:usage-reports:messaging", "create",
			"--api-key", "string",
		)
	})
}

func TestLegacyReportingUsageReportsMessagingRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:messaging", "retrieve",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLegacyReportingUsageReportsMessagingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:messaging", "list",
			"--api-key", "string",
			"--page", "1",
			"--per-page", "1",
		)
	})
}

func TestLegacyReportingUsageReportsMessagingDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:messaging", "delete",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

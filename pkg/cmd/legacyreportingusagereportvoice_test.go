// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingUsageReportsVoiceCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:voice", "create",
			"--api-key", "string",
			"--end-time", "'2024-02-01T00:00:00Z'",
			"--start-time", "'2024-02-01T00:00:00Z'",
			"--aggregation-type", "0",
			"--connection", "123",
			"--connection", "456",
			"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			"--product-breakdown", "0",
			"--select-all-managed-accounts=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"end_time: '2024-02-01T00:00:00Z'\n" +
			"start_time: '2024-02-01T00:00:00Z'\n" +
			"aggregation_type: 0\n" +
			"connections:\n" +
			"  - 123\n" +
			"  - 456\n" +
			"managed_accounts:\n" +
			"  - f47ac10b-58cc-4372-a567-0e02b2c3d479\n" +
			"  - 6ba7b810-9dad-11d1-80b4-00c04fd430c8\n" +
			"product_breakdown: 0\n" +
			"select_all_managed_accounts: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legacy:reporting:usage-reports:voice", "create",
			"--api-key", "string",
		)
	})
}

func TestLegacyReportingUsageReportsVoiceRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:voice", "retrieve",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLegacyReportingUsageReportsVoiceList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:voice", "list",
			"--api-key", "string",
			"--page", "1",
			"--per-page", "1",
		)
	})
}

func TestLegacyReportingUsageReportsVoiceDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:voice", "delete",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

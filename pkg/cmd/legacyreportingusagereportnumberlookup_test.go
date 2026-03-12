// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingUsageReportsNumberLookupCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:number-lookup", "create",
			"--api-key", "string",
			"--aggregation-type", "ALL",
			"--end-date", "'2025-02-10'",
			"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			"--start-date", "'2025-02-10'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aggregationType: ALL\n" +
			"endDate: '2025-02-10'\n" +
			"managedAccounts:\n" +
			"  - f47ac10b-58cc-4372-a567-0e02b2c3d479\n" +
			"  - 6ba7b810-9dad-11d1-80b4-00c04fd430c8\n" +
			"startDate: '2025-02-10'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legacy:reporting:usage-reports:number-lookup", "create",
			"--api-key", "string",
		)
	})
}

func TestLegacyReportingUsageReportsNumberLookupRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:number-lookup", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestLegacyReportingUsageReportsNumberLookupList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:number-lookup", "list",
			"--api-key", "string",
		)
	})
}

func TestLegacyReportingUsageReportsNumberLookupDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:usage-reports:number-lookup", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

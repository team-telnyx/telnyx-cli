// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingUsageReportsNumberLookupCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:number-lookup", "create",
		"--aggregation-type", "ALL",
		"--end-date", "2025-02-10",
		"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"--start-date", "2025-02-10",
	)
}

func TestLegacyReportingUsageReportsNumberLookupRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:number-lookup", "retrieve",
		"--id", "id",
	)
}

func TestLegacyReportingUsageReportsNumberLookupList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:number-lookup", "list",
	)
}

func TestLegacyReportingUsageReportsNumberLookupDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:number-lookup", "delete",
		"--id", "id",
	)
}

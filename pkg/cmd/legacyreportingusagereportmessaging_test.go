// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingUsageReportsMessagingCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:messaging", "create",
		"--aggregation-type", "0",
		"--end-time", "2020-01-02T00:00:00Z",
		"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"--profile", "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"--profile", "7d4e3f8a-9b2c-4e1d-8f5a-1a2b3c4d5e6f",
		"--select-all-managed-accounts=true",
		"--start-time", "2020-01-01T00:00:00Z",
	)
}

func TestLegacyReportingUsageReportsMessagingRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:messaging", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestLegacyReportingUsageReportsMessagingList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:messaging", "list",
		"--page", "1",
		"--per-page", "1",
	)
}

func TestLegacyReportingUsageReportsMessagingDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:messaging", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

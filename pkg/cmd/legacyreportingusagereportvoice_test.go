// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingUsageReportsVoiceCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:voice", "create",
		"--end-time", "2024-02-01T00:00:00Z",
		"--start-time", "2024-02-01T00:00:00Z",
		"--aggregation-type", "0",
		"--connection", "123",
		"--connection", "456",
		"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"--product-breakdown", "0",
		"--select-all-managed-accounts=false",
	)
}

func TestLegacyReportingUsageReportsVoiceRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:voice", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestLegacyReportingUsageReportsVoiceList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:voice", "list",
		"--page", "1",
		"--per-page", "1",
	)
}

func TestLegacyReportingUsageReportsVoiceDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports:voice", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

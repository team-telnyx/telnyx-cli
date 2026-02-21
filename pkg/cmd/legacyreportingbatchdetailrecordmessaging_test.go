// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestLegacyReportingBatchDetailRecordsMessagingCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:messaging", "create",
		"--end-time", "2024-02-12T23:59:59Z",
		"--start-time", "2024-02-01T00:00:00Z",
		"--connection", "123",
		"--connection", "456",
		"--direction", "1",
		"--direction", "2",
		"--filter", "{billing_group: adfaa016-f921-4b6c-97bb-e4c1dad231c5, cld: '+13129457420', cld_filter: contains, cli: '+13129457420', cli_filter: contains, filter_type: and, tags_list: tag1}",
		"--include-message-body=true",
		"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"--profile", "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"--profile", "7d4e3f8a-9b2c-4e1d-8f5a-1a2b3c4d5e6f",
		"--record-type", "1",
		"--record-type", "2",
		"--report-name", "My MDR Report",
		"--select-all-managed-accounts=false",
		"--timezone", "UTC",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(legacyReportingBatchDetailRecordsMessagingCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:messaging", "create",
		"--end-time", "2024-02-12T23:59:59Z",
		"--start-time", "2024-02-01T00:00:00Z",
		"--connection", "123",
		"--connection", "456",
		"--direction", "1",
		"--direction", "2",
		"--filter.billing-group", "adfaa016-f921-4b6c-97bb-e4c1dad231c5",
		"--filter.cld", "+13129457420",
		"--filter.cld-filter", "contains",
		"--filter.cli", "+13129457420",
		"--filter.cli-filter", "contains",
		"--filter.filter-type", "and",
		"--filter.tags-list", "tag1",
		"--include-message-body=true",
		"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"--profile", "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"--profile", "7d4e3f8a-9b2c-4e1d-8f5a-1a2b3c4d5e6f",
		"--record-type", "1",
		"--record-type", "2",
		"--report-name", "My MDR Report",
		"--select-all-managed-accounts=false",
		"--timezone", "UTC",
	)
}

func TestLegacyReportingBatchDetailRecordsMessagingRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:messaging", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestLegacyReportingBatchDetailRecordsMessagingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:messaging", "list",
	)
}

func TestLegacyReportingBatchDetailRecordsMessagingDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:messaging", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

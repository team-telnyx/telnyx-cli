// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestLegacyReportingBatchDetailRecordsMessagingCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:messaging", "create",
			"--end-time", "'2024-02-12T23:59:59Z'",
			"--start-time", "'2024-02-01T00:00:00Z'",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(legacyReportingBatchDetailRecordsMessagingCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:messaging", "create",
			"--end-time", "'2024-02-12T23:59:59Z'",
			"--start-time", "'2024-02-01T00:00:00Z'",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"end_time: '2024-02-12T23:59:59Z'\n" +
			"start_time: '2024-02-01T00:00:00Z'\n" +
			"connections:\n" +
			"  - 123\n" +
			"  - 456\n" +
			"directions:\n" +
			"  - 1\n" +
			"  - 2\n" +
			"filters:\n" +
			"  - billing_group: adfaa016-f921-4b6c-97bb-e4c1dad231c5\n" +
			"    cld: '+13129457420'\n" +
			"    cld_filter: contains\n" +
			"    cli: '+13129457420'\n" +
			"    cli_filter: contains\n" +
			"    filter_type: and\n" +
			"    tags_list: tag1\n" +
			"include_message_body: true\n" +
			"managed_accounts:\n" +
			"  - f47ac10b-58cc-4372-a567-0e02b2c3d479\n" +
			"  - 6ba7b810-9dad-11d1-80b4-00c04fd430c8\n" +
			"profiles:\n" +
			"  - 3fa85f64-5717-4562-b3fc-2c963f66afa6\n" +
			"  - 7d4e3f8a-9b2c-4e1d-8f5a-1a2b3c4d5e6f\n" +
			"record_types:\n" +
			"  - 1\n" +
			"  - 2\n" +
			"report_name: My MDR Report\n" +
			"select_all_managed_accounts: false\n" +
			"timezone: UTC\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:messaging", "create",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsMessagingRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:messaging", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsMessagingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:messaging", "list",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsMessagingDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:messaging", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

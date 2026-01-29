// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestLegacyReportingBatchDetailRecordsVoiceCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:voice", "create",
		"--end-time", "2024-02-12T23:59:59Z",
		"--start-time", "2024-02-01T00:00:00Z",
		"--call-type", "1",
		"--call-type", "2",
		"--connection", "123",
		"--connection", "456",
		"--field", "call_leg_id",
		"--field", "start_time",
		"--field", "end_time",
		"--filter", "{billing_group: adfaa016-f921-4b6c-97bb-e4c1dad231c5, cld: '+13129457420', cld_filter: contains, cli: '+13129457420', cli_filter: contains, filter_type: and, tags_list: tag1}",
		"--include-all-metadata=true",
		"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"--record-type", "1",
		"--record-type", "2",
		"--report-name", "My CDR Report",
		"--select-all-managed-accounts=false",
		"--source", "calls",
		"--timezone", "UTC",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(legacyReportingBatchDetailRecordsVoiceCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:voice", "create",
		"--end-time", "2024-02-12T23:59:59Z",
		"--start-time", "2024-02-01T00:00:00Z",
		"--call-type", "1",
		"--call-type", "2",
		"--connection", "123",
		"--connection", "456",
		"--field", "call_leg_id",
		"--field", "start_time",
		"--field", "end_time",
		"--filter.billing-group", "adfaa016-f921-4b6c-97bb-e4c1dad231c5",
		"--filter.cld", "+13129457420",
		"--filter.cld-filter", "contains",
		"--filter.cli", "+13129457420",
		"--filter.cli-filter", "contains",
		"--filter.filter-type", "and",
		"--filter.tags-list", "tag1",
		"--include-all-metadata=true",
		"--managed-account", "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"--managed-account", "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"--record-type", "1",
		"--record-type", "2",
		"--report-name", "My CDR Report",
		"--select-all-managed-accounts=false",
		"--source", "calls",
		"--timezone", "UTC",
	)
}

func TestLegacyReportingBatchDetailRecordsVoiceRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:voice", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestLegacyReportingBatchDetailRecordsVoiceList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:voice", "list",
	)
}

func TestLegacyReportingBatchDetailRecordsVoiceDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:voice", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestLegacyReportingBatchDetailRecordsVoiceRetrieveFields(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:voice", "retrieve-fields",
	)
}

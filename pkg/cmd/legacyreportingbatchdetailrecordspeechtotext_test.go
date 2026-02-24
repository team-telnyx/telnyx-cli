// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingBatchDetailRecordsSpeechToTextCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:speech-to-text", "create",
		"--end-date", "2020-07-01T00:00:00-06:00",
		"--start-date", "2020-07-01T00:00:00-06:00",
	)
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:speech-to-text", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:speech-to-text", "list",
	)
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:batch-detail-records:speech-to-text", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

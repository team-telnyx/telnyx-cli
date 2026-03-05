// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingBatchDetailRecordsSpeechToTextCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:batch-detail-records:speech-to-text", "create",
			"--api-key", "string",
			"--end-date", "'2020-07-01T00:00:00-06:00'",
			"--start-date", "'2020-07-01T00:00:00-06:00'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"end_date: '2020-07-01T00:00:00-06:00'\n" +
			"start_date: '2020-07-01T00:00:00-06:00'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "legacy:reporting:batch-detail-records:speech-to-text", "create",
			"--api-key", "string",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:batch-detail-records:speech-to-text", "retrieve",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:batch-detail-records:speech-to-text", "list",
			"--api-key", "string",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "legacy:reporting:batch-detail-records:speech-to-text", "delete",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

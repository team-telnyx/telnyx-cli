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
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:speech-to-text", "create",
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
			t, pipeData,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:speech-to-text", "create",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:speech-to-text", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:speech-to-text", "list",
		)
	})
}

func TestLegacyReportingBatchDetailRecordsSpeechToTextDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"legacy:reporting:batch-detail-records:speech-to-text", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

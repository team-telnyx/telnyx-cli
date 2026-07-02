// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRecordingTranscriptionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recording-transcriptions", "retrieve",
			"--recording-transcription-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestRecordingTranscriptionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recording-transcriptions", "list",
			"--max-items", "10",
			"--filter", "{created_at: {gte: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}, recording_id: 428c31b6-7af4-4bcb-b7f5-5013ef9657c1}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(recordingTranscriptionsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recording-transcriptions", "list",
			"--max-items", "10",
			"--filter.created-at", "{gte: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}",
			"--filter.recording-id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c1",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestRecordingTranscriptionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recording-transcriptions", "delete",
			"--recording-transcription-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

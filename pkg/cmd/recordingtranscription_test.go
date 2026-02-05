// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestRecordingTranscriptionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recording-transcriptions", "retrieve",
		"--recording-transcription-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestRecordingTranscriptionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recording-transcriptions", "list",
	)
}

func TestRecordingTranscriptionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recording-transcriptions", "delete",
		"--recording-transcription-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

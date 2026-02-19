// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsTranscriptionsJsonDeleteRecordingTranscriptionSidJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:transcriptions:json", "delete-recording-transcription-sid-json",
		"--account-sid", "account_sid",
		"--recording-transcription-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestTexmlAccountsTranscriptionsJsonRetrieveRecordingTranscriptionSidJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:transcriptions:json", "retrieve-recording-transcription-sid-json",
		"--account-sid", "account_sid",
		"--recording-transcription-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

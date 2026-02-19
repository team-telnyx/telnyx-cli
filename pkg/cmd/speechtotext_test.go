// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSpeechToTextTranscribe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"speech-to-text", "transcribe",
		"--input-format", "mp3",
		"--transcription-engine", "Azure",
		"--interim-results=true",
		"--language", "language",
		"--model", "fast",
	)
}

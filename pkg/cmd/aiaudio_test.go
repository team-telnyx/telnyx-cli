// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestAIAudioTranscribe(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:audio", "transcribe",
		"--model", "distil-whisper/distil-large-v2",
		"--file", "",
		"--file-url", "https://example.com/file.mp3",
		"--response-format", "json",
		"--timestamp-granularities", "segment",
	)
}

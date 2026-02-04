// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAudioTranscribe(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:audio", "transcribe",
		"--model", "distil-whisper/distil-large-v2",
		"--file", "",
		"--file-url", "https://example.com/file.mp3",
		"--language", "en-US",
		"--model-config", "{smart_format: bar, punctuate: bar}",
		"--response-format", "json",
		"--timestamp-granularities", "segment",
	)
}

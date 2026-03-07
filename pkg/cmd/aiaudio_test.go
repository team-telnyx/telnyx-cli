// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAudioTranscribe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:audio", "transcribe",
			"--api-key", "string",
			"--model", "distil-whisper/distil-large-v2",
			"--file", "Example data",
			"--file-url", "https://example.com/file.mp3",
			"--language", "en-US",
			"--model-config", "{smart_format: bar, punctuate: bar}",
			"--response-format", "json",
			"--timestamp-granularities", "segment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"model: distil-whisper/distil-large-v2\n" +
			"file: Example data\n" +
			"file_url: https://example.com/file.mp3\n" +
			"language: en-US\n" +
			"model_config:\n" +
			"  smart_format: bar\n" +
			"  punctuate: bar\n" +
			"response_format: json\n" +
			"timestamp_granularities[]: segment\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:audio", "transcribe",
			"--api-key", "string",
		)
	})
}

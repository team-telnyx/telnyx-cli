// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAudioTranscribe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:audio", "transcribe",
			"--model", "distil-whisper/distil-large-v2",
			"--file", mocktest.TestFile(t, "Example data"),
			"--file-url", "https://example.com/file.mp3",
			"--language", "en-US",
			"--model-config", "{smart_format: bar, punctuate: bar}",
			"--response-format", "json",
			"--timestamp-granularities", "segment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"model: distil-whisper/distil-large-v2\n" +
			"file: Example data\n" +
			"file_url: https://example.com/file.mp3\n" +
			"language: en-US\n" +
			"model_config:\n" +
			"  smart_format: bar\n" +
			"  punctuate: bar\n" +
			"response_format: json\n" +
			"timestamp_granularities[]: segment\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:audio", "transcribe",
		)
	})
}

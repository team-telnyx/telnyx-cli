// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSpeechToTextListProviders(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"speech-to-text", "list-providers",
			"--provider", "deepgram",
			"--service-type", "streaming",
		)
	})
}

func TestSpeechToTextRetrieveTranscription(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"speech-to-text", "retrieve-transcription",
			"--input-format", "mp3",
			"--transcription-engine", "Azure",
			"--endpointing", "0",
			"--interim-results=true",
			"--keyterm", "keyterm",
			"--keywords", "keywords",
			"--language", "language",
			"--model", "fast",
			"--redact", "redact",
		)
	})
}

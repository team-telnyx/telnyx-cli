// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTextToSpeechListVoices(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"text-to-speech", "list-voices",
		"--elevenlabs-api-key-ref", "elevenlabs_api_key_ref",
		"--provider", "aws",
	)
}

func TestTextToSpeechStream(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"text-to-speech", "stream",
		"--audio-format", "pcm",
		"--disable-cache=true",
		"--model-id", "model_id",
		"--provider", "aws",
		"--socket-id", "socket_id",
		"--voice", "voice",
		"--voice-id", "voice_id",
	)
}

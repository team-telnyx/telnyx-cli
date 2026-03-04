// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTextToSpeechGenerate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"text-to-speech", "generate",
		"--api-key", "string",
		"--aws", "{language_code: language_code, lexicon_names: [string], output_format: output_format, sample_rate: sample_rate, text_type: text}",
		"--azure", "{api_key: api_key, deployment_id: deployment_id, effect: effect, gender: gender, language_code: language_code, output_format: output_format, region: region, text_type: text}",
		"--disable-cache=true",
		"--elevenlabs", "{api_key: api_key, language_code: language_code, voice_settings: {foo: bar}}",
		"--language", "language",
		"--minimax", "{language_boost: language_boost, pitch: 0, response_format: response_format, speed: 0, vol: 0}",
		"--output-type", "binary_output",
		"--provider", "aws",
		"--resemble", "{api_key: api_key, format: format, precision: precision, sample_rate: sample_rate}",
		"--rime", "{response_format: response_format, sampling_rate: 0, voice_speed: 0}",
		"--telnyx", "{response_format: response_format, sampling_rate: 0, temperature: 0, voice_speed: 0}",
		"--text", "text",
		"--text-type", "text",
		"--voice", "voice",
		"--voice-settings", "{foo: bar}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(textToSpeechGenerate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"text-to-speech", "generate",
		"--aws.language-code", "language_code",
		"--aws.lexicon-names", "[string]",
		"--aws.output-format", "output_format",
		"--aws.sample-rate", "sample_rate",
		"--aws.text-type", "text",
		"--azure.api-key", "api_key",
		"--azure.deployment-id", "deployment_id",
		"--azure.effect", "effect",
		"--azure.gender", "gender",
		"--azure.language-code", "language_code",
		"--azure.output-format", "output_format",
		"--azure.region", "region",
		"--azure.text-type", "text",
		"--disable-cache=true",
		"--elevenlabs.api-key", "api_key",
		"--elevenlabs.language-code", "language_code",
		"--elevenlabs.voice-settings", "{foo: bar}",
		"--language", "language",
		"--minimax.language-boost", "language_boost",
		"--minimax.pitch", "0",
		"--minimax.response-format", "response_format",
		"--minimax.speed", "0",
		"--minimax.vol", "0",
		"--output-type", "binary_output",
		"--provider", "aws",
		"--resemble.api-key", "api_key",
		"--resemble.format", "format",
		"--resemble.precision", "precision",
		"--resemble.sample-rate", "sample_rate",
		"--rime.response-format", "response_format",
		"--rime.sampling-rate", "0",
		"--rime.voice-speed", "0",
		"--telnyx.response-format", "response_format",
		"--telnyx.sampling-rate", "0",
		"--telnyx.temperature", "0",
		"--telnyx.voice-speed", "0",
		"--text", "text",
		"--text-type", "text",
		"--voice", "voice",
		"--voice-settings", "{foo: bar}",
	)
}

func TestTextToSpeechListVoices(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"text-to-speech", "list-voices",
		"--api-key", "string",
		"--api-key", "api_key",
		"--provider", "aws",
	)
}

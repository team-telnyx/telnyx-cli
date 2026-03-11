// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTextToSpeechGenerate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "text-to-speech", "generate",
			"--api-key", "string",
			"--aws", "{language_code: language_code, lexicon_names: [string], output_format: output_format, sample_rate: sample_rate, text_type: text}",
			"--azure", "{api_key: api_key, deployment_id: deployment_id, effect: effect, gender: gender, language_code: language_code, output_format: output_format, region: region, text_type: text}",
			"--disable-cache=true",
			"--elevenlabs", "{api_key: api_key, language_code: language_code, voice_settings: {foo: bar}}",
			"--inworld", "{foo: bar}",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(textToSpeechGenerate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "text-to-speech", "generate",
			"--api-key", "string",
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
			"--inworld", "{foo: bar}",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aws:\n" +
			"  language_code: language_code\n" +
			"  lexicon_names:\n" +
			"    - string\n" +
			"  output_format: output_format\n" +
			"  sample_rate: sample_rate\n" +
			"  text_type: text\n" +
			"azure:\n" +
			"  api_key: api_key\n" +
			"  deployment_id: deployment_id\n" +
			"  effect: effect\n" +
			"  gender: gender\n" +
			"  language_code: language_code\n" +
			"  output_format: output_format\n" +
			"  region: region\n" +
			"  text_type: text\n" +
			"disable_cache: true\n" +
			"elevenlabs:\n" +
			"  api_key: api_key\n" +
			"  language_code: language_code\n" +
			"  voice_settings:\n" +
			"    foo: bar\n" +
			"inworld:\n" +
			"  foo: bar\n" +
			"language: language\n" +
			"minimax:\n" +
			"  language_boost: language_boost\n" +
			"  pitch: 0\n" +
			"  response_format: response_format\n" +
			"  speed: 0\n" +
			"  vol: 0\n" +
			"output_type: binary_output\n" +
			"provider: aws\n" +
			"resemble:\n" +
			"  api_key: api_key\n" +
			"  format: format\n" +
			"  precision: precision\n" +
			"  sample_rate: sample_rate\n" +
			"rime:\n" +
			"  response_format: response_format\n" +
			"  sampling_rate: 0\n" +
			"  voice_speed: 0\n" +
			"telnyx:\n" +
			"  response_format: response_format\n" +
			"  sampling_rate: 0\n" +
			"  temperature: 0\n" +
			"  voice_speed: 0\n" +
			"text: text\n" +
			"text_type: text\n" +
			"voice: voice\n" +
			"voice_settings:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "text-to-speech", "generate",
			"--api-key", "string",
		)
	})
}

func TestTextToSpeechListVoices(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "text-to-speech", "list-voices",
			"--api-key", "string",
			"--api-key", "api_key",
			"--provider", "aws",
		)
	})
}

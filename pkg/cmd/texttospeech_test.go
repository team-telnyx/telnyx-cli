// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTextToSpeechGenerateSpeech(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"text-to-speech", "generate-speech",
			"--aws", "{language_code: string, lexicon_names: [string], output_format: string, sample_rate: string, text_type: text}",
			"--azure", "{api_key: string, deployment_id: string, effect: string, gender: string, language_code: en-US, output_format: audio-24khz-160kbitrate-mono-mp3, region: string, text_type: text}",
			"--disable-cache=false",
			"--elevenlabs", "{api_key: string, language_code: string, voice_settings: {foo: bar}}",
			"--humain", "{voice_id: sara-en, ttfb_eagerness: 0}",
			"--language", "string",
			"--minimax", "{language_boost: string, pitch: 0, response_format: string, speed: 0, vol: 0}",
			"--output-type", "binary_output",
			"--provider", "aws",
			"--resemble", "{api_key: string, format: string, precision: string, sample_rate: string}",
			"--telnyx", "{emotion: neutral, response_format: mp3, sampling_rate: 24000, voice_speed: 1, volume: 1}",
			"--text", "string",
			"--text-type", "text",
			"--voice", "string",
			"--voice-settings", "{foo: bar}",
			"--xai", "{voice_id: eve, language: auto, output_format: mp3, sample_rate: 24000}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(textToSpeechGenerateSpeech)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"text-to-speech", "generate-speech",
			"--aws.language-code", "string",
			"--aws.lexicon-names", "[string]",
			"--aws.output-format", "string",
			"--aws.sample-rate", "string",
			"--aws.text-type", "text",
			"--azure.api-key", "string",
			"--azure.deployment-id", "string",
			"--azure.effect", "string",
			"--azure.gender", "string",
			"--azure.language-code", "en-US",
			"--azure.output-format", "audio-24khz-160kbitrate-mono-mp3",
			"--azure.region", "string",
			"--azure.text-type", "text",
			"--disable-cache=false",
			"--elevenlabs.api-key", "string",
			"--elevenlabs.language-code", "string",
			"--elevenlabs.voice-settings", "{foo: bar}",
			"--humain.voice-id", "sara-en",
			"--humain.ttfb-eagerness", "0",
			"--language", "string",
			"--minimax.language-boost", "string",
			"--minimax.pitch", "0",
			"--minimax.response-format", "string",
			"--minimax.speed", "0",
			"--minimax.vol", "0",
			"--output-type", "binary_output",
			"--provider", "aws",
			"--resemble.api-key", "string",
			"--resemble.format", "string",
			"--resemble.precision", "string",
			"--resemble.sample-rate", "string",
			"--telnyx.emotion", "neutral",
			"--telnyx.response-format", "mp3",
			"--telnyx.sampling-rate", "24000",
			"--telnyx.voice-speed", "1",
			"--telnyx.volume", "1",
			"--text", "string",
			"--text-type", "text",
			"--voice", "string",
			"--voice-settings", "{foo: bar}",
			"--xai.voice-id", "eve",
			"--xai.language", "auto",
			"--xai.output-format", "mp3",
			"--xai.sample-rate", "24000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aws:\n" +
			"  language_code: string\n" +
			"  lexicon_names:\n" +
			"    - string\n" +
			"  output_format: string\n" +
			"  sample_rate: string\n" +
			"  text_type: text\n" +
			"azure:\n" +
			"  api_key: string\n" +
			"  deployment_id: string\n" +
			"  effect: string\n" +
			"  gender: string\n" +
			"  language_code: en-US\n" +
			"  output_format: audio-24khz-160kbitrate-mono-mp3\n" +
			"  region: string\n" +
			"  text_type: text\n" +
			"disable_cache: false\n" +
			"elevenlabs:\n" +
			"  api_key: string\n" +
			"  language_code: string\n" +
			"  voice_settings:\n" +
			"    foo: bar\n" +
			"humain:\n" +
			"  voice_id: sara-en\n" +
			"  ttfb_eagerness: 0\n" +
			"language: string\n" +
			"minimax:\n" +
			"  language_boost: string\n" +
			"  pitch: 0\n" +
			"  response_format: string\n" +
			"  speed: 0\n" +
			"  vol: 0\n" +
			"output_type: binary_output\n" +
			"provider: aws\n" +
			"resemble:\n" +
			"  api_key: string\n" +
			"  format: string\n" +
			"  precision: string\n" +
			"  sample_rate: string\n" +
			"telnyx:\n" +
			"  emotion: neutral\n" +
			"  response_format: mp3\n" +
			"  sampling_rate: 24000\n" +
			"  voice_speed: 1\n" +
			"  volume: 1\n" +
			"text: string\n" +
			"text_type: text\n" +
			"voice: string\n" +
			"voice_settings:\n" +
			"  foo: bar\n" +
			"xai:\n" +
			"  voice_id: eve\n" +
			"  language: auto\n" +
			"  output_format: mp3\n" +
			"  sample_rate: 24000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"text-to-speech", "generate-speech",
		)
	})
}

func TestTextToSpeechListVoices(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"text-to-speech", "list-voices",
			"--api-key", "api_key",
			"--provider", "aws",
		)
	})
}

func TestTextToSpeechRetrieveSpeech(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"text-to-speech", "retrieve-speech",
			"--audio-format", "pcm",
			"--disable-cache=true",
			"--model-id", "model_id",
			"--provider", "aws",
			"--socket-id", "socket_id",
			"--voice", "voice",
			"--voice-id", "voice_id",
		)
	})
}

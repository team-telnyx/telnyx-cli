// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var textToSpeechGenerate = requestflag.WithInnerFlags(cli.Command{
	Name:    "generate",
	Usage:   "Generate synthesized speech audio from text input. Returns audio in the\nrequested format (binary audio stream, base64-encoded JSON, or an audio URL for\nlater retrieval).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "aws",
			Usage:    "AWS Polly provider-specific parameters.",
			BodyPath: "aws",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "azure",
			Usage:    "Azure Cognitive Services provider-specific parameters.",
			BodyPath: "azure",
		},
		&requestflag.Flag[bool]{
			Name:     "disable-cache",
			Usage:    "When `true`, bypass the audio cache and generate fresh audio.",
			Default:  false,
			BodyPath: "disable_cache",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "elevenlabs",
			Usage:    "ElevenLabs provider-specific parameters.",
			BodyPath: "elevenlabs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inworld",
			Usage:    "Inworld provider-specific parameters.",
			BodyPath: "inworld",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "Language code (e.g. `en-US`). Usage varies by provider.",
			BodyPath: "language",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "minimax",
			Usage:    "Minimax provider-specific parameters.",
			BodyPath: "minimax",
		},
		&requestflag.Flag[string]{
			Name:     "output-type",
			Usage:    "Determines the response format. `binary_output` returns raw audio bytes, `base64_output` returns base64-encoded audio in JSON.",
			Default:  "binary_output",
			BodyPath: "output_type",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "TTS provider. Required unless `voice` is provided.",
			BodyPath: "provider",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "resemble",
			Usage:    "Resemble AI provider-specific parameters.",
			BodyPath: "resemble",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "rime",
			Usage:    "Rime provider-specific parameters.",
			BodyPath: "rime",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "telnyx",
			Usage:    "Telnyx provider-specific parameters.",
			BodyPath: "telnyx",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "The text to convert to speech.",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "text-type",
			Usage:    "Text type. Use `ssml` for SSML-formatted input (supported by AWS and Azure).",
			BodyPath: "text_type",
		},
		&requestflag.Flag[string]{
			Name:     "voice",
			Usage:    "Voice identifier in the format `provider.model_id.voice_id` or `provider.voice_id`. Examples: `telnyx.NaturalHD.Alloy`, `azure.en-US-AvaMultilingualNeural`, `aws.Polly.Generative.Lucia`. When provided, `provider`, `model_id`, and `voice_id` are extracted automatically and take precedence over individual parameters.",
			BodyPath: "voice",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "voice-settings",
			Usage:    "Provider-specific voice settings. Contents vary by provider — see provider-specific parameter objects below.",
			BodyPath: "voice_settings",
		},
	},
	Action:          handleTextToSpeechGenerate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"aws": {
		&requestflag.InnerFlag[string]{
			Name:       "aws.language-code",
			Usage:      "Language code (e.g. `en-US`, `es-ES`).",
			InnerField: "language_code",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "aws.lexicon-names",
			Usage:      "List of lexicon names to apply.",
			InnerField: "lexicon_names",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aws.output-format",
			Usage:      "Audio output format.",
			InnerField: "output_format",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aws.sample-rate",
			Usage:      "Audio sample rate.",
			InnerField: "sample_rate",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aws.text-type",
			Usage:      "Input text type.",
			InnerField: "text_type",
		},
	},
	"azure": {
		&requestflag.InnerFlag[string]{
			Name:       "azure.api-key",
			Usage:      "Custom Azure API key. If not provided, the default Telnyx key is used.",
			InnerField: "api_key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "azure.deployment-id",
			Usage:      "Custom Azure deployment ID.",
			InnerField: "deployment_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "azure.effect",
			Usage:      "Azure audio effect to apply.",
			InnerField: "effect",
		},
		&requestflag.InnerFlag[string]{
			Name:       "azure.gender",
			Usage:      "Voice gender preference.",
			InnerField: "gender",
		},
		&requestflag.InnerFlag[string]{
			Name:       "azure.language-code",
			Usage:      "Language code (e.g. `en-US`).",
			InnerField: "language_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "azure.output-format",
			Usage:      "Azure audio output format.",
			InnerField: "output_format",
		},
		&requestflag.InnerFlag[string]{
			Name:       "azure.region",
			Usage:      "Azure region (e.g. `eastus`, `westeurope`).",
			InnerField: "region",
		},
		&requestflag.InnerFlag[string]{
			Name:       "azure.text-type",
			Usage:      "Input text type. Use `ssml` for SSML-formatted input.",
			InnerField: "text_type",
		},
	},
	"elevenlabs": {
		&requestflag.InnerFlag[string]{
			Name:       "elevenlabs.api-key",
			Usage:      "Custom ElevenLabs API key. If not provided, the default Telnyx key is used.",
			InnerField: "api_key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "elevenlabs.language-code",
			Usage:      "Language code.",
			InnerField: "language_code",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "elevenlabs.voice-settings",
			Usage:      "ElevenLabs voice settings (stability, similarity_boost, etc.).",
			InnerField: "voice_settings",
		},
	},
	"minimax": {
		&requestflag.InnerFlag[string]{
			Name:       "minimax.language-boost",
			Usage:      "Language code to boost pronunciation for.",
			InnerField: "language_boost",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "minimax.pitch",
			Usage:      "Pitch adjustment.",
			InnerField: "pitch",
		},
		&requestflag.InnerFlag[string]{
			Name:       "minimax.response-format",
			Usage:      "Audio output format.",
			InnerField: "response_format",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "minimax.speed",
			Usage:      "Speech speed multiplier.",
			InnerField: "speed",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "minimax.vol",
			Usage:      "Volume level.",
			InnerField: "vol",
		},
	},
	"resemble": {
		&requestflag.InnerFlag[string]{
			Name:       "resemble.api-key",
			Usage:      "Custom Resemble API key.",
			InnerField: "api_key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "resemble.format",
			Usage:      "Audio output format.",
			InnerField: "format",
		},
		&requestflag.InnerFlag[string]{
			Name:       "resemble.precision",
			Usage:      "Synthesis precision.",
			InnerField: "precision",
		},
		&requestflag.InnerFlag[string]{
			Name:       "resemble.sample-rate",
			Usage:      "Audio sample rate.",
			InnerField: "sample_rate",
		},
	},
	"rime": {
		&requestflag.InnerFlag[string]{
			Name:       "rime.response-format",
			Usage:      "Audio output format.",
			InnerField: "response_format",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "rime.sampling-rate",
			Usage:      "Audio sampling rate in Hz.",
			InnerField: "sampling_rate",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "rime.voice-speed",
			Usage:      "Voice speed multiplier.",
			InnerField: "voice_speed",
		},
	},
	"telnyx": {
		&requestflag.InnerFlag[string]{
			Name:       "telnyx.response-format",
			Usage:      "Audio response format.",
			InnerField: "response_format",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "telnyx.sampling-rate",
			Usage:      "Audio sampling rate in Hz.",
			InnerField: "sampling_rate",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "telnyx.temperature",
			Usage:      "Sampling temperature.",
			InnerField: "temperature",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "telnyx.voice-speed",
			Usage:      "Voice speed multiplier.",
			InnerField: "voice_speed",
		},
	},
})

var textToSpeechListVoices = cli.Command{
	Name:    "list-voices",
	Usage:   "Retrieve a list of available voices from one or all TTS providers. When\n`provider` is specified, returns voices for that provider only. Otherwise,\nreturns voices from all providers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "api-key",
			Usage:     "API key for providers that require one to list voices (e.g. ElevenLabs).",
			QueryPath: "api_key",
		},
		&requestflag.Flag[string]{
			Name:      "provider",
			Usage:     "Filter voices by provider. If omitted, voices from all providers are returned.",
			QueryPath: "provider",
		},
	},
	Action:          handleTextToSpeechListVoices,
	HideHelpCommand: true,
}

var textToSpeechStream = cli.Command{
	Name:    "stream",
	Usage:   "Open a WebSocket connection to stream text and receive synthesized audio in real\ntime. Authentication is provided via the standard\n`Authorization: Bearer <API_KEY>` header. Send JSON frames with text to\nsynthesize; receive JSON frames containing base64-encoded audio chunks.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "audio-format",
			Usage:     "Audio output format override. Supported for Telnyx `Natural`/`NaturalHD` models only. Accepted values: `pcm`, `wav`.",
			QueryPath: "audio_format",
		},
		&requestflag.Flag[bool]{
			Name:      "disable-cache",
			Usage:     "When `true`, bypass the audio cache and generate fresh audio.",
			Default:   false,
			QueryPath: "disable_cache",
		},
		&requestflag.Flag[string]{
			Name:      "model-id",
			Usage:     "Model identifier for the chosen provider. Examples: `Natural`, `NaturalHD` (Telnyx); `Polly.Generative` (AWS).",
			QueryPath: "model_id",
		},
		&requestflag.Flag[string]{
			Name:      "provider",
			Usage:     "TTS provider. Defaults to `telnyx` if not specified. Ignored when `voice` is provided.",
			Default:   "telnyx",
			QueryPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:      "socket-id",
			Usage:     "Client-provided socket identifier for tracking. If not provided, one is generated server-side.",
			QueryPath: "socket_id",
		},
		&requestflag.Flag[string]{
			Name:      "voice",
			Usage:     "Voice identifier in the format `provider.model_id.voice_id` or `provider.voice_id` (e.g. `telnyx.NaturalHD.Telnyx_Alloy` or `azure.en-US-AvaMultilingualNeural`). When provided, the `provider`, `model_id`, and `voice_id` are extracted automatically. Takes precedence over individual `provider`/`model_id`/`voice_id` parameters.",
			QueryPath: "voice",
		},
		&requestflag.Flag[string]{
			Name:      "voice-id",
			Usage:     "Voice identifier for the chosen provider.",
			QueryPath: "voice_id",
		},
	},
	Action:          handleTextToSpeechStream,
	HideHelpCommand: true,
}

func handleTextToSpeechGenerate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TextToSpeechGenerateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TextToSpeech.Generate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "text-to-speech generate", obj, format, transform)
}

func handleTextToSpeechListVoices(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TextToSpeechListVoicesParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TextToSpeech.ListVoices(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "text-to-speech list-voices", obj, format, transform)
}

func handleTextToSpeechStream(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TextToSpeechStreamParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.TextToSpeech.Stream(ctx, params, options...)
}

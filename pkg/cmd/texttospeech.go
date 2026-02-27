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

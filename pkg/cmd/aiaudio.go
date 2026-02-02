// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiAudioTranscribe = cli.Command{
	Name:    "transcribe",
	Usage:   "Transcribe speech to text. This endpoint is consistent with the\n[OpenAI Transcription API](https://platform.openai.com/docs/api-reference/audio/createTranscription)\nand may be used with the OpenAI JS or Python SDK.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "ID of the model to use. `distil-whisper/distil-large-v2` is lower latency but English-only. `openai/whisper-large-v3-turbo` is multi-lingual but slightly higher latency.",
			Default:  "distil-whisper/distil-large-v2",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "file",
			Usage:    "The audio file object to transcribe, in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm. File uploads are limited to 100 MB. Cannot be used together with `file_url`",
			BodyPath: "file",
		},
		&requestflag.Flag[string]{
			Name:     "file-url",
			Usage:    "Link to audio file in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm. Support for hosted files is limited to 100MB. Cannot be used together with `file`",
			BodyPath: "file_url",
		},
		&requestflag.Flag[string]{
			Name:     "response-format",
			Usage:    "The format of the transcript output. Use `verbose_json` to take advantage of timestamps.",
			Default:  "json",
			BodyPath: "response_format",
		},
		&requestflag.Flag[string]{
			Name:     "timestamp-granularities",
			Usage:    "The timestamp granularities to populate for this transcription. `response_format` must be set verbose_json to use timestamp granularities. Currently `segment` is supported.",
			BodyPath: "timestamp_granularities[]",
		},
	},
	Action:          handleAIAudioTranscribe,
	HideHelpCommand: true,
}

func handleAIAudioTranscribe(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAudioTranscribeParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Audio.Transcribe(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:audio transcribe", obj, format, transform)
}

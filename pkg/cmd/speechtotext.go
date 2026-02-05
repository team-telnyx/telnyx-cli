// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/urfave/cli/v3"
)

var speechToTextTranscribe = cli.Command{
	Name:    "transcribe",
	Usage:   "Transcribe audio streams to text over WebSocket.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "input-format",
			Usage:     "The format of input audio stream.",
			Required:  true,
			QueryPath: "input_format",
		},
		&requestflag.Flag[string]{
			Name:      "transcription-engine",
			Usage:     "The transcription engine to use for processing the audio stream.",
			Required:  true,
			QueryPath: "transcription_engine",
		},
		&requestflag.Flag[bool]{
			Name:      "interim-results",
			Usage:     "Whether to receive interim transcription results.",
			QueryPath: "interim_results",
		},
		&requestflag.Flag[string]{
			Name:      "language",
			Usage:     "The language spoken in the audio stream.",
			QueryPath: "language",
		},
		&requestflag.Flag[string]{
			Name:      "model",
			Usage:     "The specific model to use within the selected transcription engine.",
			QueryPath: "model",
		},
	},
	Action:          handleSpeechToTextTranscribe,
	HideHelpCommand: true,
}

func handleSpeechToTextTranscribe(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SpeechToTextTranscribeParams{}

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

	return client.SpeechToText.Transcribe(ctx, params, options...)
}

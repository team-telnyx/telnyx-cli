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
	Usage:   "Open a WebSocket connection to stream audio and receive transcriptions in\nreal-time. Authentication is provided via the standard\n`Authorization: Bearer <API_KEY>` header.",
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
		&requestflag.Flag[int64]{
			Name:      "endpointing",
			Usage:     "Silence duration (in milliseconds) that triggers end-of-speech detection. When set, the engine uses this value to determine when a speaker has stopped talking. Not all engines support this parameter.",
			QueryPath: "endpointing",
		},
		&requestflag.Flag[bool]{
			Name:      "interim-results",
			Usage:     "Whether to receive interim transcription results.",
			QueryPath: "interim_results",
		},
		&requestflag.Flag[string]{
			Name:      "keyterm",
			Usage:     "A key term to boost in the transcription. The engine will be more likely to recognize this term. Can be specified multiple times for multiple terms.",
			QueryPath: "keyterm",
		},
		&requestflag.Flag[string]{
			Name:      "keywords",
			Usage:     "Comma-separated list of keywords to boost in the transcription. The engine will prioritize recognition of these words.",
			QueryPath: "keywords",
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
		&requestflag.Flag[string]{
			Name:      "redact",
			Usage:     "Enable redaction of sensitive information (e.g., PCI data, SSN) from transcription results. Supported values depend on the transcription engine.",
			QueryPath: "redact",
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var speechToTextListProviders = cli.Command{
	Name:    "list-providers",
	Usage:   "Retrieve the canonical list of supported speech-to-text providers, models,\naccepted language codes, and the service types each model supports.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "provider",
			Usage:     "Filter to entries for a specific STT provider. The enum mirrors the providers advertised across the speech-to-text spec (including `google` and `telnyx`, which are accepted as WebSocket transcription engines). A provider that has no models currently registered for any service type will return an empty `data` array rather than an error.",
			QueryPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:      "service-type",
			Usage:     "Service surface a model is available on. `ai_assistant` is the STT surface configured via Call Control voice-assistant transcription; it covers both live-streaming and non-streaming/batch models (matching the `TranscriptionConfig.model` enum on `call-control` voice assistants).",
			QueryPath: "service_type",
		},
	},
	Action:          handleSpeechToTextListProviders,
	HideHelpCommand: true,
}

var speechToTextRetrieveTranscription = cli.Command{
	Name:    "retrieve-transcription",
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
			Usage:     "Silence duration (in milliseconds) that triggers end-of-speech detection. When set, the engine uses this value to determine when a speaker has stopped talking. Supported by `xAI`, `Deepgram`, `Google`, `Speechmatics`, and `Soniox`. `Soniox` accepts values between 500 and 3000. Other engines may not support this parameter.",
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
	Action:          handleSpeechToTextRetrieveTranscription,
	HideHelpCommand: true,
}

func handleSpeechToTextListProviders(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.SpeechToTextListProvidersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SpeechToText.ListProviders(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "speech-to-text list-providers",
		Transform:      transform,
	})
}

func handleSpeechToTextRetrieveTranscription(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.SpeechToTextGetTranscriptionParams{}

	return client.SpeechToText.GetTranscription(ctx, params, options...)
}

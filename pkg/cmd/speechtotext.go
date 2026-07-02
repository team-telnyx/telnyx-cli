// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/stainless-sdks/telnyx-go/v4"
	"github.com/stainless-sdks/telnyx-go/v4/option"
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
			Usage:     "Filter to entries that support the given service type. For backward compatibility with the values that briefly shipped before the product-aligned rename, the legacy aliases `file_transcription`, `in_call_transcription`, and `ai_assistant_transcription` are silently accepted and normalized to `file_based`, `in_call`, and `ai_assistant` respectively. The response always emits the canonical (post-rename) values.",
			QueryPath: "service_type",
		},
	},
	Action:          handleSpeechToTextListProviders,
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

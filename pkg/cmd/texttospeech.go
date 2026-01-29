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

var textToSpeechListVoices = cli.Command{
	Name:    "list-voices",
	Usage:   "Returns a list of voices that can be used with the text to speech commands.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "elevenlabs-api-key-ref",
			Usage:     "Reference to your ElevenLabs API key stored in the Telnyx Portal",
			QueryPath: "elevenlabs_api_key_ref",
		},
		&requestflag.Flag[string]{
			Name:      "provider",
			Usage:     "Filter voices by provider",
			QueryPath: "provider",
		},
	},
	Action:          handleTextToSpeechListVoices,
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

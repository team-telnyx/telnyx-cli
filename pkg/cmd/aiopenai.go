// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiOpenAIListModels = cli.Command{
	Name:            "list-models",
	Usage:           "This endpoint returns a list of Open Source and OpenAI models that are available\nfor use. <br /><br /> **Note**: Model `id`'s will be in the form\n`{source}/{model_name}`. For example `openai/gpt-4` or\n`mistralai/Mistral-7B-Instruct-v0.1` consistent with HuggingFace naming\nconventions.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAIOpenAIListModels,
	HideHelpCommand: true,
}

func handleAIOpenAIListModels(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.OpenAI.ListModels(ctx, options...)
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
		Title:          "ai:openai list-models",
		Transform:      transform,
	})
}

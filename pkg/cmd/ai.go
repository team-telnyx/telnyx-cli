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

var aiRetrieveModels = cli.Command{
	Name:            "retrieve-models",
	Usage:           "This endpoint returns a list of Open Source and OpenAI models that are available\nfor use. <br /><br /> **Note**: Model `id`'s will be in the form\n`{source}/{model_name}`. For example `openai/gpt-4` or\n`mistralai/Mistral-7B-Instruct-v0.1` consistent with HuggingFace naming\nconventions.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAIRetrieveModels,
	HideHelpCommand: true,
}

var aiSummarize = cli.Command{
	Name:    "summarize",
	Usage:   "Generate a summary of a file's contents.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket",
			Usage:    "The name of the bucket that contains the file to be summarized.",
			Required: true,
			BodyPath: "bucket",
		},
		&requestflag.Flag[string]{
			Name:     "filename",
			Usage:    "The name of the file to be summarized.",
			Required: true,
			BodyPath: "filename",
		},
		&requestflag.Flag[string]{
			Name:     "system-prompt",
			Usage:    "A system prompt to guide the summary generation.",
			BodyPath: "system_prompt",
		},
	},
	Action:          handleAISummarize,
	HideHelpCommand: true,
}

func handleAIRetrieveModels(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.GetModels(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai retrieve-models", obj, format, transform)
}

func handleAISummarize(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AISummarizeParams{}

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
	_, err = client.AI.Summarize(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai summarize", obj, format, transform)
}

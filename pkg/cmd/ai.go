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

var aiCreateResponse = cli.Command{
	Name:    "create-response",
	Usage:   "**Deprecated**: Use `POST /v2/ai/openai/responses` instead. This endpoint is\ncompatible with the\n[OpenAI Responses API](https://developers.openai.com/api/reference/responses/overview)\nand may be used with the OpenAI JS or Python SDK. Response id parameter is not\nsupported at the moment. Use the `conversation` parameter with a Telnyx\nConversation ID to leverage persistent conversations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleAICreateResponse,
	HideHelpCommand: true,
}

var aiRetrieveModels = cli.Command{
	Name:            "retrieve-models",
	Usage:           "**Deprecated**: Use `GET /v2/ai/openai/models` instead.",
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

func handleAICreateResponse(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.AINewResponseParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.NewResponse(ctx, params, options...)
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
		Title:          "ai create-response",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "ai retrieve-models",
		Transform:      transform,
	})
}

func handleAISummarize(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.AISummarizeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Summarize(ctx, params, options...)
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
		Title:          "ai summarize",
		Transform:      transform,
	})
}

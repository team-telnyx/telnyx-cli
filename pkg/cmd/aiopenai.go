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

var aiOpenAICreateResponse = cli.Command{
	Name:    "create-response",
	Usage:   "Create a response using Telnyx's OpenAI-compatible Responses API. This endpoint\nis compatible with the\n[OpenAI Responses API](https://developers.openai.com/api/reference/responses/overview)\nand may be used with the OpenAI JS or Python SDK by setting the base URL to\n`https://api.telnyx.com/v2/ai/openai`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "conversation",
			Usage:    "Optional Telnyx Conversation ID from `POST /ai/conversations`. When provided, Telnyx stores this turn on that conversation and uses the conversation's prior messages as context. Reuse the same ID for subsequent turns and tool-result followups. Omit it for a non-persisted, stateless response.",
			BodyPath: "conversation",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "input",
			Usage:    "The input items for this turn, using the OpenAI Responses API input format.",
			BodyPath: "input",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "Optional system/developer instructions for the model. When used with a persisted `conversation`, send these on the first request that creates the thread; subsequent turns can rely on the stored history.",
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Model identifier to use for the response, for example `zai-org/GLM-5.1-FP8` or another model available from the Telnyx OpenAI-compatible models endpoint.",
			BodyPath: "model",
		},
		&requestflag.Flag[bool]{
			Name:     "stream",
			Usage:    "Set to `true` to stream Server-Sent Events, matching OpenAI's Responses streaming format.",
			BodyPath: "stream",
		},
	},
	Action:          handleAIOpenAICreateResponse,
	HideHelpCommand: true,
}

var aiOpenAIListModels = cli.Command{
	Name:            "list-models",
	Usage:           "Lists every model currently available to your account on Telnyx Inference,\nincluding SOTA open-source LLMs hosted on Telnyx GPUs (for example\n`moonshotai/Kimi-K2.6`, `zai-org/GLM-5.1-FP8`, and `MiniMaxAI/MiniMax-M2.7`),\nembedding models, and any fine-tuned models you have created.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAIOpenAIListModels,
	HideHelpCommand: true,
}

func handleAIOpenAICreateResponse(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("unexpected extra arguments: %v", unusedArgs)
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

	params := telnyx.AIOpenAINewResponseParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.OpenAI.NewResponse(ctx, params, options...)
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
		Title:          "ai:openai create-response",
		Transform:      transform,
	})
}

func handleAIOpenAIListModels(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("unexpected extra arguments: %v", unusedArgs)
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

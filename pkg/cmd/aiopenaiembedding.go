// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiOpenAIEmbeddingsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates an embedding vector representing the input text. This endpoint is\ncompatible with the\n[OpenAI Embeddings API](https://platform.openai.com/docs/api-reference/embeddings)\nand may be used with the OpenAI JS or Python SDK by setting the base URL to\n`https://api.telnyx.com/v2/ai/openai`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "input",
			Usage:    "Input text to embed. Can be a string or array of strings.",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "ID of the model to use. Use the List embedding models endpoint to see available models.",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[int64]{
			Name:     "dimensions",
			Usage:    "The number of dimensions the resulting output embeddings should have. Only supported in some models.",
			BodyPath: "dimensions",
		},
		&requestflag.Flag[string]{
			Name:     "encoding-format",
			Usage:    "The format to return the embeddings in.",
			Default:  "float",
			BodyPath: "encoding_format",
		},
		&requestflag.Flag[string]{
			Name:     "user",
			Usage:    "A unique identifier representing your end-user for monitoring and abuse detection.",
			BodyPath: "user",
		},
	},
	Action:          handleAIOpenAIEmbeddingsCreate,
	HideHelpCommand: true,
}

var aiOpenAIEmbeddingsListModels = cli.Command{
	Name:            "list-models",
	Usage:           "Returns a list of available embedding models. This endpoint is compatible with\nthe OpenAI Models API format.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAIOpenAIEmbeddingsListModels,
	HideHelpCommand: true,
}

func handleAIOpenAIEmbeddingsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIOpenAIEmbeddingNewParams{}

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
	_, err = client.AI.OpenAI.Embeddings.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:openai:embeddings create", obj, format, transform)
}

func handleAIOpenAIEmbeddingsListModels(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.OpenAI.Embeddings.ListModels(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:openai:embeddings list-models", obj, format, transform)
}

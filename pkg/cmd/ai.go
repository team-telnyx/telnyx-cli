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

var aiCreateResponseDeprecated = cli.Command{
	Name:    "create-response-deprecated",
	Usage:   "**Deprecated**: Use `POST /v2/ai/openai/responses` instead. This endpoint is\ncompatible with the\n[OpenAI Responses API](https://developers.openai.com/api/reference/responses/overview)\nand may be used with the OpenAI JS or Python SDK. Response id parameter is not\nsupported at the moment. Use the `conversation` parameter with a Telnyx\nConversation ID to leverage persistent conversations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleAICreateResponseDeprecated,
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

var aiSearchConversationHistories = cli.Command{
	Name:    "search-conversation-histories",
	Usage:   "Performs semantic vector search across conversation history records.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Natural language search query. The text is embedded into a 1024-dimensional vector and compared against indexed record chunks using kNN cosine similarity.",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "record-type",
			Usage:     "The type of records to search. Each record type is stored in a separate vector index.",
			Required:  true,
			QueryPath: "record_type",
		},
		&requestflag.Flag[string]{
			Name:      "filter-document-id",
			Usage:     "Filter by document identifier (exact match). Populated for knowledge_base records.",
			QueryPath: "filter[document_id]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-ingested-at-gte",
			Usage:     "Only include records ingested (chunked, embedded, and indexed) on or after this ISO 8601 timestamp.",
			QueryPath: "filter[ingested_at][gte]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-ingested-at-lte",
			Usage:     "Only include records ingested (chunked, embedded, and indexed) on or before this ISO 8601 timestamp.",
			QueryPath: "filter[ingested_at][lte]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-record-created-at-gte",
			Usage:     "Only include records whose original creation time is on or after this ISO 8601 timestamp.",
			QueryPath: "filter[record_created_at][gte]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-record-created-at-lte",
			Usage:     "Only include records whose original creation time is on or before this ISO 8601 timestamp.",
			QueryPath: "filter[record_created_at][lte]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-record-id",
			Usage:     "Filter to chunks belonging to a specific parent record (exact match).",
			QueryPath: "filter[record_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-region-in",
			Usage:     "Filter by the region stored on the record. Comma-separated to match multiple regions (USA, DEU, AUS, UAE). Distinct from the `region` parameter, which selects which cluster(s) are queried.",
			QueryPath: "filter[region][in]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-retention",
			Usage:     "Filter by retention policy (exact match). Filter-only: not returned in the response body.",
			QueryPath: "filter[retention]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-user-id",
			Usage:     "Filter to records owned by a specific user (exact match).",
			QueryPath: "filter[user_id]",
		},
		&requestflag.Flag[float64]{
			Name:      "min-score",
			Usage:     "Minimum cosine similarity score threshold (0.0 to 1.0). Results below this threshold are excluded.",
			Default:   0,
			QueryPath: "min_score",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Restrict search to a specific region's OpenSearch cluster. When omitted, all regions are queried in parallel (fan-out) and results are merged by cosine similarity score.",
			QueryPath: "region",
		},
		&requestflag.Flag[int64]{
			Name:      "top-k",
			Usage:     "Maximum number of results to return. Defaults to 20, maximum 100.",
			Default:   20,
			QueryPath: "top_k",
		},
	},
	Action:          handleAISearchConversationHistories,
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

func handleAICreateResponseDeprecated(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AINewResponseDeprecatedParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.NewResponseDeprecated(ctx, params, options...)
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
		Title:          "ai create-response-deprecated",
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

func handleAISearchConversationHistories(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AISearchConversationHistoriesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.SearchConversationHistories(ctx, params, options...)
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
		Title:          "ai search-conversation-histories",
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

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

var aiKnowledgeCollectionsRetrieveDocuments = cli.Command{
	Name:    "retrieve-documents",
	Usage:   "Runs search over the documents in a collection, ranked by relevance to `query`.\nSearches currently run `vector` retrieval (semantic similarity). The\ncollection's `retrieval_type` setting is the forward-compatible selector:\n`hybrid` (vector similarity fused with keyword matching) can be set but cannot\nbe searched yet, and `keyword` (lexical BM25 matching) is not accepted yet --\nsetting it returns 422 `unsupported_retrieval_type`. A per-request\n`retrieval_type` is accepted but ignored; `meta.retrieval_type` echoes the mode\nthat actually ran. When `query` is omitted, returns a plain catalog listing of\nthe collection's documents.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "slug",
			Required:  true,
			PathParam: "slug",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Field filters applied before ranking, using `filter[field][operator]=value`. Supported operators: `eq` (default), `in`, `gte`, `gt`, `lte`, `lt`, `contains`. Known fields: `record_type`, `record_id`, `user_id`, `record_created_at`, `ingested_at`; any other name resolves to a `metadata.<field>` filter. Example: `filter[record_id][eq]=rec_123`.",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number to return (1-based). Defaults to 1.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results per page. Defaults to 20.",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "query",
			Usage:     "Natural-language search query. When provided, the text is matched against the collection's document chunks using the collection's `retrieval_type` (vector or hybrid). When omitted, documents are returned as a plain catalog listing.",
			QueryPath: "query",
		},
		&requestflag.Flag[string]{
			Name:      "retrieval-type",
			Usage:     "Reserved; not yet functional. A value supplied here is accepted but ignored — it does not override the collection's configured strategy, and it is not echoed back. Searches run `vector` retrieval, and `meta.retrieval_type` reports the mode that actually ran. To change retrieval strategy, set it on the collection's settings subresource.",
			QueryPath: "retrieval_type",
		},
		&requestflag.Flag[string]{
			Name:      "sources",
			Usage:     "Comma-separated list of source types to restrict the search to. When omitted, all of the collection's sources are searched.",
			QueryPath: "sources",
		},
		&requestflag.Flag[int64]{
			Name:      "top-k",
			Usage:     "Maximum number of ranked results to consider. When omitted, the collection's configured `top_k` setting is used.",
			QueryPath: "top_k",
		},
	},
	Action:          handleAIKnowledgeCollectionsRetrieveDocuments,
	HideHelpCommand: true,
}

func handleAIKnowledgeCollectionsRetrieveDocuments(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("slug") && len(unusedArgs) > 0 {
		cmd.Set("slug", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := telnyx.AIKnowledgeCollectionGetDocumentsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Knowledge.Collections.GetDocuments(
		ctx,
		cmd.Value("slug").(string),
		params,
		options...,
	)
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
		Title:          "ai:knowledge:collections retrieve-documents",
		Transform:      transform,
	})
}

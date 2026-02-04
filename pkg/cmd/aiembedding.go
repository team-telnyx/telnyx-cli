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

var aiEmbeddingsCreate = cli.Command{
	Name:    "create",
	Usage:   "Perform embedding on a Telnyx Storage Bucket using the a embedding model. The\ncurrent supported file types are:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
			BodyPath: "bucket_name",
		},
		&requestflag.Flag[int64]{
			Name:     "document-chunk-overlap-size",
			Default:  512,
			BodyPath: "document_chunk_overlap_size",
		},
		&requestflag.Flag[int64]{
			Name:     "document-chunk-size",
			Default:  1024,
			BodyPath: "document_chunk_size",
		},
		&requestflag.Flag[string]{
			Name:     "embedding-model",
			Usage:    "Supported models to vectorize and embed documents.",
			Default:  "thenlper/gte-large",
			BodyPath: "embedding_model",
		},
		&requestflag.Flag[string]{
			Name:     "loader",
			Usage:    "Supported types of custom document loaders for embeddings.",
			Default:  "default",
			BodyPath: "loader",
		},
	},
	Action:          handleAIEmbeddingsCreate,
	HideHelpCommand: true,
}

var aiEmbeddingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Check the status of a current embedding task. Will be one of the following:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "task-id",
			Required: true,
		},
	},
	Action:          handleAIEmbeddingsRetrieve,
	HideHelpCommand: true,
}

var aiEmbeddingsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve tasks for the user that are either `queued`, `processing`, `failed`,\n`success` or `partial_success` based on the query string. Defaults to `queued`\nand `processing`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "status",
			Usage:     "List of task statuses i.e. `status=queued&status=processing`",
			Default:   []string{"processing", "queued"},
			QueryPath: "status",
		},
	},
	Action:          handleAIEmbeddingsList,
	HideHelpCommand: true,
}

var aiEmbeddingsSimilaritySearch = cli.Command{
	Name:    "similarity-search",
	Usage:   "Perform a similarity search on a Telnyx Storage Bucket, returning the most\nsimilar `num_docs` document chunks to the query.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
			BodyPath: "bucket_name",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[int64]{
			Name:     "num-of-docs",
			Default:  3,
			BodyPath: "num_of_docs",
		},
	},
	Action:          handleAIEmbeddingsSimilaritySearch,
	HideHelpCommand: true,
}

var aiEmbeddingsURL = cli.Command{
	Name:    "url",
	Usage:   "Embed website content from a specified URL, including child pages up to 5 levels\ndeep within the same domain. The process crawls and loads content from the main\nURL and its linked pages into a Telnyx Cloud Storage bucket. As soon as each\nwebpage is added to the bucket, its content is immediately processed for\nembeddings, that can be used for\n[similarity search](https://developers.telnyx.com/api-reference/embeddings/search-for-documents)\nand [clustering](https://developers.telnyx.com/docs/inference/clusters).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Usage:    "Name of the bucket to store the embeddings. This bucket must already exist.",
			Required: true,
			BodyPath: "bucket_name",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL of the webpage to embed",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handleAIEmbeddingsURL,
	HideHelpCommand: true,
}

func handleAIEmbeddingsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIEmbeddingNewParams{}

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
	_, err = client.AI.Embeddings.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:embeddings create", obj, format, transform)
}

func handleAIEmbeddingsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("task-id") && len(unusedArgs) > 0 {
		cmd.Set("task-id", unusedArgs[0])
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Embeddings.Get(ctx, cmd.Value("task-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:embeddings retrieve", obj, format, transform)
}

func handleAIEmbeddingsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIEmbeddingListParams{}

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
	_, err = client.AI.Embeddings.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:embeddings list", obj, format, transform)
}

func handleAIEmbeddingsSimilaritySearch(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIEmbeddingSimilaritySearchParams{}

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
	_, err = client.AI.Embeddings.SimilaritySearch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:embeddings similarity-search", obj, format, transform)
}

func handleAIEmbeddingsURL(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIEmbeddingURLParams{}

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
	_, err = client.AI.Embeddings.URL(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:embeddings url", obj, format, transform)
}

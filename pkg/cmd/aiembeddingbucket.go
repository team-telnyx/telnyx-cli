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

var aiEmbeddingsBucketsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get all embedded files for a given user bucket, including their processing\nstatus.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
		},
	},
	Action:          handleAIEmbeddingsBucketsRetrieve,
	HideHelpCommand: true,
}

var aiEmbeddingsBucketsList = cli.Command{
	Name:            "list",
	Usage:           "Get all embedding buckets for a user.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAIEmbeddingsBucketsList,
	HideHelpCommand: true,
}

var aiEmbeddingsBucketsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an entire bucket's embeddings and disables the bucket for AI-use,\nreturning it to normal storage pricing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
		},
	},
	Action:          handleAIEmbeddingsBucketsDelete,
	HideHelpCommand: true,
}

func handleAIEmbeddingsBucketsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-name") && len(unusedArgs) > 0 {
		cmd.Set("bucket-name", unusedArgs[0])
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
	_, err = client.AI.Embeddings.Buckets.Get(ctx, cmd.Value("bucket-name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:embeddings:buckets retrieve", obj, format, transform)
}

func handleAIEmbeddingsBucketsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Embeddings.Buckets.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:embeddings:buckets list", obj, format, transform)
}

func handleAIEmbeddingsBucketsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-name") && len(unusedArgs) > 0 {
		cmd.Set("bucket-name", unusedArgs[0])
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

	return client.AI.Embeddings.Buckets.Delete(ctx, cmd.Value("bucket-name").(string), options...)
}

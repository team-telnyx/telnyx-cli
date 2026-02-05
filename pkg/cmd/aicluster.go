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

var aiClustersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a cluster",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "task-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "show-subclusters",
			Usage:     "Whether or not to include subclusters and their nodes in the response.",
			Default:   false,
			QueryPath: "show_subclusters",
		},
		&requestflag.Flag[int64]{
			Name:      "top-n-nodes",
			Usage:     "The number of nodes in the cluster to return in the response. Nodes will be sorted by their centrality within the cluster.",
			Default:   0,
			QueryPath: "top_n_nodes",
		},
	},
	Action:          handleAIClustersRetrieve,
	HideHelpCommand: true,
}

var aiClustersList = cli.Command{
	Name:    "list",
	Usage:   "List all clusters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleAIClustersList,
	HideHelpCommand: true,
}

var aiClustersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a cluster",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "task-id",
			Required: true,
		},
	},
	Action:          handleAIClustersDelete,
	HideHelpCommand: true,
}

var aiClustersCompute = cli.Command{
	Name:    "compute",
	Usage:   "Starts a background task to compute how the data in an\n[embedded storage bucket](https://developers.telnyx.com/api-reference/embeddings/embed-documents)\nis clustered. This helps identify common themes and patterns in the data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket",
			Usage:    "The embedded storage bucket to compute the clusters from. The bucket must already be [embedded](https://developers.telnyx.com/api-reference/embeddings/embed-documents).",
			Required: true,
			BodyPath: "bucket",
		},
		&requestflag.Flag[[]string]{
			Name:     "file",
			Usage:    "Array of files to filter which are included.",
			BodyPath: "files",
		},
		&requestflag.Flag[int64]{
			Name:     "min-cluster-size",
			Usage:    "Smallest number of related text chunks to qualify as a cluster. Top-level clusters should be thought of as identifying broad themes in your data.",
			Default:  25,
			BodyPath: "min_cluster_size",
		},
		&requestflag.Flag[int64]{
			Name:     "min-subcluster-size",
			Usage:    "Smallest number of related text chunks to qualify as a sub-cluster. Sub-clusters should be thought of as identifying more specific topics within a broader theme.",
			Default:  5,
			BodyPath: "min_subcluster_size",
		},
		&requestflag.Flag[string]{
			Name:     "prefix",
			Usage:    "Prefix to filter whcih files in the buckets are included.",
			BodyPath: "prefix",
		},
	},
	Action:          handleAIClustersCompute,
	HideHelpCommand: true,
}

func handleAIClustersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("task-id") && len(unusedArgs) > 0 {
		cmd.Set("task-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIClusterGetParams{}

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
	_, err = client.AI.Clusters.Get(
		ctx,
		cmd.Value("task-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:clusters retrieve", obj, format, transform)
}

func handleAIClustersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIClusterListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Clusters.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:clusters list", obj, format, transform)
	} else {
		iter := client.AI.Clusters.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "ai:clusters list", iter, format, transform)
	}
}

func handleAIClustersDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.AI.Clusters.Delete(ctx, cmd.Value("task-id").(string), options...)
}

func handleAIClustersCompute(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIClusterComputeParams{}

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
	_, err = client.AI.Clusters.Compute(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:clusters compute", obj, format, transform)
}

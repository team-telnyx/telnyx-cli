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

var storageBucketsUsageGetAPIUsage = requestflag.WithInnerFlags(cli.Command{
	Name:    "get-api-usage",
	Usage:   "Returns the detail on API usage on a bucket of a particular time period, group\nby method category.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[start_time], filter[end_time]",
			Required:  true,
			QueryPath: "filter",
		},
	},
	Action:          handleStorageBucketsUsageGetAPIUsage,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[any]{
			Name:       "filter.end-time",
			Usage:      "The end time of the period to filter the usage (ISO microsecond format)",
			InnerField: "end_time",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.start-time",
			Usage:      "The start time of the period to filter the usage (ISO microsecond format)",
			InnerField: "start_time",
		},
	},
})

var storageBucketsUsageGetBucketUsage = cli.Command{
	Name:    "get-bucket-usage",
	Usage:   "Returns the amount of storage space and number of files a bucket takes up.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
		},
	},
	Action:          handleStorageBucketsUsageGetBucketUsage,
	HideHelpCommand: true,
}

func handleStorageBucketsUsageGetAPIUsage(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-name") && len(unusedArgs) > 0 {
		cmd.Set("bucket-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.StorageBucketUsageGetAPIUsageParams{}

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
	_, err = client.Storage.Buckets.Usage.GetAPIUsage(
		ctx,
		cmd.Value("bucket-name").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:buckets:usage get-api-usage", obj, format, transform)
}

func handleStorageBucketsUsageGetBucketUsage(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Storage.Buckets.Usage.GetBucketUsage(ctx, cmd.Value("bucket-name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:buckets:usage get-bucket-usage", obj, format, transform)
}

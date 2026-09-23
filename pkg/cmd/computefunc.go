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

var computeFuncsRetrieveLogs = cli.Command{
	Name:    "retrieve-logs",
	Usage:   "Returns logs oldest first. `type=runtime` (default) returns function\nstdout/stderr. `type=invocations` returns one platform-generated record per HTTP\nrequest served.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:      "end-time",
			Usage:     "Return records at or before this RFC 3339 timestamp.",
			QueryPath: "end_time",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum records to return.",
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "start-time",
			Usage:     "Return records at or after this RFC 3339 timestamp.",
			QueryPath: "start_time",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Log stream to return.",
			Default:   "runtime",
			QueryPath: "type",
		},
	},
	Action:          handleComputeFuncsRetrieveLogs,
	HideHelpCommand: true,
}

var computeFuncsRetrieveMetricAggregates = cli.Command{
	Name:    "retrieve-metric-aggregates",
	Usage:   "Returns aggregate request, latency, CPU, memory, and resource-limit metrics for\na function over the requested window.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:      "end-time",
			Usage:     "Exclusive window end, UTC ISO 8601 with milliseconds",
			Required:  true,
			QueryPath: "end_time",
		},
		&requestflag.Flag[any]{
			Name:      "start-time",
			Usage:     "Inclusive window start, UTC ISO 8601 with milliseconds",
			Required:  true,
			QueryPath: "start_time",
		},
		&requestflag.Flag[string]{
			Name:      "filter-edge-site",
			Usage:     "Edge site filter",
			QueryPath: "filter[edge_site]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-namespace",
			Usage:     "Kubernetes namespace filter",
			QueryPath: "filter[namespace]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   20,
			QueryPath: "page[size]",
		},
	},
	Action:          handleComputeFuncsRetrieveMetricAggregates,
	HideHelpCommand: true,
}

var computeFuncsRetrieveRevisions = cli.Command{
	Name:    "retrieve-revisions",
	Usage:   "Lists a function's ship history newest first, including per-ship failure stage\nand reason when recorded.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   10,
			QueryPath: "page[size]",
		},
	},
	Action:          handleComputeFuncsRetrieveRevisions,
	HideHelpCommand: true,
}

var computeFuncsRetrieveShipInspection = cli.Command{
	Name:    "retrieve-ship-inspection",
	Usage:   "Returns the latest ship outcome. The stage is `none` on success, `pending` while\nbuilding, or a failure stage such as `build`, `platform`, `pre_build`, `deploy`,\nor `security_review`. This stage-neutral customer-facing path is an alias over\nthe same inspection resource as `build_log_inspection`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleComputeFuncsRetrieveShipInspection,
	HideHelpCommand: true,
}

func handleComputeFuncsRetrieveLogs(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	params := telnyx.ComputeFuncGetLogsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.Funcs.GetLogs(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "compute:funcs retrieve-logs",
		Transform:      transform,
	})
}

func handleComputeFuncsRetrieveMetricAggregates(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	params := telnyx.ComputeFuncGetMetricAggregatesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.Funcs.GetMetricAggregates(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "compute:funcs retrieve-metric-aggregates",
		Transform:      transform,
	})
}

func handleComputeFuncsRetrieveRevisions(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	params := telnyx.ComputeFuncGetRevisionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.Funcs.GetRevisions(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "compute:funcs retrieve-revisions",
		Transform:      transform,
	})
}

func handleComputeFuncsRetrieveShipInspection(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Compute.Funcs.GetShipInspection(ctx, cmd.Value("id").(string), options...)
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
		Title:          "compute:funcs retrieve-ship-inspection",
		Transform:      transform,
	})
}

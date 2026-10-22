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

var dialMetricsList = cli.Command{
	Name:    "list",
	Usage:   "Returns dials-per-second (DPS) metrics for your account's outbound dials made\nwith the Call Control `Dial` command and TeXML, as time buckets with dial\ncounts, accepted and rejected dials, average and peak DPS, and the DPS limit in\nforce.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "end-time",
			Usage:     "Exclusive end of the range, ISO 8601. Must be after `start_time`. Buckets that start before this time are returned.",
			Required:  true,
			QueryPath: "end_time",
		},
		&requestflag.Flag[string]{
			Name:      "granularity",
			Usage:     "Bucket size. `second` returns 5-second buckets (the finest resolution) and is kept for the last 48 hours. `minute` and `hour` are kept for the last 30 days.",
			Required:  true,
			QueryPath: "granularity",
		},
		&requestflag.Flag[any]{
			Name:      "start-time",
			Usage:     "Inclusive start of the range, ISO 8601 (for example `2026-09-29T12:00:00.000Z`). Buckets that start at or after this time are returned. A `start_time` older than the retention window of the chosen `granularity` is moved forward to the start of that window, and `meta.start_time` returns the start actually used.",
			Required:  true,
			QueryPath: "start_time",
		},
		&requestflag.Flag[string]{
			Name:      "page-after",
			Usage:     "Opaque cursor from `meta.cursors.after` of the previous page. Do not parse or construct it. A cursor that matches no records returns an empty page whose `meta.cursors.after` and `meta.cursors.before` both echo the given cursor.",
			QueryPath: "page[after]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-limit",
			Usage:     "Maximum number of buckets per page. Values outside the range are clamped to 1-10,000.",
			Default:   1000,
			QueryPath: "page[limit]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sort order of the buckets by `start_time`. `-start_time` (default) returns the newest first, `start_time` the oldest first. `region` follows the same direction for buckets with the same `start_time`. To read the live rate, use `granularity=second` and `sort=-start_time`, then take the first record of each `region`.",
			Default:   "-start_time",
			QueryPath: "sort",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleDialMetricsList,
	HideHelpCommand: true,
}

func handleDialMetricsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.DialMetricListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.DialMetrics.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "dial-metrics list",
			Transform:      transform,
		})
	} else {
		iter := client.DialMetrics.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "dial-metrics list",
			Transform:      transform,
		})
	}
}

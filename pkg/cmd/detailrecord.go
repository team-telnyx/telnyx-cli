// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var detailRecordsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Search for any detail record across the Telnyx Platform",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Filter records on a given record attribute and value. <br/>Example: filter[status]=delivered. <br/>Required: filter[record_type] must be specified.",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[[]string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. <br/>Example: sort=-created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleDetailRecordsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.record-type",
			Usage:      "Filter by the given record type.",
			InnerField: "record_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.date-range",
			Usage:      "Filter by the given user-friendly date range. You can specify one of the following enum values, or a dynamic one using this format: last_N_days.",
			InnerField: "date_range",
		},
	},
})

func handleDetailRecordsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DetailRecordListParams{}

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
		_, err = client.DetailRecords.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "detail-records list", obj, format, transform)
	} else {
		iter := client.DetailRecords.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "detail-records list", iter, format, transform)
	}
}

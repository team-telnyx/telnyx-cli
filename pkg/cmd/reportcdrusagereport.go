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

var reportsCdrUsageReportsFetchSync = cli.Command{
	Name:    "fetch-sync",
	Usage:   "Generate and fetch voice usage report synchronously. This endpoint will both\ngenerate and fetch the voice report over a specified time period. No polling is\nnecessary but the response may take up to a couple of minutes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "aggregation-type",
			Required:  true,
			QueryPath: "aggregation_type",
		},
		&requestflag.Flag[string]{
			Name:      "product-breakdown",
			Required:  true,
			QueryPath: "product_breakdown",
		},
		&requestflag.Flag[[]float64]{
			Name:      "connection",
			QueryPath: "connections",
		},
		&requestflag.Flag[any]{
			Name:      "end-date",
			QueryPath: "end_date",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			QueryPath: "start_date",
		},
	},
	Action:          handleReportsCdrUsageReportsFetchSync,
	HideHelpCommand: true,
}

func handleReportsCdrUsageReportsFetchSync(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ReportCdrUsageReportFetchSyncParams{}

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
	_, err = client.Reports.CdrUsageReports.FetchSync(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports:cdr-usage-reports fetch-sync", obj, format, transform)
}

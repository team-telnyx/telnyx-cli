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

var usageReportsList = cli.Command{
	Name:    "list",
	Usage:   "Get Telnyx usage data by product, broken out by the specified dimensions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "dimension",
			Usage:     "Breakout by specified product dimensions",
			Required:  true,
			QueryPath: "dimensions",
		},
		&requestflag.Flag[[]string]{
			Name:      "metric",
			Usage:     "Specified product usage values",
			Required:  true,
			QueryPath: "metrics",
		},
		&requestflag.Flag[string]{
			Name:      "product",
			Usage:     "Telnyx product",
			Required:  true,
			QueryPath: "product",
		},
		&requestflag.Flag[string]{
			Name:      "date-range",
			Usage:     "A more user-friendly way to specify the timespan you want to filter by. More options can be found in the Telnyx API Reference docs.",
			QueryPath: "date_range",
		},
		&requestflag.Flag[string]{
			Name:      "end-date",
			Usage:     "The end date for the time range you are interested in. The maximum time range is 31 days. Format: YYYY-MM-DDTHH:mm:ssZ",
			QueryPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:      "filter",
			Usage:     "Filter records on dimensions",
			QueryPath: "filter",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Specify the response format (csv or json). JSON is returned by default, even if not specified.",
			QueryPath: "format",
		},
		&requestflag.Flag[bool]{
			Name:      "managed-accounts",
			Usage:     "Return the aggregations for all Managed Accounts under the user making the request.",
			QueryPath: "managed_accounts",
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
			Usage:     "Specifies the sort order for results",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "start-date",
			Usage:     "The start date for the time range you are interested in. The maximum time range is 31 days. Format: YYYY-MM-DDTHH:mm:ssZ",
			QueryPath: "start_date",
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleUsageReportsList,
	HideHelpCommand: true,
}

var usageReportsGetOptions = cli.Command{
	Name:    "get-options",
	Usage:   "Get the Usage Reports options for querying usage, including the products\navailable and their respective metrics and dimensions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "product",
			Usage:     "Options (dimensions and metrics) for a given product. If none specified, all products will be returned.",
			QueryPath: "product",
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleUsageReportsGetOptions,
	HideHelpCommand: true,
}

func handleUsageReportsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.UsageReportListParams{}

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
		_, err = client.UsageReports.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "usage-reports list", obj, format, transform)
	} else {
		iter := client.UsageReports.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "usage-reports list", iter, format, transform)
	}
}

func handleUsageReportsGetOptions(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.UsageReportGetOptionsParams{}

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
	_, err = client.UsageReports.GetOptions(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "usage-reports get-options", obj, format, transform)
}

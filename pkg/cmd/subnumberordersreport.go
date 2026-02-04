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

var subNumberOrdersReportCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a CSV report for sub number orders. The report will be generated\nasynchronously and can be downloaded once complete.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "country-code",
			Usage:    "Filter by country code",
			BodyPath: "country_code",
		},
		&requestflag.Flag[any]{
			Name:     "created-at-gt",
			Usage:    "Filter for orders created after this date",
			BodyPath: "created_at_gt",
		},
		&requestflag.Flag[any]{
			Name:     "created-at-lt",
			Usage:    "Filter for orders created before this date",
			BodyPath: "created_at_lt",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "Filter by customer reference",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "order-request-id",
			Usage:    "Filter by specific order request ID",
			BodyPath: "order_request_id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Filter by order status",
			BodyPath: "status",
		},
	},
	Action:          handleSubNumberOrdersReportCreate,
	HideHelpCommand: true,
}

var subNumberOrdersReportRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get the status and details of a sub number orders report.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "report-id",
			Required: true,
		},
	},
	Action:          handleSubNumberOrdersReportRetrieve,
	HideHelpCommand: true,
}

func handleSubNumberOrdersReportCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SubNumberOrdersReportNewParams{}

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
	_, err = client.SubNumberOrdersReport.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sub-number-orders-report create", obj, format, transform)
}

func handleSubNumberOrdersReportRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("report-id") && len(unusedArgs) > 0 {
		cmd.Set("report-id", unusedArgs[0])
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
	_, err = client.SubNumberOrdersReport.Get(ctx, cmd.Value("report-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sub-number-orders-report retrieve", obj, format, transform)
}

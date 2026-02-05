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

var chargesBreakdownRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a detailed breakdown of monthly charges for phone numbers in a\nspecified date range. The date range cannot exceed 31 days.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "start-date",
			Usage:     "Start date for the charges breakdown in ISO date format (YYYY-MM-DD)",
			Required:  true,
			QueryPath: "start_date",
		},
		&requestflag.Flag[any]{
			Name:      "end-date",
			Usage:     "End date for the charges breakdown in ISO date format (YYYY-MM-DD). If not provided, defaults to start_date + 1 month. The date is exclusive, data for the end_date itself is not included in the report. The interval between start_date and end_date cannot exceed 31 days.",
			QueryPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format",
			Default:   "json",
			QueryPath: "format",
		},
	},
	Action:          handleChargesBreakdownRetrieve,
	HideHelpCommand: true,
}

func handleChargesBreakdownRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ChargesBreakdownGetParams{}

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
	_, err = client.ChargesBreakdown.Get(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "charges-breakdown retrieve", obj, format, transform)
}

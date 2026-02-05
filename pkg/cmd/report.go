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

var reportsListMdrs = cli.Command{
	Name:    "list-mdrs",
	Usage:   "Fetch all Mdr records",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Message uuid",
			QueryPath: "id",
		},
		&requestflag.Flag[string]{
			Name:      "cld",
			Usage:     "Destination number",
			QueryPath: "cld",
		},
		&requestflag.Flag[string]{
			Name:      "cli",
			Usage:     "Origination number",
			QueryPath: "cli",
		},
		&requestflag.Flag[string]{
			Name:      "direction",
			Usage:     "Direction (inbound or outbound)",
			QueryPath: "direction",
		},
		&requestflag.Flag[string]{
			Name:      "end-date",
			Usage:     "Pagination end date",
			QueryPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:      "message-type",
			Usage:     "Type of message",
			QueryPath: "message_type",
		},
		&requestflag.Flag[string]{
			Name:      "profile",
			Usage:     "Name of the profile",
			QueryPath: "profile",
		},
		&requestflag.Flag[string]{
			Name:      "start-date",
			Usage:     "Pagination start date",
			QueryPath: "start_date",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Message status",
			QueryPath: "status",
		},
	},
	Action:          handleReportsListMdrs,
	HideHelpCommand: true,
}

var reportsListWdrs = cli.Command{
	Name:    "list-wdrs",
	Usage:   "Fetch all Wdr records",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "WDR uuid",
			QueryPath: "id",
		},
		&requestflag.Flag[string]{
			Name:      "end-date",
			Usage:     "End date",
			QueryPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:      "imsi",
			Usage:     "International mobile subscriber identity",
			QueryPath: "imsi",
		},
		&requestflag.Flag[string]{
			Name:      "mcc",
			Usage:     "Mobile country code",
			QueryPath: "mcc",
		},
		&requestflag.Flag[string]{
			Name:      "mnc",
			Usage:     "Mobile network code",
			QueryPath: "mnc",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Usage:     "Phone number",
			QueryPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:      "sim-card-id",
			Usage:     "Sim card unique identifier",
			QueryPath: "sim_card_id",
		},
		&requestflag.Flag[string]{
			Name:      "sim-group-id",
			Usage:     "Sim group unique identifier",
			QueryPath: "sim_group_id",
		},
		&requestflag.Flag[string]{
			Name:      "sim-group-name",
			Usage:     "Sim group name",
			QueryPath: "sim_group_name",
		},
		&requestflag.Flag[[]string]{
			Name:      "sort",
			Usage:     "Field used to order the data. If no field is specified, default value is 'created_at'",
			Default:   []string{"created_at"},
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "start-date",
			Usage:     "Start date",
			QueryPath: "start_date",
		},
	},
	Action:          handleReportsListWdrs,
	HideHelpCommand: true,
}

func handleReportsListMdrs(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ReportListMdrsParams{}

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
	_, err = client.Reports.ListMdrs(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports list-mdrs", obj, format, transform)
}

func handleReportsListWdrs(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ReportListWdrsParams{}

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
		_, err = client.Reports.ListWdrs(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "reports list-wdrs", obj, format, transform)
	} else {
		iter := client.Reports.ListWdrsAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "reports list-wdrs", iter, format, transform)
	}
}

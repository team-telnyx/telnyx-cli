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

var portoutsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the portout request based on the ID provided",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePortoutsRetrieve,
	HideHelpCommand: true,
}

var portoutsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns the portout requests according to filters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[carrier_name], filter[country_code], filter[country_code_in], filter[foc_date], filter[inserted_at], filter[phone_number], filter[pon], filter[ported_out_at], filter[spid], filter[status], filter[status_in], filter[support_key]",
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
	},
	Action:          handlePortoutsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.carrier-name",
			Usage:      "Filter by new carrier name.",
			InnerField: "carrier_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "Filter by 2-letter country code",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.country-code-in",
			Usage:      "Filter by a list of 2-letter country codes",
			InnerField: "country_code_in",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.foc-date",
			Usage:      "Filter by foc_date. Matches all portouts with the same date",
			InnerField: "foc_date",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.inserted-at",
			Usage:      "Filter by inserted_at date range using nested operations",
			InnerField: "inserted_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number",
			Usage:      "Filter by a phone number on the portout. Matches all portouts with the phone number",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.pon",
			Usage:      "Filter by Port Order Number (PON).",
			InnerField: "pon",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.ported-out-at",
			Usage:      "Filter by ported_out_at date range using nested operations",
			InnerField: "ported_out_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.spid",
			Usage:      "Filter by new carrier spid.",
			InnerField: "spid",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter by portout status.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.status-in",
			Usage:      "Filter by a list of portout statuses",
			InnerField: "status_in",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.support-key",
			Usage:      "Filter by the portout's support_key",
			InnerField: "support_key",
		},
	},
})

var portoutsListRejectionCodes = requestflag.WithInnerFlags(cli.Command{
	Name:    "list-rejection-codes",
	Usage:   "Given a port-out ID, list rejection codes that are eligible for that port-out",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "portout-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[code], filter[code][in]",
			QueryPath: "filter",
		},
	},
	Action:          handlePortoutsListRejectionCodes,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[any]{
			Name:       "filter.code",
			Usage:      "Filter rejections of a specific code",
			InnerField: "code",
		},
	},
})

var portoutsUpdateStatus = cli.Command{
	Name:    "update-status",
	Usage:   "Authorize or reject portout request",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Provide a reason if rejecting the port out request",
			Required: true,
			BodyPath: "reason",
		},
		&requestflag.Flag[bool]{
			Name:     "host-messaging",
			Usage:    "Indicates whether messaging services should be maintained with Telnyx after the port out completes",
			Default:  false,
			BodyPath: "host_messaging",
		},
	},
	Action:          handlePortoutsUpdateStatus,
	HideHelpCommand: true,
}

func handlePortoutsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Portouts.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "portouts retrieve", obj, format, transform)
}

func handlePortoutsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortoutListParams{}

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
		_, err = client.Portouts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "portouts list", obj, format, transform)
	} else {
		iter := client.Portouts.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "portouts list", iter, format, transform)
	}
}

func handlePortoutsListRejectionCodes(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("portout-id") && len(unusedArgs) > 0 {
		cmd.Set("portout-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortoutListRejectionCodesParams{}

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
	_, err = client.Portouts.ListRejectionCodes(
		ctx,
		cmd.Value("portout-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "portouts list-rejection-codes", obj, format, transform)
}

func handlePortoutsUpdateStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("status") && len(unusedArgs) > 0 {
		cmd.Set("status", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortoutUpdateStatusParams{
		ID: cmd.Value("id").(string),
	}

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
	_, err = client.Portouts.UpdateStatus(
		ctx,
		telnyx.PortoutUpdateStatusParamsStatus(cmd.Value("status").(string)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "portouts update-status", obj, format, transform)
}

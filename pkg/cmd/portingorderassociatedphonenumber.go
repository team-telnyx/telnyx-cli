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

var portingOrdersAssociatedPhoneNumbersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new associated phone number for a porting order. This is used for\npartial porting in GB to specify which phone numbers should be kept or\ndisconnected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "action",
			Usage:    "Specifies the action to take with this phone number during partial porting.",
			Required: true,
			BodyPath: "action",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "phone-number-range",
			Required: true,
			BodyPath: "phone_number_range",
		},
	},
	Action:          handlePortingOrdersAssociatedPhoneNumbersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"phone-number-range": {
		&requestflag.InnerFlag[string]{
			Name:       "phone-number-range.end-at",
			Usage:      "Specifies the end of the phone number range for this associated phone number.",
			InnerField: "end_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phone-number-range.start-at",
			Usage:      "Specifies the start of the phone number range for this associated phone number.",
			InnerField: "start_at",
		},
	},
})

var portingOrdersAssociatedPhoneNumbersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of all associated phone numbers for a porting order. Associated\nphone numbers are used for partial porting in GB to specify which phone numbers\nshould be kept or disconnected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[phone_number], filter[action]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "sort",
			Usage:     "Consolidated sort parameter (deepObject style). Originally: sort[value]",
			QueryPath: "sort",
		},
	},
	Action:          handlePortingOrdersAssociatedPhoneNumbersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.action",
			Usage:      "Filter results by action type",
			InnerField: "action",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number",
			Usage:      "Filter results by a phone number. It should be in E.164 format.",
			InnerField: "phone_number",
		},
	},
	"page": {
		&requestflag.InnerFlag[int64]{
			Name:       "page.number",
			Usage:      "The page number to load",
			InnerField: "number",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.size",
			Usage:      "The size of the page",
			InnerField: "size",
		},
	},
	"sort": {
		&requestflag.InnerFlag[string]{
			Name:       "sort.value",
			Usage:      "Specifies the sort order for results. If not given, results are sorted by created_at in descending order",
			InnerField: "value",
		},
	},
})

var portingOrdersAssociatedPhoneNumbersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an associated phone number from a porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePortingOrdersAssociatedPhoneNumbersDelete,
	HideHelpCommand: true,
}

func handlePortingOrdersAssociatedPhoneNumbersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("porting-order-id") && len(unusedArgs) > 0 {
		cmd.Set("porting-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderAssociatedPhoneNumberNewParams{}

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
	_, err = client.PortingOrders.AssociatedPhoneNumbers.New(
		ctx,
		cmd.Value("porting-order-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "porting-orders:associated-phone-numbers create", obj, format, transform)
}

func handlePortingOrdersAssociatedPhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("porting-order-id") && len(unusedArgs) > 0 {
		cmd.Set("porting-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderAssociatedPhoneNumberListParams{}

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
		_, err = client.PortingOrders.AssociatedPhoneNumbers.List(
			ctx,
			cmd.Value("porting-order-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "porting-orders:associated-phone-numbers list", obj, format, transform)
	} else {
		iter := client.PortingOrders.AssociatedPhoneNumbers.ListAutoPaging(
			ctx,
			cmd.Value("porting-order-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "porting-orders:associated-phone-numbers list", iter, format, transform)
	}
}

func handlePortingOrdersAssociatedPhoneNumbersDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderAssociatedPhoneNumberDeleteParams{
		PortingOrderID: cmd.Value("porting-order-id").(string),
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
	_, err = client.PortingOrders.AssociatedPhoneNumbers.Delete(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "porting-orders:associated-phone-numbers delete", obj, format, transform)
}

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

var portingOrdersPhoneNumberBlocksCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new phone number block.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "activation-range",
			Usage:    "Specifies the activation ranges for this porting phone number block. The activation range must be within the block range and should not overlap with other activation ranges.",
			Required: true,
			BodyPath: "activation_ranges",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "phone-number-range",
			Required: true,
			BodyPath: "phone_number_range",
		},
	},
	Action:          handlePortingOrdersPhoneNumberBlocksCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"activation-range": {
		&requestflag.InnerFlag[string]{
			Name:       "activation-range.end-at",
			Usage:      "Specifies the end of the activation range. It must be no more than the end of the extension range.",
			InnerField: "end_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "activation-range.start-at",
			Usage:      "Specifies the start of the activation range. Must be greater or equal the start of the extension range.",
			InnerField: "start_at",
		},
	},
	"phone-number-range": {
		&requestflag.InnerFlag[string]{
			Name:       "phone-number-range.end-at",
			Usage:      "Specifies the end of the phone number range for this porting phone number block.",
			InnerField: "end_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phone-number-range.start-at",
			Usage:      "Specifies the start of the phone number range for this porting phone number block.",
			InnerField: "start_at",
		},
	},
})

var portingOrdersPhoneNumberBlocksList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of all phone number blocks of a porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[porting_order_id], filter[support_key], filter[status], filter[phone_number], filter[activation_status], filter[portability_status]",
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
		&requestflag.Flag[map[string]any]{
			Name:      "sort",
			Usage:     "Consolidated sort parameter (deepObject style). Originally: sort[value]",
			QueryPath: "sort",
		},
	},
	Action:          handlePortingOrdersPhoneNumberBlocksList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.activation-status",
			Usage:      "Filter results by activation status",
			InnerField: "activation_status",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.phone-number",
			Usage:      "Filter results by a list of phone numbers",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.portability-status",
			Usage:      "Filter results by portability status",
			InnerField: "portability_status",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.porting-order-id",
			Usage:      "Filter results by a list of porting order ids",
			InnerField: "porting_order_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.status",
			Usage:      "Filter porting orders by status(es). Originally: filter[status], filter[status][in][]",
			InnerField: "status",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.support-key",
			Usage:      "Filter results by support key(s). Originally: filter[support_key][eq], filter[support_key][in][]",
			InnerField: "support_key",
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

var portingOrdersPhoneNumberBlocksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a phone number block.",
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
	Action:          handlePortingOrdersPhoneNumberBlocksDelete,
	HideHelpCommand: true,
}

func handlePortingOrdersPhoneNumberBlocksCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("porting-order-id") && len(unusedArgs) > 0 {
		cmd.Set("porting-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderPhoneNumberBlockNewParams{}

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
	_, err = client.PortingOrders.PhoneNumberBlocks.New(
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
	return ShowJSON(os.Stdout, "porting-orders:phone-number-blocks create", obj, format, transform)
}

func handlePortingOrdersPhoneNumberBlocksList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("porting-order-id") && len(unusedArgs) > 0 {
		cmd.Set("porting-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderPhoneNumberBlockListParams{}

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
		_, err = client.PortingOrders.PhoneNumberBlocks.List(
			ctx,
			cmd.Value("porting-order-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "porting-orders:phone-number-blocks list", obj, format, transform)
	} else {
		iter := client.PortingOrders.PhoneNumberBlocks.ListAutoPaging(
			ctx,
			cmd.Value("porting-order-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "porting-orders:phone-number-blocks list", iter, format, transform)
	}
}

func handlePortingOrdersPhoneNumberBlocksDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderPhoneNumberBlockDeleteParams{
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
	_, err = client.PortingOrders.PhoneNumberBlocks.Delete(
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
	return ShowJSON(os.Stdout, "porting-orders:phone-number-blocks delete", obj, format, transform)
}

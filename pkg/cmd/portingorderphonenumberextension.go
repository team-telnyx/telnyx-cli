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

var portingOrdersPhoneNumberExtensionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new phone number extension.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "activation-range",
			Usage:    "Specifies the activation ranges for this porting phone number extension. The activation range must be within the extension range and should not overlap with other activation ranges.",
			Required: true,
			BodyPath: "activation_ranges",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "extension-range",
			Required: true,
			BodyPath: "extension_range",
		},
		&requestflag.Flag[string]{
			Name:     "porting-phone-number-id",
			Usage:    "Identifies the porting phone number associated with this porting phone number extension.",
			Required: true,
			BodyPath: "porting_phone_number_id",
		},
	},
	Action:          handlePortingOrdersPhoneNumberExtensionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"activation-range": {
		&requestflag.InnerFlag[int64]{
			Name:       "activation-range.end-at",
			Usage:      "Specifies the end of the activation range. It must be no more than the end of the extension range.",
			InnerField: "end_at",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "activation-range.start-at",
			Usage:      "Specifies the start of the activation range. Must be greater or equal the start of the extension range.",
			InnerField: "start_at",
		},
	},
	"extension-range": {
		&requestflag.InnerFlag[int64]{
			Name:       "extension-range.end-at",
			Usage:      "Specifies the end of the extension range for this porting phone number extension.",
			InnerField: "end_at",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "extension-range.start-at",
			Usage:      "Specifies the start of the extension range for this porting phone number extension.",
			InnerField: "start_at",
		},
	},
})

var portingOrdersPhoneNumberExtensionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of all phone number extensions of a porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[porting_phone_number_id]",
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
	Action:          handlePortingOrdersPhoneNumberExtensionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.porting-phone-number-id",
			Usage:      "Filter results by porting phone number id",
			InnerField: "porting_phone_number_id",
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

var portingOrdersPhoneNumberExtensionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a phone number extension.",
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
	Action:          handlePortingOrdersPhoneNumberExtensionsDelete,
	HideHelpCommand: true,
}

func handlePortingOrdersPhoneNumberExtensionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("porting-order-id") && len(unusedArgs) > 0 {
		cmd.Set("porting-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderPhoneNumberExtensionNewParams{}

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
	_, err = client.PortingOrders.PhoneNumberExtensions.New(
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
	return ShowJSON(os.Stdout, "porting-orders:phone-number-extensions create", obj, format, transform)
}

func handlePortingOrdersPhoneNumberExtensionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("porting-order-id") && len(unusedArgs) > 0 {
		cmd.Set("porting-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderPhoneNumberExtensionListParams{}

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
		_, err = client.PortingOrders.PhoneNumberExtensions.List(
			ctx,
			cmd.Value("porting-order-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "porting-orders:phone-number-extensions list", obj, format, transform)
	} else {
		iter := client.PortingOrders.PhoneNumberExtensions.ListAutoPaging(
			ctx,
			cmd.Value("porting-order-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "porting-orders:phone-number-extensions list", iter, format, transform)
	}
}

func handlePortingOrdersPhoneNumberExtensionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderPhoneNumberExtensionDeleteParams{
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
	_, err = client.PortingOrders.PhoneNumberExtensions.Delete(
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
	return ShowJSON(os.Stdout, "porting-orders:phone-number-extensions delete", obj, format, transform)
}

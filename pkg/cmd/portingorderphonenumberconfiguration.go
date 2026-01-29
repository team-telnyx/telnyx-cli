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

var portingOrdersPhoneNumberConfigurationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a list of phone number configurations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "phone-number-configuration",
			BodyPath: "phone_number_configurations",
		},
	},
	Action:          handlePortingOrdersPhoneNumberConfigurationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"phone-number-configuration": {
		&requestflag.InnerFlag[string]{
			Name:       "phone-number-configuration.porting-phone-number-id",
			Usage:      "Identifies the porting phone number to be configured.",
			InnerField: "porting_phone_number_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phone-number-configuration.user-bundle-id",
			Usage:      "Identifies the user bundle to be associated with the porting phone number.",
			InnerField: "user_bundle_id",
		},
	},
})

var portingOrdersPhoneNumberConfigurationsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of phone number configurations paginated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[porting_order.status][in][], filter[porting_phone_number][in][], filter[user_bundle_id][in][]",
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
	Action:          handlePortingOrdersPhoneNumberConfigurationsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.porting-order",
			InnerField: "porting_order",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.porting-phone-number",
			Usage:      "Filter results by a list of porting phone number IDs",
			InnerField: "porting_phone_number",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.user-bundle-id",
			Usage:      "Filter results by a list of user bundle IDs",
			InnerField: "user_bundle_id",
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
			Usage:      "Specifies the sort order for results. If not given, results are sorted by created_at in descending order.",
			InnerField: "value",
		},
	},
})

func handlePortingOrdersPhoneNumberConfigurationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderPhoneNumberConfigurationNewParams{}

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
	_, err = client.PortingOrders.PhoneNumberConfigurations.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "porting-orders:phone-number-configurations create", obj, format, transform)
}

func handlePortingOrdersPhoneNumberConfigurationsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderPhoneNumberConfigurationListParams{}

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
		_, err = client.PortingOrders.PhoneNumberConfigurations.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "porting-orders:phone-number-configurations list", obj, format, transform)
	} else {
		iter := client.PortingOrders.PhoneNumberConfigurations.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "porting-orders:phone-number-configurations list", iter, format, transform)
	}
}

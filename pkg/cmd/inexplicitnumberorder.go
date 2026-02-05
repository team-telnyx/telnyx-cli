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

var inexplicitNumberOrdersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an inexplicit number order to programmatically purchase phone numbers\nwithout specifying exact numbers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "ordering-group",
			Usage:    "Group(s) of numbers to order. You can have multiple ordering_groups objects added to a single request.",
			Required: true,
			BodyPath: "ordering_groups",
		},
		&requestflag.Flag[string]{
			Name:     "billing-group-id",
			Usage:    "Billing group id to apply to phone numbers that are purchased",
			BodyPath: "billing_group_id",
		},
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Usage:    "Connection id to apply to phone numbers that are purchased",
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "Reference label for the customer",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "Messaging profile id to apply to phone numbers that are purchased",
			BodyPath: "messaging_profile_id",
		},
	},
	Action:          handleInexplicitNumberOrdersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"ordering-group": {
		&requestflag.InnerFlag[string]{
			Name:       "ordering-group.count-requested",
			Usage:      "Quantity of phone numbers to order",
			InnerField: "count_requested",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ordering-group.country-iso",
			Usage:      "Country where you would like to purchase phone numbers. Allowable values: US, CA",
			InnerField: "country_iso",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ordering-group.phone-number-type",
			Usage:      "Number type (local, toll-free, etc.)",
			InnerField: "phone_number_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ordering-group.administrative-area",
			Usage:      "Filter for phone numbers in a given state / province",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "ordering-group.exclude-held-numbers",
			Usage:      "Filter to exclude phone numbers that are currently on hold/reserved for your account.",
			InnerField: "exclude_held_numbers",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "ordering-group.features",
			Usage:      `Filter for phone numbers that have the features to satisfy your use case (e.g., ["voice"])`,
			InnerField: "features",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ordering-group.locality",
			Usage:      "Filter for phone numbers in a given city / region / rate center",
			InnerField: "locality",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ordering-group.national-destination-code",
			Usage:      "Filter by area code",
			InnerField: "national_destination_code",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "ordering-group.phone-number",
			Usage:      "Phone number search criteria",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "ordering-group.quickship",
			Usage:      "Filter to exclude phone numbers that need additional time after to purchase to activate. Only applicable for +1 toll_free numbers.",
			InnerField: "quickship",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ordering-group.strategy",
			Usage:      "Ordering strategy. Define what action should be taken if we don't have enough phone numbers to fulfill your request. Allowable values are: always = proceed with ordering phone numbers, regardless of current inventory levels; never = do not place any orders unless there are enough phone numbers to satisfy the request. If not specified, the always strategy will be enforced.",
			InnerField: "strategy",
		},
	},
})

var inexplicitNumberOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an existing inexplicit number order by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleInexplicitNumberOrdersRetrieve,
	HideHelpCommand: true,
}

var inexplicitNumberOrdersList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of inexplicit number orders.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load",
			Default:   1,
			QueryPath: "page_number",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page",
			Default:   20,
			QueryPath: "page_size",
		},
	},
	Action:          handleInexplicitNumberOrdersList,
	HideHelpCommand: true,
}

func handleInexplicitNumberOrdersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.InexplicitNumberOrderNewParams{}

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
	_, err = client.InexplicitNumberOrders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "inexplicit-number-orders create", obj, format, transform)
}

func handleInexplicitNumberOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.InexplicitNumberOrders.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "inexplicit-number-orders retrieve", obj, format, transform)
}

func handleInexplicitNumberOrdersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.InexplicitNumberOrderListParams{}

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
		_, err = client.InexplicitNumberOrders.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "inexplicit-number-orders list", obj, format, transform)
	} else {
		iter := client.InexplicitNumberOrders.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "inexplicit-number-orders list", iter, format, transform)
	}
}

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

var simCardOrdersCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new order for SIM cards.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "address-id",
			Usage:    "Uniquely identifies the address for the order.",
			Required: true,
			BodyPath: "address_id",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			Usage:    "The amount of SIM cards to order.",
			Required: true,
			BodyPath: "quantity",
		},
	},
	Action:          handleSimCardOrdersCreate,
	HideHelpCommand: true,
}

var simCardOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single SIM card order by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardOrdersRetrieve,
	HideHelpCommand: true,
}

var simCardOrdersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get all SIM card orders according to filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter for SIM card orders (deepObject style). Originally: filter[created_at], filter[updated_at], filter[quantity], filter[cost.amount], filter[cost.currency], filter[address.id], filter[address.street_address], filter[address.extended_address], filter[address.locality], filter[address.administrative_area], filter[address.country_code], filter[address.postal_code]",
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
	Action:          handleSimCardOrdersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.address-administrative-area",
			Usage:      "Filter by state or province where the address is located.",
			InnerField: "address.administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.address-country-code",
			Usage:      "Filter by the mobile operator two-character (ISO 3166-1 alpha-2) origin country code.",
			InnerField: "address.country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.address-extended-address",
			Usage:      "Returns entries with matching name of the supplemental field for address information.",
			InnerField: "address.extended_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.address-id",
			Usage:      "Uniquely identifies the address for the order.",
			InnerField: "address.id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.address-locality",
			Usage:      "Filter by the name of the city where the address is located.",
			InnerField: "address.locality",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.address-postal-code",
			Usage:      "Filter by postal code for the address.",
			InnerField: "address.postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.address-street-address",
			Usage:      "Returns entries with matching name of the street where the address is located.",
			InnerField: "address.street_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.cost-amount",
			Usage:      "The total monetary amount of the order.",
			InnerField: "cost.amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.cost-currency",
			Usage:      "Filter by ISO 4217 currency string.",
			InnerField: "cost.currency",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.created-at",
			Usage:      "Filter by ISO 8601 formatted date-time string matching resource creation date-time.",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filter.quantity",
			Usage:      "Filter orders by how many SIM cards were ordered.",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.updated-at",
			Usage:      "Filter by ISO 8601 formatted date-time string matching resource last update date-time.",
			InnerField: "updated_at",
		},
	},
})

func handleSimCardOrdersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardOrderNewParams{}

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
	_, err = client.SimCardOrders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-orders create", obj, format, transform)
}

func handleSimCardOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCardOrders.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-orders retrieve", obj, format, transform)
}

func handleSimCardOrdersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardOrderListParams{}

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
		_, err = client.SimCardOrders.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "sim-card-orders list", obj, format, transform)
	} else {
		iter := client.SimCardOrders.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "sim-card-orders list", iter, format, transform)
	}
}

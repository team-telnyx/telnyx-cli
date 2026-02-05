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

var numberBlockOrdersCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a phone number block order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "range",
			Usage:    "The phone number range included in the block.",
			Required: true,
			BodyPath: "range",
		},
		&requestflag.Flag[string]{
			Name:     "starting-number",
			Usage:    "Starting phone number block",
			Required: true,
			BodyPath: "starting_number",
		},
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Usage:    "Identifies the connection associated with this phone number.",
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "A customer reference string for customer look ups.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "Identifies the messaging profile associated with the phone number.",
			BodyPath: "messaging_profile_id",
		},
	},
	Action:          handleNumberBlockOrdersCreate,
	HideHelpCommand: true,
}

var numberBlockOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an existing phone number block order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "number-block-order-id",
			Required: true,
		},
	},
	Action:          handleNumberBlockOrdersRetrieve,
	HideHelpCommand: true,
}

var numberBlockOrdersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of number block orders.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[status], filter[created_at], filter[phone_numbers.starting_number]",
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
	Action:          handleNumberBlockOrdersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.created-at",
			Usage:      "Filter number block orders by date range.",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-numbers-starting-number",
			Usage:      "Filter number block  orders having these phone numbers.",
			InnerField: "phone_numbers.starting_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter number block orders by status.",
			InnerField: "status",
		},
	},
})

func handleNumberBlockOrdersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberBlockOrderNewParams{}

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
	_, err = client.NumberBlockOrders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-block-orders create", obj, format, transform)
}

func handleNumberBlockOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("number-block-order-id") && len(unusedArgs) > 0 {
		cmd.Set("number-block-order-id", unusedArgs[0])
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
	_, err = client.NumberBlockOrders.Get(ctx, cmd.Value("number-block-order-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-block-orders retrieve", obj, format, transform)
}

func handleNumberBlockOrdersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberBlockOrderListParams{}

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
		_, err = client.NumberBlockOrders.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "number-block-orders list", obj, format, transform)
	} else {
		iter := client.NumberBlockOrders.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "number-block-orders list", iter, format, transform)
	}
}

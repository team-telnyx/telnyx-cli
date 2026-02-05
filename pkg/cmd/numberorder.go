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

var numberOrdersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a phone number order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "billing-group-id",
			Usage:    "Identifies the billing group associated with the phone number.",
			BodyPath: "billing_group_id",
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
		&requestflag.Flag[[]map[string]any]{
			Name:     "phone-number",
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleNumberOrdersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"phone-number": {
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.phone-number",
			Usage:      "e164_phone_number",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.bundle-id",
			Usage:      "ID of bundle to associate the number to",
			InnerField: "bundle_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.requirement-group-id",
			Usage:      "ID of requirement group to use to satisfy number requirements",
			InnerField: "requirement_group_id",
		},
	},
})

var numberOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an existing phone number order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "number-order-id",
			Required: true,
		},
	},
	Action:          handleNumberOrdersRetrieve,
	HideHelpCommand: true,
}

var numberOrdersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates a phone number order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "number-order-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "A customer reference string for customer look ups.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "regulatory-requirement",
			BodyPath: "regulatory_requirements",
		},
	},
	Action:          handleNumberOrdersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"regulatory-requirement": {
		&requestflag.InnerFlag[string]{
			Name:       "regulatory-requirement.field-value",
			Usage:      "The value of the requirement. For address and document requirements, this should be the ID of the resource. For textual, this should be the value of the requirement.",
			InnerField: "field_value",
		},
		&requestflag.InnerFlag[string]{
			Name:       "regulatory-requirement.requirement-id",
			Usage:      "Unique id for a requirement.",
			InnerField: "requirement_id",
		},
	},
})

var numberOrdersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of number orders.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[status], filter[created_at], filter[phone_numbers_count], filter[customer_reference], filter[requirements_met]",
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
	Action:          handleNumberOrdersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.created-at",
			Usage:      "Filter number orders by date range.",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.customer-reference",
			Usage:      "Filter number orders via the customer reference set.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-numbers-count",
			Usage:      "Filter number order with this amount of numbers",
			InnerField: "phone_numbers_count",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.requirements-met",
			Usage:      "Filter number orders by requirements met.",
			InnerField: "requirements_met",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter number orders by status.",
			InnerField: "status",
		},
	},
})

func handleNumberOrdersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberOrderNewParams{}

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
	_, err = client.NumberOrders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-orders create", obj, format, transform)
}

func handleNumberOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("number-order-id") && len(unusedArgs) > 0 {
		cmd.Set("number-order-id", unusedArgs[0])
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
	_, err = client.NumberOrders.Get(ctx, cmd.Value("number-order-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-orders retrieve", obj, format, transform)
}

func handleNumberOrdersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("number-order-id") && len(unusedArgs) > 0 {
		cmd.Set("number-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberOrderUpdateParams{}

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
	_, err = client.NumberOrders.Update(
		ctx,
		cmd.Value("number-order-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-orders update", obj, format, transform)
}

func handleNumberOrdersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberOrderListParams{}

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
		_, err = client.NumberOrders.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "number-orders list", obj, format, transform)
	} else {
		iter := client.NumberOrders.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "number-orders list", iter, format, transform)
	}
}

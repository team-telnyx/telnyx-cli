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

var subNumberOrdersRetrieve = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve",
	Usage:   "Get an existing sub number order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sub-number-order-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[include_phone_numbers]",
			QueryPath: "filter",
		},
	},
	Action:          handleSubNumberOrdersRetrieve,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[bool]{
			Name:       "filter.include-phone-numbers",
			Usage:      "Include the first 50 phone number objects in the results",
			InnerField: "include_phone_numbers",
		},
	},
})

var subNumberOrdersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates a sub number order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sub-number-order-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "regulatory-requirement",
			BodyPath: "regulatory_requirements",
		},
	},
	Action:          handleSubNumberOrdersUpdate,
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

var subNumberOrdersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of sub number orders.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[status], filter[order_request_id], filter[country_code], filter[phone_number_type], filter[phone_numbers_count]",
			QueryPath: "filter",
		},
	},
	Action:          handleSubNumberOrdersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "ISO alpha-2 country code.",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.order-request-id",
			Usage:      "ID of the number order the sub number order belongs to",
			InnerField: "order_request_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number-type",
			Usage:      "Phone Number Type",
			InnerField: "phone_number_type",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filter.phone-numbers-count",
			Usage:      "Amount of numbers in the sub number order",
			InnerField: "phone_numbers_count",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter sub number orders by status.",
			InnerField: "status",
		},
	},
})

var subNumberOrdersCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Allows you to cancel a sub number order in 'pending' status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sub-number-order-id",
			Required: true,
		},
	},
	Action:          handleSubNumberOrdersCancel,
	HideHelpCommand: true,
}

var subNumberOrdersUpdateRequirementGroup = cli.Command{
	Name:    "update-requirement-group",
	Usage:   "Update requirement group for a sub number order",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "requirement-group-id",
			Usage:    "The ID of the requirement group to associate",
			Required: true,
			BodyPath: "requirement_group_id",
		},
	},
	Action:          handleSubNumberOrdersUpdateRequirementGroup,
	HideHelpCommand: true,
}

func handleSubNumberOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sub-number-order-id") && len(unusedArgs) > 0 {
		cmd.Set("sub-number-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SubNumberOrderGetParams{}

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
	_, err = client.SubNumberOrders.Get(
		ctx,
		cmd.Value("sub-number-order-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sub-number-orders retrieve", obj, format, transform)
}

func handleSubNumberOrdersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sub-number-order-id") && len(unusedArgs) > 0 {
		cmd.Set("sub-number-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SubNumberOrderUpdateParams{}

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
	_, err = client.SubNumberOrders.Update(
		ctx,
		cmd.Value("sub-number-order-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sub-number-orders update", obj, format, transform)
}

func handleSubNumberOrdersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SubNumberOrderListParams{}

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
	_, err = client.SubNumberOrders.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sub-number-orders list", obj, format, transform)
}

func handleSubNumberOrdersCancel(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sub-number-order-id") && len(unusedArgs) > 0 {
		cmd.Set("sub-number-order-id", unusedArgs[0])
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
	_, err = client.SubNumberOrders.Cancel(ctx, cmd.Value("sub-number-order-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sub-number-orders cancel", obj, format, transform)
}

func handleSubNumberOrdersUpdateRequirementGroup(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SubNumberOrderUpdateRequirementGroupParams{}

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
	_, err = client.SubNumberOrders.UpdateRequirementGroup(
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
	return ShowJSON(os.Stdout, "sub-number-orders update-requirement-group", obj, format, transform)
}

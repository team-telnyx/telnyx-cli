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

var advancedOrdersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create Advanced Order",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "area-code",
			Default:  "",
			BodyPath: "area_code",
		},
		&requestflag.Flag[string]{
			Name:     "comments",
			Default:  "",
			BodyPath: "comments",
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Default:  "US",
			BodyPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Default:  "",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[[]string]{
			Name:     "feature",
			BodyPath: "features",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number-type",
			Default:  "local",
			BodyPath: "phone_number_type",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			Default:  1,
			BodyPath: "quantity",
		},
		&requestflag.Flag[string]{
			Name:     "requirement-group-id",
			Usage:    "The ID of the requirement group to associate with this advanced order",
			BodyPath: "requirement_group_id",
		},
	},
	Action:          handleAdvancedOrdersCreate,
	HideHelpCommand: true,
}

var advancedOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get Advanced Order",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "order-id",
			Required: true,
		},
	},
	Action:          handleAdvancedOrdersRetrieve,
	HideHelpCommand: true,
}

var advancedOrdersList = cli.Command{
	Name:            "list",
	Usage:           "List Advanced Orders",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAdvancedOrdersList,
	HideHelpCommand: true,
}

var advancedOrdersUpdateRequirementGroup = cli.Command{
	Name:    "update-requirement-group",
	Usage:   "Update Advanced Order",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "advanced-order-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "area-code",
			Default:  "",
			BodyPath: "area_code",
		},
		&requestflag.Flag[string]{
			Name:     "comments",
			Default:  "",
			BodyPath: "comments",
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Default:  "US",
			BodyPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Default:  "",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[[]string]{
			Name:     "feature",
			BodyPath: "features",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number-type",
			Default:  "local",
			BodyPath: "phone_number_type",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			Default:  1,
			BodyPath: "quantity",
		},
		&requestflag.Flag[string]{
			Name:     "requirement-group-id",
			Usage:    "The ID of the requirement group to associate with this advanced order",
			BodyPath: "requirement_group_id",
		},
	},
	Action:          handleAdvancedOrdersUpdateRequirementGroup,
	HideHelpCommand: true,
}

func handleAdvancedOrdersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AdvancedOrderNewParams{}

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
	_, err = client.AdvancedOrders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "advanced-orders create", obj, format, transform)
}

func handleAdvancedOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
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
	_, err = client.AdvancedOrders.Get(ctx, cmd.Value("order-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "advanced-orders retrieve", obj, format, transform)
}

func handleAdvancedOrdersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.AdvancedOrders.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "advanced-orders list", obj, format, transform)
}

func handleAdvancedOrdersUpdateRequirementGroup(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("advanced-order-id") && len(unusedArgs) > 0 {
		cmd.Set("advanced-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AdvancedOrderUpdateRequirementGroupParams{}

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
	_, err = client.AdvancedOrders.UpdateRequirementGroup(
		ctx,
		cmd.Value("advanced-order-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "advanced-orders update-requirement-group", obj, format, transform)
}

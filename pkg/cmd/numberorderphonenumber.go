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

var numberOrderPhoneNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an existing phone number in number order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "number-order-phone-number-id",
			Required: true,
		},
	},
	Action:          handleNumberOrderPhoneNumbersRetrieve,
	HideHelpCommand: true,
}

var numberOrderPhoneNumbersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get a list of phone numbers associated to orders.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[country_code]",
			QueryPath: "filter",
		},
	},
	Action:          handleNumberOrderPhoneNumbersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "Country code of the order phone number.",
			InnerField: "country_code",
		},
	},
})

var numberOrderPhoneNumbersUpdateRequirementGroup = cli.Command{
	Name:    "update-requirement-group",
	Usage:   "Update requirement group for a phone number order",
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
	Action:          handleNumberOrderPhoneNumbersUpdateRequirementGroup,
	HideHelpCommand: true,
}

var numberOrderPhoneNumbersUpdateRequirements = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-requirements",
	Usage:   "Updates requirements for a single phone number within a number order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "number-order-phone-number-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "regulatory-requirement",
			BodyPath: "regulatory_requirements",
		},
	},
	Action:          handleNumberOrderPhoneNumbersUpdateRequirements,
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

func handleNumberOrderPhoneNumbersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("number-order-phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("number-order-phone-number-id", unusedArgs[0])
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
	_, err = client.NumberOrderPhoneNumbers.Get(ctx, cmd.Value("number-order-phone-number-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-order-phone-numbers retrieve", obj, format, transform)
}

func handleNumberOrderPhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberOrderPhoneNumberListParams{}

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
	_, err = client.NumberOrderPhoneNumbers.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-order-phone-numbers list", obj, format, transform)
}

func handleNumberOrderPhoneNumbersUpdateRequirementGroup(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberOrderPhoneNumberUpdateRequirementGroupParams{}

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
	_, err = client.NumberOrderPhoneNumbers.UpdateRequirementGroup(
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
	return ShowJSON(os.Stdout, "number-order-phone-numbers update-requirement-group", obj, format, transform)
}

func handleNumberOrderPhoneNumbersUpdateRequirements(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("number-order-phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("number-order-phone-number-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberOrderPhoneNumberUpdateRequirementsParams{}

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
	_, err = client.NumberOrderPhoneNumbers.UpdateRequirements(
		ctx,
		cmd.Value("number-order-phone-number-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-order-phone-numbers update-requirements", obj, format, transform)
}

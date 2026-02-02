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

var messagingHostedNumberOrdersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a messaging hosted number order",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "Automatically associate the number with this messaging profile ID when the order is complete.",
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "Phone numbers to be used for hosted messaging.",
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleMessagingHostedNumberOrdersCreate,
	HideHelpCommand: true,
}

var messagingHostedNumberOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a messaging hosted number order",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMessagingHostedNumberOrdersRetrieve,
	HideHelpCommand: true,
}

var messagingHostedNumberOrdersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List messaging hosted number orders",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[number], page[size]",
			QueryPath: "page",
		},
	},
	Action:          handleMessagingHostedNumberOrdersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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
})

var messagingHostedNumberOrdersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a messaging hosted number order and all associated phone numbers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMessagingHostedNumberOrdersDelete,
	HideHelpCommand: true,
}

var messagingHostedNumberOrdersCheckEligibility = cli.Command{
	Name:    "check-eligibility",
	Usage:   "Check hosted messaging eligibility",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "List of phone numbers to check eligibility",
			Required: true,
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleMessagingHostedNumberOrdersCheckEligibility,
	HideHelpCommand: true,
}

var messagingHostedNumberOrdersCreateVerificationCodes = cli.Command{
	Name:    "create-verification-codes",
	Usage:   "Create verification codes to validate numbers of the hosted order. The\nverification codes will be sent to the numbers of the hosted order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Required: true,
			BodyPath: "phone_numbers",
		},
		&requestflag.Flag[string]{
			Name:     "verification-method",
			Required: true,
			BodyPath: "verification_method",
		},
	},
	Action:          handleMessagingHostedNumberOrdersCreateVerificationCodes,
	HideHelpCommand: true,
}

var messagingHostedNumberOrdersValidateCodes = requestflag.WithInnerFlags(cli.Command{
	Name:    "validate-codes",
	Usage:   "Validate the verification codes sent to the numbers of the hosted order. The\nverification codes must be created in the verification codes endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "verification-code",
			Required: true,
			BodyPath: "verification_codes",
		},
	},
	Action:          handleMessagingHostedNumberOrdersValidateCodes,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"verification-code": {
		&requestflag.InnerFlag[string]{
			Name:       "verification-code.code",
			InnerField: "code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "verification-code.phone-number",
			InnerField: "phone_number",
		},
	},
})

func handleMessagingHostedNumberOrdersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingHostedNumberOrderNewParams{}

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
	_, err = client.MessagingHostedNumberOrders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-hosted-number-orders create", obj, format, transform)
}

func handleMessagingHostedNumberOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MessagingHostedNumberOrders.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-hosted-number-orders retrieve", obj, format, transform)
}

func handleMessagingHostedNumberOrdersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingHostedNumberOrderListParams{}

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
		_, err = client.MessagingHostedNumberOrders.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messaging-hosted-number-orders list", obj, format, transform)
	} else {
		iter := client.MessagingHostedNumberOrders.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "messaging-hosted-number-orders list", iter, format, transform)
	}
}

func handleMessagingHostedNumberOrdersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MessagingHostedNumberOrders.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-hosted-number-orders delete", obj, format, transform)
}

func handleMessagingHostedNumberOrdersCheckEligibility(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingHostedNumberOrderCheckEligibilityParams{}

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
	_, err = client.MessagingHostedNumberOrders.CheckEligibility(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-hosted-number-orders check-eligibility", obj, format, transform)
}

func handleMessagingHostedNumberOrdersCreateVerificationCodes(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingHostedNumberOrderNewVerificationCodesParams{}

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
	_, err = client.MessagingHostedNumberOrders.NewVerificationCodes(
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
	return ShowJSON(os.Stdout, "messaging-hosted-number-orders create-verification-codes", obj, format, transform)
}

func handleMessagingHostedNumberOrdersValidateCodes(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingHostedNumberOrderValidateCodesParams{}

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
	_, err = client.MessagingHostedNumberOrders.ValidateCodes(
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
	return ShowJSON(os.Stdout, "messaging-hosted-number-orders validate-codes", obj, format, transform)
}

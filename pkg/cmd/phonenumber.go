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

var phoneNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePhoneNumbersRetrieve,
	HideHelpCommand: true,
}

var phoneNumbersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "billing-group-id",
			Usage:    "Identifies the billing group associated with the phone number.",
			BodyPath: "billing_group_id",
		},
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Usage:    "Identifies the connection associated with the phone number.",
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "A customer reference string for customer look ups.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "external-pin",
			Usage:    "If someone attempts to port your phone number away from Telnyx and your phone number has an external PIN set, we will attempt to verify that you provided the correct external PIN to the winning carrier. Note that not all carriers cooperate with this security mechanism.",
			BodyPath: "external_pin",
		},
		&requestflag.Flag[bool]{
			Name:     "hd-voice-enabled",
			Usage:    "Indicates whether HD voice is enabled for this number.",
			BodyPath: "hd_voice_enabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "A list of user-assigned tags to help organize phone numbers.",
			BodyPath: "tags",
		},
	},
	Action:          handlePhoneNumbersUpdate,
	HideHelpCommand: true,
}

var phoneNumbersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List phone numbers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[tag], filter[phone_number], filter[status], filter[country_iso_alpha2], filter[connection_id], filter[voice.connection_name], filter[voice.usage_payment_method], filter[billing_group_id], filter[emergency_address_id], filter[customer_reference], filter[number_type], filter[source]",
			QueryPath: "filter",
		},
		&requestflag.Flag[string]{
			Name:      "handle-messaging-profile-error",
			Usage:     "Although it is an infrequent occurrence, due to the highly distributed nature of the Telnyx platform, it is possible that there will be an issue when loading in Messaging Profile information. As such, when this parameter is set to `true` and an error in fetching this information occurs, messaging profile related fields will be omitted in the response and an error message will be included instead of returning a 503 error.",
			Default:   "false",
			QueryPath: "handle_messaging_profile_error",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. If not given, results are sorted by created_at in descending order.",
			QueryPath: "sort",
		},
	},
	Action:          handlePhoneNumbersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.billing-group-id",
			Usage:      "Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string 'null'.",
			InnerField: "billing_group_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.connection-id",
			Usage:      "Filter by connection_id.",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.country-iso-alpha2",
			Usage:      "Filter by phone number country ISO alpha-2 code. Can be a single value or an array of values.",
			InnerField: "country_iso_alpha2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.customer-reference",
			Usage:      "Filter numbers via the customer_reference set.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.emergency-address-id",
			Usage:      "Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string 'null'.",
			InnerField: "emergency_address_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.number-type",
			Usage:      "Filter phone numbers by phone number type.",
			InnerField: "number_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number",
			Usage:      "Filter by phone number. Requires at least three digits.\n             Non-numerical characters will result in no values being returned.",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.source",
			Usage:      "Filter phone numbers by their source. Use 'ported' for numbers ported from other carriers, or 'purchased' for numbers bought directly from Telnyx.",
			InnerField: "source",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter by phone number status.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.tag",
			Usage:      "Filter by phone number tags.",
			InnerField: "tag",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.voice-connection-name",
			Usage:      "Filter by voice connection name pattern matching.",
			InnerField: "voice.connection_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.voice-usage-payment-method",
			Usage:      "Filter by usage_payment_method.",
			InnerField: "voice.usage_payment_method",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.without-tags",
			Usage:      "When set to 'true', filters for phone numbers that do not have any tags applied. All other values are ignored.",
			InnerField: "without_tags",
		},
	},
})

var phoneNumbersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePhoneNumbersDelete,
	HideHelpCommand: true,
}

var phoneNumbersSlimList = requestflag.WithInnerFlags(cli.Command{
	Name:    "slim-list",
	Usage:   "List phone numbers, This endpoint is a lighter version of the /phone_numbers\nendpoint having higher performance and rate limit.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[tag], filter[phone_number], filter[status], filter[country_iso_alpha2], filter[connection_id], filter[voice.connection_name], filter[voice.usage_payment_method], filter[billing_group_id], filter[emergency_address_id], filter[customer_reference], filter[number_type], filter[source]",
			QueryPath: "filter",
		},
		&requestflag.Flag[bool]{
			Name:      "include-connection",
			Usage:     "Include the connection associated with the phone number.",
			Default:   false,
			QueryPath: "include_connection",
		},
		&requestflag.Flag[bool]{
			Name:      "include-tags",
			Usage:     "Include the tags associated with the phone number.",
			Default:   false,
			QueryPath: "include_tags",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. If not given, results are sorted by created_at in descending order.",
			QueryPath: "sort",
		},
	},
	Action:          handlePhoneNumbersSlimList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.billing-group-id",
			Usage:      "Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string 'null'.",
			InnerField: "billing_group_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.connection-id",
			Usage:      "Filter by connection_id.",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.country-iso-alpha2",
			Usage:      "Filter by phone number country ISO alpha-2 code. Can be a single value or an array of values.",
			InnerField: "country_iso_alpha2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.customer-reference",
			Usage:      "Filter numbers via the customer_reference set.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.emergency-address-id",
			Usage:      "Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string 'null'.",
			InnerField: "emergency_address_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.number-type",
			Usage:      "Filter phone numbers by phone number type.",
			InnerField: "number_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number",
			Usage:      "Filter by phone number. Requires at least three digits.\n             Non-numerical characters will result in no values being returned.",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.source",
			Usage:      "Filter phone numbers by their source. Use 'ported' for numbers ported from other carriers, or 'purchased' for numbers bought directly from Telnyx.",
			InnerField: "source",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter by phone number status.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.tag",
			Usage:      "Filter by phone number tags. (This requires the include_tags param)",
			InnerField: "tag",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.voice-connection-name",
			Usage:      "Filter by voice connection name pattern matching (requires include_connection param).",
			InnerField: "voice.connection_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.voice-usage-payment-method",
			Usage:      "Filter by usage_payment_method.",
			InnerField: "voice.usage_payment_method",
		},
	},
})

func handlePhoneNumbersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PhoneNumbers.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers retrieve", obj, format, transform)
}

func handlePhoneNumbersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberUpdateParams{}

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
	_, err = client.PhoneNumbers.Update(
		ctx,
		cmd.Value("phone-number-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers update", obj, format, transform)
}

func handlePhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberListParams{}

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
		_, err = client.PhoneNumbers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "phone-numbers list", obj, format, transform)
	} else {
		iter := client.PhoneNumbers.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "phone-numbers list", iter, format, transform)
	}
}

func handlePhoneNumbersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PhoneNumbers.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers delete", obj, format, transform)
}

func handlePhoneNumbersSlimList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberSlimListParams{}

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
		_, err = client.PhoneNumbers.SlimList(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "phone-numbers slim-list", obj, format, transform)
	} else {
		iter := client.PhoneNumbers.SlimListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "phone-numbers slim-list", iter, format, transform)
	}
}

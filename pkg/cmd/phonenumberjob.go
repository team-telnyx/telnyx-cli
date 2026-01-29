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

var phoneNumbersJobsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a phone numbers job",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePhoneNumbersJobsRetrieve,
	HideHelpCommand: true,
}

var phoneNumbersJobsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Lists the phone numbers jobs",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[type]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. If not given, results are sorted by created_at in descending order.",
			QueryPath: "sort",
		},
	},
	Action:          handlePhoneNumbersJobsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			Usage:      "Identifies the type of the background job.",
			InnerField: "type",
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
})

var phoneNumbersJobsDeleteBatch = cli.Command{
	Name:    "delete-batch",
	Usage:   "Creates a new background job to delete a batch of numbers. At most one thousand\nnumbers can be updated per API call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Required: true,
			BodyPath: "phone_numbers",
		},
	},
	Action:          handlePhoneNumbersJobsDeleteBatch,
	HideHelpCommand: true,
}

var phoneNumbersJobsUpdateBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-batch",
	Usage:   "Creates a new background job to update a batch of numbers. At most one thousand\nnumbers can be updated per API call. At least one of the updateable fields must\nbe submitted. IMPORTANT: You must either specify filters (using the filter\nparameters) or specific phone numbers (using the phone_numbers parameter in the\nrequest body). If you specify filters, ALL phone numbers that match the given\nfilters (up to 1000 at a time) will be updated. If you want to update only\nspecific numbers, you must use the phone_numbers parameter in the request body.\nWhen using the phone_numbers parameter, ensure you follow the correct format as\nshown in the example (either phone number IDs or phone numbers in E164 format).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "Array of phone number ids and/or phone numbers in E164 format to update. This parameter is required if no filter parameters are provided. If you want to update specific numbers rather than all numbers matching a filter, you must use this parameter. Each item must be either a valid phone number ID or a phone number in E164 format (e.g., '+13127367254').",
			Required: true,
			BodyPath: "phone_numbers",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[has_bundle], filter[tag], filter[connection_id], filter[phone_number], filter[status], filter[voice.connection_name], filter[voice.usage_payment_method], filter[billing_group_id], filter[emergency_address_id], filter[customer_reference]",
			QueryPath: "filter",
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
		&requestflag.Flag[bool]{
			Name:     "deletion-lock-enabled",
			Usage:    "Indicates whether to enable or disable the deletion lock on each phone number. When enabled, this prevents the phone number from being deleted via the API or Telnyx portal.",
			BodyPath: "deletion_lock_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "external-pin",
			Usage:    "If someone attempts to port your phone number away from Telnyx and your phone number has an external PIN set, we will attempt to verify that you provided the correct external PIN to the winning carrier. Note that not all carriers cooperate with this security mechanism.",
			BodyPath: "external_pin",
		},
		&requestflag.Flag[bool]{
			Name:     "hd-voice-enabled",
			Usage:    "Indicates whether to enable or disable HD Voice on each phone number. HD Voice is a paid feature and may not be available for all phone numbers, more details about it can be found in the Telnyx support documentation.",
			BodyPath: "hd_voice_enabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "A list of user-assigned tags to help organize phone numbers.",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "voice",
			BodyPath: "voice",
		},
	},
	Action:          handlePhoneNumbersJobsUpdateBatch,
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
		&requestflag.InnerFlag[string]{
			Name:       "filter.has-bundle",
			Usage:      "Filter by phone number that have bundles.",
			InnerField: "has_bundle",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number",
			Usage:      "Filter by phone number. Requires at least three digits.\n             Non-numerical characters will result in no values being returned.",
			InnerField: "phone_number",
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
	},
	"voice": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "voice.call-forwarding",
			Usage:      "The call forwarding settings for a phone number.",
			InnerField: "call_forwarding",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "voice.call-recording",
			Usage:      "The call recording settings for a phone number.",
			InnerField: "call_recording",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "voice.caller-id-name-enabled",
			Usage:      "Controls whether the caller ID name is enabled for this phone number.",
			InnerField: "caller_id_name_enabled",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "voice.cnam-listing",
			Usage:      "The CNAM listing settings for a phone number.",
			InnerField: "cnam_listing",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.inbound-call-screening",
			Usage:      "The inbound_call_screening setting is a phone number configuration option variable that allows users to configure their settings to block or flag fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This feature has an additional per-number monthly cost associated with it.",
			InnerField: "inbound_call_screening",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "voice.media-features",
			Usage:      "The media features settings for a phone number.",
			InnerField: "media_features",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "voice.tech-prefix-enabled",
			Usage:      "Controls whether a tech prefix is enabled for this phone number.",
			InnerField: "tech_prefix_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.translated-number",
			Usage:      "This field allows you to rewrite the destination number of an inbound call before the call is routed to you. The value of this field may be any alphanumeric value, and the value will replace the number originally dialed.",
			InnerField: "translated_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice.usage-payment-method",
			Usage:      "Controls whether a number is billed per minute or uses your concurrent channels.",
			InnerField: "usage_payment_method",
		},
	},
})

var phoneNumbersJobsUpdateEmergencySettingsBatch = cli.Command{
	Name:    "update-emergency-settings-batch",
	Usage:   "Creates a background job to update the emergency settings of a collection of\nphone numbers. At most one thousand numbers can be updated per API call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "emergency-enabled",
			Usage:    "Indicates whether to enable or disable emergency services on the numbers.",
			Required: true,
			BodyPath: "emergency_enabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Required: true,
			BodyPath: "phone_numbers",
		},
		&requestflag.Flag[any]{
			Name:     "emergency-address-id",
			Usage:    "Identifies the address to be used with emergency services. Required if emergency_enabled is true, must be null or omitted if emergency_enabled is false.",
			BodyPath: "emergency_address_id",
		},
	},
	Action:          handlePhoneNumbersJobsUpdateEmergencySettingsBatch,
	HideHelpCommand: true,
}

func handlePhoneNumbersJobsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PhoneNumbers.Jobs.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers:jobs retrieve", obj, format, transform)
}

func handlePhoneNumbersJobsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberJobListParams{}

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
		_, err = client.PhoneNumbers.Jobs.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "phone-numbers:jobs list", obj, format, transform)
	} else {
		iter := client.PhoneNumbers.Jobs.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "phone-numbers:jobs list", iter, format, transform)
	}
}

func handlePhoneNumbersJobsDeleteBatch(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberJobDeleteBatchParams{}

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
	_, err = client.PhoneNumbers.Jobs.DeleteBatch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers:jobs delete-batch", obj, format, transform)
}

func handlePhoneNumbersJobsUpdateBatch(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberJobUpdateBatchParams{}

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
	_, err = client.PhoneNumbers.Jobs.UpdateBatch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers:jobs update-batch", obj, format, transform)
}

func handlePhoneNumbersJobsUpdateEmergencySettingsBatch(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberJobUpdateEmergencySettingsBatchParams{}

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
	_, err = client.PhoneNumbers.Jobs.UpdateEmergencySettingsBatch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers:jobs update-emergency-settings-batch", obj, format, transform)
}

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

var phoneNumbersVoiceRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a phone number with voice settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePhoneNumbersVoiceRetrieve,
	HideHelpCommand: true,
}

var phoneNumbersVoiceUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a phone number with voice settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "call-forwarding",
			Usage:    "The call forwarding settings for a phone number.",
			BodyPath: "call_forwarding",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "call-recording",
			Usage:    "The call recording settings for a phone number.",
			BodyPath: "call_recording",
		},
		&requestflag.Flag[bool]{
			Name:     "caller-id-name-enabled",
			Usage:    "Controls whether the caller ID name is enabled for this phone number.",
			Default:  false,
			BodyPath: "caller_id_name_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "cnam-listing",
			Usage:    "The CNAM listing settings for a phone number.",
			BodyPath: "cnam_listing",
		},
		&requestflag.Flag[string]{
			Name:     "inbound-call-screening",
			Usage:    "The inbound_call_screening setting is a phone number configuration option variable that allows users to configure their settings to block or flag fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This feature has an additional per-number monthly cost associated with it.",
			Default:  "disabled",
			BodyPath: "inbound_call_screening",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "media-features",
			Usage:    "The media features settings for a phone number.",
			BodyPath: "media_features",
		},
		&requestflag.Flag[bool]{
			Name:     "tech-prefix-enabled",
			Usage:    "Controls whether a tech prefix is enabled for this phone number.",
			Default:  false,
			BodyPath: "tech_prefix_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "translated-number",
			Usage:    "This field allows you to rewrite the destination number of an inbound call before the call is routed to you. The value of this field may be any alphanumeric value, and the value will replace the number originally dialed.",
			BodyPath: "translated_number",
		},
		&requestflag.Flag[string]{
			Name:     "usage-payment-method",
			Usage:    "Controls whether a number is billed per minute or uses your concurrent channels.",
			Default:  "pay-per-minute",
			BodyPath: "usage_payment_method",
		},
	},
	Action:          handlePhoneNumbersVoiceUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"call-forwarding": {
		&requestflag.InnerFlag[bool]{
			Name:       "call-forwarding.call-forwarding-enabled",
			Usage:      "Indicates if call forwarding will be enabled for this number if forwards_to and forwarding_type are filled in. Defaults to true for backwards compatibility with APIV1 use of numbers endpoints.",
			InnerField: "call_forwarding_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-forwarding.forwarding-type",
			Usage:      "Call forwarding type. 'forwards_to' must be set for this to have an effect.",
			InnerField: "forwarding_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-forwarding.forwards-to",
			Usage:      "The phone number to which inbound calls to this number are forwarded. Inbound calls will not be forwarded if this field is left blank. If set, must be a +E.164-formatted phone number.",
			InnerField: "forwards_to",
		},
	},
	"call-recording": {
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.inbound-call-recording-channels",
			Usage:      "When using 'dual' channels, final audio file will be stereo recorded with the first leg on channel A, and the rest on channel B.",
			InnerField: "inbound_call_recording_channels",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "call-recording.inbound-call-recording-enabled",
			Usage:      "When enabled, any inbound call to this number will be recorded.",
			InnerField: "inbound_call_recording_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.inbound-call-recording-format",
			Usage:      "The audio file format for calls being recorded.",
			InnerField: "inbound_call_recording_format",
		},
	},
	"cnam-listing": {
		&requestflag.InnerFlag[string]{
			Name:       "cnam-listing.cnam-listing-details",
			Usage:      "The CNAM listing details for this number. Must be alphanumeric characters or spaces with a maximum length of 15. Requires cnam_listing_enabled to also be set to true.",
			InnerField: "cnam_listing_details",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "cnam-listing.cnam-listing-enabled",
			Usage:      "Enables CNAM listings for this number. Requires cnam_listing_details to also be set.",
			InnerField: "cnam_listing_enabled",
		},
	},
	"media-features": {
		&requestflag.InnerFlag[bool]{
			Name:       "media-features.accept-any-rtp-packets-enabled",
			Usage:      "When enabled, Telnyx will accept RTP packets from any customer-side IP address and port, not just those to which Telnyx is sending RTP.",
			InnerField: "accept_any_rtp_packets_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "media-features.rtp-auto-adjust-enabled",
			Usage:      "When RTP Auto-Adjust is enabled, the destination RTP address port will be automatically changed to match the source of the incoming RTP packets.",
			InnerField: "rtp_auto_adjust_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "media-features.t38-fax-gateway-enabled",
			Usage:      "Controls whether Telnyx will accept a T.38 re-INVITE for this phone number. Note that Telnyx will not send a T.38 re-INVITE; this option only controls whether one will be accepted.",
			InnerField: "t38_fax_gateway_enabled",
		},
	},
})

var phoneNumbersVoiceList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List phone numbers with voice settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[phone_number], filter[connection_name], filter[customer_reference], filter[voice.usage_payment_method]",
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
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. If not given, results are sorted by created_at in descending order.",
			QueryPath: "sort",
		},
	},
	Action:          handlePhoneNumbersVoiceList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.connection-name",
			Usage:      "Filter by connection name pattern matching.",
			InnerField: "connection_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.customer-reference",
			Usage:      "Filter numbers via the customer_reference set.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number",
			Usage:      "Filter by phone number. Requires at least three digits.\n             Non-numerical characters will result in no values being returned.",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.voice-usage-payment-method",
			Usage:      "Filter by usage_payment_method.",
			InnerField: "voice.usage_payment_method",
		},
	},
})

func handlePhoneNumbersVoiceRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PhoneNumbers.Voice.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers:voice retrieve", obj, format, transform)
}

func handlePhoneNumbersVoiceUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberVoiceUpdateParams{}

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
	_, err = client.PhoneNumbers.Voice.Update(
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
	return ShowJSON(os.Stdout, "phone-numbers:voice update", obj, format, transform)
}

func handlePhoneNumbersVoiceList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberVoiceListParams{}

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
		_, err = client.PhoneNumbers.Voice.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "phone-numbers:voice list", obj, format, transform)
	} else {
		iter := client.PhoneNumbers.Voice.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "phone-numbers:voice list", iter, format, transform)
	}
}

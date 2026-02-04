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

var mobilePhoneNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Mobile Phone Number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMobilePhoneNumbersRetrieve,
	HideHelpCommand: true,
}

var mobilePhoneNumbersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a Mobile Phone Number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "call-forwarding",
			BodyPath: "call_forwarding",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "call-recording",
			BodyPath: "call_recording",
		},
		&requestflag.Flag[bool]{
			Name:     "caller-id-name-enabled",
			BodyPath: "caller_id_name_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "cnam-listing",
			BodyPath: "cnam_listing",
		},
		&requestflag.Flag[any]{
			Name:     "connection-id",
			BodyPath: "connection_id",
		},
		&requestflag.Flag[any]{
			Name:     "customer-reference",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[string]{
			Name:     "inbound-call-screening",
			BodyPath: "inbound_call_screening",
		},
		&requestflag.Flag[bool]{
			Name:     "noise-suppression",
			BodyPath: "noise_suppression",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handleMobilePhoneNumbersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"call-forwarding": {
		&requestflag.InnerFlag[bool]{
			Name:       "call-forwarding.call-forwarding-enabled",
			InnerField: "call_forwarding_enabled",
		},
		&requestflag.InnerFlag[any]{
			Name:       "call-forwarding.forwarding-type",
			InnerField: "forwarding_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "call-forwarding.forwards-to",
			InnerField: "forwards_to",
		},
	},
	"call-recording": {
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.inbound-call-recording-channels",
			InnerField: "inbound_call_recording_channels",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "call-recording.inbound-call-recording-enabled",
			InnerField: "inbound_call_recording_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.inbound-call-recording-format",
			InnerField: "inbound_call_recording_format",
		},
	},
	"cnam-listing": {
		&requestflag.InnerFlag[any]{
			Name:       "cnam-listing.cnam-listing-details",
			InnerField: "cnam_listing_details",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "cnam-listing.cnam-listing-enabled",
			InnerField: "cnam_listing_enabled",
		},
	},
	"inbound": {
		&requestflag.InnerFlag[any]{
			Name:       "inbound.interception-app-id",
			Usage:      "The ID of the CallControl or TeXML Application that will intercept inbound calls.",
			InnerField: "interception_app_id",
		},
	},
	"outbound": {
		&requestflag.InnerFlag[any]{
			Name:       "outbound.interception-app-id",
			Usage:      "The ID of the CallControl or TeXML Application that will intercept outbound calls.",
			InnerField: "interception_app_id",
		},
	},
})

var mobilePhoneNumbersList = cli.Command{
	Name:    "list",
	Usage:   "List Mobile Phone Numbers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page",
			QueryPath: "page[size]",
		},
	},
	Action:          handleMobilePhoneNumbersList,
	HideHelpCommand: true,
}

func handleMobilePhoneNumbersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MobilePhoneNumbers.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "mobile-phone-numbers retrieve", obj, format, transform)
}

func handleMobilePhoneNumbersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MobilePhoneNumberUpdateParams{}

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
	_, err = client.MobilePhoneNumbers.Update(
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
	return ShowJSON(os.Stdout, "mobile-phone-numbers update", obj, format, transform)
}

func handleMobilePhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MobilePhoneNumberListParams{}

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
		_, err = client.MobilePhoneNumbers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "mobile-phone-numbers list", obj, format, transform)
	} else {
		iter := client.MobilePhoneNumbers.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "mobile-phone-numbers list", iter, format, transform)
	}
}

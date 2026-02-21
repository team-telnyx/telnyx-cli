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

var phoneNumbersMessagingRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a phone number with messaging settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePhoneNumbersMessagingRetrieve,
	HideHelpCommand: true,
}

var phoneNumbersMessagingUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the messaging profile and/or messaging product of a phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "messaging-product",
			Usage:    "Configure the messaging product for this number:\n\n* Omit this field or set its value to `null` to keep the current value.\n* Set this field to a quoted product ID to set this phone number to that product",
			BodyPath: "messaging_product",
		},
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "Configure the messaging profile this phone number is assigned to:\n\n* Omit this field or set its value to `null` to keep the current value.\n* Set this field to `\"\"` to unassign the number from its messaging profile\n* Set this field to a quoted UUID of a messaging profile to assign this number to that messaging profile",
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to set on this phone number.",
			BodyPath: "tags",
		},
	},
	Action:          handlePhoneNumbersMessagingUpdate,
	HideHelpCommand: true,
}

var phoneNumbersMessagingList = cli.Command{
	Name:    "list",
	Usage:   "List phone numbers with messaging settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-messaging-profile-id",
			Usage:     "Filter by messaging profile ID.",
			QueryPath: "filter[messaging_profile_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-phone-number",
			Usage:     "Filter by exact phone number (supports comma-separated list).",
			QueryPath: "filter[phone_number]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-phone-number-contains",
			Usage:     "Filter by phone number substring.",
			QueryPath: "filter[phone_number][contains]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-type",
			Usage:     "Filter by phone number type.",
			QueryPath: "filter[type]",
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
			Name:      "sort-phone-number",
			Usage:     "Sort by phone number.",
			QueryPath: "sort[phone_number]",
		},
	},
	Action:          handlePhoneNumbersMessagingList,
	HideHelpCommand: true,
}

func handlePhoneNumbersMessagingRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PhoneNumbers.Messaging.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers:messaging retrieve", obj, format, transform)
}

func handlePhoneNumbersMessagingUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberMessagingUpdateParams{}

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
	_, err = client.PhoneNumbers.Messaging.Update(
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
	return ShowJSON(os.Stdout, "phone-numbers:messaging update", obj, format, transform)
}

func handlePhoneNumbersMessagingList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberMessagingListParams{}

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
		_, err = client.PhoneNumbers.Messaging.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "phone-numbers:messaging list", obj, format, transform)
	} else {
		iter := client.PhoneNumbers.Messaging.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "phone-numbers:messaging list", iter, format, transform)
	}
}

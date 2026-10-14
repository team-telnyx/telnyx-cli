// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var whatsappPhoneNumbersCallingRoutingList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve the connection ID currently stored for a WhatsApp Calling BYON (Bring\nYour Own Number) phone number: the connection that inbound WhatsApp calls to the\nnumber are delivered to.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleWhatsappPhoneNumbersCallingRoutingList,
	HideHelpCommand: true,
}

var whatsappPhoneNumbersCallingRoutingPatchAll = cli.Command{
	Name:    "patch-all",
	Usage:   "Set or clear the connection ID for a WhatsApp Calling BYON (Bring Your Own\nNumber) phone number: the connection that inbound WhatsApp calls to the number\nare delivered to.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "connection-id",
			Usage:    "ID of the connection to deliver inbound WhatsApp calls to: a positive integer up to 9223372036854775807, sent as a decimal string or an integer. Send a string to keep large IDs exact. Non-null values are returned as strings. `null` clears the routing.",
			Required: true,
			BodyPath: "connection_id",
		},
	},
	Action:          handleWhatsappPhoneNumbersCallingRoutingPatchAll,
	HideHelpCommand: true,
}

func handleWhatsappPhoneNumbersCallingRoutingList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Whatsapp.PhoneNumbers.CallingRouting.List(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "whatsapp:phone-numbers:calling-routing list",
		Transform:      transform,
	})
}

func handleWhatsappPhoneNumbersCallingRoutingPatchAll(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.WhatsappPhoneNumberCallingRoutingPatchAllParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Whatsapp.PhoneNumbers.CallingRouting.PatchAll(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "whatsapp:phone-numbers:calling-routing patch-all",
		Transform:      transform,
	})
}

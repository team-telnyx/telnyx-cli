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

var whatsappPhoneNumbersConversationalComponentsList = cli.Command{
	Name:    "list",
	Usage:   "Get phone number conversational components",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Required:  true,
			PathParam: "phone_number",
		},
	},
	Action:          handleWhatsappPhoneNumbersConversationalComponentsList,
	HideHelpCommand: true,
}

var whatsappPhoneNumbersConversationalComponentsPatchAll = requestflag.WithInnerFlags(cli.Command{
	Name:    "patch-all",
	Usage:   "Update phone number conversational components",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Required:  true,
			PathParam: "phone_number",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "command",
			Usage:    "List of commands",
			BodyPath: "commands",
		},
		&requestflag.Flag[[]string]{
			Name:     "ice-breaker",
			Usage:    "List of ice breakers",
			BodyPath: "ice_breakers",
		},
	},
	Action:          handleWhatsappPhoneNumbersConversationalComponentsPatchAll,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"command": {
		&requestflag.InnerFlag[string]{
			Name:       "command.command",
			InnerField: "command",
		},
		&requestflag.InnerFlag[string]{
			Name:       "command.description",
			InnerField: "description",
		},
	},
})

func handleWhatsappPhoneNumbersConversationalComponentsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
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
	_, err = client.Whatsapp.PhoneNumbers.ConversationalComponents.List(ctx, cmd.Value("phone-number").(string), options...)
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
		Title:          "whatsapp:phone-numbers:conversational-components list",
		Transform:      transform,
	})
}

func handleWhatsappPhoneNumbersConversationalComponentsPatchAll(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
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

	params := telnyx.WhatsappPhoneNumberConversationalComponentPatchAllParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Whatsapp.PhoneNumbers.ConversationalComponents.PatchAll(
		ctx,
		cmd.Value("phone-number").(string),
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
		Title:          "whatsapp:phone-numbers:conversational-components patch-all",
		Transform:      transform,
	})
}

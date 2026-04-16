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

var whatsappPhoneNumbersList = cli.Command{
	Name:    "list",
	Usage:   "List Whatsapp phone numbers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleWhatsappPhoneNumbersList,
	HideHelpCommand: true,
}

var whatsappPhoneNumbersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Whatsapp phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
		},
	},
	Action:          handleWhatsappPhoneNumbersDelete,
	HideHelpCommand: true,
}

var whatsappPhoneNumbersResendVerification = cli.Command{
	Name:    "resend-verification",
	Usage:   "Resend verification code",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "verification-method",
			Usage:    `Allowed values: "sms", "voice".`,
			Default:  "sms",
			BodyPath: "verification_method",
		},
	},
	Action:          handleWhatsappPhoneNumbersResendVerification,
	HideHelpCommand: true,
}

var whatsappPhoneNumbersVerify = cli.Command{
	Name:    "verify",
	Usage:   "Submit verification code for a phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "code",
			Required: true,
			BodyPath: "code",
		},
	},
	Action:          handleWhatsappPhoneNumbersVerify,
	HideHelpCommand: true,
}

func handleWhatsappPhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.WhatsappPhoneNumberListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Whatsapp.PhoneNumbers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "whatsapp:phone-numbers list", obj, format, explicitFormat, transform)
	} else {
		iter := client.Whatsapp.PhoneNumbers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "whatsapp:phone-numbers list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleWhatsappPhoneNumbersDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Whatsapp.PhoneNumbers.Delete(ctx, cmd.Value("phone-number").(string), options...)
}

func handleWhatsappPhoneNumbersResendVerification(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.WhatsappPhoneNumberResendVerificationParams{}

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

	return client.Whatsapp.PhoneNumbers.ResendVerification(
		ctx,
		cmd.Value("phone-number").(string),
		params,
		options...,
	)
}

func handleWhatsappPhoneNumbersVerify(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.WhatsappPhoneNumberVerifyParams{}

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

	return client.Whatsapp.PhoneNumbers.Verify(
		ctx,
		cmd.Value("phone-number").(string),
		params,
		options...,
	)
}

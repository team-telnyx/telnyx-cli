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

var whatsappBusinessAccountsPhoneNumbersList = cli.Command{
	Name:    "list",
	Usage:   "List phone numbers for a WABA",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
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
	Action:          handleWhatsappBusinessAccountsPhoneNumbersList,
	HideHelpCommand: true,
}

var whatsappBusinessAccountsPhoneNumbersInitializeVerification = cli.Command{
	Name:    "initialize-verification",
	Usage:   "Initialize Whatsapp phone number verification",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Required: true,
			BodyPath: "display_name",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
			BodyPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Default:  "en_US",
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "verification-method",
			Usage:    `Allowed values: "sms", "voice".`,
			Default:  "sms",
			BodyPath: "verification_method",
		},
	},
	Action:          handleWhatsappBusinessAccountsPhoneNumbersInitializeVerification,
	HideHelpCommand: true,
}

func handleWhatsappBusinessAccountsPhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.WhatsappBusinessAccountPhoneNumberListParams{}

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
		_, err = client.Whatsapp.BusinessAccounts.PhoneNumbers.List(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "whatsapp:business-accounts:phone-numbers list", obj, format, explicitFormat, transform)
	} else {
		iter := client.Whatsapp.BusinessAccounts.PhoneNumbers.ListAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "whatsapp:business-accounts:phone-numbers list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleWhatsappBusinessAccountsPhoneNumbersInitializeVerification(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.WhatsappBusinessAccountPhoneNumberInitializeVerificationParams{}

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

	return client.Whatsapp.BusinessAccounts.PhoneNumbers.InitializeVerification(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

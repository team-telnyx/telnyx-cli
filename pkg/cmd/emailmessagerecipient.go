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

var emailMessagesRecipientsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the current delivery state of a single recipient, including status,\nbillable flag, SMTP detail, and lifecycle timestamps. BCC recipient addresses\nare redacted (returned as null).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email-id",
			Required:  true,
			PathParam: "email_id",
		},
		&requestflag.Flag[string]{
			Name:      "recipient-id",
			Required:  true,
			PathParam: "recipient_id",
		},
	},
	Action:          handleEmailMessagesRecipientsRetrieve,
	HideHelpCommand: true,
}

var emailMessagesRecipientsList = cli.Command{
	Name:    "list",
	Usage:   "Lists per-recipient delivery states for a single message with cursor pagination.\nEach recipient has an independent status, billable flag, and lifecycle\ntimestamps. BCC recipient addresses are redacted (returned as null) to protect\nBCC privacy. Default page size is 25, maximum is 100.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email-id",
			Required:  true,
			PathParam: "email_id",
		},
		&requestflag.Flag[string]{
			Name:      "kind",
			Usage:     "Filter recipients by address kind.",
			QueryPath: "kind",
		},
		&requestflag.Flag[string]{
			Name:      "page-cursor",
			Usage:     "Opaque URL-safe Base64 cursor returned by a previous list response.",
			QueryPath: "page_cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results to return. Defaults to 25; maximum is 100. Invalid values are clamped to the valid range.",
			Default:   25,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter recipients by status.",
			QueryPath: "status",
		},
	},
	Action:          handleEmailMessagesRecipientsList,
	HideHelpCommand: true,
}

func handleEmailMessagesRecipientsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recipient-id") && len(unusedArgs) > 0 {
		cmd.Set("recipient-id", unusedArgs[0])
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

	params := telnyx.EmailMessageRecipientGetParams{
		EmailID: cmd.Value("email-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailMessages.Recipients.Get(
		ctx,
		cmd.Value("recipient-id").(string),
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
		Title:          "email-messages:recipients retrieve",
		Transform:      transform,
	})
}

func handleEmailMessagesRecipientsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("email-id") && len(unusedArgs) > 0 {
		cmd.Set("email-id", unusedArgs[0])
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

	params := telnyx.EmailMessageRecipientListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailMessages.Recipients.List(
		ctx,
		cmd.Value("email-id").(string),
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
		Title:          "email-messages:recipients list",
		Transform:      transform,
	})
}

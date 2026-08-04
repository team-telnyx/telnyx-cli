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

var emailInboxesMessagesActionsForward = cli.Command{
	Name:    "forward",
	Usage:   "Sends from the inbox address through the standard email send pipeline to\ncaller-supplied To, Cc, and Bcc recipients. `to` must contain at least one\nrecipient. Optional `text` and `html` are prepended to a forwarded-message block\ncontaining the original metadata and available body content. The subject is\nprefixed with `Fwd:` unless it already has that prefix.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "message-id",
			Required:  true,
			PathParam: "message_id",
		},
		&requestflag.Flag[any]{
			Name:     "to",
			Usage:    "One recipient or a non-empty recipient array. Each recipient may be an email string or an object with `email` and optional `name`.",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[any]{
			Name:     "bcc",
			Usage:    "One recipient or a recipient array. Each recipient may be an email string or an object with `email` and optional `name`.",
			BodyPath: "bcc",
		},
		&requestflag.Flag[any]{
			Name:     "cc",
			Usage:    "One recipient or a recipient array. Each recipient may be an email string or an object with `email` and optional `name`.",
			BodyPath: "cc",
		},
		&requestflag.Flag[string]{
			Name:     "html",
			Usage:    "Optional HTML note prepended to the generated forwarded-message block. Blank values are treated as omitted.",
			BodyPath: "html",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Optional plain-text note prepended to the generated forwarded-message block. Blank values are treated as omitted.",
			BodyPath: "text",
		},
	},
	Action:          handleEmailInboxesMessagesActionsForward,
	HideHelpCommand: true,
}

var emailInboxesMessagesActionsReply = cli.Command{
	Name:    "reply",
	Usage:   "Sends from the inbox address through the standard email send pipeline. The\nrecipient is the original `Reply-To`, falling back to `From`; original Cc\nrecipients are not included. The subject is prefixed with `Re:` unless it\nalready has that prefix.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "message-id",
			Required:  true,
			PathParam: "message_id",
		},
		&requestflag.Flag[string]{
			Name:     "html",
			Usage:    "HTML reply body.",
			BodyPath: "html",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Plain-text reply body.",
			BodyPath: "text",
		},
	},
	Action:          handleEmailInboxesMessagesActionsReply,
	HideHelpCommand: true,
}

var emailInboxesMessagesActionsReplyAll = cli.Command{
	Name:    "reply-all",
	Usage:   "Sends from the inbox address through the standard email send pipeline. The To\nlist starts with the original `Reply-To` (or `From`) and includes original To\nrecipients; the Cc list includes original Cc recipients. The inbox address is\nexcluded, and recipients are de-duplicated case-insensitively across To and Cc.\nBcc is always empty. The subject is prefixed with `Re:` unless it already has\nthat prefix.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "message-id",
			Required:  true,
			PathParam: "message_id",
		},
		&requestflag.Flag[string]{
			Name:     "html",
			Usage:    "HTML reply body.",
			BodyPath: "html",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Plain-text reply body.",
			BodyPath: "text",
		},
	},
	Action:          handleEmailInboxesMessagesActionsReplyAll,
	HideHelpCommand: true,
}

func handleEmailInboxesMessagesActionsForward(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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

	params := telnyx.EmailInboxMessageActionForwardParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Messages.Actions.Forward(
		ctx,
		cmd.Value("message-id").(string),
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
		Title:          "email-inboxes:messages:actions forward",
		Transform:      transform,
	})
}

func handleEmailInboxesMessagesActionsReply(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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

	params := telnyx.EmailInboxMessageActionReplyParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Messages.Actions.Reply(
		ctx,
		cmd.Value("message-id").(string),
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
		Title:          "email-inboxes:messages:actions reply",
		Transform:      transform,
	})
}

func handleEmailInboxesMessagesActionsReplyAll(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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

	params := telnyx.EmailInboxMessageActionReplyAllParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Messages.Actions.ReplyAll(
		ctx,
		cmd.Value("message-id").(string),
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
		Title:          "email-inboxes:messages:actions reply-all",
		Transform:      transform,
	})
}

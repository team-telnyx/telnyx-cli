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

var emailInboxesMessagesLabelsCreate = cli.Command{
	Name:    "create",
	Usage:   "Adds one or more mutable labels to a message. Labels carry agent workflow state\nsuch as `spam`, `needs_review`, or `processed`.",
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
		&requestflag.Flag[[]string]{
			Name:     "label",
			Usage:    "One or more labels. Each label is a freeform, case-sensitive string of at most 255 characters; a message or thread may carry at most 50 labels. The `telnyx:` prefix is a reserved system namespace and is rejected on customer writes.",
			Required: true,
			BodyPath: "labels",
		},
	},
	Action:          handleEmailInboxesMessagesLabelsCreate,
	HideHelpCommand: true,
}

var emailInboxesMessagesLabelsDeleteAll = cli.Command{
	Name:    "delete-all",
	Usage:   "Removes one or more labels from a message. Idempotent — removing a label the\nmessage does not carry is a no-op and still returns 200. Removal is\ncase-sensitive.",
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
		&requestflag.Flag[[]string]{
			Name:     "label",
			Usage:    "One or more labels. Each label is a freeform, case-sensitive string of at most 255 characters; a message or thread may carry at most 50 labels. The `telnyx:` prefix is a reserved system namespace and is rejected on customer writes.",
			Required: true,
			BodyPath: "labels",
		},
	},
	Action:          handleEmailInboxesMessagesLabelsDeleteAll,
	HideHelpCommand: true,
}

func handleEmailInboxesMessagesLabelsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailInboxMessageLabelNewParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Messages.Labels.New(
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
		Title:          "email-inboxes:messages:labels create",
		Transform:      transform,
	})
}

func handleEmailInboxesMessagesLabelsDeleteAll(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailInboxMessageLabelDeleteAllParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Messages.Labels.DeleteAll(
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
		Title:          "email-inboxes:messages:labels delete-all",
		Transform:      transform,
	})
}

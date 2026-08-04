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

var emailInboxesDraftsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates an unsent draft in the inbox. Every field is optional — a draft is a\nwork-in-progress and may be saved incomplete. Send-time requirements (sender,\nsubject, at least one recipient) are enforced when the draft is sent, not when\nit is created.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[[]any]{
			Name:     "attachment",
			BodyPath: "attachments",
		},
		&requestflag.Flag[[]any]{
			Name:     "bcc",
			BodyPath: "bcc",
		},
		&requestflag.Flag[[]any]{
			Name:     "cc",
			BodyPath: "cc",
		},
		&requestflag.Flag[string]{
			Name:     "from-email",
			BodyPath: "from_email",
		},
		&requestflag.Flag[string]{
			Name:     "from-name",
			BodyPath: "from_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			BodyPath: "headers",
		},
		&requestflag.Flag[string]{
			Name:     "html",
			Usage:    "Alias for `html_body`, matching the send endpoint.",
			BodyPath: "html",
		},
		&requestflag.Flag[string]{
			Name:     "html-body",
			BodyPath: "html_body",
		},
		&requestflag.Flag[[]string]{
			Name:     "label",
			BodyPath: "labels",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "reply-to",
			BodyPath: "reply_to",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			BodyPath: "subject",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Alias for `text_body`, matching the send endpoint.",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "text-body",
			BodyPath: "text_body",
		},
		&requestflag.Flag[[]any]{
			Name:     "to",
			BodyPath: "to",
		},
	},
	Action:          handleEmailInboxesDraftsCreate,
	HideHelpCommand: true,
}

var emailInboxesDraftsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a single draft. Drafts that have been sent remain retrievable, so the\nexact content that was sent stays auditable.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "draft-id",
			Required:  true,
			PathParam: "draft_id",
		},
	},
	Action:          handleEmailInboxesDraftsRetrieve,
	HideHelpCommand: true,
}

var emailInboxesDraftsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the supplied fields on a draft. `account_id` and `inbox_id` are\nserver-owned and ignored if present in the body, so a draft can never be moved\nbetween accounts or inboxes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "draft-id",
			Required:  true,
			PathParam: "draft_id",
		},
		&requestflag.Flag[[]any]{
			Name:     "attachment",
			BodyPath: "attachments",
		},
		&requestflag.Flag[[]any]{
			Name:     "bcc",
			BodyPath: "bcc",
		},
		&requestflag.Flag[[]any]{
			Name:     "cc",
			BodyPath: "cc",
		},
		&requestflag.Flag[string]{
			Name:     "from-email",
			BodyPath: "from_email",
		},
		&requestflag.Flag[string]{
			Name:     "from-name",
			BodyPath: "from_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			BodyPath: "headers",
		},
		&requestflag.Flag[string]{
			Name:     "html",
			Usage:    "Alias for `html_body`, matching the send endpoint.",
			BodyPath: "html",
		},
		&requestflag.Flag[string]{
			Name:     "html-body",
			BodyPath: "html_body",
		},
		&requestflag.Flag[[]string]{
			Name:     "label",
			BodyPath: "labels",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "reply-to",
			BodyPath: "reply_to",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			BodyPath: "subject",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Alias for `text_body`, matching the send endpoint.",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "text-body",
			BodyPath: "text_body",
		},
		&requestflag.Flag[[]any]{
			Name:     "to",
			BodyPath: "to",
		},
	},
	Action:          handleEmailInboxesDraftsUpdate,
	HideHelpCommand: true,
}

var emailInboxesDraftsList = cli.Command{
	Name:    "list",
	Usage:   "Lists drafts newest first using stable cursor pagination. All access is scoped\nto the authenticated account and the given inbox.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     "Restrict results to drafts in this state.",
			QueryPath: "filter[status]",
		},
		&requestflag.Flag[string]{
			Name:      "page-after",
			Usage:     "Opaque cursor returned by the previous page.",
			QueryPath: "page[after]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results to return. Defaults to 25; maximum is 100.",
			Default:   25,
			QueryPath: "page[size]",
		},
	},
	Action:          handleEmailInboxesDraftsList,
	HideHelpCommand: true,
}

var emailInboxesDraftsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes an unsent draft. Drafts that are being sent or have been\nsent cannot be deleted; sent drafts are retained for audit.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "draft-id",
			Required:  true,
			PathParam: "draft_id",
		},
	},
	Action:          handleEmailInboxesDraftsDelete,
	HideHelpCommand: true,
}

var emailInboxesDraftsPatch = cli.Command{
	Name:    "patch",
	Usage:   "Identical to `PUT`; both apply a partial update to the supplied fields.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "draft-id",
			Required:  true,
			PathParam: "draft_id",
		},
		&requestflag.Flag[[]any]{
			Name:     "attachment",
			BodyPath: "attachments",
		},
		&requestflag.Flag[[]any]{
			Name:     "bcc",
			BodyPath: "bcc",
		},
		&requestflag.Flag[[]any]{
			Name:     "cc",
			BodyPath: "cc",
		},
		&requestflag.Flag[string]{
			Name:     "from-email",
			BodyPath: "from_email",
		},
		&requestflag.Flag[string]{
			Name:     "from-name",
			BodyPath: "from_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			BodyPath: "headers",
		},
		&requestflag.Flag[string]{
			Name:     "html",
			Usage:    "Alias for `html_body`, matching the send endpoint.",
			BodyPath: "html",
		},
		&requestflag.Flag[string]{
			Name:     "html-body",
			BodyPath: "html_body",
		},
		&requestflag.Flag[[]string]{
			Name:     "label",
			BodyPath: "labels",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "reply-to",
			BodyPath: "reply_to",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			BodyPath: "subject",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Alias for `text_body`, matching the send endpoint.",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "text-body",
			BodyPath: "text_body",
		},
		&requestflag.Flag[[]any]{
			Name:     "to",
			BodyPath: "to",
		},
	},
	Action:          handleEmailInboxesDraftsPatch,
	HideHelpCommand: true,
}

var emailInboxesDraftsSend = cli.Command{
	Name:    "send",
	Usage:   "Sends the draft through the standard send pipeline — the same domain resolution,\nsuppression, reputation, daily-quota, persistence and Detail Record behaviour as\n`POST /v2/email_messages`. The response body is the created email message.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "draft-id",
			Required:  true,
			PathParam: "draft_id",
		},
	},
	Action:          handleEmailInboxesDraftsSend,
	HideHelpCommand: true,
}

func handleEmailInboxesDraftsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
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

	params := telnyx.EmailInboxDraftNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Drafts.New(
		ctx,
		cmd.Value("inbox-id").(string),
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
		Title:          "email-inboxes:drafts create",
		Transform:      transform,
	})
}

func handleEmailInboxesDraftsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("draft-id") && len(unusedArgs) > 0 {
		cmd.Set("draft-id", unusedArgs[0])
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

	params := telnyx.EmailInboxDraftGetParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Drafts.Get(
		ctx,
		cmd.Value("draft-id").(string),
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
		Title:          "email-inboxes:drafts retrieve",
		Transform:      transform,
	})
}

func handleEmailInboxesDraftsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("draft-id") && len(unusedArgs) > 0 {
		cmd.Set("draft-id", unusedArgs[0])
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

	params := telnyx.EmailInboxDraftUpdateParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Drafts.Update(
		ctx,
		cmd.Value("draft-id").(string),
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
		Title:          "email-inboxes:drafts update",
		Transform:      transform,
	})
}

func handleEmailInboxesDraftsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
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

	params := telnyx.EmailInboxDraftListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Drafts.List(
		ctx,
		cmd.Value("inbox-id").(string),
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
		Title:          "email-inboxes:drafts list",
		Transform:      transform,
	})
}

func handleEmailInboxesDraftsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("draft-id") && len(unusedArgs) > 0 {
		cmd.Set("draft-id", unusedArgs[0])
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

	params := telnyx.EmailInboxDraftDeleteParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	return client.EmailInboxes.Drafts.Delete(
		ctx,
		cmd.Value("draft-id").(string),
		params,
		options...,
	)
}

func handleEmailInboxesDraftsPatch(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("draft-id") && len(unusedArgs) > 0 {
		cmd.Set("draft-id", unusedArgs[0])
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

	params := telnyx.EmailInboxDraftPatchParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Drafts.Patch(
		ctx,
		cmd.Value("draft-id").(string),
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
		Title:          "email-inboxes:drafts patch",
		Transform:      transform,
	})
}

func handleEmailInboxesDraftsSend(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("draft-id") && len(unusedArgs) > 0 {
		cmd.Set("draft-id", unusedArgs[0])
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

	params := telnyx.EmailInboxDraftSendParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Drafts.Send(
		ctx,
		cmd.Value("draft-id").(string),
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
		Title:          "email-inboxes:drafts send",
		Transform:      transform,
	})
}

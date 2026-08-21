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

var emailInboxesMessagesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the explicit read state of an account-scoped inbound message. Set\n`read_at` to `true` to mark the message read at the server's current time, to an\nISO 8601 timestamp to use that timestamp, or to `null` to mark the message\nunread. Repeating the same update is idempotent.",
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
			Name:     "read-at",
			Required: true,
			BodyPath: "read_at",
		},
	},
	Action:          handleEmailInboxesMessagesUpdate,
	HideHelpCommand: true,
}

var emailInboxesMessagesList = cli.Command{
	Name:    "list",
	Usage:   "Lists inbound messages newest first. All access is scoped to the authenticated\naccount. `filter[search]` performs PostgreSQL full-text search over the subject,\nplain-text body, and HTML body. Filters compose with stable cursor pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "filter-from",
			Usage:     "Case-insensitive literal substring of the sender address.",
			QueryPath: "filter[from]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-label",
			Usage:     "Returns only messages carrying this label. Matching is exact and case-sensitive. Reserved `telnyx:` labels can be filtered on even though they cannot be written by customers.",
			QueryPath: "filter[label]",
		},
		&requestflag.Flag[bool]{
			Name:      "filter-read",
			Usage:     "Whether the message has a read timestamp.",
			QueryPath: "filter[read]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-received-after",
			Usage:     "Inclusive ISO 8601 lower bound for the received timestamp.",
			QueryPath: "filter[received_after]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-received-before",
			Usage:     "Inclusive ISO 8601 upper bound for the received timestamp.",
			QueryPath: "filter[received_before]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-search",
			Usage:     "Full-text query over subject and body, up to 500 characters.",
			QueryPath: "filter[search]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-subject",
			Usage:     "Case-insensitive literal substring of the subject.",
			QueryPath: "filter[subject]",
		},
		&requestflag.Flag[bool]{
			Name:      "filter-unread",
			Usage:     "Whether the message has no read timestamp. Set to `true` to return only unread messages.",
			QueryPath: "filter[unread]",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEmailInboxesMessagesList,
	HideHelpCommand: true,
}

var emailInboxesMessagesDrafts = cli.Command{
	Name:    "drafts",
	Usage:   "Creates an unsent reply draft for an inbound message. Unlike the\n`/actions/reply` endpoint, which sends immediately, this stores a draft that can\nbe reviewed and edited before sending.",
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
		&requestflag.Flag[[]map[string]any]{
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
		&requestflag.Flag[map[string]any]{
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
	Action:          handleEmailInboxesMessagesDrafts,
	HideHelpCommand: true,
}

func handleEmailInboxesMessagesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailInboxMessageUpdateParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Messages.Update(
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
		Title:          "email-inboxes:messages update",
		Transform:      transform,
	})
}

func handleEmailInboxesMessagesList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailInboxMessageListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailInboxes.Messages.List(
			ctx,
			cmd.Value("inbox-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-inboxes:messages list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailInboxes.Messages.ListAutoPaging(
			ctx,
			cmd.Value("inbox-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-inboxes:messages list",
			Transform:      transform,
		})
	}
}

func handleEmailInboxesMessagesDrafts(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailInboxMessageDraftsParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Messages.Drafts(
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
		Title:          "email-inboxes:messages drafts",
		Transform:      transform,
	})
}

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

var emailMessagesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Queues, schedules, or sandbox-sends an email message. The legacy `/v2/emails`\nPOST route is a backward-compatible alias for this operation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "from",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[[]any]{
			Name:     "to",
			Required: true,
			BodyPath: "to",
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
		&requestflag.Flag[*string]{
			Name:     "forward-of-message-id",
			Usage:    "Telnyx message UUID of the message this send forwards. Forwarded\nmessages start a NEW thread per RFC 5322 — NO `In-Reply-To` or\n`References` headers are set on the outbound MIME. The id is\nrecorded in the message's metadata for EDR provenance only.\n\nThe id is validated as a UUID but is NOT looked up against the\nmessage store — existence is the caller's responsibility (the\nforward is pure metadata; it does not affect delivery). Cannot be\ncombined with `in_reply_to_message_id` (422).",
			BodyPath: "forward_of_message_id",
		},
		&requestflag.Flag[string]{
			Name:     "from-name",
			Usage:    "Optional display name for string `from`; overrides `from.name` when provided.",
			BodyPath: "from_name",
		},
		&requestflag.Flag[*string]{
			Name:     "group-id",
			Usage:    "Optional unsubscribe-group UUID used for group-scoped suppression checks and unsubscribe handling.",
			BodyPath: "group_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			Usage:    "Custom email headers. Write-only; not returned in responses.",
			BodyPath: "headers",
		},
		&requestflag.Flag[string]{
			Name:     "html-body",
			Usage:    "HTML email body. Returned only by `GET /email_messages/{id}`; omitted from create and list responses.",
			BodyPath: "html_body",
		},
		&requestflag.Flag[bool]{
			Name:     "ignore-suppression",
			Usage:    "When true, allows delivery to recipients whose suppressions explicitly\npermit an override. Hard bounces, spam complaints, and invalid-address\nsuppressions cannot be overridden. Requires the `email:override` API scope.\n",
			Default:  false,
			BodyPath: "ignore_suppression",
		},
		&requestflag.Flag[*string]{
			Name:     "in-reply-to-message-id",
			Usage:    "Telnyx message UUID of the message this send replies to. When provided,\nthe API sets RFC 5322 `In-Reply-To` and `References` headers on the\noutbound MIME so the recipient's mailbox (Gmail/Outlook) threads it\ncorrectly. The parent is looked up under the caller's account scope;\na UUID belonging to another account yields a non-enumerating 404.\n\nWire-only (Phase 1): the API sets the headers and does NOT resolve or\nmutate `thread_id` on the server side. Messages sent without this\nparameter are standalone (no threading headers injected).\n\nCannot be combined with `forward_of_message_id` (422).",
			BodyPath: "in_reply_to_message_id",
		},
		&requestflag.Flag[bool]{
			Name:     "inline-css",
			Default:  false,
			BodyPath: "inline_css",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Custom metadata. Write-only; not returned in responses.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "reply-to",
			BodyPath: "reply_to",
		},
		&requestflag.Flag[*bool]{
			Name:     "reply-to-all",
			Usage:    "Indicates a reply-all intent. In Phase 1 (wire-only) this does not\nchange the threading headers — recipient selection is customer-\ncontrolled (`to`/`cc`), and a thread is not defined by its audience.\nWhen the referenced message has no thread context, reply-all\ndegrades to a plain reply (parent ID only in `References`). The\nresolution engine (separate work) will expand the ancestor chain\nat a later phase with no API change.\n\nOnly meaningful alongside `in_reply_to_message_id`.",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "reply_to_all",
		},
		&requestflag.Flag[bool]{
			Name:     "sandbox-mode",
			Default:  false,
			BodyPath: "sandbox_mode",
		},
		&requestflag.Flag[any]{
			Name:     "scheduled-at",
			Usage:    "Future ISO 8601 time to schedule sending. Invalid or past timestamps\nare silently ignored and the email is sent immediately. The legacy\nalias `send_at` is still accepted for backward compatibility; when\nboth are provided, `scheduled_at` wins.\n",
			BodyPath: "scheduled_at",
		},
		&requestflag.Flag[any]{
			Name:     "send-at",
			Usage:    "Deprecated alias for `scheduled_at`.",
			BodyPath: "send_at",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Required unless `template_id` is supplied. When using a template, the template's subject is rendered; if the template has no subject or renders empty, the request returns 400.",
			BodyPath: "subject",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags for categorization and reporting. Stored on the message and propagated to Email Detail Records. Not returned in API responses.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "template-id",
			BodyPath: "template_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "template-variables",
			Usage:    "Variables for Liquid template rendering. Non-object values may cause a 422 validation error on message creation, but are silently treated as an empty object for template rendering.",
			Default:  map[string]any{},
			BodyPath: "template_variables",
		},
		&requestflag.Flag[string]{
			Name:     "text-body",
			Usage:    "Plain text email body. Returned only by `GET /email_messages/{id}`; omitted from create and list responses.",
			BodyPath: "text_body",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "tracking-settings",
			Usage:    "Per-send open and click tracking overrides. Omitted properties inherit the sender domain's tracking settings.",
			BodyPath: "tracking_settings",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleEmailMessagesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"attachment": {
		&requestflag.InnerFlag[string]{
			Name:       "attachment.content",
			Usage:      "Attachment content, typically Base64-encoded. Defaults to empty string when omitted.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "attachment.content-id",
			Usage:      "MIME Content-ID used to reference an inline attachment.",
			InnerField: "content_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "attachment.content-type",
			Usage:      `MIME content type. Defaults to "application/octet-stream" when omitted.`,
			InnerField: "content_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "attachment.disposition",
			Usage:      "MIME disposition (`attachment` or `inline`).",
			InnerField: "disposition",
		},
		&requestflag.InnerFlag[string]{
			Name:       "attachment.filename",
			Usage:      `Attachment filename. Defaults to "attachment" when omitted.`,
			InnerField: "filename",
		},
	},
	"tracking-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "tracking-settings.click-tracking",
			Usage:      "Whether to rewrite links for click tracking in this message.",
			InnerField: "click_tracking",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "tracking-settings.open-tracking",
			Usage:      "Whether to inject an open-tracking pixel for this message.",
			InnerField: "open_tracking",
		},
	},
})

var emailMessagesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "The legacy `/v2/emails/{id}` GET route is a backward-compatible alias for this\noperation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailMessagesRetrieve,
	HideHelpCommand: true,
}

var emailMessagesList = cli.Command{
	Name:    "list",
	Usage:   "Lists messages sorted newest first by `created_at desc, id desc`. No filters\nother than cursor pagination are implemented. The legacy `/v2/emails` GET route\nis a backward-compatible alias for this operation.",
	Suggest: true,
	Flags: []cli.Flag{
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
	},
	Action:          handleEmailMessagesList,
	HideHelpCommand: true,
}

var emailMessagesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes an account-scoped email message, its events, its durable\nrecipients, and unshared attachment objects. Returns 404 when the message does\nnot exist in the authenticated account. The legacy `/v2/emails/{id}` DELETE\nroute is a backward-compatible alias.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailMessagesDelete,
	HideHelpCommand: true,
}

var emailMessagesBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "batch",
	Usage:   "Creates up to 50 email messages in a single request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "message",
			Required: true,
			BodyPath: "messages",
		},
		&requestflag.Flag[bool]{
			Name:     "sandbox-mode",
			Usage:    "Applies sandbox mode to all messages in the batch. Overrides any per-message sandbox_mode in the messages array.",
			Default:  false,
			BodyPath: "sandbox_mode",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleEmailMessagesBatch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message": {
		&requestflag.InnerFlag[any]{
			Name:       "message.from",
			InnerField: "from",
		},
		&requestflag.InnerFlag[[]any]{
			Name:       "message.to",
			InnerField: "to",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "message.attachments",
			InnerField: "attachments",
		},
		&requestflag.InnerFlag[[]any]{
			Name:       "message.bcc",
			InnerField: "bcc",
		},
		&requestflag.InnerFlag[[]any]{
			Name:       "message.cc",
			InnerField: "cc",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.from-name",
			Usage:      "Optional display name for string `from`; overrides `from.name` when provided.",
			InnerField: "from_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "message.group-id",
			Usage:      "Optional unsubscribe-group UUID used for group-scoped suppression checks and unsubscribe handling.",
			InnerField: "group_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.headers",
			Usage:      "Custom email headers. Write-only; not returned in responses.",
			InnerField: "headers",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.html-body",
			Usage:      "HTML email body. Returned only by `GET /email_messages/{id}`; omitted from create and list responses.",
			InnerField: "html_body",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "message.ignore-suppression",
			Usage:      "When true, allows delivery to recipients whose suppressions explicitly\npermit an override. Hard bounces, spam complaints, and invalid-address\nsuppressions cannot be overridden. Requires the `email:override` API scope.\n",
			InnerField: "ignore_suppression",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "message.inline-css",
			InnerField: "inline_css",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.metadata",
			Usage:      "Custom metadata. Write-only; not returned in responses.",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[any]{
			Name:       "message.reply-to",
			InnerField: "reply_to",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "message.sandbox-mode",
			InnerField: "sandbox_mode",
		},
		&requestflag.InnerFlag[any]{
			Name:       "message.scheduled-at",
			Usage:      "Future ISO 8601 time to schedule sending. Invalid or past timestamps\nare silently ignored and the email is sent immediately. The legacy\nalias `send_at` is still accepted for backward compatibility; when\nboth are provided, `scheduled_at` wins.\n",
			InnerField: "scheduled_at",
		},
		&requestflag.InnerFlag[any]{
			Name:       "message.send-at",
			Usage:      "Deprecated alias for `scheduled_at`.",
			InnerField: "send_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.subject",
			Usage:      "Required unless `template_id` is supplied. When using a template, the template's subject is rendered; if the template has no subject or renders empty, the request returns 400.",
			InnerField: "subject",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "message.tags",
			Usage:      "Tags for categorization and reporting. Stored on the message and propagated to Email Detail Records. Not returned in API responses.",
			InnerField: "tags",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.template-id",
			InnerField: "template_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.template-variables",
			Usage:      "Variables for Liquid template rendering. Non-object values may cause a 422 validation error on message creation, but are silently treated as an empty object for template rendering.",
			InnerField: "template_variables",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.text-body",
			Usage:      "Plain text email body. Returned only by `GET /email_messages/{id}`; omitted from create and list responses.",
			InnerField: "text_body",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.tracking-settings",
			Usage:      "Per-send open and click tracking overrides. Omitted properties inherit the sender domain's tracking settings.",
			InnerField: "tracking_settings",
		},
	},
})

var emailMessagesDeleteAll = cli.Command{
	Name:    "delete-all",
	Usage:   "Permanently deletes every email in the authenticated account sent from or to the\nsupplied address, including retained events whose parent message has expired.\nEvents and durable recipients are deleted immediately with each message. The\noperation never searches or reports matches in another account. The legacy\n`/v2/emails` DELETE route is a backward-compatible alias.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "address",
			Usage:     "Sender or recipient address to delete. Matching is trimmed and case-insensitive.",
			Required:  true,
			QueryPath: "address",
		},
	},
	Action:          handleEmailMessagesDeleteAll,
	HideHelpCommand: true,
}

var emailMessagesDeleteSchedule = cli.Command{
	Name:    "delete-schedule",
	Usage:   "Cancels a scheduled email and returns it with status `cancelled`. The legacy\n`/v2/emails/{id}/schedule` DELETE route is an alias.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email-id",
			Required:  true,
			PathParam: "email_id",
		},
	},
	Action:          handleEmailMessagesDeleteSchedule,
	HideHelpCommand: true,
}

var emailMessagesRetrieveEvents = cli.Command{
	Name:    "retrieve-events",
	Usage:   "Lists events for a single message sorted oldest first by\n`occurred_at asc, id asc`. The legacy `/v2/emails/{id}/events` GET route is a\nbackward-compatible alias.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email-id",
			Required:  true,
			PathParam: "email_id",
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
	},
	Action:          handleEmailMessagesRetrieveEvents,
	HideHelpCommand: true,
}

func handleEmailMessagesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.EmailMessageNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailMessages.New(ctx, params, options...)
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
		Title:          "email-messages create",
		Transform:      transform,
	})
}

func handleEmailMessagesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EmailMessages.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email-messages retrieve",
		Transform:      transform,
	})
}

func handleEmailMessagesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.EmailMessageListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailMessages.List(ctx, params, options...)
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
		Title:          "email-messages list",
		Transform:      transform,
	})
}

func handleEmailMessagesDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.EmailMessages.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleEmailMessagesBatch(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.EmailMessageBatchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailMessages.Batch(ctx, params, options...)
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
		Title:          "email-messages batch",
		Transform:      transform,
	})
}

func handleEmailMessagesDeleteAll(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.EmailMessageDeleteAllParams{}

	return client.EmailMessages.DeleteAll(ctx, params, options...)
}

func handleEmailMessagesDeleteSchedule(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailMessages.DeleteSchedule(ctx, cmd.Value("email-id").(string), options...)
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
		Title:          "email-messages delete-schedule",
		Transform:      transform,
	})
}

func handleEmailMessagesRetrieveEvents(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailMessageGetEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailMessages.GetEvents(
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
		Title:          "email-messages retrieve-events",
		Transform:      transform,
	})
}

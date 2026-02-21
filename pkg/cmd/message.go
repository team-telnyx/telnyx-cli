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

var messagesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Note: This API endpoint can only retrieve messages that are no older than 10\ndays since their creation. If you require messages older than this, please\ngenerate an\n[MDR report.](https://developers.telnyx.com/api-reference/mdr-usage-reports/create-mdr-usage-report)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMessagesRetrieve,
	HideHelpCommand: true,
}

var messagesCancelScheduled = cli.Command{
	Name:    "cancel-scheduled",
	Usage:   "Cancel a scheduled message that has not yet been sent. Only messages with\n`status=scheduled` and `send_at` more than a minute from now can be cancelled.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMessagesCancelScheduled,
	HideHelpCommand: true,
}

var messagesRetrieveGroupMessages = cli.Command{
	Name:    "retrieve-group-messages",
	Usage:   "Retrieve all messages in a group MMS conversation by the group message ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
	},
	Action:          handleMessagesRetrieveGroupMessages,
	HideHelpCommand: true,
}

var messagesSchedule = cli.Command{
	Name:    "schedule",
	Usage:   "Schedule a message with a Phone Number, Alphanumeric Sender ID, Short Code or\nNumber Pool.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Receiving address (+E.164 formatted phone number or short code).",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-detect",
			Usage:    "Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts.",
			Default:  false,
			BodyPath: "auto_detect",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code).\n\n**Required if sending with a phone number, short code, or alphanumeric sender ID.**\n",
			BodyPath: "from",
		},
		&requestflag.Flag[[]string]{
			Name:     "media-url",
			Usage:    "A list of media URLs. The total media size must be less than 1 MB.\n\n**Required for MMS**",
			BodyPath: "media_urls",
		},
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "Unique identifier for a messaging profile.\n\n**Required if sending via number pool or with an alphanumeric sender ID.**\n",
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[any]{
			Name:     "send-at",
			Usage:    "ISO 8601 formatted date indicating when to send the message - accurate up till a minute.",
			BodyPath: "send_at",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Subject of multimedia message",
			BodyPath: "subject",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Message body (i.e., content) as a non-empty string.\n\n**Required for SMS**",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The protocol for sending the message, either SMS or MMS.",
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "use-profile-webhooks",
			Usage:    "If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile.",
			Default:  true,
			BodyPath: "use_profile_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "The failover URL where webhooks related to this message will be sent if sending to the primary URL fails.",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The URL where webhooks related to this message will be sent.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesSchedule,
	HideHelpCommand: true,
}

var messagesSend = cli.Command{
	Name:    "send",
	Usage:   "Send a message with a Phone Number, Alphanumeric Sender ID, Short Code or Number\nPool.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Receiving address (+E.164 formatted phone number or short code).",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-detect",
			Usage:    "Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts.",
			Default:  false,
			BodyPath: "auto_detect",
		},
		&requestflag.Flag[string]{
			Name:     "encoding",
			Usage:    "Encoding to use for the message. `auto` (default) uses smart encoding to automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding (returns 400 if message contains characters that cannot be encoded). `ucs2` forces UCS-2 encoding and disables smart encoding. When set, this overrides the messaging profile's `smart_encoding` setting.",
			Default:  "auto",
			BodyPath: "encoding",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code).\n\n**Required if sending with a phone number, short code, or alphanumeric sender ID.**\n",
			BodyPath: "from",
		},
		&requestflag.Flag[[]string]{
			Name:     "media-url",
			Usage:    "A list of media URLs. The total media size must be less than 1 MB.\n\n**Required for MMS**",
			BodyPath: "media_urls",
		},
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "Unique identifier for a messaging profile.\n\n**Required if sending via number pool or with an alphanumeric sender ID.**\n",
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[any]{
			Name:     "send-at",
			Usage:    "ISO 8601 formatted date indicating when to send the message - accurate up till a minute.",
			BodyPath: "send_at",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Subject of multimedia message",
			BodyPath: "subject",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Message body (i.e., content) as a non-empty string.\n\n**Required for SMS**",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The protocol for sending the message, either SMS or MMS.",
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "use-profile-webhooks",
			Usage:    "If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile.",
			Default:  true,
			BodyPath: "use_profile_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "The failover URL where webhooks related to this message will be sent if sending to the primary URL fails.",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The URL where webhooks related to this message will be sent.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesSend,
	HideHelpCommand: true,
}

var messagesSendGroupMms = cli.Command{
	Name:    "send-group-mms",
	Usage:   "Send a group MMS message",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Phone number, in +E.164 format, used to send the message.",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[[]string]{
			Name:     "to",
			Usage:    "A list of destinations. No more than 8 destinations are allowed.",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[[]string]{
			Name:     "media-url",
			Usage:    "A list of media URLs. The total media size must be less than 1 MB.",
			BodyPath: "media_urls",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Subject of multimedia message",
			BodyPath: "subject",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Message body (i.e., content) as a non-empty string.",
			BodyPath: "text",
		},
		&requestflag.Flag[bool]{
			Name:     "use-profile-webhooks",
			Usage:    "If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile.",
			Default:  true,
			BodyPath: "use_profile_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "The failover URL where webhooks related to this message will be sent if sending to the primary URL fails.",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The URL where webhooks related to this message will be sent.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesSendGroupMms,
	HideHelpCommand: true,
}

var messagesSendLongCode = cli.Command{
	Name:    "send-long-code",
	Usage:   "Send a long code message",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Phone number, in +E.164 format, used to send the message.",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Receiving address (+E.164 formatted phone number or short code).",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-detect",
			Usage:    "Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts.",
			Default:  false,
			BodyPath: "auto_detect",
		},
		&requestflag.Flag[string]{
			Name:     "encoding",
			Usage:    "Encoding to use for the message. `auto` (default) uses smart encoding to automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding (returns 400 if message contains characters that cannot be encoded). `ucs2` forces UCS-2 encoding and disables smart encoding. When set, this overrides the messaging profile's `smart_encoding` setting.",
			Default:  "auto",
			BodyPath: "encoding",
		},
		&requestflag.Flag[[]string]{
			Name:     "media-url",
			Usage:    "A list of media URLs. The total media size must be less than 1 MB.\n\n**Required for MMS**",
			BodyPath: "media_urls",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Subject of multimedia message",
			BodyPath: "subject",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Message body (i.e., content) as a non-empty string.\n\n**Required for SMS**",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The protocol for sending the message, either SMS or MMS.",
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "use-profile-webhooks",
			Usage:    "If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile.",
			Default:  true,
			BodyPath: "use_profile_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "The failover URL where webhooks related to this message will be sent if sending to the primary URL fails.",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The URL where webhooks related to this message will be sent.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesSendLongCode,
	HideHelpCommand: true,
}

var messagesSendNumberPool = cli.Command{
	Name:    "send-number-pool",
	Usage:   "Send a message using number pool",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "Unique identifier for a messaging profile.",
			Required: true,
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Receiving address (+E.164 formatted phone number or short code).",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-detect",
			Usage:    "Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts.",
			Default:  false,
			BodyPath: "auto_detect",
		},
		&requestflag.Flag[string]{
			Name:     "encoding",
			Usage:    "Encoding to use for the message. `auto` (default) uses smart encoding to automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding (returns 400 if message contains characters that cannot be encoded). `ucs2` forces UCS-2 encoding and disables smart encoding. When set, this overrides the messaging profile's `smart_encoding` setting.",
			Default:  "auto",
			BodyPath: "encoding",
		},
		&requestflag.Flag[[]string]{
			Name:     "media-url",
			Usage:    "A list of media URLs. The total media size must be less than 1 MB.\n\n**Required for MMS**",
			BodyPath: "media_urls",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Subject of multimedia message",
			BodyPath: "subject",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Message body (i.e., content) as a non-empty string.\n\n**Required for SMS**",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The protocol for sending the message, either SMS or MMS.",
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "use-profile-webhooks",
			Usage:    "If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile.",
			Default:  true,
			BodyPath: "use_profile_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "The failover URL where webhooks related to this message will be sent if sending to the primary URL fails.",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The URL where webhooks related to this message will be sent.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesSendNumberPool,
	HideHelpCommand: true,
}

var messagesSendShortCode = cli.Command{
	Name:    "send-short-code",
	Usage:   "Send a short code message",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Phone number, in +E.164 format, used to send the message.",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Receiving address (+E.164 formatted phone number or short code).",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-detect",
			Usage:    "Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts.",
			Default:  false,
			BodyPath: "auto_detect",
		},
		&requestflag.Flag[string]{
			Name:     "encoding",
			Usage:    "Encoding to use for the message. `auto` (default) uses smart encoding to automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding (returns 400 if message contains characters that cannot be encoded). `ucs2` forces UCS-2 encoding and disables smart encoding. When set, this overrides the messaging profile's `smart_encoding` setting.",
			Default:  "auto",
			BodyPath: "encoding",
		},
		&requestflag.Flag[[]string]{
			Name:     "media-url",
			Usage:    "A list of media URLs. The total media size must be less than 1 MB.\n\n**Required for MMS**",
			BodyPath: "media_urls",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Subject of multimedia message",
			BodyPath: "subject",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Message body (i.e., content) as a non-empty string.\n\n**Required for SMS**",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The protocol for sending the message, either SMS or MMS.",
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "use-profile-webhooks",
			Usage:    "If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile.",
			Default:  true,
			BodyPath: "use_profile_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "The failover URL where webhooks related to this message will be sent if sending to the primary URL fails.",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The URL where webhooks related to this message will be sent.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesSendShortCode,
	HideHelpCommand: true,
}

var messagesSendWhatsapp = requestflag.WithInnerFlags(cli.Command{
	Name:    "send-whatsapp",
	Usage:   "Send a Whatsapp message",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Phone number in +E.164 format associated with Whatsapp account",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Phone number in +E.164 format",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "whatsapp-message",
			Required: true,
			BodyPath: "whatsapp_message",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Message type - must be set to "WHATSAPP"`,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The URL where webhooks related to this message will be sent.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesSendWhatsapp,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"whatsapp-message": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "whatsapp-message.audio",
			InnerField: "audio",
		},
		&requestflag.InnerFlag[string]{
			Name:       "whatsapp-message.biz-opaque-callback-data",
			Usage:      "custom data to return with status update",
			InnerField: "biz_opaque_callback_data",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "whatsapp-message.contacts",
			InnerField: "contacts",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "whatsapp-message.document",
			InnerField: "document",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "whatsapp-message.image",
			InnerField: "image",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "whatsapp-message.interactive",
			InnerField: "interactive",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "whatsapp-message.location",
			InnerField: "location",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "whatsapp-message.reaction",
			InnerField: "reaction",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "whatsapp-message.sticker",
			InnerField: "sticker",
		},
		&requestflag.InnerFlag[string]{
			Name:       "whatsapp-message.type",
			InnerField: "type",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "whatsapp-message.video",
			InnerField: "video",
		},
	},
})

var messagesSendWithAlphanumericSender = cli.Command{
	Name:    "send-with-alphanumeric-sender",
	Usage:   "Send an SMS message using an alphanumeric sender ID. This is SMS only.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "A valid alphanumeric sender ID on the user's account.",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "The messaging profile ID to use.",
			Required: true,
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "The message body.",
			Required: true,
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Receiving address (+E.164 formatted phone number).",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[bool]{
			Name:     "use-profile-webhooks",
			Usage:    "If true, use the messaging profile's webhook settings.",
			BodyPath: "use_profile_webhooks",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-failover-url",
			Usage:    "Failover callback URL for delivery status updates.",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-url",
			Usage:    "Callback URL for delivery status updates.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesSendWithAlphanumericSender,
	HideHelpCommand: true,
}

func handleMessagesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Messages.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages retrieve", obj, format, transform)
}

func handleMessagesCancelScheduled(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Messages.CancelScheduled(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages cancel-scheduled", obj, format, transform)
}

func handleMessagesRetrieveGroupMessages(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.GetGroupMessages(ctx, cmd.Value("message-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages retrieve-group-messages", obj, format, transform)
}

func handleMessagesSchedule(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageScheduleParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.Schedule(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages schedule", obj, format, transform)
}

func handleMessagesSend(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageSendParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.Send(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages send", obj, format, transform)
}

func handleMessagesSendGroupMms(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageSendGroupMmsParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.SendGroupMms(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages send-group-mms", obj, format, transform)
}

func handleMessagesSendLongCode(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageSendLongCodeParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.SendLongCode(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages send-long-code", obj, format, transform)
}

func handleMessagesSendNumberPool(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageSendNumberPoolParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.SendNumberPool(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages send-number-pool", obj, format, transform)
}

func handleMessagesSendShortCode(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageSendShortCodeParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.SendShortCode(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages send-short-code", obj, format, transform)
}

func handleMessagesSendWhatsapp(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageSendWhatsappParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.SendWhatsapp(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages send-whatsapp", obj, format, transform)
}

func handleMessagesSendWithAlphanumericSender(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageSendWithAlphanumericSenderParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.SendWithAlphanumericSender(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages send-with-alphanumeric-sender", obj, format, transform)
}

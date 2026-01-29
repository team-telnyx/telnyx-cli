// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var messagesRcsGenerateDeeplink = cli.Command{
	Name:    "generate-deeplink",
	Usage:   "Generate a deeplink URL that can be used to start an RCS conversation with a\nspecific agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "body",
			Usage:     "Pre-filled message body (URL encoded)",
			QueryPath: "body",
		},
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Usage:     "Phone number in E164 format (URL encoded)",
			QueryPath: "phone_number",
		},
	},
	Action:          handleMessagesRcsGenerateDeeplink,
	HideHelpCommand: true,
}

var messagesRcsSend = requestflag.WithInnerFlags(cli.Command{
	Name:    "send",
	Usage:   "Send an RCS message",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Usage:    "RCS Agent ID",
			Required: true,
			BodyPath: "agent_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "agent-message",
			Required: true,
			BodyPath: "agent_message",
		},
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "A valid messaging profile ID",
			Required: true,
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Phone number in +E.164 format",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "mms-fallback",
			BodyPath: "mms_fallback",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sms-fallback",
			BodyPath: "sms_fallback",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Message type - must be set to "RCS"`,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The URL where webhooks related to this message will be sent.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMessagesRcsSend,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"agent-message": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "agent-message.content-message",
			InnerField: "content_message",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "agent-message.event",
			Usage:      "RCS Event to send to the recipient",
			InnerField: "event",
		},
		&requestflag.InnerFlag[any]{
			Name:       "agent-message.expire-time",
			Usage:      "Timestamp in UTC of when this message is considered expired",
			InnerField: "expire_time",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent-message.ttl",
			Usage:      "Duration in seconds ending with 's'",
			InnerField: "ttl",
		},
	},
	"mms-fallback": {
		&requestflag.InnerFlag[string]{
			Name:       "mms-fallback.from",
			Usage:      "Phone number in +E.164 format",
			InnerField: "from",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "mms-fallback.media-urls",
			Usage:      "List of media URLs",
			InnerField: "media_urls",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mms-fallback.subject",
			Usage:      "Subject of the message",
			InnerField: "subject",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mms-fallback.text",
			Usage:      "Text",
			InnerField: "text",
		},
	},
	"sms-fallback": {
		&requestflag.InnerFlag[string]{
			Name:       "sms-fallback.from",
			Usage:      "Phone number in +E.164 format",
			InnerField: "from",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sms-fallback.text",
			Usage:      "Text (maximum 3072 characters)",
			InnerField: "text",
		},
	},
})

func handleMessagesRcsGenerateDeeplink(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageRcGenerateDeeplinkParams{}

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
	_, err = client.Messages.Rcs.GenerateDeeplink(
		ctx,
		cmd.Value("agent-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages:rcs generate-deeplink", obj, format, transform)
}

func handleMessagesRcsSend(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessageRcSendParams{}

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
	_, err = client.Messages.Rcs.Send(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages:rcs send", obj, format, transform)
}

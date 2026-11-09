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

var meetingSessionsActionsSendChat = cli.Command{
	Name:    "send-chat",
	Usage:   "Send chat as the bot in an active meeting. Returns 202 with data.accepted=true.\nGoogle Meet and Teams confirm sending before success; 502 chat_not_sent means\nnot sent, while chat_delivery_unknown means delivery may have occurred. Inspect\nchat and chat.sent events before retrying an uncertain send. Each request is a\nnew message and is not idempotent. Google Meet and Teams allow at most 1000 chat\nrequests per session, including an entrance message; exceeding this returns 409\ninvalid_state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Message text, 1–4000 characters. On Google Meet and Teams, also at most 4096 UTF-8 bytes and nonblank; control characters other than tab and line breaks are rejected. CRLF is normalized to LF.",
			Required: true,
			BodyPath: "text",
		},
	},
	Action:          handleMeetingSessionsActionsSendChat,
	HideHelpCommand: true,
}

var meetingSessionsActionsSpeak = cli.Command{
	Name:    "speak",
	Usage:   "Sends audio / text-to-speech into a meeting session. With a Telnyx AI Assistant\nattached, wait for the assistant to connect before requesting speech. If audio\ndelivery fails, the response can be 502, which may arrive without an error\nenvelope, so branch on the status code before parsing a body. The assistant is\ndesigned to own the conversation, so prefer letting it speak or use `send_chat`.\nWith a digital avatar (`assistant.avatar: true`), speak is refused with 409\n`invalid_state` (message prefix `avatar_requires_assistant_speech:`): the avatar\nrenders only the assistant's own speech.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Text for the bot to speak.",
			Required: true,
			BodyPath: "text",
		},
		&requestflag.Flag[bool]{
			Name:     "interrupt",
			Usage:    "If true, interrupt any currently playing audio to speak this text immediately.",
			BodyPath: "interrupt",
		},
		&requestflag.Flag[string]{
			Name:     "voice",
			Usage:    "Voice identifier to use for this utterance. When supplied, it overrides the session-default voice configured at creation; otherwise the speak action uses that session default.",
			BodyPath: "voice",
		},
	},
	Action:          handleMeetingSessionsActionsSpeak,
	HideHelpCommand: true,
}

var meetingSessionsActionsStopSpeaking = cli.Command{
	Name:    "stop-speaking",
	Usage:   "Stops any active text-to-speech playback in a meeting session. Sessions with a\ndigital avatar (`assistant.avatar: true`) return 409 `invalid_state` (message\nprefix `avatar_requires_assistant_speech:`): the avatar renders only the\nassistant's own speech.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMeetingSessionsActionsStopSpeaking,
	HideHelpCommand: true,
}

func handleMeetingSessionsActionsSendChat(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.MeetingSessionActionSendChatParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.Actions.SendChat(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "meeting-sessions:actions send-chat",
		Transform:      transform,
	})
}

func handleMeetingSessionsActionsSpeak(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.MeetingSessionActionSpeakParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.Actions.Speak(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "meeting-sessions:actions speak",
		Transform:      transform,
	})
}

func handleMeetingSessionsActionsStopSpeaking(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MeetingSessions.Actions.StopSpeaking(ctx, cmd.Value("id").(string), options...)
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
		Title:          "meeting-sessions:actions stop-speaking",
		Transform:      transform,
	})
}

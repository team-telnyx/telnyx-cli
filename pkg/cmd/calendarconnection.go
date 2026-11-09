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

var calendarConnectionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Start Google Calendar or Microsoft Outlook OAuth. Open oauth_url in a browser.\nNo connection exists until authorization succeeds. The resulting connection is\nowned by the connecting user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "platform",
			Usage:    "Calendar provider: Google Calendar or Microsoft Outlook.",
			Required: true,
			BodyPath: "platform",
		},
	},
	Action:          handleCalendarConnectionsCreate,
	HideHelpCommand: true,
}

var calendarConnectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a calendar connection owned by the authenticated connecting user.\nMissing and foreign connections both return 404. Signing secrets are omitted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleCalendarConnectionsRetrieve,
	HideHelpCommand: true,
}

var calendarConnectionsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update policy or defaults for future sessions. Supply at least one field.\nExisting sessions retain their configuration. Null clears nullable overrides.\nAssistant, recording, and webhook defaults apply to Google Meet and Teams.\nAssistant defaults require default_barge_in to be false or null and\ndefault_voice to be null. Recording=true requires a destination; false rejects\none. A non-null default_webhook_url update returns webhook_secret once; save it\nto verify future session deliveries.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "auto-join-policy",
			Usage:    "all_meetings joins eligible events; external_organizer_excluded requires the connected user to be the organizer; per_meeting requires an explicit opt-in.",
			BodyPath: "auto_join_policy",
		},
		&requestflag.Flag[*string]{
			Name:     "bot-name",
			Usage:    "Bot display name; null clears the override.",
			BodyPath: "bot_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-assistant",
			Usage:    "Default assistant for future calendar-scheduled Google Meet and Teams sessions. The saved configuration must fit within 8 KiB. Use id, audio_gate, dynamic_variables, and leave_on_end; avatar is not accepted. audio_gate defaults to full_duplex.",
			BodyPath: "default_assistant",
		},
		&requestflag.Flag[*bool]{
			Name:     "default-barge-in",
			Usage:    "Default manual speech interruption. Cannot be true with default_assistant.",
			BodyPath: "default_barge_in",
		},
		&requestflag.Flag[*bool]{
			Name:     "default-recording",
			Usage:    "Google Meet and Teams recording intent. When null, recording is enabled only if a destination is set. True requires default_recording_export; false cannot be combined with a destination.",
			BodyPath: "default_recording",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-recording-export",
			Usage:    "Customer-owned Telnyx Storage destination for recording exports. Uses the authenticated organization and user; do not supply credentials, region, organization ID or object key.",
			BodyPath: "default_recording_export",
		},
		&requestflag.Flag[*bool]{
			Name:     "default-summarize-on-end",
			Usage:    "Generate a summary when each future session transcript finalizes.",
			BodyPath: "default_summarize_on_end",
		},
		&requestflag.Flag[*string]{
			Name:     "default-voice",
			Usage:    "Manual-speech voice; must be null when default_assistant is set.",
			BodyPath: "default_voice",
		},
		&requestflag.Flag[*string]{
			Name:     "default-webhook-url",
			Usage:    "HTTPS destination for future session webhooks. Null clears the default. A non-null update returns webhook_secret once; save it to verify signatures for sessions created with this default. Existing sessions retain their original signing secret.",
			BodyPath: "default_webhook_url",
		},
	},
	Action:          handleCalendarConnectionsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"default-assistant": {
		&requestflag.InnerFlag[string]{
			Name:       "default-assistant.id",
			Usage:      "Identifier of the assistant to attach.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default-assistant.audio-gate",
			Usage:      "Use full_duplex (default) to allow participants to interrupt, or half_duplex to pause listening while the assistant speaks. Full-duplex pricing depends on participant audio usage; see the Meeting pricing guide.",
			InnerField: "audio_gate",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "default-assistant.dynamic-variables",
			Usage:      "Per-conversation values for the [dynamic variables](/docs/inference/ai-assistants/dynamic-variables) used in the Assistant's instructions, greeting, or tools. Delivered before the Assistant's first utterance, so they resolve for the opening line as well as the rest of the conversation. At most 63 entries; keys 1-128 characters; values must be strings. The map is budgeted in aggregate at 1,047,552 bytes (1023 KiB) rather than capped per value. `streaming_audio`, `ai_assistant_streaming_audio` and `meeting_session_id` are reserved and rejected with `400 invalid_request`.",
			InnerField: "dynamic_variables",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "default-assistant.leave-on-end",
			Usage:      "Leave when the assistant conversation ends or fails. Defaults to false. Other leave-policy rules and explicit leave requests still apply.",
			InnerField: "leave_on_end",
		},
	},
	"default-recording-export": {
		&requestflag.InnerFlag[string]{
			Name:       "default-recording-export.bucket",
			Usage:      "Existing accessible bucket. Consecutive dots and IPv4-address names are rejected.",
			InnerField: "bucket",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default-recording-export.prefix",
			Usage:      "Optional object-key prefix. One trailing slash is normalized. Empty path segments, . and .. segments are rejected.",
			InnerField: "prefix",
		},
	},
})

var calendarConnectionsList = cli.Command{
	Name:            "list",
	Usage:           "List only the authenticated connecting user’s calendar connections. Session\nresources created by those connections are shared within the organization.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleCalendarConnectionsList,
	HideHelpCommand: true,
}

var calendarConnectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Disconnect the calendar, cancel future joins, and clear retained calendar\nidentity fields. Previously created session records, transcripts, artifacts, and\nStorage objects remain. The response is the updated connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleCalendarConnectionsDelete,
	HideHelpCommand: true,
}

func handleCalendarConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.CalendarConnectionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CalendarConnections.New(ctx, params, options...)
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
		Title:          "calendar-connections create",
		Transform:      transform,
	})
}

func handleCalendarConnectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.CalendarConnections.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "calendar-connections retrieve",
		Transform:      transform,
	})
}

func handleCalendarConnectionsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.CalendarConnectionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CalendarConnections.Update(
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
		Title:          "calendar-connections update",
		Transform:      transform,
	})
}

func handleCalendarConnectionsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CalendarConnections.List(ctx, options...)
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
		Title:          "calendar-connections list",
		Transform:      transform,
	})
}

func handleCalendarConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.CalendarConnections.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "calendar-connections delete",
		Transform:      transform,
	})
}

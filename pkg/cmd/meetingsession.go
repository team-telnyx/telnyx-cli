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

var meetingSessionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new meeting session. When an idempotency_key is supplied in the\nrequest body, replay lookup is scoped to the authenticated account and compares\nonly the key; the request payload is not fingerprinted or compared. If a session\nwith that key already exists for the account, the existing session is replayed\n(200); otherwise a new session is created (201). Supports bring-your-own-key\n(BYOK) configuration. The session may enter asynchronous states (e.g. joining,\nwaiting_for_admission) before becoming active. Optional `camera_image` input is\nwrite-only and applies only when no Avatar or Assistant webpage output takes\nprecedence. An ignored URL is not fetched. An effective URL source is resolved\nbefore bot creation; neither the source URL nor image bytes are persisted,\nreturned, or logged. Treat signed URLs as credentials.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "meeting-url",
			Usage:    "The meeting URL the bot should join.",
			Required: true,
			BodyPath: "meeting_url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "assistant",
			Usage:    "Request options for attaching a voice assistant to the session. Routing fields (`call_control_connection_id`, `from`, and `loopback_sip_uri`) are used only to establish the assistant call leg and are omitted from response objects. `audio_gate` is returned with `id` in the assistant response object.",
			BodyPath: "assistant",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "avatar",
			Usage:    "Request options for attaching a bring-your-own-key avatar to the session.",
			BodyPath: "avatar",
		},
		&requestflag.Flag[bool]{
			Name:     "barge-in",
			Usage:    "When enabled, a human participant `speech_on` event interrupts and stops the current bot audio; it does not bypass admission or initiate speech. Assistant sessions reject `barge_in: true`.",
			Default:  false,
			BodyPath: "barge_in",
		},
		&requestflag.Flag[string]{
			Name:     "bot-name",
			Usage:    `Display name for the bot in the meeting. Defaults to "Meeting Bot".`,
			BodyPath: "bot_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "camera-image",
			Usage:    "Write-only static camera-tile image for this session, not a native account or participant profile photo. Supply exactly one JPEG source. When effective, the image is used as the bot's static camera/video output; presentation varies by meeting platform and recording configuration and is not guaranteed in recordings. An effective Avatar or Assistant webpage output takes precedence, so this input is ignored and a URL source is not fetched.",
			BodyPath: "camera_image",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Client-supplied idempotency key to safely retry creation requests without duplicating sessions. Lookup is scoped to the authenticated account and compares the key only; the request payload is not fingerprinted or compared.",
			BodyPath: "idempotency_key",
		},
		&requestflag.Flag[any]{
			Name:     "join-at",
			Usage:    "ISO-8601 timestamp in the future at which the bot should join. If omitted, the bot joins immediately.",
			BodyPath: "join_at",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Arbitrary key-value metadata attached to the session. The serialized JSON representation must not exceed 16384 characters at runtime.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "speak-on-enter",
			Usage:    "Text the bot speaks when it enters the meeting.",
			BodyPath: "speak_on_enter",
		},
		&requestflag.Flag[bool]{
			Name:     "summarize-on-end",
			Usage:    "If true, generate a summary artifact when the session ends.",
			Default:  false,
			BodyPath: "summarize_on_end",
		},
		&requestflag.Flag[string]{
			Name:     "voice",
			Usage:    "Session-default voice identifier used for `speak_on_enter` and ordinary speak actions. A voice supplied on an individual speak action overrides this default for that utterance.",
			BodyPath: "voice",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "HTTPS endpoint to receive session lifecycle callbacks. Static validation requires HTTPS, rejects embedded credentials and blocked hosts, and enforces egress policy. Validation makes no network request to the endpoint.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleMeetingSessionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"assistant": {
		&requestflag.InnerFlag[string]{
			Name:       "assistant.id",
			Usage:      "Identifier of the assistant to attach.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assistant.call-control-connection-id",
			Usage:      "Call control connection used to bridge the assistant into the meeting audio.",
			InnerField: "call_control_connection_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assistant.from",
			Usage:      "E.164 calling number used as the originating party for the assistant call leg.",
			InnerField: "from",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assistant.loopback-sip-uri",
			Usage:      "SIP URI to which the assistant media loopback is established.",
			InnerField: "loopback_sip_uri",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assistant.audio-gate",
			Usage:      "Audio gating strategy for the assistant call leg.",
			InnerField: "audio_gate",
		},
	},
	"avatar": {
		&requestflag.InnerFlag[string]{
			Name:       "avatar.api-key",
			Usage:      "Bring-your-own-key API key for the avatar provider. The key is never stored or returned by the API.",
			InnerField: "api_key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "avatar.avatar-id",
			Usage:      "Identifier of the avatar to use.",
			InnerField: "avatar_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "avatar.provider",
			Usage:      `Avatar provider identifier. Currently only "anam" is supported.`,
			InnerField: "provider",
		},
	},
})

var meetingSessionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single meeting session by ID. A session that does not exist or that\nbelongs to a different account both return 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMeetingSessionsRetrieve,
	HideHelpCommand: true,
}

var meetingSessionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates mutable properties of a meeting session. Only sessions in the scheduled\nstate can be updated; any other state returns 409 with the invalid_state error\ncode. All request fields are optional, and an empty object is a valid no-op\nupdate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "bot-name",
			Usage:    "Updated display name for the bot.",
			BodyPath: "bot_name",
		},
		&requestflag.Flag[any]{
			Name:     "join-at",
			Usage:    "ISO-8601 timestamp for the bot to join. May be updated to reschedule.",
			BodyPath: "join_at",
		},
	},
	Action:          handleMeetingSessionsUpdate,
	HideHelpCommand: true,
}

var meetingSessionsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a list of meeting sessions, optionally filtered by status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter meeting sessions by current status.",
			QueryPath: "status",
		},
	},
	Action:          handleMeetingSessionsList,
	HideHelpCommand: true,
}

var meetingSessionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Stops a meeting session without deleting its persisted record. Scheduled bots\nare cancelled, while bots that are joining or active are asked to leave. The\npersisted meeting session record remains available.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMeetingSessionsDelete,
	HideHelpCommand: true,
}

var meetingSessionsDeleteRecordingMedia = cli.Command{
	Name:    "delete-recording-media",
	Usage:   "**Not yet available in production** — this route is not currently routed on\napi.telnyx.com and returns a generic 404; it is documented ahead of rollout.\nIrreversibly requests deletion of provider-hosted aggregate recording media\nunder the provider contract. The operation retains the Telnyx-local Meeting\nsession, transcript segments, events, artifacts, and usage records. It is\nseparate from `DELETE /meeting_sessions/{id}`, which stops or cancels\nparticipation without deleting the persisted session. A missing/foreign session\nreturns 404; provider deletion failures return 502.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMeetingSessionsDeleteRecordingMedia,
	HideHelpCommand: true,
}

var meetingSessionsRetrieveEvents = cli.Command{
	Name:    "retrieve-events",
	Usage:   "Returns stored events ordered by ascending `seq`. To continue, pass the last\nreturned item's `seq` as `after`. An empty page means no later stored events\nexisted at read time; this operation returns no separate next-page cursor.\nDefault `limit` is 100 and maximum is 1,000.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[int64]{
			Name:      "after",
			Usage:     "Return results with a cursor position after this value.",
			Default:   0,
			QueryPath: "after",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return per page.",
			Default:   100,
			QueryPath: "limit",
		},
	},
	Action:          handleMeetingSessionsRetrieveEvents,
	HideHelpCommand: true,
}

var meetingSessionsRetrieveRecordings = cli.Command{
	Name:    "retrieve-recordings",
	Usage:   "Returns recordings for a meeting session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMeetingSessionsRetrieveRecordings,
	HideHelpCommand: true,
}

var meetingSessionsRetrieveTranscript = cli.Command{
	Name:    "retrieve-transcript",
	Usage:   "Returns transcript segments ordered by ascending `seq`. Default `limit` is 100\nand maximum is 1,000. Continue with `after=meta.next_after`. A long-poll timeout\nreturns 200 with empty `data` and `meta.next_after: null`; retain the cursor\nsupplied to that request because null is not a replacement cursor.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[int64]{
			Name:      "after",
			Usage:     "Return results with a cursor position after this value.",
			Default:   0,
			QueryPath: "after",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return per page.",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "wait-seconds",
			Usage:     "Long-poll duration in seconds. The server holds the connection open for up to this many seconds, waiting for new or updated results before returning an empty response. Set to 0 for an immediate response.",
			Default:   0,
			QueryPath: "wait_seconds",
		},
	},
	Action:          handleMeetingSessionsRetrieveTranscript,
	HideHelpCommand: true,
}

func handleMeetingSessionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.MeetingSessionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.New(ctx, params, options...)
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
		Title:          "meeting-sessions create",
		Transform:      transform,
	})
}

func handleMeetingSessionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MeetingSessions.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "meeting-sessions retrieve",
		Transform:      transform,
	})
}

func handleMeetingSessionsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.MeetingSessionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.Update(
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
		Title:          "meeting-sessions update",
		Transform:      transform,
	})
}

func handleMeetingSessionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.MeetingSessionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.List(ctx, params, options...)
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
		Title:          "meeting-sessions list",
		Transform:      transform,
	})
}

func handleMeetingSessionsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MeetingSessions.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "meeting-sessions delete",
		Transform:      transform,
	})
}

func handleMeetingSessionsDeleteRecordingMedia(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MeetingSessions.DeleteRecordingMedia(ctx, cmd.Value("id").(string), options...)
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
		Title:          "meeting-sessions delete-recording-media",
		Transform:      transform,
	})
}

func handleMeetingSessionsRetrieveEvents(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.MeetingSessionGetEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.GetEvents(
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
		Title:          "meeting-sessions retrieve-events",
		Transform:      transform,
	})
}

func handleMeetingSessionsRetrieveRecordings(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MeetingSessions.GetRecordings(ctx, cmd.Value("id").(string), options...)
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
		Title:          "meeting-sessions retrieve-recordings",
		Transform:      transform,
	})
}

func handleMeetingSessionsRetrieveTranscript(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.MeetingSessionGetTranscriptParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.GetTranscript(
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
		Title:          "meeting-sessions retrieve-transcript",
		Transform:      transform,
	})
}

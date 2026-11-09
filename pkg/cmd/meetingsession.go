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
	Usage:   "Creates a meeting session and returns 201, or 200 for a replay. Use a unique\n`idempotency_key` per intended join and reuse the same key and request on\nretries. Reusing a key with a different join configuration or meeting URL can\nreturn 409 invalid_state. Always retry with the original request. The session\ncan pass through joining and waiting_for_admission before active. Optional\n`camera_image` input is write-only and applies when no assistant avatar takes\nprecedence. An ignored URL is not fetched. An effective URL source is resolved\nbefore bot creation; the source URL and image bytes are not returned in the\nsession response. Treat signed URLs as credentials. With\n`assistant.avatar: true`, the assistant is read before the bot is created: an\nassistant without a Telnyx avatar configured is rejected with 400\n`invalid_request` (message prefix `assistant_avatar_not_configured:`), an\nunknown assistant with 400 `invalid_request` (prefix `assistant_not_found:`),\nand a failed lookup with 503 `assistant_lookup_unavailable`; no bot is created\nin any of these cases.",
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
			Usage:    "Attach a Telnyx AI Assistant owned by the authenticated account. Supply its ID and optional conversation settings. Unknown fields are rejected.",
			BodyPath: "assistant",
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
			Usage:    "Not supported for Google Meet and Teams joins (422 unsupported_capability). Write-only static camera-tile image for this session, not an account or participant profile photo. Supply exactly one JPEG source. When effective, the image is used as the bot's static camera/video output; presentation varies by meeting platform and recording configuration and is not guaranteed in recordings. An attached assistant or avatar takes precedence, so this input is ignored and a URL source is not fetched.",
			BodyPath: "camera_image",
		},
		&requestflag.Flag[string]{
			Name:     "chat-on-enter",
			Usage:    "A message the bot posts to the meeting's chat as soon as it becomes active — typically a recording disclosure. Delivered at most once. Independent of `speak_on_enter`: both may be set, and the chat message posts first because it does not wait for text-to-speech or avatar startup. Rejected with 422 `unsupported_capability` on platforms without meeting chat. For Google Meet and Teams, text must be nonblank, at most 4096 UTF-8 bytes, with no control characters except tabs and line breaks. Counts toward the 1000-chat-request session limit.",
			BodyPath: "chat_on_enter",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Use a unique key per intended join and reuse the same key and request on retries. Reusing a key with a different join configuration or meeting URL can return 409 invalid_state. Always retry with the original request.",
			BodyPath: "idempotency_key",
		},
		&requestflag.Flag[any]{
			Name:     "join-at",
			Usage:    "Future ISO-8601 join timestamp including timezone. Google Meet and Teams accept dates at most three years ahead. Omit for an immediate join. Assistant sessions created through this endpoint must join immediately.",
			BodyPath: "join_at",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "leave-policy",
			Usage:    "Automatic departure settings for Google Meet and Teams. Omitted fields use the default behavior. Set at creation; cannot be changed later. Zoom ignores this policy.",
			BodyPath: "leave_policy",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Arbitrary key-value metadata attached to the session. The serialized JSON representation must not exceed 16384 characters at runtime.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[bool]{
			Name:     "recording",
			Usage:    "Google Meet and Teams only. Defaults to true when recording_export is supplied and false otherwise. True requires recording_export.bucket; false cannot be combined with recording_export. Fixed at creation, with no pause/resume action. Zoom rejects recording, recording_view, and recording_export with 422 unsupported_capability; omit these fields for Zoom.",
			BodyPath: "recording",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "recording-export",
			Usage:    "Customer-owned Telnyx Storage destination for recording exports. Uses the authenticated organization and user; do not supply credentials, region, organization ID or object key.",
			BodyPath: "recording_export",
		},
		&requestflag.Flag[string]{
			Name:     "recording-view",
			Usage:    "Recording layout for Google Meet and Teams. Defaults to speaker_view. Accepted at creation only; does not enable recording by itself.",
			Default:  "speaker_view",
			BodyPath: "recording_view",
		},
		&requestflag.Flag[string]{
			Name:     "speak-on-enter",
			Usage:    "Text the bot speaks after joining. Available for sessions without a digital avatar. With an AI Assistant, the greeting waits until the assistant is connected. Rejected with 400 invalid_request when assistant.avatar is true.",
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
			Usage:    "Session-default voice identifier used for `speak_on_enter` and ordinary speak actions. A voice supplied on an individual speak action overrides this default for that utterance. With an assistant attached, its conversational voice remains configured on the assistant.",
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
			Name:       "assistant.audio-gate",
			Usage:      "Use half_duplex (default) to pause listening while the assistant speaks, or full_duplex to allow participants to interrupt. Full-duplex pricing depends on participant audio usage; see the Meeting pricing guide.",
			InnerField: "audio_gate",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "assistant.avatar",
			Usage:      "Beta. Show the avatar configured on the assistant (its `avatar_settings`; see [AI Assistant avatars](/docs/inference/ai-assistants/avatars)) as the bot's camera, lip-synced to the assistant's own speech. The request carries no avatar ID or key: the assistant is read at creation and the request is rejected with 400 `invalid_request` if it has no Telnyx avatar configured. Available for Google Meet, Zoom, and Teams sessions. REST-only. Not echoed back on `assistant`; the session's top-level `avatar` and `avatar_state` fields report it instead. If the avatar fails, the session continues voice-only and `avatar_state` becomes `degraded`.",
			InnerField: "avatar",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "assistant.dynamic-variables",
			Usage:      "Per-conversation values for the [dynamic variables](/docs/inference/ai-assistants/dynamic-variables) used in the Assistant's instructions, greeting, or tools. Delivered before the Assistant's first utterance, so they resolve for the opening line as well as the rest of the conversation. At most 63 entries; keys 1-128 characters; values must be strings. The map is budgeted in aggregate at 1,047,552 bytes (1023 KiB) rather than capped per value. `streaming_audio`, `ai_assistant_streaming_audio` and `meeting_session_id` are reserved and rejected with `400 invalid_request`.",
			InnerField: "dynamic_variables",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "assistant.leave-on-end",
			Usage:      "Leave when the assistant conversation ends or fails. Defaults to false. Other leave-policy rules and explicit leave requests still apply.",
			InnerField: "leave_on_end",
		},
	},
	"leave-policy": {
		&requestflag.InnerFlag[int64]{
			Name:       "leave-policy.alone-seconds",
			Usage:      "Seconds alone after another participant has attended and left. Does not start just because the bot arrives first.",
			InnerField: "alone_seconds",
		},
		&requestflag.InnerFlag[any]{
			Name:       "leave-policy.bot-name-keywords",
			Usage:      "Case-insensitive names or phrases identifying other bots to exclude when deciding whether this bot is alone. Matching splits names on spaces, hyphens and underscores. Null or an empty array disables filtering.",
			InnerField: "bot_name_keywords",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "leave-policy.host-wait-seconds",
			Usage:      "Seconds to wait for the host to start the meeting.",
			InnerField: "host_wait_seconds",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "leave-policy.lobby-wait-seconds",
			Usage:      "Seconds to wait for lobby admission.",
			InnerField: "lobby_wait_seconds",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "leave-policy.max-duration-seconds",
			Usage:      "Maximum seconds after joining. Null means no per-join duration limit; service safety limits still apply.",
			InnerField: "max_duration_seconds",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "leave-policy.silence-grace-seconds",
			Usage:      "Seconds after joining before silence detection starts a fresh inactivity timer.",
			InnerField: "silence_grace_seconds",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "leave-policy.silence-seconds",
			Usage:      "Continuous audio inactivity before departure after the silence grace period.",
			InnerField: "silence_seconds",
		},
	},
	"recording-export": {
		&requestflag.InnerFlag[string]{
			Name:       "recording-export.bucket",
			Usage:      "Existing accessible bucket. Consecutive dots and IPv4-address names are rejected.",
			InnerField: "bucket",
		},
		&requestflag.InnerFlag[string]{
			Name:       "recording-export.prefix",
			Usage:      "Optional object-key prefix. One trailing slash is normalized. Empty path segments, . and .. segments are rejected.",
			InnerField: "prefix",
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
	Usage:   "Updates mutable properties of a meeting session. Only sessions in the scheduled\nstate can be updated; any other state returns 409 with the invalid_state error\ncode. All request fields are optional, and an empty object is a valid no-op\nupdate. Calendar-controlled Google Meet and Teams sessions return 409\ncalendar_managed; change the event instead. A conflicting pending edit returns\n409 edit_in_progress. On a timeout or 503, read the session and retry the same\nchanges.",
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
	Usage:   "Returns a list of meeting sessions, optionally filtered by status. List items\nomit assistant.dynamic_variables and recording_export. Some items include\nremote_state_observed_at; retrieve the individual session for current state.",
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
	Usage:   "Stops a meeting session without deleting its persisted record. Scheduled bots\nare cancelled, while bots that are joining or active are asked to leave. The\npersisted meeting session record remains available. A successful response can\nreport leaving while departure is pending. Poll until ended, failed, or\nadmission_denied. Exported Storage objects are unaffected.",
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
	Usage:   "Irreversibly requests deletion of the aggregate recording media for the session.\nThe operation retains the Telnyx-local Meeting session, transcript segments,\nevents, artifacts, and usage records. It is separate from\n`DELETE /meeting_sessions/{id}`, which stops or cancels participation without\ndeleting the persisted session. A missing/foreign session returns 404; a\ndeletion request whose outcome is unknown returns 502. Customer Storage exports\ndo not support this operation (422 unsupported_capability); delete the exported\nobject through Storage instead. Stopping a session does not delete an exported\nobject.",
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
	Usage:   "Retrieve recording export status. Customer Storage exports include\nrecording_export with the bucket, region, key and pending/exported/failed\nstatus, or null when no reference is available. The data array is empty for\nthese exports. Download or share the object using the Storage API once its\nstatus is exported. Zoom recordings and existing Google Meet or Teams recordings\nheld by the Meeting service continue returning temporary download links in data\nthrough this same endpoint and session ID. Those recordings do not require\nStorage setup or migration. Refresh this endpoint when a link expires. Select\nthe download method from the returned fields, not the platform alone.",
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
	Usage:   "Returns transcript segments ordered by ascending seq. Default limit is 100,\nmaximum 1000. Continue with after=meta.next_after. On an empty page,\nmeta.next_after preserves a positive supplied after cursor, or is null when no\npositive cursor was supplied. wait_seconds holds an empty read for up to 25\nseconds; a timeout still returns 200.",
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

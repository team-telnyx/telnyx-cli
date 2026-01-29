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

var callsDial = requestflag.WithInnerFlags(cli.Command{
	Name:    "dial",
	Usage:   "Dial a number or SIP URI from a given connection. A successful response will\ninclude a `call_leg_id` which can be used to correlate the command with\nsubsequent webhooks.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Usage:    "The ID of the Call Control App (formerly ID of the connection) to be used when dialing the destination.",
			Required: true,
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "The `from` number to be used as the caller id presented to the destination (`to` number). The number should be in +E164 format.",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[any]{
			Name:     "to",
			Usage:    "The DID or SIP URI to dial out to. Multiple DID or SIP URIs can be provided using an array of strings",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[string]{
			Name:     "answering-machine-detection",
			Usage:    "Enables Answering Machine Detection. Telnyx offers Premium and Standard detections. With Premium detection, when a call is answered, Telnyx runs real-time detection and sends a `call.machine.premium.detection.ended` webhook with one of the following results: `human_residence`, `human_business`, `machine`, `silence` or `fax_detected`. If we detect a beep, we also send a `call.machine.premium.greeting.ended` webhook with the result of `beep_detected`. If we detect a beep before `call.machine.premium.detection.ended` we only send `call.machine.premium.greeting.ended`, and if we detect a beep after `call.machine.premium.detection.ended`, we send both webhooks. With Standard detection, when a call is answered, Telnyx runs real-time detection to determine if it was picked up by a human or a machine and sends an `call.machine.detection.ended` webhook with the analysis result. If `greeting_end` or `detect_words` is used and a `machine` is detected, you will receive another `call.machine.greeting.ended` webhook when the answering machine greeting ends with a beep or silence. If `detect_beep` is used, you will only receive `call.machine.greeting.ended` if a beep is detected.",
			Default:  "disabled",
			BodyPath: "answering_machine_detection",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "answering-machine-detection-config",
			Usage:    "Optional configuration parameters to modify 'answering_machine_detection' performance.",
			BodyPath: "answering_machine_detection_config",
		},
		&requestflag.Flag[string]{
			Name:     "audio-url",
			Usage:    "The URL of a file to be played back to the callee when the call is answered. The URL can point to either a WAV or MP3 file. media_name and audio_url cannot be used together in one request.",
			BodyPath: "audio_url",
		},
		&requestflag.Flag[string]{
			Name:     "billing-group-id",
			Usage:    "Use this field to set the Billing Group ID for the call. Must be a valid and existing Billing Group ID.",
			BodyPath: "billing_group_id",
		},
		&requestflag.Flag[bool]{
			Name:     "bridge-intent",
			Usage:    "Indicates the intent to bridge this call with the call specified in link_to. When bridge_intent is true, link_to becomes required and the from number will be overwritten by the from number from the linked call.",
			Default:  false,
			BodyPath: "bridge_intent",
		},
		&requestflag.Flag[bool]{
			Name:     "bridge-on-answer",
			Usage:    "Whether to automatically bridge answered call to the call specified in link_to. When bridge_on_answer is true, link_to becomes required.",
			Default:  false,
			BodyPath: "bridge_on_answer",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore others Dial commands with the same `command_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "conference-config",
			Usage:    "Optional configuration parameters to dial new participant into a conference.",
			BodyPath: "conference_config",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-header",
			Usage:    "Custom headers to be added to the SIP INVITE.",
			BodyPath: "custom_headers",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dialogflow-config",
			BodyPath: "dialogflow_config",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-dialogflow",
			Usage:    "Enables Dialogflow for the current call. The default value is false.",
			Default:  false,
			BodyPath: "enable_dialogflow",
		},
		&requestflag.Flag[string]{
			Name:     "from-display-name",
			Usage:    "The `from_display_name` string to be used as the caller id name (SIP From Display Name) presented to the destination (`to` number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and -_~!.+ special characters. If ommited, the display name will be the same as the number in the `from` field.",
			BodyPath: "from_display_name",
		},
		&requestflag.Flag[string]{
			Name:     "link-to",
			Usage:    "Use another call's control id for sharing the same call session id",
			BodyPath: "link_to",
		},
		&requestflag.Flag[string]{
			Name:     "media-encryption",
			Usage:    "Defines whether media should be encrypted on the call.",
			Default:  "disabled",
			BodyPath: "media_encryption",
		},
		&requestflag.Flag[string]{
			Name:     "media-name",
			Usage:    "The media_name of a file to be played back to the callee when the call is answered. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file.",
			BodyPath: "media_name",
		},
		&requestflag.Flag[string]{
			Name:     "park-after-unbridge",
			Usage:    "If supplied with the value `self`, the current leg will be parked after unbridge. If not set, the default behavior is to hang up the leg. When park_after_unbridge is set, link_to becomes required.",
			BodyPath: "park_after_unbridge",
		},
		&requestflag.Flag[string]{
			Name:     "preferred-codecs",
			Usage:    "The list of comma-separated codecs in a preferred order for the forked media to be received.",
			BodyPath: "preferred_codecs",
		},
		&requestflag.Flag[string]{
			Name:     "record",
			Usage:    "Start recording automatically after an event. Disabled by default.",
			BodyPath: "record",
		},
		&requestflag.Flag[string]{
			Name:     "record-channels",
			Usage:    "Defines which channel should be recorded ('single' or 'dual') when `record` is specified.",
			Default:  "dual",
			BodyPath: "record_channels",
		},
		&requestflag.Flag[string]{
			Name:     "record-custom-file-name",
			Usage:    "The custom recording file name to be used instead of the default `call_leg_id`. Telnyx will still add a Unix timestamp suffix.",
			BodyPath: "record_custom_file_name",
		},
		&requestflag.Flag[string]{
			Name:     "record-format",
			Usage:    "Defines the format of the recording ('wav' or 'mp3') when `record` is specified.",
			Default:  "mp3",
			BodyPath: "record_format",
		},
		&requestflag.Flag[int64]{
			Name:     "record-max-length",
			Usage:    "Defines the maximum length for the recording in seconds when `record` is specified. The minimum value is 0. The maximum value is 43200. The default value is 0 (infinite).",
			Default:  0,
			BodyPath: "record_max_length",
		},
		&requestflag.Flag[int64]{
			Name:     "record-timeout-secs",
			Usage:    "The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected when `record` is specified. The timer only starts when the speech is detected. Please note that call transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite).",
			Default:  0,
			BodyPath: "record_timeout_secs",
		},
		&requestflag.Flag[string]{
			Name:     "record-track",
			Usage:    "The audio track to be recorded. Can be either `both`, `inbound` or `outbound`. If only single track is specified (`inbound`, `outbound`), `channels` configuration is ignored and it will be recorded as mono (single channel).",
			Default:  "both",
			BodyPath: "record_track",
		},
		&requestflag.Flag[string]{
			Name:     "record-trim",
			Usage:    "When set to `trim-silence`, silence will be removed from the beginning and end of the recording.",
			BodyPath: "record_trim",
		},
		&requestflag.Flag[bool]{
			Name:     "send-silence-when-idle",
			Usage:    "Generate silence RTP packets when no transmission available.",
			Default:  false,
			BodyPath: "send_silence_when_idle",
		},
		&requestflag.Flag[string]{
			Name:     "sip-auth-password",
			Usage:    "SIP Authentication password used for SIP challenges.",
			BodyPath: "sip_auth_password",
		},
		&requestflag.Flag[string]{
			Name:     "sip-auth-username",
			Usage:    "SIP Authentication username used for SIP challenges.",
			BodyPath: "sip_auth_username",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "sip-header",
			Usage:    "SIP headers to be added to the SIP INVITE request. Currently only User-to-User header is supported.",
			BodyPath: "sip_headers",
		},
		&requestflag.Flag[string]{
			Name:     "sip-region",
			Usage:    "Defines the SIP region to be used for the call.",
			Default:  "US",
			BodyPath: "sip_region",
		},
		&requestflag.Flag[string]{
			Name:     "sip-transport-protocol",
			Usage:    "Defines SIP transport protocol to be used on the call.",
			Default:  "UDP",
			BodyPath: "sip_transport_protocol",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sound-modifications",
			Usage:    "Use this field to modify sound effects, for example adjust the pitch.",
			BodyPath: "sound_modifications",
		},
		&requestflag.Flag[string]{
			Name:     "stream-bidirectional-codec",
			Usage:    "Indicates codec for bidirectional streaming RTP payloads. Used only with stream_bidirectional_mode=rtp. Case sensitive.",
			Default:  "PCMU",
			BodyPath: "stream_bidirectional_codec",
		},
		&requestflag.Flag[string]{
			Name:     "stream-bidirectional-mode",
			Usage:    "Configures method of bidirectional streaming (mp3, rtp).",
			Default:  "mp3",
			BodyPath: "stream_bidirectional_mode",
		},
		&requestflag.Flag[int64]{
			Name:     "stream-bidirectional-sampling-rate",
			Usage:    "Audio sampling rate.",
			Default:  8000,
			BodyPath: "stream_bidirectional_sampling_rate",
		},
		&requestflag.Flag[string]{
			Name:     "stream-bidirectional-target-legs",
			Usage:    "Specifies which call legs should receive the bidirectional stream audio.",
			Default:  "opposite",
			BodyPath: "stream_bidirectional_target_legs",
		},
		&requestflag.Flag[string]{
			Name:     "stream-codec",
			Usage:    "Specifies the codec to be used for the streamed audio. When set to 'default' or when transcoding is not possible, the codec from the call will be used.",
			Default:  "default",
			BodyPath: "stream_codec",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-establish-before-call-originate",
			Usage:    "Establish websocket connection before dialing the destination. This is useful for cases where the websocket connection takes a long time to establish.",
			Default:  false,
			BodyPath: "stream_establish_before_call_originate",
		},
		&requestflag.Flag[string]{
			Name:     "stream-track",
			Usage:    "Specifies which track should be streamed.",
			Default:  "inbound_track",
			BodyPath: "stream_track",
		},
		&requestflag.Flag[string]{
			Name:     "stream-url",
			Usage:    "The destination WebSocket address where the stream is going to be delivered.",
			BodyPath: "stream_url",
		},
		&requestflag.Flag[string]{
			Name:     "supervise-call-control-id",
			Usage:    "The call leg which will be supervised by the new call.",
			BodyPath: "supervise_call_control_id",
		},
		&requestflag.Flag[string]{
			Name:     "supervisor-role",
			Usage:    "The role of the supervisor call. 'barge' means that supervisor call hears and is being heard by both ends of the call (caller & callee). 'whisper' means that only supervised_call_control_id hears supervisor but supervisor can hear everything. 'monitor' means that nobody can hear supervisor call, but supervisor can hear everything on the call.",
			Default:  "barge",
			BodyPath: "supervisor_role",
		},
		&requestflag.Flag[int64]{
			Name:     "time-limit-secs",
			Usage:    "Sets the maximum duration of a Call Control Leg in seconds. If the time limit is reached, the call will hangup and a `call.hangup` webhook with a `hangup_cause` of `time_limit` will be sent. For example, by setting a time limit of 120 seconds, a Call Leg will be automatically terminated two minutes after being answered. The default time limit is 14400 seconds or 4 hours and this is also the maximum allowed call length.",
			Default:  14400,
			BodyPath: "time_limit_secs",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-secs",
			Usage:    "The number of seconds that Telnyx will wait for the call to be answered by the destination to which it is being called. If the timeout is reached before an answer is received, the call will hangup and a `call.hangup` webhook with a `hangup_cause` of `timeout` will be sent. Minimum value is 5 seconds. Maximum value is 600 seconds.",
			Default:  30,
			BodyPath: "timeout_secs",
		},
		&requestflag.Flag[bool]{
			Name:     "transcription",
			Usage:    "Enable transcription upon call answer. The default value is false.",
			Default:  false,
			BodyPath: "transcription",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transcription-config",
			BodyPath: "transcription_config",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Use this field to override the URL for which Telnyx will send subsequent webhooks to for this call.",
			BodyPath: "webhook_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url-method",
			Usage:    "HTTP request type used for `webhook_url`.",
			Default:  "POST",
			BodyPath: "webhook_url_method",
		},
	},
	Action:          handleCallsDial,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"answering-machine-detection-config": {
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.after-greeting-silence-millis",
			Usage:      "Silence duration threshold after a greeting message or voice for it be considered human.",
			InnerField: "after_greeting_silence_millis",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.between-words-silence-millis",
			Usage:      "Maximum threshold for silence between words.",
			InnerField: "between_words_silence_millis",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.greeting-duration-millis",
			Usage:      "Maximum threshold of a human greeting. If greeting longer than this value, considered machine.",
			InnerField: "greeting_duration_millis",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.greeting-silence-duration-millis",
			Usage:      "If machine already detected, maximum threshold for silence between words. If exceeded, the greeting is considered ended.",
			InnerField: "greeting_silence_duration_millis",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.greeting-total-analysis-time-millis",
			Usage:      "If machine already detected, maximum timeout threshold to determine the end of the machine greeting.",
			InnerField: "greeting_total_analysis_time_millis",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.initial-silence-millis",
			Usage:      "If initial silence duration is greater than this value, consider it a machine.",
			InnerField: "initial_silence_millis",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.maximum-number-of-words",
			Usage:      "If number of detected words is greater than this value, consder it a machine.",
			InnerField: "maximum_number_of_words",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.maximum-word-length-millis",
			Usage:      "If a single word lasts longer than this threshold, consider it a machine.",
			InnerField: "maximum_word_length_millis",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.silence-threshold",
			Usage:      "Minimum noise threshold for any analysis.",
			InnerField: "silence_threshold",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "answering-machine-detection-config.total-analysis-time-millis",
			Usage:      "Maximum timeout threshold for overall detection.",
			InnerField: "total_analysis_time_millis",
		},
	},
	"conference-config": {
		&requestflag.InnerFlag[string]{
			Name:       "conference-config.id",
			Usage:      "Conference ID to be joined",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "conference-config.beep-enabled",
			Usage:      "Whether a beep sound should be played when the participant joins and/or leaves the conference. Can be used to override the conference-level setting.",
			InnerField: "beep_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "conference-config.conference-name",
			Usage:      "Conference name to be joined",
			InnerField: "conference_name",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "conference-config.early-media",
			Usage:      "Controls the moment when dialled call is joined into conference. If set to `true` user will be joined as soon as media is available (ringback). If `false` user will be joined when call is answered. Defaults to `true`",
			InnerField: "early_media",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "conference-config.end-conference-on-exit",
			Usage:      `Whether the conference should end and all remaining participants be hung up after the participant leaves the conference. Defaults to "false".`,
			InnerField: "end_conference_on_exit",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "conference-config.hold",
			Usage:      `Whether the participant should be put on hold immediately after joining the conference. Defaults to "false".`,
			InnerField: "hold",
		},
		&requestflag.InnerFlag[string]{
			Name:       "conference-config.hold-audio-url",
			Usage:      `The URL of a file to be played to the participant when they are put on hold after joining the conference. hold_media_name and hold_audio_url cannot be used together in one request. Takes effect only when "start_conference_on_create" is set to "false". This property takes effect only if "hold" is set to "true".`,
			InnerField: "hold_audio_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "conference-config.hold-media-name",
			Usage:      `The media_name of a file to be played to the participant when they are put on hold after joining the conference. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. Takes effect only when "start_conference_on_create" is set to "false". This property takes effect only if "hold" is set to "true".`,
			InnerField: "hold_media_name",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "conference-config.mute",
			Usage:      `Whether the participant should be muted immediately after joining the conference. Defaults to "false".`,
			InnerField: "mute",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "conference-config.soft-end-conference-on-exit",
			Usage:      `Whether the conference should end after the participant leaves the conference. NOTE this doesn't hang up the other participants. Defaults to "false".`,
			InnerField: "soft_end_conference_on_exit",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "conference-config.start-conference-on-create",
			Usage:      `Whether the conference should be started on creation. If the conference isn't started all participants that join are automatically put on hold. Defaults to "true".`,
			InnerField: "start_conference_on_create",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "conference-config.start-conference-on-enter",
			Usage:      `Whether the conference should be started after the participant joins the conference. Defaults to "false".`,
			InnerField: "start_conference_on_enter",
		},
		&requestflag.InnerFlag[string]{
			Name:       "conference-config.supervisor-role",
			Usage:      `Sets the joining participant as a supervisor for the conference. A conference can have multiple supervisors. "barge" means the supervisor enters the conference as a normal participant. This is the same as "none". "monitor" means the supervisor is muted but can hear all participants. "whisper" means that only the specified "whisper_call_control_ids" can hear the supervisor. Defaults to "none".`,
			InnerField: "supervisor_role",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "conference-config.whisper-call-control-ids",
			Usage:      "Array of unique call_control_ids the joining supervisor can whisper to. If none provided, the supervisor will join the conference as a monitoring participant only.",
			InnerField: "whisper_call_control_ids",
		},
	},
	"custom-header": {
		&requestflag.InnerFlag[string]{
			Name:       "custom-header.name",
			Usage:      "The name of the header to add.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "custom-header.value",
			Usage:      "The value of the header.",
			InnerField: "value",
		},
	},
	"dialogflow-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "dialogflow-config.analyze-sentiment",
			Usage:      "Enable sentiment analysis from Dialogflow.",
			InnerField: "analyze_sentiment",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "dialogflow-config.partial-automated-agent-reply",
			Usage:      "Enable partial automated agent reply from Dialogflow.",
			InnerField: "partial_automated_agent_reply",
		},
	},
	"sip-header": {
		&requestflag.InnerFlag[string]{
			Name:       "sip-header.name",
			Usage:      "The name of the header to add.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sip-header.value",
			Usage:      "The value of the header.",
			InnerField: "value",
		},
	},
	"sound-modifications": {
		&requestflag.InnerFlag[float64]{
			Name:       "sound-modifications.octaves",
			Usage:      "Adjust the pitch in octaves, values should be between -1 and 1, default 0",
			InnerField: "octaves",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "sound-modifications.pitch",
			Usage:      "Set the pitch directly, value should be > 0, default 1 (lower = lower tone)",
			InnerField: "pitch",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "sound-modifications.semitone",
			Usage:      "Adjust the pitch in semitones, values should be between -14 and 14, default 0",
			InnerField: "semitone",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sound-modifications.track",
			Usage:      "The track to which the sound modifications will be applied. Accepted values are `inbound` or `outbound`",
			InnerField: "track",
		},
	},
	"transcription-config": {
		&requestflag.InnerFlag[string]{
			Name:       "transcription-config.client-state",
			Usage:      "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			InnerField: "client_state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "transcription-config.command-id",
			Usage:      "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			InnerField: "command_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "transcription-config.transcription-engine",
			Usage:      "Engine to use for speech recognition. Legacy values `A` - `Google`, `B` - `Telnyx` are supported for backward compatibility.",
			InnerField: "transcription_engine",
		},
		&requestflag.InnerFlag[any]{
			Name:       "transcription-config.transcription-engine-config",
			InnerField: "transcription_engine_config",
		},
		&requestflag.InnerFlag[string]{
			Name:       "transcription-config.transcription-tracks",
			Usage:      "Indicates which leg of the call will be transcribed. Use `inbound` for the leg that requested the transcription, `outbound` for the other leg, and `both` for both legs of the call. Will default to `inbound`.",
			InnerField: "transcription_tracks",
		},
	},
})

var callsRetrieveStatus = cli.Command{
	Name:    "retrieve-status",
	Usage:   "Returns the status of a call (data is available 10 minutes after call ended).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
	},
	Action:          handleCallsRetrieveStatus,
	HideHelpCommand: true,
}

func handleCallsDial(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallDialParams{}

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
	_, err = client.Calls.Dial(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls dial", obj, format, transform)
}

func handleCallsRetrieveStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
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
	_, err = client.Calls.GetStatus(ctx, cmd.Value("call-control-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls retrieve-status", obj, format, transform)
}

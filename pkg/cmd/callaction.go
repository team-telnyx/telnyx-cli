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

var callsActionsAddAIAssistantMessages = cli.Command{
	Name:    "add-ai-assistant-messages",
	Usage:   "Add messages to the conversation started by an AI assistant on the call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "message",
			Usage:    "The messages to add to the conversation.",
			BodyPath: "messages",
		},
	},
	Action:          handleCallsActionsAddAIAssistantMessages,
	HideHelpCommand: true,
}

var callsActionsAnswer = requestflag.WithInnerFlags(cli.Command{
	Name:    "answer",
	Usage:   "Answer an incoming call. You must issue this command before executing subsequent\ncommands on an incoming call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "billing-group-id",
			Usage:    "Use this field to set the Billing Group ID for the call. Must be a valid and existing Billing Group ID.",
			BodyPath: "billing_group_id",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-header",
			Usage:    "Custom headers to be added to the SIP INVITE response.",
			BodyPath: "custom_headers",
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
		&requestflag.Flag[[]map[string]any]{
			Name:     "sip-header",
			Usage:    "SIP headers to be added to the SIP INVITE response. Currently only User-to-User header is supported.",
			BodyPath: "sip_headers",
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
	Action:          handleCallsActionsAnswer,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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

var callsActionsBridge = cli.Command{
	Name:    "bridge",
	Usage:   "Bridge two call control calls.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id-to-bridge",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-control-id-to-bridge-with",
			Usage:    "The Call Control ID of the call you want to bridge with, can't be used together with queue parameter or video_room_id parameter.",
			Required: true,
			BodyPath: "call_control_id",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "mute-dtmf",
			Usage:    "When enabled, DTMF tones are not passed to the call participant. The webhooks containing the DTMF information will be sent.",
			Default:  "none",
			BodyPath: "mute_dtmf",
		},
		&requestflag.Flag[string]{
			Name:     "park-after-unbridge",
			Usage:    "Specifies behavior after the bridge ends (i.e. the opposite leg either hangs up or is transferred). If supplied with the value `self`, the current leg will be parked after unbridge. If not set, the default behavior is to hang up the leg.",
			BodyPath: "park_after_unbridge",
		},
		&requestflag.Flag[bool]{
			Name:     "play-ringtone",
			Usage:    "Specifies whether to play a ringtone if the call you want to bridge with has not yet been answered.",
			Default:  false,
			BodyPath: "play_ringtone",
		},
		&requestflag.Flag[string]{
			Name:     "queue",
			Usage:    "The name of the queue you want to bridge with, can't be used together with call_control_id parameter or video_room_id parameter. Bridging with a queue means bridging with the first call in the queue. The call will always be removed from the queue regardless of whether bridging succeeds. Returns an error when the queue is empty.",
			BodyPath: "queue",
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
		&requestflag.Flag[string]{
			Name:     "ringtone",
			Usage:    "Specifies which country ringtone to play when `play_ringtone` is set to `true`. If not set, the US ringtone will be played.",
			Default:  "us",
			BodyPath: "ringtone",
		},
		&requestflag.Flag[string]{
			Name:     "video-room-context",
			Usage:    "The additional parameter that will be passed to the video conference. It is a text field and the user can decide how to use it. For example, you can set the participant name or pass JSON text. It can be used only with video_room_id parameter.",
			BodyPath: "video_room_context",
		},
		&requestflag.Flag[string]{
			Name:     "video-room-id",
			Usage:    "The ID of the video room you want to bridge with, can't be used together with call_control_id parameter or queue parameter.",
			BodyPath: "video_room_id",
		},
	},
	Action:          handleCallsActionsBridge,
	HideHelpCommand: true,
}

var callsActionsEnqueue = cli.Command{
	Name:    "enqueue",
	Usage:   "Put the call in a queue.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Usage:    "The name of the queue the call should be put in. If a queue with a given name doesn't exist yet it will be created.",
			Required: true,
			BodyPath: "queue_name",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[bool]{
			Name:     "keep-after-hangup",
			Usage:    "If set to true, the call will remain in the queue after hangup. In this case bridging to such call will fail with necessary information needed to re-establish the call.",
			Default:  false,
			BodyPath: "keep_after_hangup",
		},
		&requestflag.Flag[int64]{
			Name:     "max-size",
			Usage:    "The maximum number of calls allowed in the queue at a given time. Can't be modified for an existing queue.",
			Default:  100,
			BodyPath: "max_size",
		},
		&requestflag.Flag[int64]{
			Name:     "max-wait-time-secs",
			Usage:    "The number of seconds after which the call will be removed from the queue.",
			BodyPath: "max_wait_time_secs",
		},
	},
	Action:          handleCallsActionsEnqueue,
	HideHelpCommand: true,
}

var callsActionsGather = cli.Command{
	Name:    "gather",
	Usage:   "Gather DTMF signals to build interactive menus.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "gather-id",
			Usage:    "An id that will be sent back in the corresponding `call.gather.ended` webhook. Will be randomly generated if not specified.",
			BodyPath: "gather_id",
		},
		&requestflag.Flag[int64]{
			Name:     "initial-timeout-millis",
			Usage:    "The number of milliseconds to wait for the first DTMF.",
			Default:  5000,
			BodyPath: "initial_timeout_millis",
		},
		&requestflag.Flag[int64]{
			Name:     "inter-digit-timeout-millis",
			Usage:    "The number of milliseconds to wait for input between digits.",
			Default:  5000,
			BodyPath: "inter_digit_timeout_millis",
		},
		&requestflag.Flag[int64]{
			Name:     "maximum-digits",
			Usage:    "The maximum number of digits to fetch. This parameter has a maximum value of 128.",
			Default:  128,
			BodyPath: "maximum_digits",
		},
		&requestflag.Flag[int64]{
			Name:     "minimum-digits",
			Usage:    "The minimum number of digits to fetch. This parameter has a minimum value of 1.",
			Default:  1,
			BodyPath: "minimum_digits",
		},
		&requestflag.Flag[string]{
			Name:     "terminating-digit",
			Usage:    "The digit used to terminate input if fewer than `maximum_digits` digits have been gathered.",
			Default:  "#",
			BodyPath: "terminating_digit",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-millis",
			Usage:    "The number of milliseconds to wait to complete the request.",
			Default:  60000,
			BodyPath: "timeout_millis",
		},
		&requestflag.Flag[string]{
			Name:     "valid-digits",
			Usage:    "A list of all digits accepted as valid.",
			Default:  "0123456789#*",
			BodyPath: "valid_digits",
		},
	},
	Action:          handleCallsActionsGather,
	HideHelpCommand: true,
}

var callsActionsGatherUsingAI = requestflag.WithInnerFlags(cli.Command{
	Name:    "gather-using-ai",
	Usage:   "Gather parameters defined in the request payload using a voice assistant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parameters",
			Usage:    "The parameters described as a JSON Schema object that needs to be gathered by the voice assistant. See the [JSON Schema reference](https://json-schema.org/understanding-json-schema) for documentation about the format",
			Required: true,
			BodyPath: "parameters",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "assistant",
			Usage:    "Assistant configuration including choice of LLM, custom instructions, and tools.",
			BodyPath: "assistant",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "gather-ended-speech",
			Usage:    "Text that will be played when the gathering has finished. There is a 3,000 character limit.",
			BodyPath: "gather_ended_speech",
		},
		&requestflag.Flag[string]{
			Name:     "greeting",
			Usage:    "Text that will be played when the gathering starts, if none then nothing will be played when the gathering starts. The greeting can be text for any voice or SSML for `AWS.Polly.<voice_id>` voices. There is a 3,000 character limit.",
			BodyPath: "greeting",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "interruption-settings",
			Usage:    "Settings for handling user interruptions during assistant speech",
			BodyPath: "interruption_settings",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "Language to use for speech recognition",
			Default:  "en",
			BodyPath: "language",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "message-history",
			Usage:    "The message history you want the voice assistant to be aware of, this can be useful to keep the context of the conversation, or to pass additional information to the voice assistant.",
			BodyPath: "message_history",
		},
		&requestflag.Flag[bool]{
			Name:     "send-message-history-updates",
			Usage:    "Default is `false`. If set to `true`, the voice assistant will send updates to the message history via the `call.ai_gather.message_history_updated` callback in real time as the message history is updated.",
			BodyPath: "send_message_history_updates",
		},
		&requestflag.Flag[bool]{
			Name:     "send-partial-results",
			Usage:    "Default is `false`. If set to `true`, the voice assistant will send partial results via the `call.ai_gather.partial_results` callback in real time as individual fields are gathered. If set to `false`, the voice assistant will only send the final result via the `call.ai_gather.ended` callback.",
			BodyPath: "send_partial_results",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transcription",
			Usage:    "The settings associated with speech to text for the voice assistant. This is only relevant if the assistant uses a text-to-text language model. Any assistant using a model with native audio support (e.g. `fixie-ai/ultravox-v0_4`) will ignore this field.",
			BodyPath: "transcription",
		},
		&requestflag.Flag[int64]{
			Name:     "user-response-timeout-ms",
			Usage:    "The maximum time in milliseconds to wait for user response before timing out.",
			Default:  10000,
			BodyPath: "user_response_timeout_ms",
		},
		&requestflag.Flag[string]{
			Name:     "voice",
			Usage:    "The voice to be used by the voice assistant. Currently we support ElevenLabs, Telnyx and AWS voices.\n\n **Supported Providers:**\n- **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural voices, which provide more realistic, human-like speech, append `-Neural` to the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility.\n- **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural, Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural, Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)\n- **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g., `ElevenLabs.BaseModel.John`). The `ModelId` part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration secret under `\"voice_settings\": {\"api_key_ref\": \"<secret_id>\"}`. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).\n - **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`",
			Default:  "Telnyx.KokoroTTS.af",
			BodyPath: "voice",
		},
		&requestflag.Flag[any]{
			Name:     "voice-settings",
			Usage:    "The settings associated with the voice selected",
			BodyPath: "voice_settings",
		},
	},
	Action:          handleCallsActionsGatherUsingAI,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"assistant": {
		&requestflag.InnerFlag[string]{
			Name:       "assistant.instructions",
			Usage:      "The system instructions that the voice assistant uses during the gather command",
			InnerField: "instructions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assistant.model",
			Usage:      "The model to be used by the voice assistant.",
			InnerField: "model",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assistant.openai-api-key-ref",
			Usage:      "This is necessary only if the model selected is from OpenAI. You would pass the `identifier` for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) that refers to your OpenAI API Key. Warning: Free plans are unlikely to work with this integration.",
			InnerField: "openai_api_key_ref",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "assistant.tools",
			Usage:      "The tools that the voice assistant can use.",
			InnerField: "tools",
		},
	},
	"interruption-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "interruption-settings.enable",
			Usage:      "When true, allows users to interrupt the assistant while speaking",
			InnerField: "enable",
		},
	},
	"message-history": {
		&requestflag.InnerFlag[string]{
			Name:       "message-history.content",
			Usage:      "The content of the message",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message-history.role",
			Usage:      "The role of the message sender",
			InnerField: "role",
		},
	},
	"transcription": {
		&requestflag.InnerFlag[string]{
			Name:       "transcription.model",
			Usage:      "The speech to text model to be used by the voice assistant.\n\n- `distil-whisper/distil-large-v2` is lower latency but English-only.\n- `openai/whisper-large-v3-turbo` is multi-lingual with automatic language detection but slightly higher latency.\n- `google` is a multi-lingual option, please describe the language in the `language` field.",
			InnerField: "model",
		},
	},
})

var callsActionsGatherUsingAudio = cli.Command{
	Name:    "gather-using-audio",
	Usage:   "Play an audio file on the call until the required DTMF signals are gathered to\nbuild interactive menus.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "audio-url",
			Usage:    "The URL of a file to be played back at the beginning of each prompt. The URL can point to either a WAV or MP3 file. media_name and audio_url cannot be used together in one request.",
			BodyPath: "audio_url",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[int64]{
			Name:     "inter-digit-timeout-millis",
			Usage:    "The number of milliseconds to wait for input between digits.",
			Default:  5000,
			BodyPath: "inter_digit_timeout_millis",
		},
		&requestflag.Flag[string]{
			Name:     "invalid-audio-url",
			Usage:    "The URL of a file to play when digits don't match the `valid_digits` parameter or the number of digits is not between `min` and `max`. The URL can point to either a WAV or MP3 file. invalid_media_name and invalid_audio_url cannot be used together in one request.",
			BodyPath: "invalid_audio_url",
		},
		&requestflag.Flag[string]{
			Name:     "invalid-media-name",
			Usage:    "The media_name of a file to be played back when digits don't match the `valid_digits` parameter or the number of digits is not between `min` and `max`. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file.",
			BodyPath: "invalid_media_name",
		},
		&requestflag.Flag[int64]{
			Name:     "maximum-digits",
			Usage:    "The maximum number of digits to fetch. This parameter has a maximum value of 128.",
			Default:  128,
			BodyPath: "maximum_digits",
		},
		&requestflag.Flag[int64]{
			Name:     "maximum-tries",
			Usage:    "The maximum number of times the file should be played if there is no input from the user on the call.",
			Default:  3,
			BodyPath: "maximum_tries",
		},
		&requestflag.Flag[string]{
			Name:     "media-name",
			Usage:    "The media_name of a file to be played back at the beginning of each prompt. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file.",
			BodyPath: "media_name",
		},
		&requestflag.Flag[int64]{
			Name:     "minimum-digits",
			Usage:    "The minimum number of digits to fetch. This parameter has a minimum value of 1.",
			Default:  1,
			BodyPath: "minimum_digits",
		},
		&requestflag.Flag[string]{
			Name:     "terminating-digit",
			Usage:    "The digit used to terminate input if fewer than `maximum_digits` digits have been gathered.",
			Default:  "#",
			BodyPath: "terminating_digit",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-millis",
			Usage:    "The number of milliseconds to wait for a DTMF response after file playback ends before a replaying the sound file.",
			Default:  60000,
			BodyPath: "timeout_millis",
		},
		&requestflag.Flag[string]{
			Name:     "valid-digits",
			Usage:    "A list of all digits accepted as valid.",
			Default:  "0123456789#*",
			BodyPath: "valid_digits",
		},
	},
	Action:          handleCallsActionsGatherUsingAudio,
	HideHelpCommand: true,
}

var callsActionsGatherUsingSpeak = cli.Command{
	Name:    "gather-using-speak",
	Usage:   "Convert text to speech and play it on the call until the required DTMF signals\nare gathered to build interactive menus.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "payload",
			Usage:    "The text or SSML to be converted into speech. There is a 3,000 character limit.",
			Required: true,
			BodyPath: "payload",
		},
		&requestflag.Flag[string]{
			Name:     "voice",
			Usage:    "Specifies the voice used in speech synthesis.\n\n- Define voices using the format `<Provider>.<Model>.<VoiceId>`. Specifying only the provider will give default values for voice_id and model_id.\n\n **Supported Providers:**\n- **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural voices, which provide more realistic, human-like speech, append `-Neural` to the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility.\n- **Azure:** Use `Azure.<VoiceId>` (e.g., `Azure.en-CA-ClaraNeural`, `Azure.en-US-BrianMultilingualNeural`, `Azure.en-US-Ava:DragonHDLatestNeural`). For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery). Use `voice_settings` to configure custom deployments, regions, or API keys.\n- **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g., `ElevenLabs.eleven_multilingual_v2.21m00Tcm4TlvDq8ikWAM`). The `ModelId` part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration identifier secret in `\"voice_settings\": {\"api_key_ref\": \"<secret_identifier>\"}`. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).\n- **Telnyx:** Use `Telnyx.<model_id>.<voice_id>` (e.g., `Telnyx.KokoroTTS.af`). Use `voice_settings` to configure voice_speed and other synthesis parameters.\n- **Minimax:** Use `Minimax.<ModelId>.<VoiceId>` (e.g., `Minimax.speech-02-hd.Wise_Woman`). Supported models: `speech-02-turbo`, `speech-02-hd`, `speech-2.6-turbo`, `speech-2.8-turbo`. Use `voice_settings` to configure speed, volume, pitch, and language_boost.\n- **Rime:** Use `Rime.<model_id>.<voice_id>` (e.g., `Rime.Arcana.cove`). Supported model_ids: `Arcana`, `Mist`. Use `voice_settings` to configure voice_speed.\n- **Resemble:** Use `Resemble.Turbo.<voice_id>` (e.g., `Resemble.Turbo.my_voice`). Only `Turbo` model is supported. Use `voice_settings` to configure precision, sample_rate, and format.\n\nFor service_level basic, you may define the gender of the speaker (male or female).",
			Required: true,
			BodyPath: "voice",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[int64]{
			Name:     "inter-digit-timeout-millis",
			Usage:    "The number of milliseconds to wait for input between digits.",
			Default:  5000,
			BodyPath: "inter_digit_timeout_millis",
		},
		&requestflag.Flag[string]{
			Name:     "invalid-payload",
			Usage:    "The text or SSML to be converted into speech when digits don't match the `valid_digits` parameter or the number of digits is not between `min` and `max`. There is a 3,000 character limit.",
			BodyPath: "invalid_payload",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "The language you want spoken. This parameter is ignored when a `Polly.*` voice is specified.",
			BodyPath: "language",
		},
		&requestflag.Flag[int64]{
			Name:     "maximum-digits",
			Usage:    "The maximum number of digits to fetch. This parameter has a maximum value of 128.",
			Default:  128,
			BodyPath: "maximum_digits",
		},
		&requestflag.Flag[int64]{
			Name:     "maximum-tries",
			Usage:    "The maximum number of times that a file should be played back if there is no input from the user on the call.",
			Default:  3,
			BodyPath: "maximum_tries",
		},
		&requestflag.Flag[int64]{
			Name:     "minimum-digits",
			Usage:    "The minimum number of digits to fetch. This parameter has a minimum value of 1.",
			Default:  1,
			BodyPath: "minimum_digits",
		},
		&requestflag.Flag[string]{
			Name:     "payload-type",
			Usage:    "The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML).",
			Default:  "text",
			BodyPath: "payload_type",
		},
		&requestflag.Flag[string]{
			Name:     "service-level",
			Usage:    "This parameter impacts speech quality, language options and payload types. When using `basic`, only the `en-US` language and payload type `text` are allowed.",
			Default:  "premium",
			BodyPath: "service_level",
		},
		&requestflag.Flag[string]{
			Name:     "terminating-digit",
			Usage:    "The digit used to terminate input if fewer than `maximum_digits` digits have been gathered.",
			Default:  "#",
			BodyPath: "terminating_digit",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-millis",
			Usage:    "The number of milliseconds to wait for a DTMF response after speak ends before a replaying the sound file.",
			Default:  60000,
			BodyPath: "timeout_millis",
		},
		&requestflag.Flag[string]{
			Name:     "valid-digits",
			Usage:    "A list of all digits accepted as valid.",
			Default:  "0123456789#*",
			BodyPath: "valid_digits",
		},
		&requestflag.Flag[any]{
			Name:     "voice-settings",
			Usage:    "The settings associated with the voice selected",
			BodyPath: "voice_settings",
		},
	},
	Action:          handleCallsActionsGatherUsingSpeak,
	HideHelpCommand: true,
}

var callsActionsHangup = cli.Command{
	Name:    "hangup",
	Usage:   "Hang up the call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsHangup,
	HideHelpCommand: true,
}

var callsActionsLeaveQueue = cli.Command{
	Name:    "leave-queue",
	Usage:   "Removes the call from a queue.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsLeaveQueue,
	HideHelpCommand: true,
}

var callsActionsPauseRecording = cli.Command{
	Name:    "pause-recording",
	Usage:   "Pause recording the call. Recording can be resumed via Resume recording command.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "recording-id",
			Usage:    "Uniquely identifies the resource.",
			BodyPath: "recording_id",
		},
	},
	Action:          handleCallsActionsPauseRecording,
	HideHelpCommand: true,
}

var callsActionsRefer = requestflag.WithInnerFlags(cli.Command{
	Name:    "refer",
	Usage:   "Initiate a SIP Refer on a Call Control call. You can initiate a SIP Refer at any\npoint in the duration of a call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "sip-address",
			Usage:    "The SIP URI to which the call will be referred to.",
			Required: true,
			BodyPath: "sip_address",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same `command_id` as one that has already been executed.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-header",
			Usage:    "Custom headers to be added to the SIP INVITE.",
			BodyPath: "custom_headers",
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
			Usage:    "SIP headers to be added to the request. Currently only User-to-User header is supported.",
			BodyPath: "sip_headers",
		},
	},
	Action:          handleCallsActionsRefer,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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
})

var callsActionsReject = cli.Command{
	Name:    "reject",
	Usage:   "Reject an incoming call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "cause",
			Usage:    "Cause for call rejection.",
			Required: true,
			BodyPath: "cause",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsReject,
	HideHelpCommand: true,
}

var callsActionsResumeRecording = cli.Command{
	Name:    "resume-recording",
	Usage:   "Resume recording the call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "recording-id",
			Usage:    "Uniquely identifies the resource.",
			BodyPath: "recording_id",
		},
	},
	Action:          handleCallsActionsResumeRecording,
	HideHelpCommand: true,
}

var callsActionsSendDtmf = cli.Command{
	Name:    "send-dtmf",
	Usage:   "Sends DTMF tones from this leg. DTMF tones will be heard by the other end of the\ncall.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "digits",
			Usage:    "DTMF digits to send. Valid digits are 0-9, A-D, *, and #. Pauses can be added using w (0.5s) and W (1s).",
			Required: true,
			BodyPath: "digits",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[int64]{
			Name:     "duration-millis",
			Usage:    "Specifies for how many milliseconds each digit will be played in the audio stream. Ranges from 100 to 500ms",
			Default:  250,
			BodyPath: "duration_millis",
		},
	},
	Action:          handleCallsActionsSendDtmf,
	HideHelpCommand: true,
}

var callsActionsSendSipInfo = cli.Command{
	Name:    "send-sip-info",
	Usage:   "Sends SIP info from this leg.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "body",
			Usage:    "Content of the SIP INFO",
			Required: true,
			BodyPath: "body",
		},
		&requestflag.Flag[string]{
			Name:     "content-type",
			Usage:    "Content type of the INFO body. Must be MIME type compliant. There is a 1,400 bytes limit",
			Required: true,
			BodyPath: "content_type",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsSendSipInfo,
	HideHelpCommand: true,
}

var callsActionsSpeak = cli.Command{
	Name:    "speak",
	Usage:   "Convert text to speech and play it back on the call. If multiple speak text\ncommands are issued consecutively, the audio files will be placed in a queue\nawaiting playback.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "payload",
			Usage:    "The text or SSML to be converted into speech. There is a 3,000 character limit.",
			Required: true,
			BodyPath: "payload",
		},
		&requestflag.Flag[string]{
			Name:     "voice",
			Usage:    "Specifies the voice used in speech synthesis.\n\n- Define voices using the format `<Provider>.<Model>.<VoiceId>`. Specifying only the provider will give default values for voice_id and model_id.\n\n **Supported Providers:**\n- **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural voices, which provide more realistic, human-like speech, append `-Neural` to the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility.\n- **Azure:** Use `Azure.<VoiceId>` (e.g., `Azure.en-CA-ClaraNeural`, `Azure.en-US-BrianMultilingualNeural`, `Azure.en-US-Ava:DragonHDLatestNeural`). For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery). Use `voice_settings` to configure custom deployments, regions, or API keys.\n- **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g., `ElevenLabs.eleven_multilingual_v2.21m00Tcm4TlvDq8ikWAM`). The `ModelId` part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration identifier secret in `\"voice_settings\": {\"api_key_ref\": \"<secret_identifier>\"}`. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).\n- **Telnyx:** Use `Telnyx.<model_id>.<voice_id>` (e.g., `Telnyx.KokoroTTS.af`). Use `voice_settings` to configure voice_speed and other synthesis parameters.\n- **Minimax:** Use `Minimax.<ModelId>.<VoiceId>` (e.g., `Minimax.speech-02-hd.Wise_Woman`). Supported models: `speech-02-turbo`, `speech-02-hd`, `speech-2.6-turbo`, `speech-2.8-turbo`. Use `voice_settings` to configure speed, volume, pitch, and language_boost.\n- **Rime:** Use `Rime.<model_id>.<voice_id>` (e.g., `Rime.Arcana.cove`). Supported model_ids: `Arcana`, `Mist`. Use `voice_settings` to configure voice_speed.\n- **Resemble:** Use `Resemble.Turbo.<voice_id>` (e.g., `Resemble.Turbo.my_voice`). Only `Turbo` model is supported. Use `voice_settings` to configure precision, sample_rate, and format.\n\nFor service_level basic, you may define the gender of the speaker (male or female).",
			Required: true,
			BodyPath: "voice",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "The language you want spoken. This parameter is ignored when a `Polly.*` voice is specified.",
			BodyPath: "language",
		},
		&requestflag.Flag[any]{
			Name:     "loop",
			BodyPath: "loop",
		},
		&requestflag.Flag[string]{
			Name:     "payload-type",
			Usage:    "The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML).",
			Default:  "text",
			BodyPath: "payload_type",
		},
		&requestflag.Flag[string]{
			Name:     "service-level",
			Usage:    "This parameter impacts speech quality, language options and payload types. When using `basic`, only the `en-US` language and payload type `text` are allowed.",
			Default:  "premium",
			BodyPath: "service_level",
		},
		&requestflag.Flag[string]{
			Name:     "stop",
			Usage:    "When specified, it stops the current audio being played. Specify `current` to stop the current audio being played, and to play the next file in the queue. Specify `all` to stop the current audio file being played and to also clear all audio files from the queue.",
			BodyPath: "stop",
		},
		&requestflag.Flag[string]{
			Name:     "target-legs",
			Usage:    "Specifies which legs of the call should receive the spoken audio.",
			Default:  "self",
			BodyPath: "target_legs",
		},
		&requestflag.Flag[any]{
			Name:     "voice-settings",
			Usage:    "The settings associated with the voice selected",
			BodyPath: "voice_settings",
		},
	},
	Action:          handleCallsActionsSpeak,
	HideHelpCommand: true,
}

var callsActionsStartAIAssistant = requestflag.WithInnerFlags(cli.Command{
	Name:    "start-ai-assistant",
	Usage:   "Start an AI assistant on the call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "assistant",
			Usage:    "AI Assistant configuration",
			BodyPath: "assistant",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "greeting",
			Usage:    "Text that will be played when the assistant starts, if none then nothing will be played when the assistant starts. The greeting can be text for any voice or SSML for `AWS.Polly.<voice_id>` voices. There is a 3,000 character limit.",
			BodyPath: "greeting",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "interruption-settings",
			Usage:    "Settings for handling user interruptions during assistant speech",
			BodyPath: "interruption_settings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transcription",
			Usage:    "The settings associated with speech to text for the voice assistant. This is only relevant if the assistant uses a text-to-text language model. Any assistant using a model with native audio support (e.g. `fixie-ai/ultravox-v0_4`) will ignore this field.",
			BodyPath: "transcription",
		},
		&requestflag.Flag[string]{
			Name:     "voice",
			Usage:    "The voice to be used by the voice assistant. Currently we support ElevenLabs, Telnyx and AWS voices.\n\n **Supported Providers:**\n- **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural voices, which provide more realistic, human-like speech, append `-Neural` to the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility.\n- **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural, Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural, Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)\n- **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g., `ElevenLabs.BaseModel.John`). The `ModelId` part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration secret under `\"voice_settings\": {\"api_key_ref\": \"<secret_id>\"}`. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).\n - **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`",
			Default:  "Telnyx.KokoroTTS.af",
			BodyPath: "voice",
		},
		&requestflag.Flag[any]{
			Name:     "voice-settings",
			Usage:    "The settings associated with the voice selected",
			BodyPath: "voice_settings",
		},
	},
	Action:          handleCallsActionsStartAIAssistant,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"assistant": {
		&requestflag.InnerFlag[string]{
			Name:       "assistant.id",
			Usage:      "The identifier of the AI assistant to use",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assistant.instructions",
			Usage:      "The system instructions that the voice assistant uses during the start assistant command. This will overwrite the instructions set in the assistant configuration.",
			InnerField: "instructions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assistant.openai-api-key-ref",
			Usage:      "Reference to the OpenAI API key. Required only when using OpenAI models",
			InnerField: "openai_api_key_ref",
		},
	},
	"interruption-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "interruption-settings.enable",
			Usage:      "When true, allows users to interrupt the assistant while speaking",
			InnerField: "enable",
		},
	},
	"transcription": {
		&requestflag.InnerFlag[string]{
			Name:       "transcription.model",
			Usage:      "The speech to text model to be used by the voice assistant.\n\n- `distil-whisper/distil-large-v2` is lower latency but English-only.\n- `openai/whisper-large-v3-turbo` is multi-lingual with automatic language detection but slightly higher latency.\n- `google` is a multi-lingual option, please describe the language in the `language` field.",
			InnerField: "model",
		},
	},
})

var callsActionsStartForking = cli.Command{
	Name:    "start-forking",
	Usage:   "Call forking allows you to stream the media from a call to a specific target in\nrealtime. This stream can be used to enable realtime audio analysis to support a\nvariety of use cases, including fraud detection, or the creation of AI-generated\naudio responses. Requests must specify either the `target` attribute or the `rx`\nand `tx` attributes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "rx",
			Usage:    "The network target, <udp:ip_address:port>, where the call's incoming RTP media packets should be forwarded.",
			BodyPath: "rx",
		},
		&requestflag.Flag[string]{
			Name:     "stream-type",
			Usage:    "Optionally specify a media type to stream. If `decrypted` selected, Telnyx will decrypt incoming SIP media before forking to the target. `rx` and `tx` are required fields if `decrypted` selected.",
			Default:  "decrypted",
			BodyPath: "stream_type",
		},
		&requestflag.Flag[string]{
			Name:     "tx",
			Usage:    "The network target, <udp:ip_address:port>, where the call's outgoing RTP media packets should be forwarded.",
			BodyPath: "tx",
		},
	},
	Action:          handleCallsActionsStartForking,
	HideHelpCommand: true,
}

var callsActionsStartNoiseSuppression = requestflag.WithInnerFlags(cli.Command{
	Name:    "start-noise-suppression",
	Usage:   "Noise Suppression Start (BETA)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "direction",
			Usage:    "The direction of the audio stream to be noise suppressed.",
			Default:  "inbound",
			BodyPath: "direction",
		},
		&requestflag.Flag[string]{
			Name:     "noise-suppression-engine",
			Usage:    "The engine to use for noise suppression.\nFor backward compatibility, engines A, B, and C are also supported, but are deprecated:\n A - Denoiser\n B - DeepFilterNet\n C - Krisp",
			Default:  "Denoiser",
			BodyPath: "noise_suppression_engine",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "noise-suppression-engine-config",
			Usage:    "Configuration parameters for noise suppression engines.",
			BodyPath: "noise_suppression_engine_config",
		},
	},
	Action:          handleCallsActionsStartNoiseSuppression,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"noise-suppression-engine-config": {
		&requestflag.InnerFlag[int64]{
			Name:       "noise-suppression-engine-config.attenuation-limit",
			Usage:      "The attenuation limit for noise suppression (0-100). Only applicable for DeepFilterNet.",
			InnerField: "attenuation_limit",
		},
	},
})

var callsActionsStartPlayback = cli.Command{
	Name:    "start-playback",
	Usage:   "Play an audio file on the call. If multiple play audio commands are issued\nconsecutively, the audio files will be placed in a queue awaiting playback.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "audio-type",
			Usage:    "Specifies the type of audio provided in `audio_url` or `playback_content`.",
			Default:  "mp3",
			BodyPath: "audio_type",
		},
		&requestflag.Flag[string]{
			Name:     "audio-url",
			Usage:    "The URL of a file to be played back on the call. The URL can point to either a WAV or MP3 file. media_name and audio_url cannot be used together in one request.",
			BodyPath: "audio_url",
		},
		&requestflag.Flag[bool]{
			Name:     "cache-audio",
			Usage:    "Caches the audio file. Useful when playing the same audio file multiple times during the call.",
			Default:  true,
			BodyPath: "cache_audio",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[any]{
			Name:     "loop",
			BodyPath: "loop",
		},
		&requestflag.Flag[string]{
			Name:     "media-name",
			Usage:    "The media_name of a file to be played back on the call. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file.",
			BodyPath: "media_name",
		},
		&requestflag.Flag[bool]{
			Name:     "overlay",
			Usage:    "When enabled, audio will be mixed on top of any other audio that is actively being played back. Note that `overlay: true` will only work if there is another audio file already being played on the call.",
			Default:  false,
			BodyPath: "overlay",
		},
		&requestflag.Flag[string]{
			Name:     "playback-content",
			Usage:    "Allows a user to provide base64 encoded mp3 or wav. Note: when using this parameter, `media_url` and `media_name` in the `playback_started` and `playback_ended` webhooks will be empty",
			BodyPath: "playback_content",
		},
		&requestflag.Flag[string]{
			Name:     "stop",
			Usage:    "When specified, it stops the current audio being played. Specify `current` to stop the current audio being played, and to play the next file in the queue. Specify `all` to stop the current audio file being played and to also clear all audio files from the queue.",
			BodyPath: "stop",
		},
		&requestflag.Flag[string]{
			Name:     "target-legs",
			Usage:    "Specifies the leg or legs on which audio will be played. If supplied, the value must be either `self`, `opposite` or `both`.",
			Default:  "self",
			BodyPath: "target_legs",
		},
	},
	Action:          handleCallsActionsStartPlayback,
	HideHelpCommand: true,
}

var callsActionsStartRecording = cli.Command{
	Name:    "start-recording",
	Usage:   "Start recording the call. Recording will stop on call hang-up, or can be\ninitiated via the Stop Recording command.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "channels",
			Usage:    "When `dual`, final audio file will be stereo recorded with the first leg on channel A, and the rest on channel B.",
			Required: true,
			BodyPath: "channels",
		},
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "The audio file format used when storing the call recording. Can be either `mp3` or `wav`.",
			Required: true,
			BodyPath: "format",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "custom-file-name",
			Usage:    "The custom recording file name to be used instead of the default `call_leg_id`. Telnyx will still add a Unix timestamp suffix.",
			BodyPath: "custom_file_name",
		},
		&requestflag.Flag[int64]{
			Name:     "max-length",
			Usage:    "Defines the maximum length for the recording in seconds. The minimum value is 0. The maximum value is 14400. The default value is 0 (infinite)",
			Default:  0,
			BodyPath: "max_length",
		},
		&requestflag.Flag[bool]{
			Name:     "play-beep",
			Usage:    "If enabled, a beep sound will be played at the start of a recording.",
			BodyPath: "play_beep",
		},
		&requestflag.Flag[string]{
			Name:     "recording-track",
			Usage:    "The audio track to be recorded. Can be either `both`, `inbound` or `outbound`. If only single track is specified (`inbound`, `outbound`), `channels` configuration is ignored and it will be recorded as mono (single channel).",
			Default:  "both",
			BodyPath: "recording_track",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-secs",
			Usage:    "The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected. The timer only starts when the speech is detected. Please note that call transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite)",
			Default:  0,
			BodyPath: "timeout_secs",
		},
		&requestflag.Flag[bool]{
			Name:     "transcription",
			Usage:    "Enable post recording transcription. The default value is false.",
			Default:  false,
			BodyPath: "transcription",
		},
		&requestflag.Flag[string]{
			Name:     "transcription-engine",
			Usage:    "Engine to use for speech recognition. `A` - `Google`, `B` - `Telnyx`, `deepgram/nova-3` - `Deepgram Nova-3`. Note: `deepgram/nova-3` supports only `en` and `en-{Region}` languages.",
			Default:  "A",
			BodyPath: "transcription_engine",
		},
		&requestflag.Flag[string]{
			Name:     "transcription-language",
			Usage:    "Language code for transcription. Note: Not all languages are supported by all transcription engines (google, telnyx, deepgram). See engine-specific documentation for supported values.",
			Default:  "en-US",
			BodyPath: "transcription_language",
		},
		&requestflag.Flag[int64]{
			Name:     "transcription-max-speaker-count",
			Usage:    "Defines maximum number of speakers in the conversation. Applies to `google` engine only.",
			Default:  6,
			BodyPath: "transcription_max_speaker_count",
		},
		&requestflag.Flag[int64]{
			Name:     "transcription-min-speaker-count",
			Usage:    "Defines minimum number of speakers in the conversation. Applies to `google` engine only.",
			Default:  2,
			BodyPath: "transcription_min_speaker_count",
		},
		&requestflag.Flag[bool]{
			Name:     "transcription-profanity-filter",
			Usage:    "Enables profanity_filter. Applies to `google` engine only.",
			Default:  false,
			BodyPath: "transcription_profanity_filter",
		},
		&requestflag.Flag[bool]{
			Name:     "transcription-speaker-diarization",
			Usage:    "Enables speaker diarization. Applies to `google` engine only.",
			Default:  false,
			BodyPath: "transcription_speaker_diarization",
		},
		&requestflag.Flag[string]{
			Name:     "trim",
			Usage:    "When set to `trim-silence`, silence will be removed from the beginning and end of the recording.",
			BodyPath: "trim",
		},
	},
	Action:          handleCallsActionsStartRecording,
	HideHelpCommand: true,
}

var callsActionsStartSiprec = cli.Command{
	Name:    "start-siprec",
	Usage:   "Start siprec session to configured in SIPREC connector SRS.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "connector-name",
			Usage:    "Name of configured SIPREC connector to be used.",
			BodyPath: "connector_name",
		},
		&requestflag.Flag[bool]{
			Name:     "include-metadata-custom-headers",
			Usage:    "When set, custom parameters will be added as metadata (recording.session.ExtensionParameters). Otherwise, they’ll be added to sip headers.",
			BodyPath: "include_metadata_custom_headers",
		},
		&requestflag.Flag[bool]{
			Name:     "secure",
			Usage:    "Controls whether to encrypt media sent to your SRS using SRTP and TLS. When set you need to configure SRS port in your connector to 5061.",
			BodyPath: "secure",
		},
		&requestflag.Flag[int64]{
			Name:     "session-timeout-secs",
			Usage:    "Sets `Session-Expires` header to the INVITE. A reinvite is sent every half the value set. Usefull for session keep alive. Minimum value is 90, set to 0 to disable.",
			Default:  1800,
			BodyPath: "session_timeout_secs",
		},
		&requestflag.Flag[string]{
			Name:     "sip-transport",
			Usage:    "Specifies SIP transport protocol.",
			Default:  "udp",
			BodyPath: "sip_transport",
		},
		&requestflag.Flag[string]{
			Name:     "siprec-track",
			Usage:    "Specifies which track should be sent on siprec session.",
			Default:  "both_tracks",
			BodyPath: "siprec_track",
		},
	},
	Action:          handleCallsActionsStartSiprec,
	HideHelpCommand: true,
}

var callsActionsStartStreaming = requestflag.WithInnerFlags(cli.Command{
	Name:    "start-streaming",
	Usage:   "Start streaming the media from a call to a specific WebSocket address or\nDialogflow connection in near-realtime. Audio will be delivered as\nbase64-encoded RTP payload (raw audio), wrapped in JSON payloads.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-parameter",
			Usage:    "Custom parameters to be sent as part of the WebSocket connection.",
			BodyPath: "custom_parameters",
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
			Name:     "stream-auth-token",
			Usage:    "An authentication token to be sent as part of the WebSocket connection. Maximum length is 4000 characters.",
			BodyPath: "stream_auth_token",
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
	},
	Action:          handleCallsActionsStartStreaming,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"custom-parameter": {
		&requestflag.InnerFlag[string]{
			Name:       "custom-parameter.name",
			Usage:      "The name of the custom parameter.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "custom-parameter.value",
			Usage:      "The value of the custom parameter.",
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
})

var callsActionsStartTranscription = cli.Command{
	Name:    "start-transcription",
	Usage:   "Start real-time transcription. Transcription will stop on call hang-up, or can\nbe initiated via the Transcription stop command.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "transcription-engine",
			Usage:    "Engine to use for speech recognition. Legacy values `A` - `Google`, `B` - `Telnyx` are supported for backward compatibility.",
			Default:  "Google",
			BodyPath: "transcription_engine",
		},
		&requestflag.Flag[any]{
			Name:     "transcription-engine-config",
			BodyPath: "transcription_engine_config",
		},
		&requestflag.Flag[string]{
			Name:     "transcription-tracks",
			Usage:    "Indicates which leg of the call will be transcribed. Use `inbound` for the leg that requested the transcription, `outbound` for the other leg, and `both` for both legs of the call. Will default to `inbound`.",
			Default:  "inbound",
			BodyPath: "transcription_tracks",
		},
	},
	Action:          handleCallsActionsStartTranscription,
	HideHelpCommand: true,
}

var callsActionsStopAIAssistant = cli.Command{
	Name:    "stop-ai-assistant",
	Usage:   "Stop an AI assistant on the call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsStopAIAssistant,
	HideHelpCommand: true,
}

var callsActionsStopForking = cli.Command{
	Name:    "stop-forking",
	Usage:   "Stop forking a call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "stream-type",
			Usage:    "Optionally specify a `stream_type`. This should match the `stream_type` that was used in `fork_start` command to properly stop the fork.",
			Default:  "raw",
			BodyPath: "stream_type",
		},
	},
	Action:          handleCallsActionsStopForking,
	HideHelpCommand: true,
}

var callsActionsStopGather = cli.Command{
	Name:    "stop-gather",
	Usage:   "Stop current gather.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsStopGather,
	HideHelpCommand: true,
}

var callsActionsStopNoiseSuppression = cli.Command{
	Name:    "stop-noise-suppression",
	Usage:   "Noise Suppression Stop (BETA)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsStopNoiseSuppression,
	HideHelpCommand: true,
}

var callsActionsStopPlayback = cli.Command{
	Name:    "stop-playback",
	Usage:   "Stop audio being played on the call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[bool]{
			Name:     "overlay",
			Usage:    "When enabled, it stops the audio being played in the overlay queue.",
			Default:  false,
			BodyPath: "overlay",
		},
		&requestflag.Flag[string]{
			Name:     "stop",
			Usage:    "Use `current` to stop the current audio being played. Use `all` to stop the current audio file being played and clear all audio files from the queue.",
			Default:  "all",
			BodyPath: "stop",
		},
	},
	Action:          handleCallsActionsStopPlayback,
	HideHelpCommand: true,
}

var callsActionsStopRecording = cli.Command{
	Name:    "stop-recording",
	Usage:   "Stop recording the call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "recording-id",
			Usage:    "Uniquely identifies the resource.",
			BodyPath: "recording_id",
		},
	},
	Action:          handleCallsActionsStopRecording,
	HideHelpCommand: true,
}

var callsActionsStopSiprec = cli.Command{
	Name:    "stop-siprec",
	Usage:   "Stop SIPREC session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsStopSiprec,
	HideHelpCommand: true,
}

var callsActionsStopStreaming = cli.Command{
	Name:    "stop-streaming",
	Usage:   "Stop streaming a call to a WebSocket.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "stream-id",
			Usage:    "Identifies the stream. If the `stream_id` is not provided the command stops all streams associated with a given `call_control_id`.",
			BodyPath: "stream_id",
		},
	},
	Action:          handleCallsActionsStopStreaming,
	HideHelpCommand: true,
}

var callsActionsStopTranscription = cli.Command{
	Name:    "stop-transcription",
	Usage:   "Stop real-time transcription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
	},
	Action:          handleCallsActionsStopTranscription,
	HideHelpCommand: true,
}

var callsActionsSwitchSupervisorRole = cli.Command{
	Name:    "switch-supervisor-role",
	Usage:   "Switch the supervisor role for a bridged call. This allows switching between\ndifferent supervisor modes during an active call",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "The supervisor role to switch to. 'barge' allows speaking to both parties, 'whisper' allows speaking to caller only, 'monitor' allows listening only.",
			Required: true,
			BodyPath: "role",
		},
	},
	Action:          handleCallsActionsSwitchSupervisorRole,
	HideHelpCommand: true,
}

var callsActionsTransfer = requestflag.WithInnerFlags(cli.Command{
	Name:    "transfer",
	Usage:   "Transfer a call to a new destination. If the transfer is unsuccessful, a\n`call.hangup` webhook for the other call (Leg B) will be sent indicating that\nthe transfer could not be completed. The original call will remain active and\nmay be issued additional commands, potentially transfering the call to an\nalternate destination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "The DID or SIP URI to dial out to.",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[string]{
			Name:     "answering-machine-detection",
			Usage:    "Enables Answering Machine Detection. When a call is answered, Telnyx runs real-time detection to determine if it was picked up by a human or a machine and sends an `call.machine.detection.ended` webhook with the analysis result. If 'greeting_end' or 'detect_words' is used and a 'machine' is detected, you will receive another 'call.machine.greeting.ended' webhook when the answering machine greeting ends with a beep or silence. If `detect_beep` is used, you will only receive 'call.machine.greeting.ended' if a beep is detected.",
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
			Usage:    "The URL of a file to be played back when the transfer destination answers before bridging the call. The URL can point to either a WAV or MP3 file. media_name and audio_url cannot be used together in one request.",
			BodyPath: "audio_url",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-header",
			Usage:    "Custom headers to be added to the SIP INVITE.",
			BodyPath: "custom_headers",
		},
		&requestflag.Flag[bool]{
			Name:     "early-media",
			Usage:    "If set to false, early media will not be passed to the originating leg.",
			Default:  true,
			BodyPath: "early_media",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "The `from` number to be used as the caller id presented to the destination (`to` number). The number should be in +E164 format. This attribute will default to the `to` number of the original call if omitted.",
			BodyPath: "from",
		},
		&requestflag.Flag[string]{
			Name:     "from-display-name",
			Usage:    "The `from_display_name` string to be used as the caller id name (SIP From Display Name) presented to the destination (`to` number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and -_~!.+ special characters. If ommited, the display name will be the same as the number in the `from` field.",
			BodyPath: "from_display_name",
		},
		&requestflag.Flag[string]{
			Name:     "media-encryption",
			Usage:    "Defines whether media should be encrypted on the new call leg.",
			Default:  "disabled",
			BodyPath: "media_encryption",
		},
		&requestflag.Flag[string]{
			Name:     "media-name",
			Usage:    "The media_name of a file to be played back when the transfer destination answers before bridging the call. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file.",
			BodyPath: "media_name",
		},
		&requestflag.Flag[string]{
			Name:     "mute-dtmf",
			Usage:    "When enabled, DTMF tones are not passed to the call participant. The webhooks containing the DTMF information will be sent.",
			Default:  "none",
			BodyPath: "mute_dtmf",
		},
		&requestflag.Flag[string]{
			Name:     "park-after-unbridge",
			Usage:    "Specifies behavior after the bridge ends (i.e. the opposite leg either hangs up or is transferred). If supplied with the value `self`, the current leg will be parked after unbridge. If not set, the default behavior is to hang up the leg.",
			BodyPath: "park_after_unbridge",
		},
		&requestflag.Flag[string]{
			Name:     "preferred-codecs",
			Usage:    "The list of comma-separated codecs in order of preference to be used during the call. The codecs supported are `G722`, `PCMU`, `PCMA`, `G729`, `OPUS`, `VP8`, `H264`, `AMR-WB`.",
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
			Usage:    "SIP headers to be added to the SIP INVITE. Currently only User-to-User header is supported.",
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
			Name:     "target-leg-client-state",
			Usage:    "Use this field to add state to every subsequent webhook for the new leg. It must be a valid Base-64 encoded string.",
			BodyPath: "target_leg_client_state",
		},
		&requestflag.Flag[int64]{
			Name:     "time-limit-secs",
			Usage:    "Sets the maximum duration of a Call Control Leg in seconds. If the time limit is reached, the call will hangup and a `call.hangup` webhook with a `hangup_cause` of `time_limit` will be sent. For example, by setting a time limit of 120 seconds, a Call Leg will be automatically terminated two minutes after being answered. The default time limit is 14400 seconds or 4 hours and this is also the maximum allowed call length.",
			Default:  14400,
			BodyPath: "time_limit_secs",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-secs",
			Usage:    "The number of seconds that Telnyx will wait for the call to be answered by the destination to which it is being transferred. If the timeout is reached before an answer is received, the call will hangup and a `call.hangup` webhook with a `hangup_cause` of `timeout` will be sent. Minimum value is 5 seconds. Maximum value is 600 seconds.",
			Default:  30,
			BodyPath: "timeout_secs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "webhook-retries-policies",
			Usage:    "A map of event types to retry policies. Each retry policy contains an array of `retries_ms` specifying the delays between retry attempts in milliseconds. Maximum 5 retries, total delay cannot exceed 60 seconds.",
			BodyPath: "webhook_retries_policies",
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
		&requestflag.Flag[map[string]any]{
			Name:     "webhook-urls",
			Usage:    "A map of event types to webhook URLs. When an event of the specified type occurs, the webhook URL associated with that event type will be called instead of `webhook_url`. Events not mapped here will use the default `webhook_url`.",
			BodyPath: "webhook_urls",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-urls-method",
			Usage:    "HTTP request method to invoke `webhook_urls`.",
			Default:  "POST",
			BodyPath: "webhook_urls_method",
		},
	},
	Action:          handleCallsActionsTransfer,
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
})

var callsActionsUpdateClientState = cli.Command{
	Name:    "update-client-state",
	Usage:   "Updates client state",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			Required: true,
			BodyPath: "client_state",
		},
	},
	Action:          handleCallsActionsUpdateClientState,
	HideHelpCommand: true,
}

func handleCallsActionsAddAIAssistantMessages(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionAddAIAssistantMessagesParams{}

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
	_, err = client.Calls.Actions.AddAIAssistantMessages(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions add-ai-assistant-messages", obj, format, transform)
}

func handleCallsActionsAnswer(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionAnswerParams{}

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
	_, err = client.Calls.Actions.Answer(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions answer", obj, format, transform)
}

func handleCallsActionsBridge(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id-to-bridge") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id-to-bridge", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionBridgeParams{}

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
	_, err = client.Calls.Actions.Bridge(
		ctx,
		cmd.Value("call-control-id-to-bridge").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions bridge", obj, format, transform)
}

func handleCallsActionsEnqueue(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionEnqueueParams{}

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
	_, err = client.Calls.Actions.Enqueue(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions enqueue", obj, format, transform)
}

func handleCallsActionsGather(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionGatherParams{}

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
	_, err = client.Calls.Actions.Gather(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions gather", obj, format, transform)
}

func handleCallsActionsGatherUsingAI(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionGatherUsingAIParams{}

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
	_, err = client.Calls.Actions.GatherUsingAI(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions gather-using-ai", obj, format, transform)
}

func handleCallsActionsGatherUsingAudio(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionGatherUsingAudioParams{}

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
	_, err = client.Calls.Actions.GatherUsingAudio(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions gather-using-audio", obj, format, transform)
}

func handleCallsActionsGatherUsingSpeak(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionGatherUsingSpeakParams{}

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
	_, err = client.Calls.Actions.GatherUsingSpeak(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions gather-using-speak", obj, format, transform)
}

func handleCallsActionsHangup(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionHangupParams{}

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
	_, err = client.Calls.Actions.Hangup(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions hangup", obj, format, transform)
}

func handleCallsActionsLeaveQueue(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionLeaveQueueParams{}

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
	_, err = client.Calls.Actions.LeaveQueue(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions leave-queue", obj, format, transform)
}

func handleCallsActionsPauseRecording(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionPauseRecordingParams{}

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
	_, err = client.Calls.Actions.PauseRecording(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions pause-recording", obj, format, transform)
}

func handleCallsActionsRefer(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionReferParams{}

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
	_, err = client.Calls.Actions.Refer(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions refer", obj, format, transform)
}

func handleCallsActionsReject(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionRejectParams{}

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
	_, err = client.Calls.Actions.Reject(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions reject", obj, format, transform)
}

func handleCallsActionsResumeRecording(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionResumeRecordingParams{}

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
	_, err = client.Calls.Actions.ResumeRecording(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions resume-recording", obj, format, transform)
}

func handleCallsActionsSendDtmf(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionSendDtmfParams{}

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
	_, err = client.Calls.Actions.SendDtmf(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions send-dtmf", obj, format, transform)
}

func handleCallsActionsSendSipInfo(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionSendSipInfoParams{}

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
	_, err = client.Calls.Actions.SendSipInfo(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions send-sip-info", obj, format, transform)
}

func handleCallsActionsSpeak(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionSpeakParams{}

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
	_, err = client.Calls.Actions.Speak(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions speak", obj, format, transform)
}

func handleCallsActionsStartAIAssistant(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStartAIAssistantParams{}

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
	_, err = client.Calls.Actions.StartAIAssistant(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions start-ai-assistant", obj, format, transform)
}

func handleCallsActionsStartForking(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStartForkingParams{}

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
	_, err = client.Calls.Actions.StartForking(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions start-forking", obj, format, transform)
}

func handleCallsActionsStartNoiseSuppression(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStartNoiseSuppressionParams{}

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
	_, err = client.Calls.Actions.StartNoiseSuppression(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions start-noise-suppression", obj, format, transform)
}

func handleCallsActionsStartPlayback(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStartPlaybackParams{}

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
	_, err = client.Calls.Actions.StartPlayback(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions start-playback", obj, format, transform)
}

func handleCallsActionsStartRecording(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStartRecordingParams{}

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
	_, err = client.Calls.Actions.StartRecording(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions start-recording", obj, format, transform)
}

func handleCallsActionsStartSiprec(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStartSiprecParams{}

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
	_, err = client.Calls.Actions.StartSiprec(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions start-siprec", obj, format, transform)
}

func handleCallsActionsStartStreaming(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStartStreamingParams{}

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
	_, err = client.Calls.Actions.StartStreaming(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions start-streaming", obj, format, transform)
}

func handleCallsActionsStartTranscription(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStartTranscriptionParams{}

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
	_, err = client.Calls.Actions.StartTranscription(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions start-transcription", obj, format, transform)
}

func handleCallsActionsStopAIAssistant(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopAIAssistantParams{}

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
	_, err = client.Calls.Actions.StopAIAssistant(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-ai-assistant", obj, format, transform)
}

func handleCallsActionsStopForking(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopForkingParams{}

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
	_, err = client.Calls.Actions.StopForking(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-forking", obj, format, transform)
}

func handleCallsActionsStopGather(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopGatherParams{}

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
	_, err = client.Calls.Actions.StopGather(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-gather", obj, format, transform)
}

func handleCallsActionsStopNoiseSuppression(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopNoiseSuppressionParams{}

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
	_, err = client.Calls.Actions.StopNoiseSuppression(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-noise-suppression", obj, format, transform)
}

func handleCallsActionsStopPlayback(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopPlaybackParams{}

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
	_, err = client.Calls.Actions.StopPlayback(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-playback", obj, format, transform)
}

func handleCallsActionsStopRecording(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopRecordingParams{}

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
	_, err = client.Calls.Actions.StopRecording(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-recording", obj, format, transform)
}

func handleCallsActionsStopSiprec(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopSiprecParams{}

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
	_, err = client.Calls.Actions.StopSiprec(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-siprec", obj, format, transform)
}

func handleCallsActionsStopStreaming(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopStreamingParams{}

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
	_, err = client.Calls.Actions.StopStreaming(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-streaming", obj, format, transform)
}

func handleCallsActionsStopTranscription(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionStopTranscriptionParams{}

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
	_, err = client.Calls.Actions.StopTranscription(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions stop-transcription", obj, format, transform)
}

func handleCallsActionsSwitchSupervisorRole(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionSwitchSupervisorRoleParams{}

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
	_, err = client.Calls.Actions.SwitchSupervisorRole(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions switch-supervisor-role", obj, format, transform)
}

func handleCallsActionsTransfer(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionTransferParams{}

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
	_, err = client.Calls.Actions.Transfer(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions transfer", obj, format, transform)
}

func handleCallsActionsUpdateClientState(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallActionUpdateClientStateParams{}

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
	_, err = client.Calls.Actions.UpdateClientState(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "calls:actions update-client-state", obj, format, transform)
}

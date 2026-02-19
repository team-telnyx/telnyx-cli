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

var conferencesActionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update conference participant supervisor_role",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Usage:    "Unique identifier and token for controlling the call",
			Required: true,
			BodyPath: "call_control_id",
		},
		&requestflag.Flag[string]{
			Name:     "supervisor-role",
			Usage:    `Sets the participant as a supervisor for the conference. A conference can have multiple supervisors. "barge" means the supervisor enters the conference as a normal participant. This is the same as "none". "monitor" means the supervisor is muted but can hear all participants. "whisper" means that only the specified "whisper_call_control_ids" can hear the supervisor. Defaults to "none".`,
			Required: true,
			BodyPath: "supervisor_role",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same `command_id` as one that has already been executed.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
		&requestflag.Flag[[]string]{
			Name:     "whisper-call-control-id",
			Usage:    "Array of unique call_control_ids the supervisor can whisper to. If none provided, the supervisor will join the conference as a monitoring participant only.",
			BodyPath: "whisper_call_control_ids",
		},
	},
	Action:          handleConferencesActionsUpdate,
	HideHelpCommand: true,
}

var conferencesActionsEnd = cli.Command{
	Name:    "end",
	Usage:   "End a conference and terminate all active participants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same conference.",
			BodyPath: "command_id",
		},
	},
	Action:          handleConferencesActionsEnd,
	HideHelpCommand: true,
}

var conferencesActionsGatherUsingAudio = cli.Command{
	Name:    "gather-using-audio",
	Usage:   "Play an audio file to a specific conference participant and gather DTMF input.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Usage:    "Unique identifier and token for controlling the call leg that will receive the gather prompt.",
			Required: true,
			BodyPath: "call_control_id",
		},
		&requestflag.Flag[string]{
			Name:     "audio-url",
			Usage:    "The URL of the audio file to play as the gather prompt. Must be WAV or MP3 format.",
			BodyPath: "audio_url",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. Must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "gather-id",
			Usage:    "Identifier for this gather command. Will be included in the gather ended webhook. Maximum 100 characters.",
			BodyPath: "gather_id",
		},
		&requestflag.Flag[int64]{
			Name:     "initial-timeout-millis",
			Usage:    "Duration in milliseconds to wait for the first digit before timing out.",
			BodyPath: "initial_timeout_millis",
		},
		&requestflag.Flag[int64]{
			Name:     "inter-digit-timeout-millis",
			Usage:    "Duration in milliseconds to wait between digits.",
			Default:  5000,
			BodyPath: "inter_digit_timeout_millis",
		},
		&requestflag.Flag[string]{
			Name:     "invalid-audio-url",
			Usage:    "URL of audio file to play when invalid input is received.",
			BodyPath: "invalid_audio_url",
		},
		&requestflag.Flag[string]{
			Name:     "invalid-media-name",
			Usage:    "Name of media file to play when invalid input is received.",
			BodyPath: "invalid_media_name",
		},
		&requestflag.Flag[int64]{
			Name:     "maximum-digits",
			Usage:    "Maximum number of digits to gather.",
			Default:  128,
			BodyPath: "maximum_digits",
		},
		&requestflag.Flag[int64]{
			Name:     "maximum-tries",
			Usage:    "Maximum number of times to play the prompt if no input is received.",
			Default:  3,
			BodyPath: "maximum_tries",
		},
		&requestflag.Flag[string]{
			Name:     "media-name",
			Usage:    "The name of the media file uploaded to the Media Storage API to play as the gather prompt.",
			BodyPath: "media_name",
		},
		&requestflag.Flag[int64]{
			Name:     "minimum-digits",
			Usage:    "Minimum number of digits to gather.",
			Default:  1,
			BodyPath: "minimum_digits",
		},
		&requestflag.Flag[bool]{
			Name:     "stop-playback-on-dtmf",
			Usage:    "Whether to stop the audio playback when a DTMF digit is received.",
			Default:  true,
			BodyPath: "stop_playback_on_dtmf",
		},
		&requestflag.Flag[string]{
			Name:     "terminating-digit",
			Usage:    "Digit that terminates gathering.",
			Default:  "#",
			BodyPath: "terminating_digit",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-millis",
			Usage:    "Duration in milliseconds to wait for input before timing out.",
			Default:  60000,
			BodyPath: "timeout_millis",
		},
		&requestflag.Flag[string]{
			Name:     "valid-digits",
			Usage:    "Digits that are valid for gathering. All other digits will be ignored.",
			Default:  "0123456789#*",
			BodyPath: "valid_digits",
		},
	},
	Action:          handleConferencesActionsGatherUsingAudio,
	HideHelpCommand: true,
}

var conferencesActionsHold = cli.Command{
	Name:    "hold",
	Usage:   "Hold a list of participants in a conference call",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "audio-url",
			Usage:    "The URL of a file to be played to the participants when they are put on hold. media_name and audio_url cannot be used together in one request.",
			BodyPath: "audio_url",
		},
		&requestflag.Flag[[]string]{
			Name:     "call-control-id",
			Usage:    "List of unique identifiers and tokens for controlling the call. When empty all participants will be placed on hold.",
			BodyPath: "call_control_ids",
		},
		&requestflag.Flag[string]{
			Name:     "media-name",
			Usage:    "The media_name of a file to be played to the participants when they are put on hold. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file.",
			BodyPath: "media_name",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsHold,
	HideHelpCommand: true,
}

var conferencesActionsJoin = cli.Command{
	Name:    "join",
	Usage:   "Join an existing call leg to a conference. Issue the Join Conference command\nwith the conference ID in the path and the `call_control_id` of the leg you wish\nto join to the conference as an attribute. The conference can have up to a\ncertain amount of active participants, as set by the `max_participants`\nparameter in conference creation request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Usage:    "Unique identifier and token for controlling the call",
			Required: true,
			BodyPath: "call_control_id",
		},
		&requestflag.Flag[string]{
			Name:     "beep-enabled",
			Usage:    "Whether a beep sound should be played when the participant joins and/or leaves the conference. Can be used to override the conference-level setting.",
			BodyPath: "beep_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. Please note that the client_state will be updated for the participient call leg and the change will not affect conferencing webhooks unless the participient is the owner of the conference.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same `command_id` as one that has already been executed.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[bool]{
			Name:     "end-conference-on-exit",
			Usage:    `Whether the conference should end and all remaining participants be hung up after the participant leaves the conference. Defaults to "false".`,
			BodyPath: "end_conference_on_exit",
		},
		&requestflag.Flag[bool]{
			Name:     "hold",
			Usage:    `Whether the participant should be put on hold immediately after joining the conference. Defaults to "false".`,
			BodyPath: "hold",
		},
		&requestflag.Flag[string]{
			Name:     "hold-audio-url",
			Usage:    `The URL of a file to be played to the participant when they are put on hold after joining the conference. hold_media_name and hold_audio_url cannot be used together in one request. Takes effect only when "start_conference_on_create" is set to "false". This property takes effect only if "hold" is set to "true".`,
			BodyPath: "hold_audio_url",
		},
		&requestflag.Flag[string]{
			Name:     "hold-media-name",
			Usage:    `The media_name of a file to be played to the participant when they are put on hold after joining the conference. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. Takes effect only when "start_conference_on_create" is set to "false". This property takes effect only if "hold" is set to "true".`,
			BodyPath: "hold_media_name",
		},
		&requestflag.Flag[bool]{
			Name:     "mute",
			Usage:    `Whether the participant should be muted immediately after joining the conference. Defaults to "false".`,
			BodyPath: "mute",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
		&requestflag.Flag[bool]{
			Name:     "soft-end-conference-on-exit",
			Usage:    `Whether the conference should end after the participant leaves the conference. NOTE this doesn't hang up the other participants. Defaults to "false".`,
			BodyPath: "soft_end_conference_on_exit",
		},
		&requestflag.Flag[bool]{
			Name:     "start-conference-on-enter",
			Usage:    `Whether the conference should be started after the participant joins the conference. Defaults to "false".`,
			BodyPath: "start_conference_on_enter",
		},
		&requestflag.Flag[string]{
			Name:     "supervisor-role",
			Usage:    `Sets the joining participant as a supervisor for the conference. A conference can have multiple supervisors. "barge" means the supervisor enters the conference as a normal participant. This is the same as "none". "monitor" means the supervisor is muted but can hear all participants. "whisper" means that only the specified "whisper_call_control_ids" can hear the supervisor. Defaults to "none".`,
			BodyPath: "supervisor_role",
		},
		&requestflag.Flag[[]string]{
			Name:     "whisper-call-control-id",
			Usage:    "Array of unique call_control_ids the joining supervisor can whisper to. If none provided, the supervisor will join the conference as a monitoring participant only.",
			BodyPath: "whisper_call_control_ids",
		},
	},
	Action:          handleConferencesActionsJoin,
	HideHelpCommand: true,
}

var conferencesActionsLeave = cli.Command{
	Name:    "leave",
	Usage:   "Removes a call leg from a conference and moves it back to parked state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Usage:    "Unique identifier and token for controlling the call",
			Required: true,
			BodyPath: "call_control_id",
		},
		&requestflag.Flag[string]{
			Name:     "beep-enabled",
			Usage:    "Whether a beep sound should be played when the participant leaves the conference. Can be used to override the conference-level setting.",
			BodyPath: "beep_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same `command_id` as one that has already been executed.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsLeave,
	HideHelpCommand: true,
}

var conferencesActionsMute = cli.Command{
	Name:    "mute",
	Usage:   "Mute a list of participants in a conference call",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "call-control-id",
			Usage:    "Array of unique identifiers and tokens for controlling the call. When empty all participants will be muted.",
			BodyPath: "call_control_ids",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsMute,
	HideHelpCommand: true,
}

var conferencesActionsPlay = cli.Command{
	Name:    "play",
	Usage:   "Play audio to all or some participants on a conference call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "audio-url",
			Usage:    "The URL of a file to be played back in the conference. media_name and audio_url cannot be used together in one request.",
			BodyPath: "audio_url",
		},
		&requestflag.Flag[[]string]{
			Name:     "call-control-id",
			Usage:    "List of call control ids identifying participants the audio file should be played to. If not given, the audio file will be played to the entire conference.",
			BodyPath: "call_control_ids",
		},
		&requestflag.Flag[any]{
			Name:     "loop",
			BodyPath: "loop",
		},
		&requestflag.Flag[string]{
			Name:     "media-name",
			Usage:    "The media_name of a file to be played back in the conference. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file.",
			BodyPath: "media_name",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsPlay,
	HideHelpCommand: true,
}

var conferencesActionsRecordPause = cli.Command{
	Name:    "record-pause",
	Usage:   "Pause conference recording.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "recording-id",
			Usage:    "Use this field to pause specific recording.",
			BodyPath: "recording_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsRecordPause,
	HideHelpCommand: true,
}

var conferencesActionsRecordResume = cli.Command{
	Name:    "record-resume",
	Usage:   "Resume conference recording.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "recording-id",
			Usage:    "Use this field to resume specific recording.",
			BodyPath: "recording_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsRecordResume,
	HideHelpCommand: true,
}

var conferencesActionsRecordStart = cli.Command{
	Name:    "record-start",
	Usage:   "Start recording the conference. Recording will stop on conference end, or via\nthe Stop Recording command.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "The audio file format used when storing the conference recording. Can be either `mp3` or `wav`.",
			Required: true,
			BodyPath: "format",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `conference_id`.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "custom-file-name",
			Usage:    "The custom recording file name to be used instead of the default `call_leg_id`. Telnyx will still add a Unix timestamp suffix.",
			BodyPath: "custom_file_name",
		},
		&requestflag.Flag[bool]{
			Name:     "play-beep",
			Usage:    "If enabled, a beep sound will be played at the start of a recording.",
			BodyPath: "play_beep",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
		&requestflag.Flag[string]{
			Name:     "trim",
			Usage:    "When set to `trim-silence`, silence will be removed from the beginning and end of the recording.",
			BodyPath: "trim",
		},
	},
	Action:          handleConferencesActionsRecordStart,
	HideHelpCommand: true,
}

var conferencesActionsRecordStop = cli.Command{
	Name:    "record-stop",
	Usage:   "Stop recording the conference.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
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
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsRecordStop,
	HideHelpCommand: true,
}

var conferencesActionsSendDtmf = cli.Command{
	Name:    "send-dtmf",
	Usage:   "Send DTMF tones to one or more conference participants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "digits",
			Usage:    "DTMF digits to send. Valid characters: 0-9, A-D, *, #, w (0.5s pause), W (1s pause).",
			Required: true,
			BodyPath: "digits",
		},
		&requestflag.Flag[[]string]{
			Name:     "call-control-id",
			Usage:    "Array of participant call control IDs to send DTMF to. When empty, DTMF will be sent to all participants.",
			BodyPath: "call_control_ids",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. Must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[int64]{
			Name:     "duration-millis",
			Usage:    "Duration of each DTMF digit in milliseconds.",
			Default:  250,
			BodyPath: "duration_millis",
		},
	},
	Action:          handleConferencesActionsSendDtmf,
	HideHelpCommand: true,
}

var conferencesActionsSpeak = cli.Command{
	Name:    "speak",
	Usage:   "Convert text to speech and play it to all or some participants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
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
			Usage:    "Specifies the voice used in speech synthesis.\n\n- Define voices using the format `<Provider>.<Model>.<VoiceId>`. Specifying only the provider will give default values for voice_id and model_id.\n\n **Supported Providers:**\n- **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural voices, which provide more realistic, human-like speech, append `-Neural` to the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility.\n- **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural, Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural, Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)\n- **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g., `ElevenLabs.eleven_multilingual_v2.21m00Tcm4TlvDq8ikWAM`). The `ModelId` part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration identifier secret in `\"voice_settings\": {\"api_key_ref\": \"<secret_identifier>\"}`.  Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).\n- **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`\n- **Minimax:** Use `Minimax.<ModelId>.<VoiceId>` (e.g., `Minimax.speech-02-hd.Wise_Woman`). Supported models: `speech-02-turbo`, `speech-02-hd`, `speech-2.6-turbo`, `speech-2.8-turbo`. Optional parameters: `speed` (float, default 1.0), `vol` (float, default 1.0), `pitch` (integer, default 0). \n- **Resemble:** Use `Resemble.<ModelId>.<VoiceId>` (e.g., `Resemble.Pro.my_voice`). Supported models: `Pro` (multilingual) and `Turbo` (English only).",
			Required: true,
			BodyPath: "voice",
		},
		&requestflag.Flag[[]string]{
			Name:     "call-control-id",
			Usage:    "Call Control IDs of participants who will hear the spoken text. When empty all participants will hear the spoken text.",
			BodyPath: "call_control_ids",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same `command_id` as one that has already been executed.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "The language you want spoken. This parameter is ignored when a `Polly.*` voice is specified.",
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "payload-type",
			Usage:    "The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML).",
			Default:  "text",
			BodyPath: "payload_type",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
		&requestflag.Flag[any]{
			Name:     "voice-settings",
			Usage:    "The settings associated with the voice selected",
			BodyPath: "voice_settings",
		},
	},
	Action:          handleConferencesActionsSpeak,
	HideHelpCommand: true,
}

var conferencesActionsStop = cli.Command{
	Name:    "stop",
	Usage:   "Stop audio being played to all or some participants on a conference call.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "call-control-id",
			Usage:    "List of call control ids identifying participants the audio file should stop be played to. If not given, the audio will be stoped to the entire conference.",
			BodyPath: "call_control_ids",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsStop,
	HideHelpCommand: true,
}

var conferencesActionsUnhold = cli.Command{
	Name:    "unhold",
	Usage:   "Unhold a list of participants in a conference call",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "call-control-id",
			Usage:    "List of unique identifiers and tokens for controlling the call. Enter each call control ID to be unheld.",
			Required: true,
			BodyPath: "call_control_ids",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsUnhold,
	HideHelpCommand: true,
}

var conferencesActionsUnmute = cli.Command{
	Name:    "unmute",
	Usage:   "Unmute a list of participants in a conference call",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "call-control-id",
			Usage:    "List of unique identifiers and tokens for controlling the call. Enter each call control ID to be unmuted. When empty all participants will be unmuted.",
			BodyPath: "call_control_ids",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the conference data is located. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
	},
	Action:          handleConferencesActionsUnmute,
	HideHelpCommand: true,
}

func handleConferencesActionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionUpdateParams{}

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
	_, err = client.Conferences.Actions.Update(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions update", obj, format, transform)
}

func handleConferencesActionsEnd(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionEndParams{}

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
	_, err = client.Conferences.Actions.End(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions end", obj, format, transform)
}

func handleConferencesActionsGatherUsingAudio(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionGatherUsingAudioParams{}

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
	_, err = client.Conferences.Actions.GatherUsingAudio(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions gather-using-audio", obj, format, transform)
}

func handleConferencesActionsHold(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionHoldParams{}

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
	_, err = client.Conferences.Actions.Hold(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions hold", obj, format, transform)
}

func handleConferencesActionsJoin(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionJoinParams{}

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
	_, err = client.Conferences.Actions.Join(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions join", obj, format, transform)
}

func handleConferencesActionsLeave(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionLeaveParams{}

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
	_, err = client.Conferences.Actions.Leave(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions leave", obj, format, transform)
}

func handleConferencesActionsMute(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionMuteParams{}

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
	_, err = client.Conferences.Actions.Mute(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions mute", obj, format, transform)
}

func handleConferencesActionsPlay(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionPlayParams{}

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
	_, err = client.Conferences.Actions.Play(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions play", obj, format, transform)
}

func handleConferencesActionsRecordPause(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionRecordPauseParams{}

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
	_, err = client.Conferences.Actions.RecordPause(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions record-pause", obj, format, transform)
}

func handleConferencesActionsRecordResume(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionRecordResumeParams{}

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
	_, err = client.Conferences.Actions.RecordResume(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions record-resume", obj, format, transform)
}

func handleConferencesActionsRecordStart(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionRecordStartParams{}

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
	_, err = client.Conferences.Actions.RecordStart(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions record-start", obj, format, transform)
}

func handleConferencesActionsRecordStop(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionRecordStopParams{}

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
	_, err = client.Conferences.Actions.RecordStop(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions record-stop", obj, format, transform)
}

func handleConferencesActionsSendDtmf(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionSendDtmfParams{}

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
	_, err = client.Conferences.Actions.SendDtmf(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions send-dtmf", obj, format, transform)
}

func handleConferencesActionsSpeak(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionSpeakParams{}

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
	_, err = client.Conferences.Actions.Speak(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions speak", obj, format, transform)
}

func handleConferencesActionsStop(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionStopParams{}

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
	_, err = client.Conferences.Actions.Stop(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions stop", obj, format, transform)
}

func handleConferencesActionsUnhold(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionUnholdParams{}

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
	_, err = client.Conferences.Actions.Unhold(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions unhold", obj, format, transform)
}

func handleConferencesActionsUnmute(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceActionUnmuteParams{}

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
	_, err = client.Conferences.Actions.Unmute(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences:actions unmute", obj, format, transform)
}

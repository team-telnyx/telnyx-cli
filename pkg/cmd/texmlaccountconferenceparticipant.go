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

var texmlAccountsConferencesParticipantsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Gets conference participant resource",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "conference-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid-or-participant-label",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsConferencesParticipantsRetrieve,
	HideHelpCommand: true,
}

var texmlAccountsConferencesParticipantsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a conference participant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "conference-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid-or-participant-label",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "announce-method",
			Usage:    "The HTTP method used to call the `AnnounceUrl`. Defaults to `POST`.",
			BodyPath: "AnnounceMethod",
		},
		&requestflag.Flag[string]{
			Name:     "announce-url",
			Usage:    "The URL to call to announce something to the participant. The URL may return an MP3 fileo a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs.",
			BodyPath: "AnnounceUrl",
		},
		&requestflag.Flag[bool]{
			Name:     "beep-on-exit",
			Usage:    "Whether to play a notification beep to the conference when the participant exits.",
			BodyPath: "BeepOnExit",
		},
		&requestflag.Flag[string]{
			Name:     "call-sid-to-coach",
			Usage:    "The SID of the participant who is being coached. The participant being coached is the only participant who can hear the participant who is coaching.",
			BodyPath: "CallSidToCoach",
		},
		&requestflag.Flag[bool]{
			Name:     "coaching",
			Usage:    "Whether the participant is coaching another call. When `true`, `CallSidToCoach` has to be given.",
			BodyPath: "Coaching",
		},
		&requestflag.Flag[bool]{
			Name:     "end-conference-on-exit",
			Usage:    "Whether to end the conference when the participant leaves.",
			BodyPath: "EndConferenceOnExit",
		},
		&requestflag.Flag[bool]{
			Name:     "hold",
			Usage:    "Whether the participant should be on hold.",
			BodyPath: "Hold",
		},
		&requestflag.Flag[string]{
			Name:     "hold-method",
			Usage:    "The HTTP method to use when calling the `HoldUrl`.",
			BodyPath: "HoldMethod",
		},
		&requestflag.Flag[string]{
			Name:     "hold-url",
			Usage:    "The URL to be called using the `HoldMethod` for music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs.",
			BodyPath: "HoldUrl",
		},
		&requestflag.Flag[bool]{
			Name:     "muted",
			Usage:    "Whether the participant should be muted.",
			BodyPath: "Muted",
		},
		&requestflag.Flag[string]{
			Name:     "wait-url",
			Usage:    "The URL to call for an audio file to play while the participant is waiting for the conference to start.",
			BodyPath: "WaitUrl",
		},
	},
	Action:          handleTexmlAccountsConferencesParticipantsUpdate,
	HideHelpCommand: true,
}

var texmlAccountsConferencesParticipantsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a conference participant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "conference-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid-or-participant-label",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsConferencesParticipantsDelete,
	HideHelpCommand: true,
}

var texmlAccountsConferencesParticipantsParticipants = requestflag.WithInnerFlags(cli.Command{
	Name:    "participants",
	Usage:   "Dials a new conference participant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "conference-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "amd-status-callback",
			Usage:    "The URL the result of answering machine detection will be sent to.",
			BodyPath: "AmdStatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "amd-status-callback-method",
			Usage:    "HTTP request type used for `AmdStatusCallback`. Defaults to `POST`.",
			BodyPath: "AmdStatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "beep",
			Usage:    "Whether to play a notification beep to the conference when the participant enters and exits.",
			BodyPath: "Beep",
		},
		&requestflag.Flag[string]{
			Name:     "caller-id",
			Usage:    "To be used as the caller id name (SIP From Display Name) presented to the destination (`To` number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and `-_~!.+` special characters. If ommited, the display name will be the same as the number in the `From` field.",
			BodyPath: "CallerId",
		},
		&requestflag.Flag[string]{
			Name:     "call-sid-to-coach",
			Usage:    "The SID of the participant who is being coached. The participant being coached is the only participant who can hear the participant who is coaching.",
			BodyPath: "CallSidToCoach",
		},
		&requestflag.Flag[bool]{
			Name:     "cancel-playback-on-detect-message-end",
			Usage:    "Whether to cancel ongoing playback on `greeting ended` detection. Defaults to `true`.",
			Default:  true,
			BodyPath: "CancelPlaybackOnDetectMessageEnd",
		},
		&requestflag.Flag[bool]{
			Name:     "cancel-playback-on-machine-detection",
			Usage:    "Whether to cancel ongoing playback on `machine` detection. Defaults to `true`.",
			Default:  true,
			BodyPath: "CancelPlaybackOnMachineDetection",
		},
		&requestflag.Flag[bool]{
			Name:     "coaching",
			Usage:    "Whether the participant is coaching another call. When `true`, `CallSidToCoach` has to be given.",
			BodyPath: "Coaching",
		},
		&requestflag.Flag[string]{
			Name:     "conference-record",
			Usage:    "Whether to record the conference the participant is joining. Defualts to `do-not-record`. The boolean values `true` and `false` are synonymous with `record-from-start` and `do-not-record` respectively.",
			BodyPath: "ConferenceRecord",
		},
		&requestflag.Flag[string]{
			Name:     "conference-recording-status-callback",
			Usage:    "The URL the conference recording callbacks will be sent to.",
			BodyPath: "ConferenceRecordingStatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "conference-recording-status-callback-event",
			Usage:    "The changes to the conference recording's state that should generate a call to `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`. Separate multiple values with a space. Defaults to `completed`. `failed` and `absent` are synonymous.",
			BodyPath: "ConferenceRecordingStatusCallbackEvent",
		},
		&requestflag.Flag[string]{
			Name:     "conference-recording-status-callback-method",
			Usage:    "HTTP request type used for `ConferenceRecordingStatusCallback`. Defaults to `POST`.",
			BodyPath: "ConferenceRecordingStatusCallbackMethod",
		},
		&requestflag.Flag[int64]{
			Name:     "conference-recording-timeout",
			Usage:    "The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected. The timer only starts when the speech is detected. Please note that the transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite)",
			Default:  0,
			BodyPath: "ConferenceRecordingTimeout",
		},
		&requestflag.Flag[string]{
			Name:     "conference-status-callback",
			Usage:    "The URL the conference callbacks will be sent to.",
			BodyPath: "ConferenceStatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "conference-status-callback-event",
			Usage:    "The changes to the conference's state that should generate a call to `ConferenceStatusCallback`. Can be: `start`, `end`, `join` and `leave`. Separate multiple values with a space. By default no callbacks are sent.",
			BodyPath: "ConferenceStatusCallbackEvent",
		},
		&requestflag.Flag[string]{
			Name:     "conference-status-callback-method",
			Usage:    "HTTP request type used for `ConferenceStatusCallback`. Defaults to `POST`.",
			BodyPath: "ConferenceStatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "conference-trim",
			Usage:    "Whether to trim any leading and trailing silence from the conference recording. Defaults to `trim-silence`.",
			BodyPath: "ConferenceTrim",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-header",
			Usage:    "Custom HTTP headers to be sent with the call. Each header should be an object with 'name' and 'value' properties.",
			BodyPath: "CustomHeaders",
		},
		&requestflag.Flag[bool]{
			Name:     "early-media",
			Usage:    "Whether participant shall be bridged to conference before the participant answers (from early media if available). Defaults to `false`.",
			Default:  false,
			BodyPath: "EarlyMedia",
		},
		&requestflag.Flag[bool]{
			Name:     "end-conference-on-exit",
			Usage:    "Whether to end the conference when the participant leaves. Defaults to `false`.",
			BodyPath: "EndConferenceOnExit",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "The phone number of the party that initiated the call. Phone numbers are formatted with a `+` and country code.",
			BodyPath: "From",
		},
		&requestflag.Flag[string]{
			Name:     "label",
			Usage:    "A unique label for the participant that will be added to the conference. The label can be used to reference the participant for updates via the TeXML REST API.",
			BodyPath: "Label",
		},
		&requestflag.Flag[string]{
			Name:     "machine-detection",
			Usage:    "Whether to detect if a human or an answering machine picked up the call. Use `Enable` if you would like to ne notified as soon as the called party is identified. Use `DetectMessageEnd`, if you would like to leave a message on an answering machine.",
			BodyPath: "MachineDetection",
		},
		&requestflag.Flag[int64]{
			Name:     "machine-detection-silence-timeout",
			Usage:    "If initial silence duration is greater than this value, consider it a machine. Ignored when `premium` detection is used.",
			Default:  3500,
			BodyPath: "MachineDetectionSilenceTimeout",
		},
		&requestflag.Flag[int64]{
			Name:     "machine-detection-speech-end-threshold",
			Usage:    "Silence duration threshold after a greeting message or voice for it be considered human. Ignored when `premium` detection is used.",
			Default:  800,
			BodyPath: "MachineDetectionSpeechEndThreshold",
		},
		&requestflag.Flag[int64]{
			Name:     "machine-detection-speech-threshold",
			Usage:    "Maximum threshold of a human greeting. If greeting longer than this value, considered machine. Ignored when `premium` detection is used.",
			Default:  3500,
			BodyPath: "MachineDetectionSpeechThreshold",
		},
		&requestflag.Flag[int64]{
			Name:     "machine-detection-timeout",
			Usage:    "How long answering machine detection should go on for before sending an `Unknown` result. Given in milliseconds.",
			BodyPath: "MachineDetectionTimeout",
		},
		&requestflag.Flag[int64]{
			Name:     "max-participants",
			Usage:    "The maximum number of participants in the conference. Can be a positive integer from 2 to 800. The default value is 250.",
			BodyPath: "MaxParticipants",
		},
		&requestflag.Flag[bool]{
			Name:     "muted",
			Usage:    "Whether the participant should be muted.",
			BodyPath: "Muted",
		},
		&requestflag.Flag[string]{
			Name:     "preferred-codecs",
			Usage:    "The list of comma-separated codecs to be offered on a call.",
			BodyPath: "PreferredCodecs",
		},
		&requestflag.Flag[bool]{
			Name:     "record",
			Usage:    "Whether to record the entire participant's call leg. Defaults to `false`.",
			BodyPath: "Record",
		},
		&requestflag.Flag[string]{
			Name:     "recording-channels",
			Usage:    "The number of channels in the final recording. Defaults to `mono`.",
			BodyPath: "RecordingChannels",
		},
		&requestflag.Flag[string]{
			Name:     "recording-status-callback",
			Usage:    "The URL the recording callbacks will be sent to.",
			BodyPath: "RecordingStatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "recording-status-callback-event",
			Usage:    "The changes to the recording's state that should generate a call to `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`. Separate multiple values with a space. Defaults to `completed`.",
			BodyPath: "RecordingStatusCallbackEvent",
		},
		&requestflag.Flag[string]{
			Name:     "recording-status-callback-method",
			Usage:    "HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.",
			BodyPath: "RecordingStatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "recording-track",
			Usage:    "The audio track to record for the call. The default is `both`.",
			BodyPath: "RecordingTrack",
		},
		&requestflag.Flag[string]{
			Name:     "sip-auth-password",
			Usage:    "The password to use for SIP authentication.",
			BodyPath: "SipAuthPassword",
		},
		&requestflag.Flag[string]{
			Name:     "sip-auth-username",
			Usage:    "The username to use for SIP authentication.",
			BodyPath: "SipAuthUsername",
		},
		&requestflag.Flag[bool]{
			Name:     "start-conference-on-enter",
			Usage:    "Whether to start the conference when the participant enters. Defaults to `true`.",
			BodyPath: "StartConferenceOnEnter",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "URL destination for Telnyx to send status callback events to for the call.",
			BodyPath: "StatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-event",
			Usage:    "The changes to the call's state that should generate a call to `StatusCallback`. Can be: `initiated`, `ringing`, `answered`, and `completed`. Separate multiple values with a space. The default value is `completed`.",
			BodyPath: "StatusCallbackEvent",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request type used for `StatusCallback`.",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[int64]{
			Name:     "time-limit",
			Usage:    "The maximum duration of the call in seconds.",
			BodyPath: "TimeLimit",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-seconds",
			Usage:    "The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between 5 and 120, inclusive. The default value is 30.",
			BodyPath: "Timeout",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "The phone number of the called party. Phone numbers are formatted with a `+` and country code.",
			BodyPath: "To",
		},
		&requestflag.Flag[string]{
			Name:     "trim",
			Usage:    "Whether to trim any leading and trailing silence from the recording. Defaults to `trim-silence`.",
			BodyPath: "Trim",
		},
		&requestflag.Flag[string]{
			Name:     "wait-url",
			Usage:    "The URL to call for an audio file to play while the participant is waiting for the conference to start.",
			BodyPath: "WaitUrl",
		},
	},
	Action:          handleTexmlAccountsConferencesParticipantsParticipants,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"custom-header": {
		&requestflag.InnerFlag[string]{
			Name:       "custom-header.name",
			Usage:      "The name of the custom header",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "custom-header.value",
			Usage:      "The value of the custom header",
			InnerField: "value",
		},
	},
})

var texmlAccountsConferencesParticipantsRetrieveParticipants = cli.Command{
	Name:    "retrieve-participants",
	Usage:   "Lists conference participants",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "conference-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsConferencesParticipantsRetrieveParticipants,
	HideHelpCommand: true,
}

func handleTexmlAccountsConferencesParticipantsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid-or-participant-label") && len(unusedArgs) > 0 {
		cmd.Set("call-sid-or-participant-label", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceParticipantGetParams{
		AccountSid:    cmd.Value("account-sid").(string),
		ConferenceSid: cmd.Value("conference-sid").(string),
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
	_, err = client.Texml.Accounts.Conferences.Participants.Get(
		ctx,
		cmd.Value("call-sid-or-participant-label").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:conferences:participants retrieve", obj, format, transform)
}

func handleTexmlAccountsConferencesParticipantsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid-or-participant-label") && len(unusedArgs) > 0 {
		cmd.Set("call-sid-or-participant-label", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceParticipantUpdateParams{
		AccountSid:    cmd.Value("account-sid").(string),
		ConferenceSid: cmd.Value("conference-sid").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Texml.Accounts.Conferences.Participants.Update(
		ctx,
		cmd.Value("call-sid-or-participant-label").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:conferences:participants update", obj, format, transform)
}

func handleTexmlAccountsConferencesParticipantsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid-or-participant-label") && len(unusedArgs) > 0 {
		cmd.Set("call-sid-or-participant-label", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceParticipantDeleteParams{
		AccountSid:    cmd.Value("account-sid").(string),
		ConferenceSid: cmd.Value("conference-sid").(string),
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

	return client.Texml.Accounts.Conferences.Participants.Delete(
		ctx,
		cmd.Value("call-sid-or-participant-label").(string),
		params,
		options...,
	)
}

func handleTexmlAccountsConferencesParticipantsParticipants(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conference-sid") && len(unusedArgs) > 0 {
		cmd.Set("conference-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceParticipantParticipantsParams{
		AccountSid: cmd.Value("account-sid").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Texml.Accounts.Conferences.Participants.Participants(
		ctx,
		cmd.Value("conference-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:conferences:participants participants", obj, format, transform)
}

func handleTexmlAccountsConferencesParticipantsRetrieveParticipants(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conference-sid") && len(unusedArgs) > 0 {
		cmd.Set("conference-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceParticipantGetParticipantsParams{
		AccountSid: cmd.Value("account-sid").(string),
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
	_, err = client.Texml.Accounts.Conferences.Participants.GetParticipants(
		ctx,
		cmd.Value("conference-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:conferences:participants retrieve-participants", obj, format, transform)
}

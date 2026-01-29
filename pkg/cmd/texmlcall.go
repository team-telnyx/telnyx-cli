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

var texmlCallsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update TeXML call. Please note that the keys present in the payload MUST BE\nformatted in CamelCase as specified in the example.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "fallback-method",
			Usage:    "HTTP request type used for `FallbackUrl`.",
			BodyPath: "FallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "fallback-url",
			Usage:    "A failover URL for which Telnyx will retrieve the TeXML call instructions if the Url is not responding.",
			BodyPath: "FallbackUrl",
		},
		&requestflag.Flag[string]{
			Name:     "method",
			Usage:    "HTTP request type used for `Url`.",
			BodyPath: "Method",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The value to set the call status to. Setting the status to completed ends the call.",
			BodyPath: "Status",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "URL destination for Telnyx to send status callback events to for the call.",
			BodyPath: "StatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request type used for `StatusCallback`.",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "texml",
			Usage:    "TeXML to replace the current one with.",
			BodyPath: "Texml",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL where TeXML will make a request to retrieve a new set of TeXML instructions to continue the call flow.",
			BodyPath: "Url",
		},
	},
	Action:          handleTexmlCallsUpdate,
	HideHelpCommand: true,
}

var texmlCallsInitiate = cli.Command{
	Name:    "initiate",
	Usage:   "Initiate an outbound TeXML call. Telnyx will request TeXML from the XML Request\nURL configured for the connection in the Mission Control Portal.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "application-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "The phone number of the party that initiated the call. Phone numbers are formatted with a `+` and country code.",
			Required: true,
			BodyPath: "From",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "The phone number of the called party. Phone numbers are formatted with a `+` and country code.",
			Required: true,
			BodyPath: "To",
		},
		&requestflag.Flag[bool]{
			Name:     "async-amd",
			Usage:    "Select whether to perform answering machine detection in the background. By default execution is blocked until Answering Machine Detection is completed.",
			Default:  false,
			BodyPath: "AsyncAmd",
		},
		&requestflag.Flag[string]{
			Name:     "async-amd-status-callback",
			Usage:    "URL destination for Telnyx to send AMD callback events to for the call.",
			BodyPath: "AsyncAmdStatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "async-amd-status-callback-method",
			Usage:    "HTTP request type used for `AsyncAmdStatusCallback`. The default value is inherited from TeXML Application setting.",
			Default:  "POST",
			BodyPath: "AsyncAmdStatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "caller-id",
			Usage:    "To be used as the caller id name (SIP From Display Name) presented to the destination (`To` number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and `-_~!.+` special characters. If ommited, the display name will be the same as the number in the `From` field.",
			BodyPath: "CallerId",
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
		&requestflag.Flag[string]{
			Name:     "detection-mode",
			Usage:    "Allows you to chose between Premium and Standard detections.",
			Default:  "Regular",
			BodyPath: "DetectionMode",
		},
		&requestflag.Flag[string]{
			Name:     "fallback-url",
			Usage:    "A failover URL for which Telnyx will retrieve the TeXML call instructions if the `Url` is not responding.",
			BodyPath: "FallbackUrl",
		},
		&requestflag.Flag[string]{
			Name:     "machine-detection",
			Usage:    "Enables Answering Machine Detection.",
			Default:  "Disable",
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
			Usage:    "Maximum timeout threshold in milliseconds for overall detection.",
			Default:  30000,
			BodyPath: "MachineDetectionTimeout",
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
		&requestflag.Flag[int64]{
			Name:     "recording-timeout",
			Usage:    "The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected. The timer only starts when the speech is detected. Please note that the transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite)",
			Default:  0,
			BodyPath: "RecordingTimeout",
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
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "URL destination for Telnyx to send status callback events to for the call.",
			BodyPath: "StatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-event",
			Usage:    "The call events for which Telnyx should send a webhook. Multiple events can be defined when separated by a space.",
			Default:  "completed",
			BodyPath: "StatusCallbackEvent",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request type used for `StatusCallback`.",
			Default:  "POST",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "trim",
			Usage:    "Whether to trim any leading and trailing silence from the recording. Defaults to `trim-silence`.",
			BodyPath: "Trim",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL from which Telnyx will retrieve the TeXML call instructions.",
			BodyPath: "Url",
		},
		&requestflag.Flag[string]{
			Name:     "url-method",
			Usage:    "HTTP request type used for `Url`. The default value is inherited from TeXML Application setting.",
			Default:  "POST",
			BodyPath: "UrlMethod",
		},
	},
	Action:          handleTexmlCallsInitiate,
	HideHelpCommand: true,
}

func handleTexmlCallsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlCallUpdateParams{}

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
	_, err = client.Texml.Calls.Update(
		ctx,
		cmd.Value("call-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:calls update", obj, format, transform)
}

func handleTexmlCallsInitiate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("application-id") && len(unusedArgs) > 0 {
		cmd.Set("application-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlCallInitiateParams{}

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
	_, err = client.Texml.Calls.Initiate(
		ctx,
		cmd.Value("application-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:calls initiate", obj, format, transform)
}

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

var texmlAccountsCallsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns an individual call identified by its CallSid. This endpoint is\neventually consistent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsCallsRetrieve,
	HideHelpCommand: true,
}

var texmlAccountsCallsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update TeXML call. Please note that the keys present in the payload MUST BE\nformatted in CamelCase as specified in the example.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
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
	Action:          handleTexmlAccountsCallsUpdate,
	HideHelpCommand: true,
}

var texmlAccountsCallsCalls = requestflag.WithInnerFlags(cli.Command{
	Name:    "calls",
	Usage:   "Initiate an outbound TeXML call. Telnyx will request TeXML from the XML Request\nURL configured for the connection in the Mission Control Portal.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "application-sid",
			Usage:    "The ID of the TeXML Application.",
			Required: true,
			BodyPath: "ApplicationSid",
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
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-header",
			Usage:    "Custom HTTP headers to be sent with the call. Each header should be an object with 'name' and 'value' properties.",
			BodyPath: "CustomHeaders",
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
		&requestflag.Flag[bool]{
			Name:     "send-recording-url",
			Usage:    "Whether to send RecordingUrl in webhooks.",
			Default:  true,
			BodyPath: "SendRecordingUrl",
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
			Name:     "sip-region",
			Usage:    "Defines the SIP region to be used for the call.",
			Default:  "US",
			BodyPath: "SipRegion",
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
			Name:     "supervise-call-sid",
			Usage:    "The call control ID of the existing call to supervise. When provided, the created leg will be added to the specified call in supervising mode. Status callbacks and action callbacks will NOT be sent for the supervising leg.",
			BodyPath: "SuperviseCallSid",
		},
		&requestflag.Flag[string]{
			Name:     "supervising-role",
			Usage:    "The supervising role for the new leg. Determines the audio behavior: barge (hear both sides), whisper (only hear supervisor), monitor (hear both sides but supervisor muted). Default: barge",
			Default:  "barge",
			BodyPath: "SupervisingRole",
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
	Action:          handleTexmlAccountsCallsCalls,
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

var texmlAccountsCallsRetrieveCalls = cli.Command{
	Name:    "retrieve-calls",
	Usage:   "Returns multiple call resouces for an account. This endpoint is eventually\nconsistent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "end-time",
			Usage:     "Filters calls by their end date. Expected format is YYYY-MM-DD",
			QueryPath: "EndTime",
		},
		&requestflag.Flag[string]{
			Name:      "end-time-gt",
			Usage:     "Filters calls by their end date (after). Expected format is YYYY-MM-DD",
			QueryPath: "EndTime_gt",
		},
		&requestflag.Flag[string]{
			Name:      "end-time-lt",
			Usage:     "Filters calls by their end date (before). Expected format is YYYY-MM-DD",
			QueryPath: "EndTime_lt",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Filters calls by the from number.",
			QueryPath: "From",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "The number of the page to be displayed, zero-indexed, should be used in conjuction with PageToken.",
			QueryPath: "Page",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of records to be displayed on a page",
			QueryPath: "PageSize",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Used to request the next page of results.",
			QueryPath: "PageToken",
		},
		&requestflag.Flag[string]{
			Name:      "start-time",
			Usage:     "Filters calls by their start date. Expected format is YYYY-MM-DD.",
			QueryPath: "StartTime",
		},
		&requestflag.Flag[string]{
			Name:      "start-time-gt",
			Usage:     "Filters calls by their start date (after). Expected format is YYYY-MM-DD",
			QueryPath: "StartTime_gt",
		},
		&requestflag.Flag[string]{
			Name:      "start-time-lt",
			Usage:     "Filters calls by their start date (before). Expected format is YYYY-MM-DD",
			QueryPath: "StartTime_lt",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filters calls by status.",
			QueryPath: "Status",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Filters calls by the to number.",
			QueryPath: "To",
		},
	},
	Action:          handleTexmlAccountsCallsRetrieveCalls,
	HideHelpCommand: true,
}

var texmlAccountsCallsSiprecJson = cli.Command{
	Name:    "siprec-json",
	Usage:   "Starts siprec session with specified parameters for call idientified by\ncall_sid.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "connector-name",
			Usage:    "The name of the connector to use for the SIPREC session.",
			BodyPath: "ConnectorName",
		},
		&requestflag.Flag[bool]{
			Name:     "include-metadata-custom-headers",
			Usage:    "When set, custom parameters will be added as metadata (recording.session.ExtensionParameters). Otherwise, they’ll be added to sip headers.",
			BodyPath: "IncludeMetadataCustomHeaders",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the SIPREC session. May be used to stop the SIPREC session from TeXML instruction.",
			BodyPath: "Name",
		},
		&requestflag.Flag[bool]{
			Name:     "secure",
			Usage:    "Controls whether to encrypt media sent to your SRS using SRTP and TLS. When set you need to configure SRS port in your connector to 5061.",
			BodyPath: "Secure",
		},
		&requestflag.Flag[int64]{
			Name:     "session-timeout-secs",
			Usage:    "Sets `Session-Expires` header to the INVITE. A reinvite is sent every half the value set. Usefull for session keep alive. Minimum value is 90, set to 0 to disable.",
			Default:  1800,
			BodyPath: "SessionTimeoutSecs",
		},
		&requestflag.Flag[string]{
			Name:     "sip-transport",
			Usage:    "Specifies SIP transport protocol.",
			Default:  "udp",
			BodyPath: "SipTransport",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "URL destination for Telnyx to send status callback events to for the siprec session.",
			BodyPath: "StatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request type used for `StatusCallback`.",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "track",
			Usage:    "The track to be used for siprec session. Can be `both_tracks`, `inbound_track` or `outbound_track`. Defaults to `both_tracks`.",
			BodyPath: "Track",
		},
	},
	Action:          handleTexmlAccountsCallsSiprecJson,
	HideHelpCommand: true,
}

var texmlAccountsCallsStreamsJson = cli.Command{
	Name:    "streams-json",
	Usage:   "Starts streaming media from a call to a specific WebSocket address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "bidirectional-codec",
			Usage:    "Indicates codec for bidirectional streaming RTP payloads. Used only with stream_bidirectional_mode=rtp. Case sensitive.",
			Default:  "PCMU",
			BodyPath: "BidirectionalCodec",
		},
		&requestflag.Flag[string]{
			Name:     "bidirectional-mode",
			Usage:    "Configures method of bidirectional streaming (mp3, rtp).",
			Default:  "mp3",
			BodyPath: "BidirectionalMode",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The user specified name of Stream.",
			BodyPath: "Name",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "Url where status callbacks will be sent.",
			BodyPath: "StatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP method used to send status callbacks.",
			Default:  "POST",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "track",
			Usage:    "Tracks to be included in the stream",
			Default:  "inbound_track",
			BodyPath: "Track",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The destination WebSocket address where the stream is going to be delivered.",
			BodyPath: "Url",
		},
	},
	Action:          handleTexmlAccountsCallsStreamsJson,
	HideHelpCommand: true,
}

func handleTexmlAccountsCallsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallGetParams{
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
	_, err = client.Texml.Accounts.Calls.Get(
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
	return ShowJSON(os.Stdout, "texml:accounts:calls retrieve", obj, format, transform)
}

func handleTexmlAccountsCallsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallUpdateParams{
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
	_, err = client.Texml.Accounts.Calls.Update(
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
	return ShowJSON(os.Stdout, "texml:accounts:calls update", obj, format, transform)
}

func handleTexmlAccountsCallsCalls(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallCallsParams{}

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
	_, err = client.Texml.Accounts.Calls.Calls(
		ctx,
		cmd.Value("account-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:calls calls", obj, format, transform)
}

func handleTexmlAccountsCallsRetrieveCalls(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallGetCallsParams{}

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
	_, err = client.Texml.Accounts.Calls.GetCalls(
		ctx,
		cmd.Value("account-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:calls retrieve-calls", obj, format, transform)
}

func handleTexmlAccountsCallsSiprecJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallSiprecJsonParams{
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
	_, err = client.Texml.Accounts.Calls.SiprecJson(
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
	return ShowJSON(os.Stdout, "texml:accounts:calls siprec-json", obj, format, transform)
}

func handleTexmlAccountsCallsStreamsJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallStreamsJsonParams{
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
	_, err = client.Texml.Accounts.Calls.StreamsJson(
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
	return ShowJSON(os.Stdout, "texml:accounts:calls streams-json", obj, format, transform)
}

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

var texmlInitiateAICall = requestflag.WithInnerFlags(cli.Command{
	Name:    "initiate-ai-call",
	Usage:   "Initiate an outbound AI call with warm-up support. Validates parameters, builds\nan internal TeXML with an AI Assistant configuration, encodes instructions into\nclient state, and calls the dial API. The Twiml, Texml, and Url parameters are\nnot allowed and will result in a 422 error.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "ai-assistant-id",
			Usage:    "The ID of the AI assistant to use for the call.",
			Required: true,
			BodyPath: "AIAssistantId",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "The phone number of the party initiating the call. Phone numbers are formatted with a `+` and country code.",
			Required: true,
			BodyPath: "From",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "The phone number of the called party. Phone numbers are formatted with a `+` and country code.",
			Required: true,
			BodyPath: "To",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "ai-assistant-dynamic-variables",
			Usage:    "Key-value map of dynamic variables to pass to the AI assistant.",
			BodyPath: "AIAssistantDynamicVariables",
		},
		&requestflag.Flag[string]{
			Name:     "ai-assistant-version",
			Usage:    "The version of the AI assistant to use.",
			BodyPath: "AIAssistantVersion",
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
			Usage:    "HTTP request type used for `AsyncAmdStatusCallback`.",
			Default:  "POST",
			BodyPath: "AsyncAmdStatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "caller-id",
			Usage:    "To be used as the caller id name (SIP From Display Name) presented to the destination (`To` number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and `-_~!.+` special characters. If omitted, the display name will be the same as the number in the `From` field.",
			BodyPath: "CallerId",
		},
		&requestflag.Flag[string]{
			Name:     "conversation-callback",
			Usage:    "URL destination for Telnyx to send conversation callback events to.",
			BodyPath: "ConversationCallback",
		},
		&requestflag.Flag[string]{
			Name:     "conversation-callback-method",
			Usage:    "HTTP request type used for `ConversationCallback`.",
			Default:  "POST",
			BodyPath: "ConversationCallbackMethod",
		},
		&requestflag.Flag[[]string]{
			Name:     "conversation-callback",
			Usage:    "An array of URL destinations for conversation callback events.",
			BodyPath: "ConversationCallbacks",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-header",
			Usage:    "Custom HTTP headers to be sent with the call. Each header should be an object with 'name' and 'value' properties.",
			BodyPath: "CustomHeaders",
		},
		&requestflag.Flag[string]{
			Name:     "detection-mode",
			Usage:    "Allows you to choose between Premium and Standard detections.",
			Default:  "Regular",
			BodyPath: "DetectionMode",
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
			Name:     "passports",
			Usage:    "A string of passport identifiers to associate with the call.",
			BodyPath: "Passports",
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
			Usage:    "The changes to the recording's state that should generate a call to `RecordingStatusCallback`. Can be: `in-progress`, `completed` and `absent`. Separate multiple values with a space. Defaults to `completed`.",
			BodyPath: "RecordingStatusCallbackEvent",
		},
		&requestflag.Flag[string]{
			Name:     "recording-status-callback-method",
			Usage:    "HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.",
			BodyPath: "RecordingStatusCallbackMethod",
		},
		&requestflag.Flag[int64]{
			Name:     "recording-timeout",
			Usage:    "The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected. The timer only starts when the speech is detected. The minimum value is 0. The default value is 0 (infinite).",
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
			Usage:    "The call events for which Telnyx should send a webhook. Multiple events can be defined when separated by a space. Valid values: initiated, ringing, answered, completed.",
			Default:  "completed",
			BodyPath: "StatusCallbackEvent",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request type used for `StatusCallback`.",
			Default:  "POST",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[[]string]{
			Name:     "status-callback",
			Usage:    "An array of URL destinations for Telnyx to send status callback events to for the call.",
			BodyPath: "StatusCallbacks",
		},
		&requestflag.Flag[int64]{
			Name:     "time-limit",
			Usage:    "The maximum duration of the call in seconds. The minimum value is 30 and the maximum value is 14400 (4 hours). Default is 14400 seconds.",
			Default:  14400,
			BodyPath: "TimeLimit",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-seconds",
			Usage:    "The number of seconds to wait for the called party to answer the call before the call is canceled. The minimum value is 5 and the maximum value is 120. Default is 30 seconds.",
			Default:  30,
			BodyPath: "Timeout",
		},
		&requestflag.Flag[string]{
			Name:     "trim",
			Usage:    "Whether to trim any leading and trailing silence from the recording. Defaults to `trim-silence`.",
			BodyPath: "Trim",
		},
	},
	Action:          handleTexmlInitiateAICall,
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

var texmlSecrets = cli.Command{
	Name:    "secrets",
	Usage:   "Create a TeXML secret which can be later used as a Dynamic Parameter for TeXML\nwhen using Mustache Templates in your TeXML. In your TeXML you will be able to\nuse your secret name, and this name will be replaced by the actual secret value\nwhen processing the TeXML on Telnyx side. The secrets are not visible in any\nlogs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name used as a reference for the secret, if the name already exists within the account its value will be replaced",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "value",
			Usage:    "Secret value which will be used when rendering the TeXML template",
			Required: true,
			BodyPath: "value",
		},
	},
	Action:          handleTexmlSecrets,
	HideHelpCommand: true,
}

func handleTexmlInitiateAICall(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlInitiateAICallParams{}

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
	_, err = client.Texml.InitiateAICall(
		ctx,
		cmd.Value("connection-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "texml initiate-ai-call", obj, format, explicitFormat, transform)
}

func handleTexmlSecrets(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlSecretsParams{}

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
	_, err = client.Texml.Secrets(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "texml secrets", obj, format, explicitFormat, transform)
}

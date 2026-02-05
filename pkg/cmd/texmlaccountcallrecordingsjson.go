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

var texmlAccountsCallsRecordingsJsonRecordingsJson = cli.Command{
	Name:    "recordings-json",
	Usage:   "Starts recording with specified parameters for call idientified by call_sid.",
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
		&requestflag.Flag[bool]{
			Name:     "play-beep",
			Usage:    "Whether to play a beep when recording is started.",
			Default:  true,
			BodyPath: "PlayBeep",
		},
		&requestflag.Flag[string]{
			Name:     "recording-channels",
			Usage:    "When `dual`, final audio file has the first leg on channel A, and the rest on channel B. `single` mixes both tracks into a single channel.",
			Default:  "dual",
			BodyPath: "RecordingChannels",
		},
		&requestflag.Flag[string]{
			Name:     "recording-status-callback",
			Usage:    "Url where status callbacks will be sent.",
			BodyPath: "RecordingStatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "recording-status-callback-event",
			Usage:    "The changes to the recording's state that should generate a call to `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`. Separate multiple values with a space. Defaults to `completed`.",
			BodyPath: "RecordingStatusCallbackEvent",
		},
		&requestflag.Flag[string]{
			Name:     "recording-status-callback-method",
			Usage:    "HTTP method used to send status callbacks.",
			Default:  "POST",
			BodyPath: "RecordingStatusCallbackMethod",
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
	},
	Action:          handleTexmlAccountsCallsRecordingsJsonRecordingsJson,
	HideHelpCommand: true,
}

var texmlAccountsCallsRecordingsJsonRetrieveRecordingsJson = cli.Command{
	Name:    "retrieve-recordings-json",
	Usage:   "Returns recordings for a call identified by call_sid.",
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
	Action:          handleTexmlAccountsCallsRecordingsJsonRetrieveRecordingsJson,
	HideHelpCommand: true,
}

func handleTexmlAccountsCallsRecordingsJsonRecordingsJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallRecordingsJsonRecordingsJsonParams{
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
	_, err = client.Texml.Accounts.Calls.RecordingsJson.RecordingsJson(
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
	return ShowJSON(os.Stdout, "texml:accounts:calls:recordings-json recordings-json", obj, format, transform)
}

func handleTexmlAccountsCallsRecordingsJsonRetrieveRecordingsJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallRecordingsJsonGetRecordingsJsonParams{
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
	_, err = client.Texml.Accounts.Calls.RecordingsJson.GetRecordingsJson(
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
	return ShowJSON(os.Stdout, "texml:accounts:calls:recordings-json retrieve-recordings-json", obj, format, transform)
}

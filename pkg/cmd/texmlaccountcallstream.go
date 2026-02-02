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

var texmlAccountsCallsStreamsStreamingSidJson = cli.Command{
	Name:    "streaming-sid-json",
	Usage:   "Updates streaming resource for particular call.",
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
			Name:     "streaming-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status of the Stream you wish to update.",
			Default:  "stopped",
			BodyPath: "Status",
		},
	},
	Action:          handleTexmlAccountsCallsStreamsStreamingSidJson,
	HideHelpCommand: true,
}

func handleTexmlAccountsCallsStreamsStreamingSidJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("streaming-sid") && len(unusedArgs) > 0 {
		cmd.Set("streaming-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallStreamStreamingSidJsonParams{
		AccountSid: cmd.Value("account-sid").(string),
		CallSid:    cmd.Value("call-sid").(string),
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
	_, err = client.Texml.Accounts.Calls.Streams.StreamingSidJson(
		ctx,
		cmd.Value("streaming-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:calls:streams streaming-sid-json", obj, format, transform)
}

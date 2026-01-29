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

var texmlAccountsRecordingsJsonDeleteRecordingSidJson = cli.Command{
	Name:    "delete-recording-sid-json",
	Usage:   "Deletes recording resource identified by recording id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "recording-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsRecordingsJsonDeleteRecordingSidJson,
	HideHelpCommand: true,
}

var texmlAccountsRecordingsJsonRetrieveRecordingSidJson = cli.Command{
	Name:    "retrieve-recording-sid-json",
	Usage:   "Returns recording resource identified by recording id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "recording-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsRecordingsJsonRetrieveRecordingSidJson,
	HideHelpCommand: true,
}

func handleTexmlAccountsRecordingsJsonDeleteRecordingSidJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-sid") && len(unusedArgs) > 0 {
		cmd.Set("recording-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountRecordingJsonDeleteRecordingSidJsonParams{
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

	return client.Texml.Accounts.Recordings.Json.DeleteRecordingSidJson(
		ctx,
		cmd.Value("recording-sid").(string),
		params,
		options...,
	)
}

func handleTexmlAccountsRecordingsJsonRetrieveRecordingSidJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-sid") && len(unusedArgs) > 0 {
		cmd.Set("recording-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountRecordingJsonGetRecordingSidJsonParams{
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
	_, err = client.Texml.Accounts.Recordings.Json.GetRecordingSidJson(
		ctx,
		cmd.Value("recording-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:recordings:json retrieve-recording-sid-json", obj, format, transform)
}

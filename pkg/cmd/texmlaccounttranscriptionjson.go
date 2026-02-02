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

var texmlAccountsTranscriptionsJsonDeleteRecordingTranscriptionSidJson = cli.Command{
	Name:    "delete-recording-transcription-sid-json",
	Usage:   "Permanently deletes a recording transcription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "recording-transcription-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsTranscriptionsJsonDeleteRecordingTranscriptionSidJson,
	HideHelpCommand: true,
}

var texmlAccountsTranscriptionsJsonRetrieveRecordingTranscriptionSidJson = cli.Command{
	Name:    "retrieve-recording-transcription-sid-json",
	Usage:   "Returns the recording transcription resource identified by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "recording-transcription-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsTranscriptionsJsonRetrieveRecordingTranscriptionSidJson,
	HideHelpCommand: true,
}

func handleTexmlAccountsTranscriptionsJsonDeleteRecordingTranscriptionSidJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-transcription-sid") && len(unusedArgs) > 0 {
		cmd.Set("recording-transcription-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountTranscriptionJsonDeleteRecordingTranscriptionSidJsonParams{
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

	return client.Texml.Accounts.Transcriptions.Json.DeleteRecordingTranscriptionSidJson(
		ctx,
		cmd.Value("recording-transcription-sid").(string),
		params,
		options...,
	)
}

func handleTexmlAccountsTranscriptionsJsonRetrieveRecordingTranscriptionSidJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-transcription-sid") && len(unusedArgs) > 0 {
		cmd.Set("recording-transcription-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonParams{
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
	_, err = client.Texml.Accounts.Transcriptions.Json.GetRecordingTranscriptionSidJson(
		ctx,
		cmd.Value("recording-transcription-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:transcriptions:json retrieve-recording-transcription-sid-json", obj, format, transform)
}

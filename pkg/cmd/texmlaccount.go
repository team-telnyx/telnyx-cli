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

var texmlAccountsRetrieveRecordingsJson = cli.Command{
	Name:    "retrieve-recordings-json",
	Usage:   "Returns multiple recording resources for an account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "date-created",
			Usage:     "Filters recording by the creation date. Expected format is ISO8601 date or date-time, ie. {YYYY}-{MM}-{DD} or {YYYY}-{MM}-{DD}T{hh}:{mm}:{ss}Z. Also accepts inequality operators, e.g. DateCreated>=2023-05-22.",
			QueryPath: "DateCreated",
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
	},
	Action:          handleTexmlAccountsRetrieveRecordingsJson,
	HideHelpCommand: true,
}

var texmlAccountsRetrieveTranscriptionsJson = cli.Command{
	Name:    "retrieve-transcriptions-json",
	Usage:   "Returns multiple recording transcription resources for an account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
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
	},
	Action:          handleTexmlAccountsRetrieveTranscriptionsJson,
	HideHelpCommand: true,
}

func handleTexmlAccountsRetrieveRecordingsJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountGetRecordingsJsonParams{}

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
	_, err = client.Texml.Accounts.GetRecordingsJson(
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
	return ShowJSON(os.Stdout, "texml:accounts retrieve-recordings-json", obj, format, transform)
}

func handleTexmlAccountsRetrieveTranscriptionsJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountGetTranscriptionsJsonParams{}

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
	_, err = client.Texml.Accounts.GetTranscriptionsJson(
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
	return ShowJSON(os.Stdout, "texml:accounts retrieve-transcriptions-json", obj, format, transform)
}

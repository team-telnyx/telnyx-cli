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

var texmlAccountsConferencesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a conference resource.",
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
	Action:          handleTexmlAccountsConferencesRetrieve,
	HideHelpCommand: true,
}

var texmlAccountsConferencesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a conference resource.",
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
			Name:     "announce-method",
			Usage:    "The HTTP method used to call the `AnnounceUrl`. Defaults to `POST`.",
			BodyPath: "AnnounceMethod",
		},
		&requestflag.Flag[string]{
			Name:     "announce-url",
			Usage:    "The URL we should call to announce something into the conference. The URL may return an MP3 file, a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs.",
			BodyPath: "AnnounceUrl",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The new status of the resource. Specifying `completed` will end the conference and hang up all participants.",
			BodyPath: "Status",
		},
	},
	Action:          handleTexmlAccountsConferencesUpdate,
	HideHelpCommand: true,
}

var texmlAccountsConferencesRetrieveConferences = cli.Command{
	Name:    "retrieve-conferences",
	Usage:   "Lists conference resources.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "date-created",
			Usage:     "Filters conferences by the creation date. Expected format is YYYY-MM-DD. Also accepts inequality operators, e.g. DateCreated>=2023-05-22.",
			QueryPath: "DateCreated",
		},
		&requestflag.Flag[string]{
			Name:      "date-updated",
			Usage:     "Filters conferences by the time they were last updated. Expected format is YYYY-MM-DD. Also accepts inequality operators, e.g. DateUpdated>=2023-05-22.",
			QueryPath: "DateUpdated",
		},
		&requestflag.Flag[string]{
			Name:      "friendly-name",
			Usage:     "Filters conferences by their friendly name.",
			QueryPath: "FriendlyName",
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
			Name:      "status",
			Usage:     "Filters conferences by status.",
			QueryPath: "Status",
		},
	},
	Action:          handleTexmlAccountsConferencesRetrieveConferences,
	HideHelpCommand: true,
}

var texmlAccountsConferencesRetrieveRecordings = cli.Command{
	Name:    "retrieve-recordings",
	Usage:   "Lists conference recordings",
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
	Action:          handleTexmlAccountsConferencesRetrieveRecordings,
	HideHelpCommand: true,
}

var texmlAccountsConferencesRetrieveRecordingsJson = cli.Command{
	Name:    "retrieve-recordings-json",
	Usage:   "Returns recordings for a conference identified by conference_sid.",
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
	Action:          handleTexmlAccountsConferencesRetrieveRecordingsJson,
	HideHelpCommand: true,
}

func handleTexmlAccountsConferencesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conference-sid") && len(unusedArgs) > 0 {
		cmd.Set("conference-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceGetParams{
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
	_, err = client.Texml.Accounts.Conferences.Get(
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
	return ShowJSON(os.Stdout, "texml:accounts:conferences retrieve", obj, format, transform)
}

func handleTexmlAccountsConferencesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conference-sid") && len(unusedArgs) > 0 {
		cmd.Set("conference-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceUpdateParams{
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
	_, err = client.Texml.Accounts.Conferences.Update(
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
	return ShowJSON(os.Stdout, "texml:accounts:conferences update", obj, format, transform)
}

func handleTexmlAccountsConferencesRetrieveConferences(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceGetConferencesParams{}

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
	_, err = client.Texml.Accounts.Conferences.GetConferences(
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
	return ShowJSON(os.Stdout, "texml:accounts:conferences retrieve-conferences", obj, format, transform)
}

func handleTexmlAccountsConferencesRetrieveRecordings(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conference-sid") && len(unusedArgs) > 0 {
		cmd.Set("conference-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceGetRecordingsParams{
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
	_, err = client.Texml.Accounts.Conferences.GetRecordings(
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
	return ShowJSON(os.Stdout, "texml:accounts:conferences retrieve-recordings", obj, format, transform)
}

func handleTexmlAccountsConferencesRetrieveRecordingsJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conference-sid") && len(unusedArgs) > 0 {
		cmd.Set("conference-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountConferenceGetRecordingsJsonParams{
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
	_, err = client.Texml.Accounts.Conferences.GetRecordingsJson(
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
	return ShowJSON(os.Stdout, "texml:accounts:conferences retrieve-recordings-json", obj, format, transform)
}

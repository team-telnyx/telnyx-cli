// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var emailInboxesFiltersList = cli.Command{
	Name:    "list",
	Usage:   "Returns the inbox's sender allowlist and blocklist. Entries are normalized to\nlowercase. A blocklist match takes precedence over an allowlist match; when both\nlists are empty, all senders are accepted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
	},
	Action:          handleEmailInboxesFiltersList,
	HideHelpCommand: true,
}

var emailInboxesFiltersAdd = cli.Command{
	Name:    "add",
	Usage:   "Adds entries to either the allowlist or blocklist. The operation is an\nidempotent set union: entries already present remain unchanged.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "entry",
			Required: true,
			BodyPath: "entries",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The list to change.",
			Required: true,
			BodyPath: "type",
		},
	},
	Action:          handleEmailInboxesFiltersAdd,
	HideHelpCommand: true,
}

var emailInboxesFiltersDeleteAll = cli.Command{
	Name:    "delete-all",
	Usage:   "Removes entries from either the allowlist or blocklist. The operation is\nidempotent: removing an entry that is not present still returns the current\nfilter lists.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "entry",
			Required: true,
			BodyPath: "entries",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The list to change.",
			Required: true,
			BodyPath: "type",
		},
	},
	Action:          handleEmailInboxesFiltersDeleteAll,
	HideHelpCommand: true,
}

var emailInboxesFiltersReplace = cli.Command{
	Name:    "replace",
	Usage:   "Replaces both sender filter lists atomically. Omitting either list clears that\nlist. Use `POST` or `DELETE` for incremental changes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowlist",
			BodyPath: "allowlist",
		},
		&requestflag.Flag[[]string]{
			Name:     "blocklist",
			BodyPath: "blocklist",
		},
	},
	Action:          handleEmailInboxesFiltersReplace,
	HideHelpCommand: true,
}

func handleEmailInboxesFiltersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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
	_, err = client.EmailInboxes.Filters.List(ctx, cmd.Value("inbox-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "email-inboxes:filters list",
		Transform:      transform,
	})
}

func handleEmailInboxesFiltersAdd(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := telnyx.EmailInboxFilterAddParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Filters.Add(
		ctx,
		cmd.Value("inbox-id").(string),
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "email-inboxes:filters add",
		Transform:      transform,
	})
}

func handleEmailInboxesFiltersDeleteAll(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := telnyx.EmailInboxFilterDeleteAllParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Filters.DeleteAll(
		ctx,
		cmd.Value("inbox-id").(string),
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "email-inboxes:filters delete-all",
		Transform:      transform,
	})
}

func handleEmailInboxesFiltersReplace(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbox-id") && len(unusedArgs) > 0 {
		cmd.Set("inbox-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := telnyx.EmailInboxFilterReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Filters.Replace(
		ctx,
		cmd.Value("inbox-id").(string),
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "email-inboxes:filters replace",
		Transform:      transform,
	})
}

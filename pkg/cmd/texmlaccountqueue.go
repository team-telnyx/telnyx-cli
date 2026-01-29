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

var texmlAccountsQueuesCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new queue resource.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "friendly-name",
			Usage:    "A human readable name for the queue.",
			BodyPath: "FriendlyName",
		},
		&requestflag.Flag[int64]{
			Name:     "max-size",
			Usage:    "The maximum size of the queue.",
			BodyPath: "MaxSize",
		},
	},
	Action:          handleTexmlAccountsQueuesCreate,
	HideHelpCommand: true,
}

var texmlAccountsQueuesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a queue resource.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "queue-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsQueuesRetrieve,
	HideHelpCommand: true,
}

var texmlAccountsQueuesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a queue resource.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "queue-sid",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "max-size",
			Usage:    "The maximum size of the queue.",
			BodyPath: "MaxSize",
		},
	},
	Action:          handleTexmlAccountsQueuesUpdate,
	HideHelpCommand: true,
}

var texmlAccountsQueuesList = cli.Command{
	Name:    "list",
	Usage:   "Lists queue resources.",
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
	},
	Action:          handleTexmlAccountsQueuesList,
	HideHelpCommand: true,
}

var texmlAccountsQueuesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a queue resource.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "queue-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsQueuesDelete,
	HideHelpCommand: true,
}

func handleTexmlAccountsQueuesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountQueueNewParams{}

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
	_, err = client.Texml.Accounts.Queues.New(
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
	return ShowJSON(os.Stdout, "texml:accounts:queues create", obj, format, transform)
}

func handleTexmlAccountsQueuesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-sid") && len(unusedArgs) > 0 {
		cmd.Set("queue-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountQueueGetParams{
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
	_, err = client.Texml.Accounts.Queues.Get(
		ctx,
		cmd.Value("queue-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:queues retrieve", obj, format, transform)
}

func handleTexmlAccountsQueuesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-sid") && len(unusedArgs) > 0 {
		cmd.Set("queue-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountQueueUpdateParams{
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
	_, err = client.Texml.Accounts.Queues.Update(
		ctx,
		cmd.Value("queue-sid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml:accounts:queues update", obj, format, transform)
}

func handleTexmlAccountsQueuesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountQueueListParams{}

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
	_, err = client.Texml.Accounts.Queues.List(
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
	return ShowJSON(os.Stdout, "texml:accounts:queues list", obj, format, transform)
}

func handleTexmlAccountsQueuesDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-sid") && len(unusedArgs) > 0 {
		cmd.Set("queue-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountQueueDeleteParams{
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

	return client.Texml.Accounts.Queues.Delete(
		ctx,
		cmd.Value("queue-sid").(string),
		params,
		options...,
	)
}

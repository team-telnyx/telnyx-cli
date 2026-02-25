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

var queuesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new call queue.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Usage:    "The name of the queue. Must be between 1 and 255 characters.",
			Required: true,
			BodyPath: "queue_name",
		},
		&requestflag.Flag[int64]{
			Name:     "max-size",
			Usage:    "The maximum number of calls allowed in the queue.",
			Default:  300,
			BodyPath: "max_size",
		},
	},
	Action:          handleQueuesCreate,
	HideHelpCommand: true,
}

var queuesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an existing call queue",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
	},
	Action:          handleQueuesRetrieve,
	HideHelpCommand: true,
}

var queuesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update properties of an existing call queue.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "max-size",
			Usage:    "The maximum number of calls allowed in the queue.",
			Required: true,
			BodyPath: "max_size",
		},
	},
	Action:          handleQueuesUpdate,
	HideHelpCommand: true,
}

var queuesList = cli.Command{
	Name:    "list",
	Usage:   "List all queues for the authenticated user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page",
			Default:   20,
			QueryPath: "page[size]",
		},
	},
	Action:          handleQueuesList,
	HideHelpCommand: true,
}

var queuesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an existing call queue.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
	},
	Action:          handleQueuesDelete,
	HideHelpCommand: true,
}

func handleQueuesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.QueueNewParams{}

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
	_, err = client.Queues.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "queues create", obj, format, transform)
}

func handleQueuesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-name") && len(unusedArgs) > 0 {
		cmd.Set("queue-name", unusedArgs[0])
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
	_, err = client.Queues.Get(ctx, cmd.Value("queue-name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "queues retrieve", obj, format, transform)
}

func handleQueuesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-name") && len(unusedArgs) > 0 {
		cmd.Set("queue-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.QueueUpdateParams{}

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
	_, err = client.Queues.Update(
		ctx,
		cmd.Value("queue-name").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "queues update", obj, format, transform)
}

func handleQueuesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.QueueListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Queues.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "queues list", obj, format, transform)
	} else {
		iter := client.Queues.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "queues list", iter, format, transform)
	}
}

func handleQueuesDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-name") && len(unusedArgs) > 0 {
		cmd.Set("queue-name", unusedArgs[0])
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

	return client.Queues.Delete(ctx, cmd.Value("queue-name").(string), options...)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var queuesCallsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an existing call from an existing queue",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
	},
	Action:          handleQueuesCallsRetrieve,
	HideHelpCommand: true,
}

var queuesCallsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update queued call's keep_after_hangup flag",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "keep-after-hangup",
			Usage:    "Whether the call should remain in queue after hangup.",
			BodyPath: "keep_after_hangup",
		},
	},
	Action:          handleQueuesCallsUpdate,
	HideHelpCommand: true,
}

var queuesCallsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve the list of calls in an existing queue",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleQueuesCallsList,
	HideHelpCommand: true,
}

var queuesCallsRemove = cli.Command{
	Name:    "remove",
	Usage:   "Removes an inactive call from a queue. If the call is no longer active, use this\ncommand to remove it from the queue.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Required: true,
		},
	},
	Action:          handleQueuesCallsRemove,
	HideHelpCommand: true,
}

func handleQueuesCallsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.QueueCallGetParams{
		QueueName: cmd.Value("queue-name").(string),
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
	_, err = client.Queues.Calls.Get(
		ctx,
		cmd.Value("call-control-id").(string),
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
		Title:          "queues:calls retrieve",
		Transform:      transform,
	})
}

func handleQueuesCallsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.QueueCallUpdateParams{
		QueueName: cmd.Value("queue-name").(string),
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

	return client.Queues.Calls.Update(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
}

func handleQueuesCallsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-name") && len(unusedArgs) > 0 {
		cmd.Set("queue-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.QueueCallListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Queues.Calls.List(
			ctx,
			cmd.Value("queue-name").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "queues:calls list",
			Transform:      transform,
		})
	} else {
		iter := client.Queues.Calls.ListAutoPaging(
			ctx,
			cmd.Value("queue-name").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "queues:calls list",
			Transform:      transform,
		})
	}
}

func handleQueuesCallsRemove(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-control-id") && len(unusedArgs) > 0 {
		cmd.Set("call-control-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.QueueCallRemoveParams{
		QueueName: cmd.Value("queue-name").(string),
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

	return client.Queues.Calls.Remove(
		ctx,
		cmd.Value("call-control-id").(string),
		params,
		options...,
	)
}

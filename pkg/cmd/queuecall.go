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

var queuesCallsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieve the list of calls in an existing queue",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[after], page[before], page[limit], page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleQueuesCallsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"page": {
		&requestflag.InnerFlag[string]{
			Name:       "page.after",
			Usage:      "Opaque identifier of next page",
			InnerField: "after",
		},
		&requestflag.InnerFlag[string]{
			Name:       "page.before",
			Usage:      "Opaque identifier of previous page",
			InnerField: "before",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.limit",
			Usage:      "Limit of records per single page",
			InnerField: "limit",
		},
	},
})

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "queues:calls retrieve", obj, format, transform)
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
		return ShowJSON(os.Stdout, "queues:calls list", obj, format, transform)
	} else {
		iter := client.Queues.Calls.ListAutoPaging(
			ctx,
			cmd.Value("queue-name").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "queues:calls list", iter, format, transform)
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

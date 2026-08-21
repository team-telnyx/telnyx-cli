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

var emailThreadsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a thread and a bounded page of its inbound and outbound messages,\ninterleaved in chronological order. The `inbox_id` returned by the list endpoint\nis required because a thread ID can occur in multiple inboxes. Only messages\nmatching that `(inbox_id, thread_id)` pair are returned. Threads outside the\naccount return an opaque 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "thread-id",
			Required:  true,
			PathParam: "thread_id",
		},
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Usage:     "Inbox UUID that, together with `thread_id`, identifies the thread.",
			Required:  true,
			QueryPath: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "page-after",
			Usage:     "Opaque message cursor returned by the previous thread-detail page.",
			QueryPath: "page[after]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of thread messages to return. Defaults to 25; maximum is 100.",
			Default:   25,
			QueryPath: "page[size]",
		},
	},
	Action:          handleEmailThreadsRetrieve,
	HideHelpCommand: true,
}

var emailThreadsList = cli.Command{
	Name:    "list",
	Usage:   "Lists thread summaries for the whole account, newest first, using stable cursor\npagination. An agent operating many inboxes gets every conversation in one call\ninstead of one call per inbox. Each thread carries its own `inbox_id` so a reply\ncan be routed back to the right inbox. Use `filter[inbox_id]` (repeatable) to\nnarrow the result to specific inboxes. Because a thread ID can be delivered to\nmultiple inboxes, each result is identified by its `(inbox_id, id)` pair.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "filter-inbox-id",
			Usage:     "Restrict results to one or more inboxes. Repeat the parameter\n(`filter[inbox_id][]=...&filter[inbox_id][]=...`) or pass a\ncomma-separated list. Omit to list every inbox in the account.\nInboxes outside the account are silently excluded. If the filter\nis present, it must contain at least one non-empty UUID.\n",
			QueryPath: "filter[inbox_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-label",
			Usage:     "Returns only threads carrying this label. Matching is exact and case-sensitive. Thread labels are independent of the labels on the thread's messages.",
			QueryPath: "filter[label]",
		},
		&requestflag.Flag[string]{
			Name:      "page-after",
			Usage:     "Opaque cursor returned by the previous page.",
			QueryPath: "page[after]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results to return. Defaults to 25; maximum is 100.",
			Default:   25,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEmailThreadsList,
	HideHelpCommand: true,
}

func handleEmailThreadsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("thread-id") && len(unusedArgs) > 0 {
		cmd.Set("thread-id", unusedArgs[0])
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

	params := telnyx.EmailThreadGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailThreads.Get(
		ctx,
		cmd.Value("thread-id").(string),
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
		Title:          "email-threads retrieve",
		Transform:      transform,
	})
}

func handleEmailThreadsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.EmailThreadListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailThreads.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-threads list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailThreads.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-threads list",
			Transform:      transform,
		})
	}
}

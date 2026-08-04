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

var emailInboxesThreadsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a bounded page of inbound and outbound thread messages interleaved in\nchronological order using stable cursor pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "thread-id",
			Required:  true,
			PathParam: "thread_id",
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
	Action:          handleEmailInboxesThreadsRetrieve,
	HideHelpCommand: true,
}

var emailInboxesThreadsList = cli.Command{
	Name:    "list",
	Usage:   "Lists thread summaries newest first using stable cursor pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbox-id",
			Required:  true,
			PathParam: "inbox_id",
		},
		&requestflag.Flag[string]{
			Name:      "filter-label",
			Usage:     "Returns only threads carrying this label. Thread labels are independent of the labels on the thread's messages.",
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
	},
	Action:          handleEmailInboxesThreadsList,
	HideHelpCommand: true,
}

func handleEmailInboxesThreadsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailInboxThreadGetParams{
		InboxID: cmd.Value("inbox-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Threads.Get(
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
		Title:          "email-inboxes:threads retrieve",
		Transform:      transform,
	})
}

func handleEmailInboxesThreadsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailInboxThreadListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailInboxes.Threads.List(
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
		Title:          "email-inboxes:threads list",
		Transform:      transform,
	})
}

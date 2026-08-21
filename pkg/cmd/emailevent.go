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

var emailEventsList = cli.Command{
	Name:    "list",
	Usage:   "Lists account-level email events sorted oldest first by\n`occurred_at asc, id asc`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email-id",
			Usage:     "Filter events for a specific email message UUID. Invalid UUID values are silently ignored (no filter applied).",
			QueryPath: "email_id",
		},
		&requestflag.Flag[any]{
			Name:      "event-type",
			Usage:     "Comma-separated list of event types to include. Also accepts repeated query parameters (e.g. event_type=delivered&event_type=bounced). Unknown values return no matches.",
			QueryPath: "event_type",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "Inclusive ISO 8601 start timestamp. Defaults to 30 days ago when omitted.",
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "page-cursor",
			Usage:     "Opaque URL-safe Base64 cursor returned by a previous list response.",
			QueryPath: "page_cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results to return. Defaults to 25; maximum is 100. Invalid values are clamped to the valid range.",
			Default:   25,
			QueryPath: "page_size",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "Inclusive ISO 8601 end timestamp. When `from` is provided without `to`, defaults to `from + 30 days`.",
			QueryPath: "to",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEmailEventsList,
	HideHelpCommand: true,
}

var emailEventsRetrieveStats = cli.Command{
	Name:    "retrieve-stats",
	Usage:   "Returns counts and rates for email events over a time range. The default start\ntime is 30 days ago.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "Inclusive ISO 8601 start timestamp. Defaults to 30 days ago when omitted.",
			QueryPath: "from",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "Inclusive ISO 8601 end timestamp. When `from` is provided without `to`, defaults to `from + 30 days`.",
			QueryPath: "to",
		},
	},
	Action:          handleEmailEventsRetrieveStats,
	HideHelpCommand: true,
}

func handleEmailEventsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailEventListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailEvents.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-events list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailEvents.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-events list",
			Transform:      transform,
		})
	}
}

func handleEmailEventsRetrieveStats(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailEventGetStatsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailEvents.GetStats(ctx, params, options...)
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
		Title:          "email-events retrieve-stats",
		Transform:      transform,
	})
}

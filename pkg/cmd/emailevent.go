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
	Usage:   "Lists account-level email events sorted oldest first by\n`occurred_at asc, id asc`. Each row contains a legacy email.-prefixed event_type\nand an additive canonical_event_type. Gateway rejection renders email.failed\nwith canonical email.gw_reject; ambiguous injection timeout renders\nemail.injection_timeout in both; MTA expiration renders email.bounced with\ncanonical email.expired. Message-scoped queued, sending, sandbox, cancelled, and\ndaily_limit_exceeded rows fan out per durable recipient with stable derived IDs\nmatching webhook delivery. Scheduled is the cardinality exception: account\npolling retains one message-scoped scheduled row with its stored event ID, while\nscheduled webhook publication fans out per recipient with derived IDs; reconcile\nscheduled events by message ID, event type, and occurrence time rather than\nevent UUID. Recipient-scoped stored rows retain their stored UUIDs across\npolling and webhook delivery. Legacy names are derived from stored rows; an\nAdminBounce row stored as failed renders email.failed in polling while its\nwebhook retains email.bounced, both with canonical email.failed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email-id",
			Usage:     "Filter events for a specific email message UUID. Invalid UUID values are silently ignored (no filter applied).",
			QueryPath: "email_id",
		},
		&requestflag.Flag[any]{
			Name:      "event-type",
			Usage:     "Comma-separated list of event types to include. Also accepts repeated\nquery parameters (e.g. event_type=delivered&event_type=bounced).\nUnknown values return no matches.\n\nDual-name compatibility: values are accepted\nbare or `email.`-prefixed. A legacy value keeps matching the\nrows it matched pre-rename — no widening: `failed` also\nmatches the rows that now store the canonical names of the\noutcomes it covered (`gw_reject`, `injection_timeout`,\n`expired`); `bounced` matches stored `bounced` rows only\n(recipient-scoped Expirations stored `failed` pre-rename and\nnever matched `bounced`, so `expired` is deliberately not a\n`bounced` expansion). A canonical value matches its own rows\nplus legacy rows whose recorded payload evidence proves that\noutcome (`expired` also surfaces legacy `bounced` rows with\n`bounce_category: transient`). The additive\n`canonical_event_type` field in each response row names the\ncanonical outcome.\n",
			QueryPath: "event_type",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "Inclusive ISO 8601 start timestamp. Defaults to 30 days ago when omitted.",
			QueryPath: "from",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results to return. Defaults to 25; maximum is 100. Invalid values are clamped to the valid range.",
			Default:   25,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-cursor",
			Usage:     "Opaque URL-safe Base64 cursor returned by a previous event list response. The legacy `page[after]` and flat `page_cursor` forms are also accepted.",
			QueryPath: "page[cursor]",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "Inclusive ISO 8601 end timestamp. When `from` is provided without `to`, defaults to `from + 30 days`.",
			QueryPath: "to",
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailEvents.List(ctx, params, options...)
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
		Title:          "email-events list",
		Transform:      transform,
	})
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

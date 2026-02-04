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

var callEventsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Filters call events by given filter parameters. Events are ordered by\n`occurred_at`. If filter for `leg_id` or `application_session_id` is not\npresent, it only filters events from the last 24 hours.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[application_name][contains], filter[outbound.outbound_voice_profile_id], filter[leg_id], filter[application_session_id], filter[connection_id], filter[product], filter[failed], filter[from], filter[to], filter[name], filter[type], filter[occurred_at][eq/gt/gte/lt/lte], filter[status]",
			QueryPath: "filter",
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
	Action:          handleCallEventsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.application-name",
			Usage:      "Application name filters",
			InnerField: "application_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.application-session-id",
			Usage:      "The unique identifier of the call session. A session may include multiple call leg events.",
			InnerField: "application_session_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.connection-id",
			Usage:      "The unique identifier of the conection.",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.failed",
			Usage:      "Delivery failed or not.",
			InnerField: "failed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.from",
			Usage:      "Filter by From number.",
			InnerField: "from",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.leg-id",
			Usage:      "The unique identifier of an individual call leg.",
			InnerField: "leg_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.name",
			Usage:      "If present, conferences will be filtered to those with a matching `name` attribute. Matching is case-sensitive",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.occurred-at",
			Usage:      "Event occurred_at filters",
			InnerField: "occurred_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.outbound-outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound.outbound_voice_profile_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.product",
			Usage:      "Filter by product.",
			InnerField: "product",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "If present, conferences will be filtered by status.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.to",
			Usage:      "Filter by To number.",
			InnerField: "to",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			Usage:      "Event type",
			InnerField: "type",
		},
	},
})

func handleCallEventsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallEventListParams{}

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
		_, err = client.CallEvents.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "call-events list", obj, format, transform)
	} else {
		iter := client.CallEvents.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "call-events list", iter, format, transform)
	}
}

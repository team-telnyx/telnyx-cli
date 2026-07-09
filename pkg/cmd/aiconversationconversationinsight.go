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

var aiConversationsConversationInsightsRetrieveAggregates = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-aggregates",
	Usage:   "Aggregate conversation insights by specified fields",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "created-at",
			Usage:     "Filter by creation datetime to scope the aggregation window. Supports range operators (e.g., `created_at=gte.2025-01-01T00:00:00Z` for the start of the range, `created_at=lt.2025-01-02T00:00:00Z` for the end). To build per-day time series (as the portal does for the 'Insights Over Time' chart), issue one request per day bounded by `created_at=gte.<day_start>` and `created_at=lt.<next_day_start>`.",
			QueryPath: "created_at",
		},
		&requestflag.Flag[[]string]{
			Name:      "group-by",
			Usage:     "Fields to group by (can be comma-separated or multiple parameters). Prefix a field with 'metadata.' (e.g. 'metadata.assistant_id') to group by the conversation's metadata instead of the insight result.\n\nCommon fields used for over-time charts:\n- `score` — Group by the insight's score value (e.g. for Agent Instruction Following, User Satisfaction).\n- `metadata.assistant_id` — Group by the assistant that handled the conversation.\n- `metadata.assistant_version_id` — Group by the assistant version, useful for comparing performance across versions in the portal's 'Insights Over Time' chart.\n- `metadata.telnyx_conversation_channel` — Group by conversation channel (phone_call, web_chat, etc.).",
			QueryPath: "group_by",
		},
		&requestflag.Flag[string]{
			Name:      "insight-id",
			Usage:     "Optional insight ID to filter conversation insights. Only insights matching this ID will be included in the aggregation.",
			QueryPath: "insight_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "metadata",
			QueryPath: "metadata",
		},
		&requestflag.Flag[[]string]{
			Name:      "show",
			Usage:     "Fields to include in the result (can be comma-separated or multiple parameters). Supports the same 'metadata.<key>' prefix as group_by. Each returned row will contain the grouped field values plus a `record_count` indicating how many conversation insights match that combination.",
			QueryPath: "show",
		},
	},
	Action:          handleAIConversationsConversationInsightsRetrieveAggregates,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"metadata": {
		&requestflag.InnerFlag[string]{
			Name:       "metadata.assistant-id",
			Usage:      "Filter by assistant ID (e.g., `metadata.assistant_id=eq.<assistant_id>`). When provided, only conversation insights for the specified assistant are aggregated. Used by the portal to scope the 'Insights Over Time' chart to a single assistant.",
			InnerField: "assistant_id",
		},
	},
})

func handleAIConversationsConversationInsightsRetrieveAggregates(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AIConversationConversationInsightGetAggregatesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Conversations.ConversationInsights.GetAggregates(ctx, params, options...)
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
		Title:          "ai:conversations:conversation-insights retrieve-aggregates",
		Transform:      transform,
	})
}

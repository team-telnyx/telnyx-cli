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

var llmTokenGatewayUsageRetrieveExport = cli.Command{
	Name:    "retrieve-export",
	Usage:   "Durable, replayable, content-free usage export for customer pipelines. Requires\nthe dedicated llm_token_gateway.usage_export.read V2 action; spend.read and\nusage.read do not grant it. Records are inference.request (one per immutable\nledger revision; a correction adds a new id with supersedes, keep the highest\nrevision per request_id) and guardrail.evaluation (one per guardrail event,\nincluding blocked prompts with no inference record; correlate on request_id).\nDelivery is at least once in commit order: deduplicate on id. Records appear\nafter settlement and after every older database transaction has finished\n(meta.horizon_lag_seconds); keep polling with next_cursor, which is always\nreturned and is reusable after has_more is false. After a logical database\nrestore, restart from start_time and deduplicate on id; a cursor from before it\nmay return 409 cursor_expired (no Retry-After). Retention promise 30 days. No\nprompts, responses, provider credentials, supplier or customer rates. Gateway\nreference/enforcement usage, not invoice truth or BYOK provider charges.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque next_cursor from a previous page; not combinable with start_time. Bound to the requested types; the account always comes from the credential.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   500,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "start-time",
			Usage:     "First page only: RFC 3339 with offset, at most 30 days ago and not in the future; default now - 24h.",
			QueryPath: "start_time",
		},
		&requestflag.Flag[string]{
			Name:      "types",
			Usage:     "Comma-separated subset of inference.request,guardrail.evaluation; default both.",
			QueryPath: "types",
		},
	},
	Action:          handleLlmTokenGatewayUsageRetrieveExport,
	HideHelpCommand: true,
}

var llmTokenGatewayUsageRetrieveSummary = cli.Command{
	Name:    "retrieve-summary",
	Usage:   "Gateway reference/enforcement usage, not invoice truth or BYOK provider charges.\nRequires the dedicated llm_token_gateway.usage.read V2 action for combined spend\nand guardrail reporting; existing spend.read and guardrail_events.read grants do\nnot grant this endpoint. All totals, UTC daily buckets, models and latest 20\nguardrail events share one database snapshot. UTC half-open dates, at most 31\ndays. Zero-filled days; at most 1000 models (503 rather than truncation on\noverflow/timeout). Known tokens sum independently known categories across latest\nattempt evidence, including corrections; unknown spend is excluded and\nunknown_requests/reserved_spend remain explicit. Guardrail counts count events,\nnot distinct requests. No raw spend events, prompts or secrets. This read-only\nreport creates no pagination snapshot and accepts no pagination parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "end-date",
			Required:  true,
			QueryPath: "end_date",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			Required:  true,
			QueryPath: "start_date",
		},
		&requestflag.Flag[string]{
			Name:      "token-group-id",
			Required:  true,
			QueryPath: "token_group_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "metadata",
			Usage:     `Request tag filters as metadata[<key>]=<value>, at most 5, combined with AND. A value matches the tag's text form exactly (booleans as true/false, integers in decimal), so metadata[flag]=true matches both boolean true and the string "true". Keys are case-sensitive; reserved-prefix keys are rejected. Unindexed: bounded by the account, the date range and the report timeout.`,
			QueryPath: "metadata",
		},
	},
	Action:          handleLlmTokenGatewayUsageRetrieveSummary,
	HideHelpCommand: true,
}

func handleLlmTokenGatewayUsageRetrieveExport(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewayUsageGetExportParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.Usage.GetExport(ctx, params, options...)
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
		Title:          "llm-token-gateway:usage retrieve-export",
		Transform:      transform,
	})
}

func handleLlmTokenGatewayUsageRetrieveSummary(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewayUsageGetSummaryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.Usage.GetSummary(ctx, params, options...)
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
		Title:          "llm-token-gateway:usage retrieve-summary",
		Transform:      transform,
	})
}

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

var llmTokenGatewaySpendRetrieveEvents = cli.Command{
	Name:    "retrieve-events",
	Usage:   "Return request-level usage history for the authenticated account. Supply both\nUTC dates or neither; omitting both returns the previous full UTC day. Dates use\na half-open range [start_date, end_date) of at most 31 days. Cost reflects\nbudget reference rates, not invoice charges. Unknown spend is excluded and\nreported separately. The report snapshot supports at most 10,000 matching\nrequests, independently of page size. Exceeding this bound returns HTTP 503\nwithout partial results. Narrow the date range or resource filters before\nretrying. A failed report must not be interpreted as zero usage. Request tags\ncan filter either report and metadata:<key> can group the summary.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "end-date",
			QueryPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:      "end-user-id",
			QueryPath: "end_user_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "metadata",
			Usage:     `Request tag filters as metadata[<key>]=<value>, at most 5, combined with AND. A value matches the tag's text form exactly (booleans as true/false, integers in decimal), so metadata[flag]=true matches both boolean true and the string "true". Keys are case-sensitive; reserved-prefix keys are rejected. Unindexed: bounded by the account, the date range and the report timeout.`,
			QueryPath: "metadata",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "page-snapshot",
			Usage:     "Required after page 1; account/filter-bound 15-minute stable snapshot; expired/mismatch 409.",
			QueryPath: "page[snapshot]",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			QueryPath: "start_date",
		},
		&requestflag.Flag[string]{
			Name:      "token-group-id",
			QueryPath: "token_group_id",
		},
		&requestflag.Flag[string]{
			Name:      "token-key-id",
			QueryPath: "token_key_id",
		},
		&requestflag.Flag[string]{
			Name:      "token-user-id",
			QueryPath: "token_user_id",
		},
	},
	Action:          handleLlmTokenGatewaySpendRetrieveEvents,
	HideHelpCommand: true,
}

var llmTokenGatewaySpendRetrieveSummary = cli.Command{
	Name:    "retrieve-summary",
	Usage:   "Return request counts, token counts and reference spend grouped by one\nattribution dimension. Both UTC dates and group_by are required. Dates use a\nhalf-open range [start_date, end_date) of at most 31 days. Cost reflects budget\nreference rates, not invoice charges. Unknown spend is excluded and reported\nseparately. The report snapshot supports at most 10,000 matching requests,\nindependently of page size. Exceeding this bound returns HTTP 503 without\npartial results. Narrow the date range or resource filters before retrying. A\nfailed report must not be interpreted as zero usage. Request tags can filter\neither report and metadata:<key> can group the summary.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "end-date",
			Required:  true,
			QueryPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:      "group-by",
			Usage:     "Group by one request tag key. Keys starting with ltg_, telnyx_, litellm_ or cf_ (any case) are rejected.",
			Required:  true,
			QueryPath: "group_by",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			Required:  true,
			QueryPath: "start_date",
		},
		&requestflag.Flag[string]{
			Name:      "end-user-id",
			QueryPath: "end_user_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "metadata",
			Usage:     `Request tag filters as metadata[<key>]=<value>, at most 5, combined with AND. A value matches the tag's text form exactly (booleans as true/false, integers in decimal), so metadata[flag]=true matches both boolean true and the string "true". Keys are case-sensitive; reserved-prefix keys are rejected. Unindexed: bounded by the account, the date range and the report timeout.`,
			QueryPath: "metadata",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "page-snapshot",
			Usage:     "Required after page 1; account/filter-bound 15-minute stable snapshot; expired/mismatch 409.",
			QueryPath: "page[snapshot]",
		},
		&requestflag.Flag[string]{
			Name:      "token-group-id",
			QueryPath: "token_group_id",
		},
		&requestflag.Flag[string]{
			Name:      "token-key-id",
			QueryPath: "token_key_id",
		},
		&requestflag.Flag[string]{
			Name:      "token-user-id",
			QueryPath: "token_user_id",
		},
	},
	Action:          handleLlmTokenGatewaySpendRetrieveSummary,
	HideHelpCommand: true,
}

func handleLlmTokenGatewaySpendRetrieveEvents(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewaySpendGetEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.Spend.GetEvents(ctx, params, options...)
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
		Title:          "llm-token-gateway:spend retrieve-events",
		Transform:      transform,
	})
}

func handleLlmTokenGatewaySpendRetrieveSummary(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewaySpendGetSummaryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.Spend.GetSummary(ctx, params, options...)
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
		Title:          "llm-token-gateway:spend retrieve-summary",
		Transform:      transform,
	})
}

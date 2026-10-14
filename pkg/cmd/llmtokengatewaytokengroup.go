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

var llmTokenGatewayTokenGroupsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create gateway token group within the authenticated Telnyx account. Foreign and\nunknown resource IDs return the same 404 response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "allowed-model",
			Required: true,
			BodyPath: "allowed_models",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[bool]{
			Name:     "blocked",
			BodyPath: "blocked",
		},
		&requestflag.Flag[*string]{
			Name:     "budget-duration",
			Usage:    `Allowed values: "1d", "7d", "30d".`,
			BodyPath: "budget_duration",
		},
		&requestflag.Flag[*int64]{
			Name:     "cache-ttl",
			BodyPath: "cache_ttl",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fallbacks",
			Usage:    "Per-group model fallbacks: primary alias -> ordered list of 1..3 other aliases. Every key and listed alias must be in allowed_models (a PATCH removing a referenced alias is 400); a chain never lists its own key; wildcards are rejected. Chains are one level deep (a step's own chain is not expanded). Fallback is attempted only for managed routes, only before the first provider event is forwarded, and only to steps the key may use; each attempt is reserved and recorded. {} and null both disable fallbacks.",
			BodyPath: "fallbacks",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "guardrails",
			Usage:    "Per-group guardrail policy. Missing actions default to ignore. profiles may be listed with both dlp actions ignore (staged, inactive) but a non-ignore dlp action requires at least one profile. streaming defaults to buffered when any response action is block and passthrough otherwise; passthrough with a response block is 400. Findings are codes and counts only; content is never stored. The stored form is echoed back canonicalised (all actions explicit, profiles sorted); {} is stored as an all-ignore policy, null disables.",
			BodyPath: "guardrails",
		},
		&requestflag.Flag[*float64]{
			Name:     "max-budget",
			Usage:    "USD, decimal precision at most 6 places; parse exactly to integer micro-USD, never binary float accounting.",
			BodyPath: "max_budget",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-parallel-requests",
			Usage:    "Optional maximum concurrent reserved/dispatched cache misses across replicas. Null means no concurrency cap. Cache hits do not consume slots; unknown usage retains its budget hold without a slot.",
			BodyPath: "max_parallel_requests",
		},
		&requestflag.Flag[[]string]{
			Name:     "provider-key-id",
			BodyPath: "provider_key_ids",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "retries",
			Usage:    "Per-group same-route retries. max_attempts counts the first try of each step (the primary and each fallback) and is the ceiling for the x-ltg-max-attempts request header. A step is retried only when its attempt definitely did no provider work (connection failure before dispatch, or 429/529) before the first provider event; an uncertain failure (timeout, 408/500/502/503/504) moves to the next fallback step instead. max_attempts x (1 + the longest fallback chain) must not exceed 8, else 400. The n-th retry waits retry_delay_ms (constant), retry_delay_ms*n (linear) or retry_delay_ms*2^(n-1) (exponential), at most 5 s and never past the request deadline, then jittered uniformly down to half that value so concurrent retries spread out; a provider Retry-After of at most 5 s is honoured and a longer one moves to the next step. request_timeout_ms applies to streaming requests only: it bounds each attempt's time to its first provider event and never applies after it. Non-stream requests keep the per-attempt hop and idle bounds (their headers arrive only with the finished completion). retry_delay_ms and backoff are floors for the request knobs. Each planned attempt is reserved (at most 8 per request); unused reservation is released at settlement. Managed routes only. The stored form is echoed back with defaults filled in; null disables retries.",
			BodyPath: "retries",
		},
		&requestflag.Flag[*int64]{
			Name:     "rpm-limit",
			BodyPath: "rpm_limit",
		},
		&requestflag.Flag[*int64]{
			Name:     "tpm-limit",
			BodyPath: "tpm_limit",
		},
	},
	Action:          handleLlmTokenGatewayTokenGroupsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"guardrails": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "guardrails.dlp",
			InnerField: "dlp",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "guardrails.secrets",
			InnerField: "secrets",
		},
		&requestflag.InnerFlag[string]{
			Name:       "guardrails.streaming",
			Usage:      `Allowed values: "buffered", "passthrough".`,
			InnerField: "streaming",
		},
	},
	"retries": {
		&requestflag.InnerFlag[int64]{
			Name:       "retries.max-attempts",
			InnerField: "max_attempts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "retries.backoff",
			Usage:      `Allowed values: "constant", "linear", "exponential".`,
			InnerField: "backoff",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "retries.request-timeout-ms",
			InnerField: "request_timeout_ms",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "retries.retry-delay-ms",
			InnerField: "retry_delay_ms",
		},
	},
})

var llmTokenGatewayTokenGroupsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get gateway token group within the authenticated Telnyx account. Foreign and\nunknown resource IDs return the same 404 response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleLlmTokenGatewayTokenGroupsRetrieve,
	HideHelpCommand: true,
}

var llmTokenGatewayTokenGroupsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update gateway token group within the authenticated Telnyx account. Foreign and\nunknown resource IDs return the same 404 response. Omitted fields are preserved.\nNull clears the optional concurrency, fallback and retry settings. A stale\nIf-Match returns 412; read the resource again before updating.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[string]{
			Name:       "if-match",
			Required:   true,
			HeaderPath: "If-Match",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-model",
			BodyPath: "allowed_models",
		},
		&requestflag.Flag[bool]{
			Name:     "blocked",
			BodyPath: "blocked",
		},
		&requestflag.Flag[*string]{
			Name:     "budget-duration",
			Usage:    `Allowed values: "1d", "7d", "30d".`,
			BodyPath: "budget_duration",
		},
		&requestflag.Flag[*int64]{
			Name:     "cache-ttl",
			BodyPath: "cache_ttl",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fallbacks",
			Usage:    "Per-group model fallbacks: primary alias -> ordered list of 1..3 other aliases. Every key and listed alias must be in allowed_models (a PATCH removing a referenced alias is 400); a chain never lists its own key; wildcards are rejected. Chains are one level deep (a step's own chain is not expanded). Fallback is attempted only for managed routes, only before the first provider event is forwarded, and only to steps the key may use; each attempt is reserved and recorded. {} and null both disable fallbacks.",
			BodyPath: "fallbacks",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "guardrails",
			Usage:    "Per-group guardrail policy. Missing actions default to ignore. profiles may be listed with both dlp actions ignore (staged, inactive) but a non-ignore dlp action requires at least one profile. streaming defaults to buffered when any response action is block and passthrough otherwise; passthrough with a response block is 400. Findings are codes and counts only; content is never stored. The stored form is echoed back canonicalised (all actions explicit, profiles sorted); {} is stored as an all-ignore policy, null disables.",
			BodyPath: "guardrails",
		},
		&requestflag.Flag[*float64]{
			Name:     "max-budget",
			Usage:    "USD, decimal precision at most 6 places; parse exactly to integer micro-USD, never binary float accounting.",
			BodyPath: "max_budget",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-parallel-requests",
			Usage:    "Optional maximum concurrent reserved/dispatched cache misses across replicas. Null means no concurrency cap. Cache hits do not consume slots; unknown usage retains its budget hold without a slot.",
			BodyPath: "max_parallel_requests",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "provider-key-id",
			BodyPath: "provider_key_ids",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "retries",
			Usage:    "Per-group same-route retries. max_attempts counts the first try of each step (the primary and each fallback) and is the ceiling for the x-ltg-max-attempts request header. A step is retried only when its attempt definitely did no provider work (connection failure before dispatch, or 429/529) before the first provider event; an uncertain failure (timeout, 408/500/502/503/504) moves to the next fallback step instead. max_attempts x (1 + the longest fallback chain) must not exceed 8, else 400. The n-th retry waits retry_delay_ms (constant), retry_delay_ms*n (linear) or retry_delay_ms*2^(n-1) (exponential), at most 5 s and never past the request deadline, then jittered uniformly down to half that value so concurrent retries spread out; a provider Retry-After of at most 5 s is honoured and a longer one moves to the next step. request_timeout_ms applies to streaming requests only: it bounds each attempt's time to its first provider event and never applies after it. Non-stream requests keep the per-attempt hop and idle bounds (their headers arrive only with the finished completion). retry_delay_ms and backoff are floors for the request knobs. Each planned attempt is reserved (at most 8 per request); unused reservation is released at settlement. Managed routes only. The stored form is echoed back with defaults filled in; null disables retries.",
			BodyPath: "retries",
		},
		&requestflag.Flag[*int64]{
			Name:     "rpm-limit",
			BodyPath: "rpm_limit",
		},
		&requestflag.Flag[*int64]{
			Name:     "tpm-limit",
			BodyPath: "tpm_limit",
		},
	},
	Action:          handleLlmTokenGatewayTokenGroupsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"guardrails": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "guardrails.dlp",
			InnerField: "dlp",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "guardrails.secrets",
			InnerField: "secrets",
		},
		&requestflag.InnerFlag[string]{
			Name:       "guardrails.streaming",
			Usage:      `Allowed values: "buffered", "passthrough".`,
			InnerField: "streaming",
		},
	},
	"retries": {
		&requestflag.InnerFlag[int64]{
			Name:       "retries.max-attempts",
			InnerField: "max_attempts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "retries.backoff",
			Usage:      `Allowed values: "constant", "linear", "exponential".`,
			InnerField: "backoff",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "retries.request-timeout-ms",
			InnerField: "request_timeout_ms",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "retries.retry-delay-ms",
			InnerField: "retry_delay_ms",
		},
	},
})

var llmTokenGatewayTokenGroupsList = cli.Command{
	Name:    "list",
	Usage:   "List gateway token groups within the authenticated Telnyx account. Foreign and\nunknown resource IDs return the same 404 response.",
	Suggest: true,
	Flags: []cli.Flag{
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleLlmTokenGatewayTokenGroupsList,
	HideHelpCommand: true,
}

var llmTokenGatewayTokenGroupsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Revoke resource and its keys atomically; retain historical attribution, never\ndetach into an ungoverned service key.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[string]{
			Name:       "if-match",
			Required:   true,
			HeaderPath: "If-Match",
		},
	},
	Action:          handleLlmTokenGatewayTokenGroupsDelete,
	HideHelpCommand: true,
}

func handleLlmTokenGatewayTokenGroupsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := telnyx.LlmTokenGatewayTokenGroupNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.TokenGroups.New(ctx, params, options...)
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
		Title:          "llm-token-gateway:token-groups create",
		Transform:      transform,
	})
}

func handleLlmTokenGatewayTokenGroupsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.TokenGroups.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "llm-token-gateway:token-groups retrieve",
		Transform:      transform,
	})
}

func handleLlmTokenGatewayTokenGroupsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := telnyx.LlmTokenGatewayTokenGroupUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.TokenGroups.Update(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "llm-token-gateway:token-groups update",
		Transform:      transform,
	})
}

func handleLlmTokenGatewayTokenGroupsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewayTokenGroupListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.LlmTokenGateway.TokenGroups.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "llm-token-gateway:token-groups list",
			Transform:      transform,
		})
	} else {
		iter := client.LlmTokenGateway.TokenGroups.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "llm-token-gateway:token-groups list",
			Transform:      transform,
		})
	}
}

func handleLlmTokenGatewayTokenGroupsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	params := telnyx.LlmTokenGatewayTokenGroupDeleteParams{}

	return client.LlmTokenGateway.TokenGroups.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

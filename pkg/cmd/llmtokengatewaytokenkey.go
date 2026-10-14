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

var llmTokenGatewayTokenKeysCreate = cli.Command{
	Name:    "create",
	Usage:   "Create gateway token key within the authenticated Telnyx account. Foreign and\nunknown resource IDs return the same 404 response. The token secret is returned\nonly on the original create response. A replay returns metadata without the\nsecret. Provisioning is advisory; ready is not a guarantee of inference\navailability.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "token-group-id",
			Required: true,
			BodyPath: "token_group_id",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[any]{
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
		&requestflag.Flag[any]{
			Name:     "expires-at",
			BodyPath: "expires_at",
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
		&requestflag.Flag[bool]{
			Name:     "required-end-user-id",
			Usage:    "Require a nonempty caller-asserted end-user ID; not proof of identity.",
			Default:  false,
			BodyPath: "required_end_user_id",
		},
		&requestflag.Flag[*int64]{
			Name:     "rpm-limit",
			Usage:    "Omitted or null defaults to the maximum (6000) (on a blocked key, null is kept until the key is unblocked); a value below the minimum (1) or above the maximum is rejected with 400 limit_out_of_range; a malformed value (negative, wrong type or precision) is 400 invalid_request. Block the key (blocked: true) to deny it; blocking is always accepted.",
			BodyPath: "rpm_limit",
		},
		&requestflag.Flag[*string]{
			Name:     "token-user-id",
			BodyPath: "token_user_id",
		},
		&requestflag.Flag[*int64]{
			Name:     "tpm-limit",
			Usage:    "Omitted or null defaults to the maximum (10000000) (on a blocked key, null is kept until the key is unblocked); a value below the minimum (1) or above the maximum is rejected with 400 limit_out_of_range; a malformed value (negative, wrong type or precision) is 400 invalid_request. Block the key (blocked: true) to deny it; blocking is always accepted.",
			BodyPath: "tpm_limit",
		},
	},
	Action:          handleLlmTokenGatewayTokenKeysCreate,
	HideHelpCommand: true,
}

var llmTokenGatewayTokenKeysRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get gateway token key within the authenticated Telnyx account. Foreign and\nunknown resource IDs return the same 404 response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleLlmTokenGatewayTokenKeysRetrieve,
	HideHelpCommand: true,
}

var llmTokenGatewayTokenKeysUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update gateway token key within the authenticated Telnyx account. Foreign and\nunknown resource IDs return the same 404 response. Omitted fields are preserved.\nNull clears max_parallel_requests; null resets budget and rate limits to their\nmaximum. A stale If-Match returns 412.",
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
		&requestflag.Flag[any]{
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
		&requestflag.Flag[any]{
			Name:     "expires-at",
			BodyPath: "expires_at",
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
		&requestflag.Flag[bool]{
			Name:     "required-end-user-id",
			Usage:    "Require a nonempty caller-asserted end-user ID; not proof of identity. Omission preserves the current requirement.",
			BodyPath: "required_end_user_id",
		},
		&requestflag.Flag[*int64]{
			Name:     "rpm-limit",
			Usage:    "Omitted or null defaults to the maximum (6000) (on a blocked key, null is kept until the key is unblocked); a value below the minimum (1) or above the maximum is rejected with 400 limit_out_of_range; a malformed value (negative, wrong type or precision) is 400 invalid_request. Block the key (blocked: true) to deny it; blocking is always accepted.",
			BodyPath: "rpm_limit",
		},
		&requestflag.Flag[*int64]{
			Name:     "tpm-limit",
			Usage:    "Omitted or null defaults to the maximum (10000000) (on a blocked key, null is kept until the key is unblocked); a value below the minimum (1) or above the maximum is rejected with 400 limit_out_of_range; a malformed value (negative, wrong type or precision) is 400 invalid_request. Block the key (blocked: true) to deny it; blocking is always accepted.",
			BodyPath: "tpm_limit",
		},
	},
	Action:          handleLlmTokenGatewayTokenKeysUpdate,
	HideHelpCommand: true,
}

var llmTokenGatewayTokenKeysList = cli.Command{
	Name:    "list",
	Usage:   "List gateway token keys within the authenticated Telnyx account. Foreign and\nunknown resource IDs return the same 404 response.",
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
		&requestflag.Flag[string]{
			Name:      "token-group-id",
			QueryPath: "token_group_id",
		},
		&requestflag.Flag[string]{
			Name:      "token-user-id",
			QueryPath: "token_user_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleLlmTokenGatewayTokenKeysList,
	HideHelpCommand: true,
}

var llmTokenGatewayTokenKeysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Revoke this token key. Acknowledged revocation blocks new admissions with the\nkey; requests already admitted may complete. Historical usage attribution is\nretained. Other keys in the group are unchanged.",
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
	Action:          handleLlmTokenGatewayTokenKeysDelete,
	HideHelpCommand: true,
}

func handleLlmTokenGatewayTokenKeysCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewayTokenKeyNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.TokenKeys.New(ctx, params, options...)
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
		Title:          "llm-token-gateway:token-keys create",
		Transform:      transform,
	})
}

func handleLlmTokenGatewayTokenKeysRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.LlmTokenGateway.TokenKeys.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "llm-token-gateway:token-keys retrieve",
		Transform:      transform,
	})
}

func handleLlmTokenGatewayTokenKeysUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewayTokenKeyUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.LlmTokenGateway.TokenKeys.Update(
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
		Title:          "llm-token-gateway:token-keys update",
		Transform:      transform,
	})
}

func handleLlmTokenGatewayTokenKeysList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewayTokenKeyListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.LlmTokenGateway.TokenKeys.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "llm-token-gateway:token-keys list",
			Transform:      transform,
		})
	} else {
		iter := client.LlmTokenGateway.TokenKeys.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "llm-token-gateway:token-keys list",
			Transform:      transform,
		})
	}
}

func handleLlmTokenGatewayTokenKeysDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.LlmTokenGatewayTokenKeyDeleteParams{}

	return client.LlmTokenGateway.TokenKeys.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

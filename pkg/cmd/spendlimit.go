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

var spendLimitsCreate = cli.Command{
	Name:    "create",
	Usage:   "Sets a limit for a product and period that has none. Send exactly one of\n`amount` and `unlimited: true`. The period's spend is checked at once: if it is\nalready above the new limit, the product is blocked immediately\n(`evaluation.blocked_now`). Returns 409 when a limit already exists for the\nproduct and period; update it instead. With `credential_id`, the limit applies\nto that API key only: it needs an `amount`, and the period's API key limits\ncannot add up to more than the total limit when the total has an amount.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "Limit in USD. `0` blocks at the first cent of spend.",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "product",
			Usage:    "Product to limit, as returned in `product` by the list operation.",
			Required: true,
			BodyPath: "product",
		},
		&requestflag.Flag[*string]{
			Name:     "credential-id",
			Usage:    "Optional. Identifier of one API key as recorded on usage, never the key's secret: the limit then applies to that key only. Absent or `null` sets the total limit. A blank value, or one with control characters, is refused; it never means the total limit. An API key limit always has an `amount`.",
			BodyPath: "credential_id",
		},
		&requestflag.Flag[string]{
			Name:     "period",
			Usage:    "`daily` is the current UTC day; `monthly` is the current UTC calendar month.",
			Default:  "daily",
			BodyPath: "period",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Why the limit is set or changed, kept for audit.",
			BodyPath: "reason",
		},
		&requestflag.Flag[bool]{
			Name:     "unlimited",
			Usage:    "Optional; only `false` is allowed together with `amount`.",
			BodyPath: "unlimited",
		},
	},
	Action:          handleSpendLimitsCreate,
	HideHelpCommand: true,
}

var spendLimitsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Replaces the value of the existing limit for the product and period. Send\nexactly one of `amount` and `unlimited: true`. The period's spend is checked at\nonce: raising the limit above the spend lifts the period's block\n(`evaluation.released`), and lowering it below the spend blocks the product\n(`evaluation.blocked_now`). Returns 404 when no limit is set; create it instead.\nWith the `credential_id` query parameter, the request changes that API key's\nlimit; it needs an `amount`, and it can be at most the period's total limit.\nLowering the total limit below the period's highest API key limit is refused.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "product",
			Required:  true,
			PathParam: "product",
		},
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "Limit in USD. `0` blocks at the first cent of spend.",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:      "credential-id",
			Usage:     "Identifier of one API key as recorded on usage, never the key's secret, percent-encoded: the request is then about that key's limit. Absent means the total limit; present but blank is refused.",
			QueryPath: "credential_id",
		},
		&requestflag.Flag[string]{
			Name:      "period",
			Usage:     "`daily` is the current UTC day; `monthly` is the current UTC calendar month.",
			Default:   "daily",
			QueryPath: "period",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Why the limit is set or changed, kept for audit.",
			BodyPath: "reason",
		},
		&requestflag.Flag[bool]{
			Name:     "unlimited",
			Usage:    "Optional; only `false` is allowed together with `amount`.",
			BodyPath: "unlimited",
		},
	},
	Action:          handleSpendLimitsUpdate,
	HideHelpCommand: true,
}

var spendLimitsList = cli.Command{
	Name:            "list",
	Usage:           "Returns one entry per product and period you can set a limit on, with the limit,\nthe spend so far in the period and whether the product is blocked. An entry\nwithout a limit is still listed (`limit: null`). When the spend cannot be read,\nthe entry is returned with `spend_usd: null` and `spend_error` set. The list is\nnot paginated. Each entry also lists the API key limits of its product and\nperiod in `credential_limits`, with each key's spend and block. Each API key\nlimit can be at most the total limit; together they may add up to more than it.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleSpendLimitsList,
	HideHelpCommand: true,
}

var spendLimitsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Removes the limit for the product and period. For `inference`, which has no\ndefault limit, the product becomes unlimited for the period and the period's\nblock is lifted (`evaluation.released`). The response carries `limit: null` and\nthe `effective_limit_usd` that applies after the removal. Returns 404 when no\nlimit is set. With the `credential_id` query parameter, the request removes that\nAPI key's limit for the period and lifts the key's block in that period at once.\nA block under the key's limit in the other period stays\n(`evaluation.still_blocked_other_period`). Removing the total limit leaves the\nAPI key limits in place.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "product",
			Required:  true,
			PathParam: "product",
		},
		&requestflag.Flag[string]{
			Name:      "credential-id",
			Usage:     "Identifier of one API key as recorded on usage, never the key's secret, percent-encoded: the request is then about that key's limit. Absent means the total limit; present but blank is refused.",
			QueryPath: "credential_id",
		},
		&requestflag.Flag[string]{
			Name:      "period",
			Usage:     "`daily` is the current UTC day; `monthly` is the current UTC calendar month.",
			Default:   "daily",
			QueryPath: "period",
		},
		&requestflag.Flag[string]{
			Name:      "reason",
			Usage:     "Why the limit is removed, kept for audit. At most 500 characters.",
			QueryPath: "reason",
		},
	},
	Action:          handleSpendLimitsDelete,
	HideHelpCommand: true,
}

func handleSpendLimitsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.SpendLimitNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SpendLimits.New(ctx, params, options...)
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
		Title:          "spend-limits create",
		Transform:      transform,
	})
}

func handleSpendLimitsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("product") && len(unusedArgs) > 0 {
		cmd.Set("product", unusedArgs[0])
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

	params := telnyx.SpendLimitUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SpendLimits.Update(
		ctx,
		cmd.Value("product").(string),
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
		Title:          "spend-limits update",
		Transform:      transform,
	})
}

func handleSpendLimitsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SpendLimits.List(ctx, options...)
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
		Title:          "spend-limits list",
		Transform:      transform,
	})
}

func handleSpendLimitsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("product") && len(unusedArgs) > 0 {
		cmd.Set("product", unusedArgs[0])
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

	params := telnyx.SpendLimitDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SpendLimits.Delete(
		ctx,
		cmd.Value("product").(string),
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
		Title:          "spend-limits delete",
		Transform:      transform,
	})
}

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

var emailUnsubscribeGroupsSuppressionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a suppression with `reason: unsubscribe`, `source: manual`,\n`group_id: <this group>`. All other body fields are ignored; only `to` is read.\nIdempotent (same dedupe key → `200`, no new event).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Required: true,
			BodyPath: "to",
		},
	},
	Action:          handleEmailUnsubscribeGroupsSuppressionsCreate,
	HideHelpCommand: true,
}

var emailUnsubscribeGroupsSuppressionsList = cli.Command{
	Name:    "list",
	Usage:   "Account + group scoped. Offset pagination only (`page[number]` default 1,\n`page[size]` default 25, max 100). No `sort`/`filter`/ cursor — ordering fixed\n`desc created_at, desc id`. Uses the shared `QueryParser.parse_offset/1` — a\nmalformed `page` returns `400` (code `10015`), consistent with\n`GET /v2/email_blocks`. `meta` includes `total_pages`. Rows reuse the standard\nsuppression shape (`group_id` set to this group).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Offset page number (≥1, default 1).",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Page size (1–100, default 25).",
			Default:   25,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEmailUnsubscribeGroupsSuppressionsList,
	HideHelpCommand: true,
}

var emailUnsubscribeGroupsSuppressionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Soft-deletes all active blocks for (account, group, normalized email) — one\n`removed` audit event per block (`actor: manual`). The `email` path segment is\nnormalized (trim + lower-case) before matching. Idempotent on already-removed\nrows (returns `404` since they're no longer `active`).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "email",
			Required:  true,
			PathParam: "email",
		},
	},
	Action:          handleEmailUnsubscribeGroupsSuppressionsDelete,
	HideHelpCommand: true,
}

func handleEmailUnsubscribeGroupsSuppressionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailUnsubscribeGroupSuppressionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailUnsubscribeGroups.Suppressions.New(
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
		Title:          "email-unsubscribe-groups:suppressions create",
		Transform:      transform,
	})
}

func handleEmailUnsubscribeGroupsSuppressionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailUnsubscribeGroupSuppressionListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailUnsubscribeGroups.Suppressions.List(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-unsubscribe-groups:suppressions list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailUnsubscribeGroups.Suppressions.ListAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-unsubscribe-groups:suppressions list",
			Transform:      transform,
		})
	}
}

func handleEmailUnsubscribeGroupsSuppressionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("email") && len(unusedArgs) > 0 {
		cmd.Set("email", unusedArgs[0])
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

	params := telnyx.EmailUnsubscribeGroupSuppressionDeleteParams{
		ID: cmd.Value("id").(string),
	}

	return client.EmailUnsubscribeGroups.Suppressions.Delete(
		ctx,
		cmd.Value("email").(string),
		params,
		options...,
	)
}

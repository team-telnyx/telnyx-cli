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

var emailBlocksCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a suppression with `reason: manual_block` and `source: manual`.\nCaller-supplied `reason` / `source` are **ignored**; `scope` is **derived**\nserver-side from `domain_id` / `from` and is never trusted. Idempotent: if a\nmatching row already exists (NULL-safe dedupe key: account_id, scope, to,\nreason, domain_id, from), returns the existing record with `200` (no new audit\nevent).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "Recipient address (normalized: trim + lower-case).",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[*string]{
			Name:     "domain-id",
			Usage:    "`null` ⇒ account scope.",
			BodyPath: "domain_id",
		},
		&requestflag.Flag[any]{
			Name:     "expires-at",
			BodyPath: "expires_at",
		},
		&requestflag.Flag[*string]{
			Name:     "from",
			Usage:    "Sender address (normalized). `null` ⇒ account/domain scope.",
			BodyPath: "from",
		},
	},
	Action:          handleEmailBlocksCreate,
	HideHelpCommand: true,
}

var emailBlocksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the account-owned suppression identified by ID. Cross-account lookups\nand malformed IDs return `404` without exposing another account’s data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailBlocksRetrieve,
	HideHelpCommand: true,
}

var emailBlocksList = cli.Command{
	Name:    "list",
	Usage:   "Account-scoped list. Two mutually exclusive pagination modes:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "filter-created-after",
			Usage:     "`created_at > value` (ISO 8601).",
			QueryPath: "filter[created_after]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-created-before",
			Usage:     "`created_at < value` (ISO 8601).",
			QueryPath: "filter[created_before]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-domain-id",
			Usage:     "Exact-match filter on domain_id (UUID).",
			QueryPath: "filter[domain_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-reason",
			Usage:     "Exact-match filter on reason.",
			QueryPath: "filter[reason]",
		},
		&requestflag.Flag[string]{
			Name:      "page-after",
			Usage:     "Opaque cursor (`Base.url_encode64` of `{\"created_at\",\"id\"}`). Cursor mode; mutually exclusive with `page[number]` and `page[before]`.",
			QueryPath: "page[after]",
		},
		&requestflag.Flag[string]{
			Name:      "page-before",
			Usage:     "Opaque cursor (see `page[after]`). Mutually exclusive with `page[after]` and `page[number]`.",
			QueryPath: "page[before]",
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
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sort field. Leading `-` = desc; only `created_at` is sortable. Default `-created_at`. `--` is an error.",
			Default:   "-created_at",
			QueryPath: "sort",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEmailBlocksList,
	HideHelpCommand: true,
}

var emailBlocksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Soft-deletes (status → `removed`; tombstone retained). A `removed` audit event\nis appended unless the block was already `removed` (idempotent — returns the\nexisting row with `200` and no new event). Mutates `updated_at`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailBlocksDelete,
	HideHelpCommand: true,
}

var emailBlocksRetrieveEvents = cli.Command{
	Name:    "retrieve-events",
	Usage:   "Offset pagination only (`page[number]` default 1, `page[size]` default **50**,\nmax 100). No `sort`, no `filter`, no cursor — ordering is fixed\n`desc occurred_at, desc id`. Verifies the block belongs to the account first\n(cross-account → 404).",
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
			Usage:     "Page size (default 50, max 100).",
			Default:   50,
			QueryPath: "page[size]",
		},
	},
	Action:          handleEmailBlocksRetrieveEvents,
	HideHelpCommand: true,
}

func handleEmailBlocksCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailBlockNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailBlocks.New(ctx, params, options...)
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
		Title:          "email-blocks create",
		Transform:      transform,
	})
}

func handleEmailBlocksRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EmailBlocks.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email-blocks retrieve",
		Transform:      transform,
	})
}

func handleEmailBlocksList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailBlockListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailBlocks.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-blocks list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailBlocks.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-blocks list",
			Transform:      transform,
		})
	}
}

func handleEmailBlocksDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EmailBlocks.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email-blocks delete",
		Transform:      transform,
	})
}

func handleEmailBlocksRetrieveEvents(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailBlockGetEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailBlocks.GetEvents(
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
		Title:          "email-blocks retrieve-events",
		Transform:      transform,
	})
}

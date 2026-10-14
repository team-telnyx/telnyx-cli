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

var emailTemplatesCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a Liquid email template. Variables are auto-extracted when omitted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Letters, numbers, spaces, hyphens, and underscores only.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "autoescape",
			Usage:    "Per-template HTML autoescaping setting. Defaults to `false` for backward compatibility. When `true`, the rendered `html_body` HTML-escapes each Liquid expression's output at the output boundary (after its filters run, before concatenation with literal template markup). Input values are never mutated and `subject`/`text_body` are never autoescaped. The boundary escape is idempotent: HTML entities already present in the output (e.g. from an explicit `escape` filter) are preserved, so an explicit `escape`/`escape_once` is never double-escaped, and markup introduced by any later filter in the chain is still escaped.",
			Default:  false,
			BodyPath: "autoescape",
		},
		&requestflag.Flag[*string]{
			Name:     "html-body",
			Usage:    "Liquid template HTML body.",
			BodyPath: "html_body",
		},
		&requestflag.Flag[bool]{
			Name:     "strict-variables",
			Usage:    "Per-template strict variable-validation setting. Defaults to `false` for backward compatibility. When `true`, a send or render that is missing a variable marked `required: true` in `variable_schema` fails with 422 naming the variable. Missing optional variables never fail; their schema `default` (when set) is applied to the render.",
			Default:  false,
			BodyPath: "strict_variables",
		},
		&requestflag.Flag[*string]{
			Name:     "subject",
			Usage:    "Liquid template subject.",
			BodyPath: "subject",
		},
		&requestflag.Flag[*string]{
			Name:     "text-body",
			Usage:    "Liquid template text body.",
			BodyPath: "text_body",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "variable-schema",
			Usage:    "Structured variable requirements. Required variables cannot define defaults; invalid combinations return 422. This is independent of the legacy `variables` array. On render with `strict_variables` enabled: `required` variables must be supplied as non-empty values — absent, `null`, empty string, empty object `{}`, and empty array `[]` all fail with 422 naming the variable, while present values such as `false` and `0` pass (they are present, not empty). Optional variables fall back to their `default` when absent.",
			BodyPath: "variable_schema",
		},
		&requestflag.Flag[any]{
			Name:     "variable",
			Usage:    "Persisted variable names. In `automatic` mode a non-null array is rejected; in `explicit` mode this is the exact manual array and each item is limited to 255 Unicode code points.",
			BodyPath: "variables",
		},
		&requestflag.Flag[string]{
			Name:     "variables-mode",
			Usage:    "Controls how `variables` is maintained. If both fields are omitted (or `variables` is null), the template uses `automatic` mode and extracts variable names from the effective subject, HTML body, and text body. A non-null `variables` array with no mode selects `explicit`. `automatic` rejects a simultaneous non-null array; `explicit` requires a non-null array, including valid `[]`, and preserves its exact order, duplicates, and string contents.",
			BodyPath: "variables_mode",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleEmailTemplatesCreate,
	HideHelpCommand: true,
}

var emailTemplatesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the account-owned template identified by ID, including its Liquid\nsubject and bodies, declared variables, and timestamps.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailTemplatesRetrieve,
	HideHelpCommand: true,
}

var emailTemplatesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates one or more fields of the specified email template and returns the\nupdated template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "autoescape",
			Usage:    "Per-template HTML autoescaping setting.",
			BodyPath: "autoescape",
		},
		&requestflag.Flag[*string]{
			Name:     "html-body",
			Usage:    "Liquid template HTML body.",
			BodyPath: "html_body",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "strict-variables",
			Usage:    "Per-template strict variable-validation setting.",
			BodyPath: "strict_variables",
		},
		&requestflag.Flag[*string]{
			Name:     "subject",
			Usage:    "Liquid template subject.",
			BodyPath: "subject",
		},
		&requestflag.Flag[*string]{
			Name:     "text-body",
			Usage:    "Liquid template text body.",
			BodyPath: "text_body",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "variable-schema",
			Usage:    "Structured variable requirements. Required variables cannot define defaults; invalid combinations return 422. Set to `null` to clear the schema.",
			BodyPath: "variable_schema",
		},
		&requestflag.Flag[any]{
			Name:     "variable",
			Usage:    "Persisted variable names. In `automatic` mode a non-null array is rejected; in `explicit` mode this is the exact manual array and each item is limited to 255 Unicode code points.",
			BodyPath: "variables",
		},
		&requestflag.Flag[string]{
			Name:     "variables-mode",
			Usage:    "Controls how `variables` is maintained. Setting `automatic` re-extracts immediately from the effective subject, HTML body, and text body and rejects a simultaneous non-null array. Setting `explicit` requires the exact non-null array, including valid `[]`. With the mode omitted, a non-null array selects explicit mode, null selects automatic mode, and omitting both fields preserves the current mode. Source-content changes re-extract only in automatic mode.",
			BodyPath: "variables_mode",
		},
	},
	Action:          handleEmailTemplatesUpdate,
	HideHelpCommand: true,
}

var emailTemplatesList = cli.Command{
	Name:    "list",
	Usage:   "Lists templates sorted newest first by `created_at desc, id desc`. Optional\nfilters match a case-insensitive literal name/subject query, stored body\npresence, and the `strict_variables` and `autoescape` settings. Filters\nintersect, remain scoped to the authenticated account, and apply before cursor\npagination. Invalid filter names or values return 400.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-autoescape",
			Usage:     "Exact match on the template's `autoescape` option. Only the exact query strings `true` and `false` are accepted; other values, including blank, return 400.",
			QueryPath: "filter[autoescape]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-content-type",
			Usage:     "Matches templates by stored body presence. `html_and_text` requires both an HTML and text body; `html_only` requires HTML and no text body; `text_only` requires text and no HTML body; `no_body` requires neither. `null` and exactly `\"\"` count as absent; whitespace-only strings count as content. An empty or unsupported value returns a 400 response.",
			QueryPath: "filter[content_type]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-query",
			Usage:     "Case-insensitive literal substring match over the template `name` or stored `subject` only; bodies are never searched. `%`, `_`, and `\\` are literal characters with no wildcard meaning. A blank value is treated as omitted.",
			QueryPath: "filter[query]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-strict-variables",
			Usage:     "Exact match on the template's `strict_variables` option. Only the exact query strings `true` and `false` are accepted; other values, including blank, return 400.",
			QueryPath: "filter[strict_variables]",
		},
		&requestflag.Flag[string]{
			Name:      "page-cursor",
			Usage:     "Opaque URL-safe Base64 cursor returned by a previous list response.",
			QueryPath: "page_cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results to return. Defaults to 25; maximum is 100. Invalid values are clamped to the valid range.",
			Default:   25,
			QueryPath: "page_size",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEmailTemplatesList,
	HideHelpCommand: true,
}

var emailTemplatesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the account-owned template. The operation returns `204` with no body and\nprevents future sends or renders from using the deleted template ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailTemplatesDelete,
	HideHelpCommand: true,
}

var emailTemplatesRender = cli.Command{
	Name:    "render",
	Usage:   "Renders a template using the provided Liquid variables. Missing\n`template_variables` defaults to `{}`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "template-variables",
			Usage:    "Variables for Liquid template rendering. Non-object values are silently treated as an empty object.",
			Default:  map[string]any{},
			BodyPath: "template_variables",
		},
	},
	Action:          handleEmailTemplatesRender,
	HideHelpCommand: true,
}

var emailTemplatesReplace = cli.Command{
	Name:    "replace",
	Usage:   "Replaces template fields. Behaves identically to PATCH; provided for\ncompatibility with Phoenix resource routes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "autoescape",
			Usage:    "Per-template HTML autoescaping setting.",
			BodyPath: "autoescape",
		},
		&requestflag.Flag[*string]{
			Name:     "html-body",
			Usage:    "Liquid template HTML body.",
			BodyPath: "html_body",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "strict-variables",
			Usage:    "Per-template strict variable-validation setting.",
			BodyPath: "strict_variables",
		},
		&requestflag.Flag[*string]{
			Name:     "subject",
			Usage:    "Liquid template subject.",
			BodyPath: "subject",
		},
		&requestflag.Flag[*string]{
			Name:     "text-body",
			Usage:    "Liquid template text body.",
			BodyPath: "text_body",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "variable-schema",
			Usage:    "Structured variable requirements. Required variables cannot define defaults; invalid combinations return 422. Set to `null` to clear the schema.",
			BodyPath: "variable_schema",
		},
		&requestflag.Flag[any]{
			Name:     "variable",
			Usage:    "Persisted variable names. In `automatic` mode a non-null array is rejected; in `explicit` mode this is the exact manual array and each item is limited to 255 Unicode code points.",
			BodyPath: "variables",
		},
		&requestflag.Flag[string]{
			Name:     "variables-mode",
			Usage:    "Controls how `variables` is maintained. Setting `automatic` re-extracts immediately from the effective subject, HTML body, and text body and rejects a simultaneous non-null array. Setting `explicit` requires the exact non-null array, including valid `[]`. With the mode omitted, a non-null array selects explicit mode, null selects automatic mode, and omitting both fields preserves the current mode. Source-content changes re-extract only in automatic mode.",
			BodyPath: "variables_mode",
		},
	},
	Action:          handleEmailTemplatesReplace,
	HideHelpCommand: true,
}

func handleEmailTemplatesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailTemplateNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailTemplates.New(ctx, params, options...)
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
		Title:          "email-templates create",
		Transform:      transform,
	})
}

func handleEmailTemplatesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EmailTemplates.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email-templates retrieve",
		Transform:      transform,
	})
}

func handleEmailTemplatesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailTemplateUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailTemplates.Update(
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
		Title:          "email-templates update",
		Transform:      transform,
	})
}

func handleEmailTemplatesList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailTemplateListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailTemplates.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-templates list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailTemplates.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-templates list",
			Transform:      transform,
		})
	}
}

func handleEmailTemplatesDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.EmailTemplates.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleEmailTemplatesRender(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailTemplateRenderParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailTemplates.Render(
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
		Title:          "email-templates render",
		Transform:      transform,
	})
}

func handleEmailTemplatesReplace(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailTemplateReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailTemplates.Replace(
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
		Title:          "email-templates replace",
		Transform:      transform,
	})
}

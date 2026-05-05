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

var whatsappTemplatesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Whatsapp message template",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "Template category: AUTHENTICATION, UTILITY, or MARKETING.",
			Required: true,
			BodyPath: "category",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "component",
			Usage:    "Template components defining message structure. Passed through to Meta Graph API. Templates with variables must include example values. Supports HEADER, BODY, FOOTER, BUTTONS, CAROUSEL and any future Meta component types.",
			Required: true,
			BodyPath: "components",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "Template language code (e.g. en_US, es, pt_BR).",
			Required: true,
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Template name. Lowercase letters, numbers, and underscores only.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "waba-id",
			Usage:    "The WhatsApp Business Account ID.",
			Required: true,
			BodyPath: "waba_id",
		},
	},
	Action:          handleWhatsappTemplatesCreate,
	HideHelpCommand: true,
}

var whatsappTemplatesList = cli.Command{
	Name:    "list",
	Usage:   "List Whatsapp message templates",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-category",
			Usage:     "Filter by category",
			QueryPath: "filter[category]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-search",
			Usage:     "Search templates by name",
			QueryPath: "filter[search]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     "Filter by template status",
			QueryPath: "filter[status]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-waba-id",
			Usage:     "Filter by WABA ID",
			QueryPath: "filter[waba_id]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleWhatsappTemplatesList,
	HideHelpCommand: true,
}

func handleWhatsappTemplatesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.WhatsappTemplateNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Whatsapp.Templates.New(ctx, params, options...)
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
		Title:          "whatsapp:templates create",
		Transform:      transform,
	})
}

func handleWhatsappTemplatesList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.WhatsappTemplateListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Whatsapp.Templates.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "whatsapp:templates list",
			Transform:      transform,
		})
	} else {
		iter := client.Whatsapp.Templates.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "whatsapp:templates list",
			Transform:      transform,
		})
	}
}

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

var aiCollectionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new collection scoped to your organization. Optionally attach sources\nand retrieval settings at creation time. If `slug` is omitted, one is derived\nfrom `name` and must be unique within your organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable collection name.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			BodyPath: "settings",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			Usage:    "Optional slug (unique per organization). Derived from `name` when omitted.",
			BodyPath: "slug",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "source",
			Usage:    "Optional sources to attach at creation time.",
			BodyPath: "sources",
		},
	},
	Action:          handleAICollectionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"settings": {
		&requestflag.InnerFlag[string]{
			Name:       "settings.record-type",
			Usage:      "Identifies the record type. Always `ai_collection_settings`.",
			InnerField: "record_type",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "settings.retrieval",
			Usage:      "How documents are retrieved when searching the collection.",
			InnerField: "retrieval",
		},
	},
	"source": {
		&requestflag.InnerFlag[string]{
			Name:       "source.source-type",
			Usage:      "The type of Telnyx data attached as a source. `bucket` requires an additional `bucket_id`. Only `voice` is searchable today; `meeting_bot`, `message`, and `bucket` attach but are not yet searchable (Coming soon).",
			InnerField: "source_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "source.bucket-id",
			Usage:      "The Telnyx Storage bucket name. Required when `source_type` is `bucket`; ignored otherwise.",
			InnerField: "bucket_id",
		},
	},
})

var aiCollectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetches a single collection by its `slug`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "slug",
			Required:  true,
			PathParam: "slug",
		},
	},
	Action:          handleAICollectionsRetrieve,
	HideHelpCommand: true,
}

var aiCollectionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a collection's metadata (`name` and/or `description`). Sources and\nsettings are managed through their own sub-resources.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleAICollectionsUpdate,
	HideHelpCommand: true,
}

var aiCollectionsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of collections in your organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number to return (1-based). Defaults to 1.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results per page. Defaults to 20.",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAICollectionsList,
	HideHelpCommand: true,
}

var aiCollectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Soft-deletes a collection. Its `slug` is freed and may be reused by a new\ncollection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
	},
	Action:          handleAICollectionsDelete,
	HideHelpCommand: true,
}

var aiCollectionsRetrieveByID = cli.Command{
	Name:    "retrieve-by-id",
	Usage:   "Fetches a single collection by its `uuid`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
	},
	Action:          handleAICollectionsRetrieveByID,
	HideHelpCommand: true,
}

func handleAICollectionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AICollectionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Collections.New(ctx, params, options...)
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
		Title:          "ai:collections create",
		Transform:      transform,
	})
}

func handleAICollectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("slug") && len(unusedArgs) > 0 {
		cmd.Set("slug", unusedArgs[0])
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
	_, err = client.AI.Collections.Get(ctx, cmd.Value("slug").(string), options...)
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
		Title:          "ai:collections retrieve",
		Transform:      transform,
	})
}

func handleAICollectionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("uuid") && len(unusedArgs) > 0 {
		cmd.Set("uuid", unusedArgs[0])
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

	params := telnyx.AICollectionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Collections.Update(
		ctx,
		cmd.Value("uuid").(string),
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
		Title:          "ai:collections update",
		Transform:      transform,
	})
}

func handleAICollectionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AICollectionListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Collections.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "ai:collections list",
			Transform:      transform,
		})
	} else {
		iter := client.AI.Collections.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "ai:collections list",
			Transform:      transform,
		})
	}
}

func handleAICollectionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("uuid") && len(unusedArgs) > 0 {
		cmd.Set("uuid", unusedArgs[0])
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

	return client.AI.Collections.Delete(ctx, cmd.Value("uuid").(string), options...)
}

func handleAICollectionsRetrieveByID(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("uuid") && len(unusedArgs) > 0 {
		cmd.Set("uuid", unusedArgs[0])
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
	_, err = client.AI.Collections.GetByID(ctx, cmd.Value("uuid").(string), options...)
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
		Title:          "ai:collections retrieve-by-id",
		Transform:      transform,
	})
}

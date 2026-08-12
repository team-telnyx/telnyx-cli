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

var aiCollectionsSourcesCreate = cli.Command{
	Name:    "create",
	Usage:   "Attaches a new source to a collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
		&requestflag.Flag[string]{
			Name:     "source-type",
			Usage:    "The type of Telnyx data attached as a source. `bucket` requires an additional `bucket_id`. Only `voice` is searchable today; `meeting_bot`, `message`, and `bucket` attach but are not yet searchable (Coming soon).",
			Required: true,
			BodyPath: "source_type",
		},
		&requestflag.Flag[string]{
			Name:     "bucket-id",
			Usage:    "The Telnyx Storage bucket name. Required when `source_type` is `bucket`; ignored otherwise.",
			BodyPath: "bucket_id",
		},
	},
	Action:          handleAICollectionsSourcesCreate,
	HideHelpCommand: true,
}

var aiCollectionsSourcesList = cli.Command{
	Name:    "list",
	Usage:   "Returns the sources attached to a collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
	},
	Action:          handleAICollectionsSourcesList,
	HideHelpCommand: true,
}

var aiCollectionsSourcesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Removes a single source from a collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
		&requestflag.Flag[string]{
			Name:      "source-id",
			Required:  true,
			PathParam: "sourceId",
		},
	},
	Action:          handleAICollectionsSourcesDelete,
	HideHelpCommand: true,
}

var aiCollectionsSourcesReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Replaces the collection's entire source set. The response `meta` reports which\nsources were added, retained, and removed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "source",
			Required: true,
			BodyPath: "sources",
		},
	},
	Action:          handleAICollectionsSourcesReplace,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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

func handleAICollectionsSourcesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AICollectionSourceNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Collections.Sources.New(
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
		Title:          "ai:collections:sources create",
		Transform:      transform,
	})
}

func handleAICollectionsSourcesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Collections.Sources.List(ctx, cmd.Value("uuid").(string), options...)
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
		Title:          "ai:collections:sources list",
		Transform:      transform,
	})
}

func handleAICollectionsSourcesDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("source-id") && len(unusedArgs) > 0 {
		cmd.Set("source-id", unusedArgs[0])
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

	params := telnyx.AICollectionSourceDeleteParams{
		Uuid: cmd.Value("uuid").(string),
	}

	return client.AI.Collections.Sources.Delete(
		ctx,
		cmd.Value("source-id").(string),
		params,
		options...,
	)
}

func handleAICollectionsSourcesReplace(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AICollectionSourceReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Collections.Sources.Replace(
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
		Title:          "ai:collections:sources replace",
		Transform:      transform,
	})
}

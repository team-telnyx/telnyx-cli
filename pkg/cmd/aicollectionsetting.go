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

var aiCollectionsSettingsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Replaces the collection's retrieval settings.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "retrieval",
			Usage:    "How documents are retrieved when searching the collection.",
			BodyPath: "retrieval",
		},
	},
	Action:          handleAICollectionsSettingsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"retrieval": {
		&requestflag.InnerFlag[string]{
			Name:       "retrieval.retrieval-type",
			Usage:      "Retrieval strategy. `vector` runs semantic similarity search; `hybrid` combines vector similarity with keyword matching; `keyword` runs lexical (BM25) matching. `keyword` is not accepted yet: setting it returns 422 `unsupported_retrieval_type`. A collection set to `hybrid` is accepted here but cannot be searched until hybrid execution ships.",
			InnerField: "retrieval_type",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "retrieval.top-k",
			Usage:      "Number of top results to retrieve (1–50).",
			InnerField: "top_k",
		},
	},
})

var aiCollectionsSettingsList = cli.Command{
	Name:    "list",
	Usage:   "Returns the retrieval settings for a collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
	},
	Action:          handleAICollectionsSettingsList,
	HideHelpCommand: true,
}

var aiCollectionsSettingsPatchAll = requestflag.WithInnerFlags(cli.Command{
	Name:    "patch-all",
	Usage:   "Partially updates the collection's retrieval settings.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "uuid",
			Required:  true,
			PathParam: "uuid",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "retrieval",
			Usage:    "How documents are retrieved when searching the collection.",
			BodyPath: "retrieval",
		},
	},
	Action:          handleAICollectionsSettingsPatchAll,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"retrieval": {
		&requestflag.InnerFlag[string]{
			Name:       "retrieval.retrieval-type",
			Usage:      "Retrieval strategy. `vector` runs semantic similarity search; `hybrid` combines vector similarity with keyword matching; `keyword` runs lexical (BM25) matching. `keyword` is not accepted yet: setting it returns 422 `unsupported_retrieval_type`. A collection set to `hybrid` is accepted here but cannot be searched until hybrid execution ships.",
			InnerField: "retrieval_type",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "retrieval.top-k",
			Usage:      "Number of top results to retrieve (1–50).",
			InnerField: "top_k",
		},
	},
})

func handleAICollectionsSettingsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AICollectionSettingNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Collections.Settings.New(
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
		Title:          "ai:collections:settings create",
		Transform:      transform,
	})
}

func handleAICollectionsSettingsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Collections.Settings.List(ctx, cmd.Value("uuid").(string), options...)
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
		Title:          "ai:collections:settings list",
		Transform:      transform,
	})
}

func handleAICollectionsSettingsPatchAll(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AICollectionSettingPatchAllParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Collections.Settings.PatchAll(
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
		Title:          "ai:collections:settings patch-all",
		Transform:      transform,
	})
}

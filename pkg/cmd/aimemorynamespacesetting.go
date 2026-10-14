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

var aiMemoryNamespacesSettingsList = cli.Command{
	Name:    "list",
	Usage:   "What is currently set for this namespace. `instructions: null` means none are\nset and summaries use the neutral default.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "namespace",
			Usage:     "The namespace. `default` exists for every organization.",
			Required:  true,
			PathParam: "namespace",
		},
	},
	Action:          handleAIMemoryNamespacesSettingsList,
	HideHelpCommand: true,
}

var aiMemoryNamespacesSettingsPatchAll = requestflag.WithInnerFlags(cli.Command{
	Name:    "patch-all",
	Usage:   "Only the fields you send are changed; anything omitted is left as it is, so `{}`\nchanges nothing. Sending `instructions: null`, or an empty or whitespace-only\nstring, clears them and returns summaries to the neutral default.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "namespace",
			Usage:     "The namespace. `default` exists for every organization.",
			Required:  true,
			PathParam: "namespace",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "summary",
			Usage:    "A partial update to a namespace's summary settings.\n\nOnly the fields present in the request are changed; the rest are left as\nthey are. Sending `instructions: null` (or empty) clears the instructions.",
			BodyPath: "summary",
		},
	},
	Action:          handleAIMemoryNamespacesSettingsPatchAll,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"summary": {
		&requestflag.InnerFlag[*string]{
			Name:       "summary.instructions",
			Usage:      "Replace the namespace's summary instructions. Null or empty clears them and returns to the neutral default. Omit the field to leave the current instructions unchanged.",
			InnerField: "instructions",
		},
	},
})

func handleAIMemoryNamespacesSettingsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("namespace") && len(unusedArgs) > 0 {
		cmd.Set("namespace", unusedArgs[0])
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
	_, err = client.AI.Memory.Namespaces.Settings.List(ctx, cmd.Value("namespace").(string), options...)
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
		Title:          "ai:memory:namespaces:settings list",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesSettingsPatchAll(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("namespace") && len(unusedArgs) > 0 {
		cmd.Set("namespace", unusedArgs[0])
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

	params := telnyx.AIMemoryNamespaceSettingPatchAllParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Settings.PatchAll(
		ctx,
		cmd.Value("namespace").(string),
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
		Title:          "ai:memory:namespaces:settings patch-all",
		Transform:      transform,
	})
}

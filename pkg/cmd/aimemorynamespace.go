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

var aiMemoryNamespacesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a namespace. An organization can have at most five, `default` among them\n— a sixth returns `403`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name for the new namespace, unique within your organization.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleAIMemoryNamespacesCreate,
	HideHelpCommand: true,
}

var aiMemoryNamespacesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Whether a write has finished. Both `ingest` and `remember` return an\n`operation_id`, and a memory is not recallable until its operation completes —\nextraction, embedding and consolidation all run first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "namespace",
			Required:  true,
			PathParam: "namespace",
		},
		&requestflag.Flag[string]{
			Name:      "operation-id",
			Required:  true,
			PathParam: "operation_id",
		},
	},
	Action:          handleAIMemoryNamespacesRetrieve,
	HideHelpCommand: true,
}

var aiMemoryNamespacesList = cli.Command{
	Name:            "list",
	Usage:           "Every namespace in your organization, `default` among them.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAIMemoryNamespacesList,
	HideHelpCommand: true,
}

var aiMemoryNamespacesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a namespace and every profile and memory in it. `default` cannot be\ndeleted. This cannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "namespace",
			Usage:     "The namespace to delete. `default` cannot be deleted.",
			Required:  true,
			PathParam: "namespace",
		},
	},
	Action:          handleAIMemoryNamespacesDelete,
	HideHelpCommand: true,
}

func handleAIMemoryNamespacesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AIMemoryNamespaceNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.New(ctx, params, options...)
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
		Title:          "ai:memory:namespaces create",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("operation-id") && len(unusedArgs) > 0 {
		cmd.Set("operation-id", unusedArgs[0])
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

	params := telnyx.AIMemoryNamespaceGetParams{
		Namespace: cmd.Value("namespace").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Get(
		ctx,
		cmd.Value("operation-id").(string),
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
		Title:          "ai:memory:namespaces retrieve",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Memory.Namespaces.List(ctx, options...)
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
		Title:          "ai:memory:namespaces list",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.AI.Memory.Namespaces.Delete(ctx, cmd.Value("namespace").(string), options...)
}

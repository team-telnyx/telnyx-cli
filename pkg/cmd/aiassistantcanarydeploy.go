// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/stainless-sdks/telnyx-go/v4"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiAssistantsCanaryDeploysCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Endpoint to create a canary deploy configuration for an assistant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "assistant-id",
			Required:  true,
			PathParam: "assistant_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "rule",
			BodyPath: "rules",
		},
	},
	Action:          handleAIAssistantsCanaryDeploysCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"rule": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "rule.serve",
			Usage:      "What a rule serves when matched.\n\nExactly one of:\n- ``version_id`` — serve a specific version\n- ``rollout`` — weighted random across versions; weights must sum to\n  less than 100, with the leftover routing to the main version",
			InnerField: "serve",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "rule.match",
			InnerField: "match",
		},
	},
})

var aiAssistantsCanaryDeploysRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Endpoint to get a canary deploy configuration for an assistant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "assistant-id",
			Required:  true,
			PathParam: "assistant_id",
		},
	},
	Action:          handleAIAssistantsCanaryDeploysRetrieve,
	HideHelpCommand: true,
}

var aiAssistantsCanaryDeploysUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Endpoint to update a canary deploy configuration for an assistant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "assistant-id",
			Required:  true,
			PathParam: "assistant_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "rule",
			BodyPath: "rules",
		},
	},
	Action:          handleAIAssistantsCanaryDeploysUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"rule": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "rule.serve",
			Usage:      "What a rule serves when matched.\n\nExactly one of:\n- ``version_id`` — serve a specific version\n- ``rollout`` — weighted random across versions; weights must sum to\n  less than 100, with the leftover routing to the main version",
			InnerField: "serve",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "rule.match",
			InnerField: "match",
		},
	},
})

var aiAssistantsCanaryDeploysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Endpoint to delete a canary deploy configuration for an assistant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "assistant-id",
			Required:  true,
			PathParam: "assistant_id",
		},
	},
	Action:          handleAIAssistantsCanaryDeploysDelete,
	HideHelpCommand: true,
}

func handleAIAssistantsCanaryDeploysCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
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

	params := telnyx.AIAssistantCanaryDeployNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Assistants.CanaryDeploys.New(
		ctx,
		cmd.Value("assistant-id").(string),
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
		Title:          "ai:assistants:canary-deploys create",
		Transform:      transform,
	})
}

func handleAIAssistantsCanaryDeploysRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
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
	_, err = client.AI.Assistants.CanaryDeploys.Get(ctx, cmd.Value("assistant-id").(string), options...)
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
		Title:          "ai:assistants:canary-deploys retrieve",
		Transform:      transform,
	})
}

func handleAIAssistantsCanaryDeploysUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
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

	params := telnyx.AIAssistantCanaryDeployUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Assistants.CanaryDeploys.Update(
		ctx,
		cmd.Value("assistant-id").(string),
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
		Title:          "ai:assistants:canary-deploys update",
		Transform:      transform,
	})
}

func handleAIAssistantsCanaryDeploysDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
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

	return client.AI.Assistants.CanaryDeploys.Delete(ctx, cmd.Value("assistant-id").(string), options...)
}

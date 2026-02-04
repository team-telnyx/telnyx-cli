// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiAssistantsCanaryDeploysCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Endpoint to create a canary deploy configuration for an assistant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "version",
			Usage:    "List of version configurations",
			Required: true,
			BodyPath: "versions",
		},
	},
	Action:          handleAIAssistantsCanaryDeploysCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"version": {
		&requestflag.InnerFlag[float64]{
			Name:       "version.percentage",
			Usage:      "Percentage of traffic for this version [1-99]",
			InnerField: "percentage",
		},
		&requestflag.InnerFlag[string]{
			Name:       "version.version-id",
			Usage:      "Version ID string that references assistant_versions.version_id",
			InnerField: "version_id",
		},
	},
})

var aiAssistantsCanaryDeploysRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Endpoint to get a canary deploy configuration for an assistant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
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
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "version",
			Usage:    "List of version configurations",
			Required: true,
			BodyPath: "versions",
		},
	},
	Action:          handleAIAssistantsCanaryDeploysUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"version": {
		&requestflag.InnerFlag[float64]{
			Name:       "version.percentage",
			Usage:      "Percentage of traffic for this version [1-99]",
			InnerField: "percentage",
		},
		&requestflag.InnerFlag[string]{
			Name:       "version.version-id",
			Usage:      "Version ID string that references assistant_versions.version_id",
			InnerField: "version_id",
		},
	},
})

var aiAssistantsCanaryDeploysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Endpoint to delete a canary deploy configuration for an assistant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
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

	params := telnyx.AIAssistantCanaryDeployNewParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:canary-deploys create", obj, format, transform)
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:canary-deploys retrieve", obj, format, transform)
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

	params := telnyx.AIAssistantCanaryDeployUpdateParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:canary-deploys update", obj, format, transform)
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

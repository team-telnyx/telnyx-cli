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

var aiMissionsToolsCreateTool = cli.Command{
	Name:    "create-tool",
	Usage:   "Create a new tool for a mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsToolsCreateTool,
	HideHelpCommand: true,
}

var aiMissionsToolsDeleteTool = cli.Command{
	Name:    "delete-tool",
	Usage:   "Delete a tool from a mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsToolsDeleteTool,
	HideHelpCommand: true,
}

var aiMissionsToolsGetTool = cli.Command{
	Name:    "get-tool",
	Usage:   "Get a specific tool by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsToolsGetTool,
	HideHelpCommand: true,
}

var aiMissionsToolsListTools = cli.Command{
	Name:    "list-tools",
	Usage:   "List all tools for a mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsToolsListTools,
	HideHelpCommand: true,
}

var aiMissionsToolsUpdateTool = cli.Command{
	Name:    "update-tool",
	Usage:   "Update a tool definition",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsToolsUpdateTool,
	HideHelpCommand: true,
}

func handleAIMissionsToolsCreateTool(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mission-id") && len(unusedArgs) > 0 {
		cmd.Set("mission-id", unusedArgs[0])
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
	_, err = client.AI.Missions.Tools.NewTool(ctx, cmd.Value("mission-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:tools create-tool", obj, format, transform)
}

func handleAIMissionsToolsDeleteTool(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionToolDeleteToolParams{
		MissionID: cmd.Value("mission-id").(string),
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

	return client.AI.Missions.Tools.DeleteTool(
		ctx,
		cmd.Value("tool-id").(string),
		params,
		options...,
	)
}

func handleAIMissionsToolsGetTool(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionToolGetToolParams{
		MissionID: cmd.Value("mission-id").(string),
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
	_, err = client.AI.Missions.Tools.GetTool(
		ctx,
		cmd.Value("tool-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:tools get-tool", obj, format, transform)
}

func handleAIMissionsToolsListTools(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mission-id") && len(unusedArgs) > 0 {
		cmd.Set("mission-id", unusedArgs[0])
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
	_, err = client.AI.Missions.Tools.ListTools(ctx, cmd.Value("mission-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:tools list-tools", obj, format, transform)
}

func handleAIMissionsToolsUpdateTool(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionToolUpdateToolParams{
		MissionID: cmd.Value("mission-id").(string),
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
	_, err = client.AI.Missions.Tools.UpdateTool(
		ctx,
		cmd.Value("tool-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:tools update-tool", obj, format, transform)
}

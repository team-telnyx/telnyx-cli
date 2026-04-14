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

var aiMissionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new mission definition",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "execution-mode",
			Usage:    `Allowed values: "external", "managed".`,
			Default:  "external",
			BodyPath: "execution_mode",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
	},
	Action:          handleAIMissionsCreate,
	HideHelpCommand: true,
}

var aiMissionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a mission by ID (includes tools, knowledge_bases, mcp_servers)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsRetrieve,
	HideHelpCommand: true,
}

var aiMissionsList = cli.Command{
	Name:    "list",
	Usage:   "List all missions for the organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (1-based)",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of items per page",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAIMissionsList,
	HideHelpCommand: true,
}

var aiMissionsCloneMission = cli.Command{
	Name:    "clone-mission",
	Usage:   "Clone an existing mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsCloneMission,
	HideHelpCommand: true,
}

var aiMissionsDeleteMission = cli.Command{
	Name:    "delete-mission",
	Usage:   "Delete a mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsDeleteMission,
	HideHelpCommand: true,
}

var aiMissionsListEvents = cli.Command{
	Name:    "list-events",
	Usage:   "List recent events across all missions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (1-based)",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of items per page",
			Default:   50,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			QueryPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAIMissionsListEvents,
	HideHelpCommand: true,
}

var aiMissionsUpdateMission = cli.Command{
	Name:    "update-mission",
	Usage:   "Update a mission definition",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "execution-mode",
			Usage:    `Allowed values: "external", "managed".`,
			BodyPath: "execution_mode",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleAIMissionsUpdateMission,
	HideHelpCommand: true,
}

func handleAIMissionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionNewParams{}

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
	_, err = client.AI.Missions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "ai:missions create", obj, format, explicitFormat, transform)
}

func handleAIMissionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Missions.Get(ctx, cmd.Value("mission-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "ai:missions retrieve", obj, format, explicitFormat, transform)
}

func handleAIMissionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionListParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Missions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "ai:missions list", obj, format, explicitFormat, transform)
	} else {
		iter := client.AI.Missions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "ai:missions list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleAIMissionsCloneMission(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Missions.CloneMission(ctx, cmd.Value("mission-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "ai:missions clone-mission", obj, format, explicitFormat, transform)
}

func handleAIMissionsDeleteMission(ctx context.Context, cmd *cli.Command) error {
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

	return client.AI.Missions.DeleteMission(ctx, cmd.Value("mission-id").(string), options...)
}

func handleAIMissionsListEvents(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionListEventsParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Missions.ListEvents(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "ai:missions list-events", obj, format, explicitFormat, transform)
	} else {
		iter := client.AI.Missions.ListEventsAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "ai:missions list-events", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleAIMissionsUpdateMission(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mission-id") && len(unusedArgs) > 0 {
		cmd.Set("mission-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionUpdateMissionParams{}

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
	_, err = client.AI.Missions.UpdateMission(
		ctx,
		cmd.Value("mission-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "ai:missions update-mission", obj, format, explicitFormat, transform)
}

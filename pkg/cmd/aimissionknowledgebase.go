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

var aiMissionsKnowledgeBasesCreateKnowledgeBase = cli.Command{
	Name:    "create-knowledge-base",
	Usage:   "Create a new knowledge base for a mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsKnowledgeBasesCreateKnowledgeBase,
	HideHelpCommand: true,
}

var aiMissionsKnowledgeBasesDeleteKnowledgeBase = cli.Command{
	Name:    "delete-knowledge-base",
	Usage:   "Delete a knowledge base from a mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "knowledge-base-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsKnowledgeBasesDeleteKnowledgeBase,
	HideHelpCommand: true,
}

var aiMissionsKnowledgeBasesGetKnowledgeBase = cli.Command{
	Name:    "get-knowledge-base",
	Usage:   "Get a specific knowledge base by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "knowledge-base-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsKnowledgeBasesGetKnowledgeBase,
	HideHelpCommand: true,
}

var aiMissionsKnowledgeBasesListKnowledgeBases = cli.Command{
	Name:    "list-knowledge-bases",
	Usage:   "List all knowledge bases for a mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsKnowledgeBasesListKnowledgeBases,
	HideHelpCommand: true,
}

var aiMissionsKnowledgeBasesUpdateKnowledgeBase = cli.Command{
	Name:    "update-knowledge-base",
	Usage:   "Update a knowledge base definition",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "knowledge-base-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsKnowledgeBasesUpdateKnowledgeBase,
	HideHelpCommand: true,
}

func handleAIMissionsKnowledgeBasesCreateKnowledgeBase(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Missions.KnowledgeBases.NewKnowledgeBase(ctx, cmd.Value("mission-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:knowledge-bases create-knowledge-base", obj, format, transform)
}

func handleAIMissionsKnowledgeBasesDeleteKnowledgeBase(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("knowledge-base-id") && len(unusedArgs) > 0 {
		cmd.Set("knowledge-base-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionKnowledgeBaseDeleteKnowledgeBaseParams{
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

	return client.AI.Missions.KnowledgeBases.DeleteKnowledgeBase(
		ctx,
		cmd.Value("knowledge-base-id").(string),
		params,
		options...,
	)
}

func handleAIMissionsKnowledgeBasesGetKnowledgeBase(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("knowledge-base-id") && len(unusedArgs) > 0 {
		cmd.Set("knowledge-base-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionKnowledgeBaseGetKnowledgeBaseParams{
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
	_, err = client.AI.Missions.KnowledgeBases.GetKnowledgeBase(
		ctx,
		cmd.Value("knowledge-base-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:knowledge-bases get-knowledge-base", obj, format, transform)
}

func handleAIMissionsKnowledgeBasesListKnowledgeBases(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Missions.KnowledgeBases.ListKnowledgeBases(ctx, cmd.Value("mission-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:knowledge-bases list-knowledge-bases", obj, format, transform)
}

func handleAIMissionsKnowledgeBasesUpdateKnowledgeBase(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("knowledge-base-id") && len(unusedArgs) > 0 {
		cmd.Set("knowledge-base-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionKnowledgeBaseUpdateKnowledgeBaseParams{
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
	_, err = client.AI.Missions.KnowledgeBases.UpdateKnowledgeBase(
		ctx,
		cmd.Value("knowledge-base-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:knowledge-bases update-knowledge-base", obj, format, transform)
}

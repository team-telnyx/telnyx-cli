// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/urfave/cli/v3"
)

var aiConversationsInsightGroupsInsightsAssign = cli.Command{
	Name:    "assign",
	Usage:   "Assign an insight to a group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "group-id",
			Usage:     "The ID of the insight group",
			Required:  true,
			PathParam: "group_id",
		},
		&requestflag.Flag[string]{
			Name:      "insight-id",
			Usage:     "The ID of the insight",
			Required:  true,
			PathParam: "insight_id",
		},
	},
	Action:          handleAIConversationsInsightGroupsInsightsAssign,
	HideHelpCommand: true,
}

var aiConversationsInsightGroupsInsightsDeleteUnassign = cli.Command{
	Name:    "delete-unassign",
	Usage:   "Remove an insight from a group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "group-id",
			Usage:     "The ID of the insight group",
			Required:  true,
			PathParam: "group_id",
		},
		&requestflag.Flag[string]{
			Name:      "insight-id",
			Usage:     "The ID of the insight",
			Required:  true,
			PathParam: "insight_id",
		},
	},
	Action:          handleAIConversationsInsightGroupsInsightsDeleteUnassign,
	HideHelpCommand: true,
}

func handleAIConversationsInsightGroupsInsightsAssign(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("insight-id") && len(unusedArgs) > 0 {
		cmd.Set("insight-id", unusedArgs[0])
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

	params := telnyx.AIConversationInsightGroupInsightAssignParams{
		GroupID: cmd.Value("group-id").(string),
	}

	return client.AI.Conversations.InsightGroups.Insights.Assign(
		ctx,
		cmd.Value("insight-id").(string),
		params,
		options...,
	)
}

func handleAIConversationsInsightGroupsInsightsDeleteUnassign(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("insight-id") && len(unusedArgs) > 0 {
		cmd.Set("insight-id", unusedArgs[0])
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

	params := telnyx.AIConversationInsightGroupInsightDeleteUnassignParams{
		GroupID: cmd.Value("group-id").(string),
	}

	return client.AI.Conversations.InsightGroups.Insights.DeleteUnassign(
		ctx,
		cmd.Value("insight-id").(string),
		params,
		options...,
	)
}

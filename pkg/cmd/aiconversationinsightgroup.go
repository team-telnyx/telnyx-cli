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

var aiConversationsInsightGroupsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get insight group by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group-id",
			Usage:    "The ID of the insight group",
			Required: true,
		},
	},
	Action:          handleAIConversationsInsightGroupsRetrieve,
	HideHelpCommand: true,
}

var aiConversationsInsightGroupsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an insight template group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group-id",
			Usage:    "The ID of the insight group",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "webhook",
			BodyPath: "webhook",
		},
	},
	Action:          handleAIConversationsInsightGroupsUpdate,
	HideHelpCommand: true,
}

var aiConversationsInsightGroupsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete insight group by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group-id",
			Usage:    "The ID of the insight group",
			Required: true,
		},
	},
	Action:          handleAIConversationsInsightGroupsDelete,
	HideHelpCommand: true,
}

var aiConversationsInsightGroupsInsightGroups = cli.Command{
	Name:    "insight-groups",
	Usage:   "Create a new insight group",
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
			Name:     "webhook",
			Default:  "",
			BodyPath: "webhook",
		},
	},
	Action:          handleAIConversationsInsightGroupsInsightGroups,
	HideHelpCommand: true,
}

var aiConversationsInsightGroupsRetrieveInsightGroups = cli.Command{
	Name:    "retrieve-insight-groups",
	Usage:   "Get all insight groups",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleAIConversationsInsightGroupsRetrieveInsightGroups,
	HideHelpCommand: true,
}

func handleAIConversationsInsightGroupsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group-id") && len(unusedArgs) > 0 {
		cmd.Set("group-id", unusedArgs[0])
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
	_, err = client.AI.Conversations.InsightGroups.Get(ctx, cmd.Value("group-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations:insight-groups retrieve", obj, format, transform)
}

func handleAIConversationsInsightGroupsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group-id") && len(unusedArgs) > 0 {
		cmd.Set("group-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationInsightGroupUpdateParams{}

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
	_, err = client.AI.Conversations.InsightGroups.Update(
		ctx,
		cmd.Value("group-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations:insight-groups update", obj, format, transform)
}

func handleAIConversationsInsightGroupsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group-id") && len(unusedArgs) > 0 {
		cmd.Set("group-id", unusedArgs[0])
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

	return client.AI.Conversations.InsightGroups.Delete(ctx, cmd.Value("group-id").(string), options...)
}

func handleAIConversationsInsightGroupsInsightGroups(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationInsightGroupInsightGroupsParams{}

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
	_, err = client.AI.Conversations.InsightGroups.InsightGroups(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations:insight-groups insight-groups", obj, format, transform)
}

func handleAIConversationsInsightGroupsRetrieveInsightGroups(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationInsightGroupGetInsightGroupsParams{}

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
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Conversations.InsightGroups.GetInsightGroups(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:conversations:insight-groups retrieve-insight-groups", obj, format, transform)
	} else {
		iter := client.AI.Conversations.InsightGroups.GetInsightGroupsAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "ai:conversations:insight-groups retrieve-insight-groups", iter, format, transform)
	}
}

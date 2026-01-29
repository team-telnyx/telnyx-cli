// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiConversationsInsightsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new insight",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "instructions",
			Required: true,
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "json-schema",
			Usage:    "If specified, the output will follow the JSON schema.",
			BodyPath: "json_schema",
		},
		&requestflag.Flag[string]{
			Name:     "webhook",
			Default:  "",
			BodyPath: "webhook",
		},
	},
	Action:          handleAIConversationsInsightsCreate,
	HideHelpCommand: true,
}

var aiConversationsInsightsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get insight by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "insight-id",
			Usage:    "The ID of the insight",
			Required: true,
		},
	},
	Action:          handleAIConversationsInsightsRetrieve,
	HideHelpCommand: true,
}

var aiConversationsInsightsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an insight template",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "insight-id",
			Usage:    "The ID of the insight",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[any]{
			Name:     "json-schema",
			BodyPath: "json_schema",
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
	Action:          handleAIConversationsInsightsUpdate,
	HideHelpCommand: true,
}

var aiConversationsInsightsList = cli.Command{
	Name:    "list",
	Usage:   "Get all insights",
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
	Action:          handleAIConversationsInsightsList,
	HideHelpCommand: true,
}

var aiConversationsInsightsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete insight by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "insight-id",
			Usage:    "The ID of the insight",
			Required: true,
		},
	},
	Action:          handleAIConversationsInsightsDelete,
	HideHelpCommand: true,
}

func handleAIConversationsInsightsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationInsightNewParams{}

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
	_, err = client.AI.Conversations.Insights.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations:insights create", obj, format, transform)
}

func handleAIConversationsInsightsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Conversations.Insights.Get(ctx, cmd.Value("insight-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations:insights retrieve", obj, format, transform)
}

func handleAIConversationsInsightsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("insight-id") && len(unusedArgs) > 0 {
		cmd.Set("insight-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationInsightUpdateParams{}

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
	_, err = client.AI.Conversations.Insights.Update(
		ctx,
		cmd.Value("insight-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations:insights update", obj, format, transform)
}

func handleAIConversationsInsightsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationInsightListParams{}

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
		_, err = client.AI.Conversations.Insights.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:conversations:insights list", obj, format, transform)
	} else {
		iter := client.AI.Conversations.Insights.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "ai:conversations:insights list", iter, format, transform)
	}
}

func handleAIConversationsInsightsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.AI.Conversations.Insights.Delete(ctx, cmd.Value("insight-id").(string), options...)
}

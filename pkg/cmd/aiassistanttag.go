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

var aiAssistantsTagsList = cli.Command{
	Name:            "list",
	Usage:           "Get All Tags",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAIAssistantsTagsList,
	HideHelpCommand: true,
}

var aiAssistantsTagsAdd = cli.Command{
	Name:    "add",
	Usage:   "Add Assistant Tag",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tag",
			Required: true,
			BodyPath: "tag",
		},
	},
	Action:          handleAIAssistantsTagsAdd,
	HideHelpCommand: true,
}

var aiAssistantsTagsRemove = cli.Command{
	Name:    "remove",
	Usage:   "Remove Assistant Tag",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tag",
			Required: true,
		},
	},
	Action:          handleAIAssistantsTagsRemove,
	HideHelpCommand: true,
}

func handleAIAssistantsTagsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Assistants.Tags.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:tags list", obj, format, transform)
}

func handleAIAssistantsTagsAdd(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTagAddParams{}

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
	_, err = client.AI.Assistants.Tags.Add(
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
	return ShowJSON(os.Stdout, "ai:assistants:tags add", obj, format, transform)
}

func handleAIAssistantsTagsRemove(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tag") && len(unusedArgs) > 0 {
		cmd.Set("tag", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTagRemoveParams{
		AssistantID: cmd.Value("assistant-id").(string),
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
	_, err = client.AI.Assistants.Tags.Remove(
		ctx,
		cmd.Value("tag").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:tags remove", obj, format, transform)
}

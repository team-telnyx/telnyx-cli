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

var aiToolsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create Tool",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "display-name",
			Required: true,
			BodyPath: "display_name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "function",
			BodyPath: "function",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "handoff",
			BodyPath: "handoff",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "invite",
			BodyPath: "invite",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "retrieval",
			BodyPath: "retrieval",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Default:  5000,
			BodyPath: "timeout_ms",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "webhook",
			BodyPath: "webhook",
		},
	},
	Action:          handleAIToolsCreate,
	HideHelpCommand: true,
}

var aiToolsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get Tool",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
	},
	Action:          handleAIToolsRetrieve,
	HideHelpCommand: true,
}

var aiToolsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update Tool",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			BodyPath: "display_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "function",
			BodyPath: "function",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "handoff",
			BodyPath: "handoff",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "invite",
			BodyPath: "invite",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "retrieval",
			BodyPath: "retrieval",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			BodyPath: "timeout_ms",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			BodyPath: "type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "webhook",
			BodyPath: "webhook",
		},
	},
	Action:          handleAIToolsUpdate,
	HideHelpCommand: true,
}

var aiToolsList = cli.Command{
	Name:    "list",
	Usage:   "List Tools",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-name",
			QueryPath: "filter[name]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-type",
			QueryPath: "filter[type]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAIToolsList,
	HideHelpCommand: true,
}

var aiToolsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Tool",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tool-id",
			Required: true,
		},
	},
	Action:          handleAIToolsDelete,
	HideHelpCommand: true,
}

func handleAIToolsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIToolNewParams{}

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
	_, err = client.AI.Tools.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:tools create", obj, format, transform)
}

func handleAIToolsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
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
	_, err = client.AI.Tools.Get(ctx, cmd.Value("tool-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:tools retrieve", obj, format, transform)
}

func handleAIToolsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIToolUpdateParams{}

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
	_, err = client.AI.Tools.Update(
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
	return ShowJSON(os.Stdout, "ai:tools update", obj, format, transform)
}

func handleAIToolsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIToolListParams{}

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
		_, err = client.AI.Tools.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:tools list", obj, format, transform)
	} else {
		iter := client.AI.Tools.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "ai:tools list", iter, format, transform, maxItems)
	}
}

func handleAIToolsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-id", unusedArgs[0])
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
	_, err = client.AI.Tools.Delete(ctx, cmd.Value("tool-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:tools delete", obj, format, transform)
}

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

var aiMcpServersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new MCP server.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[any]{
			Name:     "allowed-tool",
			BodyPath: "allowed_tools",
		},
		&requestflag.Flag[any]{
			Name:     "api-key-ref",
			BodyPath: "api_key_ref",
		},
	},
	Action:          handleAIMcpServersCreate,
	HideHelpCommand: true,
}

var aiMcpServersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve details for a specific MCP server.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mcp-server-id",
			Required: true,
		},
	},
	Action:          handleAIMcpServersRetrieve,
	HideHelpCommand: true,
}

var aiMcpServersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an existing MCP server.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mcp-server-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[any]{
			Name:     "allowed-tool",
			BodyPath: "allowed_tools",
		},
		&requestflag.Flag[any]{
			Name:     "api-key-ref",
			BodyPath: "api_key_ref",
		},
		&requestflag.Flag[any]{
			Name:     "created-at",
			BodyPath: "created_at",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			BodyPath: "url",
		},
	},
	Action:          handleAIMcpServersUpdate,
	HideHelpCommand: true,
}

var aiMcpServersList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of MCP servers.",
	Suggest: true,
	Flags: []cli.Flag{
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
		&requestflag.Flag[string]{
			Name:      "type",
			QueryPath: "type",
		},
		&requestflag.Flag[string]{
			Name:      "url",
			QueryPath: "url",
		},
	},
	Action:          handleAIMcpServersList,
	HideHelpCommand: true,
}

var aiMcpServersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a specific MCP server.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mcp-server-id",
			Required: true,
		},
	},
	Action:          handleAIMcpServersDelete,
	HideHelpCommand: true,
}

func handleAIMcpServersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMcpServerNewParams{}

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
	_, err = client.AI.McpServers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:mcp-servers create", obj, format, transform)
}

func handleAIMcpServersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mcp-server-id") && len(unusedArgs) > 0 {
		cmd.Set("mcp-server-id", unusedArgs[0])
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
	_, err = client.AI.McpServers.Get(ctx, cmd.Value("mcp-server-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:mcp-servers retrieve", obj, format, transform)
}

func handleAIMcpServersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mcp-server-id") && len(unusedArgs) > 0 {
		cmd.Set("mcp-server-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMcpServerUpdateParams{}

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
	_, err = client.AI.McpServers.Update(
		ctx,
		cmd.Value("mcp-server-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:mcp-servers update", obj, format, transform)
}

func handleAIMcpServersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMcpServerListParams{}

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
		_, err = client.AI.McpServers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:mcp-servers list", obj, format, transform)
	} else {
		iter := client.AI.McpServers.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "ai:mcp-servers list", iter, format, transform)
	}
}

func handleAIMcpServersDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mcp-server-id") && len(unusedArgs) > 0 {
		cmd.Set("mcp-server-id", unusedArgs[0])
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

	return client.AI.McpServers.Delete(ctx, cmd.Value("mcp-server-id").(string), options...)
}

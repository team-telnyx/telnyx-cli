// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var agentGatewaysMcpServersCreate = cli.Command{
	Name:    "create",
	Usage:   "Adds an MCP server. The instance serves its tools at `{gateway_url}/mcp`, each\nnamed with the server's `name` and a hyphen, to the keys granted the server and,\nwith `open_to_every_key`, to every key. The instance must be `ready`; otherwise\nthe request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name for the server, unique on the instance across MCP servers and APIs, regardless of case. Up to 100 letters, digits, `_`, and `.`. The gateway prefixes the server's tools with it and a hyphen.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The server's MCP endpoint, `http` or `https`.",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "auth",
			Usage:    "How the gateway authenticates to the server. Replaced whole on an update; `{\"type\": \"none\"}` removes the credential.",
			BodyPath: "auth",
		},
		&requestflag.Flag[[]string]{
			Name:     "forwarded-header",
			Usage:    "Headers copied from the agent's request, by name.",
			BodyPath: "forwarded_headers",
		},
		&requestflag.Flag[bool]{
			Name:     "open-to-every-key",
			Usage:    "`true` lets every key on the instance use the server, granted or not.",
			Default:  false,
			BodyPath: "open_to_every_key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "static-headers",
			Usage:    "Headers sent on every request.",
			BodyPath: "static_headers",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool",
			Usage:    "The tools the server exposes to every key, by name. `[]` exposes every tool.",
			BodyPath: "tools",
		},
		&requestflag.Flag[string]{
			Name:     "transport",
			Usage:    "`http` for Streamable HTTP, or `sse` for Server-Sent Events.",
			Default:  "http",
			BodyPath: "transport",
		},
	},
	Action:          handleAgentGatewaysMcpServersCreate,
	HideHelpCommand: true,
}

var agentGatewaysMcpServersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Changes an MCP server's settings. A field left out stays as it is. `auth` is\nreplaced whole, `tools: []` exposes every tool, and `static_headers: {}` clears\nthe headers. The instance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "server-id",
			Required:  true,
			PathParam: "server_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "auth",
			Usage:    "How the gateway authenticates to the server. Replaced whole on an update; `{\"type\": \"none\"}` removes the credential.",
			BodyPath: "auth",
		},
		&requestflag.Flag[[]string]{
			Name:     "forwarded-header",
			Usage:    "Headers copied from the agent's request, by name.",
			BodyPath: "forwarded_headers",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name for the server, unique on the instance across MCP servers and APIs, regardless of case. Letters, digits, `_`, and `.`. The gateway prefixes the server's tools with it and a hyphen.",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "open-to-every-key",
			Usage:    "`true` lets every key on the instance use the server, granted or not.",
			Default:  false,
			BodyPath: "open_to_every_key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "static-headers",
			Usage:    "Headers sent on every request.",
			BodyPath: "static_headers",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool",
			Usage:    "The tools the server exposes to every key, by name. `[]` exposes every tool.",
			BodyPath: "tools",
		},
		&requestflag.Flag[string]{
			Name:     "transport",
			Usage:    "`http` for Streamable HTTP, or `sse` for Server-Sent Events.",
			Default:  "http",
			BodyPath: "transport",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The server's MCP endpoint, `http` or `https`.",
			BodyPath: "url",
		},
	},
	Action:          handleAgentGatewaysMcpServersUpdate,
	HideHelpCommand: true,
}

var agentGatewaysMcpServersList = cli.Command{
	Name:    "list",
	Usage:   "Lists the MCP servers the instance serves at `{gateway_url}/mcp`, without their\ncredentials. The instance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
	},
	Action:          handleAgentGatewaysMcpServersList,
	HideHelpCommand: true,
}

var agentGatewaysMcpServersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Removes an MCP server and returns it as it was. Its `id` stays in the tool\npermissions that name it until they are next set. The instance must be `ready`;\notherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "server-id",
			Required:  true,
			PathParam: "server_id",
		},
	},
	Action:          handleAgentGatewaysMcpServersDelete,
	HideHelpCommand: true,
}

func handleAgentGatewaysMcpServersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instance-id") && len(unusedArgs) > 0 {
		cmd.Set("instance-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.AgentGatewayMcpServerNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.McpServers.New(
		ctx,
		cmd.Value("instance-id").(string),
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent-gateways:mcp-servers create",
		Transform:      transform,
	})
}

func handleAgentGatewaysMcpServersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("server-id") && len(unusedArgs) > 0 {
		cmd.Set("server-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.AgentGatewayMcpServerUpdateParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.McpServers.Update(
		ctx,
		cmd.Value("server-id").(string),
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent-gateways:mcp-servers update",
		Transform:      transform,
	})
}

func handleAgentGatewaysMcpServersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instance-id") && len(unusedArgs) > 0 {
		cmd.Set("instance-id", unusedArgs[0])
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
	_, err = client.AgentGateways.McpServers.List(ctx, cmd.Value("instance-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent-gateways:mcp-servers list",
		Transform:      transform,
	})
}

func handleAgentGatewaysMcpServersDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("server-id") && len(unusedArgs) > 0 {
		cmd.Set("server-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayMcpServerDeleteParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.McpServers.Delete(
		ctx,
		cmd.Value("server-id").(string),
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent-gateways:mcp-servers delete",
		Transform:      transform,
	})
}

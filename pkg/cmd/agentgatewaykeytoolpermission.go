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

var agentGatewaysKeysToolPermissionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Replaces the grant whole and returns it. Every server must be one of the\ninstance's MCP servers or APIs, named once. A key reaches the servers granted to\nit or to its team, and the servers open to every key; when both grant, only the\nservers in both. The instance must be `ready`; otherwise the request answers\n`409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "key-id",
			Required:  true,
			PathParam: "key_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "server",
			Usage:    "The servers granted, each with its tools. `[]` grants none.",
			Required: true,
			BodyPath: "servers",
		},
	},
	Action:          handleAgentGatewaysKeysToolPermissionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"server": {
		&requestflag.InnerFlag[string]{
			Name:       "server.id",
			Usage:      "An MCP server's or API's `id` on the instance.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "server.tools",
			Usage:      "The tools the grant covers, by name without the server prefix. `[]` covers every tool the server exposes.",
			InnerField: "tools",
		},
	},
})

var agentGatewaysKeysToolPermissionsList = cli.Command{
	Name:    "list",
	Usage:   "Returns the MCP servers and APIs the key may use, and the tools within each.\n`servers: []` means none. The instance must be `ready`; otherwise the request\nanswers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "key-id",
			Required:  true,
			PathParam: "key_id",
		},
	},
	Action:          handleAgentGatewaysKeysToolPermissionsList,
	HideHelpCommand: true,
}

func handleAgentGatewaysKeysToolPermissionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key-id") && len(unusedArgs) > 0 {
		cmd.Set("key-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayKeyToolPermissionNewParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Keys.ToolPermissions.New(
		ctx,
		cmd.Value("key-id").(string),
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
		Title:          "agent-gateways:keys:tool-permissions create",
		Transform:      transform,
	})
}

func handleAgentGatewaysKeysToolPermissionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key-id") && len(unusedArgs) > 0 {
		cmd.Set("key-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayKeyToolPermissionListParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Keys.ToolPermissions.List(
		ctx,
		cmd.Value("key-id").(string),
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
		Title:          "agent-gateways:keys:tool-permissions list",
		Transform:      transform,
	})
}

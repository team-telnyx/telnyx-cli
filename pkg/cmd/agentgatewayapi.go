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

var agentGatewaysAPIsCreate = cli.Command{
	Name:    "create",
	Usage:   "Adds a REST API by its OpenAPI description. Each operation becomes a tool at\n`{gateway_url}/mcp`, named with the API's `name`, a hyphen, and the lowercased\n`operationId`. A description the gateway cannot fetch yields an API with no\ntools. The instance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name for the API, unique on the instance across MCP servers and APIs, regardless of case. Up to 100 letters, digits, `_`, and `.`. The gateway prefixes the API's tools with it and a hyphen.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "spec-url",
			Usage:    "The URL of the API's OpenAPI description: OpenAPI 3.0, OpenAPI 3.1, or Swagger 2.0.",
			Required: true,
			BodyPath: "spec_url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "auth",
			Usage:    "How the gateway authenticates to the API. AWS Signature Version 4 is for MCP servers only. Replaced whole on an update.",
			BodyPath: "auth",
		},
		&requestflag.Flag[*string]{
			Name:     "base-url",
			Usage:    "Where the gateway sends the API's requests. Without it, the gateway takes the description's first `servers` entry.",
			BodyPath: "base_url",
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
			Usage:    "The operations the API exposes as tools, by tool name. `[]` exposes every operation.",
			BodyPath: "tools",
		},
	},
	Action:          handleAgentGatewaysAPIsCreate,
	HideHelpCommand: true,
}

var agentGatewaysAPIsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Changes an API's settings. A field left out stays as it is, `\"base_url\": null`\nclears the base URL, and `auth` is replaced whole. The instance must be `ready`;\notherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "api-id",
			Required:  true,
			PathParam: "api_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "auth",
			Usage:    "How the gateway authenticates to the API. AWS Signature Version 4 is for MCP servers only. Replaced whole on an update.",
			BodyPath: "auth",
		},
		&requestflag.Flag[*string]{
			Name:     "base-url",
			Usage:    "Where the gateway sends the API's requests. Without it, the gateway takes the description's first `servers` entry.",
			BodyPath: "base_url",
		},
		&requestflag.Flag[[]string]{
			Name:     "forwarded-header",
			Usage:    "Headers copied from the agent's request, by name.",
			BodyPath: "forwarded_headers",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name for the API, unique on the instance across MCP servers and APIs, regardless of case. Letters, digits, `_`, and `.`. The gateway prefixes the API's tools with it and a hyphen.",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "open-to-every-key",
			Usage:    "`true` lets every key on the instance use the server, granted or not.",
			Default:  false,
			BodyPath: "open_to_every_key",
		},
		&requestflag.Flag[string]{
			Name:     "spec-url",
			Usage:    "The URL of the API's OpenAPI description: OpenAPI 3.0, OpenAPI 3.1, or Swagger 2.0.",
			BodyPath: "spec_url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "static-headers",
			Usage:    "Headers sent on every request.",
			BodyPath: "static_headers",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool",
			Usage:    "The operations the API exposes as tools, by tool name. `[]` exposes every operation.",
			BodyPath: "tools",
		},
	},
	Action:          handleAgentGatewaysAPIsUpdate,
	HideHelpCommand: true,
}

var agentGatewaysAPIsList = cli.Command{
	Name:    "list",
	Usage:   "Lists the REST APIs the instance serves as tools, without their credentials. The\ninstance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
	},
	Action:          handleAgentGatewaysAPIsList,
	HideHelpCommand: true,
}

var agentGatewaysAPIsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Removes an API and returns it as it was. An MCP server's `id` answers `404`. The\ninstance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "api-id",
			Required:  true,
			PathParam: "api_id",
		},
	},
	Action:          handleAgentGatewaysAPIsDelete,
	HideHelpCommand: true,
}

func handleAgentGatewaysAPIsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayAPINewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.APIs.New(
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
		Title:          "agent-gateways:apis create",
		Transform:      transform,
	})
}

func handleAgentGatewaysAPIsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("api-id") && len(unusedArgs) > 0 {
		cmd.Set("api-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayAPIUpdateParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.APIs.Update(
		ctx,
		cmd.Value("api-id").(string),
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
		Title:          "agent-gateways:apis update",
		Transform:      transform,
	})
}

func handleAgentGatewaysAPIsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AgentGateways.APIs.List(ctx, cmd.Value("instance-id").(string), options...)
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
		Title:          "agent-gateways:apis list",
		Transform:      transform,
	})
}

func handleAgentGatewaysAPIsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("api-id") && len(unusedArgs) > 0 {
		cmd.Set("api-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayAPIDeleteParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.APIs.Delete(
		ctx,
		cmd.Value("api-id").(string),
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
		Title:          "agent-gateways:apis delete",
		Transform:      transform,
	})
}

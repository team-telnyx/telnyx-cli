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

var agentGatewaysCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates an Agent Gateway instance on one of the account's WireGuard Interfaces.\nThe instance starts `pending`, moves to `provisioning`, and then to `ready` or\n`failed`. It joins the WireGuard Interface as a peer named `agw-<instance id>`,\nand its `gateway_url` is set when it is `ready`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wireguard-interface-id",
			Usage:    "One of the account's WireGuard Interfaces. The instance joins it as a peer, and the interface's other peers reach the instance.",
			Required: true,
			BodyPath: "wireguard_interface_id",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "A label for people.",
			BodyPath: "name",
		},
	},
	Action:          handleAgentGatewaysCreate,
	HideHelpCommand: true,
}

var agentGatewaysRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns one of the account's Agent Gateway instances.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
	},
	Action:          handleAgentGatewaysRetrieve,
	HideHelpCommand: true,
}

var agentGatewaysList = cli.Command{
	Name:    "list",
	Usage:   "Lists the account's Agent Gateway instances, newest first, one page at a time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page to return, counting from 1.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Results per page.",
			Default:   50,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAgentGatewaysList,
	HideHelpCommand: true,
}

var agentGatewaysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an Agent Gateway instance in any status. The response shows it\n`deleting`; repeating the request answers the same until the instance is gone,\nand then `404`. The instance's engine and everything stored on it are removed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
	},
	Action:          handleAgentGatewaysDelete,
	HideHelpCommand: true,
}

var agentGatewaysRetrieveModels = cli.Command{
	Name:    "retrieve-models",
	Usage:   "Lists the models the instance serves. Each has the name a request uses, the\nTelnyx Inference model behind it, and its price per million tokens.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
	},
	Action:          handleAgentGatewaysRetrieveModels,
	HideHelpCommand: true,
}

var agentGatewaysRetrieveSpendLogs = cli.Command{
	Name:    "retrieve-spend-logs",
	Usage:   "Lists the requests the instance served with its keys, newest first, one page at\na time. The window defaults to the last 24 hours and spans at most 31 days. A\nrow appears a few seconds after its request. The instance must be `ready`;\notherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[any]{
			Name:      "end",
			Usage:     "The window's end, an RFC 3339 timestamp with an offset. Defaults to now.",
			QueryPath: "end",
		},
		&requestflag.Flag[string]{
			Name:      "key-id",
			Usage:     "Only the requests made with this key.",
			QueryPath: "key_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page to return, counting from 1.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Results per page.",
			Default:   50,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[any]{
			Name:      "start",
			Usage:     "The window's start, an RFC 3339 timestamp with an offset. Defaults to 24 hours before `end`.",
			QueryPath: "start",
		},
		&requestflag.Flag[string]{
			Name:      "team-id",
			Usage:     "Only the requests made with this team's keys.",
			QueryPath: "team_id",
		},
	},
	Action:          handleAgentGatewaysRetrieveSpendLogs,
	HideHelpCommand: true,
}

func handleAgentGatewaysCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.AgentGatewayNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.New(ctx, params, options...)
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
		Title:          "agent-gateways create",
		Transform:      transform,
	})
}

func handleAgentGatewaysRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AgentGateways.Get(ctx, cmd.Value("instance-id").(string), options...)
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
		Title:          "agent-gateways retrieve",
		Transform:      transform,
	})
}

func handleAgentGatewaysList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AgentGateways.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "agent-gateways list",
			Transform:      transform,
		})
	} else {
		iter := client.AgentGateways.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "agent-gateways list",
			Transform:      transform,
		})
	}
}

func handleAgentGatewaysDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AgentGateways.Delete(ctx, cmd.Value("instance-id").(string), options...)
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
		Title:          "agent-gateways delete",
		Transform:      transform,
	})
}

func handleAgentGatewaysRetrieveModels(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AgentGateways.GetModels(ctx, cmd.Value("instance-id").(string), options...)
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
		Title:          "agent-gateways retrieve-models",
		Transform:      transform,
	})
}

func handleAgentGatewaysRetrieveSpendLogs(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayGetSpendLogsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.GetSpendLogs(
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
		Title:          "agent-gateways retrieve-spend-logs",
		Transform:      transform,
	})
}

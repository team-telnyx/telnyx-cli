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

var agentGatewaysKeysCreate = cli.Command{
	Name:    "create",
	Usage:   "Issues a key on the instance. The key's value is in this response only. Every\nfield is optional; a key without `models` may use every model the instance\nserves. The instance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[*string]{
			Name:     "budget-resets-every",
			Usage:    "How often the spend resets: a number and a unit, `s`, `m`, `h`, or `d`.",
			BodyPath: "budget_resets_every",
		},
		&requestflag.Flag[*string]{
			Name:     "expires-in",
			Usage:    "How long until the key stops working: a number and a unit, `s`, `m`, `h`, or `d`.",
			BodyPath: "expires_in",
		},
		&requestflag.Flag[*float64]{
			Name:     "max-budget",
			Usage:    "The most the key may spend, in USD.",
			BodyPath: "max_budget",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-parallel-requests",
			Usage:    "The most requests the key may have in flight at once.",
			BodyPath: "max_parallel_requests",
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "The models the key may use. Leave it out to allow every model.",
			BodyPath: "models",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "A label for people, unique on the instance.",
			BodyPath: "name",
		},
		&requestflag.Flag[*int64]{
			Name:     "requests-per-minute",
			Usage:    "The most requests the key may make per minute.",
			BodyPath: "requests_per_minute",
		},
		&requestflag.Flag[*string]{
			Name:     "team-id",
			Usage:    "A team on the instance. The key shares the team's model list, budget, and rate limits.",
			BodyPath: "team_id",
		},
		&requestflag.Flag[*int64]{
			Name:     "tokens-per-minute",
			Usage:    "The most tokens the key may use per minute.",
			BodyPath: "tokens_per_minute",
		},
	},
	Action:          handleAgentGatewaysKeysCreate,
	HideHelpCommand: true,
}

var agentGatewaysKeysUpdate = cli.Command{
	Name:    "update",
	Usage:   "Changes a key's settings without reissuing it. A field left out stays as it is,\nand `null` clears it. `models: []` lifts the model restriction. `expires_in`\ncounts from now. Setting `blocked` to `true` makes the instance refuse the key\nwith `401` while keeping its settings and spend; `false` lets it through again.\nThe instance must be `ready`; otherwise the request answers `409`.",
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
		&requestflag.Flag[bool]{
			Name:     "blocked",
			Usage:    "`true` makes the instance refuse the key; `false` lets it through again.",
			BodyPath: "blocked",
		},
		&requestflag.Flag[*string]{
			Name:     "budget-resets-every",
			Usage:    "How often the spend resets.",
			BodyPath: "budget_resets_every",
		},
		&requestflag.Flag[*string]{
			Name:     "expires-in",
			Usage:    "How long from now until the key stops working.",
			BodyPath: "expires_in",
		},
		&requestflag.Flag[*float64]{
			Name:     "max-budget",
			Usage:    "The most the key may spend, in USD.",
			BodyPath: "max_budget",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-parallel-requests",
			Usage:    "The most requests the key may have in flight at once.",
			BodyPath: "max_parallel_requests",
		},
		&requestflag.Flag[[]string]{
			Name:     "model",
			Usage:    "The models the key may use. `[]` allows every model.",
			BodyPath: "models",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A label for people.",
			BodyPath: "name",
		},
		&requestflag.Flag[*int64]{
			Name:     "requests-per-minute",
			Usage:    "The most requests the key may make per minute.",
			BodyPath: "requests_per_minute",
		},
		&requestflag.Flag[*string]{
			Name:     "team-id",
			Usage:    "Move the key to this team, or `null` to leave its team.",
			BodyPath: "team_id",
		},
		&requestflag.Flag[*int64]{
			Name:     "tokens-per-minute",
			Usage:    "The most tokens the key may use per minute.",
			BodyPath: "tokens_per_minute",
		},
	},
	Action:          handleAgentGatewaysKeysUpdate,
	HideHelpCommand: true,
}

var agentGatewaysKeysList = cli.Command{
	Name:    "list",
	Usage:   "Lists the instance's keys, without their values. Returns the first 100. The\ninstance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "team-id",
			Usage:     "Only the keys of this team.",
			QueryPath: "team_id",
		},
	},
	Action:          handleAgentGatewaysKeysList,
	HideHelpCommand: true,
}

var agentGatewaysKeysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Revokes a key and returns it as it was, without its value. The instance must be\n`ready`; otherwise the request answers `409`.",
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
	Action:          handleAgentGatewaysKeysDelete,
	HideHelpCommand: true,
}

func handleAgentGatewaysKeysCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayKeyNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Keys.New(
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
		Title:          "agent-gateways:keys create",
		Transform:      transform,
	})
}

func handleAgentGatewaysKeysUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayKeyUpdateParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Keys.Update(
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
		Title:          "agent-gateways:keys update",
		Transform:      transform,
	})
}

func handleAgentGatewaysKeysList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayKeyListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Keys.List(
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
		Title:          "agent-gateways:keys list",
		Transform:      transform,
	})
}

func handleAgentGatewaysKeysDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayKeyDeleteParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Keys.Delete(
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
		Title:          "agent-gateways:keys delete",
		Transform:      transform,
	})
}

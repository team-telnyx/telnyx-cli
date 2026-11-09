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

var agentGatewaysTeamsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a team on the instance. A team is a model list, a budget, and rate\nlimits that its keys share. Names need not be unique; the `id` identifies the\nteam. The instance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A label for people.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "budget-resets-every",
			Usage:    "How often the spend resets: a number and a unit, `s`, `m`, `h`, or `d`.",
			BodyPath: "budget_resets_every",
		},
		&requestflag.Flag[*float64]{
			Name:     "max-budget",
			Usage:    "The most the team's keys may spend together, in USD.",
			BodyPath: "max_budget",
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "The models the team's keys may use. Leave it out to allow every model.",
			BodyPath: "models",
		},
		&requestflag.Flag[*int64]{
			Name:     "requests-per-minute",
			Usage:    "The most requests the team's keys may make per minute, together.",
			BodyPath: "requests_per_minute",
		},
		&requestflag.Flag[*int64]{
			Name:     "tokens-per-minute",
			Usage:    "The most tokens the team's keys may use per minute, together.",
			BodyPath: "tokens_per_minute",
		},
	},
	Action:          handleAgentGatewaysTeamsCreate,
	HideHelpCommand: true,
}

var agentGatewaysTeamsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Changes a team's settings. A field left out stays as it is, and `null` clears\nit. `models: []` lifts the model restriction. Setting `blocked` to `true` makes\nthe instance refuse every key in the team with `401`. The instance must be\n`ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "team_id",
		},
		&requestflag.Flag[bool]{
			Name:     "blocked",
			Usage:    "`true` makes the instance refuse every key in the team; `false` lets them through again.",
			BodyPath: "blocked",
		},
		&requestflag.Flag[*string]{
			Name:     "budget-resets-every",
			Usage:    "How often the spend resets.",
			BodyPath: "budget_resets_every",
		},
		&requestflag.Flag[*float64]{
			Name:     "max-budget",
			Usage:    "The most the team's keys may spend together, in USD.",
			BodyPath: "max_budget",
		},
		&requestflag.Flag[[]string]{
			Name:     "model",
			Usage:    "The models the team's keys may use. `[]` allows every model.",
			BodyPath: "models",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A label for people.",
			BodyPath: "name",
		},
		&requestflag.Flag[*int64]{
			Name:     "requests-per-minute",
			Usage:    "The most requests the team's keys may make per minute, together.",
			BodyPath: "requests_per_minute",
		},
		&requestflag.Flag[*int64]{
			Name:     "tokens-per-minute",
			Usage:    "The most tokens the team's keys may use per minute, together.",
			BodyPath: "tokens_per_minute",
		},
	},
	Action:          handleAgentGatewaysTeamsUpdate,
	HideHelpCommand: true,
}

var agentGatewaysTeamsList = cli.Command{
	Name:    "list",
	Usage:   "Lists the instance's teams. The instance must be `ready`; otherwise the request\nanswers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
	},
	Action:          handleAgentGatewaysTeamsList,
	HideHelpCommand: true,
}

var agentGatewaysTeamsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a team that has no keys and returns it. A team that still has keys\nanswers `409`; delete or move its keys first. The instance must be `ready`;\notherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "team_id",
		},
	},
	Action:          handleAgentGatewaysTeamsDelete,
	HideHelpCommand: true,
}

func handleAgentGatewaysTeamsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayTeamNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Teams.New(
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
		Title:          "agent-gateways:teams create",
		Transform:      transform,
	})
}

func handleAgentGatewaysTeamsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("team-id") && len(unusedArgs) > 0 {
		cmd.Set("team-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayTeamUpdateParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Teams.Update(
		ctx,
		cmd.Value("team-id").(string),
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
		Title:          "agent-gateways:teams update",
		Transform:      transform,
	})
}

func handleAgentGatewaysTeamsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AgentGateways.Teams.List(ctx, cmd.Value("instance-id").(string), options...)
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
		Title:          "agent-gateways:teams list",
		Transform:      transform,
	})
}

func handleAgentGatewaysTeamsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("team-id") && len(unusedArgs) > 0 {
		cmd.Set("team-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayTeamDeleteParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Teams.Delete(
		ctx,
		cmd.Value("team-id").(string),
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
		Title:          "agent-gateways:teams delete",
		Transform:      transform,
	})
}

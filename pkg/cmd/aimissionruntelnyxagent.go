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

var aiMissionsRunsTelnyxAgentsList = cli.Command{
	Name:    "list",
	Usage:   "List all Telnyx agents linked to a run",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "run-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsRunsTelnyxAgentsList,
	HideHelpCommand: true,
}

var aiMissionsRunsTelnyxAgentsLink = cli.Command{
	Name:    "link",
	Usage:   "Link a Telnyx AI agent (voice/messaging) to a run",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "run-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "telnyx-agent-id",
			Usage:    "The Telnyx AI agent ID to link",
			Required: true,
			BodyPath: "telnyx_agent_id",
		},
	},
	Action:          handleAIMissionsRunsTelnyxAgentsLink,
	HideHelpCommand: true,
}

var aiMissionsRunsTelnyxAgentsUnlink = cli.Command{
	Name:    "unlink",
	Usage:   "Unlink a Telnyx agent from a run",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "run-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "telnyx-agent-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsRunsTelnyxAgentsUnlink,
	HideHelpCommand: true,
}

func handleAIMissionsRunsTelnyxAgentsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunTelnyxAgentListParams{
		MissionID: cmd.Value("mission-id").(string),
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
	_, err = client.AI.Missions.Runs.TelnyxAgents.List(
		ctx,
		cmd.Value("run-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs:telnyx-agents list", obj, format, transform)
}

func handleAIMissionsRunsTelnyxAgentsLink(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunTelnyxAgentLinkParams{
		MissionID: cmd.Value("mission-id").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Missions.Runs.TelnyxAgents.Link(
		ctx,
		cmd.Value("run-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs:telnyx-agents link", obj, format, transform)
}

func handleAIMissionsRunsTelnyxAgentsUnlink(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("telnyx-agent-id") && len(unusedArgs) > 0 {
		cmd.Set("telnyx-agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunTelnyxAgentUnlinkParams{
		MissionID: cmd.Value("mission-id").(string),
		RunID:     cmd.Value("run-id").(string),
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

	return client.AI.Missions.Runs.TelnyxAgents.Unlink(
		ctx,
		cmd.Value("telnyx-agent-id").(string),
		params,
		options...,
	)
}

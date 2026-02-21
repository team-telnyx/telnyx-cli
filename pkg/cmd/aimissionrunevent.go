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

var aiMissionsRunsEventsList = cli.Command{
	Name:    "list",
	Usage:   "List events for a run (paginated)",
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
			Name:      "agent-id",
			QueryPath: "agent_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (1-based)",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of items per page",
			Default:   50,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "step-id",
			QueryPath: "step_id",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			QueryPath: "type",
		},
	},
	Action:          handleAIMissionsRunsEventsList,
	HideHelpCommand: true,
}

var aiMissionsRunsEventsGetEventDetails = cli.Command{
	Name:    "get-event-details",
	Usage:   "Get details of a specific event",
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
			Name:     "event-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsRunsEventsGetEventDetails,
	HideHelpCommand: true,
}

var aiMissionsRunsEventsLog = cli.Command{
	Name:    "log",
	Usage:   "Log an event for a run",
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
			Name:     "summary",
			Required: true,
			BodyPath: "summary",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "agent-id",
			BodyPath: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Prevents duplicate events on retry",
			BodyPath: "idempotency_key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "payload",
			BodyPath: "payload",
		},
		&requestflag.Flag[string]{
			Name:     "step-id",
			BodyPath: "step_id",
		},
	},
	Action:          handleAIMissionsRunsEventsLog,
	HideHelpCommand: true,
}

func handleAIMissionsRunsEventsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunEventListParams{
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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Missions.Runs.Events.List(
			ctx,
			cmd.Value("run-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:missions:runs:events list", obj, format, transform)
	} else {
		iter := client.AI.Missions.Runs.Events.ListAutoPaging(
			ctx,
			cmd.Value("run-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "ai:missions:runs:events list", iter, format, transform)
	}
}

func handleAIMissionsRunsEventsGetEventDetails(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunEventGetEventDetailsParams{
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Missions.Runs.Events.GetEventDetails(
		ctx,
		cmd.Value("event-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs:events get-event-details", obj, format, transform)
}

func handleAIMissionsRunsEventsLog(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunEventLogParams{
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
	_, err = client.AI.Missions.Runs.Events.Log(
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
	return ShowJSON(os.Stdout, "ai:missions:runs:events log", obj, format, transform)
}

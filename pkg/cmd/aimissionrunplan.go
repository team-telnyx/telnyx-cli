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

var aiMissionsRunsPlanCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates the initial plan for the specified run from the provided steps and\nreturns the created plan steps. Progress is subsequently reported by updating\nindividual steps.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "mission-id",
			Required:  true,
			PathParam: "mission_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			Required: true,
			BodyPath: "steps",
		},
	},
	Action:          handleAIMissionsRunsPlanCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"step": {
		&requestflag.InnerFlag[string]{
			Name:       "step.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "step.sequence",
			InnerField: "sequence",
		},
		&requestflag.InnerFlag[string]{
			Name:       "step.step-id",
			InnerField: "step_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "step.metadata",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "step.parent-step-id",
			InnerField: "parent_step_id",
		},
	},
})

var aiMissionsRunsPlanRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the plan for the specified run, including all plan steps and their\nstatuses, so you can see how the mission was decomposed and how far execution\nhas progressed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "mission-id",
			Required:  true,
			PathParam: "mission_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
	},
	Action:          handleAIMissionsRunsPlanRetrieve,
	HideHelpCommand: true,
}

var aiMissionsRunsPlanAddStepsToPlan = requestflag.WithInnerFlags(cli.Command{
	Name:    "add-steps-to-plan",
	Usage:   "Add one or more steps to an existing plan",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "mission-id",
			Required:  true,
			PathParam: "mission_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			Required: true,
			BodyPath: "steps",
		},
	},
	Action:          handleAIMissionsRunsPlanAddStepsToPlan,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"step": {
		&requestflag.InnerFlag[string]{
			Name:       "step.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "step.sequence",
			InnerField: "sequence",
		},
		&requestflag.InnerFlag[string]{
			Name:       "step.step-id",
			InnerField: "step_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "step.metadata",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "step.parent-step-id",
			InnerField: "parent_step_id",
		},
	},
})

var aiMissionsRunsPlanGetStepDetails = cli.Command{
	Name:    "get-step-details",
	Usage:   "Returns the details of a single plan step within a run's plan, including its\nstatus.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "mission-id",
			Required:  true,
			PathParam: "mission_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
		&requestflag.Flag[string]{
			Name:      "step-id",
			Required:  true,
			PathParam: "step_id",
		},
	},
	Action:          handleAIMissionsRunsPlanGetStepDetails,
	HideHelpCommand: true,
}

var aiMissionsRunsPlanUpdateStep = cli.Command{
	Name:    "update-step",
	Usage:   "Updates the status of a single plan step and returns the updated step. Typically\ncalled by the executing agent as it works through the plan.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "mission-id",
			Required:  true,
			PathParam: "mission_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
		&requestflag.Flag[string]{
			Name:      "step-id",
			Required:  true,
			PathParam: "step_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "pending", "in_progress", "completed", "skipped", "failed".`,
			BodyPath: "status",
		},
	},
	Action:          handleAIMissionsRunsPlanUpdateStep,
	HideHelpCommand: true,
}

func handleAIMissionsRunsPlanCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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

	params := telnyx.AIMissionRunPlanNewParams{
		MissionID: cmd.Value("mission-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Missions.Runs.Plan.New(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "ai:missions:runs:plan create",
		Transform:      transform,
	})
}

func handleAIMissionsRunsPlanRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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

	params := telnyx.AIMissionRunPlanGetParams{
		MissionID: cmd.Value("mission-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Missions.Runs.Plan.Get(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "ai:missions:runs:plan retrieve",
		Transform:      transform,
	})
}

func handleAIMissionsRunsPlanAddStepsToPlan(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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

	params := telnyx.AIMissionRunPlanAddStepsToPlanParams{
		MissionID: cmd.Value("mission-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Missions.Runs.Plan.AddStepsToPlan(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "ai:missions:runs:plan add-steps-to-plan",
		Transform:      transform,
	})
}

func handleAIMissionsRunsPlanGetStepDetails(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("step-id") && len(unusedArgs) > 0 {
		cmd.Set("step-id", unusedArgs[0])
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

	params := telnyx.AIMissionRunPlanGetStepDetailsParams{
		MissionID: cmd.Value("mission-id").(string),
		RunID:     cmd.Value("run-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Missions.Runs.Plan.GetStepDetails(
		ctx,
		cmd.Value("step-id").(string),
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
		Title:          "ai:missions:runs:plan get-step-details",
		Transform:      transform,
	})
}

func handleAIMissionsRunsPlanUpdateStep(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("step-id") && len(unusedArgs) > 0 {
		cmd.Set("step-id", unusedArgs[0])
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

	params := telnyx.AIMissionRunPlanUpdateStepParams{
		MissionID: cmd.Value("mission-id").(string),
		RunID:     cmd.Value("run-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Missions.Runs.Plan.UpdateStep(
		ctx,
		cmd.Value("step-id").(string),
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
		Title:          "ai:missions:runs:plan update-step",
		Transform:      transform,
	})
}

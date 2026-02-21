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

var aiMissionsRunsPlanCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create the initial plan for a run",
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
	Usage:   "Get the plan (all steps) for a run",
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
	Action:          handleAIMissionsRunsPlanRetrieve,
	HideHelpCommand: true,
}

var aiMissionsRunsPlanAddStepsToPlan = requestflag.WithInnerFlags(cli.Command{
	Name:    "add-steps-to-plan",
	Usage:   "Add one or more steps to an existing plan",
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
	Usage:   "Get details of a specific plan step",
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
			Name:     "step-id",
			Required: true,
		},
	},
	Action:          handleAIMissionsRunsPlanGetStepDetails,
	HideHelpCommand: true,
}

var aiMissionsRunsPlanUpdateStep = cli.Command{
	Name:    "update-step",
	Usage:   "Update the status of a plan step",
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
			Name:     "step-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "status",
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

	params := telnyx.AIMissionRunPlanNewParams{
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs:plan create", obj, format, transform)
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

	params := telnyx.AIMissionRunPlanGetParams{
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs:plan retrieve", obj, format, transform)
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

	params := telnyx.AIMissionRunPlanAddStepsToPlanParams{
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs:plan add-steps-to-plan", obj, format, transform)
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

	params := telnyx.AIMissionRunPlanGetStepDetailsParams{
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs:plan get-step-details", obj, format, transform)
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

	params := telnyx.AIMissionRunPlanUpdateStepParams{
		MissionID: cmd.Value("mission-id").(string),
		RunID:     cmd.Value("run-id").(string),
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs:plan update-step", obj, format, transform)
}

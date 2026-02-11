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

var aiMissionsRunsCreate = cli.Command{
	Name:    "create",
	Usage:   "Start a new run for a mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "input",
			BodyPath: "input",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
	},
	Action:          handleAIMissionsRunsCreate,
	HideHelpCommand: true,
}

var aiMissionsRunsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get details of a specific run",
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
	Action:          handleAIMissionsRunsRetrieve,
	HideHelpCommand: true,
}

var aiMissionsRunsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update run status and/or result",
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
			Name:     "error",
			BodyPath: "error",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "result-payload",
			BodyPath: "result_payload",
		},
		&requestflag.Flag[string]{
			Name:     "result-summary",
			BodyPath: "result_summary",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			BodyPath: "status",
		},
	},
	Action:          handleAIMissionsRunsUpdate,
	HideHelpCommand: true,
}

var aiMissionsRunsList = cli.Command{
	Name:    "list",
	Usage:   "List all runs for a specific mission",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mission-id",
			Required: true,
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
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			QueryPath: "status",
		},
	},
	Action:          handleAIMissionsRunsList,
	HideHelpCommand: true,
}

var aiMissionsRunsCancelRun = cli.Command{
	Name:    "cancel-run",
	Usage:   "Cancel a running or paused run",
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
	Action:          handleAIMissionsRunsCancelRun,
	HideHelpCommand: true,
}

var aiMissionsRunsListRuns = cli.Command{
	Name:    "list-runs",
	Usage:   "List recent runs across all missions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (1-based)",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of items per page",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			QueryPath: "status",
		},
	},
	Action:          handleAIMissionsRunsListRuns,
	HideHelpCommand: true,
}

var aiMissionsRunsPauseRun = cli.Command{
	Name:    "pause-run",
	Usage:   "Pause a running run",
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
	Action:          handleAIMissionsRunsPauseRun,
	HideHelpCommand: true,
}

var aiMissionsRunsResumeRun = cli.Command{
	Name:    "resume-run",
	Usage:   "Resume a paused run",
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
	Action:          handleAIMissionsRunsResumeRun,
	HideHelpCommand: true,
}

func handleAIMissionsRunsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mission-id") && len(unusedArgs) > 0 {
		cmd.Set("mission-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunNewParams{}

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
	_, err = client.AI.Missions.Runs.New(
		ctx,
		cmd.Value("mission-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:missions:runs create", obj, format, transform)
}

func handleAIMissionsRunsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunGetParams{
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
	_, err = client.AI.Missions.Runs.Get(
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
	return ShowJSON(os.Stdout, "ai:missions:runs retrieve", obj, format, transform)
}

func handleAIMissionsRunsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunUpdateParams{
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
	_, err = client.AI.Missions.Runs.Update(
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
	return ShowJSON(os.Stdout, "ai:missions:runs update", obj, format, transform)
}

func handleAIMissionsRunsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mission-id") && len(unusedArgs) > 0 {
		cmd.Set("mission-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunListParams{}

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
		_, err = client.AI.Missions.Runs.List(
			ctx,
			cmd.Value("mission-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:missions:runs list", obj, format, transform)
	} else {
		iter := client.AI.Missions.Runs.ListAutoPaging(
			ctx,
			cmd.Value("mission-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "ai:missions:runs list", iter, format, transform)
	}
}

func handleAIMissionsRunsCancelRun(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunCancelRunParams{
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
	_, err = client.AI.Missions.Runs.CancelRun(
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
	return ShowJSON(os.Stdout, "ai:missions:runs cancel-run", obj, format, transform)
}

func handleAIMissionsRunsListRuns(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunListRunsParams{}

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
		_, err = client.AI.Missions.Runs.ListRuns(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:missions:runs list-runs", obj, format, transform)
	} else {
		iter := client.AI.Missions.Runs.ListRunsAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "ai:missions:runs list-runs", iter, format, transform)
	}
}

func handleAIMissionsRunsPauseRun(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunPauseRunParams{
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
	_, err = client.AI.Missions.Runs.PauseRun(
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
	return ShowJSON(os.Stdout, "ai:missions:runs pause-run", obj, format, transform)
}

func handleAIMissionsRunsResumeRun(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIMissionRunResumeRunParams{
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
	_, err = client.AI.Missions.Runs.ResumeRun(
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
	return ShowJSON(os.Stdout, "ai:missions:runs resume-run", obj, format, transform)
}

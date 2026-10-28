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

var aiTestsSchedulesV1CronJobsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the specified scheduled test job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAITestsSchedulesV1CronJobsRetrieve,
	HideHelpCommand: true,
}

var aiTestsSchedulesV1CronJobsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the specified scheduled test job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*int64]{
			Name:     "agent",
			BodyPath: "agent",
		},
		&requestflag.Flag[string]{
			Name:     "crontab-expression",
			Usage:    "Cron expression controlling when the job runs.",
			BodyPath: "crontab_expression",
		},
		&requestflag.Flag[string]{
			Name:     "execution-mode",
			Usage:    "Execution mode for the test runs, such as voice, text, or sip.",
			BodyPath: "execution_mode",
		},
		&requestflag.Flag[int64]{
			Name:     "frequency",
			BodyPath: "frequency",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			BodyPath: "is_active",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "notify-on",
			Usage:    `Allowed values: "never", "success", "failure", "both".`,
			BodyPath: "notify_on",
		},
		&requestflag.Flag[[]int64]{
			Name:     "personality",
			BodyPath: "personalities",
		},
		&requestflag.Flag[bool]{
			Name:     "run-as-text",
			BodyPath: "run_as_text",
		},
		&requestflag.Flag[[]int64]{
			Name:     "scenario",
			BodyPath: "scenarios",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "timezone",
			BodyPath: "timezone",
		},
	},
	Action:          handleAITestsSchedulesV1CronJobsUpdate,
	HideHelpCommand: true,
}

var aiTestsSchedulesV1CronJobsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the specified scheduled test job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAITestsSchedulesV1CronJobsDelete,
	HideHelpCommand: true,
}

var aiTestsSchedulesV1CronJobsCronJobs = cli.Command{
	Name:    "cron-jobs",
	Usage:   "Creates a scheduled test job that runs the selected scenarios on a cron\nexpression.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "crontab-expression",
			Usage:    "Cron expression controlling when the job runs.",
			Required: true,
			BodyPath: "crontab_expression",
		},
		&requestflag.Flag[int64]{
			Name:     "agent",
			BodyPath: "agent",
		},
		&requestflag.Flag[string]{
			Name:     "execution-mode",
			Usage:    "Execution mode for the test runs, such as voice, text, or sip.",
			BodyPath: "execution_mode",
		},
		&requestflag.Flag[int64]{
			Name:     "frequency",
			BodyPath: "frequency",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			BodyPath: "is_active",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "notify-on",
			Usage:    `Allowed values: "never", "success", "failure", "both".`,
			BodyPath: "notify_on",
		},
		&requestflag.Flag[[]int64]{
			Name:     "personality",
			BodyPath: "personalities",
		},
		&requestflag.Flag[bool]{
			Name:     "run-as-text",
			BodyPath: "run_as_text",
		},
		&requestflag.Flag[[]int64]{
			Name:     "scenario",
			BodyPath: "scenarios",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "timezone",
			BodyPath: "timezone",
		},
	},
	Action:          handleAITestsSchedulesV1CronJobsCronJobs,
	HideHelpCommand: true,
}

var aiTestsSchedulesV1CronJobsRetrieveCronJobs = cli.Command{
	Name:    "retrieve-cron-jobs",
	Usage:   "Returns the scheduled test jobs in your organization, optionally filtered by\ntest agent. A scheduled job runs selected scenarios on a cron expression.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "agent-id",
			Usage:     "Filter jobs belonging to the specified test agent.",
			QueryPath: "agent_id",
		},
	},
	Action:          handleAITestsSchedulesV1CronJobsRetrieveCronJobs,
	HideHelpCommand: true,
}

func handleAITestsSchedulesV1CronJobsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.AI.Tests.Schedules.V1.CronJobs.Get(ctx, cmd.Value("id").(int64), options...)
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
		Title:          "ai:tests:schedules:v1:cron-jobs retrieve",
		Transform:      transform,
	})
}

func handleAITestsSchedulesV1CronJobsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	params := telnyx.AITestScheduleV1CronJobUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.Schedules.V1.CronJobs.Update(
		ctx,
		cmd.Value("id").(int64),
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
		Title:          "ai:tests:schedules:v1:cron-jobs update",
		Transform:      transform,
	})
}

func handleAITestsSchedulesV1CronJobsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.AI.Tests.Schedules.V1.CronJobs.Delete(ctx, cmd.Value("id").(int64), options...)
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
		Title:          "ai:tests:schedules:v1:cron-jobs delete",
		Transform:      transform,
	})
}

func handleAITestsSchedulesV1CronJobsCronJobs(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestScheduleV1CronJobCronJobsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.Schedules.V1.CronJobs.CronJobs(ctx, params, options...)
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
		Title:          "ai:tests:schedules:v1:cron-jobs cron-jobs",
		Transform:      transform,
	})
}

func handleAITestsSchedulesV1CronJobsRetrieveCronJobs(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestScheduleV1CronJobGetCronJobsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.Schedules.V1.CronJobs.GetCronJobs(ctx, params, options...)
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
		Title:          "ai:tests:schedules:v1:cron-jobs retrieve-cron-jobs",
		Transform:      transform,
	})
}

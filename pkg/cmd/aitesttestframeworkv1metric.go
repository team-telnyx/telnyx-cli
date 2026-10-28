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

var aiTestsTestFrameworkV1MetricsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a metric in your organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*bool]{
			Name:     "add-to-new-agents",
			BodyPath: "add_to_new_agents",
		},
		&requestflag.Flag[[]int64]{
			Name:     "agent",
			BodyPath: "agents",
		},
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			BodyPath: "assistant_id",
		},
		&requestflag.Flag[bool]{
			Name:     "audio-enabled",
			BodyPath: "audio_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[int64]{
			Name:     "display-order",
			BodyPath: "display_order",
		},
		&requestflag.Flag[[]string]{
			Name:     "enum-value",
			BodyPath: "enum_values",
		},
		&requestflag.Flag[string]{
			Name:     "eval-type",
			Usage:    "What the metric scores: numeric, continuous_qualitative, or binary_workflow_adherence.",
			BodyPath: "eval_type",
		},
		&requestflag.Flag[bool]{
			Name:     "observability-enabled",
			BodyPath: "observability_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Judge prompt used by llm_judge metrics.",
			BodyPath: "prompt",
		},
		&requestflag.Flag[bool]{
			Name:     "sampling-enabled",
			BodyPath: "sampling_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "simulation-enabled",
			BodyPath: "simulation_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "How the metric is evaluated: custom_code or llm_judge.",
			BodyPath: "type",
		},
	},
	Action:          handleAITestsTestFrameworkV1MetricsCreate,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1MetricsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the specified metric.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAITestsTestFrameworkV1MetricsRetrieve,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1MetricsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the specified metric.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*bool]{
			Name:     "add-to-new-agents",
			BodyPath: "add_to_new_agents",
		},
		&requestflag.Flag[[]int64]{
			Name:     "agent",
			BodyPath: "agents",
		},
		&requestflag.Flag[bool]{
			Name:     "audio-enabled",
			BodyPath: "audio_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[int64]{
			Name:     "display-order",
			BodyPath: "display_order",
		},
		&requestflag.Flag[[]string]{
			Name:     "enum-value",
			BodyPath: "enum_values",
		},
		&requestflag.Flag[string]{
			Name:     "eval-type",
			Usage:    "What the metric scores: numeric, continuous_qualitative, or binary_workflow_adherence.",
			BodyPath: "eval_type",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "observability-enabled",
			BodyPath: "observability_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Judge prompt used by llm_judge metrics.",
			BodyPath: "prompt",
		},
		&requestflag.Flag[bool]{
			Name:     "sampling-enabled",
			BodyPath: "sampling_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "simulation-enabled",
			BodyPath: "simulation_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "How the metric is evaluated: custom_code or llm_judge.",
			BodyPath: "type",
		},
	},
	Action:          handleAITestsTestFrameworkV1MetricsUpdate,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1MetricsList = cli.Command{
	Name:    "list",
	Usage:   "Returns the metrics available in your organization, optionally filtered by test\nagent. Each metric scores one aspect of a conversation or scenario run.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "agent-id",
			Usage:     "Filter metrics attached to the specified test agent.",
			QueryPath: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:      "assistant-id",
			Usage:     "Filter metrics attached to the specified Telnyx assistant.",
			QueryPath: "assistant_id",
		},
	},
	Action:          handleAITestsTestFrameworkV1MetricsList,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1MetricsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the specified metric.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAITestsTestFrameworkV1MetricsDelete,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1MetricsBulkCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "bulk-create",
	Usage:   "Creates several metrics in one request. Accepts an array of metric objects and\nreturns the created metrics.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "body",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleAITestsTestFrameworkV1MetricsBulkCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"body": {
		&requestflag.InnerFlag[string]{
			Name:       "body.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "body.add-to-new-agents",
			InnerField: "add_to_new_agents",
		},
		&requestflag.InnerFlag[[]int64]{
			Name:       "body.agents",
			InnerField: "agents",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.assistant-id",
			InnerField: "assistant_id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "body.audio-enabled",
			InnerField: "audio_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "body.display-order",
			InnerField: "display_order",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "body.enum-values",
			InnerField: "enum_values",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.eval-type",
			Usage:      "What the metric scores: numeric, continuous_qualitative, or binary_workflow_adherence.",
			InnerField: "eval_type",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "body.observability-enabled",
			InnerField: "observability_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.prompt",
			Usage:      "Judge prompt used by llm_judge metrics.",
			InnerField: "prompt",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "body.sampling-enabled",
			InnerField: "sampling_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "body.simulation-enabled",
			InnerField: "simulation_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.type",
			Usage:      "How the metric is evaluated: custom_code or llm_judge.",
			InnerField: "type",
		},
	},
})

var aiTestsTestFrameworkV1MetricsGenerateMetrics = cli.Command{
	Name:    "generate-metrics",
	Usage:   "Starts background generation of suggested metrics for a test agent. Returns a\nprogress ID to poll with the get-test-metric-generation-progress endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "agent-id",
			Usage:    "Test agent to suggest metrics for.",
			Required: true,
			BodyPath: "agent_id",
		},
		&requestflag.Flag[int64]{
			Name:     "num-metrics",
			Usage:    "How many metrics to suggest.",
			Required: true,
			BodyPath: "num_metrics",
		},
		&requestflag.Flag[[]int64]{
			Name:     "scenario",
			Usage:    "Scenario IDs to base the suggestions on.",
			BodyPath: "scenarios",
		},
	},
	Action:          handleAITestsTestFrameworkV1MetricsGenerateMetrics,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1MetricsRetrieveGenerateMetricsProgress = cli.Command{
	Name:    "retrieve-generate-metrics-progress",
	Usage:   "Returns the status of a background metric-generation job started with the\ngenerate-test-metrics endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "progress-id",
			Usage:     "Progress ID returned by the generate-test-metrics endpoint.",
			Required:  true,
			QueryPath: "progress_id",
		},
	},
	Action:          handleAITestsTestFrameworkV1MetricsRetrieveGenerateMetricsProgress,
	HideHelpCommand: true,
}

func handleAITestsTestFrameworkV1MetricsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1MetricNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Metrics.New(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:metrics create",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1MetricsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Tests.TestFramework.V1.Metrics.Get(ctx, cmd.Value("id").(int64), options...)
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
		Title:          "ai:tests:test-framework:v1:metrics retrieve",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1MetricsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1MetricUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Metrics.Update(
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
		Title:          "ai:tests:test-framework:v1:metrics update",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1MetricsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1MetricListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Metrics.List(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:metrics list",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1MetricsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Tests.TestFramework.V1.Metrics.Delete(ctx, cmd.Value("id").(int64), options...)
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
		Title:          "ai:tests:test-framework:v1:metrics delete",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1MetricsBulkCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1MetricBulkNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Metrics.BulkNew(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:metrics bulk-create",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1MetricsGenerateMetrics(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1MetricGenerateMetricsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Metrics.GenerateMetrics(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:metrics generate-metrics",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1MetricsRetrieveGenerateMetricsProgress(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1MetricGetGenerateMetricsProgressParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Metrics.GetGenerateMetricsProgress(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:metrics retrieve-generate-metrics-progress",
		Transform:      transform,
	})
}

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

var aiTestsTestFrameworkV1ScenariosCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates an evaluator scenario in your organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "personality",
			Usage:    "ID of the personality the caller simulates.",
			Required: true,
			BodyPath: "personality",
		},
		&requestflag.Flag[*int64]{
			Name:     "agent",
			BodyPath: "agent",
		},
		&requestflag.Flag[string]{
			Name:     "expected-outcome-prompt",
			BodyPath: "expected_outcome_prompt",
		},
		&requestflag.Flag[string]{
			Name:     "first-message",
			BodyPath: "first_message",
		},
		&requestflag.Flag[*string]{
			Name:     "folder-path",
			BodyPath: "folder_path",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[[]int64]{
			Name:     "metric",
			BodyPath: "metrics",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[int64]{
			Name:     "test-profile",
			BodyPath: "test_profile",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosCreate,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the specified evaluator scenario with full detail.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosRetrieve,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the specified evaluator scenario.",
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
			Name:     "expected-outcome-prompt",
			BodyPath: "expected_outcome_prompt",
		},
		&requestflag.Flag[string]{
			Name:     "first-message",
			BodyPath: "first_message",
		},
		&requestflag.Flag[*string]{
			Name:     "folder-path",
			BodyPath: "folder_path",
		},
		&requestflag.Flag[int64]{
			Name:     "inbound-phone-number",
			BodyPath: "inbound_phone_number",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-duration",
			BodyPath: "max_duration",
		},
		&requestflag.Flag[[]int64]{
			Name:     "metric",
			BodyPath: "metrics",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "personality",
			BodyPath: "personality",
		},
		&requestflag.Flag[int64]{
			Name:     "phone-number",
			BodyPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:     "scenario-language",
			BodyPath: "scenario_language",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[int64]{
			Name:     "test-profile",
			BodyPath: "test_profile",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool-id",
			BodyPath: "tool_ids",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosUpdate,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of evaluator scenarios in your organization, optionally\nfiltered by test agent or tags.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "agent-id",
			Usage:     "Filter scenarios belonging to the specified test agent.",
			QueryPath: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:      "assistant-id",
			Usage:     "Filter scenarios belonging to the specified Telnyx assistant.",
			QueryPath: "assistant_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number (1-based).",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results per page.",
			Default:   25,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "tags",
			Usage:     "Filter scenarios by a comma-separated list of tags.",
			QueryPath: "tags",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosList,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the specified evaluator scenario.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosDelete,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosGenerateBg = cli.Command{
	Name:    "generate-bg",
	Usage:   "Starts background generation of evaluator scenarios for a test agent. Returns a\nprogress ID to poll with the generate-scenarios-progress endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "num-scenarios",
			Usage:    "How many scenarios to generate.",
			Required: true,
			BodyPath: "num_scenarios",
		},
		&requestflag.Flag[int64]{
			Name:     "agent-id",
			BodyPath: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:     "extra-instructions",
			Usage:    "Optional guidance steering scenario generation.",
			BodyPath: "extra_instructions",
		},
		&requestflag.Flag[string]{
			Name:     "first-message",
			BodyPath: "first_message",
		},
		&requestflag.Flag[*string]{
			Name:     "folder-path",
			BodyPath: "folder_path",
		},
		&requestflag.Flag[bool]{
			Name:     "generate-expected-outcomes",
			Usage:    "Whether to generate expected-outcome prompts for the scenarios.",
			BodyPath: "generate_expected_outcomes",
		},
		&requestflag.Flag[*int64]{
			Name:     "inbound-phone-number",
			BodyPath: "inbound_phone_number",
		},
		&requestflag.Flag[[]int64]{
			Name:     "personality",
			BodyPath: "personalities",
		},
		&requestflag.Flag[string]{
			Name:     "scenario-type",
			Usage:    "Type of scenario to generate.",
			BodyPath: "scenario_type",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[int64]{
			Name:     "test-profile",
			BodyPath: "test_profile",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool-id",
			BodyPath: "tool_ids",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosGenerateBg,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosRetrieveExportScenarios = cli.Command{
	Name:    "retrieve-export-scenarios",
	Usage:   "Exports the selected evaluator scenarios as a CSV or XLSX file download.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "export-type",
			Usage:     "File format of the export.",
			Required:  true,
			QueryPath: "export_type",
		},
		&requestflag.Flag[string]{
			Name:      "scenario-ids",
			Usage:     "Comma-separated scenario IDs to export.",
			Required:  true,
			QueryPath: "scenario_ids",
		},
		&requestflag.Flag[int64]{
			Name:      "agent-id",
			Usage:     "Export scenarios belonging to the specified test agent.",
			QueryPath: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosRetrieveExportScenarios,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosRetrieveGenerateProgress = cli.Command{
	Name:    "retrieve-generate-progress",
	Usage:   "Returns the status of a background scenario-generation job started with the\ngenerate-evaluator-scenarios endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "progress-id",
			Usage:     "Progress ID returned by the generate-evaluator-scenarios endpoint.",
			Required:  true,
			QueryPath: "progress_id",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosRetrieveGenerateProgress,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosRetrieveScenarioAgentProgress = cli.Command{
	Name:    "retrieve-scenario-agent-progress",
	Usage:   "Returns the status of an improve-scenarios session started with the\nimprove-evaluator-scenarios endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "progress-id",
			Usage:     "Progress ID returned by the improve-evaluator-scenarios endpoint.",
			Required:  true,
			QueryPath: "progress_id",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosRetrieveScenarioAgentProgress,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosRunScenarios = cli.Command{
	Name:    "run-scenarios",
	Usage:   "Runs the selected evaluator scenarios against their test agents and creates a\nrun batch. Returns the ID of the created batch, whose results can be listed with\nthe test-run-results endpoints.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "agent-id",
			BodyPath: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:     "agent-number",
			BodyPath: "agent_number",
		},
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			BodyPath: "assistant_id",
		},
		&requestflag.Flag[int64]{
			Name:     "concurrency-limit",
			Usage:    "Maximum number of concurrent simulated calls.",
			BodyPath: "concurrency_limit",
		},
		&requestflag.Flag[int64]{
			Name:     "frequency",
			BodyPath: "frequency",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    `Allowed values: "same_number", "different_numbers".`,
			BodyPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "outbound-phone-number",
			BodyPath: "outbound_phone_number",
		},
		&requestflag.Flag[[]int64]{
			Name:     "personality-id",
			BodyPath: "personality_ids",
		},
		&requestflag.Flag[any]{
			Name:     "scenarios",
			Usage:    "Scenario IDs to run, or the literal string 'all'.",
			BodyPath: "scenarios",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[[]int64]{
			Name:     "test-profile-id",
			BodyPath: "test_profile_ids",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosRunScenarios,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1ScenariosScenarioAgent = requestflag.WithInnerFlags(cli.Command{
	Name:    "scenario-agent",
	Usage:   "Starts an interactive session that improves selected evaluator scenarios from a\nnatural-language instruction. Returns a progress ID to poll with the\nimprove-scenarios-progress endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "scenarios",
			Usage:    "Scenario IDs to improve, or the literal string 'all'.",
			Required: true,
			BodyPath: "scenarios",
		},
		&requestflag.Flag[string]{
			Name:     "user-message",
			Usage:    "What the assistant should improve about the selected scenarios.",
			Required: true,
			BodyPath: "user_message",
		},
		&requestflag.Flag[*int64]{
			Name:     "agent-id",
			BodyPath: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:     "agent-resolution-response",
			BodyPath: "agent_resolution_response",
		},
		&requestflag.Flag[bool]{
			Name:     "apply-changes",
			BodyPath: "apply_changes",
		},
		&requestflag.Flag[any]{
			Name:     "context",
			BodyPath: "context",
		},
		&requestflag.Flag[[]int64]{
			Name:     "knowledge-base",
			BodyPath: "knowledge_base",
		},
		&requestflag.Flag[int64]{
			Name:     "num-scenarios",
			BodyPath: "num_scenarios",
		},
		&requestflag.Flag[any]{
			Name:     "pending-tool-call",
			BodyPath: "pending_tool_calls",
		},
		&requestflag.Flag[[]int64]{
			Name:     "scenarios-needing-agent",
			BodyPath: "scenarios_needing_agent",
		},
		&requestflag.Flag[*int64]{
			Name:     "session-id",
			BodyPath: "session_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "updated-scenario",
			BodyPath: "updated_scenarios",
		},
	},
	Action:          handleAITestsTestFrameworkV1ScenariosScenarioAgent,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"context": {
		&requestflag.InnerFlag[string]{
			Name:                  "context.content",
			InnerField:            "content",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "context.role",
			InnerField:            "role",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "context.timestamp",
			InnerField:            "timestamp",
			OuterIsArrayOfObjects: true,
		},
	},
})

func handleAITestsTestFrameworkV1ScenariosCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.New(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios create",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.Get(ctx, cmd.Value("id").(int64), options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios retrieve",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.Update(
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
		Title:          "ai:tests:test-framework:v1:scenarios update",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.List(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios list",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.Delete(ctx, cmd.Value("id").(int64), options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios delete",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosGenerateBg(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioGenerateBgParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.GenerateBg(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios generate-bg",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosRetrieveExportScenarios(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioGetExportScenariosParams{}

	response, err := client.AI.Tests.TestFramework.V1.Scenarios.GetExportScenarios(ctx, params, options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}

func handleAITestsTestFrameworkV1ScenariosRetrieveGenerateProgress(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioGetGenerateProgressParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.GetGenerateProgress(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios retrieve-generate-progress",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosRetrieveScenarioAgentProgress(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioGetScenarioAgentProgressParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.GetScenarioAgentProgress(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios retrieve-scenario-agent-progress",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosRunScenarios(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioRunScenariosParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.RunScenarios(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios run-scenarios",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1ScenariosScenarioAgent(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1ScenarioScenarioAgentParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Scenarios.ScenarioAgent(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:scenarios scenario-agent",
		Transform:      transform,
	})
}

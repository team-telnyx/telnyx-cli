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

var aiTestsTestFrameworkV1AiagentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a test agent in your organization. The agent wraps a Telnyx AI assistant\nso simulated callers can be sent against it.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-name",
			Usage:    "Name of the test agent.",
			Required: true,
			BodyPath: "agent_name",
		},
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Usage:    "ID of the Telnyx AI assistant under test.",
			Required: true,
			BodyPath: "assistant_id",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-fetch-calls-enabled",
			BodyPath: "auto_fetch_calls_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-sync-prompt-enabled",
			BodyPath: "auto_sync_prompt_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-update-dropoff-nodes",
			BodyPath: "auto_update_dropoff_nodes",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-update-topic-nodes",
			BodyPath: "auto_update_topic_nodes",
		},
		&requestflag.Flag[string]{
			Name:     "contact-number",
			BodyPath: "contact_number",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "ISO language code the test agent converses in.",
			BodyPath: "language",
		},
		&requestflag.Flag[int64]{
			Name:     "llm-max-tokens",
			BodyPath: "llm_max_tokens",
		},
		&requestflag.Flag[string]{
			Name:     "llm-model",
			BodyPath: "llm_model",
		},
		&requestflag.Flag[string]{
			Name:     "llm-system-prompt",
			BodyPath: "llm_system_prompt",
		},
		&requestflag.Flag[float64]{
			Name:     "llm-temperature",
			BodyPath: "llm_temperature",
		},
		&requestflag.Flag[bool]{
			Name:     "outbound-auto-call",
			BodyPath: "outbound_auto_call",
		},
		&requestflag.Flag[string]{
			Name:     "sip-endpoint",
			BodyPath: "sip_endpoint",
		},
	},
	Action:          handleAITestsTestFrameworkV1AiagentsCreate,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1AiagentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the specified test agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAITestsTestFrameworkV1AiagentsRetrieve,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1AiagentsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the specified test agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "agent-name",
			Usage:    "Name of the test agent.",
			BodyPath: "agent_name",
		},
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Usage:    "ID of the Telnyx AI assistant under test.",
			BodyPath: "assistant_id",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-fetch-calls-enabled",
			BodyPath: "auto_fetch_calls_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-sync-prompt-enabled",
			BodyPath: "auto_sync_prompt_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-update-dropoff-nodes",
			BodyPath: "auto_update_dropoff_nodes",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-update-topic-nodes",
			BodyPath: "auto_update_topic_nodes",
		},
		&requestflag.Flag[string]{
			Name:     "contact-number",
			BodyPath: "contact_number",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "ISO language code the test agent converses in.",
			BodyPath: "language",
		},
		&requestflag.Flag[int64]{
			Name:     "llm-max-tokens",
			BodyPath: "llm_max_tokens",
		},
		&requestflag.Flag[string]{
			Name:     "llm-model",
			BodyPath: "llm_model",
		},
		&requestflag.Flag[string]{
			Name:     "llm-system-prompt",
			BodyPath: "llm_system_prompt",
		},
		&requestflag.Flag[float64]{
			Name:     "llm-temperature",
			BodyPath: "llm_temperature",
		},
		&requestflag.Flag[bool]{
			Name:     "outbound-auto-call",
			BodyPath: "outbound_auto_call",
		},
		&requestflag.Flag[string]{
			Name:     "sip-endpoint",
			BodyPath: "sip_endpoint",
		},
	},
	Action:          handleAITestsTestFrameworkV1AiagentsUpdate,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1AiagentsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of test agents in your organization. A test agent wraps\nthe Telnyx AI assistant under test and drives simulated callers against it.",
	Suggest: true,
	Flags: []cli.Flag{
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
	},
	Action:          handleAITestsTestFrameworkV1AiagentsList,
	HideHelpCommand: true,
}

var aiTestsTestFrameworkV1AiagentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the specified test agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAITestsTestFrameworkV1AiagentsDelete,
	HideHelpCommand: true,
}

func handleAITestsTestFrameworkV1AiagentsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1AiagentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Aiagents.New(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:aiagents create",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1AiagentsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Tests.TestFramework.V1.Aiagents.Get(ctx, cmd.Value("id").(int64), options...)
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
		Title:          "ai:tests:test-framework:v1:aiagents retrieve",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1AiagentsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1AiagentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Aiagents.Update(
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
		Title:          "ai:tests:test-framework:v1:aiagents update",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1AiagentsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITestTestFrameworkV1AiagentListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Tests.TestFramework.V1.Aiagents.List(ctx, params, options...)
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
		Title:          "ai:tests:test-framework:v1:aiagents list",
		Transform:      transform,
	})
}

func handleAITestsTestFrameworkV1AiagentsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AI.Tests.TestFramework.V1.Aiagents.Delete(ctx, cmd.Value("id").(int64), options...)
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
		Title:          "ai:tests:test-framework:v1:aiagents delete",
		Transform:      transform,
	})
}

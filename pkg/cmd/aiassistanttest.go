// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiAssistantsTestsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a comprehensive test configuration for evaluating AI assistant\nperformance",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "destination",
			Usage:    "The target destination for the test conversation. Format depends on the channel: phone number for SMS/voice, webhook URL for web chat, etc.",
			Required: true,
			BodyPath: "destination",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "Detailed instructions that define the test scenario and what the assistant should accomplish. This guides the test execution and evaluation.",
			Required: true,
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A descriptive name for the assistant test. This will be used to identify the test in the UI and reports.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "rubric",
			Usage:    "Evaluation criteria used to assess the assistant's performance. Each rubric item contains a name and specific criteria for evaluation.",
			Required: true,
			BodyPath: "rubric",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional detailed description of what this test evaluates and its purpose. Helps team members understand the test's objectives.",
			BodyPath: "description",
		},
		&requestflag.Flag[int64]{
			Name:     "max-duration-seconds",
			Usage:    "Maximum duration in seconds that the test conversation should run before timing out. If not specified, uses system default timeout.",
			BodyPath: "max_duration_seconds",
		},
		&requestflag.Flag[string]{
			Name:     "telnyx-conversation-channel",
			BodyPath: "telnyx_conversation_channel",
		},
		&requestflag.Flag[string]{
			Name:     "test-suite",
			Usage:    "Optional test suite name to group related tests together. Useful for organizing tests by feature, team, or release cycle.",
			BodyPath: "test_suite",
		},
	},
	Action:          handleAIAssistantsTestsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"rubric": {
		&requestflag.InnerFlag[string]{
			Name:       "rubric.criteria",
			Usage:      "Specific guidance on how to assess the assistant’s performance for this rubric item.",
			InnerField: "criteria",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rubric.name",
			Usage:      "Label for the evaluation criterion, e.g., Empathy, Accuracy, Clarity.",
			InnerField: "name",
		},
	},
})

var aiAssistantsTestsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves detailed information about a specific assistant test",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "test-id",
			Required: true,
		},
	},
	Action:          handleAIAssistantsTestsRetrieve,
	HideHelpCommand: true,
}

var aiAssistantsTestsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing assistant test configuration with new settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "test-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Updated description of the test's purpose and evaluation criteria.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "destination",
			Usage:    "Updated target destination for test conversations.",
			BodyPath: "destination",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "Updated test scenario instructions and objectives.",
			BodyPath: "instructions",
		},
		&requestflag.Flag[int64]{
			Name:     "max-duration-seconds",
			Usage:    "Updated maximum test duration in seconds.",
			BodyPath: "max_duration_seconds",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Updated name for the assistant test. Must be unique and descriptive.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "rubric",
			Usage:    "Updated evaluation criteria for assessing assistant performance.",
			BodyPath: "rubric",
		},
		&requestflag.Flag[string]{
			Name:     "telnyx-conversation-channel",
			BodyPath: "telnyx_conversation_channel",
		},
		&requestflag.Flag[string]{
			Name:     "test-suite",
			Usage:    "Updated test suite assignment for better organization.",
			BodyPath: "test_suite",
		},
	},
	Action:          handleAIAssistantsTestsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"rubric": {
		&requestflag.InnerFlag[string]{
			Name:       "rubric.criteria",
			Usage:      "Specific guidance on how to assess the assistant’s performance for this rubric item.",
			InnerField: "criteria",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rubric.name",
			Usage:      "Label for the evaluation criterion, e.g., Empathy, Accuracy, Clarity.",
			InnerField: "name",
		},
	},
})

var aiAssistantsTestsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of assistant tests with optional filtering\ncapabilities",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "destination",
			Usage:     "Filter tests by destination (phone number, webhook URL, etc.)",
			QueryPath: "destination",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "telnyx-conversation-channel",
			Usage:     "Filter tests by communication channel (e.g., 'web_chat', 'sms')",
			QueryPath: "telnyx_conversation_channel",
		},
		&requestflag.Flag[string]{
			Name:      "test-suite",
			Usage:     "Filter tests by test suite name",
			QueryPath: "test_suite",
		},
	},
	Action:          handleAIAssistantsTestsList,
	HideHelpCommand: true,
}

var aiAssistantsTestsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently removes an assistant test and all associated data",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "test-id",
			Required: true,
		},
	},
	Action:          handleAIAssistantsTestsDelete,
	HideHelpCommand: true,
}

func handleAIAssistantsTestsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTestNewParams{}

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
	_, err = client.AI.Assistants.Tests.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:tests create", obj, format, transform)
}

func handleAIAssistantsTestsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("test-id") && len(unusedArgs) > 0 {
		cmd.Set("test-id", unusedArgs[0])
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
	_, err = client.AI.Assistants.Tests.Get(ctx, cmd.Value("test-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:tests retrieve", obj, format, transform)
}

func handleAIAssistantsTestsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("test-id") && len(unusedArgs) > 0 {
		cmd.Set("test-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTestUpdateParams{}

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
	_, err = client.AI.Assistants.Tests.Update(
		ctx,
		cmd.Value("test-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:tests update", obj, format, transform)
}

func handleAIAssistantsTestsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTestListParams{}

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
		_, err = client.AI.Assistants.Tests.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:assistants:tests list", obj, format, transform)
	} else {
		iter := client.AI.Assistants.Tests.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "ai:assistants:tests list", iter, format, transform)
	}
}

func handleAIAssistantsTestsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("test-id") && len(unusedArgs) > 0 {
		cmd.Set("test-id", unusedArgs[0])
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

	return client.AI.Assistants.Tests.Delete(ctx, cmd.Value("test-id").(string), options...)
}

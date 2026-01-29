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

var aiAssistantsTestsTestSuitesRunsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves paginated history of test runs for a specific test suite with\nfiltering options",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "suite-name",
			Required: true,
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
			Name:      "status",
			Usage:     "Filter runs by execution status (pending, running, completed, failed, timeout)",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "test-suite-run-id",
			Usage:     "Filter runs by specific suite execution batch ID",
			QueryPath: "test_suite_run_id",
		},
	},
	Action:          handleAIAssistantsTestsTestSuitesRunsList,
	HideHelpCommand: true,
}

var aiAssistantsTestsTestSuitesRunsTrigger = cli.Command{
	Name:    "trigger",
	Usage:   "Executes all tests within a specific test suite as a batch operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "suite-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "destination-version-id",
			Usage:    "Optional assistant version ID to use for all test runs in this suite. If provided, the version must exist or a 400 error will be returned. If not provided, test will run on main version",
			BodyPath: "destination_version_id",
		},
	},
	Action:          handleAIAssistantsTestsTestSuitesRunsTrigger,
	HideHelpCommand: true,
}

func handleAIAssistantsTestsTestSuitesRunsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("suite-name") && len(unusedArgs) > 0 {
		cmd.Set("suite-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTestTestSuiteRunListParams{}

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
		_, err = client.AI.Assistants.Tests.TestSuites.Runs.List(
			ctx,
			cmd.Value("suite-name").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:assistants:tests:test-suites:runs list", obj, format, transform)
	} else {
		iter := client.AI.Assistants.Tests.TestSuites.Runs.ListAutoPaging(
			ctx,
			cmd.Value("suite-name").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "ai:assistants:tests:test-suites:runs list", iter, format, transform)
	}
}

func handleAIAssistantsTestsTestSuitesRunsTrigger(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("suite-name") && len(unusedArgs) > 0 {
		cmd.Set("suite-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTestTestSuiteRunTriggerParams{}

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
	_, err = client.AI.Assistants.Tests.TestSuites.Runs.Trigger(
		ctx,
		cmd.Value("suite-name").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:tests:test-suites:runs trigger", obj, format, transform)
}

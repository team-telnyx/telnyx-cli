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

var aiAssistantsTestsRunsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves detailed information about a specific test run execution",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "test-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "run-id",
			Required: true,
		},
	},
	Action:          handleAIAssistantsTestsRunsRetrieve,
	HideHelpCommand: true,
}

var aiAssistantsTestsRunsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves paginated execution history for a specific assistant test with\nfiltering options",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "test-id",
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
	},
	Action:          handleAIAssistantsTestsRunsList,
	HideHelpCommand: true,
}

var aiAssistantsTestsRunsTrigger = cli.Command{
	Name:    "trigger",
	Usage:   "Initiates immediate execution of a specific assistant test",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "test-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "destination-version-id",
			Usage:    "Optional assistant version ID to use for this test run. If provided, the version must exist or a 400 error will be returned. If not provided, test will run on main version",
			BodyPath: "destination_version_id",
		},
	},
	Action:          handleAIAssistantsTestsRunsTrigger,
	HideHelpCommand: true,
}

func handleAIAssistantsTestsRunsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTestRunGetParams{
		TestID: cmd.Value("test-id").(string),
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
	_, err = client.AI.Assistants.Tests.Runs.Get(
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
	return ShowJSON(os.Stdout, "ai:assistants:tests:runs retrieve", obj, format, transform)
}

func handleAIAssistantsTestsRunsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("test-id") && len(unusedArgs) > 0 {
		cmd.Set("test-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTestRunListParams{}

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
		_, err = client.AI.Assistants.Tests.Runs.List(
			ctx,
			cmd.Value("test-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:assistants:tests:runs list", obj, format, transform)
	} else {
		iter := client.AI.Assistants.Tests.Runs.ListAutoPaging(
			ctx,
			cmd.Value("test-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "ai:assistants:tests:runs list", iter, format, transform)
	}
}

func handleAIAssistantsTestsRunsTrigger(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("test-id") && len(unusedArgs) > 0 {
		cmd.Set("test-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantTestRunTriggerParams{}

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
	_, err = client.AI.Assistants.Tests.Runs.Trigger(
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
	return ShowJSON(os.Stdout, "ai:assistants:tests:runs trigger", obj, format, transform)
}

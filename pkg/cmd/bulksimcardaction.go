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

var bulkSimCardActionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "This API fetches information about a bulk SIM card action. A bulk SIM card\naction contains details about a collection of individual SIM card actions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleBulkSimCardActionsRetrieve,
	HideHelpCommand: true,
}

var bulkSimCardActionsList = cli.Command{
	Name:    "list",
	Usage:   "This API lists a paginated collection of bulk SIM card actions. A bulk SIM card\naction contains details about a collection of individual SIM card actions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-action-type",
			Usage:     "Filter by action type.",
			QueryPath: "filter[action_type]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page.",
			Default:   20,
			QueryPath: "page[size]",
		},
	},
	Action:          handleBulkSimCardActionsList,
	HideHelpCommand: true,
}

func handleBulkSimCardActionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.BulkSimCardActions.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "bulk-sim-card-actions retrieve", obj, format, transform)
}

func handleBulkSimCardActionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BulkSimCardActionListParams{}

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
		_, err = client.BulkSimCardActions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "bulk-sim-card-actions list", obj, format, transform)
	} else {
		iter := client.BulkSimCardActions.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "bulk-sim-card-actions list", iter, format, transform)
	}
}

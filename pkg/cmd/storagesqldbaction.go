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

var storageSqldbsActionsQuery = cli.Command{
	Name:    "query",
	Usage:   "Runs SQL against the database and returns the resulting rows — empty for\nstatements that return none, such as DDL. Bind positional `?` placeholders with\n`params` rather than interpolating values into the SQL string.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "sql",
			Usage:    "The SQL to run. Use positional `?` placeholders and supply the values in `params` rather than interpolating them into this string.",
			Required: true,
			BodyPath: "sql",
		},
		&requestflag.Flag[[]any]{
			Name:     "param",
			Usage:    "Positional bind parameters, in placeholder order. Each value is a string, a number, a boolean, or null; booleans are cast to `1`/`0`. The count must match the number of `?` placeholders exactly — a mismatch is rejected with 422 rather than binding null for the ones you left out. (Not enforced for multi-statement scripts or named parameters, where the placeholder count is not the number bound.)",
			BodyPath: "params",
		},
	},
	Action:          handleStorageSqldbsActionsQuery,
	HideHelpCommand: true,
}

func handleStorageSqldbsActionsQuery(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.StorageSqldbActionQueryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Sqldbs.Actions.Query(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "storage:sqldbs:actions query",
		Transform:      transform,
	})
}

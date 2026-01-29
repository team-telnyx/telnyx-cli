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

var storageMigrationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Initiate a migration of data from an external provider into Telnyx Cloud\nStorage. Currently, only S3 is supported.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source-id",
			Usage:    "ID of the Migration Source from which to migrate data.",
			Required: true,
			BodyPath: "source_id",
		},
		&requestflag.Flag[string]{
			Name:     "target-bucket-name",
			Usage:    "Bucket name to migrate the data into. Will default to the same name as the `source_bucket_name`.",
			Required: true,
			BodyPath: "target_bucket_name",
		},
		&requestflag.Flag[string]{
			Name:     "target-region",
			Usage:    "Telnyx Cloud Storage region to migrate the data to.",
			Required: true,
			BodyPath: "target_region",
		},
		&requestflag.Flag[bool]{
			Name:     "refresh",
			Usage:    "If true, will continue to poll the source bucket to ensure new data is continually migrated over.",
			BodyPath: "refresh",
		},
	},
	Action:          handleStorageMigrationsCreate,
	HideHelpCommand: true,
}

var storageMigrationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Migration",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleStorageMigrationsRetrieve,
	HideHelpCommand: true,
}

var storageMigrationsList = cli.Command{
	Name:            "list",
	Usage:           "List all Migrations",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleStorageMigrationsList,
	HideHelpCommand: true,
}

func handleStorageMigrationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.StorageMigrationNewParams{}

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
	_, err = client.Storage.Migrations.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:migrations create", obj, format, transform)
}

func handleStorageMigrationsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Storage.Migrations.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:migrations retrieve", obj, format, transform)
}

func handleStorageMigrationsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Migrations.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:migrations list", obj, format, transform)
}

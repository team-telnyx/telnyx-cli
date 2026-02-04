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

var storageMigrationSourcesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a source from which data can be migrated from.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Usage:    "Bucket name to migrate the data from.",
			Required: true,
			BodyPath: "bucket_name",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "Cloud provider from which to migrate data. Use 'telnyx' if you want to migrate data from one Telnyx bucket to another.",
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "provider-auth",
			Required: true,
			BodyPath: "provider_auth",
		},
		&requestflag.Flag[string]{
			Name:     "source-region",
			Usage:    "For intra-Telnyx buckets migration, specify the source bucket region in this field.",
			BodyPath: "source_region",
		},
	},
	Action:          handleStorageMigrationSourcesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"provider-auth": {
		&requestflag.InnerFlag[string]{
			Name:       "provider-auth.access-key",
			Usage:      "AWS Access Key. For Telnyx-to-Telnyx migrations, use your Telnyx API key here.",
			InnerField: "access_key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "provider-auth.secret-access-key",
			Usage:      "AWS Secret Access Key. For Telnyx-to-Telnyx migrations, use your Telnyx API key here as well.",
			InnerField: "secret_access_key",
		},
	},
})

var storageMigrationSourcesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Migration Source",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleStorageMigrationSourcesRetrieve,
	HideHelpCommand: true,
}

var storageMigrationSourcesList = cli.Command{
	Name:            "list",
	Usage:           "List all Migration Sources",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleStorageMigrationSourcesList,
	HideHelpCommand: true,
}

var storageMigrationSourcesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Migration Source",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleStorageMigrationSourcesDelete,
	HideHelpCommand: true,
}

func handleStorageMigrationSourcesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.StorageMigrationSourceNewParams{}

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
	_, err = client.Storage.MigrationSources.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:migration-sources create", obj, format, transform)
}

func handleStorageMigrationSourcesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Storage.MigrationSources.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:migration-sources retrieve", obj, format, transform)
}

func handleStorageMigrationSourcesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Storage.MigrationSources.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:migration-sources list", obj, format, transform)
}

func handleStorageMigrationSourcesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Storage.MigrationSources.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:migration-sources delete", obj, format, transform)
}

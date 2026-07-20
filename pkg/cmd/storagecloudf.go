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

var storageCloudfsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a CloudFS filesystem. Provisioning is synchronous — typically a few\nseconds, up to a few minutes — and the filesystem is returned with status\n`ready`, together with its S3 bucket and metadata connection details. This\nresponse is the only time the filesystem's `meta_token` — and the\ncredential-bearing `meta_url` — are returned; store them securely. If the token\nis lost, issue a new one with the rotate-meta-token action. Names are unique\nwithin your organization: creating with an existing name returns a `422`.\nRequests are idempotent: retrying with the same `Idempotency-Key` within 24\nhours replays the original response instead of creating another filesystem.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Filesystem name, unique within your organization. Names are trimmed and lowercased; after normalization they may contain lowercase letters, numbers, `.`, `_`, and `-` only.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region where the filesystem's storage and metadata are provisioned.",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleStorageCloudfsCreate,
	HideHelpCommand: true,
}

var storageCloudfsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a CloudFS filesystem by its ID. The returned `meta_url` omits the\ncredential — the metadata token is only ever returned by create and\nrotate-meta-token. A filesystem whose last lifecycle action failed includes a\ncustomer-safe `error` message.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleStorageCloudfsRetrieve,
	HideHelpCommand: true,
}

var storageCloudfsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a CloudFS filesystem. Only `name` can be changed; other fields are\nimmutable and unknown fields are rejected with a `400`. Renaming to a name that\nalready exists in your organization returns a `422`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New filesystem name, unique within your organization. Names are trimmed and lowercased; after normalization they may contain lowercase letters, numbers, `.`, `_`, and `-` only.",
			BodyPath: "name",
		},
	},
	Action:          handleStorageCloudfsUpdate,
	HideHelpCommand: true,
}

var storageCloudfsList = cli.Command{
	Name:    "list",
	Usage:   "Lists the CloudFS filesystems for the authenticated user's organization. Results\nuse cursor-based pagination: fetch the next page by passing `meta.cursors.after`\nas `page[after]`, or follow the `meta.next` URL.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "Return only the filesystem whose name matches exactly.",
			QueryPath: "filter[name]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-region",
			Usage:     "Return only filesystems in this region.",
			QueryPath: "filter[region]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     "Return only filesystems with this status. Unrecognized values are ignored.",
			QueryPath: "filter[status]",
		},
		&requestflag.Flag[string]{
			Name:      "page-after",
			Usage:     "Opaque cursor from a previous response's `meta.cursors.after`; returns the page after it. Mutually exclusive with `page[before]`.",
			QueryPath: "page[after]",
		},
		&requestflag.Flag[string]{
			Name:      "page-before",
			Usage:     "Opaque cursor from a previous response's `meta.cursors.before`; returns the page before it. Mutually exclusive with `page[after]`.",
			QueryPath: "page[before]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-limit",
			Usage:     "The number of filesystems to return per page. Values above 250 are treated as 250.",
			Default:   20,
			QueryPath: "page[limit]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sort order for the results: a field name for ascending, or the field name prefixed with `-` for descending.",
			Default:   "-created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleStorageCloudfsList,
	HideHelpCommand: true,
}

var storageCloudfsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a CloudFS filesystem, removing its S3 bucket and its\nmetadata database. Deletion is synchronous: the response returns the\nfilesystem's final state with status `deleted`. There is no restore. A\nfilesystem that is still `provisioning` returns a `409`. If the filesystem still\ncontains data, the request may be rejected with a `409` — drain the bucket and\nretry.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleStorageCloudfsDelete,
	HideHelpCommand: true,
}

func handleStorageCloudfsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.StorageCloudfNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Cloudfs.New(ctx, params, options...)
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
		Title:          "storage:cloudfs create",
		Transform:      transform,
	})
}

func handleStorageCloudfsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Storage.Cloudfs.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "storage:cloudfs retrieve",
		Transform:      transform,
	})
}

func handleStorageCloudfsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.StorageCloudfUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Cloudfs.Update(
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
		Title:          "storage:cloudfs update",
		Transform:      transform,
	})
}

func handleStorageCloudfsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.StorageCloudfListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Cloudfs.List(ctx, params, options...)
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
		Title:          "storage:cloudfs list",
		Transform:      transform,
	})
}

func handleStorageCloudfsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Storage.Cloudfs.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "storage:cloudfs delete",
		Transform:      transform,
	})
}

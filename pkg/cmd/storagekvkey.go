// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/stainless-sdks/telnyx-go/v4"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var storageKvsKeysRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the raw stored value for a key. The response body is the value exactly\nas it was written; the `Content-Type` header echoes the value's stored content\ntype (defaults to `application/octet-stream`).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "key",
			Required:  true,
			PathParam: "key",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleStorageKvsKeysRetrieve,
	HideHelpCommand: true,
}

var storageKvsKeysList = cli.Command{
	Name:    "list",
	Usage:   "Lists the keys in a namespace. Returns key names and metadata only, never\nvalues. Results are paginated with `limit` and an opaque `cursor`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response's `meta.cursor`.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of keys to return. Values above 1000 are treated as 1000.",
			Default:   1000,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "prefix",
			Usage:     "Return only keys that start with this prefix.",
			QueryPath: "prefix",
		},
	},
	Action:          handleStorageKvsKeysList,
	HideHelpCommand: true,
}

var storageKvsKeysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a key. Idempotent: deleting a key that does not exist still succeeds.\nThe namespace itself must exist and be provisioned.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "key",
			Required:  true,
			PathParam: "key",
		},
	},
	Action:          handleStorageKvsKeysDelete,
	HideHelpCommand: true,
}

var storageKvsKeysSet = cli.Command{
	Name:    "set",
	Usage:   "Creates or replaces the value for a key. The request body is stored verbatim as\nthe value — no base64, no JSON envelope — up to 1 MiB. The request's\n`Content-Type` header is stored with the value and echoed back on retrieval.\nReturns `201` when the key is created and `200` when an existing key is updated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "key",
			Required:  true,
			PathParam: "key",
		},
		&requestflag.Flag[string]{
			Name:      "body",
			Usage:     "Raw value bytes, stored verbatim.",
			Required:  true,
			BodyRoot:  true,
			FileInput: true,
		},
		&requestflag.Flag[int64]{
			Name:      "ttl-secs",
			Usage:     "Time-to-live in seconds. When set, the key expires and is deleted after this duration. Requires a namespace provisioned with TTL support; namespaces without it return a `409`.",
			QueryPath: "ttl_secs",
		},
	},
	Action:          handleStorageKvsKeysSet,
	HideHelpCommand: true,
}

func handleStorageKvsKeysRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key") && len(unusedArgs) > 0 {
		cmd.Set("key", unusedArgs[0])
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

	params := telnyx.StorageKvKeyGetParams{
		ID: cmd.Value("id").(string),
	}

	response, err := client.Storage.Kvs.Keys.Get(
		ctx,
		cmd.Value("key").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}

func handleStorageKvsKeysList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.StorageKvKeyListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Kvs.Keys.List(
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
		Title:          "storage:kvs:keys list",
		Transform:      transform,
	})
}

func handleStorageKvsKeysDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key") && len(unusedArgs) > 0 {
		cmd.Set("key", unusedArgs[0])
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

	params := telnyx.StorageKvKeyDeleteParams{
		ID: cmd.Value("id").(string),
	}

	return client.Storage.Kvs.Keys.Delete(
		ctx,
		cmd.Value("key").(string),
		params,
		options...,
	)
}

func handleStorageKvsKeysSet(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key") && len(unusedArgs) > 0 {
		cmd.Set("key", unusedArgs[0])
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

	params := telnyx.StorageKvKeySetParams{
		ID: cmd.Value("id").(string),
	}

	return client.Storage.Kvs.Keys.Set(
		ctx,
		cmd.Value("key").(string),
		params,
		options...,
	)
}

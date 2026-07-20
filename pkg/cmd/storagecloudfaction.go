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

var storageCloudfsActionsRotateMetaToken = cli.Command{
	Name:    "rotate-meta-token",
	Usage:   "Issues a new metadata access token for the filesystem and returns the full\nfilesystem, including the new `meta_token` and credential-bearing `meta_url`.\nThe previous token stops authenticating immediately; the metadata database and\nS3 bucket are unchanged. The request takes no body. Allowed while the filesystem\nis `ready` or `needs_format`; otherwise returns a `409`. Retrying with the same\n`Idempotency-Key` within 24 hours replays the original response — including the\nsame token — instead of rotating again.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleStorageCloudfsActionsRotateMetaToken,
	HideHelpCommand: true,
}

func handleStorageCloudfsActionsRotateMetaToken(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.StorageCloudfActionRotateMetaTokenParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Cloudfs.Actions.RotateMetaToken(
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
		Title:          "storage:cloudfs:actions rotate-meta-token",
		Transform:      transform,
	})
}

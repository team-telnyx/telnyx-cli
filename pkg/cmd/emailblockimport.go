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

var emailBlocksImportCreate = cli.Command{
	Name:    "create",
	Usage:   "Accepts `multipart/form-data` with a `file` field (the CSV) and an optional\n`block_ttl_days` (integer >0, default 30). Validates:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file",
			Usage:     "The CSV file (Plug.Upload). Missing/non-upload → 400.",
			Required:  true,
			BodyPath:  "file",
			FileInput: true,
		},
		&requestflag.Flag[int64]{
			Name:     "block-ttl-days",
			Usage:    "TTL for imported `manual_block` rows; other reasons get `expires_at: null`. Invalid/missing → falls back to 30.",
			Default:  30,
			BodyPath: "block_ttl_days",
		},
	},
	Action:          handleEmailBlocksImportCreate,
	HideHelpCommand: true,
}

var emailBlocksImportRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Account-scoped fetch (cross-account → 404; malformed UUID → 404). Nullable\nfields are omitted until terminal: `provider`/`completed_at` when nil;\n`processed_rows`/`created_count`/`existing_count`/ `skipped_count`/`error_count`\nonly when `status == completed`; `errors` only when non-empty; `failure_reason`\nonly on terminal failure.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailBlocksImportRetrieve,
	HideHelpCommand: true,
}

func handleEmailBlocksImportCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EmailBlockImportNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailBlocks.Import.New(ctx, params, options...)
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
		Title:          "email-blocks:import create",
		Transform:      transform,
	})
}

func handleEmailBlocksImportRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EmailBlocks.Import.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email-blocks:import retrieve",
		Transform:      transform,
	})
}

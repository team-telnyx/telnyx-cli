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

var faxesCreate = cli.Command{
	Name:    "create",
	Usage:   "Send a fax. Files have size limits and page count limit validations. If a file\nis bigger than 50MB or has more than 350 pages it will fail with\n`file_size_limit_exceeded` and `page_count_limit_exceeded` respectively.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Usage:    "The connection ID to send the fax with.",
			Required: true,
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "The phone number, in E.164 format, the fax will be sent from.",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "The phone number, in E.164 format, the fax will be sent to or SIP URI",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[int64]{
			Name:     "black-threshold",
			Usage:    "The black threshold percentage for monochrome faxes. Only applicable if `monochrome` is set to `true`.",
			Default:  95,
			BodyPath: "black_threshold",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[string]{
			Name:     "from-display-name",
			Usage:    "The `from_display_name` string to be used as the caller id name (SIP From Display Name) presented to the destination (`to` number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and -_~!.+ special characters. If ommited, the display name will be the same as the number in the `from` field.",
			BodyPath: "from_display_name",
		},
		&requestflag.Flag[string]{
			Name:     "media-name",
			Usage:    "The media_name used for the fax's media. Must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. media_name and media_url/contents can't be submitted together.",
			BodyPath: "media_name",
		},
		&requestflag.Flag[string]{
			Name:     "media-url",
			Usage:    "The URL (or list of URLs) to the PDF used for the fax's media. media_url and media_name/contents can't be submitted together.",
			BodyPath: "media_url",
		},
		&requestflag.Flag[bool]{
			Name:     "monochrome",
			Usage:    "The flag to enable monochrome, true black and white fax results.",
			Default:  false,
			BodyPath: "monochrome",
		},
		&requestflag.Flag[string]{
			Name:     "preview-format",
			Usage:    "The format for the preview file in case the `store_preview` is `true`.",
			Default:  "tiff",
			BodyPath: "preview_format",
		},
		&requestflag.Flag[string]{
			Name:     "quality",
			Usage:    "The quality of the fax. The `ultra` settings provides the highest quality available, but also present longer fax processing times. `ultra_light` is best suited for images, wihle `ultra_dark` is best suited for text.",
			Default:  "high",
			BodyPath: "quality",
		},
		&requestflag.Flag[bool]{
			Name:     "store-media",
			Usage:    "Should fax media be stored on temporary URL. It does not support media_name, they can't be submitted together.",
			Default:  false,
			BodyPath: "store_media",
		},
		&requestflag.Flag[bool]{
			Name:     "store-preview",
			Usage:    "Should fax preview be stored on temporary URL.",
			Default:  false,
			BodyPath: "store_preview",
		},
		&requestflag.Flag[bool]{
			Name:     "t38-enabled",
			Usage:    "The flag to disable the T.38 protocol.",
			Default:  true,
			BodyPath: "t38_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Use this field to override the URL to which Telnyx will send subsequent webhooks for this fax.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleFaxesCreate,
	HideHelpCommand: true,
}

var faxesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "View a fax",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFaxesRetrieve,
	HideHelpCommand: true,
}

var faxesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "View a list of faxes",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[created_at][gte], filter[created_at][gt], filter[created_at][lte], filter[created_at][lt], filter[direction][eq], filter[from][eq], filter[to][eq]",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleFaxesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.created-at",
			Usage:      "Date range filtering operations for fax creation timestamp",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.direction",
			Usage:      "Direction filtering operations",
			InnerField: "direction",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.from",
			Usage:      "From number filtering operations",
			InnerField: "from",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.to",
			Usage:      "To number filtering operations",
			InnerField: "to",
		},
	},
})

var faxesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a fax",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFaxesDelete,
	HideHelpCommand: true,
}

func handleFaxesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FaxNewParams{}

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
	_, err = client.Faxes.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "faxes create", obj, format, transform)
}

func handleFaxesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Faxes.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "faxes retrieve", obj, format, transform)
}

func handleFaxesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FaxListParams{}

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
		_, err = client.Faxes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "faxes list", obj, format, transform)
	} else {
		iter := client.Faxes.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "faxes list", iter, format, transform)
	}
}

func handleFaxesDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Faxes.Delete(ctx, cmd.Value("id").(string), options...)
}

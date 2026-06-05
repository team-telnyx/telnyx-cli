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

var dirPhoneNumbersList = cli.Command{
	Name:    "list",
	Usage:   "List the phone numbers registered under a DIR. The enterprise is resolved\nserver-side from the DIR id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "1-based page number. Out-of-range values return an empty page with correct meta.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Items per page. Maximum 250; values above are clamped to 250.",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by phone-number status.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleDirPhoneNumbersList,
	HideHelpCommand: true,
}

var dirPhoneNumbersAdd = requestflag.WithInnerFlags(cli.Command{
	Name:    "add",
	Usage:   "Register phone numbers under a DIR. The enterprise is resolved server-side from\nthe DIR id. Same body, failure modes, and batch semantics whichever path form\nyou use.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "document",
			Usage:    "Supporting documents covering this batch. At least one entry with `document_type: letter_of_authorization` is required — the LOA authorises Telnyx to register these numbers under the DIR. Each `document_id` must come from the Telnyx Documents API. Additional document types (e.g. business registration) may be included alongside the LOA.",
			Required: true,
			BodyPath: "documents",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "1–15 phone numbers in E.164 format. 10-digit US numbers are auto-prefixed with `1`.",
			Required: true,
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleDirPhoneNumbersAdd,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"document": {
		&requestflag.InnerFlag[string]{
			Name:       "document.document-id",
			Usage:      "Id returned by the Telnyx Documents API after you upload the file (upload via `POST /v2/documents`; see https://developers.telnyx.com/api/documents).",
			InnerField: "document_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.document-type",
			Usage:      "Type of supporting document. Pick the closest match to what the file actually contains; `other` triggers manual vetting and may slow approval. The matching short_name reference list is at `GET /v2/dir/document_types`.",
			InnerField: "document_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.description",
			InnerField: "description",
		},
	},
})

var dirPhoneNumbersRemove = cli.Command{
	Name:    "remove",
	Usage:   "Deregister phone numbers from a DIR. The enterprise is resolved server-side from\nthe DIR id. Returns a partial-success envelope.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Required: true,
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleDirPhoneNumbersRemove,
	HideHelpCommand: true,
}

func handleDirPhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dir-id") && len(unusedArgs) > 0 {
		cmd.Set("dir-id", unusedArgs[0])
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

	params := telnyx.DirPhoneNumberListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Dir.PhoneNumbers.List(
			ctx,
			cmd.Value("dir-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "dir:phone-numbers list",
			Transform:      transform,
		})
	} else {
		iter := client.Dir.PhoneNumbers.ListAutoPaging(
			ctx,
			cmd.Value("dir-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "dir:phone-numbers list",
			Transform:      transform,
		})
	}
}

func handleDirPhoneNumbersAdd(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dir-id") && len(unusedArgs) > 0 {
		cmd.Set("dir-id", unusedArgs[0])
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

	params := telnyx.DirPhoneNumberAddParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.PhoneNumbers.Add(
		ctx,
		cmd.Value("dir-id").(string),
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
		Title:          "dir:phone-numbers add",
		Transform:      transform,
	})
}

func handleDirPhoneNumbersRemove(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dir-id") && len(unusedArgs) > 0 {
		cmd.Set("dir-id", unusedArgs[0])
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

	params := telnyx.DirPhoneNumberRemoveParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.PhoneNumbers.Remove(
		ctx,
		cmd.Value("dir-id").(string),
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
		Title:          "dir:phone-numbers remove",
		Transform:      transform,
	})
}

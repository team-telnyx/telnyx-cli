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

var callReasonsList = cli.Command{
	Name:    "list",
	Usage:   "Telnyx maintains a library of pre-vetted call-reason phrases (e.g. \"Appointment\nreminders\", \"Billing inquiries\") that carry through DIR vetting smoothly. You\ncan use any string that fits your use case in `DirCreateRequest.call_reasons`,\nbut matching one of these reduces the chance the vetting team flags the phrasing\nfor clarification.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "1-based page number. Out-of-range values return an empty page with correct meta.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Items per page. Default `100` for this endpoint (the call-reason library is small and most callers want the whole list in one call). Maximum 250; values above are clamped to 250.",
			Default:   100,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleCallReasonsList,
	HideHelpCommand: true,
}

var callReasonsValidate = cli.Command{
	Name:    "validate",
	Usage:   "Check up to 10 candidate `call_reasons` strings against Telnyx's vetting\nheuristics before sending them on a DIR create or update. The endpoint flags\nstrings that are likely to be rejected during vetting (too generic, banned\nphrases, length issues, etc.) so you can fix them up front.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "body",
			Usage:    "**Bare JSON array** of candidate call-reason strings (NOT an object - there is no top-level `call_reasons` key on this endpoint). 1–10 strings, each ≤64 characters.",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleCallReasonsValidate,
	HideHelpCommand: true,
}

func handleCallReasonsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.CallReasonListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.CallReasons.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "call-reasons list",
			Transform:      transform,
		})
	} else {
		iter := client.CallReasons.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "call-reasons list",
			Transform:      transform,
		})
	}
}

func handleCallReasonsValidate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.CallReasonValidateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CallReasons.Validate(ctx, params, options...)
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
		Title:          "call-reasons validate",
		Transform:      transform,
	})
}

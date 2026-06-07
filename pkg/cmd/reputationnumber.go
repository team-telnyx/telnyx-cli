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

var reputationNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Convenience alias for\n`GET /v2/enterprises/{enterprise_id}/reputation/numbers/{phone_number}`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Required:  true,
			PathParam: "phone_number",
		},
		&requestflag.Flag[bool]{
			Name:      "fresh",
			Usage:     "When true, fetches fresh reputation data (incurs API cost). When false (default), returns cached data.",
			Default:   false,
			QueryPath: "fresh",
		},
	},
	Action:          handleReputationNumbersRetrieve,
	HideHelpCommand: true,
}

var reputationNumbersList = cli.Command{
	Name:    "list",
	Usage:   "Convenience alias for `GET /v2/enterprises/{enterprise_id}/reputation/numbers`\nthat returns numbers across every enterprise you own. Useful when you don't want\nto look up the enterprise id first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-enterprise-id",
			Usage:     "Filter by enterprise ID.",
			QueryPath: "filter[enterprise_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-phone-number-contains",
			Usage:     "Partial match on phone number. Must contain at least 5 digits.",
			QueryPath: "filter[phone_number][contains]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-phone-number-eq",
			Usage:     "Exact phone-number match (E.164).",
			QueryPath: "filter[phone_number][eq]",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleReputationNumbersList,
	HideHelpCommand: true,
}

var reputationNumbersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Convenience alias for\n`DELETE /v2/enterprises/{enterprise_id}/reputation/numbers/{phone_number}`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Required:  true,
			PathParam: "phone_number",
		},
	},
	Action:          handleReputationNumbersDelete,
	HideHelpCommand: true,
}

func handleReputationNumbersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
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

	params := telnyx.ReputationNumberGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Reputation.Numbers.Get(
		ctx,
		cmd.Value("phone-number").(string),
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
		Title:          "reputation:numbers retrieve",
		Transform:      transform,
	})
}

func handleReputationNumbersList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.ReputationNumberListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Reputation.Numbers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "reputation:numbers list",
			Transform:      transform,
		})
	} else {
		iter := client.Reputation.Numbers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "reputation:numbers list",
			Transform:      transform,
		})
	}
}

func handleReputationNumbersDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
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

	return client.Reputation.Numbers.Delete(ctx, cmd.Value("phone-number").(string), options...)
}

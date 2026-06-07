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

var enterprisesReputationNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve one registered number with its latest reputation snapshot. The\n`phone_number` path parameter is in E.164 format and must be URL-encoded (e.g.\n`%2B19493253498`).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
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
	Action:          handleEnterprisesReputationNumbersRetrieve,
	HideHelpCommand: true,
}

var enterprisesReputationNumbersList = cli.Command{
	Name:    "list",
	Usage:   "Paginated list of phone numbers registered for reputation monitoring under this\nenterprise. The response includes the latest reputation snapshot per number\nwhere one has been collected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
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
			Usage:     "Items per page. Default 10. Maximum 250; values above are clamped to 250.",
			Default:   10,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Usage:     "Filter by specific phone number (E.164 format).",
			QueryPath: "phone_number",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEnterprisesReputationNumbersList,
	HideHelpCommand: true,
}

var enterprisesReputationNumbersAssociate = cli.Command{
	Name:    "associate",
	Usage:   "Add up to 100 phone numbers to reputation monitoring on this enterprise. Each\nmust be in E.164 format (`+1NPANXXXXXX` for US/CA) and belong to your Telnyx\nphone-number inventory.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "1–100 phone numbers in E.164 format with a leading `+`.",
			Required: true,
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleEnterprisesReputationNumbersAssociate,
	HideHelpCommand: true,
}

var enterprisesReputationNumbersDisassociate = cli.Command{
	Name:    "disassociate",
	Usage:   "Stop tracking the reputation of this phone number. The number itself remains in\nyour inventory; only the reputation registration is removed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Required:  true,
			PathParam: "phone_number",
		},
	},
	Action:          handleEnterprisesReputationNumbersDisassociate,
	HideHelpCommand: true,
}

var enterprisesReputationNumbersRefresh = cli.Command{
	Name:    "refresh",
	Usage:   "Immediately refresh the stored reputation data for the listed numbers. This is\nin addition to the periodic refresh determined by `check_frequency`. Up to 100\nnumbers per call. The response carries the kicked-off jobs; the actual refresh\nruns asynchronously.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "Phone numbers to refresh reputation data for. 1–100 numbers per request, each in E.164 format. Reputation refreshes are subject to per-enterprise rate limits.",
			Required: true,
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleEnterprisesReputationNumbersRefresh,
	HideHelpCommand: true,
}

func handleEnterprisesReputationNumbersRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseReputationNumberGetParams{
		EnterpriseID: cmd.Value("enterprise-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.Numbers.Get(
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
		Title:          "enterprises:reputation:numbers retrieve",
		Transform:      transform,
	})
}

func handleEnterprisesReputationNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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

	params := telnyx.EnterpriseReputationNumberListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Enterprises.Reputation.Numbers.List(
			ctx,
			cmd.Value("enterprise-id").(string),
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
			Title:          "enterprises:reputation:numbers list",
			Transform:      transform,
		})
	} else {
		iter := client.Enterprises.Reputation.Numbers.ListAutoPaging(
			ctx,
			cmd.Value("enterprise-id").(string),
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
			Title:          "enterprises:reputation:numbers list",
			Transform:      transform,
		})
	}
}

func handleEnterprisesReputationNumbersAssociate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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

	params := telnyx.EnterpriseReputationNumberAssociateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.Numbers.Associate(
		ctx,
		cmd.Value("enterprise-id").(string),
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
		Title:          "enterprises:reputation:numbers associate",
		Transform:      transform,
	})
}

func handleEnterprisesReputationNumbersDisassociate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseReputationNumberDisassociateParams{
		EnterpriseID: cmd.Value("enterprise-id").(string),
	}

	return client.Enterprises.Reputation.Numbers.Disassociate(
		ctx,
		cmd.Value("phone-number").(string),
		params,
		options...,
	)
}

func handleEnterprisesReputationNumbersRefresh(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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

	params := telnyx.EnterpriseReputationNumberRefreshParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.Numbers.Refresh(
		ctx,
		cmd.Value("enterprise-id").(string),
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
		Title:          "enterprises:reputation:numbers refresh",
		Transform:      transform,
	})
}

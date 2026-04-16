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

var enterprisesReputationNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get detailed reputation data for a specific phone number associated with an\nenterprise.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "enterprise-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "fresh",
			Usage:     "When true, fetches fresh reputation data (incurs API cost). When false, returns cached data.",
			Default:   false,
			QueryPath: "fresh",
		},
	},
	Action:          handleEnterprisesReputationNumbersRetrieve,
	HideHelpCommand: true,
}

var enterprisesReputationNumbersList = cli.Command{
	Name:    "list",
	Usage:   "List all phone numbers associated with an enterprise for Number Reputation\nmonitoring.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "enterprise-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (1-indexed)",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of items per page",
			Default:   10,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "phone-number",
			Usage:     "Filter by specific phone number (E.164 format)",
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
	Usage:   "Associate one or more phone numbers with an enterprise for Number Reputation\nmonitoring.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "enterprise-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "List of phone numbers to associate for reputation monitoring (max 100)",
			Required: true,
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleEnterprisesReputationNumbersAssociate,
	HideHelpCommand: true,
}

var enterprisesReputationNumbersDisassociate = cli.Command{
	Name:    "disassociate",
	Usage:   "Remove a phone number from Number Reputation monitoring for an enterprise.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "enterprise-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
		},
	},
	Action:          handleEnterprisesReputationNumbersDisassociate,
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

	params := telnyx.EnterpriseReputationNumberGetParams{
		EnterpriseID: cmd.Value("enterprise-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "enterprises:reputation:numbers retrieve", obj, format, explicitFormat, transform)
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

	params := telnyx.EnterpriseReputationNumberListParams{}

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
		return ShowJSON(os.Stdout, os.Stderr, "enterprises:reputation:numbers list", obj, format, explicitFormat, transform)
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
		return ShowJSONIterator(os.Stdout, os.Stderr, "enterprises:reputation:numbers list", iter, format, explicitFormat, transform, maxItems)
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

	params := telnyx.EnterpriseReputationNumberNewParams{}

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
	_, err = client.Enterprises.Reputation.Numbers.New(
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
	return ShowJSON(os.Stdout, os.Stderr, "enterprises:reputation:numbers associate", obj, format, explicitFormat, transform)
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

	params := telnyx.EnterpriseReputationNumberDeleteParams{
		EnterpriseID: cmd.Value("enterprise-id").(string),
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

	return client.Enterprises.Reputation.Numbers.Delete(
		ctx,
		cmd.Value("phone-number").(string),
		params,
		options...,
	)
}

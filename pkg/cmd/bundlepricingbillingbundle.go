// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var bundlePricingBillingBundlesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single bundle by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bundle-id",
			Usage:    "Billing bundle's ID, this is used to identify the billing bundle in the API.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleBundlePricingBillingBundlesRetrieve,
	HideHelpCommand: true,
}

var bundlePricingBillingBundlesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get all allowed bundles.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Supports filtering by country_iso and resource. Examples: filter[country_iso]=US or filter[resource]=+15617819942",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleBundlePricingBillingBundlesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.country-iso",
			Usage:      "Filter by country code.",
			InnerField: "country_iso",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.resource",
			Usage:      "Filter by resource.",
			InnerField: "resource",
		},
	},
	"page": {
		&requestflag.InnerFlag[int64]{
			Name:       "page.number",
			Usage:      "The page number to load.",
			InnerField: "number",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.size",
			Usage:      "The size of the page.",
			InnerField: "size",
		},
	},
})

func handleBundlePricingBillingBundlesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bundle-id") && len(unusedArgs) > 0 {
		cmd.Set("bundle-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BundlePricingBillingBundleGetParams{}

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
	_, err = client.BundlePricing.BillingBundles.Get(
		ctx,
		cmd.Value("bundle-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "bundle-pricing:billing-bundles retrieve", obj, format, transform)
}

func handleBundlePricingBillingBundlesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BundlePricingBillingBundleListParams{}

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
		_, err = client.BundlePricing.BillingBundles.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "bundle-pricing:billing-bundles list", obj, format, transform)
	} else {
		iter := client.BundlePricing.BillingBundles.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "bundle-pricing:billing-bundles list", iter, format, transform)
	}
}

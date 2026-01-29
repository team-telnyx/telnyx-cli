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

var bundlePricingUserBundlesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates multiple user bundles for the user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Idempotency key for the request. Can be any UUID, but should always be unique for each request.",
			BodyPath: "idempotency_key",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "item",
			BodyPath: "items",
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleBundlePricingUserBundlesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"item": {
		&requestflag.InnerFlag[string]{
			Name:       "item.billing-bundle-id",
			Usage:      "Quantity of user bundles to order.",
			InnerField: "billing_bundle_id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "item.quantity",
			Usage:      "Quantity of user bundles to order.",
			InnerField: "quantity",
		},
	},
})

var bundlePricingUserBundlesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a user bundle by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-bundle-id",
			Usage:    "User bundle's ID, this is used to identify the user bundle in the API.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleBundlePricingUserBundlesRetrieve,
	HideHelpCommand: true,
}

var bundlePricingUserBundlesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of user bundles.",
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
	Action:          handleBundlePricingUserBundlesList,
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

var bundlePricingUserBundlesDeactivate = cli.Command{
	Name:    "deactivate",
	Usage:   "Deactivates a user bundle by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-bundle-id",
			Usage:    "User bundle's ID, this is used to identify the user bundle in the API.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleBundlePricingUserBundlesDeactivate,
	HideHelpCommand: true,
}

var bundlePricingUserBundlesListResources = cli.Command{
	Name:    "list-resources",
	Usage:   "Retrieves the resources of a user bundle by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-bundle-id",
			Usage:    "User bundle's ID, this is used to identify the user bundle in the API.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleBundlePricingUserBundlesListResources,
	HideHelpCommand: true,
}

var bundlePricingUserBundlesListUnused = requestflag.WithInnerFlags(cli.Command{
	Name:    "list-unused",
	Usage:   "Returns all user bundles that aren't in use.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Supports filtering by country_iso and resource. Examples: filter[country_iso]=US or filter[resource]=+15617819942",
			QueryPath: "filter",
		},
		&requestflag.Flag[string]{
			Name:       "authorization-bearer",
			Usage:      "Authenticates the request with your Telnyx API V2 KEY",
			HeaderPath: "authorization_bearer",
		},
	},
	Action:          handleBundlePricingUserBundlesListUnused,
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
})

func handleBundlePricingUserBundlesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BundlePricingUserBundleNewParams{}

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
	_, err = client.BundlePricing.UserBundles.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "bundle-pricing:user-bundles create", obj, format, transform)
}

func handleBundlePricingUserBundlesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-bundle-id") && len(unusedArgs) > 0 {
		cmd.Set("user-bundle-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BundlePricingUserBundleGetParams{}

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
	_, err = client.BundlePricing.UserBundles.Get(
		ctx,
		cmd.Value("user-bundle-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "bundle-pricing:user-bundles retrieve", obj, format, transform)
}

func handleBundlePricingUserBundlesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BundlePricingUserBundleListParams{}

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
		_, err = client.BundlePricing.UserBundles.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "bundle-pricing:user-bundles list", obj, format, transform)
	} else {
		iter := client.BundlePricing.UserBundles.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "bundle-pricing:user-bundles list", iter, format, transform)
	}
}

func handleBundlePricingUserBundlesDeactivate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-bundle-id") && len(unusedArgs) > 0 {
		cmd.Set("user-bundle-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BundlePricingUserBundleDeactivateParams{}

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
	_, err = client.BundlePricing.UserBundles.Deactivate(
		ctx,
		cmd.Value("user-bundle-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "bundle-pricing:user-bundles deactivate", obj, format, transform)
}

func handleBundlePricingUserBundlesListResources(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-bundle-id") && len(unusedArgs) > 0 {
		cmd.Set("user-bundle-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BundlePricingUserBundleListResourcesParams{}

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
	_, err = client.BundlePricing.UserBundles.ListResources(
		ctx,
		cmd.Value("user-bundle-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "bundle-pricing:user-bundles list-resources", obj, format, transform)
}

func handleBundlePricingUserBundlesListUnused(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.BundlePricingUserBundleListUnusedParams{}

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
	_, err = client.BundlePricing.UserBundles.ListUnused(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "bundle-pricing:user-bundles list-unused", obj, format, transform)
}

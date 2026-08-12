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

var rcsBrandsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates an editable RCS brand draft. Creating the draft does not begin external\nreview.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "addresses",
			Required: true,
			BodyPath: "addresses",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "contacts",
			Usage:    "Named business contacts. Use the `brand` key for the required BRAND contact.",
			Required: true,
			BodyPath: "contacts",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Required: true,
			BodyPath: "display_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "identifiers",
			Usage:    "Named business identifiers. Use the `ein` key for the required EIN and `stock_symbol` for a public-profit brand's stock symbol.",
			Required: true,
			BodyPath: "identifiers",
		},
		&requestflag.Flag[string]{
			Name:     "legal-entity-type",
			Usage:    `Allowed values: "LIMITED_LIABILITY_COMPANY", "SOLE_PROPRIETORSHIP", "PARTNERSHIP", "CORPORATION", "S_CORPORATION".`,
			Required: true,
			BodyPath: "legal_entity_type",
		},
		&requestflag.Flag[string]{
			Name:     "legal-name",
			Required: true,
			BodyPath: "legal_name",
		},
		&requestflag.Flag[string]{
			Name:     "organization-type",
			Usage:    `Allowed values: "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT", "UNKNOWN".`,
			Required: true,
			BodyPath: "organization_type",
		},
		&requestflag.Flag[string]{
			Name:     "website-url",
			Required: true,
			BodyPath: "website_url",
		},
		&requestflag.Flag[*string]{
			Name:     "profile-id",
			Usage:    "A Messaging Profile owned by the authenticated organization. Agents inherit this value when they do not provide their own profile.",
			BodyPath: "profile_id",
		},
	},
	Action:          handleRcsBrandsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"contacts": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "contacts.brand",
			InnerField: "brand",
		},
	},
	"identifiers": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "identifiers.ein",
			InnerField: "ein",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "identifiers.stock-symbol",
			InnerField: "stock_symbol",
		},
	},
})

var rcsBrandsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves an RCS brand and its current lifecycle status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRcsBrandsRetrieve,
	HideHelpCommand: true,
}

var rcsBrandsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates one or more fields on a brand while its status is `CREATED`. Submitted\nbrands cannot be changed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "addresses",
			BodyPath: "addresses",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "contacts",
			Usage:    "Named business contacts. Use the `brand` key for the required BRAND contact.",
			BodyPath: "contacts",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			BodyPath: "display_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "identifiers",
			Usage:    "Named business identifiers. Use the `ein` key for the required EIN and `stock_symbol` for a public-profit brand's stock symbol.",
			BodyPath: "identifiers",
		},
		&requestflag.Flag[string]{
			Name:     "legal-entity-type",
			Usage:    `Allowed values: "LIMITED_LIABILITY_COMPANY", "SOLE_PROPRIETORSHIP", "PARTNERSHIP", "CORPORATION", "S_CORPORATION".`,
			BodyPath: "legal_entity_type",
		},
		&requestflag.Flag[string]{
			Name:     "legal-name",
			BodyPath: "legal_name",
		},
		&requestflag.Flag[string]{
			Name:     "organization-type",
			Usage:    `Allowed values: "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT", "UNKNOWN".`,
			BodyPath: "organization_type",
		},
		&requestflag.Flag[string]{
			Name:     "profile-id",
			BodyPath: "profile_id",
		},
		&requestflag.Flag[string]{
			Name:     "website-url",
			BodyPath: "website_url",
		},
	},
	Action:          handleRcsBrandsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"contacts": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "contacts.brand",
			InnerField: "brand",
		},
	},
	"identifiers": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "identifiers.ein",
			InnerField: "ein",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "identifiers.stock-symbol",
			InnerField: "stock_symbol",
		},
	},
})

var rcsBrandsList = cli.Command{
	Name:            "list",
	Usage:           "Lists RCS brands owned by the authenticated organization.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRcsBrandsList,
	HideHelpCommand: true,
}

var rcsBrandsSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Starts asynchronous provider provisioning and external review for a brand.\nRepeating this request for an in-progress brand returns its current state\nwithout creating new work.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRcsBrandsSubmit,
	HideHelpCommand: true,
}

func handleRcsBrandsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.RcBrandNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Rcs.Brands.New(ctx, params, options...)
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
		Title:          "rcs:brands create",
		Transform:      transform,
	})
}

func handleRcsBrandsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rcs.Brands.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "rcs:brands retrieve",
		Transform:      transform,
	})
}

func handleRcsBrandsUpdate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.RcBrandUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Rcs.Brands.Update(
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
		Title:          "rcs:brands update",
		Transform:      transform,
	})
}

func handleRcsBrandsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Rcs.Brands.List(ctx, options...)
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
		Title:          "rcs:brands list",
		Transform:      transform,
	})
}

func handleRcsBrandsSubmit(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rcs.Brands.Submit(ctx, cmd.Value("id").(string), options...)
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
		Title:          "rcs:brands submit",
		Transform:      transform,
	})
}

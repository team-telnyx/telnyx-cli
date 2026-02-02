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

var dynamicEmergencyAddressesCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a dynamic emergency address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "administrative-area",
			Required: true,
			BodyPath: "administrative_area",
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Required: true,
			BodyPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:     "house-number",
			Required: true,
			BodyPath: "house_number",
		},
		&requestflag.Flag[string]{
			Name:     "locality",
			Required: true,
			BodyPath: "locality",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Required: true,
			BodyPath: "postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "street-name",
			Required: true,
			BodyPath: "street_name",
		},
		&requestflag.Flag[string]{
			Name:     "extended-address",
			BodyPath: "extended_address",
		},
		&requestflag.Flag[string]{
			Name:     "house-suffix",
			BodyPath: "house_suffix",
		},
		&requestflag.Flag[string]{
			Name:     "street-post-directional",
			BodyPath: "street_post_directional",
		},
		&requestflag.Flag[string]{
			Name:     "street-pre-directional",
			BodyPath: "street_pre_directional",
		},
		&requestflag.Flag[string]{
			Name:     "street-suffix",
			BodyPath: "street_suffix",
		},
	},
	Action:          handleDynamicEmergencyAddressesCreate,
	HideHelpCommand: true,
}

var dynamicEmergencyAddressesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the dynamic emergency address based on the ID provided",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDynamicEmergencyAddressesRetrieve,
	HideHelpCommand: true,
}

var dynamicEmergencyAddressesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns the dynamic emergency addresses according to filters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[status], filter[country_code]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
	},
	Action:          handleDynamicEmergencyAddressesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "Filter by country code.",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter by status.",
			InnerField: "status",
		},
	},
	"page": {
		&requestflag.InnerFlag[int64]{
			Name:       "page.number",
			Usage:      "The page number to load",
			InnerField: "number",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.size",
			Usage:      "The size of the page",
			InnerField: "size",
		},
	},
})

var dynamicEmergencyAddressesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the dynamic emergency address based on the ID provided",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDynamicEmergencyAddressesDelete,
	HideHelpCommand: true,
}

func handleDynamicEmergencyAddressesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DynamicEmergencyAddressNewParams{}

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
	_, err = client.DynamicEmergencyAddresses.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "dynamic-emergency-addresses create", obj, format, transform)
}

func handleDynamicEmergencyAddressesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.DynamicEmergencyAddresses.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "dynamic-emergency-addresses retrieve", obj, format, transform)
}

func handleDynamicEmergencyAddressesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DynamicEmergencyAddressListParams{}

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
		_, err = client.DynamicEmergencyAddresses.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "dynamic-emergency-addresses list", obj, format, transform)
	} else {
		iter := client.DynamicEmergencyAddresses.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "dynamic-emergency-addresses list", iter, format, transform)
	}
}

func handleDynamicEmergencyAddressesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.DynamicEmergencyAddresses.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "dynamic-emergency-addresses delete", obj, format, transform)
}

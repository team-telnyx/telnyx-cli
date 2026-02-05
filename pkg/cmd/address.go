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

var addressesCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates an address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "business-name",
			Usage:    "The business name associated with the address. An address must have either a first last name or a business name.",
			Required: true,
			BodyPath: "business_name",
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Usage:    "The two-character (ISO 3166-1 alpha-2) country code of the address.",
			Required: true,
			BodyPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			Usage:    "The first name associated with the address. An address must have either a first last name or a business name.",
			Required: true,
			BodyPath: "first_name",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			Usage:    "The last name associated with the address. An address must have either a first last name or a business name.",
			Required: true,
			BodyPath: "last_name",
		},
		&requestflag.Flag[string]{
			Name:     "locality",
			Usage:    "The locality of the address. For US addresses, this corresponds to the city of the address.",
			Required: true,
			BodyPath: "locality",
		},
		&requestflag.Flag[string]{
			Name:     "street-address",
			Usage:    "The primary street address information about the address.",
			Required: true,
			BodyPath: "street_address",
		},
		&requestflag.Flag[bool]{
			Name:     "address-book",
			Usage:    "Indicates whether or not the address should be considered part of your list of addresses that appear for regular use.",
			Default:  true,
			BodyPath: "address_book",
		},
		&requestflag.Flag[string]{
			Name:     "administrative-area",
			Usage:    "The locality of the address. For US addresses, this corresponds to the state of the address.",
			BodyPath: "administrative_area",
		},
		&requestflag.Flag[string]{
			Name:     "borough",
			Usage:    "The borough of the address. This field is not used for addresses in the US but is used for some international addresses.",
			BodyPath: "borough",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "A customer reference string for customer look ups.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "extended-address",
			Usage:    "Additional street address information about the address such as, but not limited to, unit number or apartment number.",
			BodyPath: "extended_address",
		},
		&requestflag.Flag[string]{
			Name:     "neighborhood",
			Usage:    "The neighborhood of the address. This field is not used for addresses in the US but is used for some international addresses.",
			BodyPath: "neighborhood",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "The phone number associated with the address.",
			BodyPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "The postal code of the address.",
			BodyPath: "postal_code",
		},
		&requestflag.Flag[bool]{
			Name:     "validate-address",
			Usage:    "Indicates whether or not the address should be validated for emergency use upon creation or not. This should be left with the default value of `true` unless you have used the `/addresses/actions/validate` endpoint to validate the address separately prior to creation. If an address is not validated for emergency use upon creation and it is not valid, it will not be able to be used for emergency services.",
			Default:  true,
			BodyPath: "validate_address",
		},
	},
	Action:          handleAddressesCreate,
	HideHelpCommand: true,
}

var addressesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAddressesRetrieve,
	HideHelpCommand: true,
}

var addressesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your addresses.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[customer_reference][eq], filter[customer_reference][contains], filter[used_as_emergency], filter[street_address][contains], filter[address_book][eq]",
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
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/>\nThat is: <ul>\n  <li>\n    <code>street_address</code>: sorts the result by the\n    <code>street_address</code> field in ascending order.\n  </li>\n\n  <li>\n    <code>-street_address</code>: sorts the result by the\n    <code>street_address</code> field in descending order.\n  </li>\n</ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.",
			Default:   "created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleAddressesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.address-book",
			InnerField: "address_book",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.customer-reference",
			Usage:      "If present, addresses with <code>customer_reference</code> containing the given value will be returned. Matching is not case-sensitive.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.street-address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.used-as-emergency",
			Usage:      "If set as 'true', only addresses used as the emergency address for at least one active phone-number will be returned. When set to 'false', the opposite happens: only addresses not used as the emergency address from phone-numbers will be returned.",
			InnerField: "used_as_emergency",
		},
	},
})

var addressesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an existing address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAddressesDelete,
	HideHelpCommand: true,
}

func handleAddressesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AddressNewParams{}

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
	_, err = client.Addresses.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addresses create", obj, format, transform)
}

func handleAddressesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Addresses.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addresses retrieve", obj, format, transform)
}

func handleAddressesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AddressListParams{}

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
		_, err = client.Addresses.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "addresses list", obj, format, transform)
	} else {
		iter := client.Addresses.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "addresses list", iter, format, transform)
	}
}

func handleAddressesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Addresses.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addresses delete", obj, format, transform)
}

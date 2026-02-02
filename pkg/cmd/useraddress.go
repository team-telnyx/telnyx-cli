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

var userAddressesCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a user address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "business-name",
			Usage:    "The business name associated with the user address.",
			Required: true,
			BodyPath: "business_name",
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Usage:    "The two-character (ISO 3166-1 alpha-2) country code of the user address.",
			Required: true,
			BodyPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			Usage:    "The first name associated with the user address.",
			Required: true,
			BodyPath: "first_name",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			Usage:    "The last name associated with the user address.",
			Required: true,
			BodyPath: "last_name",
		},
		&requestflag.Flag[string]{
			Name:     "locality",
			Usage:    "The locality of the user address. For US addresses, this corresponds to the city of the address.",
			Required: true,
			BodyPath: "locality",
		},
		&requestflag.Flag[string]{
			Name:     "street-address",
			Usage:    "The primary street address information about the user address.",
			Required: true,
			BodyPath: "street_address",
		},
		&requestflag.Flag[string]{
			Name:     "administrative-area",
			Usage:    "The locality of the user address. For US addresses, this corresponds to the state of the address.",
			BodyPath: "administrative_area",
		},
		&requestflag.Flag[string]{
			Name:     "borough",
			Usage:    "The borough of the user address. This field is not used for addresses in the US but is used for some international addresses.",
			BodyPath: "borough",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "A customer reference string for customer look ups.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "extended-address",
			Usage:    "Additional street address information about the user address such as, but not limited to, unit number or apartment number.",
			BodyPath: "extended_address",
		},
		&requestflag.Flag[string]{
			Name:     "neighborhood",
			Usage:    "The neighborhood of the user address. This field is not used for addresses in the US but is used for some international addresses.",
			BodyPath: "neighborhood",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "The phone number associated with the user address.",
			BodyPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "The postal code of the user address.",
			BodyPath: "postal_code",
		},
		&requestflag.Flag[bool]{
			Name:     "skip-address-verification",
			Usage:    "An optional boolean value specifying if verification of the address should be skipped or not. UserAddresses are generally used for shipping addresses, and failure to validate your shipping address will likely result in a failure to deliver SIM cards or other items ordered from Telnyx. Do not use this parameter unless you are sure that the address is correct even though it cannot be validated. If this is set to any value other than true, verification of the address will be attempted, and the user address will not be allowed if verification fails. If verification fails but suggested values are available that might make the address correct, they will be present in the response as well. If this value is set to true, then the verification will not be attempted. Defaults to false (verification will be performed).",
			Default:  false,
			BodyPath: "skip_address_verification",
		},
	},
	Action:          handleUserAddressesCreate,
	HideHelpCommand: true,
}

var userAddressesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing user address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleUserAddressesRetrieve,
	HideHelpCommand: true,
}

var userAddressesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your user addresses.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[customer_reference][eq], filter[customer_reference][contains], filter[street_address][contains]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/>\nThat is: <ul>\n  <li>\n    <code>street_address</code>: sorts the result by the\n    <code>street_address</code> field in ascending order.\n  </li>\n\n  <li>\n    <code>-street_address</code>: sorts the result by the\n    <code>street_address</code> field in descending order.\n  </li>\n</ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.",
			Default:   "created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleUserAddressesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.customer-reference",
			Usage:      "Filter user addresses via the customer reference. Supports both exact matching (eq) and partial matching (contains). Matching is not case-sensitive.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.street-address",
			Usage:      "Filter user addresses via street address. Supports partial matching (contains). Matching is not case-sensitive.",
			InnerField: "street_address",
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

func handleUserAddressesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.UserAddressNewParams{}

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
	_, err = client.UserAddresses.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "user-addresses create", obj, format, transform)
}

func handleUserAddressesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.UserAddresses.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "user-addresses retrieve", obj, format, transform)
}

func handleUserAddressesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.UserAddressListParams{}

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
		_, err = client.UserAddresses.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "user-addresses list", obj, format, transform)
	} else {
		iter := client.UserAddresses.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "user-addresses list", iter, format, transform)
	}
}

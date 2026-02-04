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

var addressesActionsAcceptSuggestions = cli.Command{
	Name:    "accept-suggestions",
	Usage:   "Accepts this address suggestion as a new emergency address for Operator Connect\nand finishes the uploads of the numbers associated with it to Microsoft.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "address-uuid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The ID of the address.",
			BodyPath: "id",
		},
	},
	Action:          handleAddressesActionsAcceptSuggestions,
	HideHelpCommand: true,
}

var addressesActionsValidate = cli.Command{
	Name:    "validate",
	Usage:   "Validates an address for emergency services.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "country-code",
			Usage:    "The two-character (ISO 3166-1 alpha-2) country code of the address.",
			Required: true,
			BodyPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "The postal code of the address.",
			Required: true,
			BodyPath: "postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "street-address",
			Usage:    "The primary street address information about the address.",
			Required: true,
			BodyPath: "street_address",
		},
		&requestflag.Flag[string]{
			Name:     "administrative-area",
			Usage:    "The locality of the address. For US addresses, this corresponds to the state of the address.",
			BodyPath: "administrative_area",
		},
		&requestflag.Flag[string]{
			Name:     "extended-address",
			Usage:    "Additional street address information about the address such as, but not limited to, unit number or apartment number.",
			BodyPath: "extended_address",
		},
		&requestflag.Flag[string]{
			Name:     "locality",
			Usage:    "The locality of the address. For US addresses, this corresponds to the city of the address.",
			BodyPath: "locality",
		},
	},
	Action:          handleAddressesActionsValidate,
	HideHelpCommand: true,
}

func handleAddressesActionsAcceptSuggestions(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("address-uuid") && len(unusedArgs) > 0 {
		cmd.Set("address-uuid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AddressActionAcceptSuggestionsParams{}

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
	_, err = client.Addresses.Actions.AcceptSuggestions(
		ctx,
		cmd.Value("address-uuid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addresses:actions accept-suggestions", obj, format, transform)
}

func handleAddressesActionsValidate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AddressActionValidateParams{}

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
	_, err = client.Addresses.Actions.Validate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addresses:actions validate", obj, format, transform)
}

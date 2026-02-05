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

var availablePhoneNumberBlocksList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List available phone number blocks",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[locality], filter[country_code], filter[national_destination_code], filter[phone_number_type]",
			QueryPath: "filter",
		},
	},
	Action:          handleAvailablePhoneNumberBlocksList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "Filter phone numbers by country.",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.locality",
			Usage:      "Filter phone numbers by city.",
			InnerField: "locality",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.national-destination-code",
			Usage:      "Filter by the national destination code of the number.",
			InnerField: "national_destination_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number-type",
			Usage:      "Filter phone numbers by number type.",
			InnerField: "phone_number_type",
		},
	},
})

func handleAvailablePhoneNumberBlocksList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AvailablePhoneNumberBlockListParams{}

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
	_, err = client.AvailablePhoneNumberBlocks.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "available-phone-number-blocks list", obj, format, transform)
}

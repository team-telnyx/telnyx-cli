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

var availablePhoneNumbersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List available phone numbers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[phone_number], filter[locality], filter[administrative_area], filter[country_code], filter[national_destination_code], filter[rate_center], filter[phone_number_type], filter[features], filter[limit], filter[best_effort], filter[quickship], filter[reservable], filter[exclude_held_numbers]",
			QueryPath: "filter",
		},
	},
	Action:          handleAvailablePhoneNumbersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.administrative-area",
			Usage:      "Find numbers in a particular US state or CA province.",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.best-effort",
			Usage:      "Filter to determine if best effort results should be included. Only available in USA/CANADA.",
			InnerField: "best_effort",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "Filter phone numbers by country.",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.exclude-held-numbers",
			Usage:      "Filter to exclude phone numbers that are currently on hold/reserved for your account.",
			InnerField: "exclude_held_numbers",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.features",
			Usage:      "Filter phone numbers with specific features.",
			InnerField: "features",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filter.limit",
			Usage:      "Limits the number of results.",
			InnerField: "limit",
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.phone-number",
			Usage:      "Filter phone numbers by pattern matching.",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number-type",
			Usage:      "Filter phone numbers by number type.",
			InnerField: "phone_number_type",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.quickship",
			Usage:      "Filter to exclude phone numbers that need additional time after to purchase to activate. Only applicable for +1 toll_free numbers.",
			InnerField: "quickship",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.rate-center",
			Usage:      "Filter phone numbers by rate center. This filter is only applicable to USA and Canada numbers.",
			InnerField: "rate_center",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.reservable",
			Usage:      "Filter to ensure only numbers that can be reserved are included in the results.",
			InnerField: "reservable",
		},
	},
})

func handleAvailablePhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AvailablePhoneNumberListParams{}

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
	_, err = client.AvailablePhoneNumbers.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "available-phone-numbers list", obj, format, transform)
}

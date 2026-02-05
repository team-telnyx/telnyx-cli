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

var inventoryCoverageList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Creates an inventory coverage request. If locality, npa or\nnational_destination_code is used in groupBy, and no region or locality filters\nare used, the whole paginated set is returned.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[npa], filter[nxx], filter[administrative_area], filter[phone_number_type], filter[country_code], filter[count], filter[features], filter[groupBy]",
			QueryPath: "filter",
		},
	},
	Action:          handleInventoryCoverageList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.administrative-area",
			Usage:      "Filter by administrative area",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.count",
			Usage:      "Include count in the result",
			InnerField: "count",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "Filter by country. Defaults to US",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.features",
			Usage:      "Filter if the phone number should be used for voice, fax, mms, sms, emergency. Returns features in the response when used.",
			InnerField: "features",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.group-by",
			Usage:      "Filter to group results",
			InnerField: "groupBy",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filter.npa",
			Usage:      "Filter by npa",
			InnerField: "npa",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filter.nxx",
			Usage:      "Filter by nxx",
			InnerField: "nxx",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number-type",
			Usage:      "Filter by phone number type",
			InnerField: "phone_number_type",
		},
	},
})

func handleInventoryCoverageList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.InventoryCoverageListParams{}

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
	_, err = client.InventoryCoverage.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "inventory-coverage list", obj, format, transform)
}

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

var mobileNetworkOperatorsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Telnyx has a set of GSM mobile operators partners that are available through our\nmobile network roaming. This resource is entirely managed by Telnyx and may\nchange over time. That means that this resource won't allow any write operations\nfor it. Still, it's available so it can be used as a support resource that can\nbe related to other resources or become a configuration option.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter for mobile network operators (deepObject style). Originally: filter[name][starts_with], filter[name][contains], filter[name][ends_with], filter[country_code], filter[mcc], filter[mnc], filter[tadig], filter[network_preferences_enabled]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated pagination parameter (deepObject style). Originally: page[number], page[size]",
			QueryPath: "page",
		},
	},
	Action:          handleMobileNetworkOperatorsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "Filter by exact country_code.",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.mcc",
			Usage:      "Filter by exact MCC.",
			InnerField: "mcc",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.mnc",
			Usage:      "Filter by exact MNC.",
			InnerField: "mnc",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.name",
			Usage:      "Advanced name filtering operations",
			InnerField: "name",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.network-preferences-enabled",
			Usage:      "Filter by network_preferences_enabled.",
			InnerField: "network_preferences_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.tadig",
			Usage:      "Filter by exact TADIG.",
			InnerField: "tadig",
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

func handleMobileNetworkOperatorsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MobileNetworkOperatorListParams{}

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
		_, err = client.MobileNetworkOperators.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "mobile-network-operators list", obj, format, transform)
	} else {
		iter := client.MobileNetworkOperators.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "mobile-network-operators list", iter, format, transform)
	}
}

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

var networkCoverageList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List all locations and the interfaces that region supports",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[location.region], filter[location.site], filter[location.pop], filter[location.code]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filters",
			Usage:     "Consolidated filters parameter (deepObject style). Originally: filters[available_services][contains]",
			QueryPath: "filters",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleNetworkCoverageList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.location-code",
			Usage:      "The code of associated location to filter on.",
			InnerField: "location.code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.location-pop",
			Usage:      "The POP of associated location to filter on.",
			InnerField: "location.pop",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.location-region",
			Usage:      "The region of associated location to filter on.",
			InnerField: "location.region",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.location-site",
			Usage:      "The site of associated location to filter on.",
			InnerField: "location.site",
		},
	},
	"filters": {
		&requestflag.InnerFlag[any]{
			Name:       "filters.available-services",
			Usage:      "Filter by exact available service match",
			InnerField: "available_services",
		},
	},
})

func handleNetworkCoverageList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NetworkCoverageListParams{}

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
		_, err = client.NetworkCoverage.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "network-coverage list", obj, format, transform)
	} else {
		iter := client.NetworkCoverage.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "network-coverage list", iter, format, transform)
	}
}

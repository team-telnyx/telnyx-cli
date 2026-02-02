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

var virtualCrossConnectsCoverageList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Virtual Cross Connects Cloud Coverage.<br /><br />This endpoint shows which\ncloud regions are available for the `location_code` your Virtual Cross Connect\nwill be provisioned in.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[cloud_provider], filter[cloud_provider_region], filter[location.region], filter[location.site], filter[location.pop], filter[location.code]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filters",
			Usage:     "Consolidated filters parameter (deepObject style). Originally: filters[available_bandwidth][contains]",
			QueryPath: "filters",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[number], page[size]",
			QueryPath: "page",
		},
	},
	Action:          handleVirtualCrossConnectsCoverageList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.cloud-provider",
			Usage:      "The Virtual Private Cloud provider.",
			InnerField: "cloud_provider",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.cloud-provider-region",
			Usage:      "The region of specific cloud provider.",
			InnerField: "cloud_provider_region",
		},
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
			Name:       "filters.available-bandwidth",
			Usage:      "Filter by exact available bandwidth match",
			InnerField: "available_bandwidth",
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

func handleVirtualCrossConnectsCoverageList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VirtualCrossConnectsCoverageListParams{}

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
		_, err = client.VirtualCrossConnectsCoverage.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "virtual-cross-connects-coverage list", obj, format, transform)
	} else {
		iter := client.VirtualCrossConnectsCoverage.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "virtual-cross-connects-coverage list", iter, format, transform)
	}
}

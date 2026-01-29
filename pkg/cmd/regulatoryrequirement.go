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

var regulatoryRequirementsRetrieve = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve regulatory requirements",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[phone_number], filter[requirement_group_id], filter[country_code], filter[phone_number_type], filter[action]",
			QueryPath: "filter",
		},
	},
	Action:          handleRegulatoryRequirementsRetrieve,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.action",
			Usage:      "Action to check requirements for",
			InnerField: "action",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.country-code",
			Usage:      "Country code(iso2) to check requirements for",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number",
			Usage:      "Phone number to check requirements for",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number-type",
			Usage:      "Phone number type to check requirements for",
			InnerField: "phone_number_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.requirement-group-id",
			Usage:      "ID of requirement group to check requirements for",
			InnerField: "requirement_group_id",
		},
	},
})

func handleRegulatoryRequirementsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RegulatoryRequirementGetParams{}

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
	_, err = client.RegulatoryRequirements.Get(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "regulatory-requirements retrieve", obj, format, transform)
}

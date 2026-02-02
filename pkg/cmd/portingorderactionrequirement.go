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

var portingOrdersActionRequirementsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of action requirements for a specific porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[id][in][], filter[requirement_type_id], filter[action_type], filter[status]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "sort",
			Usage:     "Consolidated sort parameter (deepObject style). Originally: sort[value]",
			QueryPath: "sort",
		},
	},
	Action:          handlePortingOrdersActionRequirementsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.id",
			Usage:      "Filter action requirements by a list of IDs",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.action-type",
			Usage:      "Filter action requirements by action type",
			InnerField: "action_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.requirement-type-id",
			Usage:      "Filter action requirements by requirement type ID",
			InnerField: "requirement_type_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter action requirements by status",
			InnerField: "status",
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
	"sort": {
		&requestflag.InnerFlag[string]{
			Name:       "sort.value",
			Usage:      "Specifies the sort order for results. If not given, results are sorted by created_at in descending order.",
			InnerField: "value",
		},
	},
})

var portingOrdersActionRequirementsInitiate = requestflag.WithInnerFlags(cli.Command{
	Name:    "initiate",
	Usage:   "Initiates a specific action requirement for a porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "porting-order-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Usage:    "Required information for initiating the action requirement for AU ID verification.",
			Required: true,
			BodyPath: "params",
		},
	},
	Action:          handlePortingOrdersActionRequirementsInitiate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"params": {
		&requestflag.InnerFlag[string]{
			Name:       "params.first-name",
			Usage:      "The first name of the person that will perform the verification flow.",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "params.last-name",
			Usage:      "The last name of the person that will perform the verification flow.",
			InnerField: "last_name",
		},
	},
})

func handlePortingOrdersActionRequirementsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("porting-order-id") && len(unusedArgs) > 0 {
		cmd.Set("porting-order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderActionRequirementListParams{}

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
		_, err = client.PortingOrders.ActionRequirements.List(
			ctx,
			cmd.Value("porting-order-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "porting-orders:action-requirements list", obj, format, transform)
	} else {
		iter := client.PortingOrders.ActionRequirements.ListAutoPaging(
			ctx,
			cmd.Value("porting-order-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "porting-orders:action-requirements list", iter, format, transform)
	}
}

func handlePortingOrdersActionRequirementsInitiate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderActionRequirementInitiateParams{
		PortingOrderID: cmd.Value("porting-order-id").(string),
	}

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
	_, err = client.PortingOrders.ActionRequirements.Initiate(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "porting-orders:action-requirements initiate", obj, format, transform)
}

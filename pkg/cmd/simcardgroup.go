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

var simCardGroupsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new SIM card group object",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A user friendly name for the SIM card group.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data-limit",
			Usage:    "Upper limit on the amount of data the SIM cards, within the group, can use.",
			BodyPath: "data_limit",
		},
	},
	Action:          handleSimCardGroupsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data-limit": {
		&requestflag.InnerFlag[string]{
			Name:       "data-limit.amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-limit.unit",
			InnerField: "unit",
		},
	},
})

var simCardGroupsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the details regarding a specific SIM card group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "include-iccids",
			Usage:     "It includes a list of associated ICCIDs.",
			Default:   false,
			QueryPath: "include_iccids",
		},
	},
	Action:          handleSimCardGroupsRetrieve,
	HideHelpCommand: true,
}

var simCardGroupsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates a SIM card group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data-limit",
			Usage:    "Upper limit on the amount of data the SIM cards, within the group, can use.",
			BodyPath: "data_limit",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A user friendly name for the SIM card group.",
			BodyPath: "name",
		},
	},
	Action:          handleSimCardGroupsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data-limit": {
		&requestflag.InnerFlag[string]{
			Name:       "data-limit.amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-limit.unit",
			InnerField: "unit",
		},
	},
})

var simCardGroupsList = cli.Command{
	Name:    "list",
	Usage:   "Get all SIM card groups belonging to the user that match the given filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "A valid SIM card group name.",
			QueryPath: "filter[name]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-private-wireless-gateway-id",
			Usage:     "A Private Wireless Gateway ID associated with the group.",
			QueryPath: "filter[private_wireless_gateway_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-wireless-blocklist-id",
			Usage:     "A Wireless Blocklist ID associated with the group.",
			QueryPath: "filter[wireless_blocklist_id]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page.",
			Default:   20,
			QueryPath: "page[size]",
		},
	},
	Action:          handleSimCardGroupsList,
	HideHelpCommand: true,
}

var simCardGroupsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a SIM card group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardGroupsDelete,
	HideHelpCommand: true,
}

func handleSimCardGroupsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardGroupNewParams{}

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
	_, err = client.SimCardGroups.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-groups create", obj, format, transform)
}

func handleSimCardGroupsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardGroupGetParams{}

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
	_, err = client.SimCardGroups.Get(
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
	return ShowJSON(os.Stdout, "sim-card-groups retrieve", obj, format, transform)
}

func handleSimCardGroupsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardGroupUpdateParams{}

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
	_, err = client.SimCardGroups.Update(
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
	return ShowJSON(os.Stdout, "sim-card-groups update", obj, format, transform)
}

func handleSimCardGroupsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardGroupListParams{}

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
		_, err = client.SimCardGroups.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "sim-card-groups list", obj, format, transform)
	} else {
		iter := client.SimCardGroups.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "sim-card-groups list", iter, format, transform)
	}
}

func handleSimCardGroupsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCardGroups.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-groups delete", obj, format, transform)
}

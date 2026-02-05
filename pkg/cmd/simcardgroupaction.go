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

var simCardGroupsActionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "This API allows fetching detailed information about a SIM card group action\nresource to make follow-ups in an existing asynchronous operation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardGroupsActionsRetrieve,
	HideHelpCommand: true,
}

var simCardGroupsActionsList = cli.Command{
	Name:    "list",
	Usage:   "This API allows listing a paginated collection a SIM card group actions. It\nallows to explore a collection of existing asynchronous operation using specific\nfilters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-sim-card-group-id",
			Usage:     "A valid SIM card group ID.",
			QueryPath: "filter[sim_card_group_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     "Filter by a specific status of the resource's lifecycle.",
			QueryPath: "filter[status]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-type",
			Usage:     "Filter by action type.",
			QueryPath: "filter[type]",
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
	Action:          handleSimCardGroupsActionsList,
	HideHelpCommand: true,
}

var simCardGroupsActionsRemovePrivateWirelessGateway = cli.Command{
	Name:    "remove-private-wireless-gateway",
	Usage:   "This action will asynchronously remove an existing Private Wireless Gateway\ndefinition from a SIM card group. Completing this operation defines that all SIM\ncards in the SIM card group will get their traffic handled by Telnyx's default\nmobile network configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardGroupsActionsRemovePrivateWirelessGateway,
	HideHelpCommand: true,
}

var simCardGroupsActionsRemoveWirelessBlocklist = cli.Command{
	Name:    "remove-wireless-blocklist",
	Usage:   "This action will asynchronously remove an existing Wireless Blocklist to all the\nSIMs in the SIM card group.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardGroupsActionsRemoveWirelessBlocklist,
	HideHelpCommand: true,
}

var simCardGroupsActionsSetPrivateWirelessGateway = cli.Command{
	Name:    "set-private-wireless-gateway",
	Usage:   "This action will asynchronously assign a provisioned Private Wireless Gateway to\nthe SIM card group. Completing this operation defines that all SIM cards in the\nSIM card group will get their traffic controlled by the associated Private\nWireless Gateway. This operation will also imply that new SIM cards assigned to\na group will inherit its network definitions. If it's moved to a different group\nthat doesn't have a Private Wireless Gateway, it'll use Telnyx's default mobile\nnetwork configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "private-wireless-gateway-id",
			Usage:    "The identification of the related Private Wireless Gateway resource.",
			Required: true,
			BodyPath: "private_wireless_gateway_id",
		},
	},
	Action:          handleSimCardGroupsActionsSetPrivateWirelessGateway,
	HideHelpCommand: true,
}

var simCardGroupsActionsSetWirelessBlocklist = cli.Command{
	Name:    "set-wireless-blocklist",
	Usage:   "This action will asynchronously assign a Wireless Blocklist to all the SIMs in\nthe SIM card group.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "wireless-blocklist-id",
			Usage:    "The identification of the related Wireless Blocklist resource.",
			Required: true,
			BodyPath: "wireless_blocklist_id",
		},
	},
	Action:          handleSimCardGroupsActionsSetWirelessBlocklist,
	HideHelpCommand: true,
}

func handleSimCardGroupsActionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCardGroups.Actions.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-groups:actions retrieve", obj, format, transform)
}

func handleSimCardGroupsActionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardGroupActionListParams{}

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
		_, err = client.SimCardGroups.Actions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "sim-card-groups:actions list", obj, format, transform)
	} else {
		iter := client.SimCardGroups.Actions.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "sim-card-groups:actions list", iter, format, transform)
	}
}

func handleSimCardGroupsActionsRemovePrivateWirelessGateway(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCardGroups.Actions.RemovePrivateWirelessGateway(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-groups:actions remove-private-wireless-gateway", obj, format, transform)
}

func handleSimCardGroupsActionsRemoveWirelessBlocklist(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCardGroups.Actions.RemoveWirelessBlocklist(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-groups:actions remove-wireless-blocklist", obj, format, transform)
}

func handleSimCardGroupsActionsSetPrivateWirelessGateway(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardGroupActionSetPrivateWirelessGatewayParams{}

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
	_, err = client.SimCardGroups.Actions.SetPrivateWirelessGateway(
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
	return ShowJSON(os.Stdout, "sim-card-groups:actions set-private-wireless-gateway", obj, format, transform)
}

func handleSimCardGroupsActionsSetWirelessBlocklist(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardGroupActionSetWirelessBlocklistParams{}

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
	_, err = client.SimCardGroups.Actions.SetWirelessBlocklist(
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
	return ShowJSON(os.Stdout, "sim-card-groups:actions set-wireless-blocklist", obj, format, transform)
}

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

var simCardsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the details regarding a specific SIM card.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "include-pin-puk-codes",
			Usage:     "When set to true, includes the PIN and PUK codes in the response. These codes are used for SIM card security and unlocking purposes. Available for both physical SIM cards and eSIMs.",
			Default:   false,
			QueryPath: "include_pin_puk_codes",
		},
		&requestflag.Flag[bool]{
			Name:      "include-sim-card-group",
			Usage:     "It includes the associated SIM card group object in the response when present.",
			Default:   false,
			QueryPath: "include_sim_card_group",
		},
	},
	Action:          handleSimCardsRetrieve,
	HideHelpCommand: true,
}

var simCardsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates SIM card data",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sim-card-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "authorized-imei",
			Usage:    "List of IMEIs authorized to use a given SIM card.",
			BodyPath: "authorized_imeis",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data-limit",
			Usage:    "The SIM card individual data limit configuration.",
			BodyPath: "data_limit",
		},
		&requestflag.Flag[string]{
			Name:     "sim-card-group-id",
			Usage:    "The group SIMCardGroup identification. This attribute can be <code>null</code> when it's present in an associated resource.",
			BodyPath: "sim_card_group_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "status",
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Searchable tags associated with the SIM card",
			BodyPath: "tags",
		},
	},
	Action:          handleSimCardsUpdate,
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
	"status": {
		&requestflag.InnerFlag[string]{
			Name:       "status.reason",
			Usage:      "It describes why the SIM card is in the current status.",
			InnerField: "reason",
		},
		&requestflag.InnerFlag[string]{
			Name:       "status.value",
			Usage:      "The current status of the SIM card. It will be one of the following: <br/>\n<ul>\n <li><code>registering</code> - the card is being registered</li>\n <li><code>enabling</code> - the card is being enabled</li>\n <li><code>enabled</code> - the card is enabled and ready for use</li>\n <li><code>disabling</code> - the card is being disabled</li>\n <li><code>disabled</code> - the card has been disabled and cannot be used</li>\n <li><code>data_limit_exceeded</code> - the card has exceeded its data consumption limit</li>\n <li><code>setting_standby</code> - the process to set the card in stand by is in progress</li>\n <li><code>standby</code> - the card is in stand by</li>\n</ul>\nTransitioning between the enabled and disabled states may take a period of time.\n",
			InnerField: "value",
		},
	},
})

var simCardsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get all SIM cards belonging to the user that match the given filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter for SIM cards (deepObject style). Originally: filter[iccid], filter[msisdn], filter[status], filter[tags]",
			QueryPath: "filter",
		},
		&requestflag.Flag[string]{
			Name:      "filter-sim-card-group-id",
			Usage:     "A valid SIM card group ID.",
			QueryPath: "filter[sim_card_group_id]",
		},
		&requestflag.Flag[bool]{
			Name:      "include-sim-card-group",
			Usage:     "It includes the associated SIM card group object in the response when present.",
			Default:   false,
			QueryPath: "include_sim_card_group",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated pagination parameter (deepObject style). Originally: page[number], page[size]",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sorts SIM cards by the given field. Defaults to ascending order unless field is prefixed with a minus sign.",
			QueryPath: "sort",
		},
	},
	Action:          handleSimCardsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.iccid",
			Usage:      "A search string to partially match for the SIM card's ICCID.",
			InnerField: "iccid",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.msisdn",
			Usage:      "A search string to match for the SIM card's MSISDN.",
			InnerField: "msisdn",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.status",
			Usage:      "Filter by a SIM card's status.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.tags",
			Usage:      "A list of SIM card tags to filter on.<br/><br/>\n If the SIM card contains <b><i>all</i></b> of the given <code>tags</code> they will be found.<br/><br/>\nFor example, if the SIM cards have the following tags: <ul>\n  <li><code>['customers', 'staff', 'test']</code>\n  <li><code>['test']</code></li>\n  <li><code>['customers']</code></li>\n</ul>\nSearching for <code>['customers', 'test']</code> returns only the first because it's the only one with both tags.<br/> Searching for <code>test</code> returns the first two SIMs, because both of them have such tag.<br/> Searching for <code>customers</code> returns the first and last SIMs.<br/>\n",
			InnerField: "tags",
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

var simCardsDelete = cli.Command{
	Name:    "delete",
	Usage:   "The SIM card will be decommissioned, removed from your account and you will stop\nbeing charged.<br />The SIM card won't be able to connect to the network after\nthe deletion is completed, thus making it impossible to consume data.<br/>\nTransitioning to the disabled state may take a period of time. Until the\ntransition is completed, the SIM card status will be disabling\n<code>disabling</code>.<br />In order to re-enable the SIM card, you will need\nto re-register it.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "report-lost",
			Usage:     "Enables deletion of disabled eSIMs that can't be uninstalled from a device. This is irreversible and the eSIM cannot be re-registered.",
			Default:   false,
			QueryPath: "report_lost",
		},
	},
	Action:          handleSimCardsDelete,
	HideHelpCommand: true,
}

var simCardsGetActivationCode = cli.Command{
	Name:    "get-activation-code",
	Usage:   "It returns the activation code for an eSIM.<br/><br/> This API is only available\nfor eSIMs. If the given SIM is a physical SIM card, or has already been\ninstalled, an error will be returned.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardsGetActivationCode,
	HideHelpCommand: true,
}

var simCardsGetDeviceDetails = cli.Command{
	Name:    "get-device-details",
	Usage:   "It returns the device details where a SIM card is currently being used.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardsGetDeviceDetails,
	HideHelpCommand: true,
}

var simCardsGetPublicIP = cli.Command{
	Name:    "get-public-ip",
	Usage:   "It returns the public IP requested for a SIM card.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardsGetPublicIP,
	HideHelpCommand: true,
}

var simCardsListWirelessConnectivityLogs = cli.Command{
	Name:    "list-wireless-connectivity-logs",
	Usage:   "This API allows listing a paginated collection of Wireless Connectivity Logs\nassociated with a SIM Card, for troubleshooting purposes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
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
	Action:          handleSimCardsListWirelessConnectivityLogs,
	HideHelpCommand: true,
}

func handleSimCardsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardGetParams{}

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
	_, err = client.SimCards.Get(
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
	return ShowJSON(os.Stdout, "sim-cards retrieve", obj, format, transform)
}

func handleSimCardsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sim-card-id") && len(unusedArgs) > 0 {
		cmd.Set("sim-card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardUpdateParams{}

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
	_, err = client.SimCards.Update(
		ctx,
		cmd.Value("sim-card-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards update", obj, format, transform)
}

func handleSimCardsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardListParams{}

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
		_, err = client.SimCards.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "sim-cards list", obj, format, transform)
	} else {
		iter := client.SimCards.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "sim-cards list", iter, format, transform)
	}
}

func handleSimCardsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardDeleteParams{}

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
	_, err = client.SimCards.Delete(
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
	return ShowJSON(os.Stdout, "sim-cards delete", obj, format, transform)
}

func handleSimCardsGetActivationCode(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCards.GetActivationCode(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards get-activation-code", obj, format, transform)
}

func handleSimCardsGetDeviceDetails(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCards.GetDeviceDetails(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards get-device-details", obj, format, transform)
}

func handleSimCardsGetPublicIP(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCards.GetPublicIP(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards get-public-ip", obj, format, transform)
}

func handleSimCardsListWirelessConnectivityLogs(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardListWirelessConnectivityLogsParams{}

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
		_, err = client.SimCards.ListWirelessConnectivityLogs(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "sim-cards list-wireless-connectivity-logs", obj, format, transform)
	} else {
		iter := client.SimCards.ListWirelessConnectivityLogsAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "sim-cards list-wireless-connectivity-logs", iter, format, transform)
	}
}

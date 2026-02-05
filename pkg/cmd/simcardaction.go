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

var simCardsActionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "This API fetches detailed information about a SIM card action to follow-up on an\nexisting asynchronous operation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardsActionsRetrieve,
	HideHelpCommand: true,
}

var simCardsActionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "This API lists a paginated collection of SIM card actions. It enables exploring\na collection of existing asynchronous operations using specific filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter for SIM card actions (deepObject style). Originally: filter[sim_card_id], filter[status], filter[bulk_sim_card_action_id], filter[action_type]",
			QueryPath: "filter",
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
	Action:          handleSimCardsActionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.action-type",
			Usage:      "Filter by action type.",
			InnerField: "action_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.bulk-sim-card-action-id",
			Usage:      "Filter by a bulk SIM card action ID.",
			InnerField: "bulk_sim_card_action_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.sim-card-id",
			Usage:      "A valid SIM card ID.",
			InnerField: "sim_card_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter by a specific status of the resource's lifecycle.",
			InnerField: "status",
		},
	},
})

var simCardsActionsBulkSetPublicIPs = cli.Command{
	Name:    "bulk-set-public-ips",
	Usage:   "This API triggers an asynchronous operation to set a public IP for each of the\nspecified SIM cards.<br/> For each SIM Card a SIM Card Action will be generated.\nThe status of the SIM Card Action can be followed through the\n[List SIM Card Action](https://developers.telnyx.com/api-reference/sim-card-actions/list-sim-card-actions)\nAPI.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "sim-card-id",
			Required: true,
			BodyPath: "sim_card_ids",
		},
	},
	Action:          handleSimCardsActionsBulkSetPublicIPs,
	HideHelpCommand: true,
}

var simCardsActionsDisable = cli.Command{
	Name:    "disable",
	Usage:   "This API disables a SIM card, disconnecting it from the network and making it\nimpossible to consume data.<br/> The API will trigger an asynchronous operation\ncalled a SIM Card Action. Transitioning to the disabled state may take a period\nof time. The status of the SIM Card Action can be followed through the\n[List SIM Card Action](https://developers.telnyx.com/api-reference/sim-card-actions/list-sim-card-actions)\nAPI.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardsActionsDisable,
	HideHelpCommand: true,
}

var simCardsActionsEnable = cli.Command{
	Name:    "enable",
	Usage:   "This API enables a SIM card, connecting it to the network and making it possible\nto consume data.<br/> To enable a SIM card, it must be associated with a SIM\ncard group.<br/> The API will trigger an asynchronous operation called a SIM\nCard Action. Transitioning to the enabled state may take a period of time. The\nstatus of the SIM Card Action can be followed through the\n[List SIM Card Action](https://developers.telnyx.com/api-reference/sim-card-actions/list-sim-card-actions)\nAPI.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardsActionsEnable,
	HideHelpCommand: true,
}

var simCardsActionsRemovePublicIP = cli.Command{
	Name:    "remove-public-ip",
	Usage:   "This API removes an existing public IP from a SIM card. <br/><br/> The API will\ntrigger an asynchronous operation called a SIM Card Action. The status of the\nSIM Card Action can be followed through the\n[List SIM Card Action](https://developers.telnyx.com/api-reference/sim-card-actions/list-sim-card-actions)\nAPI.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardsActionsRemovePublicIP,
	HideHelpCommand: true,
}

var simCardsActionsSetPublicIP = cli.Command{
	Name:    "set-public-ip",
	Usage:   "This API makes a SIM card reachable on the public internet by mapping a random\npublic IP to the SIM card. <br/><br/> The API will trigger an asynchronous\noperation called a SIM Card Action. The status of the SIM Card Action can be\nfollowed through the\n[List SIM Card Action](https://developers.telnyx.com/api-reference/sim-card-actions/list-sim-card-actions)\nAPI. <br/><br/> Setting a Public IP to a SIM Card incurs a charge and will only\nsucceed if the account has sufficient funds.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "region-code",
			Usage:     "The code of the region where the public IP should be assigned. A list of available regions can be found at the regions endpoint",
			QueryPath: "region_code",
		},
	},
	Action:          handleSimCardsActionsSetPublicIP,
	HideHelpCommand: true,
}

var simCardsActionsSetStandby = cli.Command{
	Name:    "set-standby",
	Usage:   "The SIM card will be able to connect to the network once the process to set it\nto standby has been completed, thus making it possible to consume data.<br/> To\nset a SIM card to standby, it must be associated with SIM card group.<br/> The\nAPI will trigger an asynchronous operation called a SIM Card Action.\nTransitioning to the standby state may take a period of time. The status of the\nSIM Card Action can be followed through the\n[List SIM Card Action](https://developers.telnyx.com/api-reference/sim-card-actions/list-sim-card-actions)\nAPI.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardsActionsSetStandby,
	HideHelpCommand: true,
}

var simCardsActionsValidateRegistrationCodes = cli.Command{
	Name:    "validate-registration-codes",
	Usage:   "It validates whether SIM card registration codes are valid or not.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "registration-code",
			BodyPath: "registration_codes",
		},
	},
	Action:          handleSimCardsActionsValidateRegistrationCodes,
	HideHelpCommand: true,
}

func handleSimCardsActionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCards.Actions.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards:actions retrieve", obj, format, transform)
}

func handleSimCardsActionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardActionListParams{}

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
		_, err = client.SimCards.Actions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "sim-cards:actions list", obj, format, transform)
	} else {
		iter := client.SimCards.Actions.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "sim-cards:actions list", iter, format, transform)
	}
}

func handleSimCardsActionsBulkSetPublicIPs(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardActionBulkSetPublicIPsParams{}

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
	_, err = client.SimCards.Actions.BulkSetPublicIPs(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards:actions bulk-set-public-ips", obj, format, transform)
}

func handleSimCardsActionsDisable(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCards.Actions.Disable(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards:actions disable", obj, format, transform)
}

func handleSimCardsActionsEnable(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCards.Actions.Enable(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards:actions enable", obj, format, transform)
}

func handleSimCardsActionsRemovePublicIP(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCards.Actions.RemovePublicIP(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards:actions remove-public-ip", obj, format, transform)
}

func handleSimCardsActionsSetPublicIP(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardActionSetPublicIPParams{}

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
	_, err = client.SimCards.Actions.SetPublicIP(
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
	return ShowJSON(os.Stdout, "sim-cards:actions set-public-ip", obj, format, transform)
}

func handleSimCardsActionsSetStandby(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCards.Actions.SetStandby(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards:actions set-standby", obj, format, transform)
}

func handleSimCardsActionsValidateRegistrationCodes(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardActionValidateRegistrationCodesParams{}

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
	_, err = client.SimCards.Actions.ValidateRegistrationCodes(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-cards:actions validate-registration-codes", obj, format, transform)
}

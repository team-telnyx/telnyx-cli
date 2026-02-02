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

var paymentAutoRechargePrefsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update payment auto recharge preferences.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Usage:    "Whether auto recharge is enabled.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "invoice-enabled",
			BodyPath: "invoice_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "preference",
			Usage:    "The payment preference for auto recharge.",
			BodyPath: "preference",
		},
		&requestflag.Flag[string]{
			Name:     "recharge-amount",
			Usage:    "The amount to recharge the account, the actual recharge amount will be the amount necessary to reach the threshold amount plus the recharge amount.",
			BodyPath: "recharge_amount",
		},
		&requestflag.Flag[string]{
			Name:     "threshold-amount",
			Usage:    "The threshold amount at which the account will be recharged.",
			BodyPath: "threshold_amount",
		},
	},
	Action:          handlePaymentAutoRechargePrefsUpdate,
	HideHelpCommand: true,
}

var paymentAutoRechargePrefsList = cli.Command{
	Name:            "list",
	Usage:           "Returns the payment auto recharge preferences.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePaymentAutoRechargePrefsList,
	HideHelpCommand: true,
}

func handlePaymentAutoRechargePrefsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PaymentAutoRechargePrefUpdateParams{}

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
	_, err = client.Payment.AutoRechargePrefs.Update(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payment:auto-recharge-prefs update", obj, format, transform)
}

func handlePaymentAutoRechargePrefsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.Payment.AutoRechargePrefs.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payment:auto-recharge-prefs list", obj, format, transform)
}

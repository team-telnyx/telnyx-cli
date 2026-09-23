// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var machinePaymentsAccountCredit = cli.Command{
	Name:    "account-credit",
	Usage:   "Creates an account credit using the Machine Payment Protocol (MPP), an HTTP-402\npayment flow for machines and agents.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "amount-usd",
			Usage:    "Amount to credit in USD, as a decimal string with up to two fractional digits (by default between 5.00 and 500.00). The request body is required on the initial challenge request and remains required on a paid retry, where you re-send the identical body plus the payment credential — the credential, not the body, selects the payment, and the retried body is not re-validated.",
			Required: true,
			BodyPath: "amount_usd",
		},
	},
	Action:          handleMachinePaymentsAccountCredit,
	HideHelpCommand: true,
}

func handleMachinePaymentsAccountCredit(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := telnyx.MachinePaymentAccountCreditParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MachinePayments.AccountCredit(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "machine-payments account-credit",
		Transform:      transform,
	})
}

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

var x402CreditAccountCreatePaymentQuote = cli.Command{
	Name:    "create-payment-quote",
	Usage:   "Creates a payment quote for the specified USD amount. Returns payment details\nincluding the x402 payment requirements, network, and expiration time. The quote\nmust be settled before it expires.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "amount-usd",
			Usage:    "Amount in USD to fund (minimum 5.00, maximum 10000.00).",
			Required: true,
			BodyPath: "amount_usd",
		},
	},
	Action:          handleX402CreditAccountCreatePaymentQuote,
	HideHelpCommand: true,
}

var x402CreditAccountSettlePayment = cli.Command{
	Name:    "settle-payment",
	Usage:   "Settles an x402 payment using the quote ID and a signed payment authorization.\nThe payment signature can be provided via the `PAYMENT-SIGNATURE` header or the\n`payment_signature` body parameter. Settlement is idempotent — submitting the\nsame quote ID multiple times returns the existing transaction.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The quote ID to settle.",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "payment-signature",
			Usage:    "Base64-encoded signed payment authorization (x402 PaymentPayload). Can alternatively be provided via the PAYMENT-SIGNATURE header.",
			BodyPath: "payment_signature",
		},
		&requestflag.Flag[string]{
			Name:       "payment-signature",
			HeaderPath: "PAYMENT-SIGNATURE",
		},
	},
	Action:          handleX402CreditAccountSettlePayment,
	HideHelpCommand: true,
}

func handleX402CreditAccountCreatePaymentQuote(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.X402CreditAccountNewPaymentQuoteParams{}

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
	_, err = client.X402.CreditAccount.NewPaymentQuote(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "x402:credit-account create-payment-quote", obj, format, transform)
}

func handleX402CreditAccountSettlePayment(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.X402CreditAccountSettlePaymentParams{}

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
	_, err = client.X402.CreditAccount.SettlePayment(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "x402:credit-account settle-payment", obj, format, transform)
}

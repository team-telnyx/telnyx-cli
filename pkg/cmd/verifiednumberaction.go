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

var verifiedNumbersActionsSubmitVerificationCode = cli.Command{
	Name:    "submit-verification-code",
	Usage:   "Submit verification code",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "+E164 formatted phone number.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "verification-code",
			Required: true,
			BodyPath: "verification_code",
		},
	},
	Action:          handleVerifiedNumbersActionsSubmitVerificationCode,
	HideHelpCommand: true,
}

func handleVerifiedNumbersActionsSubmitVerificationCode(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VerifiedNumberActionSubmitVerificationCodeParams{}

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
	_, err = client.VerifiedNumbers.Actions.SubmitVerificationCode(
		ctx,
		cmd.Value("phone-number").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "verified-numbers:actions submit-verification-code", obj, format, transform)
}

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

var verificationsActionsVerify = cli.Command{
	Name:    "verify",
	Usage:   "Verify verification code by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "verification-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "code",
			Usage:    "This is the code the user submits for verification.",
			BodyPath: "code",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Identifies if the verification code has been accepted or rejected. Only permitted if custom_code was used for the verification.",
			BodyPath: "status",
		},
	},
	Action:          handleVerificationsActionsVerify,
	HideHelpCommand: true,
}

func handleVerificationsActionsVerify(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("verification-id") && len(unusedArgs) > 0 {
		cmd.Set("verification-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VerificationActionVerifyParams{}

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
	_, err = client.Verifications.Actions.Verify(
		ctx,
		cmd.Value("verification-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "verifications:actions verify", obj, format, transform)
}

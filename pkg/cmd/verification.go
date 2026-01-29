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

var verificationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve verification",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "verification-id",
			Required: true,
		},
	},
	Action:          handleVerificationsRetrieve,
	HideHelpCommand: true,
}

var verificationsTriggerCall = cli.Command{
	Name:    "trigger-call",
	Usage:   "Trigger Call verification",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "+E164 formatted phone number.",
			Required: true,
			BodyPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:     "verify-profile-id",
			Usage:    "The identifier of the associated Verify profile.",
			Required: true,
			BodyPath: "verify_profile_id",
		},
		&requestflag.Flag[any]{
			Name:     "custom-code",
			Usage:    "Send a self-generated numeric code to the end-user",
			Default:  nil,
			BodyPath: "custom_code",
		},
		&requestflag.Flag[any]{
			Name:     "extension",
			Usage:    "Optional extension to dial after call is answered using DTMF digits. Valid digits are 0-9, A-D, *, and #. Pauses can be added using w (0.5s) and W (1s).",
			Default:  nil,
			BodyPath: "extension",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-secs",
			Usage:    "The number of seconds the verification code is valid for.",
			BodyPath: "timeout_secs",
		},
	},
	Action:          handleVerificationsTriggerCall,
	HideHelpCommand: true,
}

var verificationsTriggerFlashcall = cli.Command{
	Name:    "trigger-flashcall",
	Usage:   "Trigger Flash call verification",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "+E164 formatted phone number.",
			Required: true,
			BodyPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:     "verify-profile-id",
			Usage:    "The identifier of the associated Verify profile.",
			Required: true,
			BodyPath: "verify_profile_id",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-secs",
			Usage:    "The number of seconds the verification code is valid for.",
			BodyPath: "timeout_secs",
		},
	},
	Action:          handleVerificationsTriggerFlashcall,
	HideHelpCommand: true,
}

var verificationsTriggerSMS = cli.Command{
	Name:    "trigger-sms",
	Usage:   "Trigger SMS verification",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "+E164 formatted phone number.",
			Required: true,
			BodyPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:     "verify-profile-id",
			Usage:    "The identifier of the associated Verify profile.",
			Required: true,
			BodyPath: "verify_profile_id",
		},
		&requestflag.Flag[any]{
			Name:     "custom-code",
			Usage:    "Send a self-generated numeric code to the end-user",
			Default:  nil,
			BodyPath: "custom_code",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-secs",
			Usage:    "The number of seconds the verification code is valid for.",
			BodyPath: "timeout_secs",
		},
	},
	Action:          handleVerificationsTriggerSMS,
	HideHelpCommand: true,
}

func handleVerificationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("verification-id") && len(unusedArgs) > 0 {
		cmd.Set("verification-id", unusedArgs[0])
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
	_, err = client.Verifications.Get(ctx, cmd.Value("verification-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "verifications retrieve", obj, format, transform)
}

func handleVerificationsTriggerCall(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VerificationTriggerCallParams{}

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
	_, err = client.Verifications.TriggerCall(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "verifications trigger-call", obj, format, transform)
}

func handleVerificationsTriggerFlashcall(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VerificationTriggerFlashcallParams{}

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
	_, err = client.Verifications.TriggerFlashcall(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "verifications trigger-flashcall", obj, format, transform)
}

func handleVerificationsTriggerSMS(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VerificationTriggerSMSParams{}

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
	_, err = client.Verifications.TriggerSMS(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "verifications trigger-sms", obj, format, transform)
}

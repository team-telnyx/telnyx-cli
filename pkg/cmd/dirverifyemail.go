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

var dirVerifyEmailConfirmCode = cli.Command{
	Name:    "confirm-code",
	Usage:   "Submit the 6-digit code that was emailed to the DIR's authorizer email. On\nsuccess the authorizer email is marked verified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[string]{
			Name:     "code",
			Usage:    "The 6-digit code sent to the authorizer email.",
			Required: true,
			BodyPath: "code",
		},
	},
	Action:          handleDirVerifyEmailConfirmCode,
	HideHelpCommand: true,
}

var dirVerifyEmailSendCode = cli.Command{
	Name:    "send-code",
	Usage:   "Email a 6-digit code to the DIR's authorizer email to confirm ownership of that\naddress.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
	},
	Action:          handleDirVerifyEmailSendCode,
	HideHelpCommand: true,
}

var dirVerifyEmailStatus = cli.Command{
	Name:    "status",
	Usage:   "Whether the DIR's current authorizer email has been verified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
	},
	Action:          handleDirVerifyEmailStatus,
	HideHelpCommand: true,
}

func handleDirVerifyEmailConfirmCode(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dir-id") && len(unusedArgs) > 0 {
		cmd.Set("dir-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := telnyx.DirVerifyEmailConfirmCodeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.VerifyEmail.ConfirmCode(
		ctx,
		cmd.Value("dir-id").(string),
		params,
		options...,
	)
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
		Title:          "dir:verify-email confirm-code",
		Transform:      transform,
	})
}

func handleDirVerifyEmailSendCode(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dir-id") && len(unusedArgs) > 0 {
		cmd.Set("dir-id", unusedArgs[0])
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
	_, err = client.Dir.VerifyEmail.SendCode(ctx, cmd.Value("dir-id").(string), options...)
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
		Title:          "dir:verify-email send-code",
		Transform:      transform,
	})
}

func handleDirVerifyEmailStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dir-id") && len(unusedArgs) > 0 {
		cmd.Set("dir-id", unusedArgs[0])
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
	_, err = client.Dir.VerifyEmail.Status(ctx, cmd.Value("dir-id").(string), options...)
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
		Title:          "dir:verify-email status",
		Transform:      transform,
	})
}

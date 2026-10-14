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

var enterprisesVerifyEmailCreate = cli.Command{
	Name:    "create",
	Usage:   "Email a 6-digit code to the enterprise account's contact email to confirm\nownership of that address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
	},
	Action:          handleEnterprisesVerifyEmailCreate,
	HideHelpCommand: true,
}

var enterprisesVerifyEmailConfirm = cli.Command{
	Name:    "confirm",
	Usage:   "Submit the 6-digit code that was emailed to the enterprise account's contact\nemail. On success the contact email is marked verified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:     "code",
			Usage:    "The 6-digit code sent to the enterprise account's contact email.",
			Required: true,
			BodyPath: "code",
		},
	},
	Action:          handleEnterprisesVerifyEmailConfirm,
	HideHelpCommand: true,
}

func handleEnterprisesVerifyEmailCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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
	_, err = client.Enterprises.VerifyEmail.New(ctx, cmd.Value("enterprise-id").(string), options...)
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
		Title:          "enterprises:verify-email create",
		Transform:      transform,
	})
}

func handleEnterprisesVerifyEmailConfirm(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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

	params := telnyx.EnterpriseVerifyEmailConfirmParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.VerifyEmail.Confirm(
		ctx,
		cmd.Value("enterprise-id").(string),
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
		Title:          "enterprises:verify-email confirm",
		Transform:      transform,
	})
}

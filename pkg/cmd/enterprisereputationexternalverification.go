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

var enterprisesReputationExternalVerificationCreate = cli.Command{
	Name:    "create",
	Usage:   "Request permission to add phone numbers that the enterprise owns but that are\nhosted with another provider. Submit one or more uploaded documents proving\nownership, such as a carrier invoice or account statement. The request is\nreviewed once per enterprise. While it is `pending` or `rejected` you can submit\nagain to replace the documents; once `approved` it cannot be changed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "document-id",
			Usage:    "Ids of uploaded documents (Documents API) proving the enterprise owns its phone numbers hosted with another provider, for example a carrier invoice or account statement.",
			Required: true,
			BodyPath: "document_ids",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "Optional. The phone numbers hosted with another provider that the documents cover, in E.164 format. Shown to the reviewer to match the documents to your numbers; you can still add numbers that are not listed here, and each is checked on its own.",
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleEnterprisesReputationExternalVerificationCreate,
	HideHelpCommand: true,
}

var enterprisesReputationExternalVerificationList = cli.Command{
	Name:    "list",
	Usage:   "Get the review status of the enterprise's ownership proof for phone numbers\nhosted with another provider. Returns 404 when no proof has been submitted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
	},
	Action:          handleEnterprisesReputationExternalVerificationList,
	HideHelpCommand: true,
}

func handleEnterprisesReputationExternalVerificationCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseReputationExternalVerificationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.ExternalVerification.New(
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
		Title:          "enterprises:reputation:external-verification create",
		Transform:      transform,
	})
}

func handleEnterprisesReputationExternalVerificationList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Enterprises.Reputation.ExternalVerification.List(ctx, cmd.Value("enterprise-id").(string), options...)
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
		Title:          "enterprises:reputation:external-verification list",
		Transform:      transform,
	})
}

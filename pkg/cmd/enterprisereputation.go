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

var enterprisesReputationRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Phone Number Reputation tracks how your outbound caller-IDs are perceived (spam\nrisk, engagement, etc.) across the call-screening ecosystem. This endpoint reads\nthe enterprise-level settings: whether the product is enabled, the refresh\ncadence, and the stored Letter of Authorization document id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
	},
	Action:          handleEnterprisesReputationRetrieve,
	HideHelpCommand: true,
}

var enterprisesReputationDisable = cli.Command{
	Name:    "disable",
	Usage:   "Disable Phone Number Reputation. All registered numbers are de-registered as a\ncascade. The enterprise itself is unaffected. Returns `204` on success, `404` if\nreputation is not enabled for this enterprise.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
	},
	Action:          handleEnterprisesReputationDisable,
	HideHelpCommand: true,
}

var enterprisesReputationEnable = cli.Command{
	Name:    "enable",
	Usage:   "Activate Phone Number Reputation for the given enterprise. Requires an uploaded\nLetter of Authorization document (the `loa_document_id` references the Telnyx\nDocuments API) and a refresh-frequency selection. After activation, individual\nphone numbers can be registered via `POST .../reputation/numbers`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:     "loa-document-id",
			Usage:    "Id of the signed Letter of Authorization document, returned by the Telnyx Documents API after upload (upload via `POST /v2/documents`; see https://developers.telnyx.com/api/documents).",
			Required: true,
			BodyPath: "loa_document_id",
		},
		&requestflag.Flag[string]{
			Name:     "check-frequency",
			Usage:    "How often Telnyx refreshes the stored reputation data for this enterprise's registered numbers.",
			BodyPath: "check_frequency",
		},
	},
	Action:          handleEnterprisesReputationEnable,
	HideHelpCommand: true,
}

var enterprisesReputationUpdateFrequency = cli.Command{
	Name:    "update-frequency",
	Usage:   "Update how often Telnyx refreshes the reputation data for this enterprise's\nregistered numbers. The new frequency takes effect on the next scheduled\nrefresh.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:     "check-frequency",
			Usage:    "How often Telnyx refreshes the stored reputation data for this enterprise's registered numbers.",
			Required: true,
			BodyPath: "check_frequency",
		},
	},
	Action:          handleEnterprisesReputationUpdateFrequency,
	HideHelpCommand: true,
}

func handleEnterprisesReputationRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Enterprises.Reputation.Get(ctx, cmd.Value("enterprise-id").(string), options...)
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
		Title:          "enterprises:reputation retrieve",
		Transform:      transform,
	})
}

func handleEnterprisesReputationDisable(ctx context.Context, cmd *cli.Command) error {
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

	return client.Enterprises.Reputation.Disable(ctx, cmd.Value("enterprise-id").(string), options...)
}

func handleEnterprisesReputationEnable(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseReputationEnableParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.Enable(
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
		Title:          "enterprises:reputation enable",
		Transform:      transform,
	})
}

func handleEnterprisesReputationUpdateFrequency(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseReputationUpdateFrequencyParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.UpdateFrequency(
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
		Title:          "enterprises:reputation update-frequency",
		Transform:      transform,
	})
}

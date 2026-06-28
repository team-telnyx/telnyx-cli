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

var termsOfServiceRetrieveInfo = cli.Command{
	Name:    "retrieve-info",
	Usage:   "Returns the available Terms of Service agreements (product, current version,\nterms URL, effective date). Omit `product_type` to return all products; pass it\nto scope to one.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "product-type",
			Usage:     "Telnyx product the Terms of Service apply to.",
			QueryPath: "product_type",
		},
	},
	Action:          handleTermsOfServiceRetrieveInfo,
	HideHelpCommand: true,
}

var termsOfServiceRetrieveStatus = cli.Command{
	Name:    "retrieve-status",
	Usage:   "Returns whether the authenticated user has agreed to the current Terms of\nService for the product given by `product_type`. Used during onboarding to\ndecide whether to prompt the user with the ToS dialog before continuing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "product-type",
			Usage:     "Telnyx product the Terms of Service apply to.",
			QueryPath: "product_type",
		},
	},
	Action:          handleTermsOfServiceRetrieveStatus,
	HideHelpCommand: true,
}

func handleTermsOfServiceRetrieveInfo(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.TermsOfServiceGetInfoParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TermsOfService.GetInfo(ctx, params, options...)
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
		Title:          "terms-of-service retrieve-info",
		Transform:      transform,
	})
}

func handleTermsOfServiceRetrieveStatus(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.TermsOfServiceGetStatusParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TermsOfService.GetStatus(ctx, params, options...)
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
		Title:          "terms-of-service retrieve-status",
		Transform:      transform,
	})
}

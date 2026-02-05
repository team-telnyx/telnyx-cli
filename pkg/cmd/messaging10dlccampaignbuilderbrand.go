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

var messaging10dlcCampaignBuilderBrandQualifyByUsecase = cli.Command{
	Name:    "qualify-by-usecase",
	Usage:   "This endpoint allows you to see whether or not the supplied brand is suitable\nfor your desired campaign use case.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "usecase",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcCampaignBuilderBrandQualifyByUsecase,
	HideHelpCommand: true,
}

func handleMessaging10dlcCampaignBuilderBrandQualifyByUsecase(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("usecase") && len(unusedArgs) > 0 {
		cmd.Set("usecase", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcCampaignBuilderBrandQualifyByUsecaseParams{
		BrandID: cmd.Value("brand-id").(string),
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
	_, err = client.Messaging10dlc.CampaignBuilder.Brand.QualifyByUsecase(
		ctx,
		cmd.Value("usecase").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign-builder:brand qualify-by-usecase", obj, format, transform)
}

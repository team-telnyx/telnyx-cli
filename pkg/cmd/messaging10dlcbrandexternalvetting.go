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

var messaging10dlcBrandExternalVettingList = cli.Command{
	Name:    "list",
	Usage:   "Get list of valid external vetting record for a given brand",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcBrandExternalVettingList,
	HideHelpCommand: true,
}

var messaging10dlcBrandExternalVettingImports = cli.Command{
	Name:    "imports",
	Usage:   "This operation can be used to import an external vetting record from a\nTCR-approved vetting provider. If the vetting provider confirms validity of the\nrecord, it will be saved with the brand and will be considered for future\ncampaign qualification.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "evp-id",
			Usage:    "External vetting provider ID for the brand.",
			Required: true,
			BodyPath: "evpId",
		},
		&requestflag.Flag[string]{
			Name:     "vetting-id",
			Usage:    "Unique ID that identifies a vetting transaction performed by a vetting provider. This ID is provided by the vetting provider at time of vetting.",
			Required: true,
			BodyPath: "vettingId",
		},
		&requestflag.Flag[string]{
			Name:     "vetting-token",
			Usage:    "Required by some providers for vetting record confirmation.",
			BodyPath: "vettingToken",
		},
	},
	Action:          handleMessaging10dlcBrandExternalVettingImports,
	HideHelpCommand: true,
}

var messaging10dlcBrandExternalVettingOrder = cli.Command{
	Name:    "order",
	Usage:   "Order new external vetting for a brand",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "evp-id",
			Usage:    "External vetting provider ID for the brand.",
			Required: true,
			BodyPath: "evpId",
		},
		&requestflag.Flag[string]{
			Name:     "vetting-class",
			Usage:    "Identifies the vetting classification.",
			Required: true,
			BodyPath: "vettingClass",
		},
	},
	Action:          handleMessaging10dlcBrandExternalVettingOrder,
	HideHelpCommand: true,
}

func handleMessaging10dlcBrandExternalVettingList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Brand.ExternalVetting.List(ctx, cmd.Value("brand-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand:external-vetting list", obj, format, transform)
}

func handleMessaging10dlcBrandExternalVettingImports(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcBrandExternalVettingImportsParams{}

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
	_, err = client.Messaging10dlc.Brand.ExternalVetting.Imports(
		ctx,
		cmd.Value("brand-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand:external-vetting imports", obj, format, transform)
}

func handleMessaging10dlcBrandExternalVettingOrder(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcBrandExternalVettingOrderParams{}

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
	_, err = client.Messaging10dlc.Brand.ExternalVetting.Order(
		ctx,
		cmd.Value("brand-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand:external-vetting order", obj, format, transform)
}

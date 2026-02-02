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

var messaging10dlcPhoneNumberCampaignsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create New Phone Number Campaign",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Usage:    "The ID of the campaign you want to link to the specified phone number.",
			Required: true,
			BodyPath: "campaignId",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "The phone number you want to link to a specified campaign.",
			Required: true,
			BodyPath: "phoneNumber",
		},
	},
	Action:          handleMessaging10dlcPhoneNumberCampaignsCreate,
	HideHelpCommand: true,
}

var messaging10dlcPhoneNumberCampaignsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an individual phone number/campaign assignment by `phoneNumber`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcPhoneNumberCampaignsRetrieve,
	HideHelpCommand: true,
}

var messaging10dlcPhoneNumberCampaignsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Create New Phone Number Campaign",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-phone-number",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Usage:    "The ID of the campaign you want to link to the specified phone number.",
			Required: true,
			BodyPath: "campaignId",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "The phone number you want to link to a specified campaign.",
			Required: true,
			BodyPath: "phoneNumber",
		},
	},
	Action:          handleMessaging10dlcPhoneNumberCampaignsUpdate,
	HideHelpCommand: true,
}

var messaging10dlcPhoneNumberCampaignsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List phone number campaigns",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[telnyx_campaign_id], filter[telnyx_brand_id], filter[tcr_campaign_id], filter[tcr_brand_id]",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "records-per-page",
			Default:   20,
			QueryPath: "recordsPerPage",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. If not given, results are sorted by createdAt in descending order.",
			Default:   "-createdAt",
			QueryPath: "sort",
		},
	},
	Action:          handleMessaging10dlcPhoneNumberCampaignsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.tcr-brand-id",
			Usage:      "Filter results by the TCR Brand id",
			InnerField: "tcr_brand_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.tcr-campaign-id",
			Usage:      "Filter results by the TCR Campaign id",
			InnerField: "tcr_campaign_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.telnyx-brand-id",
			Usage:      "Filter results by the Telnyx Brand id",
			InnerField: "telnyx_brand_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.telnyx-campaign-id",
			Usage:      "Filter results by the Telnyx Campaign id",
			InnerField: "telnyx_campaign_id",
		},
	},
})

var messaging10dlcPhoneNumberCampaignsDelete = cli.Command{
	Name:    "delete",
	Usage:   "This endpoint allows you to remove a campaign assignment from the supplied\n`phoneNumber`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcPhoneNumberCampaignsDelete,
	HideHelpCommand: true,
}

func handleMessaging10dlcPhoneNumberCampaignsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPhoneNumberCampaignNewParams{}

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
	_, err = client.Messaging10dlc.PhoneNumberCampaigns.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-campaigns create", obj, format, transform)
}

func handleMessaging10dlcPhoneNumberCampaignsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
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
	_, err = client.Messaging10dlc.PhoneNumberCampaigns.Get(ctx, cmd.Value("phone-number").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-campaigns retrieve", obj, format, transform)
}

func handleMessaging10dlcPhoneNumberCampaignsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-phone-number") && len(unusedArgs) > 0 {
		cmd.Set("campaign-phone-number", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPhoneNumberCampaignUpdateParams{}

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
	_, err = client.Messaging10dlc.PhoneNumberCampaigns.Update(
		ctx,
		cmd.Value("campaign-phone-number").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-campaigns update", obj, format, transform)
}

func handleMessaging10dlcPhoneNumberCampaignsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPhoneNumberCampaignListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Messaging10dlc.PhoneNumberCampaigns.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-campaigns list", obj, format, transform)
	} else {
		iter := client.Messaging10dlc.PhoneNumberCampaigns.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "messaging-10dlc:phone-number-campaigns list", iter, format, transform)
	}
}

func handleMessaging10dlcPhoneNumberCampaignsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number") && len(unusedArgs) > 0 {
		cmd.Set("phone-number", unusedArgs[0])
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
	_, err = client.Messaging10dlc.PhoneNumberCampaigns.Delete(ctx, cmd.Value("phone-number").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-campaigns delete", obj, format, transform)
}

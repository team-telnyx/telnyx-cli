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

var messaging10dlcCampaignRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve campaign details by `campaignId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcCampaignRetrieve,
	HideHelpCommand: true,
}

var messaging10dlcCampaignUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a campaign's properties by `campaignId`. **Please note:** only sample\nmessages are editable.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "auto-renewal",
			Usage:    "Help message of the campaign.",
			Default:  true,
			BodyPath: "autoRenewal",
		},
		&requestflag.Flag[string]{
			Name:     "help-message",
			Usage:    "Help message of the campaign.",
			BodyPath: "helpMessage",
		},
		&requestflag.Flag[string]{
			Name:     "message-flow",
			Usage:    "Message flow description.",
			BodyPath: "messageFlow",
		},
		&requestflag.Flag[string]{
			Name:     "reseller-id",
			Usage:    "Alphanumeric identifier of the reseller that you want to associate with this campaign.",
			BodyPath: "resellerId",
		},
		&requestflag.Flag[string]{
			Name:     "sample1",
			Usage:    "Message sample. Some campaign tiers require 1 or more message samples.",
			BodyPath: "sample1",
		},
		&requestflag.Flag[string]{
			Name:     "sample2",
			Usage:    "Message sample. Some campaign tiers require 2 or more message samples.",
			BodyPath: "sample2",
		},
		&requestflag.Flag[string]{
			Name:     "sample3",
			Usage:    "Message sample. Some campaign tiers require 3 or more message samples.",
			BodyPath: "sample3",
		},
		&requestflag.Flag[string]{
			Name:     "sample4",
			Usage:    "Message sample. Some campaign tiers require 4 or more message samples.",
			BodyPath: "sample4",
		},
		&requestflag.Flag[string]{
			Name:     "sample5",
			Usage:    "Message sample. Some campaign tiers require 5 or more message samples.",
			BodyPath: "sample5",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "Webhook failover to which campaign status updates are sent.",
			BodyPath: "webhookFailoverURL",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Webhook to which campaign status updates are sent.",
			BodyPath: "webhookURL",
		},
	},
	Action:          handleMessaging10dlcCampaignUpdate,
	HideHelpCommand: true,
}

var messaging10dlcCampaignList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of campaigns associated with a supplied `brandId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "brand-id",
			Required:  true,
			QueryPath: "brandId",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "The 1-indexed page number to get. The default value is `1`.",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "records-per-page",
			Usage:     "The amount of records per page, limited to between 1 and 500 inclusive. The default value is `10`.",
			Default:   10,
			QueryPath: "recordsPerPage",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. If not given, results are sorted by createdAt in descending order.",
			Default:   "-createdAt",
			QueryPath: "sort",
		},
	},
	Action:          handleMessaging10dlcCampaignList,
	HideHelpCommand: true,
}

var messaging10dlcCampaignAcceptSharing = cli.Command{
	Name:    "accept-sharing",
	Usage:   "Manually accept a campaign shared with Telnyx",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Usage:    "TCR's ID for the campaign to import",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcCampaignAcceptSharing,
	HideHelpCommand: true,
}

var messaging10dlcCampaignDeactivate = cli.Command{
	Name:    "deactivate",
	Usage:   "Terminate a campaign. Note that once deactivated, a campaign cannot be restored.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcCampaignDeactivate,
	HideHelpCommand: true,
}

var messaging10dlcCampaignGetMnoMetadata = cli.Command{
	Name:    "get-mno-metadata",
	Usage:   "Get the campaign metadata for each MNO it was submitted to.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Usage:    "ID of the campaign in question",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcCampaignGetMnoMetadata,
	HideHelpCommand: true,
}

var messaging10dlcCampaignGetOperationStatus = cli.Command{
	Name:    "get-operation-status",
	Usage:   "Retrieve campaign's operation status at MNO level.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcCampaignGetOperationStatus,
	HideHelpCommand: true,
}

var messaging10dlcCampaignGetSharingStatus = cli.Command{
	Name:    "get-sharing-status",
	Usage:   "Get Sharing Status",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Usage:    "ID of the campaign in question",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcCampaignGetSharingStatus,
	HideHelpCommand: true,
}

var messaging10dlcCampaignSubmitAppeal = cli.Command{
	Name:    "submit-appeal",
	Usage:   "Submits an appeal for rejected native campaigns in TELNYX_FAILED or MNO_REJECTED\nstatus. The appeal is recorded for manual compliance team review and the\ncampaign status is reset to TCR_ACCEPTED. Note: Appeal forwarding is handled\nmanually to allow proper review before incurring upstream charges.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "appeal-reason",
			Usage:    "Detailed explanation of why the campaign should be reconsidered and what changes have been made to address the rejection reason.",
			Required: true,
			BodyPath: "appeal_reason",
		},
	},
	Action:          handleMessaging10dlcCampaignSubmitAppeal,
	HideHelpCommand: true,
}

func handleMessaging10dlcCampaignRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Campaign.Get(ctx, cmd.Value("campaign-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign retrieve", obj, format, transform)
}

func handleMessaging10dlcCampaignUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcCampaignUpdateParams{}

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
	_, err = client.Messaging10dlc.Campaign.Update(
		ctx,
		cmd.Value("campaign-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign update", obj, format, transform)
}

func handleMessaging10dlcCampaignList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcCampaignListParams{}

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
		_, err = client.Messaging10dlc.Campaign.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messaging-10dlc:campaign list", obj, format, transform)
	} else {
		iter := client.Messaging10dlc.Campaign.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "messaging-10dlc:campaign list", iter, format, transform)
	}
}

func handleMessaging10dlcCampaignAcceptSharing(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Campaign.AcceptSharing(ctx, cmd.Value("campaign-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign accept-sharing", obj, format, transform)
}

func handleMessaging10dlcCampaignDeactivate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Campaign.Deactivate(ctx, cmd.Value("campaign-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign deactivate", obj, format, transform)
}

func handleMessaging10dlcCampaignGetMnoMetadata(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Campaign.GetMnoMetadata(ctx, cmd.Value("campaign-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign get-mno-metadata", obj, format, transform)
}

func handleMessaging10dlcCampaignGetOperationStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Campaign.GetOperationStatus(ctx, cmd.Value("campaign-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign get-operation-status", obj, format, transform)
}

func handleMessaging10dlcCampaignGetSharingStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Campaign.GetSharingStatus(ctx, cmd.Value("campaign-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign get-sharing-status", obj, format, transform)
}

func handleMessaging10dlcCampaignSubmitAppeal(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcCampaignSubmitAppealParams{}

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
	_, err = client.Messaging10dlc.Campaign.SubmitAppeal(
		ctx,
		cmd.Value("campaign-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:campaign submit-appeal", obj, format, transform)
}

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

var messaging10dlcPartnerCampaignsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve campaign details by `campaignId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcPartnerCampaignsRetrieve,
	HideHelpCommand: true,
}

var messaging10dlcPartnerCampaignsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update campaign details by `campaignId`. **Please note:** Only webhook urls are\neditable.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Required: true,
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
	Action:          handleMessaging10dlcPartnerCampaignsUpdate,
	HideHelpCommand: true,
}

var messaging10dlcPartnerCampaignsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve all partner campaigns you have shared to Telnyx in a paginated fashion.",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleMessaging10dlcPartnerCampaignsList,
	HideHelpCommand: true,
}

var messaging10dlcPartnerCampaignsListSharedByMe = cli.Command{
	Name:    "list-shared-by-me",
	Usage:   "Get all partner campaigns you have shared to Telnyx in a paginated fashion",
	Suggest: true,
	Flags: []cli.Flag{
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
	},
	Action:          handleMessaging10dlcPartnerCampaignsListSharedByMe,
	HideHelpCommand: true,
}

var messaging10dlcPartnerCampaignsRetrieveSharingStatus = cli.Command{
	Name:    "retrieve-sharing-status",
	Usage:   "Get Sharing Status",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Usage:    "ID of the campaign in question",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcPartnerCampaignsRetrieveSharingStatus,
	HideHelpCommand: true,
}

func handleMessaging10dlcPartnerCampaignsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Messaging10dlc.PartnerCampaigns.Get(ctx, cmd.Value("campaign-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:partner-campaigns retrieve", obj, format, transform)
}

func handleMessaging10dlcPartnerCampaignsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("campaign-id") && len(unusedArgs) > 0 {
		cmd.Set("campaign-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPartnerCampaignUpdateParams{}

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
	_, err = client.Messaging10dlc.PartnerCampaigns.Update(
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
	return ShowJSON(os.Stdout, "messaging-10dlc:partner-campaigns update", obj, format, transform)
}

func handleMessaging10dlcPartnerCampaignsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPartnerCampaignListParams{}

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
		_, err = client.Messaging10dlc.PartnerCampaigns.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messaging-10dlc:partner-campaigns list", obj, format, transform)
	} else {
		iter := client.Messaging10dlc.PartnerCampaigns.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "messaging-10dlc:partner-campaigns list", iter, format, transform)
	}
}

func handleMessaging10dlcPartnerCampaignsListSharedByMe(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPartnerCampaignListSharedByMeParams{}

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
		_, err = client.Messaging10dlc.PartnerCampaigns.ListSharedByMe(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messaging-10dlc:partner-campaigns list-shared-by-me", obj, format, transform)
	} else {
		iter := client.Messaging10dlc.PartnerCampaigns.ListSharedByMeAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "messaging-10dlc:partner-campaigns list-shared-by-me", iter, format, transform)
	}
}

func handleMessaging10dlcPartnerCampaignsRetrieveSharingStatus(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Messaging10dlc.PartnerCampaigns.GetSharingStatus(ctx, cmd.Value("campaign-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:partner-campaigns retrieve-sharing-status", obj, format, transform)
}

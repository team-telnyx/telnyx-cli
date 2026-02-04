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

var messaging10dlcPhoneNumberAssignmentByProfileAssign = cli.Command{
	Name:    "assign",
	Usage:   "This endpoint allows you to link all phone numbers associated with a Messaging\nProfile to a campaign. **Please note:** if you want to assign phone numbers to a\ncampaign that you did not create with Telnyx 10DLC services, this endpoint\nallows that provided that you've shared the campaign with Telnyx. In this case,\nonly provide the parameter, `tcrCampaignId`, and not `campaignId`. In all other\ncases (where the campaign you're assigning was created with Telnyx 10DLC\nservices), only provide `campaignId`, not `tcrCampaignId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "The ID of the messaging profile that you want to link to the specified campaign.",
			Required: true,
			BodyPath: "messagingProfileId",
		},
		&requestflag.Flag[string]{
			Name:     "campaign-id",
			Usage:    "The ID of the campaign you want to link to the specified messaging profile. If you supply this ID in the request, do not also include a tcrCampaignId.",
			BodyPath: "campaignId",
		},
		&requestflag.Flag[string]{
			Name:     "tcr-campaign-id",
			Usage:    "The TCR ID of the shared campaign you want to link to the specified messaging profile (for campaigns not created using Telnyx 10DLC services only). If you supply this ID in the request, do not also include a campaignId.",
			BodyPath: "tcrCampaignId",
		},
	},
	Action:          handleMessaging10dlcPhoneNumberAssignmentByProfileAssign,
	HideHelpCommand: true,
}

var messaging10dlcPhoneNumberAssignmentByProfileListPhoneNumberStatus = cli.Command{
	Name:    "list-phone-number-status",
	Usage:   "Check the status of the individual phone number/campaign assignments associated\nwith the supplied `taskId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "task-id",
			Required: true,
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
	},
	Action:          handleMessaging10dlcPhoneNumberAssignmentByProfileListPhoneNumberStatus,
	HideHelpCommand: true,
}

var messaging10dlcPhoneNumberAssignmentByProfileRetrievePhoneNumberStatus = cli.Command{
	Name:    "retrieve-phone-number-status",
	Usage:   "Check the status of the individual phone number/campaign assignments associated\nwith the supplied `taskId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "task-id",
			Required: true,
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
	},
	Action:          handleMessaging10dlcPhoneNumberAssignmentByProfileRetrievePhoneNumberStatus,
	HideHelpCommand: true,
}

var messaging10dlcPhoneNumberAssignmentByProfileRetrieveStatus = cli.Command{
	Name:    "retrieve-status",
	Usage:   "Check the status of the task associated with assigning all phone numbers on a\nmessaging profile to a campaign by `taskId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "task-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcPhoneNumberAssignmentByProfileRetrieveStatus,
	HideHelpCommand: true,
}

func handleMessaging10dlcPhoneNumberAssignmentByProfileAssign(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPhoneNumberAssignmentByProfileAssignParams{}

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
	_, err = client.Messaging10dlc.PhoneNumberAssignmentByProfile.Assign(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-assignment-by-profile assign", obj, format, transform)
}

func handleMessaging10dlcPhoneNumberAssignmentByProfileListPhoneNumberStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("task-id") && len(unusedArgs) > 0 {
		cmd.Set("task-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPhoneNumberAssignmentByProfileListPhoneNumberStatusParams{}

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
	_, err = client.Messaging10dlc.PhoneNumberAssignmentByProfile.ListPhoneNumberStatus(
		ctx,
		cmd.Value("task-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-assignment-by-profile list-phone-number-status", obj, format, transform)
}

func handleMessaging10dlcPhoneNumberAssignmentByProfileRetrievePhoneNumberStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("task-id") && len(unusedArgs) > 0 {
		cmd.Set("task-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusParams{}

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
	_, err = client.Messaging10dlc.PhoneNumberAssignmentByProfile.GetPhoneNumberStatus(
		ctx,
		cmd.Value("task-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-assignment-by-profile retrieve-phone-number-status", obj, format, transform)
}

func handleMessaging10dlcPhoneNumberAssignmentByProfileRetrieveStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("task-id") && len(unusedArgs) > 0 {
		cmd.Set("task-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.PhoneNumberAssignmentByProfile.GetStatus(ctx, cmd.Value("task-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:phone-number-assignment-by-profile retrieve-status", obj, format, transform)
}

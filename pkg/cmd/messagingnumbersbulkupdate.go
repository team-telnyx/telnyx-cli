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

var messagingNumbersBulkUpdatesCreate = cli.Command{
	Name:    "create",
	Usage:   "Bulk update phone number profiles",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "Configure the messaging profile these phone numbers are assigned to:\n\n* Set this field to `\"\"` to unassign each number from their respective messaging profile\n* Set this field to a quoted UUID of a messaging profile to assign these numbers to that messaging profile",
			Required: true,
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "number",
			Usage:    "The list of phone numbers to update.",
			Required: true,
			BodyPath: "numbers",
		},
		&requestflag.Flag[bool]{
			Name:     "assign-only",
			Usage:    "If true, only assign numbers to the profile without changing other settings.",
			Default:  false,
			BodyPath: "assign_only",
		},
	},
	Action:          handleMessagingNumbersBulkUpdatesCreate,
	HideHelpCommand: true,
}

var messagingNumbersBulkUpdatesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve bulk update status",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "order-id",
			Required: true,
		},
	},
	Action:          handleMessagingNumbersBulkUpdatesRetrieve,
	HideHelpCommand: true,
}

func handleMessagingNumbersBulkUpdatesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingNumbersBulkUpdateNewParams{}

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
	_, err = client.MessagingNumbersBulkUpdates.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-numbers-bulk-updates create", obj, format, transform)
}

func handleMessagingNumbersBulkUpdatesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
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
	_, err = client.MessagingNumbersBulkUpdates.Get(ctx, cmd.Value("order-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-numbers-bulk-updates retrieve", obj, format, transform)
}

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

var notificationSettingsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Add a notification setting.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "notification-channel-id",
			Usage:    "A UUID reference to the associated Notification Channel.",
			BodyPath: "notification_channel_id",
		},
		&requestflag.Flag[string]{
			Name:     "notification-event-condition-id",
			Usage:    "A UUID reference to the associated Notification Event Condition.",
			BodyPath: "notification_event_condition_id",
		},
		&requestflag.Flag[string]{
			Name:     "notification-profile-id",
			Usage:    "A UUID reference to the associated Notification Profile.",
			BodyPath: "notification_profile_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "parameter",
			BodyPath: "parameters",
		},
	},
	Action:          handleNotificationSettingsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"parameter": {
		&requestflag.InnerFlag[string]{
			Name:       "parameter.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "parameter.value",
			InnerField: "value",
		},
	},
})

var notificationSettingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a notification setting.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleNotificationSettingsRetrieve,
	HideHelpCommand: true,
}

var notificationSettingsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List notification settings.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[associated_record_type][eq], filter[channel_type_id][eq], filter[notification_profile_id][eq], filter[notification_channel][eq], filter[notification_event_condition_id][eq], filter[status][eq]",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleNotificationSettingsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.associated-record-type",
			InnerField: "associated_record_type",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.channel-type-id",
			InnerField: "channel_type_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.notification-channel",
			InnerField: "notification_channel",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.notification-event-condition-id",
			InnerField: "notification_event_condition_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.notification-profile-id",
			InnerField: "notification_profile_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.status",
			InnerField: "status",
		},
	},
})

var notificationSettingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a notification setting.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleNotificationSettingsDelete,
	HideHelpCommand: true,
}

func handleNotificationSettingsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NotificationSettingNewParams{}

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
	_, err = client.NotificationSettings.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notification-settings create", obj, format, transform)
}

func handleNotificationSettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.NotificationSettings.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notification-settings retrieve", obj, format, transform)
}

func handleNotificationSettingsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NotificationSettingListParams{}

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
		_, err = client.NotificationSettings.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "notification-settings list", obj, format, transform)
	} else {
		iter := client.NotificationSettings.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "notification-settings list", iter, format, transform)
	}
}

func handleNotificationSettingsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.NotificationSettings.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notification-settings delete", obj, format, transform)
}

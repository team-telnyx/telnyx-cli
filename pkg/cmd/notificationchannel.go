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

var notificationChannelsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a notification channel.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "channel-destination",
			Usage:    "The destination associated with the channel type.",
			BodyPath: "channel_destination",
		},
		&requestflag.Flag[string]{
			Name:     "channel-type-id",
			Usage:    "A Channel Type ID",
			BodyPath: "channel_type_id",
		},
		&requestflag.Flag[string]{
			Name:     "notification-profile-id",
			Usage:    "A UUID reference to the associated Notification Profile.",
			BodyPath: "notification_profile_id",
		},
	},
	Action:          handleNotificationChannelsCreate,
	HideHelpCommand: true,
}

var notificationChannelsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a notification channel.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleNotificationChannelsRetrieve,
	HideHelpCommand: true,
}

var notificationChannelsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a notification channel.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "notification-channel-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "channel-destination",
			Usage:    "The destination associated with the channel type.",
			BodyPath: "channel_destination",
		},
		&requestflag.Flag[string]{
			Name:     "channel-type-id",
			Usage:    "A Channel Type ID",
			BodyPath: "channel_type_id",
		},
		&requestflag.Flag[string]{
			Name:     "notification-profile-id",
			Usage:    "A UUID reference to the associated Notification Profile.",
			BodyPath: "notification_profile_id",
		},
	},
	Action:          handleNotificationChannelsUpdate,
	HideHelpCommand: true,
}

var notificationChannelsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List notification channels.",
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
	Action:          handleNotificationChannelsList,
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

var notificationChannelsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a notification channel.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleNotificationChannelsDelete,
	HideHelpCommand: true,
}

func handleNotificationChannelsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NotificationChannelNewParams{}

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
	_, err = client.NotificationChannels.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notification-channels create", obj, format, transform)
}

func handleNotificationChannelsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.NotificationChannels.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notification-channels retrieve", obj, format, transform)
}

func handleNotificationChannelsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("notification-channel-id") && len(unusedArgs) > 0 {
		cmd.Set("notification-channel-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NotificationChannelUpdateParams{}

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
	_, err = client.NotificationChannels.Update(
		ctx,
		cmd.Value("notification-channel-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notification-channels update", obj, format, transform)
}

func handleNotificationChannelsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NotificationChannelListParams{}

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
		_, err = client.NotificationChannels.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "notification-channels list", obj, format, transform)
	} else {
		iter := client.NotificationChannels.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "notification-channels list", iter, format, transform)
	}
}

func handleNotificationChannelsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.NotificationChannels.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notification-channels delete", obj, format, transform)
}

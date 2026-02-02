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

var simCardDataUsageNotificationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new SIM card data usage notification.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sim-card-id",
			Usage:    "The identification UUID of the related SIM card resource.",
			Required: true,
			BodyPath: "sim_card_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "threshold",
			Usage:    "Data usage threshold that will trigger the notification.",
			Required: true,
			BodyPath: "threshold",
		},
	},
	Action:          handleSimCardDataUsageNotificationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"threshold": {
		&requestflag.InnerFlag[string]{
			Name:       "threshold.amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "threshold.unit",
			InnerField: "unit",
		},
	},
})

var simCardDataUsageNotificationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single SIM Card Data Usage Notification.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardDataUsageNotificationsRetrieve,
	HideHelpCommand: true,
}

var simCardDataUsageNotificationsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates information for a SIM Card Data Usage Notification.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "sim-card-data-usage-notification-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "sim-card-id",
			Usage:    "The identification UUID of the related SIM card resource.",
			BodyPath: "sim_card_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "threshold",
			Usage:    "Data usage threshold that will trigger the notification.",
			BodyPath: "threshold",
		},
	},
	Action:          handleSimCardDataUsageNotificationsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"threshold": {
		&requestflag.InnerFlag[string]{
			Name:       "threshold.amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "threshold.unit",
			InnerField: "unit",
		},
	},
})

var simCardDataUsageNotificationsList = cli.Command{
	Name:    "list",
	Usage:   "Lists a paginated collection of SIM card data usage notifications. It enables\nexploring the collection using specific filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-sim-card-id",
			Usage:     "A valid SIM card ID.",
			QueryPath: "filter[sim_card_id]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page.",
			Default:   20,
			QueryPath: "page[size]",
		},
	},
	Action:          handleSimCardDataUsageNotificationsList,
	HideHelpCommand: true,
}

var simCardDataUsageNotificationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete the SIM Card Data Usage Notification.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSimCardDataUsageNotificationsDelete,
	HideHelpCommand: true,
}

func handleSimCardDataUsageNotificationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardDataUsageNotificationNewParams{}

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
	_, err = client.SimCardDataUsageNotifications.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-data-usage-notifications create", obj, format, transform)
}

func handleSimCardDataUsageNotificationsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCardDataUsageNotifications.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-data-usage-notifications retrieve", obj, format, transform)
}

func handleSimCardDataUsageNotificationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sim-card-data-usage-notification-id") && len(unusedArgs) > 0 {
		cmd.Set("sim-card-data-usage-notification-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardDataUsageNotificationUpdateParams{}

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
	_, err = client.SimCardDataUsageNotifications.Update(
		ctx,
		cmd.Value("sim-card-data-usage-notification-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-data-usage-notifications update", obj, format, transform)
}

func handleSimCardDataUsageNotificationsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardDataUsageNotificationListParams{}

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
		_, err = client.SimCardDataUsageNotifications.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "sim-card-data-usage-notifications list", obj, format, transform)
	} else {
		iter := client.SimCardDataUsageNotifications.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "sim-card-data-usage-notifications list", iter, format, transform)
	}
}

func handleSimCardDataUsageNotificationsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SimCardDataUsageNotifications.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-data-usage-notifications delete", obj, format, transform)
}

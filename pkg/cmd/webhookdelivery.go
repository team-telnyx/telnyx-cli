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

var webhookDeliveriesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Provides webhook_delivery debug data, such as timestamps, delivery status and\nattempts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleWebhookDeliveriesRetrieve,
	HideHelpCommand: true,
}

var webhookDeliveriesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Lists webhook_deliveries for the authenticated user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[status][eq], filter[event_type], filter[webhook][contains], filter[attempts][contains], filter[started_at][gte], filter[started_at][lte], filter[finished_at][gte], filter[finished_at][lte]",
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
	Action:          handleWebhookDeliveriesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.attempts",
			InnerField: "attempts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.event-type",
			Usage:      "Return only webhook_deliveries matching the given value of `event_type`. Accepts multiple values separated by a `,`.",
			InnerField: "event_type",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.finished-at",
			InnerField: "finished_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.started-at",
			InnerField: "started_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.status",
			InnerField: "status",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.webhook",
			InnerField: "webhook",
		},
	},
})

func handleWebhookDeliveriesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WebhookDeliveries.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhook-deliveries retrieve", obj, format, transform)
}

func handleWebhookDeliveriesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.WebhookDeliveryListParams{}

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
		_, err = client.WebhookDeliveries.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "webhook-deliveries list", obj, format, transform)
	} else {
		iter := client.WebhookDeliveries.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "webhook-deliveries list", iter, format, transform)
	}
}

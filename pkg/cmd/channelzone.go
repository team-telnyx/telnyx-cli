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

var channelZonesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the number of Voice Channels for the Non-US Zones. This allows your\naccount to handle multiple simultaneous inbound calls to Non-US numbers. Use\nthis endpoint to increase or decrease your capacity based on expected call\nvolume.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "channel-zone-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "channels",
			Usage:    "The number of reserved channels",
			Required: true,
			BodyPath: "channels",
		},
	},
	Action:          handleChannelZonesUpdate,
	HideHelpCommand: true,
}

var channelZonesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns the non-US voice channels for your account. voice channels allow you to\nuse Channel Billing for calls to your Telnyx phone numbers. Please check the\n<a href=\"https://support.telnyx.com/en/articles/8428806-global-channel-billing\">Telnyx\nSupport Articles</a> section for full information and examples of how to utilize\nChannel Billing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
	},
	Action:          handleChannelZonesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"page": {
		&requestflag.InnerFlag[int64]{
			Name:       "page.number",
			Usage:      "The page number to load",
			InnerField: "number",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.size",
			Usage:      "The size of the page",
			InnerField: "size",
		},
	},
})

func handleChannelZonesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("channel-zone-id") && len(unusedArgs) > 0 {
		cmd.Set("channel-zone-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ChannelZoneUpdateParams{}

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
	_, err = client.ChannelZones.Update(
		ctx,
		cmd.Value("channel-zone-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "channel-zones update", obj, format, transform)
}

func handleChannelZonesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ChannelZoneListParams{}

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
		_, err = client.ChannelZones.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "channel-zones list", obj, format, transform)
	} else {
		iter := client.ChannelZones.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "channel-zones list", iter, format, transform)
	}
}

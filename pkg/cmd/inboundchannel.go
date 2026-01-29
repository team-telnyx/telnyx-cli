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

var inboundChannelsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the number of Voice Channels for the US Zone. This allows your account to\nhandle multiple simultaneous inbound calls to US numbers. Use this endpoint to\nincrease or decrease your capacity based on expected call volume.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "channels",
			Usage:    "The new number of concurrent channels for the account",
			Required: true,
			BodyPath: "channels",
		},
	},
	Action:          handleInboundChannelsUpdate,
	HideHelpCommand: true,
}

var inboundChannelsList = cli.Command{
	Name:            "list",
	Usage:           "Returns the US Zone voice channels for your account. voice channels allows you\nto use Channel Billing for calls to your Telnyx phone numbers. Please check the\n<a href=\"https://support.telnyx.com/en/articles/8428806-global-channel-billing\">Telnyx\nSupport Articles</a> section for full information and examples of how to utilize\nChannel Billing.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleInboundChannelsList,
	HideHelpCommand: true,
}

func handleInboundChannelsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.InboundChannelUpdateParams{}

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
	_, err = client.InboundChannels.Update(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "inbound-channels update", obj, format, transform)
}

func handleInboundChannelsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.InboundChannels.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "inbound-channels list", obj, format, transform)
}

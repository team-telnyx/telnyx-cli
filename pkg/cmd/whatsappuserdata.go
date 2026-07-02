// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/stainless-sdks/telnyx-go/v4"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var whatsappUserDataRetrieve = cli.Command{
	Name:            "retrieve",
	Usage:           "Fetch Whatsapp user data",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWhatsappUserDataRetrieve,
	HideHelpCommand: true,
}

var whatsappUserDataUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update Whatsapp user data",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "Failover URL to send Whatsapp signup events",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "URL to send Whatsapp signup events",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleWhatsappUserDataUpdate,
	HideHelpCommand: true,
}

func handleWhatsappUserDataRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Whatsapp.UserData.Get(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "whatsapp:user-data retrieve",
		Transform:      transform,
	})
}

func handleWhatsappUserDataUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.WhatsappUserDataUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Whatsapp.UserData.Update(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "whatsapp:user-data update",
		Transform:      transform,
	})
}

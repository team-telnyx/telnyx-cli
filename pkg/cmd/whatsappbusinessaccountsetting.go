// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var whatsappBusinessAccountsSettingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get WABA settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleWhatsappBusinessAccountsSettingsRetrieve,
	HideHelpCommand: true,
}

var whatsappBusinessAccountsSettingsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update WABA settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "timezone",
			Usage:    "IANA timezone identifier",
			BodyPath: "timezone",
		},
		&requestflag.Flag[bool]{
			Name:     "webhook-enabled",
			Usage:    "Enable/disable receiving Whatsapp events",
			BodyPath: "webhook_enabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "webhook-event",
			BodyPath: "webhook_events",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "Failover URL to send Whatsapp events",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "URL to send Whatsapp events",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleWhatsappBusinessAccountsSettingsUpdate,
	HideHelpCommand: true,
}

func handleWhatsappBusinessAccountsSettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Whatsapp.BusinessAccounts.Settings.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "whatsapp:business-accounts:settings retrieve",
		Transform:      transform,
	})
}

func handleWhatsappBusinessAccountsSettingsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.WhatsappBusinessAccountSettingUpdateParams{}

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
	_, err = client.Whatsapp.BusinessAccounts.Settings.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
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
		Title:          "whatsapp:business-accounts:settings update",
		Transform:      transform,
	})
}

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

var fqdnConnectionsFqdnAuthenticationList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves the details of an existing FQDN authentication strategy for a specific\nFQDN connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "fqdn-connection-id",
			Required:  true,
			PathParam: "fqdn_connection_id",
		},
	},
	Action:          handleFqdnConnectionsFqdnAuthenticationList,
	HideHelpCommand: true,
}

var fqdnConnectionsFqdnAuthenticationPatchAll = cli.Command{
	Name:    "patch-all",
	Usage:   "Updates the FQDN authentication strategy for a specific FQDN connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "fqdn-connection-id",
			Required:  true,
			PathParam: "fqdn_connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "failover-url",
			Usage:    "The failover webhook URL.",
			BodyPath: "failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "fqdn-outbound-authentication",
			Usage:    "The outbound authentication type.",
			BodyPath: "fqdn_outbound_authentication",
		},
		&requestflag.Flag[string]{
			Name:     "ip-authentication-method",
			Usage:    "The IP authentication method.",
			BodyPath: "ip_authentication_method",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			Usage:    "The password for authentication.",
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "txt-name",
			Usage:    "The TXT record name for Microsoft Teams SBC DNS verification.",
			BodyPath: "txt_name",
		},
		&requestflag.Flag[int64]{
			Name:     "txt-ttl",
			Usage:    "The TTL for the TXT record.",
			BodyPath: "txt_ttl",
		},
		&requestflag.Flag[string]{
			Name:     "txt-value",
			Usage:    "The TXT record value for Microsoft Teams SBC DNS verification.",
			BodyPath: "txt_value",
		},
		&requestflag.Flag[string]{
			Name:     "user-name",
			Usage:    "The username for authentication.",
			BodyPath: "user_name",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "The webhook URL for authentication events.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleFqdnConnectionsFqdnAuthenticationPatchAll,
	HideHelpCommand: true,
}

func handleFqdnConnectionsFqdnAuthenticationList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fqdn-connection-id") && len(unusedArgs) > 0 {
		cmd.Set("fqdn-connection-id", unusedArgs[0])
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
	_, err = client.FqdnConnections.FqdnAuthentication.List(ctx, cmd.Value("fqdn-connection-id").(string), options...)
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
		Title:          "fqdn-connections:fqdn-authentication list",
		Transform:      transform,
	})
}

func handleFqdnConnectionsFqdnAuthenticationPatchAll(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fqdn-connection-id") && len(unusedArgs) > 0 {
		cmd.Set("fqdn-connection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := telnyx.FqdnConnectionFqdnAuthenticationPatchAllParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FqdnConnections.FqdnAuthentication.PatchAll(
		ctx,
		cmd.Value("fqdn-connection-id").(string),
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "fqdn-connections:fqdn-authentication patch-all",
		Transform:      transform,
	})
}

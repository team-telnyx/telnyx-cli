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

var externalConnectionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new External Connection based on the parameters sent in the request.\nThe external_sip_connection and outbound voice profile id are required. Once\ncreated, you can assign phone numbers to your application using the\n`/phone_numbers` endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "external-sip-connection",
			Usage:    "The service that will be consuming this connection.",
			Default:  "zoom",
			Required: true,
			BodyPath: "external_sip_connection",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			Required: true,
			BodyPath: "outbound",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "Specifies whether the connection can be used.",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags associated with the connection.",
			BodyPath: "tags",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event-failover-url",
			Usage:    "The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.",
			Default:  "",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-url",
			Usage:    "The URL where webhooks related to this connection will be sent. Must include a scheme, such as 'https'.",
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a webhook.",
			Default:  nil,
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleExternalConnectionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"outbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.channel-limit",
			Usage:      "When set, this will limit the number of concurrent outbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
	},
	"inbound": {
		&requestflag.InnerFlag[string]{
			Name:       "inbound.outbound-voice-profile-id",
			Usage:      "The ID of the outbound voice profile to use for inbound calls.",
			InnerField: "outbound_voice_profile_id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.channel-limit",
			Usage:      "When set, this will limit the number of concurrent inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
	},
})

var externalConnectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Return the details of an existing External Connection inside the 'data'\nattribute of the response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleExternalConnectionsRetrieve,
	HideHelpCommand: true,
}

var externalConnectionsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates settings of an existing External Connection based on the parameters of\nthe request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			Required: true,
			BodyPath: "outbound",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "Specifies whether the connection can be used.",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags associated with the connection.",
			BodyPath: "tags",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event-failover-url",
			Usage:    "The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.",
			Default:  "",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-url",
			Usage:    "The URL where webhooks related to this connection will be sent. Must include a scheme, such as 'https'.",
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a webhook.",
			Default:  nil,
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleExternalConnectionsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"outbound": {
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.channel-limit",
			Usage:      "When set, this will limit the number of concurrent outbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
	},
	"inbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.channel-limit",
			Usage:      "When set, this will limit the number of concurrent inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
	},
})

var externalConnectionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "This endpoint returns a list of your External Connections inside the 'data'\nattribute of the response. External Connections are used by Telnyx customers to\nseamless configure SIP trunking integrations with Telnyx Partners, through\nExternal Voice Integrations in Mission Control Portal.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Filter parameter for external connections (deepObject style). Supports filtering by connection_name, external_sip_connection, id, created_at, and phone_number.",
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
	Action:          handleExternalConnectionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.id",
			Usage:      "If present, connections with <code>id</code> matching the given value will be returned.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.connection-name",
			InnerField: "connection_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.created-at",
			Usage:      "If present, connections with <code>created_at</code> date matching the given YYYY-MM-DD date will be returned.",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.external-sip-connection",
			Usage:      "If present, connections with <code>external_sip_connection</code> matching the given value will be returned.",
			InnerField: "external_sip_connection",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.phone-number",
			Usage:      "Phone number filter for connections. Note: Despite the 'contains' name, this requires a full E164 match per the original specification.",
			InnerField: "phone_number",
		},
	},
})

var externalConnectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes an External Connection. Deletion may be prevented if the\napplication is in use by phone numbers, is active, or if it is an Operator\nConnect connection. To remove an Operator Connect integration please contact\nTelnyx support.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleExternalConnectionsDelete,
	HideHelpCommand: true,
}

var externalConnectionsUpdateLocation = cli.Command{
	Name:    "update-location",
	Usage:   "Update a location's static emergency address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "location-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "static-emergency-address-id",
			Usage:    "A new static emergency address ID to update the location with",
			Required: true,
			BodyPath: "static_emergency_address_id",
		},
	},
	Action:          handleExternalConnectionsUpdateLocation,
	HideHelpCommand: true,
}

func handleExternalConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionNewParams{}

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
	_, err = client.ExternalConnections.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections create", obj, format, transform)
}

func handleExternalConnectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExternalConnections.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections retrieve", obj, format, transform)
}

func handleExternalConnectionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionUpdateParams{}

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
	_, err = client.ExternalConnections.Update(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections update", obj, format, transform)
}

func handleExternalConnectionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionListParams{}

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
		_, err = client.ExternalConnections.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "external-connections list", obj, format, transform)
	} else {
		iter := client.ExternalConnections.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "external-connections list", iter, format, transform)
	}
}

func handleExternalConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExternalConnections.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections delete", obj, format, transform)
}

func handleExternalConnectionsUpdateLocation(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("location-id") && len(unusedArgs) > 0 {
		cmd.Set("location-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionUpdateLocationParams{
		ID: cmd.Value("id").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExternalConnections.UpdateLocation(
		ctx,
		cmd.Value("location-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections update-location", obj, format, transform)
}

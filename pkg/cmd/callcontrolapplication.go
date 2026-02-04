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

var callControlApplicationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a call control application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "application-name",
			Usage:    "A user-assigned name to help manage the application.",
			Required: true,
			BodyPath: "application_name",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-url",
			Usage:    "The URL where webhooks related to this connection will be sent. Must include a scheme, such as 'https'.",
			Required: true,
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "Specifies whether the connection can be used.",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "anchorsite-override",
			Usage:    "<code>Latency</code> directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.\n",
			Default:  "Latency",
			BodyPath: "anchorsite_override",
		},
		&requestflag.Flag[bool]{
			Name:     "call-cost-in-webhooks",
			Usage:    "Specifies if call cost webhooks should be sent for this Call Control Application.",
			Default:  false,
			BodyPath: "call_cost_in_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "dtmf-type",
			Usage:    "Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats.",
			Default:  "RFC 2833",
			BodyPath: "dtmf_type",
		},
		&requestflag.Flag[bool]{
			Name:     "first-command-timeout",
			Usage:    "Specifies whether calls to phone numbers associated with this connection should hangup after timing out.",
			Default:  false,
			BodyPath: "first_command_timeout",
		},
		&requestflag.Flag[int64]{
			Name:     "first-command-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a dial command.",
			Default:  30,
			BodyPath: "first_command_timeout_secs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[bool]{
			Name:     "redact-dtmf-debug-logging",
			Usage:    "When enabled, DTMF digits entered by users will be redacted in debug logs to protect PII data entered through IVR interactions.",
			Default:  false,
			BodyPath: "redact_dtmf_debug_logging",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-api-version",
			Usage:    "Determines which webhook format will be used, Telnyx API v1 or v2.",
			Default:  "1",
			BodyPath: "webhook_api_version",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event-failover-url",
			Usage:    "The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.",
			Default:  "",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a webhook.",
			Default:  nil,
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleCallControlApplicationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.channel-limit",
			Usage:      "When set, this will limit the total number of inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.shaken-stir-enabled",
			Usage:      "When enabled Telnyx will include Shaken/Stir data in the Webhook for new inbound calls.",
			InnerField: "shaken_stir_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.sip-subdomain",
			Usage:      `Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls.`,
			InnerField: "sip_subdomain",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.sip-subdomain-receive-settings",
			Usage:      `This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in the public Internet) or "Only my connections" (any connection assigned to the same Telnyx user).`,
			InnerField: "sip_subdomain_receive_settings",
		},
	},
	"outbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.channel-limit",
			Usage:      "When set, this will limit the total number of outbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var callControlApplicationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing call control application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleCallControlApplicationsRetrieve,
	HideHelpCommand: true,
}

var callControlApplicationsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates settings of an existing call control application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "application-name",
			Usage:    "A user-assigned name to help manage the application.",
			Required: true,
			BodyPath: "application_name",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-url",
			Usage:    "The URL where webhooks related to this connection will be sent. Must include a scheme, such as 'https'.",
			Required: true,
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "Specifies whether the connection can be used.",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "anchorsite-override",
			Usage:    "<code>Latency</code> directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.\n",
			Default:  "Latency",
			BodyPath: "anchorsite_override",
		},
		&requestflag.Flag[bool]{
			Name:     "call-cost-in-webhooks",
			Usage:    "Specifies if call cost webhooks should be sent for this Call Control Application.",
			Default:  false,
			BodyPath: "call_cost_in_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "dtmf-type",
			Usage:    "Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats.",
			Default:  "RFC 2833",
			BodyPath: "dtmf_type",
		},
		&requestflag.Flag[bool]{
			Name:     "first-command-timeout",
			Usage:    "Specifies whether calls to phone numbers associated with this connection should hangup after timing out.",
			Default:  false,
			BodyPath: "first_command_timeout",
		},
		&requestflag.Flag[int64]{
			Name:     "first-command-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a dial command.",
			Default:  30,
			BodyPath: "first_command_timeout_secs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[bool]{
			Name:     "redact-dtmf-debug-logging",
			Usage:    "When enabled, DTMF digits entered by users will be redacted in debug logs to protect PII data entered through IVR interactions.",
			Default:  false,
			BodyPath: "redact_dtmf_debug_logging",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags assigned to the Call Control Application.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-api-version",
			Usage:    "Determines which webhook format will be used, Telnyx API v1 or v2.",
			Default:  "1",
			BodyPath: "webhook_api_version",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event-failover-url",
			Usage:    "The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.",
			Default:  "",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a webhook.",
			Default:  nil,
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleCallControlApplicationsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.channel-limit",
			Usage:      "When set, this will limit the total number of inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.shaken-stir-enabled",
			Usage:      "When enabled Telnyx will include Shaken/Stir data in the Webhook for new inbound calls.",
			InnerField: "shaken_stir_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.sip-subdomain",
			Usage:      `Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls.`,
			InnerField: "sip_subdomain",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.sip-subdomain-receive-settings",
			Usage:      `This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in the public Internet) or "Only my connections" (any connection assigned to the same Telnyx user).`,
			InnerField: "sip_subdomain_receive_settings",
		},
	},
	"outbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.channel-limit",
			Usage:      "When set, this will limit the total number of outbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var callControlApplicationsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Return a list of call control applications.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[application_name][contains], filter[outbound.outbound_voice_profile_id], filter[leg_id], filter[application_session_id], filter[connection_id], filter[product], filter[failed], filter[from], filter[to], filter[name], filter[type], filter[occurred_at][eq/gt/gte/lt/lte], filter[status]",
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
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/>\nThat is: <ul>\n  <li>\n    <code>connection_name</code>: sorts the result by the\n    <code>connection_name</code> field in ascending order.\n  </li>\n\n  <li>\n    <code>-connection_name</code>: sorts the result by the\n    <code>connection_name</code> field in descending order.\n  </li>\n</ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.",
			Default:   "created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleCallControlApplicationsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.application-name",
			Usage:      "Application name filters",
			InnerField: "application_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.application-session-id",
			Usage:      "The unique identifier of the call session. A session may include multiple call leg events.",
			InnerField: "application_session_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.connection-id",
			Usage:      "The unique identifier of the conection.",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.failed",
			Usage:      "Delivery failed or not.",
			InnerField: "failed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.from",
			Usage:      "Filter by From number.",
			InnerField: "from",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.leg-id",
			Usage:      "The unique identifier of an individual call leg.",
			InnerField: "leg_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.name",
			Usage:      "If present, conferences will be filtered to those with a matching `name` attribute. Matching is case-sensitive",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.occurred-at",
			Usage:      "Event occurred_at filters",
			InnerField: "occurred_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.outbound-outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound.outbound_voice_profile_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.product",
			Usage:      "Filter by product.",
			InnerField: "product",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "If present, conferences will be filtered by status.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.to",
			Usage:      "Filter by To number.",
			InnerField: "to",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			Usage:      "Event type",
			InnerField: "type",
		},
	},
})

var callControlApplicationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a call control application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleCallControlApplicationsDelete,
	HideHelpCommand: true,
}

func handleCallControlApplicationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallControlApplicationNewParams{}

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
	_, err = client.CallControlApplications.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "call-control-applications create", obj, format, transform)
}

func handleCallControlApplicationsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.CallControlApplications.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "call-control-applications retrieve", obj, format, transform)
}

func handleCallControlApplicationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallControlApplicationUpdateParams{}

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
	_, err = client.CallControlApplications.Update(
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
	return ShowJSON(os.Stdout, "call-control-applications update", obj, format, transform)
}

func handleCallControlApplicationsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CallControlApplicationListParams{}

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
		_, err = client.CallControlApplications.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "call-control-applications list", obj, format, transform)
	} else {
		iter := client.CallControlApplications.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "call-control-applications list", iter, format, transform)
	}
}

func handleCallControlApplicationsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.CallControlApplications.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "call-control-applications delete", obj, format, transform)
}

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

var texmlApplicationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a TeXML Application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "friendly-name",
			Usage:    "A user-assigned name to help manage the application.",
			Required: true,
			BodyPath: "friendly_name",
		},
		&requestflag.Flag[string]{
			Name:     "voice-url",
			Usage:    "URL to which Telnyx will deliver your XML Translator webhooks.",
			Required: true,
			BodyPath: "voice_url",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "Specifies whether the connection can be used.",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "anchorsite-override",
			Usage:    "`Latency` directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.",
			Default:  "Latency",
			BodyPath: "anchorsite_override",
		},
		&requestflag.Flag[bool]{
			Name:     "call-cost-in-webhooks",
			Usage:    "Specifies if call cost webhooks should be sent for this TeXML Application.",
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
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "URL for Telnyx to send requests to containing information about call progress events.",
			BodyPath: "status_callback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request method Telnyx should use when requesting the status_callback URL.",
			Default:  "post",
			BodyPath: "status_callback_method",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags associated with the Texml Application.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "voice-fallback-url",
			Usage:    "URL to which Telnyx will deliver your XML Translator webhooks if we get an error response from your voice_url.",
			BodyPath: "voice_fallback_url",
		},
		&requestflag.Flag[string]{
			Name:     "voice-method",
			Usage:    "HTTP request method Telnyx will use to interact with your XML Translator webhooks. Either 'get' or 'post'.",
			Default:  "post",
			BodyPath: "voice_method",
		},
	},
	Action:          handleTexmlApplicationsCreate,
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

var texmlApplicationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing TeXML Application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleTexmlApplicationsRetrieve,
	HideHelpCommand: true,
}

var texmlApplicationsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates settings of an existing TeXML Application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "friendly-name",
			Usage:    "A user-assigned name to help manage the application.",
			Required: true,
			BodyPath: "friendly_name",
		},
		&requestflag.Flag[string]{
			Name:     "voice-url",
			Usage:    "URL to which Telnyx will deliver your XML Translator webhooks.",
			Required: true,
			BodyPath: "voice_url",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "Specifies whether the connection can be used.",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "anchorsite-override",
			Usage:    "`Latency` directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.",
			Default:  "Latency",
			BodyPath: "anchorsite_override",
		},
		&requestflag.Flag[bool]{
			Name:     "call-cost-in-webhooks",
			Usage:    "Specifies if call cost webhooks should be sent for this TeXML Application.",
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
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "URL for Telnyx to send requests to containing information about call progress events.",
			BodyPath: "status_callback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request method Telnyx should use when requesting the status_callback URL.",
			Default:  "post",
			BodyPath: "status_callback_method",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags associated with the Texml Application.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "voice-fallback-url",
			Usage:    "URL to which Telnyx will deliver your XML Translator webhooks if we get an error response from your voice_url.",
			BodyPath: "voice_fallback_url",
		},
		&requestflag.Flag[string]{
			Name:     "voice-method",
			Usage:    "HTTP request method Telnyx will use to interact with your XML Translator webhooks. Either 'get' or 'post'.",
			Default:  "post",
			BodyPath: "voice_method",
		},
	},
	Action:          handleTexmlApplicationsUpdate,
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

var texmlApplicationsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your TeXML Applications.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[outbound_voice_profile_id], filter[friendly_name]",
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
			Usage:     "Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/>\nThat is: <ul>\n  <li>\n    <code>friendly_name</code>: sorts the result by the\n    <code>friendly_name</code> field in ascending order.\n  </li>\n\n  <li>\n    <code>-friendly_name</code>: sorts the result by the\n    <code>friendly_name</code> field in descending order.\n  </li>\n</ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.",
			Default:   "created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleTexmlApplicationsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.friendly-name",
			Usage:      "If present, applications with <code>friendly_name</code> containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters.",
			InnerField: "friendly_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var texmlApplicationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a TeXML Application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleTexmlApplicationsDelete,
	HideHelpCommand: true,
}

func handleTexmlApplicationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlApplicationNewParams{}

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
	_, err = client.TexmlApplications.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml-applications create", obj, format, transform)
}

func handleTexmlApplicationsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TexmlApplications.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml-applications retrieve", obj, format, transform)
}

func handleTexmlApplicationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlApplicationUpdateParams{}

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
	_, err = client.TexmlApplications.Update(
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
	return ShowJSON(os.Stdout, "texml-applications update", obj, format, transform)
}

func handleTexmlApplicationsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlApplicationListParams{}

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
		_, err = client.TexmlApplications.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "texml-applications list", obj, format, transform)
	} else {
		iter := client.TexmlApplications.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "texml-applications list", iter, format, transform)
	}
}

func handleTexmlApplicationsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TexmlApplications.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml-applications delete", obj, format, transform)
}

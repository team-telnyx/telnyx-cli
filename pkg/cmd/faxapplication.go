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

var faxApplicationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new Fax Application based on the parameters sent in the request. The\napplication name and webhook URL are required. Once created, you can assign\nphone numbers to your application using the `/phone_numbers` endpoint.",
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
			Usage:    "`Latency` directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.",
			Default:  "Latency",
			BodyPath: "anchorsite_override",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags associated with the Fax Application.",
			Default:  []string{},
			BodyPath: "tags",
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
	Action:          handleFaxApplicationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.channel-limit",
			Usage:      "When set, this will limit the number of concurrent inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
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
			Usage:      "When set, this will limit the number of concurrent outbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var faxApplicationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Return the details of an existing Fax Application inside the 'data' attribute of\nthe response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFaxApplicationsRetrieve,
	HideHelpCommand: true,
}

var faxApplicationsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates settings of an existing Fax Application based on the parameters of the\nrequest.",
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
			Usage:    "`Latency` directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.",
			Default:  "Latency",
			BodyPath: "anchorsite_override",
		},
		&requestflag.Flag[any]{
			Name:     "fax-email-recipient",
			Usage:    "Specifies an email address where faxes sent to this application will be forwarded to (as pdf or tiff attachments)",
			Default:  nil,
			BodyPath: "fax_email_recipient",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags associated with the Fax Application.",
			BodyPath: "tags",
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
	Action:          handleFaxApplicationsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.channel-limit",
			Usage:      "When set, this will limit the number of concurrent inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
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
			Usage:      "When set, this will limit the number of concurrent outbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var faxApplicationsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "This endpoint returns a list of your Fax Applications inside the 'data'\nattribute of the response. You can adjust which applications are listed by using\nfilters. Fax Applications are used to configure how you send and receive faxes\nusing the Programmable Fax API with Telnyx.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[application_name][contains], filter[outbound_voice_profile_id]",
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
			Usage:     "Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/>\nThat is: <ul>\n  <li>\n    <code>application_name</code>: sorts the result by the\n    <code>application_name</code> field in ascending order.\n  </li>\n\n  <li>\n    <code>-application_name</code>: sorts the result by the\n    <code>application_name</code> field in descending order.\n  </li>\n</ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.",
			Default:   "created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleFaxApplicationsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.application-name",
			Usage:      "Application name filtering operations",
			InnerField: "application_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var faxApplicationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a Fax Application. Deletion may be prevented if the\napplication is in use by phone numbers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFaxApplicationsDelete,
	HideHelpCommand: true,
}

func handleFaxApplicationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FaxApplicationNewParams{}

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
	_, err = client.FaxApplications.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fax-applications create", obj, format, transform)
}

func handleFaxApplicationsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.FaxApplications.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fax-applications retrieve", obj, format, transform)
}

func handleFaxApplicationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FaxApplicationUpdateParams{}

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
	_, err = client.FaxApplications.Update(
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
	return ShowJSON(os.Stdout, "fax-applications update", obj, format, transform)
}

func handleFaxApplicationsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FaxApplicationListParams{}

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
		_, err = client.FaxApplications.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "fax-applications list", obj, format, transform)
	} else {
		iter := client.FaxApplications.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "fax-applications list", iter, format, transform)
	}
}

func handleFaxApplicationsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.FaxApplications.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fax-applications delete", obj, format, transform)
}

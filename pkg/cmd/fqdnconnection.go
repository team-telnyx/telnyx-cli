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

var fqdnConnectionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a FQDN connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-name",
			Usage:    "A user-assigned name to help manage the connection.",
			Required: true,
			BodyPath: "connection_name",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "Defaults to true",
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
			Name:     "android-push-credential-id",
			Usage:    "The uuid of the push credential for Android",
			Default:  nil,
			BodyPath: "android_push_credential_id",
		},
		&requestflag.Flag[bool]{
			Name:     "call-cost-in-webhooks",
			Usage:    "Specifies if call cost webhooks should be sent for this connection.",
			Default:  false,
			BodyPath: "call_cost_in_webhooks",
		},
		&requestflag.Flag[bool]{
			Name:     "default-on-hold-comfort-noise-enabled",
			Usage:    "When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout.",
			Default:  true,
			BodyPath: "default_on_hold_comfort_noise_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "dtmf-type",
			Usage:    "Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats.",
			Default:  "RFC 2833",
			BodyPath: "dtmf_type",
		},
		&requestflag.Flag[bool]{
			Name:     "encode-contact-header-enabled",
			Usage:    "Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios.",
			Default:  false,
			BodyPath: "encode_contact_header_enabled",
		},
		&requestflag.Flag[any]{
			Name:     "encrypted-media",
			Usage:    "Enable use of SRTP for encryption. Cannot be set if the transport_portocol is TLS.",
			BodyPath: "encrypted_media",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[any]{
			Name:     "ios-push-credential-id",
			Usage:    "The uuid of the push credential for Ios",
			Default:  nil,
			BodyPath: "ios_push_credential_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "jitter-buffer",
			Usage:    "Configuration options for Jitter Buffer. Enables Jitter Buffer for RTP streams of SIP Trunking calls. The feature is off unless enabled. You may define min and max values in msec for customized buffering behaviors. Larger values add latency but tolerate more jitter, while smaller values reduce latency but are more sensitive to jitter and reordering.",
			BodyPath: "jitter_buffer",
		},
		&requestflag.Flag[bool]{
			Name:     "microsoft-teams-sbc",
			Usage:    "When enabled, the connection will be created for Microsoft Teams Direct Routing. A *.mstsbc.telnyx.tech FQDN will be created for the connection automatically.",
			Default:  false,
			BodyPath: "microsoft_teams_sbc",
		},
		&requestflag.Flag[string]{
			Name:     "noise-suppression",
			Usage:    "Controls when noise suppression is applied to calls. When set to 'inbound', noise suppression is applied to incoming audio. When set to 'outbound', it's applied to outgoing audio. When set to 'both', it's applied in both directions. When set to 'disabled', noise suppression is turned off.",
			BodyPath: "noise_suppression",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "noise-suppression-details",
			Usage:    "Configuration options for noise suppression. These settings are stored regardless of the noise_suppression value, but only take effect when noise_suppression is not 'disabled'. If you disable noise suppression and later re-enable it, the previously configured settings will be used.",
			BodyPath: "noise_suppression_details",
		},
		&requestflag.Flag[bool]{
			Name:     "onnet-t38-passthrough-enabled",
			Usage:    "Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg's settings.",
			Default:  false,
			BodyPath: "onnet_t38_passthrough_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "rtcp-settings",
			BodyPath: "rtcp_settings",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags associated with the connection.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "transport-protocol",
			Usage:    "One of UDP, TLS, or TCP. Applies only to connections with IP authentication or FQDN authentication.",
			Default:  "UDP",
			BodyPath: "transport_protocol",
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
	Action:          handleFqdnConnectionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound": {
		&requestflag.InnerFlag[string]{
			Name:       "inbound.ani-number-format",
			Usage:      "This setting allows you to set the format with which the caller's number (ANI) is sent for inbound phone calls.",
			InnerField: "ani_number_format",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.channel-limit",
			Usage:      "When set, this will limit the total number of inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "inbound.codecs",
			Usage:      "Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP.",
			InnerField: "codecs",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.default-primary-fqdn-id",
			Usage:      "The default primary FQDN to use for the number. Only settable if the connection is\nof FQDN type. Value must be the ID of an FQDN set on the connection.",
			InnerField: "default_primary_fqdn_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.default-routing-method",
			Usage:      "Default routing method to be used when a number is associated with the connection. Must be one of the routing method types or null, other values are not allowed.",
			InnerField: "default_routing_method",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.default-secondary-fqdn-id",
			Usage:      "The default secondary FQDN to use for the number. Only settable if the connection is\nof FQDN type. Value must be the ID of an FQDN set on the connection.",
			InnerField: "default_secondary_fqdn_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.default-tertiary-fqdn-id",
			Usage:      "The default tertiary FQDN to use for the number. Only settable if the connection is\nof FQDN type. Value must be the ID of an FQDN set on the connection.",
			InnerField: "default_tertiary_fqdn_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.dnis-number-format",
			InnerField: "dnis_number_format",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.generate-ringback-tone",
			Usage:      "Generate ringback tone through 183 session progress message with early media.",
			InnerField: "generate_ringback_tone",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.isup-headers-enabled",
			Usage:      "When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.)",
			InnerField: "isup_headers_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.prack-enabled",
			Usage:      "Enable PRACK messages as defined in RFC3262.",
			InnerField: "prack_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.shaken-stir-enabled",
			Usage:      "When enabled the SIP Connection will receive the Identity header with Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP transport.",
			InnerField: "shaken_stir_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.sip-compact-headers-enabled",
			Usage:      "Defaults to true.",
			InnerField: "sip_compact_headers_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.sip-region",
			Usage:      "Selects which `sip_region` to receive inbound calls from. If null, the default region (US) will be used.",
			InnerField: "sip_region",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.sip-subdomain",
			Usage:      `Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls.`,
			InnerField: "sip_subdomain",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.sip-subdomain-receive-settings",
			Usage:      `This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in the public Internet) or "Only my connections" (any connection assigned to the same Telnyx user).`,
			InnerField: "sip_subdomain_receive_settings",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.timeout-1xx-secs",
			Usage:      "Time(sec) before aborting if connection is not made.",
			InnerField: "timeout_1xx_secs",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.timeout-2xx-secs",
			Usage:      "Time(sec) before aborting if call is unanswered (min: 1, max: 600).",
			InnerField: "timeout_2xx_secs",
		},
	},
	"jitter-buffer": {
		&requestflag.InnerFlag[bool]{
			Name:       "jitter-buffer.enable-jitter-buffer",
			Usage:      "Enables Jitter Buffer for RTP streams of SIP Trunking calls. The feature is off unless enabled.",
			InnerField: "enable_jitter_buffer",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "jitter-buffer.jitterbuffer-msec-max",
			Usage:      "The maximum jitter buffer size in milliseconds. Must be between 40 and 400. Has no effect if enable_jitter_buffer is not true.",
			InnerField: "jitterbuffer_msec_max",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "jitter-buffer.jitterbuffer-msec-min",
			Usage:      "The minimum jitter buffer size in milliseconds. Must be between 40 and 400. Has no effect if enable_jitter_buffer is not true.",
			InnerField: "jitterbuffer_msec_min",
		},
	},
	"noise-suppression-details": {
		&requestflag.InnerFlag[int64]{
			Name:       "noise-suppression-details.attenuation-limit",
			Usage:      "The attenuation limit value for the selected engine. Default values vary by engine: 0 for 'denoiser', 80 for 'deep_filter_net', 'deep_filter_net_large', and all Krisp engines ('krisp_viva_tel', 'krisp_viva_tel_lite', 'krisp_viva_promodel', 'krisp_viva_ss').",
			InnerField: "attenuation_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "noise-suppression-details.engine",
			Usage:      "The noise suppression engine to use. 'denoiser' is the default engine. 'deep_filter_net' and 'deep_filter_net_large' are alternative engines with different performance characteristics. Krisp engines ('krisp_viva_tel', 'krisp_viva_tel_lite', 'krisp_viva_promodel', 'krisp_viva_ss') provide advanced noise suppression capabilities.",
			InnerField: "engine",
		},
	},
	"outbound": {
		&requestflag.InnerFlag[string]{
			Name:       "outbound.ani-override",
			Usage:      "Set a phone number as the ani_override value to override caller id number on outbound calls.",
			InnerField: "ani_override",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.ani-override-type",
			Usage:      "Specifies when we should apply your ani_override setting. Only applies when ani_override is not blank.",
			InnerField: "ani_override_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "outbound.call-parking-enabled",
			Usage:      `Forces all SIP calls originated on this connection to be \"parked\" instead of \"bridged\" to the destination specified on the URI. Parked calls will return ringback to the caller and will await for a Call Control command to define which action will be taken next.`,
			InnerField: "call_parking_enabled",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.channel-limit",
			Usage:      "When set, this will limit the total number of inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[any]{
			Name:       "outbound.encrypted-media",
			Usage:      "Enable use of SRTP for encryption. Cannot be set if the transport_portocol is TLS.",
			InnerField: "encrypted_media",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "outbound.generate-ringback-tone",
			Usage:      "Generate ringback tone through 183 session progress message with early media.",
			InnerField: "generate_ringback_tone",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "outbound.instant-ringback-enabled",
			Usage:      "When set, ringback will not wait for indication before sending ringback tone to calling party.",
			InnerField: "instant_ringback_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.ip-authentication-method",
			InnerField: "ip_authentication_method",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.ip-authentication-token",
			InnerField: "ip_authentication_token",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.localization",
			Usage:      "A 2-character country code specifying the country whose national dialing rules should be used. For example, if set to `US` then any US number can be dialed without preprending +1 to the number. When left blank, Telnyx will try US and GB dialing rules, in that order, by default.\",",
			InnerField: "localization",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.t38-reinvite-source",
			Usage:      "This setting only affects connections with Fax-type Outbound Voice Profiles. The setting dictates whether or not Telnyx sends a t.38 reinvite. By default, Telnyx will send the re-invite. If set to `customer`, the caller is expected to send the t.38 reinvite.",
			InnerField: "t38_reinvite_source",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.tech-prefix",
			Usage:      "Numerical chars only, exactly 4 characters.",
			InnerField: "tech_prefix",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.timeout-1xx-secs",
			Usage:      "Time(sec) before aborting if connection is not made.",
			InnerField: "timeout_1xx_secs",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.timeout-2xx-secs",
			Usage:      "Time(sec) before aborting if call is unanswered (min: 1, max: 600).",
			InnerField: "timeout_2xx_secs",
		},
	},
	"rtcp-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "rtcp-settings.capture-enabled",
			Usage:      "BETA - Enable the capture and storage of RTCP messages to create QoS reports on the Telnyx Mission Control Portal.",
			InnerField: "capture_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rtcp-settings.port",
			Usage:      "RTCP port by default is rtp+1, it can also be set to rtcp-mux",
			InnerField: "port",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "rtcp-settings.report-frequency-secs",
			Usage:      "RTCP reports are sent to customers based on the frequency set. Frequency is in seconds and it can be set to values from 5 to 3000 seconds.",
			InnerField: "report_frequency_secs",
		},
	},
})

var fqdnConnectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing FQDN connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFqdnConnectionsRetrieve,
	HideHelpCommand: true,
}

var fqdnConnectionsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates settings of an existing FQDN connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "Defaults to true",
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "anchorsite-override",
			Usage:    "`Latency` directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.",
			Default:  "Latency",
			BodyPath: "anchorsite_override",
		},
		&requestflag.Flag[any]{
			Name:     "android-push-credential-id",
			Usage:    "The uuid of the push credential for Android",
			Default:  nil,
			BodyPath: "android_push_credential_id",
		},
		&requestflag.Flag[bool]{
			Name:     "call-cost-in-webhooks",
			Usage:    "Specifies if call cost webhooks should be sent for this connection.",
			Default:  false,
			BodyPath: "call_cost_in_webhooks",
		},
		&requestflag.Flag[string]{
			Name:     "connection-name",
			Usage:    "A user-assigned name to help manage the connection.",
			BodyPath: "connection_name",
		},
		&requestflag.Flag[bool]{
			Name:     "default-on-hold-comfort-noise-enabled",
			Usage:    "When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout.",
			Default:  true,
			BodyPath: "default_on_hold_comfort_noise_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "dtmf-type",
			Usage:    "Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats.",
			Default:  "RFC 2833",
			BodyPath: "dtmf_type",
		},
		&requestflag.Flag[bool]{
			Name:     "encode-contact-header-enabled",
			Usage:    "Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios.",
			Default:  false,
			BodyPath: "encode_contact_header_enabled",
		},
		&requestflag.Flag[any]{
			Name:     "encrypted-media",
			Usage:    "Enable use of SRTP for encryption. Cannot be set if the transport_portocol is TLS.",
			BodyPath: "encrypted_media",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[any]{
			Name:     "ios-push-credential-id",
			Usage:    "The uuid of the push credential for Ios",
			Default:  nil,
			BodyPath: "ios_push_credential_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "jitter-buffer",
			Usage:    "Configuration options for Jitter Buffer. Enables Jitter Buffer for RTP streams of SIP Trunking calls. The feature is off unless enabled. You may define min and max values in msec for customized buffering behaviors. Larger values add latency but tolerate more jitter, while smaller values reduce latency but are more sensitive to jitter and reordering.",
			BodyPath: "jitter_buffer",
		},
		&requestflag.Flag[string]{
			Name:     "noise-suppression",
			Usage:    "Controls when noise suppression is applied to calls. When set to 'inbound', noise suppression is applied to incoming audio. When set to 'outbound', it's applied to outgoing audio. When set to 'both', it's applied in both directions. When set to 'disabled', noise suppression is turned off.",
			BodyPath: "noise_suppression",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "noise-suppression-details",
			Usage:    "Configuration options for noise suppression. These settings are stored regardless of the noise_suppression value, but only take effect when noise_suppression is not 'disabled'. If you disable noise suppression and later re-enable it, the previously configured settings will be used.",
			BodyPath: "noise_suppression_details",
		},
		&requestflag.Flag[bool]{
			Name:     "onnet-t38-passthrough-enabled",
			Usage:    "Enable on-net T38 if you prefer that the sender and receiver negotiate T38 directly when both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call according to each leg's settings.",
			Default:  false,
			BodyPath: "onnet_t38_passthrough_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "rtcp-settings",
			BodyPath: "rtcp_settings",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags associated with the connection.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "transport-protocol",
			Usage:    "One of UDP, TLS, or TCP. Applies only to connections with IP authentication or FQDN authentication.",
			Default:  "UDP",
			BodyPath: "transport_protocol",
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
	Action:          handleFqdnConnectionsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound": {
		&requestflag.InnerFlag[string]{
			Name:       "inbound.ani-number-format",
			Usage:      "This setting allows you to set the format with which the caller's number (ANI) is sent for inbound phone calls.",
			InnerField: "ani_number_format",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.channel-limit",
			Usage:      "When set, this will limit the total number of inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "inbound.codecs",
			Usage:      "Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP.",
			InnerField: "codecs",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.default-primary-fqdn-id",
			Usage:      "The default primary FQDN to use for the number. Only settable if the connection is\nof FQDN type. Value must be the ID of an FQDN set on the connection.",
			InnerField: "default_primary_fqdn_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.default-routing-method",
			Usage:      "Default routing method to be used when a number is associated with the connection. Must be one of the routing method types or null, other values are not allowed.",
			InnerField: "default_routing_method",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.default-secondary-fqdn-id",
			Usage:      "The default secondary FQDN to use for the number. Only settable if the connection is\nof FQDN type. Value must be the ID of an FQDN set on the connection.",
			InnerField: "default_secondary_fqdn_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.default-tertiary-fqdn-id",
			Usage:      "The default tertiary FQDN to use for the number. Only settable if the connection is\nof FQDN type. Value must be the ID of an FQDN set on the connection.",
			InnerField: "default_tertiary_fqdn_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.dnis-number-format",
			InnerField: "dnis_number_format",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.generate-ringback-tone",
			Usage:      "Generate ringback tone through 183 session progress message with early media.",
			InnerField: "generate_ringback_tone",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.isup-headers-enabled",
			Usage:      "When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.)",
			InnerField: "isup_headers_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.prack-enabled",
			Usage:      "Enable PRACK messages as defined in RFC3262.",
			InnerField: "prack_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.shaken-stir-enabled",
			Usage:      "When enabled the SIP Connection will receive the Identity header with Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP transport.",
			InnerField: "shaken_stir_enabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "inbound.sip-compact-headers-enabled",
			Usage:      "Defaults to true.",
			InnerField: "sip_compact_headers_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.sip-region",
			Usage:      "Selects which `sip_region` to receive inbound calls from. If null, the default region (US) will be used.",
			InnerField: "sip_region",
		},
		&requestflag.InnerFlag[any]{
			Name:       "inbound.sip-subdomain",
			Usage:      `Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain "example.sip.telnyx.com" can be called from any SIP endpoint by using the SIP URI "sip:@example.sip.telnyx.com" where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls.`,
			InnerField: "sip_subdomain",
		},
		&requestflag.InnerFlag[string]{
			Name:       "inbound.sip-subdomain-receive-settings",
			Usage:      `This option can be enabled to receive calls from: "Anyone" (any SIP endpoint in the public Internet) or "Only my connections" (any connection assigned to the same Telnyx user).`,
			InnerField: "sip_subdomain_receive_settings",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.timeout-1xx-secs",
			Usage:      "Time(sec) before aborting if connection is not made.",
			InnerField: "timeout_1xx_secs",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.timeout-2xx-secs",
			Usage:      "Time(sec) before aborting if call is unanswered (min: 1, max: 600).",
			InnerField: "timeout_2xx_secs",
		},
	},
	"jitter-buffer": {
		&requestflag.InnerFlag[bool]{
			Name:       "jitter-buffer.enable-jitter-buffer",
			Usage:      "Enables Jitter Buffer for RTP streams of SIP Trunking calls. The feature is off unless enabled.",
			InnerField: "enable_jitter_buffer",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "jitter-buffer.jitterbuffer-msec-max",
			Usage:      "The maximum jitter buffer size in milliseconds. Must be between 40 and 400. Has no effect if enable_jitter_buffer is not true.",
			InnerField: "jitterbuffer_msec_max",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "jitter-buffer.jitterbuffer-msec-min",
			Usage:      "The minimum jitter buffer size in milliseconds. Must be between 40 and 400. Has no effect if enable_jitter_buffer is not true.",
			InnerField: "jitterbuffer_msec_min",
		},
	},
	"noise-suppression-details": {
		&requestflag.InnerFlag[int64]{
			Name:       "noise-suppression-details.attenuation-limit",
			Usage:      "The attenuation limit value for the selected engine. Default values vary by engine: 0 for 'denoiser', 80 for 'deep_filter_net', 'deep_filter_net_large', and all Krisp engines ('krisp_viva_tel', 'krisp_viva_tel_lite', 'krisp_viva_promodel', 'krisp_viva_ss').",
			InnerField: "attenuation_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "noise-suppression-details.engine",
			Usage:      "The noise suppression engine to use. 'denoiser' is the default engine. 'deep_filter_net' and 'deep_filter_net_large' are alternative engines with different performance characteristics. Krisp engines ('krisp_viva_tel', 'krisp_viva_tel_lite', 'krisp_viva_promodel', 'krisp_viva_ss') provide advanced noise suppression capabilities.",
			InnerField: "engine",
		},
	},
	"outbound": {
		&requestflag.InnerFlag[string]{
			Name:       "outbound.ani-override",
			Usage:      "Set a phone number as the ani_override value to override caller id number on outbound calls.",
			InnerField: "ani_override",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.ani-override-type",
			Usage:      "Specifies when we should apply your ani_override setting. Only applies when ani_override is not blank.",
			InnerField: "ani_override_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "outbound.call-parking-enabled",
			Usage:      `Forces all SIP calls originated on this connection to be \"parked\" instead of \"bridged\" to the destination specified on the URI. Parked calls will return ringback to the caller and will await for a Call Control command to define which action will be taken next.`,
			InnerField: "call_parking_enabled",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.channel-limit",
			Usage:      "When set, this will limit the total number of inbound calls to phone numbers associated with this connection.",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[any]{
			Name:       "outbound.encrypted-media",
			Usage:      "Enable use of SRTP for encryption. Cannot be set if the transport_portocol is TLS.",
			InnerField: "encrypted_media",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "outbound.generate-ringback-tone",
			Usage:      "Generate ringback tone through 183 session progress message with early media.",
			InnerField: "generate_ringback_tone",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "outbound.instant-ringback-enabled",
			Usage:      "When set, ringback will not wait for indication before sending ringback tone to calling party.",
			InnerField: "instant_ringback_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.ip-authentication-method",
			InnerField: "ip_authentication_method",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.ip-authentication-token",
			InnerField: "ip_authentication_token",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.localization",
			Usage:      "A 2-character country code specifying the country whose national dialing rules should be used. For example, if set to `US` then any US number can be dialed without preprending +1 to the number. When left blank, Telnyx will try US and GB dialing rules, in that order, by default.\",",
			InnerField: "localization",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.t38-reinvite-source",
			Usage:      "This setting only affects connections with Fax-type Outbound Voice Profiles. The setting dictates whether or not Telnyx sends a t.38 reinvite. By default, Telnyx will send the re-invite. If set to `customer`, the caller is expected to send the t.38 reinvite.",
			InnerField: "t38_reinvite_source",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.tech-prefix",
			Usage:      "Numerical chars only, exactly 4 characters.",
			InnerField: "tech_prefix",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.timeout-1xx-secs",
			Usage:      "Time(sec) before aborting if connection is not made.",
			InnerField: "timeout_1xx_secs",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.timeout-2xx-secs",
			Usage:      "Time(sec) before aborting if call is unanswered (min: 1, max: 600).",
			InnerField: "timeout_2xx_secs",
		},
	},
	"rtcp-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "rtcp-settings.capture-enabled",
			Usage:      "BETA - Enable the capture and storage of RTCP messages to create QoS reports on the Telnyx Mission Control Portal.",
			InnerField: "capture_enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rtcp-settings.port",
			Usage:      "RTCP port by default is rtp+1, it can also be set to rtcp-mux",
			InnerField: "port",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "rtcp-settings.report-frequency-secs",
			Usage:      "RTCP reports are sent to customers based on the frequency set. Frequency is in seconds and it can be set to values from 5 to 3000 seconds.",
			InnerField: "report_frequency_secs",
		},
	},
})

var fqdnConnectionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your FQDN connections.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[connection_name], filter[fqdn], filter[outbound_voice_profile_id], filter[outbound.outbound_voice_profile_id]",
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
	Action:          handleFqdnConnectionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.connection-name",
			Usage:      "Filter by connection_name using nested operations",
			InnerField: "connection_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.fqdn",
			Usage:      "If present, connections with an `fqdn` that equals the given value will be returned. Matching is case-sensitive, and the full string must match.",
			InnerField: "fqdn",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var fqdnConnectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an FQDN connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFqdnConnectionsDelete,
	HideHelpCommand: true,
}

func handleFqdnConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FqdnConnectionNewParams{}

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
	_, err = client.FqdnConnections.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fqdn-connections create", obj, format, transform)
}

func handleFqdnConnectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.FqdnConnections.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fqdn-connections retrieve", obj, format, transform)
}

func handleFqdnConnectionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FqdnConnectionUpdateParams{}

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
	_, err = client.FqdnConnections.Update(
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
	return ShowJSON(os.Stdout, "fqdn-connections update", obj, format, transform)
}

func handleFqdnConnectionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FqdnConnectionListParams{}

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
		_, err = client.FqdnConnections.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "fqdn-connections list", obj, format, transform)
	} else {
		iter := client.FqdnConnections.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "fqdn-connections list", iter, format, transform)
	}
}

func handleFqdnConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.FqdnConnections.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fqdn-connections delete", obj, format, transform)
}

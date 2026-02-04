// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestFqdnConnectionsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdn-connections", "create",
		"--connection-name", "string",
		"--active=true",
		"--anchorsite-override", "Latency",
		"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
		"--call-cost-in-webhooks=false",
		"--default-on-hold-comfort-noise-enabled=true",
		"--dtmf-type", "RFC 2833",
		"--encode-contact-header-enabled=true",
		"--encrypted-media", "SRTP",
		"--inbound", "{ani_number_format: +E.164, channel_limit: 10, codecs: [G722], default_primary_fqdn_id: default_primary_fqdn_id, default_routing_method: sequential, default_secondary_fqdn_id: default_secondary_fqdn_id, default_tertiary_fqdn_id: default_tertiary_fqdn_id, dnis_number_format: +e164, generate_ringback_tone: true, isup_headers_enabled: true, prack_enabled: true, shaken_stir_enabled: true, sip_compact_headers_enabled: true, sip_region: US, sip_subdomain: string, sip_subdomain_receive_settings: only_my_connections, timeout_1xx_secs: 10, timeout_2xx_secs: 10}",
		"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
		"--jitter-buffer", "{enable_jitter_buffer: true, jitterbuffer_msec_max: 200, jitterbuffer_msec_min: 60}",
		"--microsoft-teams-sbc=true",
		"--noise-suppression", "both",
		"--noise-suppression-details", "{attenuation_limit: 80, engine: deep_filter_net}",
		"--onnet-t38-passthrough-enabled=true",
		"--outbound", "{ani_override: '+1234567890', ani_override_type: always, call_parking_enabled: true, channel_limit: 10, encrypted_media: SRTP, generate_ringback_tone: true, instant_ringback_enabled: true, ip_authentication_method: credential-authentication, ip_authentication_token: aBcD1234, localization: string, outbound_voice_profile_id: outbound_voice_profile_id, t38_reinvite_source: customer, tech_prefix: '0123', timeout_1xx_secs: 10, timeout_2xx_secs: 10}",
		"--rtcp-settings", "{capture_enabled: true, port: rtcp-mux, report_frequency_secs: 10}",
		"--tag", "tag1",
		"--tag", "tag2",
		"--transport-protocol", "UDP",
		"--webhook-api-version", "1",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(fqdnConnectionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdn-connections", "create",
		"--connection-name", "string",
		"--active=true",
		"--anchorsite-override", "Latency",
		"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
		"--call-cost-in-webhooks=false",
		"--default-on-hold-comfort-noise-enabled=true",
		"--dtmf-type", "RFC 2833",
		"--encode-contact-header-enabled=true",
		"--encrypted-media", "SRTP",
		"--inbound.ani-number-format", "+E.164",
		"--inbound.channel-limit", "10",
		"--inbound.codecs", "[G722]",
		"--inbound.default-primary-fqdn-id", "default_primary_fqdn_id",
		"--inbound.default-routing-method", "sequential",
		"--inbound.default-secondary-fqdn-id", "default_secondary_fqdn_id",
		"--inbound.default-tertiary-fqdn-id", "default_tertiary_fqdn_id",
		"--inbound.dnis-number-format", "+e164",
		"--inbound.generate-ringback-tone=true",
		"--inbound.isup-headers-enabled=true",
		"--inbound.prack-enabled=true",
		"--inbound.shaken-stir-enabled=true",
		"--inbound.sip-compact-headers-enabled=true",
		"--inbound.sip-region", "US",
		"--inbound.sip-subdomain", "string",
		"--inbound.sip-subdomain-receive-settings", "only_my_connections",
		"--inbound.timeout-1xx-secs", "10",
		"--inbound.timeout-2xx-secs", "10",
		"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
		"--jitter-buffer.enable-jitter-buffer=true",
		"--jitter-buffer.jitterbuffer-msec-max", "200",
		"--jitter-buffer.jitterbuffer-msec-min", "60",
		"--microsoft-teams-sbc=true",
		"--noise-suppression", "both",
		"--noise-suppression-details.attenuation-limit", "80",
		"--noise-suppression-details.engine", "deep_filter_net",
		"--onnet-t38-passthrough-enabled=true",
		"--outbound.ani-override", "+1234567890",
		"--outbound.ani-override-type", "always",
		"--outbound.call-parking-enabled=true",
		"--outbound.channel-limit", "10",
		"--outbound.encrypted-media", "SRTP",
		"--outbound.generate-ringback-tone=true",
		"--outbound.instant-ringback-enabled=true",
		"--outbound.ip-authentication-method", "credential-authentication",
		"--outbound.ip-authentication-token", "aBcD1234",
		"--outbound.localization", "string",
		"--outbound.outbound-voice-profile-id", "outbound_voice_profile_id",
		"--outbound.t38-reinvite-source", "customer",
		"--outbound.tech-prefix", "0123",
		"--outbound.timeout-1xx-secs", "10",
		"--outbound.timeout-2xx-secs", "10",
		"--rtcp-settings.capture-enabled=true",
		"--rtcp-settings.port", "rtcp-mux",
		"--rtcp-settings.report-frequency-secs", "10",
		"--tag", "tag1",
		"--tag", "tag2",
		"--transport-protocol", "UDP",
		"--webhook-api-version", "1",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)
}

func TestFqdnConnectionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdn-connections", "retrieve",
		"--id", "id",
	)
}

func TestFqdnConnectionsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdn-connections", "update",
		"--id", "id",
		"--active=true",
		"--anchorsite-override", "Latency",
		"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
		"--call-cost-in-webhooks=true",
		"--connection-name", "string",
		"--default-on-hold-comfort-noise-enabled=true",
		"--dtmf-type", "RFC 2833",
		"--encode-contact-header-enabled=true",
		"--encrypted-media", "SRTP",
		"--inbound", "{ani_number_format: +E.164, channel_limit: 10, codecs: [G722], default_primary_fqdn_id: default_primary_fqdn_id, default_routing_method: sequential, default_secondary_fqdn_id: default_secondary_fqdn_id, default_tertiary_fqdn_id: default_tertiary_fqdn_id, dnis_number_format: +e164, generate_ringback_tone: true, isup_headers_enabled: true, prack_enabled: true, shaken_stir_enabled: true, sip_compact_headers_enabled: true, sip_region: US, sip_subdomain: string, sip_subdomain_receive_settings: only_my_connections, timeout_1xx_secs: 10, timeout_2xx_secs: 10}",
		"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
		"--jitter-buffer", "{enable_jitter_buffer: true, jitterbuffer_msec_max: 200, jitterbuffer_msec_min: 60}",
		"--noise-suppression", "both",
		"--noise-suppression-details", "{attenuation_limit: 80, engine: deep_filter_net}",
		"--onnet-t38-passthrough-enabled=true",
		"--outbound", "{ani_override: ani_override, ani_override_type: normal, call_parking_enabled: true, channel_limit: 0, encrypted_media: SRTP, generate_ringback_tone: true, instant_ringback_enabled: true, ip_authentication_method: credential-authentication, ip_authentication_token: ip_authentication_token, localization: US, outbound_voice_profile_id: outbound_voice_profile_id, t38_reinvite_source: telnyx, tech_prefix: tech_prefix, timeout_1xx_secs: 1, timeout_2xx_secs: 1}",
		"--rtcp-settings", "{capture_enabled: true, port: rtcp-mux, report_frequency_secs: 10}",
		"--tag", "tag1",
		"--tag", "tag2",
		"--transport-protocol", "UDP",
		"--webhook-api-version", "1",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(fqdnConnectionsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdn-connections", "update",
		"--id", "id",
		"--active=true",
		"--anchorsite-override", "Latency",
		"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
		"--call-cost-in-webhooks=true",
		"--connection-name", "string",
		"--default-on-hold-comfort-noise-enabled=true",
		"--dtmf-type", "RFC 2833",
		"--encode-contact-header-enabled=true",
		"--encrypted-media", "SRTP",
		"--inbound.ani-number-format", "+E.164",
		"--inbound.channel-limit", "10",
		"--inbound.codecs", "[G722]",
		"--inbound.default-primary-fqdn-id", "default_primary_fqdn_id",
		"--inbound.default-routing-method", "sequential",
		"--inbound.default-secondary-fqdn-id", "default_secondary_fqdn_id",
		"--inbound.default-tertiary-fqdn-id", "default_tertiary_fqdn_id",
		"--inbound.dnis-number-format", "+e164",
		"--inbound.generate-ringback-tone=true",
		"--inbound.isup-headers-enabled=true",
		"--inbound.prack-enabled=true",
		"--inbound.shaken-stir-enabled=true",
		"--inbound.sip-compact-headers-enabled=true",
		"--inbound.sip-region", "US",
		"--inbound.sip-subdomain", "string",
		"--inbound.sip-subdomain-receive-settings", "only_my_connections",
		"--inbound.timeout-1xx-secs", "10",
		"--inbound.timeout-2xx-secs", "10",
		"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
		"--jitter-buffer.enable-jitter-buffer=true",
		"--jitter-buffer.jitterbuffer-msec-max", "200",
		"--jitter-buffer.jitterbuffer-msec-min", "60",
		"--noise-suppression", "both",
		"--noise-suppression-details.attenuation-limit", "80",
		"--noise-suppression-details.engine", "deep_filter_net",
		"--onnet-t38-passthrough-enabled=true",
		"--outbound.ani-override", "ani_override",
		"--outbound.ani-override-type", "normal",
		"--outbound.call-parking-enabled=true",
		"--outbound.channel-limit", "0",
		"--outbound.encrypted-media", "SRTP",
		"--outbound.generate-ringback-tone=true",
		"--outbound.instant-ringback-enabled=true",
		"--outbound.ip-authentication-method", "credential-authentication",
		"--outbound.ip-authentication-token", "ip_authentication_token",
		"--outbound.localization", "US",
		"--outbound.outbound-voice-profile-id", "outbound_voice_profile_id",
		"--outbound.t38-reinvite-source", "telnyx",
		"--outbound.tech-prefix", "tech_prefix",
		"--outbound.timeout-1xx-secs", "1",
		"--outbound.timeout-2xx-secs", "1",
		"--rtcp-settings.capture-enabled=true",
		"--rtcp-settings.port", "rtcp-mux",
		"--rtcp-settings.report-frequency-secs", "10",
		"--tag", "tag1",
		"--tag", "tag2",
		"--transport-protocol", "UDP",
		"--webhook-api-version", "1",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)
}

func TestFqdnConnectionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdn-connections", "list",
		"--filter", "{connection_name: {contains: contains}, fqdn: fqdn, outbound_voice_profile_id: outbound_voice_profile_id}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "connection_name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(fqdnConnectionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdn-connections", "list",
		"--filter.connection-name", "{contains: contains}",
		"--filter.fqdn", "fqdn",
		"--filter.outbound-voice-profile-id", "outbound_voice_profile_id",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "connection_name",
	)
}

func TestFqdnConnectionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdn-connections", "delete",
		"--id", "id",
	)
}

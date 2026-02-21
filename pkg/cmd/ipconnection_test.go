// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestIPConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ip-connections", "create",
		"--active=true",
		"--anchorsite-override", "Latency",
		"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
		"--call-cost-in-webhooks=false",
		"--connection-name", "string",
		"--default-on-hold-comfort-noise-enabled=true",
		"--dtmf-type", "RFC 2833",
		"--encode-contact-header-enabled=true",
		"--encrypted-media", "SRTP",
		"--inbound", "{ani_number_format: +E.164, channel_limit: 10, codecs: [G722], default_routing_method: sequential, dnis_number_format: +e164, generate_ringback_tone: true, isup_headers_enabled: true, prack_enabled: true, shaken_stir_enabled: true, sip_compact_headers_enabled: true, sip_region: US, sip_subdomain: test, sip_subdomain_receive_settings: only_my_connections, timeout_1xx_secs: 10, timeout_2xx_secs: 20}",
		"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
		"--jitter-buffer", "{enable_jitter_buffer: true, jitterbuffer_msec_max: 200, jitterbuffer_msec_min: 60}",
		"--noise-suppression", "both",
		"--noise-suppression-details", "{attenuation_limit: 80, engine: deep_filter_net}",
		"--onnet-t38-passthrough-enabled=false",
		"--outbound", "{ani_override: string, ani_override_type: always, call_parking_enabled: true, channel_limit: 10, generate_ringback_tone: true, instant_ringback_enabled: true, ip_authentication_method: token, ip_authentication_token: string, localization: string, outbound_voice_profile_id: '1293384261075731499', t38_reinvite_source: customer, tech_prefix: string}",
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
	requestflag.CheckInnerFlags(ipConnectionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ip-connections", "create",
		"--active=true",
		"--anchorsite-override", "Latency",
		"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
		"--call-cost-in-webhooks=false",
		"--connection-name", "string",
		"--default-on-hold-comfort-noise-enabled=true",
		"--dtmf-type", "RFC 2833",
		"--encode-contact-header-enabled=true",
		"--encrypted-media", "SRTP",
		"--inbound.ani-number-format", "+E.164",
		"--inbound.channel-limit", "10",
		"--inbound.codecs", "[G722]",
		"--inbound.default-routing-method", "sequential",
		"--inbound.dnis-number-format", "+e164",
		"--inbound.generate-ringback-tone=true",
		"--inbound.isup-headers-enabled=true",
		"--inbound.prack-enabled=true",
		"--inbound.shaken-stir-enabled=true",
		"--inbound.sip-compact-headers-enabled=true",
		"--inbound.sip-region", "US",
		"--inbound.sip-subdomain", "test",
		"--inbound.sip-subdomain-receive-settings", "only_my_connections",
		"--inbound.timeout-1xx-secs", "10",
		"--inbound.timeout-2xx-secs", "20",
		"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
		"--jitter-buffer.enable-jitter-buffer=true",
		"--jitter-buffer.jitterbuffer-msec-max", "200",
		"--jitter-buffer.jitterbuffer-msec-min", "60",
		"--noise-suppression", "both",
		"--noise-suppression-details.attenuation-limit", "80",
		"--noise-suppression-details.engine", "deep_filter_net",
		"--onnet-t38-passthrough-enabled=false",
		"--outbound.ani-override", "string",
		"--outbound.ani-override-type", "always",
		"--outbound.call-parking-enabled=true",
		"--outbound.channel-limit", "10",
		"--outbound.generate-ringback-tone=true",
		"--outbound.instant-ringback-enabled=true",
		"--outbound.ip-authentication-method", "token",
		"--outbound.ip-authentication-token", "string",
		"--outbound.localization", "string",
		"--outbound.outbound-voice-profile-id", "1293384261075731499",
		"--outbound.t38-reinvite-source", "customer",
		"--outbound.tech-prefix", "string",
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

func TestIPConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ip-connections", "retrieve",
		"--id", "id",
	)
}

func TestIPConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ip-connections", "update",
		"--id", "id",
		"--active=true",
		"--anchorsite-override", "Latency",
		"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
		"--call-cost-in-webhooks=false",
		"--connection-name", "string",
		"--default-on-hold-comfort-noise-enabled=true",
		"--dtmf-type", "RFC 2833",
		"--encode-contact-header-enabled=true",
		"--encrypted-media", "SRTP",
		"--inbound", "{ani_number_format: +E.164, channel_limit: 10, codecs: [G722], default_primary_ip_id: 192.168.0.0, default_routing_method: sequential, default_secondary_ip_id: 192.168.0.0, default_tertiary_ip_id: 192.168.0.0, dnis_number_format: +e164, generate_ringback_tone: true, isup_headers_enabled: true, prack_enabled: true, shaken_stir_enabled: true, sip_compact_headers_enabled: true, sip_region: US, sip_subdomain: test, sip_subdomain_receive_settings: only_my_connections, timeout_1xx_secs: 10, timeout_2xx_secs: 20}",
		"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
		"--jitter-buffer", "{enable_jitter_buffer: true, jitterbuffer_msec_max: 200, jitterbuffer_msec_min: 60}",
		"--noise-suppression", "both",
		"--noise-suppression-details", "{attenuation_limit: 80, engine: deep_filter_net}",
		"--onnet-t38-passthrough-enabled=false",
		"--outbound", "{ani_override: string, ani_override_type: always, call_parking_enabled: true, channel_limit: 10, generate_ringback_tone: true, instant_ringback_enabled: true, ip_authentication_method: token, ip_authentication_token: string, localization: string, outbound_voice_profile_id: '1293384261075731499', t38_reinvite_source: customer, tech_prefix: string}",
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
	requestflag.CheckInnerFlags(ipConnectionsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ip-connections", "update",
		"--id", "id",
		"--active=true",
		"--anchorsite-override", "Latency",
		"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
		"--call-cost-in-webhooks=false",
		"--connection-name", "string",
		"--default-on-hold-comfort-noise-enabled=true",
		"--dtmf-type", "RFC 2833",
		"--encode-contact-header-enabled=true",
		"--encrypted-media", "SRTP",
		"--inbound.ani-number-format", "+E.164",
		"--inbound.channel-limit", "10",
		"--inbound.codecs", "[G722]",
		"--inbound.default-primary-ip-id", "192.168.0.0",
		"--inbound.default-routing-method", "sequential",
		"--inbound.default-secondary-ip-id", "192.168.0.0",
		"--inbound.default-tertiary-ip-id", "192.168.0.0",
		"--inbound.dnis-number-format", "+e164",
		"--inbound.generate-ringback-tone=true",
		"--inbound.isup-headers-enabled=true",
		"--inbound.prack-enabled=true",
		"--inbound.shaken-stir-enabled=true",
		"--inbound.sip-compact-headers-enabled=true",
		"--inbound.sip-region", "US",
		"--inbound.sip-subdomain", "test",
		"--inbound.sip-subdomain-receive-settings", "only_my_connections",
		"--inbound.timeout-1xx-secs", "10",
		"--inbound.timeout-2xx-secs", "20",
		"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
		"--jitter-buffer.enable-jitter-buffer=true",
		"--jitter-buffer.jitterbuffer-msec-max", "200",
		"--jitter-buffer.jitterbuffer-msec-min", "60",
		"--noise-suppression", "both",
		"--noise-suppression-details.attenuation-limit", "80",
		"--noise-suppression-details.engine", "deep_filter_net",
		"--onnet-t38-passthrough-enabled=false",
		"--outbound.ani-override", "string",
		"--outbound.ani-override-type", "always",
		"--outbound.call-parking-enabled=true",
		"--outbound.channel-limit", "10",
		"--outbound.generate-ringback-tone=true",
		"--outbound.instant-ringback-enabled=true",
		"--outbound.ip-authentication-method", "token",
		"--outbound.ip-authentication-token", "string",
		"--outbound.localization", "string",
		"--outbound.outbound-voice-profile-id", "1293384261075731499",
		"--outbound.t38-reinvite-source", "customer",
		"--outbound.tech-prefix", "string",
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

func TestIPConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ip-connections", "list",
		"--filter", "{connection_name: {contains: contains}, fqdn: fqdn, outbound_voice_profile_id: '1293384261075731499'}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "connection_name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(ipConnectionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ip-connections", "list",
		"--filter.connection-name", "{contains: contains}",
		"--filter.fqdn", "fqdn",
		"--filter.outbound-voice-profile-id", "1293384261075731499",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "connection_name",
	)
}

func TestIPConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ip-connections", "delete",
		"--id", "id",
	)
}

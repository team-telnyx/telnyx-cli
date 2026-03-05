// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestCredentialConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "credential-connections", "create",
			"--api-key", "string",
			"--connection-name", "my name",
			"--password", "my123secure456password789",
			"--user-name", "myusername123",
			"--active=true",
			"--anchorsite-override", "Latency",
			"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
			"--call-cost-in-webhooks=false",
			"--default-on-hold-comfort-noise-enabled=false",
			"--dtmf-type", "RFC 2833",
			"--encode-contact-header-enabled=true",
			"--encrypted-media", "SRTP",
			"--inbound", "{ani_number_format: +E.164, channel_limit: 10, codecs: [G722], default_routing_method: sequential, dnis_number_format: +e164, generate_ringback_tone: true, isup_headers_enabled: true, prack_enabled: true, shaken_stir_enabled: true, simultaneous_ringing: disabled, sip_compact_headers_enabled: true, timeout_1xx_secs: 10, timeout_2xx_secs: 20}",
			"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
			"--jitter-buffer", "{enable_jitter_buffer: true, jitterbuffer_msec_max: 200, jitterbuffer_msec_min: 60}",
			"--noise-suppression", "both",
			"--noise-suppression-details", "{attenuation_limit: 80, engine: deep_filter_net}",
			"--onnet-t38-passthrough-enabled=true",
			"--outbound", "{ani_override: always, ani_override_type: always, call_parking_enabled: true, channel_limit: 10, generate_ringback_tone: true, instant_ringback_enabled: true, localization: US, outbound_voice_profile_id: '1293384261075731499', t38_reinvite_source: customer}",
			"--rtcp-settings", "{capture_enabled: true, port: rtcp-mux, report_frequency_secs: 10}",
			"--sip-uri-calling-preference", "disabled",
			"--tag", "tag1",
			"--tag", "tag2",
			"--webhook-api-version", "1",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-event-url", "https://example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(credentialConnectionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "credential-connections", "create",
			"--api-key", "string",
			"--connection-name", "my name",
			"--password", "my123secure456password789",
			"--user-name", "myusername123",
			"--active=true",
			"--anchorsite-override", "Latency",
			"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
			"--call-cost-in-webhooks=false",
			"--default-on-hold-comfort-noise-enabled=false",
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
			"--inbound.simultaneous-ringing", "disabled",
			"--inbound.sip-compact-headers-enabled=true",
			"--inbound.timeout-1xx-secs", "10",
			"--inbound.timeout-2xx-secs", "20",
			"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
			"--jitter-buffer.enable-jitter-buffer=true",
			"--jitter-buffer.jitterbuffer-msec-max", "200",
			"--jitter-buffer.jitterbuffer-msec-min", "60",
			"--noise-suppression", "both",
			"--noise-suppression-details.attenuation-limit", "80",
			"--noise-suppression-details.engine", "deep_filter_net",
			"--onnet-t38-passthrough-enabled=true",
			"--outbound.ani-override", "always",
			"--outbound.ani-override-type", "always",
			"--outbound.call-parking-enabled=true",
			"--outbound.channel-limit", "10",
			"--outbound.generate-ringback-tone=true",
			"--outbound.instant-ringback-enabled=true",
			"--outbound.localization", "US",
			"--outbound.outbound-voice-profile-id", "1293384261075731499",
			"--outbound.t38-reinvite-source", "customer",
			"--rtcp-settings.capture-enabled=true",
			"--rtcp-settings.port", "rtcp-mux",
			"--rtcp-settings.report-frequency-secs", "10",
			"--sip-uri-calling-preference", "disabled",
			"--tag", "tag1",
			"--tag", "tag2",
			"--webhook-api-version", "1",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-event-url", "https://example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connection_name: my name\n" +
			"password: my123secure456password789\n" +
			"user_name: myusername123\n" +
			"active: true\n" +
			"anchorsite_override: Latency\n" +
			"android_push_credential_id: 06b09dfd-7154-4980-8b75-cebf7a9d4f8e\n" +
			"call_cost_in_webhooks: false\n" +
			"default_on_hold_comfort_noise_enabled: false\n" +
			"dtmf_type: RFC 2833\n" +
			"encode_contact_header_enabled: true\n" +
			"encrypted_media: SRTP\n" +
			"inbound:\n" +
			"  ani_number_format: +E.164\n" +
			"  channel_limit: 10\n" +
			"  codecs:\n" +
			"    - G722\n" +
			"  default_routing_method: sequential\n" +
			"  dnis_number_format: +e164\n" +
			"  generate_ringback_tone: true\n" +
			"  isup_headers_enabled: true\n" +
			"  prack_enabled: true\n" +
			"  shaken_stir_enabled: true\n" +
			"  simultaneous_ringing: disabled\n" +
			"  sip_compact_headers_enabled: true\n" +
			"  timeout_1xx_secs: 10\n" +
			"  timeout_2xx_secs: 20\n" +
			"ios_push_credential_id: ec0c8e5d-439e-4620-a0c1-9d9c8d02a836\n" +
			"jitter_buffer:\n" +
			"  enable_jitter_buffer: true\n" +
			"  jitterbuffer_msec_max: 200\n" +
			"  jitterbuffer_msec_min: 60\n" +
			"noise_suppression: both\n" +
			"noise_suppression_details:\n" +
			"  attenuation_limit: 80\n" +
			"  engine: deep_filter_net\n" +
			"onnet_t38_passthrough_enabled: true\n" +
			"outbound:\n" +
			"  ani_override: always\n" +
			"  ani_override_type: always\n" +
			"  call_parking_enabled: true\n" +
			"  channel_limit: 10\n" +
			"  generate_ringback_tone: true\n" +
			"  instant_ringback_enabled: true\n" +
			"  localization: US\n" +
			"  outbound_voice_profile_id: '1293384261075731499'\n" +
			"  t38_reinvite_source: customer\n" +
			"rtcp_settings:\n" +
			"  capture_enabled: true\n" +
			"  port: rtcp-mux\n" +
			"  report_frequency_secs: 10\n" +
			"sip_uri_calling_preference: disabled\n" +
			"tags:\n" +
			"  - tag1\n" +
			"  - tag2\n" +
			"webhook_api_version: '1'\n" +
			"webhook_event_failover_url: https://failover.example.com\n" +
			"webhook_event_url: https://example.com\n" +
			"webhook_timeout_secs: 25\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "credential-connections", "create",
			"--api-key", "string",
		)
	})
}

func TestCredentialConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "credential-connections", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestCredentialConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "credential-connections", "update",
			"--api-key", "string",
			"--id", "id",
			"--active=true",
			"--anchorsite-override", "Latency",
			"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
			"--call-cost-in-webhooks=false",
			"--connection-name", "my name",
			"--default-on-hold-comfort-noise-enabled=false",
			"--dtmf-type", "RFC 2833",
			"--encode-contact-header-enabled=true",
			"--encrypted-media", "SRTP",
			"--inbound", "{ani_number_format: +E.164, channel_limit: 10, codecs: [G722], default_routing_method: sequential, dnis_number_format: +e164, generate_ringback_tone: true, isup_headers_enabled: true, prack_enabled: true, shaken_stir_enabled: true, simultaneous_ringing: disabled, sip_compact_headers_enabled: true, timeout_1xx_secs: 10, timeout_2xx_secs: 20}",
			"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
			"--jitter-buffer", "{enable_jitter_buffer: true, jitterbuffer_msec_max: 200, jitterbuffer_msec_min: 60}",
			"--noise-suppression", "both",
			"--noise-suppression-details", "{attenuation_limit: 80, engine: deep_filter_net}",
			"--onnet-t38-passthrough-enabled=true",
			"--outbound", "{ani_override: always, ani_override_type: always, call_parking_enabled: true, channel_limit: 10, generate_ringback_tone: true, instant_ringback_enabled: true, localization: US, outbound_voice_profile_id: '1293384261075731499', t38_reinvite_source: customer}",
			"--password", "my123secure456password789",
			"--rtcp-settings", "{capture_enabled: true, port: rtcp-mux, report_frequency_secs: 10}",
			"--sip-uri-calling-preference", "disabled",
			"--tag", "tag1",
			"--tag", "tag2",
			"--user-name", "myusername123",
			"--webhook-api-version", "1",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-event-url", "https://example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(credentialConnectionsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "credential-connections", "update",
			"--api-key", "string",
			"--id", "id",
			"--active=true",
			"--anchorsite-override", "Latency",
			"--android-push-credential-id", "06b09dfd-7154-4980-8b75-cebf7a9d4f8e",
			"--call-cost-in-webhooks=false",
			"--connection-name", "my name",
			"--default-on-hold-comfort-noise-enabled=false",
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
			"--inbound.simultaneous-ringing", "disabled",
			"--inbound.sip-compact-headers-enabled=true",
			"--inbound.timeout-1xx-secs", "10",
			"--inbound.timeout-2xx-secs", "20",
			"--ios-push-credential-id", "ec0c8e5d-439e-4620-a0c1-9d9c8d02a836",
			"--jitter-buffer.enable-jitter-buffer=true",
			"--jitter-buffer.jitterbuffer-msec-max", "200",
			"--jitter-buffer.jitterbuffer-msec-min", "60",
			"--noise-suppression", "both",
			"--noise-suppression-details.attenuation-limit", "80",
			"--noise-suppression-details.engine", "deep_filter_net",
			"--onnet-t38-passthrough-enabled=true",
			"--outbound.ani-override", "always",
			"--outbound.ani-override-type", "always",
			"--outbound.call-parking-enabled=true",
			"--outbound.channel-limit", "10",
			"--outbound.generate-ringback-tone=true",
			"--outbound.instant-ringback-enabled=true",
			"--outbound.localization", "US",
			"--outbound.outbound-voice-profile-id", "1293384261075731499",
			"--outbound.t38-reinvite-source", "customer",
			"--password", "my123secure456password789",
			"--rtcp-settings.capture-enabled=true",
			"--rtcp-settings.port", "rtcp-mux",
			"--rtcp-settings.report-frequency-secs", "10",
			"--sip-uri-calling-preference", "disabled",
			"--tag", "tag1",
			"--tag", "tag2",
			"--user-name", "myusername123",
			"--webhook-api-version", "1",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-event-url", "https://example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"active: true\n" +
			"anchorsite_override: Latency\n" +
			"android_push_credential_id: 06b09dfd-7154-4980-8b75-cebf7a9d4f8e\n" +
			"call_cost_in_webhooks: false\n" +
			"connection_name: my name\n" +
			"default_on_hold_comfort_noise_enabled: false\n" +
			"dtmf_type: RFC 2833\n" +
			"encode_contact_header_enabled: true\n" +
			"encrypted_media: SRTP\n" +
			"inbound:\n" +
			"  ani_number_format: +E.164\n" +
			"  channel_limit: 10\n" +
			"  codecs:\n" +
			"    - G722\n" +
			"  default_routing_method: sequential\n" +
			"  dnis_number_format: +e164\n" +
			"  generate_ringback_tone: true\n" +
			"  isup_headers_enabled: true\n" +
			"  prack_enabled: true\n" +
			"  shaken_stir_enabled: true\n" +
			"  simultaneous_ringing: disabled\n" +
			"  sip_compact_headers_enabled: true\n" +
			"  timeout_1xx_secs: 10\n" +
			"  timeout_2xx_secs: 20\n" +
			"ios_push_credential_id: ec0c8e5d-439e-4620-a0c1-9d9c8d02a836\n" +
			"jitter_buffer:\n" +
			"  enable_jitter_buffer: true\n" +
			"  jitterbuffer_msec_max: 200\n" +
			"  jitterbuffer_msec_min: 60\n" +
			"noise_suppression: both\n" +
			"noise_suppression_details:\n" +
			"  attenuation_limit: 80\n" +
			"  engine: deep_filter_net\n" +
			"onnet_t38_passthrough_enabled: true\n" +
			"outbound:\n" +
			"  ani_override: always\n" +
			"  ani_override_type: always\n" +
			"  call_parking_enabled: true\n" +
			"  channel_limit: 10\n" +
			"  generate_ringback_tone: true\n" +
			"  instant_ringback_enabled: true\n" +
			"  localization: US\n" +
			"  outbound_voice_profile_id: '1293384261075731499'\n" +
			"  t38_reinvite_source: customer\n" +
			"password: my123secure456password789\n" +
			"rtcp_settings:\n" +
			"  capture_enabled: true\n" +
			"  port: rtcp-mux\n" +
			"  report_frequency_secs: 10\n" +
			"sip_uri_calling_preference: disabled\n" +
			"tags:\n" +
			"  - tag1\n" +
			"  - tag2\n" +
			"user_name: myusername123\n" +
			"webhook_api_version: '1'\n" +
			"webhook_event_failover_url: https://failover.example.com\n" +
			"webhook_event_url: https://example.com\n" +
			"webhook_timeout_secs: 25\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "credential-connections", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestCredentialConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "credential-connections", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{connection_name: {contains: contains}, fqdn: fqdn, outbound_voice_profile_id: '1293384261075731499'}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "connection_name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(credentialConnectionsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "credential-connections", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.connection-name", "{contains: contains}",
			"--filter.fqdn", "fqdn",
			"--filter.outbound-voice-profile-id", "1293384261075731499",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "connection_name",
		)
	})
}

func TestCredentialConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "credential-connections", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

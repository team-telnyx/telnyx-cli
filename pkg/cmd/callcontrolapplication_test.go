// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestCallControlApplicationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "call-control-applications", "create",
			"--api-key", "string",
			"--application-name", "call-router",
			"--webhook-event-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Latency",
			"--call-cost-in-webhooks=true",
			"--dtmf-type", "Inband",
			"--first-command-timeout=true",
			"--first-command-timeout-secs", "10",
			"--inbound", "{channel_limit: 10, shaken_stir_enabled: true, sip_subdomain: example, sip_subdomain_receive_settings: only_my_connections}",
			"--outbound", "{channel_limit: 10, outbound_voice_profile_id: '1293384261075731499'}",
			"--redact-dtmf-debug-logging=true",
			"--webhook-api-version", "1",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callControlApplicationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "call-control-applications", "create",
			"--api-key", "string",
			"--application-name", "call-router",
			"--webhook-event-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Latency",
			"--call-cost-in-webhooks=true",
			"--dtmf-type", "Inband",
			"--first-command-timeout=true",
			"--first-command-timeout-secs", "10",
			"--inbound.channel-limit", "10",
			"--inbound.shaken-stir-enabled=true",
			"--inbound.sip-subdomain", "example",
			"--inbound.sip-subdomain-receive-settings", "only_my_connections",
			"--outbound.channel-limit", "10",
			"--outbound.outbound-voice-profile-id", "1293384261075731499",
			"--redact-dtmf-debug-logging=true",
			"--webhook-api-version", "1",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"application_name: call-router\n" +
			"webhook_event_url: https://example.com\n" +
			"active: false\n" +
			"anchorsite_override: Latency\n" +
			"call_cost_in_webhooks: true\n" +
			"dtmf_type: Inband\n" +
			"first_command_timeout: true\n" +
			"first_command_timeout_secs: 10\n" +
			"inbound:\n" +
			"  channel_limit: 10\n" +
			"  shaken_stir_enabled: true\n" +
			"  sip_subdomain: example\n" +
			"  sip_subdomain_receive_settings: only_my_connections\n" +
			"outbound:\n" +
			"  channel_limit: 10\n" +
			"  outbound_voice_profile_id: '1293384261075731499'\n" +
			"redact_dtmf_debug_logging: true\n" +
			"webhook_api_version: '1'\n" +
			"webhook_event_failover_url: https://failover.example.com\n" +
			"webhook_timeout_secs: 25\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "call-control-applications", "create",
			"--api-key", "string",
		)
	})
}

func TestCallControlApplicationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "call-control-applications", "retrieve",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestCallControlApplicationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "call-control-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--application-name", "call-router",
			"--webhook-event-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Latency",
			"--call-cost-in-webhooks=true",
			"--dtmf-type", "Inband",
			"--first-command-timeout=true",
			"--first-command-timeout-secs", "10",
			"--inbound", "{channel_limit: 10, shaken_stir_enabled: true, sip_subdomain: example, sip_subdomain_receive_settings: only_my_connections}",
			"--outbound", "{channel_limit: 10, outbound_voice_profile_id: '1293384261075731499'}",
			"--redact-dtmf-debug-logging=true",
			"--tag", "tag1",
			"--tag", "tag2",
			"--webhook-api-version", "1",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callControlApplicationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "call-control-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--application-name", "call-router",
			"--webhook-event-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Latency",
			"--call-cost-in-webhooks=true",
			"--dtmf-type", "Inband",
			"--first-command-timeout=true",
			"--first-command-timeout-secs", "10",
			"--inbound.channel-limit", "10",
			"--inbound.shaken-stir-enabled=true",
			"--inbound.sip-subdomain", "example",
			"--inbound.sip-subdomain-receive-settings", "only_my_connections",
			"--outbound.channel-limit", "10",
			"--outbound.outbound-voice-profile-id", "1293384261075731499",
			"--redact-dtmf-debug-logging=true",
			"--tag", "tag1",
			"--tag", "tag2",
			"--webhook-api-version", "1",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"application_name: call-router\n" +
			"webhook_event_url: https://example.com\n" +
			"active: false\n" +
			"anchorsite_override: Latency\n" +
			"call_cost_in_webhooks: true\n" +
			"dtmf_type: Inband\n" +
			"first_command_timeout: true\n" +
			"first_command_timeout_secs: 10\n" +
			"inbound:\n" +
			"  channel_limit: 10\n" +
			"  shaken_stir_enabled: true\n" +
			"  sip_subdomain: example\n" +
			"  sip_subdomain_receive_settings: only_my_connections\n" +
			"outbound:\n" +
			"  channel_limit: 10\n" +
			"  outbound_voice_profile_id: '1293384261075731499'\n" +
			"redact_dtmf_debug_logging: true\n" +
			"tags:\n" +
			"  - tag1\n" +
			"  - tag2\n" +
			"webhook_api_version: '1'\n" +
			"webhook_event_failover_url: https://failover.example.com\n" +
			"webhook_timeout_secs: 25\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "call-control-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestCallControlApplicationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "call-control-applications", "list",
			"--api-key", "string",
			"--filter", "{application_name: {contains: contains}, application_session_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, connection_id: connection_id, failed: false, from: '+12025550142', leg_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, name: name, occurred_at: {eq: '2019-03-29T11:10:00Z', gt: '2019-03-29T11:10:00Z', gte: '2019-03-29T11:10:00Z', lt: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}, outbound.outbound_voice_profile_id: '1293384261075731499', product: texml, status: init, to: '+12025550142', type: webhook}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "connection_name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callControlApplicationsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "call-control-applications", "list",
			"--api-key", "string",
			"--filter.application-name", "{contains: contains}",
			"--filter.application-session-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter.connection-id", "connection_id",
			"--filter.failed=false",
			"--filter.from", "+12025550142",
			"--filter.leg-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter.name", "name",
			"--filter.occurred-at", "{eq: '2019-03-29T11:10:00Z', gt: '2019-03-29T11:10:00Z', gte: '2019-03-29T11:10:00Z', lt: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}",
			"--filter.outbound-outbound-voice-profile-id", "1293384261075731499",
			"--filter.product", "texml",
			"--filter.status", "init",
			"--filter.to", "+12025550142",
			"--filter.type", "webhook",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "connection_name",
		)
	})
}

func TestCallControlApplicationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "call-control-applications", "delete",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

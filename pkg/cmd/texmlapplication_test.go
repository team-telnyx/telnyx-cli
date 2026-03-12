// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTexmlApplicationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml-applications", "create",
			"--api-key", "string",
			"--friendly-name", "call-router",
			"--voice-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Amsterdam, Netherlands",
			"--call-cost-in-webhooks=false",
			"--dtmf-type", "Inband",
			"--first-command-timeout=true",
			"--first-command-timeout-secs", "10",
			"--inbound", "{channel_limit: 10, shaken_stir_enabled: true, sip_subdomain: example, sip_subdomain_receive_settings: only_my_connections}",
			"--outbound", "{channel_limit: 10, outbound_voice_profile_id: '1293384261075731499'}",
			"--status-callback", "https://example.com",
			"--status-callback-method", "get",
			"--tag", "tag1",
			"--tag", "tag2",
			"--voice-fallback-url", "https://fallback.example.com",
			"--voice-method", "get",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(texmlApplicationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "texml-applications", "create",
			"--api-key", "string",
			"--friendly-name", "call-router",
			"--voice-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Amsterdam, Netherlands",
			"--call-cost-in-webhooks=false",
			"--dtmf-type", "Inband",
			"--first-command-timeout=true",
			"--first-command-timeout-secs", "10",
			"--inbound.channel-limit", "10",
			"--inbound.shaken-stir-enabled=true",
			"--inbound.sip-subdomain", "example",
			"--inbound.sip-subdomain-receive-settings", "only_my_connections",
			"--outbound.channel-limit", "10",
			"--outbound.outbound-voice-profile-id", "1293384261075731499",
			"--status-callback", "https://example.com",
			"--status-callback-method", "get",
			"--tag", "tag1",
			"--tag", "tag2",
			"--voice-fallback-url", "https://fallback.example.com",
			"--voice-method", "get",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"friendly_name: call-router\n" +
			"voice_url: https://example.com\n" +
			"active: false\n" +
			"anchorsite_override: Amsterdam, Netherlands\n" +
			"call_cost_in_webhooks: false\n" +
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
			"status_callback: https://example.com\n" +
			"status_callback_method: get\n" +
			"tags:\n" +
			"  - tag1\n" +
			"  - tag2\n" +
			"voice_fallback_url: https://fallback.example.com\n" +
			"voice_method: get\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml-applications", "create",
			"--api-key", "string",
		)
	})
}

func TestTexmlApplicationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml-applications", "retrieve",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestTexmlApplicationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--friendly-name", "call-router",
			"--voice-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Amsterdam, Netherlands",
			"--call-cost-in-webhooks=false",
			"--dtmf-type", "Inband",
			"--first-command-timeout=true",
			"--first-command-timeout-secs", "10",
			"--inbound", "{channel_limit: 10, shaken_stir_enabled: true, sip_subdomain: example, sip_subdomain_receive_settings: only_my_connections}",
			"--outbound", "{channel_limit: 10, outbound_voice_profile_id: '1293384261075731499'}",
			"--status-callback", "https://example.com",
			"--status-callback-method", "get",
			"--tag", "tag1",
			"--tag", "tag2",
			"--voice-fallback-url", "https://fallback.example.com",
			"--voice-method", "get",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(texmlApplicationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "texml-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--friendly-name", "call-router",
			"--voice-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Amsterdam, Netherlands",
			"--call-cost-in-webhooks=false",
			"--dtmf-type", "Inband",
			"--first-command-timeout=true",
			"--first-command-timeout-secs", "10",
			"--inbound.channel-limit", "10",
			"--inbound.shaken-stir-enabled=true",
			"--inbound.sip-subdomain", "example",
			"--inbound.sip-subdomain-receive-settings", "only_my_connections",
			"--outbound.channel-limit", "10",
			"--outbound.outbound-voice-profile-id", "1293384261075731499",
			"--status-callback", "https://example.com",
			"--status-callback-method", "get",
			"--tag", "tag1",
			"--tag", "tag2",
			"--voice-fallback-url", "https://fallback.example.com",
			"--voice-method", "get",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"friendly_name: call-router\n" +
			"voice_url: https://example.com\n" +
			"active: false\n" +
			"anchorsite_override: Amsterdam, Netherlands\n" +
			"call_cost_in_webhooks: false\n" +
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
			"status_callback: https://example.com\n" +
			"status_callback_method: get\n" +
			"tags:\n" +
			"  - tag1\n" +
			"  - tag2\n" +
			"voice_fallback_url: https://fallback.example.com\n" +
			"voice_method: get\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestTexmlApplicationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml-applications", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{friendly_name: friendly_name, outbound_voice_profile_id: '1293384261075731499'}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "friendly_name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(texmlApplicationsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "texml-applications", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.friendly-name", "friendly_name",
			"--filter.outbound-voice-profile-id", "1293384261075731499",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "friendly_name",
		)
	})
}

func TestTexmlApplicationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml-applications", "delete",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

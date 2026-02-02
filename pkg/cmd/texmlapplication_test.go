// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestTexmlApplicationsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml-applications", "create",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(texmlApplicationsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml-applications", "create",
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
}

func TestTexmlApplicationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml-applications", "retrieve",
		"--id", "1293384261075731499",
	)
}

func TestTexmlApplicationsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml-applications", "update",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(texmlApplicationsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml-applications", "update",
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
}

func TestTexmlApplicationsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml-applications", "list",
		"--filter", "{friendly_name: friendly_name, outbound_voice_profile_id: '1293384261075731499'}",
		"--page", "{number: 1, size: 1}",
		"--sort", "friendly_name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(texmlApplicationsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml-applications", "list",
		"--filter.friendly-name", "friendly_name",
		"--filter.outbound-voice-profile-id", "1293384261075731499",
		"--page.number", "1",
		"--page.size", "1",
		"--sort", "friendly_name",
	)
}

func TestTexmlApplicationsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml-applications", "delete",
		"--id", "1293384261075731499",
	)
}

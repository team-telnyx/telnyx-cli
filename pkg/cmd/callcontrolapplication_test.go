// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestCallControlApplicationsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-control-applications", "create",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callControlApplicationsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-control-applications", "create",
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
}

func TestCallControlApplicationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-control-applications", "retrieve",
		"--id", "id",
	)
}

func TestCallControlApplicationsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-control-applications", "update",
		"--id", "id",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callControlApplicationsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-control-applications", "update",
		"--id", "id",
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
}

func TestCallControlApplicationsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-control-applications", "list",
		"--filter", "{application_name: {contains: contains}, application_session_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, connection_id: connection_id, failed: false, from: '+12025550142', leg_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, name: name, occurred_at: {eq: '2019-03-29T11:10:00Z', gt: '2019-03-29T11:10:00Z', gte: '2019-03-29T11:10:00Z', lt: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}, outbound.outbound_voice_profile_id: outbound.outbound_voice_profile_id, product: texml, status: init, to: '+12025550142', type: webhook}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "connection_name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callControlApplicationsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-control-applications", "list",
		"--filter.application-name", "{contains: contains}",
		"--filter.application-session-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.connection-id", "connection_id",
		"--filter.failed=false",
		"--filter.from", "+12025550142",
		"--filter.leg-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.name", "name",
		"--filter.occurred-at", "{eq: '2019-03-29T11:10:00Z', gt: '2019-03-29T11:10:00Z', gte: '2019-03-29T11:10:00Z', lt: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}",
		"--filter.outbound-outbound-voice-profile-id", "outbound.outbound_voice_profile_id",
		"--filter.product", "texml",
		"--filter.status", "init",
		"--filter.to", "+12025550142",
		"--filter.type", "webhook",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "connection_name",
	)
}

func TestCallControlApplicationsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-control-applications", "delete",
		"--id", "id",
	)
}

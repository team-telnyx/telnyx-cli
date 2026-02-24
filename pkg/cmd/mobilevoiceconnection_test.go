// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMobileVoiceConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-voice-connections", "create",
		"--active=true",
		"--connection-name", "connection_name",
		"--inbound", "{channel_limit: 0}",
		"--outbound", "{channel_limit: 0, outbound_voice_profile_id: outbound_voice_profile_id}",
		"--tag", "string",
		"--webhook-api-version", "1",
		"--webhook-event-failover-url", "webhook_event_failover_url",
		"--webhook-event-url", "webhook_event_url",
		"--webhook-timeout-secs", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(mobileVoiceConnectionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-voice-connections", "create",
		"--active=true",
		"--connection-name", "connection_name",
		"--inbound.channel-limit", "0",
		"--outbound.channel-limit", "0",
		"--outbound.outbound-voice-profile-id", "outbound_voice_profile_id",
		"--tag", "string",
		"--webhook-api-version", "1",
		"--webhook-event-failover-url", "webhook_event_failover_url",
		"--webhook-event-url", "webhook_event_url",
		"--webhook-timeout-secs", "0",
	)
}

func TestMobileVoiceConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-voice-connections", "retrieve",
		"--id", "id",
	)
}

func TestMobileVoiceConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-voice-connections", "update",
		"--id", "id",
		"--active=true",
		"--connection-name", "connection_name",
		"--inbound", "{channel_limit: 0}",
		"--outbound", "{channel_limit: 0, outbound_voice_profile_id: outbound_voice_profile_id}",
		"--tag", "string",
		"--webhook-api-version", "1",
		"--webhook-event-failover-url", "webhook_event_failover_url",
		"--webhook-event-url", "webhook_event_url",
		"--webhook-timeout-secs", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(mobileVoiceConnectionsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-voice-connections", "update",
		"--id", "id",
		"--active=true",
		"--connection-name", "connection_name",
		"--inbound.channel-limit", "0",
		"--outbound.channel-limit", "0",
		"--outbound.outbound-voice-profile-id", "outbound_voice_profile_id",
		"--tag", "string",
		"--webhook-api-version", "1",
		"--webhook-event-failover-url", "webhook_event_failover_url",
		"--webhook-event-url", "webhook_event_url",
		"--webhook-timeout-secs", "0",
	)
}

func TestMobileVoiceConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-voice-connections", "list",
		"--filter-connection-name-contains", "filter[connection_name][contains]",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "sort",
	)
}

func TestMobileVoiceConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-voice-connections", "delete",
		"--id", "id",
	)
}

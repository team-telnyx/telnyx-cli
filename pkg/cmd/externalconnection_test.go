// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestExternalConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "create",
		"--external-sip-connection", "zoom",
		"--outbound", "{channel_limit: 10, outbound_voice_profile_id: outbound_voice_profile_id}",
		"--active=false",
		"--inbound", "{outbound_voice_profile_id: 12345678-1234-1234-1234-123456789012, channel_limit: 10}",
		"--tag", "tag1",
		"--tag", "tag2",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(externalConnectionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "create",
		"--external-sip-connection", "zoom",
		"--outbound.channel-limit", "10",
		"--outbound.outbound-voice-profile-id", "outbound_voice_profile_id",
		"--active=false",
		"--inbound.outbound-voice-profile-id", "12345678-1234-1234-1234-123456789012",
		"--inbound.channel-limit", "10",
		"--tag", "tag1",
		"--tag", "tag2",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)
}

func TestExternalConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "retrieve",
		"--id", "id",
	)
}

func TestExternalConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "update",
		"--id", "id",
		"--outbound", "{outbound_voice_profile_id: outbound_voice_profile_id, channel_limit: 10}",
		"--active=false",
		"--inbound", "{channel_limit: 10}",
		"--tag", "tag1",
		"--tag", "tag2",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(externalConnectionsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "update",
		"--id", "id",
		"--outbound.outbound-voice-profile-id", "outbound_voice_profile_id",
		"--outbound.channel-limit", "10",
		"--active=false",
		"--inbound.channel-limit", "10",
		"--tag", "tag1",
		"--tag", "tag2",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)
}

func TestExternalConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "list",
		"--filter", "{id: '1930241863466354012', connection_name: {contains: My Connection}, created_at: '2022-12-31', external_sip_connection: zoom, phone_number: {contains: '+15555555555'}}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(externalConnectionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "list",
		"--filter.id", "1930241863466354012",
		"--filter.connection-name", "{contains: My Connection}",
		"--filter.created-at", "2022-12-31",
		"--filter.external-sip-connection", "zoom",
		"--filter.phone-number", "{contains: '+15555555555'}",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestExternalConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "delete",
		"--id", "id",
	)
}

func TestExternalConnectionsUpdateLocation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections", "update-location",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--location-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--static-emergency-address-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMobileVoiceConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mobile-voice-connections", "create",
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mobileVoiceConnectionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "mobile-voice-connections", "create",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"active: true\n" +
			"connection_name: connection_name\n" +
			"inbound:\n" +
			"  channel_limit: 0\n" +
			"outbound:\n" +
			"  channel_limit: 0\n" +
			"  outbound_voice_profile_id: outbound_voice_profile_id\n" +
			"tags:\n" +
			"  - string\n" +
			"webhook_api_version: '1'\n" +
			"webhook_event_failover_url: webhook_event_failover_url\n" +
			"webhook_event_url: webhook_event_url\n" +
			"webhook_timeout_secs: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "mobile-voice-connections", "create",
			"--api-key", "string",
		)
	})
}

func TestMobileVoiceConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mobile-voice-connections", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestMobileVoiceConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mobile-voice-connections", "update",
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mobileVoiceConnectionsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "mobile-voice-connections", "update",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"active: true\n" +
			"connection_name: connection_name\n" +
			"inbound:\n" +
			"  channel_limit: 0\n" +
			"outbound:\n" +
			"  channel_limit: 0\n" +
			"  outbound_voice_profile_id: outbound_voice_profile_id\n" +
			"tags:\n" +
			"  - string\n" +
			"webhook_api_version: '1'\n" +
			"webhook_event_failover_url: webhook_event_failover_url\n" +
			"webhook_event_url: webhook_event_url\n" +
			"webhook_timeout_secs: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "mobile-voice-connections", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestMobileVoiceConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mobile-voice-connections", "list",
			"--api-key", "string",
			"--filter-connection-name-contains", "filter[connection_name][contains]",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "sort",
		)
	})
}

func TestMobileVoiceConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mobile-voice-connections", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

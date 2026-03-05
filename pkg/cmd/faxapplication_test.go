// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestFaxApplicationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "fax-applications", "create",
			"--api-key", "string",
			"--application-name", "fax-router",
			"--webhook-event-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Amsterdam, Netherlands",
			"--inbound", "{channel_limit: 10, sip_subdomain: example, sip_subdomain_receive_settings: only_my_connections}",
			"--outbound", "{channel_limit: 10, outbound_voice_profile_id: '1293384261075731499'}",
			"--tag", "tag1",
			"--tag", "tag2",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(faxApplicationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "fax-applications", "create",
			"--api-key", "string",
			"--application-name", "fax-router",
			"--webhook-event-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Amsterdam, Netherlands",
			"--inbound.channel-limit", "10",
			"--inbound.sip-subdomain", "example",
			"--inbound.sip-subdomain-receive-settings", "only_my_connections",
			"--outbound.channel-limit", "10",
			"--outbound.outbound-voice-profile-id", "1293384261075731499",
			"--tag", "tag1",
			"--tag", "tag2",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"application_name: fax-router\n" +
			"webhook_event_url: https://example.com\n" +
			"active: false\n" +
			"anchorsite_override: Amsterdam, Netherlands\n" +
			"inbound:\n" +
			"  channel_limit: 10\n" +
			"  sip_subdomain: example\n" +
			"  sip_subdomain_receive_settings: only_my_connections\n" +
			"outbound:\n" +
			"  channel_limit: 10\n" +
			"  outbound_voice_profile_id: '1293384261075731499'\n" +
			"tags:\n" +
			"  - tag1\n" +
			"  - tag2\n" +
			"webhook_event_failover_url: https://failover.example.com\n" +
			"webhook_timeout_secs: 25\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "fax-applications", "create",
			"--api-key", "string",
		)
	})
}

func TestFaxApplicationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "fax-applications", "retrieve",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestFaxApplicationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "fax-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--application-name", "fax-router",
			"--webhook-event-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Amsterdam, Netherlands",
			"--fax-email-recipient", "user@example.com",
			"--inbound", "{channel_limit: 10, sip_subdomain: example, sip_subdomain_receive_settings: only_my_connections}",
			"--outbound", "{channel_limit: 10, outbound_voice_profile_id: '1293384261075731499'}",
			"--tag", "tag1",
			"--tag", "tag2",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(faxApplicationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "fax-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--application-name", "fax-router",
			"--webhook-event-url", "https://example.com",
			"--active=false",
			"--anchorsite-override", "Amsterdam, Netherlands",
			"--fax-email-recipient", "user@example.com",
			"--inbound.channel-limit", "10",
			"--inbound.sip-subdomain", "example",
			"--inbound.sip-subdomain-receive-settings", "only_my_connections",
			"--outbound.channel-limit", "10",
			"--outbound.outbound-voice-profile-id", "1293384261075731499",
			"--tag", "tag1",
			"--tag", "tag2",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"application_name: fax-router\n" +
			"webhook_event_url: https://example.com\n" +
			"active: false\n" +
			"anchorsite_override: Amsterdam, Netherlands\n" +
			"fax_email_recipient: user@example.com\n" +
			"inbound:\n" +
			"  channel_limit: 10\n" +
			"  sip_subdomain: example\n" +
			"  sip_subdomain_receive_settings: only_my_connections\n" +
			"outbound:\n" +
			"  channel_limit: 10\n" +
			"  outbound_voice_profile_id: '1293384261075731499'\n" +
			"tags:\n" +
			"  - tag1\n" +
			"  - tag2\n" +
			"webhook_event_failover_url: https://failover.example.com\n" +
			"webhook_timeout_secs: 25\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "fax-applications", "update",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestFaxApplicationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "fax-applications", "list",
			"--api-key", "string",
			"--filter", "{application_name: {contains: fax-app}, outbound_voice_profile_id: '1293384261075731499'}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "application_name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(faxApplicationsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "fax-applications", "list",
			"--api-key", "string",
			"--filter.application-name", "{contains: fax-app}",
			"--filter.outbound-voice-profile-id", "1293384261075731499",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "application_name",
		)
	})
}

func TestFaxApplicationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "fax-applications", "delete",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

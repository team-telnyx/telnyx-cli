// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestFaxApplicationsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fax-applications", "create",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(faxApplicationsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"fax-applications", "create",
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
}

func TestFaxApplicationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fax-applications", "retrieve",
		"--id", "1293384261075731499",
	)
}

func TestFaxApplicationsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fax-applications", "update",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(faxApplicationsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"fax-applications", "update",
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
}

func TestFaxApplicationsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fax-applications", "list",
		"--filter", "{application_name: {contains: fax-app}, outbound_voice_profile_id: '1293384261075731499'}",
		"--page", "{number: 1, size: 1}",
		"--sort", "application_name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(faxApplicationsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"fax-applications", "list",
		"--filter.application-name", "{contains: fax-app}",
		"--filter.outbound-voice-profile-id", "1293384261075731499",
		"--page.number", "1",
		"--page.size", "1",
		"--sort", "application_name",
	)
}

func TestFaxApplicationsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fax-applications", "delete",
		"--id", "1293384261075731499",
	)
}

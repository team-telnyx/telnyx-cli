// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcPartnerCampaignsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:partner-campaigns", "retrieve",
		"--campaign-id", "campaignId",
	)
}

func TestMessaging10dlcPartnerCampaignsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:partner-campaigns", "update",
		"--campaign-id", "campaignId",
		"--webhook-failover-url", "https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6",
		"--webhook-url", "https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93",
	)
}

func TestMessaging10dlcPartnerCampaignsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:partner-campaigns", "list",
		"--page", "0",
		"--records-per-page", "0",
		"--sort", "assignedPhoneNumbersCount",
	)
}

func TestMessaging10dlcPartnerCampaignsListSharedByMe(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:partner-campaigns", "list-shared-by-me",
		"--page", "0",
		"--records-per-page", "0",
	)
}

func TestMessaging10dlcPartnerCampaignsRetrieveSharingStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:partner-campaigns", "retrieve-sharing-status",
		"--campaign-id", "campaignId",
	)
}

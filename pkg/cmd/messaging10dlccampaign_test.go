// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcCampaignRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "retrieve",
		"--campaign-id", "campaignId",
	)
}

func TestMessaging10dlcCampaignUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "update",
		"--campaign-id", "campaignId",
		"--auto-renewal=true",
		"--help-message", "helpMessage",
		"--message-flow", "messageFlow",
		"--reseller-id", "resellerId",
		"--sample1", "sample1",
		"--sample2", "sample2",
		"--sample3", "sample3",
		"--sample4", "sample4",
		"--sample5", "sample5",
		"--webhook-failover-url", "webhookFailoverURL",
		"--webhook-url", "webhookURL",
	)
}

func TestMessaging10dlcCampaignList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "list",
		"--brand-id", "brandId",
		"--page", "0",
		"--records-per-page", "0",
		"--sort", "assignedPhoneNumbersCount",
	)
}

func TestMessaging10dlcCampaignAcceptSharing(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "accept-sharing",
		"--campaign-id", "C26F1KLZN",
	)
}

func TestMessaging10dlcCampaignDeactivate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "deactivate",
		"--campaign-id", "campaignId",
	)
}

func TestMessaging10dlcCampaignGetMnoMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "get-mno-metadata",
		"--campaign-id", "campaignId",
	)
}

func TestMessaging10dlcCampaignGetOperationStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "get-operation-status",
		"--campaign-id", "campaignId",
	)
}

func TestMessaging10dlcCampaignGetSharingStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "get-sharing-status",
		"--campaign-id", "campaignId",
	)
}

func TestMessaging10dlcCampaignSubmitAppeal(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign", "submit-appeal",
		"--campaign-id", "5eb13888-32b7-4cab-95e6-d834dde21d64",
		"--appeal-reason", "The website has been updated to include the required privacy policy and terms of service.",
	)
}

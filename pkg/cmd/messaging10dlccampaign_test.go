// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcCampaignRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "retrieve",
			"--campaign-id", "campaignId",
		)
	})
}

func TestMessaging10dlcCampaignUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "update",
			"--campaign-id", "campaignId",
			"--auto-renewal=true",
			"--help-message", "Helpmessage",
			"--message-flow", "Messageflow",
			"--reseller-id", "RESELLER",
			"--sample1", "Sample1",
			"--sample2", "Sample2",
			"--sample3", "Sample3",
			"--sample4", "Sample4",
			"--sample5", "Sample5",
			"--webhook-failover-url", "WebhookURL",
			"--webhook-url", "WebhookURL",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"autoRenewal: true\n" +
			"helpMessage: Helpmessage\n" +
			"messageFlow: Messageflow\n" +
			"resellerId: RESELLER\n" +
			"sample1: Sample1\n" +
			"sample2: Sample2\n" +
			"sample3: Sample3\n" +
			"sample4: Sample4\n" +
			"sample5: Sample5\n" +
			"webhookFailoverURL: WebhookURL\n" +
			"webhookURL: WebhookURL\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-10dlc:campaign", "update",
			"--campaign-id", "campaignId",
		)
	})
}

func TestMessaging10dlcCampaignList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "list",
			"--max-items", "10",
			"--brand-id", "brandId",
			"--page", "0",
			"--records-per-page", "0",
			"--sort", "assignedPhoneNumbersCount",
		)
	})
}

func TestMessaging10dlcCampaignAcceptSharing(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "accept-sharing",
			"--campaign-id", "C26F1KLZN",
		)
	})
}

func TestMessaging10dlcCampaignDeactivate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "deactivate",
			"--campaign-id", "campaignId",
		)
	})
}

func TestMessaging10dlcCampaignGetMnoMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "get-mno-metadata",
			"--campaign-id", "campaignId",
		)
	})
}

func TestMessaging10dlcCampaignGetOperationStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "get-operation-status",
			"--campaign-id", "campaignId",
		)
	})
}

func TestMessaging10dlcCampaignGetSharingStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "get-sharing-status",
			"--campaign-id", "campaignId",
		)
	})
}

func TestMessaging10dlcCampaignSubmitAppeal(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:campaign", "submit-appeal",
			"--campaign-id", "5eb13888-32b7-4cab-95e6-d834dde21d64",
			"--appeal-reason", "The website has been updated to include the required privacy policy and terms of service.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"appeal_reason: >-\n" +
			"  The website has been updated to include the required privacy policy and terms\n" +
			"  of service.\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-10dlc:campaign", "submit-appeal",
			"--campaign-id", "5eb13888-32b7-4cab-95e6-d834dde21d64",
		)
	})
}

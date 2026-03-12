// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcPartnerCampaignsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:partner-campaigns", "retrieve",
			"--api-key", "string",
			"--campaign-id", "campaignId",
		)
	})
}

func TestMessaging10dlcPartnerCampaignsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:partner-campaigns", "update",
			"--api-key", "string",
			"--campaign-id", "campaignId",
			"--webhook-failover-url", "https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6",
			"--webhook-url", "https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"webhookFailoverURL: https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6\n" +
			"webhookURL: https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-10dlc:partner-campaigns", "update",
			"--api-key", "string",
			"--campaign-id", "campaignId",
		)
	})
}

func TestMessaging10dlcPartnerCampaignsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:partner-campaigns", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--page", "0",
			"--records-per-page", "0",
			"--sort", "assignedPhoneNumbersCount",
		)
	})
}

func TestMessaging10dlcPartnerCampaignsListSharedByMe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:partner-campaigns", "list-shared-by-me",
			"--api-key", "string",
			"--max-items", "10",
			"--page", "0",
			"--records-per-page", "0",
		)
	})
}

func TestMessaging10dlcPartnerCampaignsRetrieveSharingStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:partner-campaigns", "retrieve-sharing-status",
			"--api-key", "string",
			"--campaign-id", "campaignId",
		)
	})
}

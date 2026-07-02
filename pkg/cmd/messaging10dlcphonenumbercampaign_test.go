// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMessaging10dlcPhoneNumberCampaignsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-campaigns", "create",
			"--campaign-id", "4b300178-131c-d902-d54e-72d90ba1620j",
			"--phone-number", "+18005550199",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"campaignId: 4b300178-131c-d902-d54e-72d90ba1620j\n" +
			"phoneNumber: '+18005550199'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-10dlc:phone-number-campaigns", "create",
		)
	})
}

func TestMessaging10dlcPhoneNumberCampaignsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-campaigns", "retrieve",
			"--phone-number", "phoneNumber",
		)
	})
}

func TestMessaging10dlcPhoneNumberCampaignsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-campaigns", "update",
			"--campaign-phone-number", "phoneNumber",
			"--campaign-id", "4b300178-131c-d902-d54e-72d90ba1620j",
			"--phone-number", "+18005550199",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"campaignId: 4b300178-131c-d902-d54e-72d90ba1620j\n" +
			"phoneNumber: '+18005550199'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-10dlc:phone-number-campaigns", "update",
			"--campaign-phone-number", "phoneNumber",
		)
	})
}

func TestMessaging10dlcPhoneNumberCampaignsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-campaigns", "list",
			"--max-items", "10",
			"--filter", "{tcr_brand_id: BRANDID, tcr_campaign_id: CAMPID3, telnyx_brand_id: f3575e15-32ce-400e-a4c0-dd78800c20b0, telnyx_campaign_id: f3575e15-32ce-400e-a4c0-dd78800c20b0}",
			"--page", "0",
			"--records-per-page", "0",
			"--sort", "assignmentStatus",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messaging10dlcPhoneNumberCampaignsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-campaigns", "list",
			"--max-items", "10",
			"--filter.tcr-brand-id", "BRANDID",
			"--filter.tcr-campaign-id", "CAMPID3",
			"--filter.telnyx-brand-id", "f3575e15-32ce-400e-a4c0-dd78800c20b0",
			"--filter.telnyx-campaign-id", "f3575e15-32ce-400e-a4c0-dd78800c20b0",
			"--page", "0",
			"--records-per-page", "0",
			"--sort", "assignmentStatus",
		)
	})
}

func TestMessaging10dlcPhoneNumberCampaignsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-campaigns", "delete",
			"--phone-number", "phoneNumber",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcCampaignBuilderSubmit(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign-builder", "submit",
		"--brand-id", "brandId",
		"--description", "description",
		"--usecase", "usecase",
		"--age-gated=true",
		"--auto-renewal=true",
		"--direct-lending=true",
		"--embedded-link=true",
		"--embedded-link-sample", "embeddedLinkSample",
		"--embedded-phone=true",
		"--help-keywords", "helpKeywords",
		"--help-message", "helpMessage",
		"--message-flow", "messageFlow",
		"--mno-id", "0",
		"--number-pool=true",
		"--optin-keywords", "optinKeywords",
		"--optin-message", "optinMessage",
		"--optout-keywords", "optoutKeywords",
		"--optout-message", "optoutMessage",
		"--privacy-policy-link", "privacyPolicyLink",
		"--reference-id", "referenceId",
		"--reseller-id", "resellerId",
		"--sample1", "sample1",
		"--sample2", "sample2",
		"--sample3", "sample3",
		"--sample4", "sample4",
		"--sample5", "sample5",
		"--subscriber-help=true",
		"--subscriber-optin=true",
		"--subscriber-optout=true",
		"--sub-usecase", "string",
		"--tag", "string",
		"--terms-and-conditions=true",
		"--terms-and-conditions-link", "termsAndConditionsLink",
		"--webhook-failover-url", "https://webhook.com/93711262-23e5-4048-a966-c0b2a16d5963",
		"--webhook-url", "https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcCampaignOsrGetAttributes(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign:osr", "get-attributes",
		"--campaign-id", "campaignId",
	)
}

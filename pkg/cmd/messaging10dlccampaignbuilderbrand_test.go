// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcCampaignBuilderBrandQualifyByUsecase(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:campaign-builder:brand", "qualify-by-usecase",
		"--brand-id", "brandId",
		"--usecase", "usecase",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLegacyReportingUsageReportsRetrieveSpeechToText(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"legacy:reporting:usage-reports", "retrieve-speech-to-text",
		"--end-date", "2020-07-01T00:00:00-06:00",
		"--start-date", "2020-07-01T00:00:00-06:00",
	)
}

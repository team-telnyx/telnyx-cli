// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestDetailRecordsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"detail-records", "list",
		"--filter", "{record_type: ai-voice-assistant, date_range: yesterday}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "string",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(detailRecordsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"detail-records", "list",
		"--filter.record-type", "ai-voice-assistant",
		"--filter.date-range", "yesterday",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "string",
	)
}

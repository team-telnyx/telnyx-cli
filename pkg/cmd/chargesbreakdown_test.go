// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestChargesBreakdownRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"charges-breakdown", "retrieve",
		"--start-date", "2025-05-01",
		"--end-date", "2025-06-01",
		"--format", "json",
	)
}

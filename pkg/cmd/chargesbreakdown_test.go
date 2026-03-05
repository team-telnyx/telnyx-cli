// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestChargesBreakdownRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "charges-breakdown", "retrieve",
			"--api-key", "string",
			"--start-date", "'2025-05-01'",
			"--end-date", "'2025-06-01'",
			"--format", "json",
		)
	})
}

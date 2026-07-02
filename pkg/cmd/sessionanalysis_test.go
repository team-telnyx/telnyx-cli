// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSessionAnalysisRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"session-analysis", "retrieve",
			"--record-type", "record_type",
			"--event-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--date-time", "'2019-12-27T18:11:19.117Z'",
			"--expand", "record",
			"--include-children=true",
			"--max-depth", "1",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestDialMetricsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dial-metrics", "list",
			"--max-items", "10",
			"--end-time", "'2019-12-27T18:11:19.117Z'",
			"--granularity", "second",
			"--start-time", "'2019-12-27T18:11:19.117Z'",
			"--page-after", "page[after]",
			"--page-limit", "1",
			"--sort", "-start_time",
		)
	})
}

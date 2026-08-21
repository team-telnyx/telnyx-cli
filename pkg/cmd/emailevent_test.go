// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailEventsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-events", "list",
			"--max-items", "10",
			"--email-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--event-type", "string",
			"--from", "'2019-12-27T18:11:19.117Z'",
			"--page-cursor", "page_cursor",
			"--page-size", "1",
			"--to", "'2019-12-27T18:11:19.117Z'",
		)
	})
}

func TestEmailEventsRetrieveStats(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-events", "retrieve-stats",
			"--from", "'2019-12-27T18:11:19.117Z'",
			"--to", "'2019-12-27T18:11:19.117Z'",
		)
	})
}

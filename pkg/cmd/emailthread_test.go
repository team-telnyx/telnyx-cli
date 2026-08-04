// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailThreadsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-threads", "retrieve",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-after", "page[after]",
			"--page-size", "1",
		)
	})
}

func TestEmailThreadsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-threads", "list",
			"--filter-inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-label", "filter[label]",
			"--page-after", "page[after]",
			"--page-size", "1",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestBulkSimCardActionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk-sim-card-actions", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestBulkSimCardActionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk-sim-card-actions", "list",
			"--max-items", "10",
			"--filter-action-type", "bulk_set_public_ips",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

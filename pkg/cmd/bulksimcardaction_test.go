// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestBulkSimCardActionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk-sim-card-actions", "retrieve",
		"--api-key", "string",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestBulkSimCardActionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk-sim-card-actions", "list",
		"--api-key", "string",
		"--filter-action-type", "bulk_set_public_ips",
		"--page-number", "1",
		"--page-size", "1",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAuditEventsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audit-events", "list",
		"--filter", "{created_after: '2021-01-01T00:00:00Z', created_before: '2021-01-01T00:00:00Z'}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "desc",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(auditEventsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"audit-events", "list",
		"--filter.created-after", "2021-01-01T00:00:00Z",
		"--filter.created-before", "2021-01-01T00:00:00Z",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "desc",
	)
}

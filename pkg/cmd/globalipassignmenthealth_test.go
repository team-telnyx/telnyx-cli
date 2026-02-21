// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestGlobalIPAssignmentHealthRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignment-health", "retrieve",
		"--filter", "{global_ip_assignment_id: string, global_ip_id: string}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(globalIPAssignmentHealthRetrieve)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignment-health", "retrieve",
		"--filter.global-ip-assignment-id", "string",
		"--filter.global-ip-id", "string",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestGlobalIPAssignmentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "create",
		"--api-key", "string",
	)
}

func TestGlobalIPAssignmentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "retrieve",
		"--api-key", "string",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestGlobalIPAssignmentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "update",
		"--api-key", "string",
		"--global-ip-assignment-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--global-ip-assignment-update-request", "{}",
	)
}

func TestGlobalIPAssignmentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "list",
		"--api-key", "string",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestGlobalIPAssignmentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "delete",
		"--api-key", "string",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestGlobalIPAssignmentsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "create",
	)
}

func TestGlobalIPAssignmentsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestGlobalIPAssignmentsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "update",
		"--global-ip-assignment-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--global-ip-assignment-update-request", "{}",
	)
}

func TestGlobalIPAssignmentsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "list",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(globalIPAssignmentsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "list",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestGlobalIPAssignmentsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-assignments", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

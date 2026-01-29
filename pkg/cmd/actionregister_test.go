// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestActionsRegisterCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"actions:register", "create",
		"--registration-code", "0000000001",
		"--registration-code", "0000000002",
		"--registration-code", "0000000003",
		"--sim-card-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--status", "standby",
		"--tag", "personal",
		"--tag", "customers",
		"--tag", "active-customers",
	)
}

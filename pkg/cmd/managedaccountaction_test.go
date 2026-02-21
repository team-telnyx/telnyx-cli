// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestManagedAccountsActionsDisable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts:actions", "disable",
		"--id", "id",
	)
}

func TestManagedAccountsActionsEnable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts:actions", "enable",
		"--id", "id",
		"--reenable-all-connections=true",
	)
}

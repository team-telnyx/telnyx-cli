// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestActionsPurchaseCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"actions:purchase", "create",
		"--amount", "10",
		"--product", "whitelabel",
		"--sim-card-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--status", "standby",
		"--tag", "personal",
		"--tag", "customers",
		"--tag", "active-customers",
		"--whitelabel-name", "Custom SPN",
	)
}

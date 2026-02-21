// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSimCardsActionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardsActionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "list",
		"--filter", "{action_type: disable, bulk_sim_card_action_id: 47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9, sim_card_id: 47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9, status: in-progress}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(simCardsActionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "list",
		"--filter.action-type", "disable",
		"--filter.bulk-sim-card-action-id", "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9",
		"--filter.sim-card-id", "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9",
		"--filter.status", "in-progress",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestSimCardsActionsBulkSetPublicIPs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "bulk-set-public-ips",
		"--sim-card-id", "6b14e151-8493-4fa1-8664-1cc4e6d14158",
	)
}

func TestSimCardsActionsDisable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "disable",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardsActionsEnable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "enable",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardsActionsRemovePublicIP(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "remove-public-ip",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardsActionsSetPublicIP(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "set-public-ip",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--region-code", "dc2",
	)
}

func TestSimCardsActionsSetStandby(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "set-standby",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardsActionsValidateRegistrationCodes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards:actions", "validate-registration-codes",
		"--registration-code", "123456780",
		"--registration-code", "1231231230",
	)
}

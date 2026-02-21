// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSimCardGroupsActionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-groups:actions", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardGroupsActionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-groups:actions", "list",
		"--filter-sim-card-group-id", "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9",
		"--filter-status", "in-progress",
		"--filter-type", "set_private_wireless_gateway",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestSimCardGroupsActionsRemovePrivateWirelessGateway(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-groups:actions", "remove-private-wireless-gateway",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardGroupsActionsRemoveWirelessBlocklist(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-groups:actions", "remove-wireless-blocklist",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardGroupsActionsSetPrivateWirelessGateway(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-groups:actions", "set-private-wireless-gateway",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--private-wireless-gateway-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardGroupsActionsSetWirelessBlocklist(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-groups:actions", "set-wireless-blocklist",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--wireless-blocklist-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

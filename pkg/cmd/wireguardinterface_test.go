// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestWireguardInterfacesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-interfaces", "create",
		"--region-code", "ashburn-va",
		"--enable-sip-trunking=false",
		"--name", "test interface",
		"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestWireguardInterfacesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-interfaces", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestWireguardInterfacesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-interfaces", "list",
		"--filter", "{network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(wireguardInterfacesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-interfaces", "list",
		"--filter.network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestWireguardInterfacesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-interfaces", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

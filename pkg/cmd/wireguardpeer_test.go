// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestWireguardPeersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-peers", "create",
		"--wireguard-interface-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--public-key", "qF4EqlZq+5JL2IKYY8ij49daYyfKVhevJrcDxdqC8GU=",
	)
}

func TestWireguardPeersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-peers", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestWireguardPeersUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-peers", "update",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--public-key", "qF4EqlZq+5JL2IKYY8ij49daYyfKVhevJrcDxdqC8GU=",
	)
}

func TestWireguardPeersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-peers", "list",
		"--filter", "{wireguard_interface_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(wireguardPeersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-peers", "list",
		"--filter.wireguard-interface-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestWireguardPeersDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireguard-peers", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

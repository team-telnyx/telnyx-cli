// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestWireguardPeersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireguard-peers", "create",
			"--api-key", "string",
			"--wireguard-interface-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("wireguard_interface_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "wireguard-peers", "create",
			"--api-key", "string",
		)
	})
}

func TestWireguardPeersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireguard-peers", "retrieve",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestWireguardPeersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireguard-peers", "update",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--public-key", "qF4EqlZq+5JL2IKYY8ij49daYyfKVhevJrcDxdqC8GU=",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("public_key: qF4EqlZq+5JL2IKYY8ij49daYyfKVhevJrcDxdqC8GU=")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "wireguard-peers", "update",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestWireguardPeersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireguard-peers", "list",
			"--api-key", "string",
			"--filter", "{wireguard_interface_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(wireguardPeersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "wireguard-peers", "list",
			"--api-key", "string",
			"--filter.wireguard-interface-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestWireguardPeersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireguard-peers", "delete",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

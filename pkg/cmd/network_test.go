// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNetworksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networks", "create",
			"--name", "test network",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: test network")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"networks", "create",
		)
	})
}

func TestNetworksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networks", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestNetworksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networks", "update",
			"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--name", "test network",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: test network")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"networks", "update",
			"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestNetworksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networks", "list",
			"--max-items", "10",
			"--filter", "{name: test network}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(networksList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networks", "list",
			"--max-items", "10",
			"--filter.name", "test network",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestNetworksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networks", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestNetworksListInterfaces(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networks", "list-interfaces",
			"--max-items", "10",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--filter", "{name: test interface, type: wireguard_interface}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(networksListInterfaces)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networks", "list-interfaces",
			"--max-items", "10",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--filter.name", "test interface",
			"--filter.type", "wireguard_interface",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

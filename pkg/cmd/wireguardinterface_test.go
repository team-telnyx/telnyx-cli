// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestWireguardInterfacesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wireguard-interfaces", "create",
			"--region-code", "ashburn-va",
			"--enable-sip-trunking=false",
			"--name", "test interface",
			"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"region_code: ashburn-va\n" +
			"enable_sip_trunking: false\n" +
			"name: test interface\n" +
			"network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"wireguard-interfaces", "create",
		)
	})
}

func TestWireguardInterfacesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wireguard-interfaces", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestWireguardInterfacesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wireguard-interfaces", "list",
			"--max-items", "10",
			"--filter", "{network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(wireguardInterfacesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wireguard-interfaces", "list",
			"--max-items", "10",
			"--filter.network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestWireguardInterfacesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wireguard-interfaces", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestNetworksDefaultGatewayCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "networks:default-gateway", "create",
			"--api-key", "string",
			"--network-identifier", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--wireguard-peer-id", "e66c496d-4a85-423b-8b2a-8e63fac20320",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("wireguard_peer_id: e66c496d-4a85-423b-8b2a-8e63fac20320")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "networks:default-gateway", "create",
			"--api-key", "string",
			"--network-identifier", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestNetworksDefaultGatewayRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "networks:default-gateway", "retrieve",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestNetworksDefaultGatewayDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "networks:default-gateway", "delete",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

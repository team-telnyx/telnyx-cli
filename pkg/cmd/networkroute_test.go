// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNetworkRoutesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"network-routes", "create",
			"--cidr", "10.20.0.0/16",
			"--network-id", "91446ee3-5b1c-4f0e-8d2a-6c7e9b0a1f24",
			"--next-hop", "{id: 62d438d2-8a51-4c5e-9d0c-1f6b2a7e4c39, type: tunnel}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(networkRoutesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"network-routes", "create",
			"--cidr", "10.20.0.0/16",
			"--network-id", "91446ee3-5b1c-4f0e-8d2a-6c7e9b0a1f24",
			"--next-hop.id", "62d438d2-8a51-4c5e-9d0c-1f6b2a7e4c39",
			"--next-hop.type", "tunnel",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cidr: 10.20.0.0/16\n" +
			"network_id: 91446ee3-5b1c-4f0e-8d2a-6c7e9b0a1f24\n" +
			"next_hop:\n" +
			"  id: 62d438d2-8a51-4c5e-9d0c-1f6b2a7e4c39\n" +
			"  type: tunnel\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"network-routes", "create",
		)
	})
}

func TestNetworkRoutesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"network-routes", "retrieve",
			"--id", "e8a2c5d0-3f7b-4a19-b6d2-9c1e0f8a7b53",
		)
	})
}

func TestNetworkRoutesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"network-routes", "list",
			"--filter", "{network_id: 91446ee3-5b1c-4f0e-8d2a-6c7e9b0a1f24, cidr: 10.20.0.0/16, next_hop.type: tunnel, origin: static, status: active}",
			"--page", "{after: after, limit: 1}",
			"--sort", "-created_at",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(networkRoutesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"network-routes", "list",
			"--filter.network-id", "91446ee3-5b1c-4f0e-8d2a-6c7e9b0a1f24",
			"--filter.cidr", "10.20.0.0/16",
			"--filter.next-hop-type", "tunnel",
			"--filter.origin", "static",
			"--filter.status", "active",
			"--page.after", "after",
			"--page.limit", "1",
			"--sort", "-created_at",
		)
	})
}

func TestNetworkRoutesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"network-routes", "delete",
			"--id", "e8a2c5d0-3f7b-4a19-b6d2-9c1e0f8a7b53",
		)
	})
}

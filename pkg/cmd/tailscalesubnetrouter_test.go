// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTailscaleSubnetRoutersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tailscale-subnet-routers", "create",
			"--advertised-route", "172.27.0.0/24",
			"--auth-key", "tskey-auth-k1AbCdEfGh2IjKlMnOpQrStUv",
			"--name", "tailnet-ashburn",
			"--network-id", "0de011ae-ee53-4df4-ba96-754728eb5aad",
			"--region-code", "ashburn-va",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"advertised_routes:\n" +
			"  - 172.27.0.0/24\n" +
			"auth_key: tskey-auth-k1AbCdEfGh2IjKlMnOpQrStUv\n" +
			"name: tailnet-ashburn\n" +
			"network_id: 0de011ae-ee53-4df4-ba96-754728eb5aad\n" +
			"region_code: ashburn-va\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tailscale-subnet-routers", "create",
		)
	})
}

func TestTailscaleSubnetRoutersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tailscale-subnet-routers", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestTailscaleSubnetRoutersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tailscale-subnet-routers", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--advertised-route", "172.27.0.0/24",
			"--advertised-route", "172.27.1.0/24",
			"--auth-key", "tskey-auth-k1AbCdEfGh2IjKlMnOpQrStUv",
			"--name", "tailnet-ashburn",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"advertised_routes:\n" +
			"  - 172.27.0.0/24\n" +
			"  - 172.27.1.0/24\n" +
			"auth_key: tskey-auth-k1AbCdEfGh2IjKlMnOpQrStUv\n" +
			"name: tailnet-ashburn\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tailscale-subnet-routers", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestTailscaleSubnetRoutersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tailscale-subnet-routers", "list",
			"--max-items", "10",
			"--filter", "{network_id: 0de011ae-ee53-4df4-ba96-754728eb5aad}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(tailscaleSubnetRoutersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tailscale-subnet-routers", "list",
			"--max-items", "10",
			"--filter.network-id", "0de011ae-ee53-4df4-ba96-754728eb5aad",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestTailscaleSubnetRoutersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tailscale-subnet-routers", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

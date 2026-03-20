// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPrivateWirelessGatewaysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"private-wireless-gateways", "create",
			"--name", "My private wireless gateway",
			"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--region-code", "dc2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My private wireless gateway\n" +
			"network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"region_code: dc2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"private-wireless-gateways", "create",
		)
	})
}

func TestPrivateWirelessGatewaysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"private-wireless-gateways", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestPrivateWirelessGatewaysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"private-wireless-gateways", "list",
			"--max-items", "10",
			"--filter-created-at", "filter[created_at]",
			"--filter-ip-range", "filter[ip_range]",
			"--filter-name", "filter[name]",
			"--filter-region-code", "filter[region_code]",
			"--filter-updated-at", "filter[updated_at]",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestPrivateWirelessGatewaysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"private-wireless-gateways", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

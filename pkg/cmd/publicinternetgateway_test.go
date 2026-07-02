// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPublicInternetGatewaysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"public-internet-gateways", "create",
			"--name", "test interface",
			"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--region-code", "ashburn-va",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: test interface\n" +
			"network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"region_code: ashburn-va\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"public-internet-gateways", "create",
		)
	})
}

func TestPublicInternetGatewaysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"public-internet-gateways", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestPublicInternetGatewaysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"public-internet-gateways", "list",
			"--max-items", "10",
			"--filter", "{network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(publicInternetGatewaysList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"public-internet-gateways", "list",
			"--max-items", "10",
			"--filter.network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestPublicInternetGatewaysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"public-internet-gateways", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

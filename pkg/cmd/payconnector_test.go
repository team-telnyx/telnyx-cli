// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPayConnectorsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pay-connectors", "create",
			"--config", "{auth_type: none, token: {}, endpoint_url: https://processor.example.com/pay, password: {}, timeout: 1, username: {}}",
			"--name", "my-payment-connector",
			"--mode", "test",
			"--type", "generic",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"config:\n" +
			"  auth_type: none\n" +
			"  token: {}\n" +
			"  endpoint_url: https://processor.example.com/pay\n" +
			"  password: {}\n" +
			"  timeout: 1\n" +
			"  username: {}\n" +
			"name: my-payment-connector\n" +
			"mode: test\n" +
			"type: generic\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pay-connectors", "create",
		)
	})
}

func TestPayConnectorsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pay-connectors", "retrieve",
			"--name", "my-payment-connector",
		)
	})
}

func TestPayConnectorsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pay-connectors", "update",
			"--name", "my-payment-connector",
			"--config", "{password: P, username: P, token: {}, auth_type: basic, endpoint_url: https://processor.example.com/pay, timeout: 1}",
			"--mode", "live",
			"--type", "generic",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"config:\n" +
			"  password: P\n" +
			"  username: P\n" +
			"  token: {}\n" +
			"  auth_type: basic\n" +
			"  endpoint_url: https://processor.example.com/pay\n" +
			"  timeout: 1\n" +
			"mode: live\n" +
			"type: generic\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pay-connectors", "update",
			"--name", "my-payment-connector",
		)
	})
}

func TestPayConnectorsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pay-connectors", "list",
			"--max-items", "10",
			"--filter-created-at-eq", "2026-10-01T00:00:00Z",
			"--filter-created-at-gt", "2026-10-01T00:00:00Z",
			"--filter-created-at-gte", "2026-10-01T00:00:00Z",
			"--filter-created-at-lt", "2026-10-01T00:00:00Z",
			"--filter-created-at-lte", "2026-10-01T00:00:00Z",
			"--filter-mode", "test",
			"--filter-name", "my-payment-connector",
			"--filter-type", "generic",
			"--filter-updated-at-eq", "2026-10-01T00:00:00Z",
			"--filter-updated-at-gt", "2026-10-01T00:00:00Z",
			"--filter-updated-at-gte", "2026-10-01T00:00:00Z",
			"--filter-updated-at-lt", "2026-10-01T00:00:00Z",
			"--filter-updated-at-lte", "2026-10-01T00:00:00Z",
			"--page-after", "page[after]",
			"--page-before", "page[before]",
			"--page-size", "1",
			"--sort", "-created_at",
		)
	})
}

func TestPayConnectorsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pay-connectors", "delete",
			"--name", "my-payment-connector",
		)
	})
}

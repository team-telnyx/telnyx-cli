// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestIPsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ips", "create",
			"--ip-address", "192.168.0.0",
			"--connection-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--port", "5060",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ip_address: 192.168.0.0\n" +
			"connection_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"port: 5060\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ips", "create",
		)
	})
}

func TestIPsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ips", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestIPsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ips", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--ip-address", "192.168.0.0",
			"--connection-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--port", "5060",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ip_address: 192.168.0.0\n" +
			"connection_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"port: 5060\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ips", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestIPsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ips", "list",
			"--max-items", "10",
			"--filter", "{connection_id: connection_id, ip_address: 192.168.0.0, port: 5060}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(ipsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ips", "list",
			"--max-items", "10",
			"--filter.connection-id", "connection_id",
			"--filter.ip-address", "192.168.0.0",
			"--filter.port", "5060",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestIPsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ips", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

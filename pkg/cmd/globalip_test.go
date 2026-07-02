// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestGlobalIPsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-ips", "create",
			"--description", "test interface",
			"--name", "test interface",
			"--ports", "{tcp: bar, udp: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: test interface\n" +
			"name: test interface\n" +
			"ports:\n" +
			"  tcp: bar\n" +
			"  udp: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"global-ips", "create",
		)
	})
}

func TestGlobalIPsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-ips", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestGlobalIPsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-ips", "list",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestGlobalIPsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-ips", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

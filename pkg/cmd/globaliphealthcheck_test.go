// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestGlobalIPHealthChecksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "global-ip-health-checks", "create",
			"--api-key", "string",
			"--global-ip-id", "a836125b-20b6-452e-9c03-2653f09c7ed7",
			"--health-check-params", "{path: bar, port: bar}",
			"--health-check-type", "http_status_2xx",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"global_ip_id: a836125b-20b6-452e-9c03-2653f09c7ed7\n" +
			"health_check_params:\n" +
			"  path: bar\n" +
			"  port: bar\n" +
			"health_check_type: http_status_2xx\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "global-ip-health-checks", "create",
			"--api-key", "string",
		)
	})
}

func TestGlobalIPHealthChecksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "global-ip-health-checks", "retrieve",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestGlobalIPHealthChecksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "global-ip-health-checks", "list",
			"--api-key", "string",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestGlobalIPHealthChecksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "global-ip-health-checks", "delete",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

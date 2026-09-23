// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestComputeFuncsExportCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:funcs:export", "create",
			"--id", "id",
			"--endpoint", "https://api.honeycomb.io/v1/logs",
			"--headers", "{x-honeycomb-team: abc123}",
			"--invocation-export-enabled=true",
			"--runtime-export-enabled=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"endpoint: https://api.honeycomb.io/v1/logs\n" +
			"headers:\n" +
			"  x-honeycomb-team: abc123\n" +
			"invocation_export_enabled: true\n" +
			"runtime_export_enabled: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:funcs:export", "create",
			"--id", "id",
		)
	})
}

func TestComputeFuncsExportList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:funcs:export", "list",
			"--id", "id",
		)
	})
}

func TestComputeFuncsExportDeleteAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:funcs:export", "delete-all",
			"--id", "id",
		)
	})
}

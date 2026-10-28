// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestBrowserSessionsActionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"browser-sessions:actions", "create",
			"--id", "id",
			"--type", "navigate",
			"--url", "https://example.com",
			"--wait-seconds", "0",
			"--idempotency-key", "x",
			"--include-observation=true",
			"--wait-until", "load",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: navigate\n" +
			"url: https://example.com\n" +
			"idempotency_key: x\n" +
			"include_observation: true\n" +
			"wait_until: load\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"browser-sessions:actions", "create",
			"--id", "id",
			"--wait-seconds", "0",
		)
	})
}

func TestBrowserSessionsActionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"browser-sessions:actions", "retrieve",
			"--id", "id",
			"--action-id", "actionId",
		)
	})
}

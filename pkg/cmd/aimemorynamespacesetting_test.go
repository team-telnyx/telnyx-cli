// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIMemoryNamespacesSettingsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:settings", "list",
			"--namespace", "namespace",
		)
	})
}

func TestAIMemoryNamespacesSettingsPatchAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:settings", "patch-all",
			"--namespace", "namespace",
			"--summary", "{instructions: Lead with the customer's plan tier. Keep it under 100 words.}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiMemoryNamespacesSettingsPatchAll)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:settings", "patch-all",
			"--namespace", "namespace",
			"--summary.instructions", "Lead with the customer's plan tier. Keep it under 100 words.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"summary:\n" +
			"  instructions: Lead with the customer's plan tier. Keep it under 100 words.\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:memory:namespaces:settings", "patch-all",
			"--namespace", "namespace",
		)
	})
}

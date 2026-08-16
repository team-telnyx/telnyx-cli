// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAICollectionsSettingsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:settings", "create",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--retrieval", "{retrieval_type: vector, top_k: 5}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiCollectionsSettingsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:settings", "create",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--retrieval.retrieval-type", "vector",
			"--retrieval.top-k", "5",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"retrieval:\n" +
			"  retrieval_type: vector\n" +
			"  top_k: 5\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:collections:settings", "create",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}

func TestAICollectionsSettingsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:settings", "list",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}

func TestAICollectionsSettingsPatchAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:settings", "patch-all",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--retrieval", "{retrieval_type: vector, top_k: 5}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiCollectionsSettingsPatchAll)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:settings", "patch-all",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--retrieval.retrieval-type", "vector",
			"--retrieval.top-k", "5",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"retrieval:\n" +
			"  retrieval_type: vector\n" +
			"  top_k: 5\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:collections:settings", "patch-all",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}

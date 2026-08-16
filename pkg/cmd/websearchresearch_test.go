// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWebSearchResearchCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-search:research", "create",
			"--query", "Compare the performance of RAG vs fine-tuning for domain-specific QA",
			"--background=false",
			"--max-sources", "20",
			"--research-effort", "standard",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: Compare the performance of RAG vs fine-tuning for domain-specific QA\n" +
			"background: false\n" +
			"max_sources: 20\n" +
			"research_effort: standard\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web-search:research", "create",
		)
	})
}

func TestWebSearchResearchRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-search:research", "retrieve",
			"--task-id", "bf3026a5-dd57-44dd-b922-200041be3a4b",
		)
	})
}

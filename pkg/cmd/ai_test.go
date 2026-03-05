// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIRetrieveModels(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai", "retrieve-models",
			"--api-key", "string",
		)
	})
}

func TestAISummarize(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai", "summarize",
			"--api-key", "string",
			"--bucket", "bucket",
			"--filename", "filename",
			"--system-prompt", "system_prompt",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket: bucket\n" +
			"filename: filename\n" +
			"system_prompt: system_prompt\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai", "summarize",
			"--api-key", "string",
		)
	})
}

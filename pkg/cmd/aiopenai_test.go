// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIOpenAICreateResponse(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:openai", "create-response",
			"--conversation", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--input", "{'0': bar}",
			"--instructions", "You are a friendly chatbot.",
			"--model", "zai-org/GLM-5.1-FP8",
			"--stream=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"conversation: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"input:\n" +
			"  '0': bar\n" +
			"instructions: You are a friendly chatbot.\n" +
			"model: zai-org/GLM-5.1-FP8\n" +
			"stream: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:openai", "create-response",
		)
	})
}

func TestAIOpenAIListModels(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:openai", "list-models",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
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
			"--reasoning", "{effort: none}",
			"--service-tier", "service_tier",
			"--stream=false",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiOpenAICreateResponse)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:openai", "create-response",
			"--conversation", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--input", "{'0': bar}",
			"--instructions", "You are a friendly chatbot.",
			"--model", "zai-org/GLM-5.1-FP8",
			"--reasoning.effort", "none",
			"--service-tier", "service_tier",
			"--stream=false",
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
			"reasoning:\n" +
			"  effort: none\n" +
			"service_tier: service_tier\n" +
			"stream: false\n")
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestBotChallengeCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bot-challenge", "create",
			"--llm-model-name", "claude-opus-4",
			"--llm-parameter-count", "175B",
			"--llm-quantization", "int8",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"llm_model_name: claude-opus-4\n" +
			"llm_parameter_count: 175B\n" +
			"llm_quantization: int8\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bot-challenge", "create",
		)
	})
}

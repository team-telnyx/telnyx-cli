// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAnthropicV1Messages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:anthropic:v1", "messages",
			"--max-tokens", "1024",
			"--message", "{role: bar, content: bar}",
			"--model", "zai-org/GLM-5.2",
			"--api-key-ref", "api_key_ref",
			"--billing-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--fallback-config", "{foo: bar}",
			"--max-retries", "0",
			"--mcp-server", "{foo: bar}",
			"--metadata", "{foo: bar}",
			"--service-tier", "service_tier",
			"--stop-sequence", "string",
			"--stream=true",
			"--system", "You are a friendly chatbot.",
			"--temperature", "0",
			"--thinking", "{foo: bar}",
			"--timeout", "0",
			"--tool-choice", "{foo: bar}",
			"--tool", "{foo: bar}",
			"--top-k", "0",
			"--top-p", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"max_tokens: 1024\n" +
			"messages:\n" +
			"  - role: bar\n" +
			"    content: bar\n" +
			"model: zai-org/GLM-5.2\n" +
			"api_key_ref: api_key_ref\n" +
			"billing_group_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"fallback_config:\n" +
			"  foo: bar\n" +
			"max_retries: 0\n" +
			"mcp_servers:\n" +
			"  - foo: bar\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"service_tier: service_tier\n" +
			"stop_sequences:\n" +
			"  - string\n" +
			"stream: true\n" +
			"system: You are a friendly chatbot.\n" +
			"temperature: 0\n" +
			"thinking:\n" +
			"  foo: bar\n" +
			"timeout: 0\n" +
			"tool_choice:\n" +
			"  foo: bar\n" +
			"tools:\n" +
			"  - foo: bar\n" +
			"top_k: 0\n" +
			"top_p: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:anthropic:v1", "messages",
		)
	})
}

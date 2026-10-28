// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAITestsTestFrameworkV1AiagentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:aiagents", "create",
			"--agent-name", "Support Bot Under Test",
			"--assistant-id", "e1a2b3c4-d5e6-f7a8-b9c0-d1e2f3a4b5c6",
			"--auto-fetch-calls-enabled=true",
			"--auto-sync-prompt-enabled=true",
			"--auto-update-dropoff-nodes=false",
			"--auto-update-topic-nodes=false",
			"--contact-number", "contact_number",
			"--description", "Evaluates the support assistant against generated callers.",
			"--inbound=false",
			"--language", "en",
			"--llm-max-tokens", "512",
			"--llm-model", "openai/gpt-oss-120b",
			"--llm-system-prompt", "llm_system_prompt",
			"--llm-temperature", "0.7",
			"--outbound-auto-call=false",
			"--sip-endpoint", "",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_name: Support Bot Under Test\n" +
			"assistant_id: e1a2b3c4-d5e6-f7a8-b9c0-d1e2f3a4b5c6\n" +
			"auto_fetch_calls_enabled: true\n" +
			"auto_sync_prompt_enabled: true\n" +
			"auto_update_dropoff_nodes: false\n" +
			"auto_update_topic_nodes: false\n" +
			"contact_number: contact_number\n" +
			"description: Evaluates the support assistant against generated callers.\n" +
			"inbound: false\n" +
			"language: en\n" +
			"llm_max_tokens: 512\n" +
			"llm_model: openai/gpt-oss-120b\n" +
			"llm_system_prompt: llm_system_prompt\n" +
			"llm_temperature: 0.7\n" +
			"outbound_auto_call: false\n" +
			"sip_endpoint: ''\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:aiagents", "create",
		)
	})
}

func TestAITestsTestFrameworkV1AiagentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:aiagents", "retrieve",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1AiagentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:aiagents", "update",
			"--id", "0",
			"--agent-name", "Support Bot Under Test - renamed",
			"--assistant-id", "assistant_id",
			"--auto-fetch-calls-enabled=true",
			"--auto-sync-prompt-enabled=true",
			"--auto-update-dropoff-nodes=true",
			"--auto-update-topic-nodes=true",
			"--contact-number", "contact_number",
			"--description", "description",
			"--inbound=true",
			"--language", "language",
			"--llm-max-tokens", "256",
			"--llm-model", "llm_model",
			"--llm-system-prompt", "llm_system_prompt",
			"--llm-temperature", "0.5",
			"--outbound-auto-call=true",
			"--sip-endpoint", "sip_endpoint",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_name: Support Bot Under Test - renamed\n" +
			"assistant_id: assistant_id\n" +
			"auto_fetch_calls_enabled: true\n" +
			"auto_sync_prompt_enabled: true\n" +
			"auto_update_dropoff_nodes: true\n" +
			"auto_update_topic_nodes: true\n" +
			"contact_number: contact_number\n" +
			"description: description\n" +
			"inbound: true\n" +
			"language: language\n" +
			"llm_max_tokens: 256\n" +
			"llm_model: llm_model\n" +
			"llm_system_prompt: llm_system_prompt\n" +
			"llm_temperature: 0.5\n" +
			"outbound_auto_call: true\n" +
			"sip_endpoint: sip_endpoint\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:aiagents", "update",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1AiagentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:aiagents", "list",
			"--page", "1",
			"--page-size", "1",
		)
	})
}

func TestAITestsTestFrameworkV1AiagentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:aiagents", "delete",
			"--id", "0",
		)
	})
}

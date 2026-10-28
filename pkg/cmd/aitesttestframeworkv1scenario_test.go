// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAITestsTestFrameworkV1ScenariosCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "create",
			"--name", "Refund escalation flow",
			"--personality", "4643",
			"--agent", "1259",
			"--expected-outcome-prompt", "The assistant explains the refund policy and offers escalation.",
			"--first-message", "Hi, I want a refund.",
			"--folder-path", "folder_path",
			"--instructions", "Call the support line and demand a refund for a damaged product.",
			"--metric", "266881",
			"--tag", "support",
			"--test-profile", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Refund escalation flow\n" +
			"personality: 4643\n" +
			"agent: 1259\n" +
			"expected_outcome_prompt: The assistant explains the refund policy and offers escalation.\n" +
			"first_message: Hi, I want a refund.\n" +
			"folder_path: folder_path\n" +
			"instructions: Call the support line and demand a refund for a damaged product.\n" +
			"metrics:\n" +
			"  - 266881\n" +
			"tags:\n" +
			"  - support\n" +
			"test_profile: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "create",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "retrieve",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "update",
			"--id", "0",
			"--agent", "0",
			"--expected-outcome-prompt", "expected_outcome_prompt",
			"--first-message", "first_message",
			"--folder-path", "folder_path",
			"--inbound-phone-number", "0",
			"--instructions", "Call the support line, demand a refund, and mention the damaged packaging.",
			"--max-duration", "0",
			"--metric", "266881",
			"--metric", "266882",
			"--name", "Refund escalation flow - updated",
			"--personality", "0",
			"--phone-number", "0",
			"--scenario-language", "scenario_language",
			"--tag", "string",
			"--test-profile", "0",
			"--tool-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent: 0\n" +
			"expected_outcome_prompt: expected_outcome_prompt\n" +
			"first_message: first_message\n" +
			"folder_path: folder_path\n" +
			"inbound_phone_number: 0\n" +
			"instructions: Call the support line, demand a refund, and mention the damaged packaging.\n" +
			"max_duration: 0\n" +
			"metrics:\n" +
			"  - 266881\n" +
			"  - 266882\n" +
			"name: Refund escalation flow - updated\n" +
			"personality: 0\n" +
			"phone_number: 0\n" +
			"scenario_language: scenario_language\n" +
			"tags:\n" +
			"  - string\n" +
			"test_profile: 0\n" +
			"tool_ids:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "update",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "list",
			"--agent-id", "0",
			"--assistant-id", "assistant_id",
			"--page", "1",
			"--page-size", "1",
			"--tags", "tags",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "delete",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosGenerateBg(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "generate-bg",
			"--num-scenarios", "5",
			"--agent-id", "1259",
			"--extra-instructions", "Focus on angry customers asking about refunds.",
			"--first-message", "first_message",
			"--folder-path", "folder_path",
			"--generate-expected-outcomes=true",
			"--inbound-phone-number", "0",
			"--personality", "0",
			"--scenario-type", "workflow",
			"--tag", "string",
			"--test-profile", "0",
			"--tool-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"num_scenarios: 5\n" +
			"agent_id: 1259\n" +
			"extra_instructions: Focus on angry customers asking about refunds.\n" +
			"first_message: first_message\n" +
			"folder_path: folder_path\n" +
			"generate_expected_outcomes: true\n" +
			"inbound_phone_number: 0\n" +
			"personalities:\n" +
			"  - 0\n" +
			"scenario_type: workflow\n" +
			"tags:\n" +
			"  - string\n" +
			"test_profile: 0\n" +
			"tool_ids:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "generate-bg",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosRetrieveExportScenarios(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "retrieve-export-scenarios",
			"--export-type", "csv",
			"--scenario-ids", "scenario_ids",
			"--agent-id", "0",
			"--output", "/dev/null",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosRetrieveGenerateProgress(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "retrieve-generate-progress",
			"--progress-id", "progress_id",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosRetrieveScenarioAgentProgress(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "retrieve-scenario-agent-progress",
			"--progress-id", "progress_id",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosRunScenarios(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "run-scenarios",
			"--agent-id", "1259",
			"--agent-number", "agent_number",
			"--assistant-id", "assistant_id",
			"--concurrency-limit", "3",
			"--frequency", "0",
			"--mode", "same_number",
			"--name", "name",
			"--outbound-phone-number", "outbound_phone_number",
			"--personality-id", "0",
			"--scenarios", "[9512]",
			"--tag", "regression",
			"--test-profile-id", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_id: 1259\n" +
			"agent_number: agent_number\n" +
			"assistant_id: assistant_id\n" +
			"concurrency_limit: 3\n" +
			"frequency: 0\n" +
			"mode: same_number\n" +
			"name: name\n" +
			"outbound_phone_number: outbound_phone_number\n" +
			"personality_ids:\n" +
			"  - 0\n" +
			"scenarios:\n" +
			"  - 9512\n" +
			"tags:\n" +
			"  - regression\n" +
			"test_profile_ids:\n" +
			"  - 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "run-scenarios",
		)
	})
}

func TestAITestsTestFrameworkV1ScenariosScenarioAgent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "scenario-agent",
			"--scenarios", "all",
			"--user-message", "Make the callers ask about damaged packaging too.",
			"--agent-id", "1259",
			"--agent-resolution-response", "agent_resolution_response",
			"--apply-changes=true",
			"--context", "[{content: content, role: role, timestamp: timestamp}]",
			"--knowledge-base", "0",
			"--num-scenarios", "0",
			"--pending-tool-call", "[{foo: bar}]",
			"--scenarios-needing-agent", "0",
			"--session-id", "0",
			"--updated-scenario", "{foo: bar}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiTestsTestFrameworkV1ScenariosScenarioAgent)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "scenario-agent",
			"--scenarios", "all",
			"--user-message", "Make the callers ask about damaged packaging too.",
			"--agent-id", "1259",
			"--agent-resolution-response", "agent_resolution_response",
			"--apply-changes=true",
			"--context.content", "content",
			"--context.role", "role",
			"--context.timestamp", "timestamp",
			"--knowledge-base", "0",
			"--num-scenarios", "0",
			"--pending-tool-call", "[{foo: bar}]",
			"--scenarios-needing-agent", "0",
			"--session-id", "0",
			"--updated-scenario", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"scenarios: all\n" +
			"user_message: Make the callers ask about damaged packaging too.\n" +
			"agent_id: 1259\n" +
			"agent_resolution_response: agent_resolution_response\n" +
			"apply_changes: true\n" +
			"context:\n" +
			"  - content: content\n" +
			"    role: role\n" +
			"    timestamp: timestamp\n" +
			"knowledge_base:\n" +
			"  - 0\n" +
			"num_scenarios: 0\n" +
			"pending_tool_calls:\n" +
			"  - foo: bar\n" +
			"scenarios_needing_agent:\n" +
			"  - 0\n" +
			"session_id: 0\n" +
			"updated_scenarios:\n" +
			"  - foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:scenarios", "scenario-agent",
		)
	})
}

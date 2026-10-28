// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAITestsTestFrameworkV1MetricsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "create",
			"--name", "Escalates on refund denial",
			"--add-to-new-agents=false",
			"--agent", "1259",
			"--assistant-id", "assistant_id",
			"--audio-enabled=true",
			"--description", "Returns a score for whether the assistant escalates when the caller's refund is denied.",
			"--display-order", "0",
			"--enum-value", "string",
			"--eval-type", "continuous_qualitative",
			"--observability-enabled=true",
			"--prompt", "Score the assistant's escalation behaviour between 0.0 and 1.0.",
			"--sampling-enabled=true",
			"--simulation-enabled=true",
			"--type", "llm_judge",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Escalates on refund denial\n" +
			"add_to_new_agents: false\n" +
			"agents:\n" +
			"  - 1259\n" +
			"assistant_id: assistant_id\n" +
			"audio_enabled: true\n" +
			"description: >-\n" +
			"  Returns a score for whether the assistant escalates when the caller's refund\n" +
			"  is denied.\n" +
			"display_order: 0\n" +
			"enum_values:\n" +
			"  - string\n" +
			"eval_type: continuous_qualitative\n" +
			"observability_enabled: true\n" +
			"prompt: Score the assistant's escalation behaviour between 0.0 and 1.0.\n" +
			"sampling_enabled: true\n" +
			"simulation_enabled: true\n" +
			"type: llm_judge\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "create",
		)
	})
}

func TestAITestsTestFrameworkV1MetricsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "retrieve",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1MetricsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "update",
			"--id", "0",
			"--add-to-new-agents=true",
			"--agent", "0",
			"--audio-enabled=true",
			"--description", "description",
			"--display-order", "0",
			"--enum-value", "string",
			"--eval-type", "eval_type",
			"--name", "Escalates on refund denial",
			"--observability-enabled=true",
			"--prompt", "Score the assistant's escalation behaviour between 0.0 and 1.0, considering tone.",
			"--sampling-enabled=true",
			"--simulation-enabled=true",
			"--type", "type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"add_to_new_agents: true\n" +
			"agents:\n" +
			"  - 0\n" +
			"audio_enabled: true\n" +
			"description: description\n" +
			"display_order: 0\n" +
			"enum_values:\n" +
			"  - string\n" +
			"eval_type: eval_type\n" +
			"name: Escalates on refund denial\n" +
			"observability_enabled: true\n" +
			"prompt: >-\n" +
			"  Score the assistant's escalation behaviour between 0.0 and 1.0, considering\n" +
			"  tone.\n" +
			"sampling_enabled: true\n" +
			"simulation_enabled: true\n" +
			"type: type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "update",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1MetricsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "list",
			"--agent-id", "0",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAITestsTestFrameworkV1MetricsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "delete",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1MetricsBulkCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "bulk-create",
			"--body", "{name: Greets within 5 seconds, add_to_new_agents: false, agents: [1259], assistant_id: assistant_id, audio_enabled: true, description: Returns whether the assistant greeted the caller within five seconds., display_order: 0, enum_values: [string], eval_type: continuous_qualitative, observability_enabled: true, prompt: Score whether the greeting was timely and natural., sampling_enabled: true, simulation_enabled: true, type: llm_judge}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiTestsTestFrameworkV1MetricsBulkCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "bulk-create",
			"--body.name", "Greets within 5 seconds",
			"--body.add-to-new-agents=false",
			"--body.agents", "[1259]",
			"--body.assistant-id", "assistant_id",
			"--body.audio-enabled=true",
			"--body.description", "Returns whether the assistant greeted the caller within five seconds.",
			"--body.display-order", "0",
			"--body.enum-values", "[string]",
			"--body.eval-type", "continuous_qualitative",
			"--body.observability-enabled=true",
			"--body.prompt", "Score whether the greeting was timely and natural.",
			"--body.sampling-enabled=true",
			"--body.simulation-enabled=true",
			"--body.type", "llm_judge",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- name: Greets within 5 seconds\n" +
			"  add_to_new_agents: false\n" +
			"  agents:\n" +
			"    - 1259\n" +
			"  assistant_id: assistant_id\n" +
			"  audio_enabled: true\n" +
			"  description: Returns whether the assistant greeted the caller within five seconds.\n" +
			"  display_order: 0\n" +
			"  enum_values:\n" +
			"    - string\n" +
			"  eval_type: continuous_qualitative\n" +
			"  observability_enabled: true\n" +
			"  prompt: Score whether the greeting was timely and natural.\n" +
			"  sampling_enabled: true\n" +
			"  simulation_enabled: true\n" +
			"  type: llm_judge\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "bulk-create",
		)
	})
}

func TestAITestsTestFrameworkV1MetricsGenerateMetrics(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "generate-metrics",
			"--agent-id", "1259",
			"--num-metrics", "5",
			"--scenario", "9512",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_id: 1259\n" +
			"num_metrics: 5\n" +
			"scenarios:\n" +
			"  - 9512\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "generate-metrics",
		)
	})
}

func TestAITestsTestFrameworkV1MetricsRetrieveGenerateMetricsProgress(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:metrics", "retrieve-generate-metrics-progress",
			"--progress-id", "progress_id",
		)
	})
}

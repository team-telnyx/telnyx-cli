// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAgentGatewaysKeysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:keys", "create",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--budget-resets-every", "30d",
			"--expires-in", "30d",
			"--max-budget", "20",
			"--max-parallel-requests", "4",
			"--model", "[GLM-5.3-Flash]",
			"--name", "support-agent",
			"--requests-per-minute", "60",
			"--team-id", "0b6c5d4e-3f2a-4b1c-9d8e-7f6a5b4c3d2e",
			"--tokens-per-minute", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"budget_resets_every: 30d\n" +
			"expires_in: 30d\n" +
			"max_budget: 20\n" +
			"max_parallel_requests: 4\n" +
			"models:\n" +
			"  - GLM-5.3-Flash\n" +
			"name: support-agent\n" +
			"requests_per_minute: 60\n" +
			"team_id: 0b6c5d4e-3f2a-4b1c-9d8e-7f6a5b4c3d2e\n" +
			"tokens_per_minute: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:keys", "create",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
		)
	})
}

func TestAgentGatewaysKeysUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:keys", "update",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--key-id", "5f1c3a0e9b7d4c2a8e6f0b1d3c5a7e9f2b4d6c8a0e1f3a5c7e9b1d3f5a7c9e0b",
			"--blocked=true",
			"--budget-resets-every", "30d",
			"--expires-in", "30d",
			"--max-budget", "0",
			"--max-parallel-requests", "1",
			"--model", "x",
			"--name", "x",
			"--requests-per-minute", "1",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--tokens-per-minute", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"blocked: true\n" +
			"budget_resets_every: 30d\n" +
			"expires_in: 30d\n" +
			"max_budget: 0\n" +
			"max_parallel_requests: 1\n" +
			"models:\n" +
			"  - x\n" +
			"name: x\n" +
			"requests_per_minute: 1\n" +
			"team_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"tokens_per_minute: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:keys", "update",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--key-id", "5f1c3a0e9b7d4c2a8e6f0b1d3c5a7e9f2b4d6c8a0e1f3a5c7e9b1d3f5a7c9e0b",
		)
	})
}

func TestAgentGatewaysKeysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:keys", "list",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAgentGatewaysKeysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:keys", "delete",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--key-id", "5f1c3a0e9b7d4c2a8e6f0b1d3c5a7e9f2b4d6c8a0e1f3a5c7e9b1d3f5a7c9e0b",
		)
	})
}

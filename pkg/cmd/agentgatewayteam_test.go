// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAgentGatewaysTeamsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:teams", "create",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--name", "support",
			"--budget-resets-every", "30d",
			"--max-budget", "50",
			"--model", "[GLM-5.3-Flash, Qwen3.8-27B-nothink]",
			"--requests-per-minute", "120",
			"--tokens-per-minute", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: support\n" +
			"budget_resets_every: 30d\n" +
			"max_budget: 50\n" +
			"models:\n" +
			"  - GLM-5.3-Flash\n" +
			"  - Qwen3.8-27B-nothink\n" +
			"requests_per_minute: 120\n" +
			"tokens_per_minute: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:teams", "create",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
		)
	})
}

func TestAgentGatewaysTeamsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:teams", "update",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--team-id", "0b6c5d4e-3f2a-4b1c-9d8e-7f6a5b4c3d2e",
			"--blocked=true",
			"--budget-resets-every", "30d",
			"--max-budget", "100",
			"--model", "x",
			"--name", "x",
			"--requests-per-minute", "1",
			"--tokens-per-minute", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"blocked: true\n" +
			"budget_resets_every: 30d\n" +
			"max_budget: 100\n" +
			"models:\n" +
			"  - x\n" +
			"name: x\n" +
			"requests_per_minute: 1\n" +
			"tokens_per_minute: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:teams", "update",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--team-id", "0b6c5d4e-3f2a-4b1c-9d8e-7f6a5b4c3d2e",
		)
	})
}

func TestAgentGatewaysTeamsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:teams", "list",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
		)
	})
}

func TestAgentGatewaysTeamsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:teams", "delete",
			"--instance-id", "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			"--team-id", "0b6c5d4e-3f2a-4b1c-9d8e-7f6a5b4c3d2e",
		)
	})
}

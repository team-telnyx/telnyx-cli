// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsRunsTelnyxAgentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:telnyx-agents", "list",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsTelnyxAgentsLink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:telnyx-agents", "link",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--telnyx-agent-id", "telnyx_agent_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("telnyx_agent_id: telnyx_agent_id")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:missions:runs:telnyx-agents", "link",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsTelnyxAgentsUnlink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:telnyx-agents", "unlink",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--telnyx-agent-id", "telnyx_agent_id",
		)
	})
}

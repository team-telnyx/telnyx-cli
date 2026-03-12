// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsToolsCreateTool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:tools", "create-tool",
			"--api-key", "string",
			"--mission-id", "mission_id",
		)
	})
}

func TestAIMissionsToolsDeleteTool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:tools", "delete-tool",
			"--api-key", "string",
			"--mission-id", "mission_id",
			"--tool-id", "tool_id",
		)
	})
}

func TestAIMissionsToolsGetTool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:tools", "get-tool",
			"--api-key", "string",
			"--mission-id", "mission_id",
			"--tool-id", "tool_id",
		)
	})
}

func TestAIMissionsToolsListTools(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:tools", "list-tools",
			"--api-key", "string",
			"--mission-id", "mission_id",
		)
	})
}

func TestAIMissionsToolsUpdateTool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:tools", "update-tool",
			"--api-key", "string",
			"--mission-id", "mission_id",
			"--tool-id", "tool_id",
		)
	})
}

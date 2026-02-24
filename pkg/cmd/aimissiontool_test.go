// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsToolsCreateTool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:tools", "create-tool",
		"--mission-id", "mission_id",
	)
}

func TestAIMissionsToolsDeleteTool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:tools", "delete-tool",
		"--mission-id", "mission_id",
		"--tool-id", "tool_id",
	)
}

func TestAIMissionsToolsGetTool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:tools", "get-tool",
		"--mission-id", "mission_id",
		"--tool-id", "tool_id",
	)
}

func TestAIMissionsToolsListTools(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:tools", "list-tools",
		"--mission-id", "mission_id",
	)
}

func TestAIMissionsToolsUpdateTool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:tools", "update-tool",
		"--mission-id", "mission_id",
		"--tool-id", "tool_id",
	)
}

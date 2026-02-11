// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsRunsTelnyxAgentsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs:telnyx-agents", "list",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIMissionsRunsTelnyxAgentsLink(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs:telnyx-agents", "link",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--telnyx-agent-id", "telnyx_agent_id",
	)
}

func TestAIMissionsRunsTelnyxAgentsUnlink(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs:telnyx-agents", "unlink",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--telnyx-agent-id", "telnyx_agent_id",
	)
}

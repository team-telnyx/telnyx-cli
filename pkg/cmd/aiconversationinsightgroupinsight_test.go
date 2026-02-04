// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIConversationsInsightGroupsInsightsAssign(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insight-groups:insights", "assign",
		"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIConversationsInsightGroupsInsightsDeleteUnassign(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insight-groups:insights", "delete-unassign",
		"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

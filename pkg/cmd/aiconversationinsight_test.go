// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIConversationsInsightsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insights", "create",
		"--instructions", "instructions",
		"--name", "name",
		"--json-schema", "string",
		"--webhook", "webhook",
	)
}

func TestAIConversationsInsightsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insights", "retrieve",
		"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIConversationsInsightsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insights", "update",
		"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--instructions", "instructions",
		"--json-schema", "string",
		"--name", "name",
		"--webhook", "webhook",
	)
}

func TestAIConversationsInsightsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insights", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestAIConversationsInsightsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insights", "delete",
		"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

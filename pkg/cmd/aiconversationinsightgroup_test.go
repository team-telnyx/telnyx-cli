// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestAIConversationsInsightGroupsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insight-groups", "retrieve",
		"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIConversationsInsightGroupsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insight-groups", "update",
		"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--description", "description",
		"--name", "name",
		"--webhook", "webhook",
	)
}

func TestAIConversationsInsightGroupsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insight-groups", "delete",
		"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIConversationsInsightGroupsInsightGroups(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insight-groups", "insight-groups",
		"--name", "name",
		"--description", "description",
		"--webhook", "webhook",
	)
}

func TestAIConversationsInsightGroupsRetrieveInsightGroups(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations:insight-groups", "retrieve-insight-groups",
		"--page-number", "0",
		"--page-size", "0",
	)
}

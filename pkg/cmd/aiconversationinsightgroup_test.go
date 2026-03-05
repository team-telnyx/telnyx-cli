// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIConversationsInsightGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insight-groups", "retrieve",
			"--api-key", "string",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsInsightGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insight-groups", "update",
			"--api-key", "string",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--description", "description",
			"--name", "name",
			"--webhook", "webhook",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"name: name\n" +
			"webhook: webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:conversations:insight-groups", "update",
			"--api-key", "string",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsInsightGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insight-groups", "delete",
			"--api-key", "string",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsInsightGroupsInsightGroups(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insight-groups", "insight-groups",
			"--api-key", "string",
			"--name", "name",
			"--description", "description",
			"--webhook", "webhook",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"description: description\n" +
			"webhook: webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:conversations:insight-groups", "insight-groups",
			"--api-key", "string",
		)
	})
}

func TestAIConversationsInsightGroupsRetrieveInsightGroups(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insight-groups", "retrieve-insight-groups",
			"--api-key", "string",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

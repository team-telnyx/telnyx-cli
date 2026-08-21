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
			t,
			"--api-key", "string",
			"ai:conversations:insight-groups", "retrieve",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsInsightGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:conversations:insight-groups", "update",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--description", "Description",
			"--name", "Name",
			"--webhook", "Webhook",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: Description\n" +
			"name: Name\n" +
			"webhook: Webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:conversations:insight-groups", "update",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsInsightGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:conversations:insight-groups", "delete",
			"--group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsInsightGroupsInsightGroups(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:conversations:insight-groups", "insight-groups",
			"--name", "Name",
			"--description", "Description",
			"--webhook", "",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Name\n" +
			"description: Description\n" +
			"webhook: ''\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:conversations:insight-groups", "insight-groups",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})
}

func TestAIConversationsInsightGroupsRetrieveInsightGroups(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:conversations:insight-groups", "retrieve-insight-groups",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIConversationsInsightsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insights", "create",
			"--api-key", "string",
			"--instructions", "instructions",
			"--name", "name",
			"--json-schema", "string",
			"--webhook", "webhook",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: instructions\n" +
			"name: name\n" +
			"json_schema: string\n" +
			"webhook: webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:conversations:insights", "create",
			"--api-key", "string",
		)
	})
}

func TestAIConversationsInsightsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insights", "retrieve",
			"--api-key", "string",
			"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsInsightsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insights", "update",
			"--api-key", "string",
			"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--instructions", "instructions",
			"--json-schema", "string",
			"--name", "name",
			"--webhook", "webhook",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: instructions\n" +
			"json_schema: string\n" +
			"name: name\n" +
			"webhook: webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:conversations:insights", "update",
			"--api-key", "string",
			"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsInsightsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insights", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestAIConversationsInsightsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations:insights", "delete",
			"--api-key", "string",
			"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

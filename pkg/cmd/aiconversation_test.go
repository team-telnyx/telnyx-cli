// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIConversationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations", "create",
			"--api-key", "string",
			"--metadata", "{foo: string}",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  foo: string\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:conversations", "create",
			"--api-key", "string",
		)
	})
}

func TestAIConversationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations", "retrieve",
			"--api-key", "string",
			"--conversation-id", "conversation_id",
		)
	})
}

func TestAIConversationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations", "update",
			"--api-key", "string",
			"--conversation-id", "conversation_id",
			"--metadata", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:conversations", "update",
			"--api-key", "string",
			"--conversation-id", "conversation_id",
		)
	})
}

func TestAIConversationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations", "list",
			"--api-key", "string",
			"--id", "id",
			"--created-at", "created_at",
			"--last-message-at", "last_message_at",
			"--limit", "1",
			"--metadata-assistant-id", "metadata->assistant_id",
			"--metadata-call-control-id", "metadata->call_control_id",
			"--metadata-telnyx-agent-target", "metadata->telnyx_agent_target",
			"--metadata-telnyx-conversation-channel", "metadata->telnyx_conversation_channel",
			"--metadata-telnyx-end-user-target", "metadata->telnyx_end_user_target",
			"--name", "name",
			"--or", "or",
			"--order", "order",
		)
	})
}

func TestAIConversationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations", "delete",
			"--api-key", "string",
			"--conversation-id", "conversation_id",
		)
	})
}

func TestAIConversationsAddMessage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations", "add-message",
			"--api-key", "string",
			"--conversation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--role", "role",
			"--content", "content",
			"--metadata", "{foo: string}",
			"--name", "name",
			"--sent-at", "'2019-12-27T18:11:19.117Z'",
			"--tool-call-id", "tool_call_id",
			"--tool-call", "{foo: bar}",
			"--tool-choice", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"role: role\n" +
			"content: content\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"name: name\n" +
			"sent_at: '2019-12-27T18:11:19.117Z'\n" +
			"tool_call_id: tool_call_id\n" +
			"tool_calls:\n" +
			"  - foo: bar\n" +
			"tool_choice: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:conversations", "add-message",
			"--api-key", "string",
			"--conversation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIConversationsRetrieveConversationsInsights(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:conversations", "retrieve-conversations-insights",
			"--api-key", "string",
			"--conversation-id", "conversation_id",
		)
	})
}

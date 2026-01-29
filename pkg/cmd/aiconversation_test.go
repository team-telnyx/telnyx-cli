// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestAIConversationsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations", "create",
		"--metadata", "{foo: string}",
		"--name", "name",
	)
}

func TestAIConversationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations", "retrieve",
		"--conversation-id", "conversation_id",
	)
}

func TestAIConversationsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations", "update",
		"--conversation-id", "conversation_id",
		"--metadata", "{foo: string}",
	)
}

func TestAIConversationsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations", "list",
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
}

func TestAIConversationsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations", "delete",
		"--conversation-id", "conversation_id",
	)
}

func TestAIConversationsAddMessage(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations", "add-message",
		"--conversation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--role", "role",
		"--content", "content",
		"--metadata", "{foo: string}",
		"--name", "name",
		"--sent-at", "2019-12-27T18:11:19.117Z",
		"--tool-call-id", "tool_call_id",
		"--tool-call", "{foo: bar}",
		"--tool-choice", "string",
	)
}

func TestAIConversationsRetrieveConversationsInsights(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:conversations", "retrieve-conversations-insights",
		"--conversation-id", "conversation_id",
	)
}

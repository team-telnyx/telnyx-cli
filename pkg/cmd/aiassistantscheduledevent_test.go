// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAssistantsScheduledEventsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:scheduled-events", "create",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--scheduled-at-fixed-datetime", "'2025-04-15T13:07:28.764Z'",
			"--telnyx-agent-target", "telnyx_agent_target",
			"--telnyx-conversation-channel", "phone_call",
			"--telnyx-end-user-target", "telnyx_end_user_target",
			"--conversation-metadata", "{foo: string}",
			"--dynamic-variables", "{foo: string}",
			"--text", "text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"scheduled_at_fixed_datetime: '2025-04-15T13:07:28.764Z'\n" +
			"telnyx_agent_target: telnyx_agent_target\n" +
			"telnyx_conversation_channel: phone_call\n" +
			"telnyx_end_user_target: telnyx_end_user_target\n" +
			"conversation_metadata:\n" +
			"  foo: string\n" +
			"dynamic_variables:\n" +
			"  foo: string\n" +
			"text: text\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:assistants:scheduled-events", "create",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAIAssistantsScheduledEventsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:scheduled-events", "retrieve",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--event-id", "event_id",
		)
	})
}

func TestAIAssistantsScheduledEventsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:scheduled-events", "list",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--conversation-channel", "phone_call",
			"--from-date", "'2019-12-27T18:11:19.117Z'",
			"--page-number", "0",
			"--page-size", "0",
			"--to-date", "'2019-12-27T18:11:19.117Z'",
		)
	})
}

func TestAIAssistantsScheduledEventsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:scheduled-events", "delete",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--event-id", "event_id",
		)
	})
}

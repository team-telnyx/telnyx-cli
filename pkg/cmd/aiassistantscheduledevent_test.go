// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAssistantsScheduledEventsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:scheduled-events", "create",
		"--assistant-id", "assistant_id",
		"--scheduled-at-fixed-datetime", "2025-04-15T13:07:28.764Z",
		"--telnyx-agent-target", "telnyx_agent_target",
		"--telnyx-conversation-channel", "phone_call",
		"--telnyx-end-user-target", "telnyx_end_user_target",
		"--conversation-metadata", "{foo: string}",
		"--text", "text",
	)
}

func TestAIAssistantsScheduledEventsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:scheduled-events", "retrieve",
		"--assistant-id", "assistant_id",
		"--event-id", "event_id",
	)
}

func TestAIAssistantsScheduledEventsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:scheduled-events", "list",
		"--assistant-id", "assistant_id",
		"--conversation-channel", "phone_call",
		"--from-date", "2019-12-27T18:11:19.117Z",
		"--page-number", "0",
		"--page-size", "0",
		"--to-date", "2019-12-27T18:11:19.117Z",
	)
}

func TestAIAssistantsScheduledEventsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:scheduled-events", "delete",
		"--assistant-id", "assistant_id",
		"--event-id", "event_id",
	)
}

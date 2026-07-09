// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIConversationsConversationInsightsRetrieveAggregates(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:conversations:conversation-insights", "retrieve-aggregates",
			"--created-at", "created_at",
			"--group-by", "string",
			"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata", "{assistant_id: assistant_id}",
			"--show", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiConversationsConversationInsightsRetrieveAggregates)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:conversations:conversation-insights", "retrieve-aggregates",
			"--created-at", "created_at",
			"--group-by", "string",
			"--insight-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata.assistant-id", "assistant_id",
			"--show", "string",
		)
	})
}

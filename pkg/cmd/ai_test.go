// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIRetrieveConversationHistories(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "retrieve-conversation-histories",
			"--q", "customer called about billing issue",
			"--filter-ingested-at-gte", "'2026-01-01T00:00:00Z'",
			"--filter-ingested-at-lte", "'2026-12-31T23:59:59Z'",
			"--filter-record-created-at-gte", "'2026-01-01T00:00:00Z'",
			"--filter-record-created-at-lte", "'2026-12-31T23:59:59Z'",
			"--filter-record-id", "rec-001",
			"--filter-region-in", "USA,DEU",
			"--filter-retention", "filter[retention]",
			"--filter-user-id", "user-123",
			"--min-score", "0.5",
			"--page-number", "1",
			"--page-size", "10",
			"--region", "USA",
		)
	})
}

func TestAISummarize(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai", "summarize",
			"--bucket", "string",
			"--filename", "string",
			"--system-prompt", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket: string\n" +
			"filename: string\n" +
			"system_prompt: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai", "summarize",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIKnowledgeCollectionsRetrieveDocuments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:knowledge:collections", "retrieve-documents",
			"--slug", "support-transcripts",
			"--filter", "{foo: bar}",
			"--page-number", "1",
			"--page-size", "20",
			"--query", "customer called about billing issue",
			"--retrieval-type", "vector",
			"--sources", "voice,message",
			"--top-k", "10",
		)
	})
}

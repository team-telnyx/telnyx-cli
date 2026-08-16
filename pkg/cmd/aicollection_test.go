// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAICollectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections", "create",
			"--name", "Support Transcripts",
			"--description", "All customer support voice transcripts.",
			"--settings", "{retrieval: {retrieval_type: vector, top_k: 5}}",
			"--slug", "support-transcripts",
			"--source", "{source_type: voice, bucket_id: policy-docs}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiCollectionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections", "create",
			"--name", "Support Transcripts",
			"--description", "All customer support voice transcripts.",
			"--settings.retrieval", "{retrieval_type: vector, top_k: 5}",
			"--slug", "support-transcripts",
			"--source.source-type", "voice",
			"--source.bucket-id", "policy-docs",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Support Transcripts\n" +
			"description: All customer support voice transcripts.\n" +
			"settings:\n" +
			"  retrieval:\n" +
			"    retrieval_type: vector\n" +
			"    top_k: 5\n" +
			"slug: support-transcripts\n" +
			"sources:\n" +
			"  - source_type: voice\n" +
			"    bucket_id: policy-docs\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:collections", "create",
		)
	})
}

func TestAICollectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections", "retrieve",
			"--slug", "support-transcripts",
		)
	})
}

func TestAICollectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections", "update",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--description", "Updated description.",
			"--name", "Support Transcripts (2026)",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: Updated description.\n" +
			"name: Support Transcripts (2026)\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:collections", "update",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}

func TestAICollectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "20",
		)
	})
}

func TestAICollectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections", "delete",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}

func TestAICollectionsRetrieveByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections", "retrieve-by-id",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}

func TestAICollectionsRetrieveDocuments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections", "retrieve-documents",
			"--slug", "support-transcripts",
			"--filter", "{foo: bar}",
			"--page-number", "1",
			"--page-size", "20",
			"--query", "customer called about billing issue",
			"--retrieval-type", "hybrid",
			"--sources", "voice,message",
			"--top-k", "10",
		)
	})
}

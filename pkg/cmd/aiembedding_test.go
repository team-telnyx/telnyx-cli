// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIEmbeddingsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:embeddings", "create",
			"--bucket-name", "Bucket Name",
			"--document-chunk-overlap-size", "512",
			"--document-chunk-size", "1024",
			"--embedding-model", "thenlper/gte-large",
			"--loader", "default",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket_name: Bucket Name\n" +
			"document_chunk_overlap_size: 512\n" +
			"document_chunk_size: 1024\n" +
			"embedding_model: thenlper/gte-large\n" +
			"loader: default\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:embeddings", "create",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})
}

func TestAIEmbeddingsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:embeddings", "retrieve",
			"--task-id", "task_id",
		)
	})
}

func TestAIEmbeddingsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:embeddings", "list",
			"--status", "string",
		)
	})
}

func TestAIEmbeddingsSimilaritySearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:embeddings", "similarity-search",
			"--bucket-name", "Bucket Name",
			"--query", "Query",
			"--num-of-docs", "3",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket_name: Bucket Name\n" +
			"query: Query\n" +
			"num_of_docs: 3\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:embeddings", "similarity-search",
		)
	})
}

func TestAIEmbeddingsURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:embeddings", "url",
			"--bucket-name", "Bucket Name",
			"--url", "URL",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket_name: Bucket Name\n" +
			"url: URL\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:embeddings", "url",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})
}

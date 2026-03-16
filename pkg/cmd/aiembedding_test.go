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
			"--bucket-name", "bucket_name",
			"--document-chunk-overlap-size", "0",
			"--document-chunk-size", "0",
			"--embedding-model", "thenlper/gte-large",
			"--loader", "default",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket_name: bucket_name\n" +
			"document_chunk_overlap_size: 0\n" +
			"document_chunk_size: 0\n" +
			"embedding_model: thenlper/gte-large\n" +
			"loader: default\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:embeddings", "create",
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
			"--bucket-name", "bucket_name",
			"--query", "query",
			"--num-of-docs", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket_name: bucket_name\n" +
			"query: query\n" +
			"num_of_docs: 0\n")
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
			"--bucket-name", "bucket_name",
			"--url", "url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket_name: bucket_name\n" +
			"url: url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:embeddings", "url",
		)
	})
}

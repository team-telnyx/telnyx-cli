// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIEmbeddingsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:embeddings", "create",
		"--bucket-name", "bucket_name",
		"--document-chunk-overlap-size", "0",
		"--document-chunk-size", "0",
		"--embedding-model", "thenlper/gte-large",
		"--loader", "default",
	)
}

func TestAIEmbeddingsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:embeddings", "retrieve",
		"--task-id", "task_id",
	)
}

func TestAIEmbeddingsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:embeddings", "list",
		"--status", "string",
	)
}

func TestAIEmbeddingsSimilaritySearch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:embeddings", "similarity-search",
		"--bucket-name", "bucket_name",
		"--query", "query",
		"--num-of-docs", "0",
	)
}

func TestAIEmbeddingsURL(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:embeddings", "url",
		"--bucket-name", "bucket_name",
		"--url", "url",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIOpenAIEmbeddingsCreateEmbeddings(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:openai:embeddings", "create-embeddings",
		"--input", "The quick brown fox jumps over the lazy dog",
		"--model", "thenlper/gte-large",
		"--dimensions", "0",
		"--encoding-format", "float",
		"--user", "user",
	)
}

func TestAIOpenAIEmbeddingsListEmbeddingModels(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:openai:embeddings", "list-embedding-models",
	)
}

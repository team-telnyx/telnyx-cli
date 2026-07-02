// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIOpenAIEmbeddingsCreateEmbeddings(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:openai:embeddings", "create-embeddings",
			"--input", "The quick brown fox jumps over the lazy dog",
			"--model", "thenlper/gte-large",
			"--dimensions", "0",
			"--encoding-format", "float",
			"--user", "user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input: The quick brown fox jumps over the lazy dog\n" +
			"model: thenlper/gte-large\n" +
			"dimensions: 0\n" +
			"encoding_format: float\n" +
			"user: user\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:openai:embeddings", "create-embeddings",
		)
	})
}

func TestAIOpenAIEmbeddingsListEmbeddingModels(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:openai:embeddings", "list-embedding-models",
		)
	})
}

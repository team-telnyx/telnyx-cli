// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIEmbeddingsBucketsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:embeddings:buckets", "retrieve",
			"--api-key", "string",
			"--bucket-name", "bucket_name",
		)
	})
}

func TestAIEmbeddingsBucketsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:embeddings:buckets", "list",
			"--api-key", "string",
		)
	})
}

func TestAIEmbeddingsBucketsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:embeddings:buckets", "delete",
			"--api-key", "string",
			"--bucket-name", "bucket_name",
		)
	})
}

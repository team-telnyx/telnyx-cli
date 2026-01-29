// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestAIEmbeddingsBucketsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:embeddings:buckets", "retrieve",
		"--bucket-name", "bucket_name",
	)
}

func TestAIEmbeddingsBucketsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:embeddings:buckets", "list",
	)
}

func TestAIEmbeddingsBucketsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:embeddings:buckets", "delete",
		"--bucket-name", "bucket_name",
	)
}

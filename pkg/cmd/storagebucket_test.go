// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestStorageBucketsCreatePresignedURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "storage:buckets", "create-presigned-url",
			"--api-key", "string",
			"--bucket-name", "",
			"--object-name", "",
			"--ttl", "60",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("ttl: 60")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "storage:buckets", "create-presigned-url",
			"--api-key", "string",
			"--bucket-name", "",
			"--object-name", "",
		)
	})
}

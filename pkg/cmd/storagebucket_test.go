// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestStorageBucketsCreatePresignedURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:buckets", "create-presigned-url",
		"--bucket-name", "",
		"--object-name", "",
		"--ttl", "60",
	)
}

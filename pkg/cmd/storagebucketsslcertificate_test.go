// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestStorageBucketsSslCertificateCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:buckets:ssl-certificate", "create",
		"--api-key", "string",
		"--bucket-name", "",
		"--certificate", "...",
		"--private-key", "...",
	)
}

func TestStorageBucketsSslCertificateRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:buckets:ssl-certificate", "retrieve",
		"--api-key", "string",
		"--bucket-name", "",
	)
}

func TestStorageBucketsSslCertificateDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:buckets:ssl-certificate", "delete",
		"--api-key", "string",
		"--bucket-name", "",
	)
}

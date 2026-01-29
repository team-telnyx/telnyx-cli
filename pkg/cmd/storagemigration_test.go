// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestStorageMigrationsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:migrations", "create",
		"--source-id", "source_id",
		"--target-bucket-name", "target_bucket_name",
		"--target-region", "target_region",
		"--refresh=true",
	)
}

func TestStorageMigrationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:migrations", "retrieve",
		"--id", "",
	)
}

func TestStorageMigrationsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:migrations", "list",
	)
}

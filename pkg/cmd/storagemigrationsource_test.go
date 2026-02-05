// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestStorageMigrationSourcesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:migration-sources", "create",
		"--bucket-name", "bucket_name",
		"--provider", "aws",
		"--provider-auth", "{access_key: access_key, secret_access_key: secret_access_key}",
		"--source-region", "source_region",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(storageMigrationSourcesCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:migration-sources", "create",
		"--bucket-name", "bucket_name",
		"--provider", "aws",
		"--provider-auth.access-key", "access_key",
		"--provider-auth.secret-access-key", "secret_access_key",
		"--source-region", "source_region",
	)
}

func TestStorageMigrationSourcesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:migration-sources", "retrieve",
		"--id", "",
	)
}

func TestStorageMigrationSourcesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:migration-sources", "list",
	)
}

func TestStorageMigrationSourcesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:migration-sources", "delete",
		"--id", "",
	)
}

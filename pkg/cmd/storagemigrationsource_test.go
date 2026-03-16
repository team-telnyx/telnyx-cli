// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestStorageMigrationSourcesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:migration-sources", "create",
			"--bucket-name", "bucket_name",
			"--provider", "aws",
			"--provider-auth", "{access_key: access_key, secret_access_key: secret_access_key}",
			"--source-region", "source_region",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(storageMigrationSourcesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:migration-sources", "create",
			"--bucket-name", "bucket_name",
			"--provider", "aws",
			"--provider-auth.access-key", "access_key",
			"--provider-auth.secret-access-key", "secret_access_key",
			"--source-region", "source_region",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket_name: bucket_name\n" +
			"provider: aws\n" +
			"provider_auth:\n" +
			"  access_key: access_key\n" +
			"  secret_access_key: secret_access_key\n" +
			"source_region: source_region\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"storage:migration-sources", "create",
		)
	})
}

func TestStorageMigrationSourcesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:migration-sources", "retrieve",
			"--id", "",
		)
	})
}

func TestStorageMigrationSourcesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:migration-sources", "list",
		)
	})
}

func TestStorageMigrationSourcesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:migration-sources", "delete",
			"--id", "",
		)
	})
}

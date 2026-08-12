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
			"--bucket-name", "string",
			"--provider", "aws",
			"--provider-auth", "{access_key: string, secret_access_key: string}",
			"--source-region", "string",
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
			"--bucket-name", "string",
			"--provider", "aws",
			"--provider-auth.access-key", "string",
			"--provider-auth.secret-access-key", "string",
			"--source-region", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket_name: string\n" +
			"provider: aws\n" +
			"provider_auth:\n" +
			"  access_key: string\n" +
			"  secret_access_key: string\n" +
			"source_region: string\n")
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

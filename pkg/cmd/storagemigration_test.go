// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestStorageMigrationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "storage:migrations", "create",
			"--api-key", "string",
			"--source-id", "source_id",
			"--target-bucket-name", "target_bucket_name",
			"--target-region", "target_region",
			"--refresh=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"source_id: source_id\n" +
			"target_bucket_name: target_bucket_name\n" +
			"target_region: target_region\n" +
			"refresh: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "storage:migrations", "create",
			"--api-key", "string",
		)
	})
}

func TestStorageMigrationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "storage:migrations", "retrieve",
			"--api-key", "string",
			"--id", "",
		)
	})
}

func TestStorageMigrationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "storage:migrations", "list",
			"--api-key", "string",
		)
	})
}

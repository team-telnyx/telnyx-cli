// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestStorageKvsKeysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:kvs:keys", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--key", "key",
			"--output", "/dev/null",
		)
	})
}

func TestStorageKvsKeysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:kvs:keys", "list",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--cursor", "cursor",
			"--limit", "1",
			"--prefix", "prefix",
		)
	})
}

func TestStorageKvsKeysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:kvs:keys", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--key", "key",
		)
	})
}

func TestStorageKvsKeysSet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:kvs:keys", "set",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--key", "key",
			"--body", mocktest.TestFile(t, "Example data"),
			"--ttl-secs", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("Example data")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"storage:kvs:keys", "set",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--key", "key",
			"--ttl-secs", "1",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestStorageSqldbsActionsQuery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:sqldbs:actions", "query",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sql", "SELECT * FROM users WHERE name = ?",
			"--param", "alice",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sql: SELECT * FROM users WHERE name = ?\n" +
			"params:\n" +
			"  - alice\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"storage:sqldbs:actions", "query",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

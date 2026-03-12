// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestRecordingsActionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "recordings:actions", "delete",
			"--api-key", "string",
			"--id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c1",
			"--id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ids:\n" +
			"  - 428c31b6-7af4-4bcb-b7f5-5013ef9657c1\n" +
			"  - 428c31b6-7af4-4bcb-b7f5-5013ef9657c2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "recordings:actions", "delete",
			"--api-key", "string",
		)
	})
}

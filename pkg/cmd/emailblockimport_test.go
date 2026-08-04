// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailBlocksImportCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-blocks:import", "create",
			"--file", mocktest.TestFile(t, "Example data"),
			"--block-ttl-days", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"file: Example data\n" +
			"block_ttl_days: 1\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-blocks:import", "create",
		)
	})
}

func TestEmailBlocksImportRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-blocks:import", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

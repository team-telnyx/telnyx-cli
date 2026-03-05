// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPortingOrdersCommentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:comments", "create",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body", "Please, let me know when the port completes",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("body: Please, let me know when the port completes")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting-orders:comments", "create",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingOrdersCommentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:comments", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

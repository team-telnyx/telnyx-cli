// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestDirCommentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:comments", "create",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--content", "Re-uploaded the certificate. New document_id: 89450109-ee35-411c-b5bb-14f1d806fca2.",
			"--parent-comment-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: >-\n" +
			"  Re-uploaded the certificate. New document_id:\n" +
			"  89450109-ee35-411c-b5bb-14f1d806fca2.\n" +
			"parent_comment_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir:comments", "create",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirCommentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:comments", "list",
			"--max-items", "10",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--comment-type", "vetting_comment",
			"--page-number", "1",
			"--page-size", "20",
		)
	})
}

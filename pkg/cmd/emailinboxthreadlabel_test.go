// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailInboxesThreadsLabelsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:threads:labels", "create",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--label", "needs_review",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"labels:\n" +
			"  - needs_review\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:threads:labels", "create",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesThreadsLabelsDeleteAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:threads:labels", "delete-all",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--label", "needs_review",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"labels:\n" +
			"  - needs_review\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:threads:labels", "delete-all",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

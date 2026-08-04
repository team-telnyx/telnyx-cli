// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailInboxesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes", "create",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--username", "username",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domain_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"username: username\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes", "create",
		)
	})
}

func TestEmailInboxesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes", "list",
			"--page-cursor", "page_cursor",
			"--page-size", "1",
		)
	})
}

func TestEmailInboxesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

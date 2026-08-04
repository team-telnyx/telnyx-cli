// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailInboxesFiltersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:filters", "list",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesFiltersAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:filters", "add",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--entry", "@spam.example",
			"--type", "blocklist",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entries:\n" +
			"  - '@spam.example'\n" +
			"type: blocklist\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:filters", "add",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesFiltersDeleteAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:filters", "delete-all",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--entry", "former-partner@example.com",
			"--type", "allowlist",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entries:\n" +
			"  - former-partner@example.com\n" +
			"type: allowlist\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:filters", "delete-all",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailInboxesFiltersReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-inboxes:filters", "replace",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--allowlist", "trusted@example.com",
			"--allowlist", "@partner.example",
			"--blocklist", "@spam.example",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowlist:\n" +
			"  - trusted@example.com\n" +
			"  - '@partner.example'\n" +
			"blocklist:\n" +
			"  - '@spam.example'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-inboxes:filters", "replace",
			"--inbox-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

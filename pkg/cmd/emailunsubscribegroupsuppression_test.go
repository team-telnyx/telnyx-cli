// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailUnsubscribeGroupsSuppressionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-unsubscribe-groups:suppressions", "create",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--to", "to",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("to: to")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-unsubscribe-groups:suppressions", "create",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailUnsubscribeGroupsSuppressionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-unsubscribe-groups:suppressions", "list",
			"--max-items", "10",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestEmailUnsubscribeGroupsSuppressionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-unsubscribe-groups:suppressions", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--email", "email",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAlphanumericSenderIDsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "alphanumeric-sender-ids", "create",
			"--api-key", "string",
			"--alphanumeric-sender-id", "MyCompany",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--us-long-code-fallback", "+15551234567",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"alphanumeric_sender_id: MyCompany\n" +
			"messaging_profile_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"us_long_code_fallback: '+15551234567'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "alphanumeric-sender-ids", "create",
			"--api-key", "string",
		)
	})
}

func TestAlphanumericSenderIDsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "alphanumeric-sender-ids", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAlphanumericSenderIDsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "alphanumeric-sender-ids", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter-messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestAlphanumericSenderIDsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "alphanumeric-sender-ids", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

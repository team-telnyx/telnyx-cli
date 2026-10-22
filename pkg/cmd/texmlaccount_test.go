// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsRetrieveRecordingsJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts", "retrieve-recordings-json",
			"--account-sid", "f5586561-8ff0-4291-a51b-c7dfe4139ff9",
			"--date-created", "'2023-05-22T00:00:00Z'",
			"--page", "0",
			"--page-size", "0",
		)
	})
}

func TestTexmlAccountsRetrieveTranscriptionsJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts", "retrieve-transcriptions-json",
			"--account-sid", "f5586561-8ff0-4291-a51b-c7dfe4139ff9",
			"--page-size", "0",
			"--page-token", "PageToken",
		)
	})
}

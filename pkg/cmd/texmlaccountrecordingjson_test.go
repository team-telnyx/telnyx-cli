// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsRecordingsJsonDeleteRecordingSidJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:recordings:json", "delete-recording-sid-json",
			"--account-sid", "f5586561-8ff0-4291-a51b-c7dfe4139ff9",
			"--recording-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestTexmlAccountsRecordingsJsonRetrieveRecordingSidJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:recordings:json", "retrieve-recording-sid-json",
			"--account-sid", "f5586561-8ff0-4291-a51b-c7dfe4139ff9",
			"--recording-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

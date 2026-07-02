// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsConferencesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences", "retrieve",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
		)
	})
}

func TestTexmlAccountsConferencesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences", "update",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
			"--announce-method", "GET",
			"--announce-url", "https://www.example.com/announce.xml",
			"--status", "completed",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"AnnounceMethod: GET\n" +
			"AnnounceUrl: https://www.example.com/announce.xml\n" +
			"Status: completed\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:conferences", "update",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
		)
	})
}

func TestTexmlAccountsConferencesRetrieveConferences(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences", "retrieve-conferences",
			"--account-sid", "account_sid",
			"--date-created", "DateCreated",
			"--date-updated", "DateUpdated",
			"--friendly-name", "FriendlyName",
			"--page", "0",
			"--page-size", "0",
			"--page-token", "PageToken",
			"--status", "init",
		)
	})
}

func TestTexmlAccountsConferencesRetrieveRecordings(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences", "retrieve-recordings",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
		)
	})
}

func TestTexmlAccountsConferencesRetrieveRecordingsJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences", "retrieve-recordings-json",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
		)
	})
}

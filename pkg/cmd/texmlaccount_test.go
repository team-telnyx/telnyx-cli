// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsRetrieveRecordingsJson(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts", "retrieve-recordings-json",
		"--account-sid", "account_sid",
		"--date-created", "2023-05-22T00:00:00Z",
		"--page", "0",
		"--page-size", "0",
	)
}

func TestTexmlAccountsRetrieveTranscriptionsJson(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts", "retrieve-transcriptions-json",
		"--account-sid", "account_sid",
		"--page-size", "0",
		"--page-token", "PageToken",
	)
}

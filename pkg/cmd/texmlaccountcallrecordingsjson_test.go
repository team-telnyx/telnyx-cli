// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsCallsRecordingsJsonRecordingsJson(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:calls:recordings-json", "recordings-json",
		"--account-sid", "account_sid",
		"--call-sid", "call_sid",
		"--play-beep=false",
		"--recording-channels", "single",
		"--recording-status-callback", "http://webhook.com/callback",
		"--recording-status-callback-event", "in-progress completed absent",
		"--recording-status-callback-method", "GET",
		"--recording-track", "inbound",
		"--send-recording-url=false",
	)
}

func TestTexmlAccountsCallsRecordingsJsonRetrieveRecordingsJson(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:calls:recordings-json", "retrieve-recordings-json",
		"--account-sid", "account_sid",
		"--call-sid", "call_sid",
	)
}

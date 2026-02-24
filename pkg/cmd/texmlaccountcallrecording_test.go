// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsCallsRecordingsRecordingSidJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:calls:recordings", "recording-sid-json",
		"--account-sid", "account_sid",
		"--call-sid", "call_sid",
		"--recording-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--status", "paused",
	)
}

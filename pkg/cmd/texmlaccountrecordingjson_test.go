// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsRecordingsJsonDeleteRecordingSidJson(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:recordings:json", "delete-recording-sid-json",
		"--account-sid", "account_sid",
		"--recording-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestTexmlAccountsRecordingsJsonRetrieveRecordingSidJson(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:recordings:json", "retrieve-recording-sid-json",
		"--account-sid", "account_sid",
		"--recording-sid", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

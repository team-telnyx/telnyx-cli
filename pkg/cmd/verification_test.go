// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVerificationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verifications", "retrieve",
		"--api-key", "string",
		"--verification-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
	)
}

func TestVerificationsTriggerCall(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verifications", "trigger-call",
		"--api-key", "string",
		"--phone-number", "+13035551234",
		"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--custom-code", "43612",
		"--extension", "1www2WABCDw9",
		"--timeout-secs", "300",
	)
}

func TestVerificationsTriggerFlashcall(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verifications", "trigger-flashcall",
		"--api-key", "string",
		"--phone-number", "+13035551234",
		"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--timeout-secs", "300",
	)
}

func TestVerificationsTriggerSMS(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verifications", "trigger-sms",
		"--api-key", "string",
		"--phone-number", "+13035551234",
		"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--custom-code", "43612",
		"--timeout-secs", "300",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVerificationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verifications", "retrieve",
			"--api-key", "string",
			"--verification-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}

func TestVerificationsTriggerCall(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verifications", "trigger-call",
			"--api-key", "string",
			"--phone-number", "+13035551234",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--custom-code", "43612",
			"--extension", "1www2WABCDw9",
			"--timeout-secs", "300",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_number: '+13035551234'\n" +
			"verify_profile_id: 12ade33a-21c0-473b-b055-b3c836e1c292\n" +
			"custom_code: '43612'\n" +
			"extension: 1www2WABCDw9\n" +
			"timeout_secs: 300\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "verifications", "trigger-call",
			"--api-key", "string",
		)
	})
}

func TestVerificationsTriggerFlashcall(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verifications", "trigger-flashcall",
			"--api-key", "string",
			"--phone-number", "+13035551234",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--timeout-secs", "300",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_number: '+13035551234'\n" +
			"verify_profile_id: 12ade33a-21c0-473b-b055-b3c836e1c292\n" +
			"timeout_secs: 300\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "verifications", "trigger-flashcall",
			"--api-key", "string",
		)
	})
}

func TestVerificationsTriggerSMS(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verifications", "trigger-sms",
			"--api-key", "string",
			"--phone-number", "+13035551234",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--custom-code", "43612",
			"--timeout-secs", "300",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_number: '+13035551234'\n" +
			"verify_profile_id: 12ade33a-21c0-473b-b055-b3c836e1c292\n" +
			"custom_code: '43612'\n" +
			"timeout_secs: 300\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "verifications", "trigger-sms",
			"--api-key", "string",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestBotSignupCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bot-signup", "create",
			"--bot-challenge-answer", "35",
			"--bot-challenge-nonce", "c6feda4e-6501-4db9-a21f-665e5b4ce2ba",
			"--privacy-policy-url", "https://telnyx.com/privacy-policy",
			"--terms-and-conditions-url", "https://telnyx.com/terms-and-conditions-of-service",
			"--terms-of-service=true",
			"--email", "agent-owner@example.com",
			"--terms-and-conditions-eu-url", "https://telnyx.com/terms-and-conditions-of-service-eu",
			"--terms-of-service-eu=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bot_challenge_answer: '35'\n" +
			"bot_challenge_nonce: c6feda4e-6501-4db9-a21f-665e5b4ce2ba\n" +
			"privacy_policy_url: https://telnyx.com/privacy-policy\n" +
			"terms_and_conditions_url: https://telnyx.com/terms-and-conditions-of-service\n" +
			"terms_of_service: true\n" +
			"email: agent-owner@example.com\n" +
			"terms_and_conditions_eu_url: https://telnyx.com/terms-and-conditions-of-service-eu\n" +
			"terms_of_service_eu: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bot-signup", "create",
		)
	})
}

func TestBotSignupResendMagicLink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bot-signup", "resend-magic-link",
			"--email", "agent-owner@example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("email: agent-owner@example.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bot-signup", "resend-magic-link",
		)
	})
}

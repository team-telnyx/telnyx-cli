// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappBusinessAccountsPhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp:business-accounts:phone-numbers", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "id",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestWhatsappBusinessAccountsPhoneNumbersInitializeVerification(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp:business-accounts:phone-numbers", "initialize-verification",
			"--api-key", "string",
			"--id", "id",
			"--display-name", "display_name",
			"--phone-number", "phone_number",
			"--language", "language",
			"--verification-method", "sms",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"display_name: display_name\n" +
			"phone_number: phone_number\n" +
			"language: language\n" +
			"verification_method: sms\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "whatsapp:business-accounts:phone-numbers", "initialize-verification",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

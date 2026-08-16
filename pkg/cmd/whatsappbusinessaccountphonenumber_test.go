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
			t,
			"--api-key", "string",
			"whatsapp:business-accounts:phone-numbers", "list",
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
			t,
			"--api-key", "string",
			"whatsapp:business-accounts:phone-numbers", "initialize-verification",
			"--id", "id",
			"--display-name", "string",
			"--phone-number", "string",
			"--language", "en_US",
			"--verification-method", "sms",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"display_name: string\n" +
			"phone_number: string\n" +
			"language: en_US\n" +
			"verification_method: sms\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"whatsapp:business-accounts:phone-numbers", "initialize-verification",
			"--id", "id",
		)
	})
}

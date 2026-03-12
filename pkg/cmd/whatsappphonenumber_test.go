// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappPhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp:phone-numbers", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestWhatsappPhoneNumbersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp:phone-numbers", "delete",
			"--api-key", "string",
			"--phone-number", "phone_number",
		)
	})
}

func TestWhatsappPhoneNumbersResendVerification(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp:phone-numbers", "resend-verification",
			"--api-key", "string",
			"--phone-number", "phone_number",
			"--verification-method", "sms",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("verification_method: sms")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "whatsapp:phone-numbers", "resend-verification",
			"--api-key", "string",
			"--phone-number", "phone_number",
		)
	})
}

func TestWhatsappPhoneNumbersVerify(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp:phone-numbers", "verify",
			"--api-key", "string",
			"--phone-number", "phone_number",
			"--code", "code",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("code: code")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "whatsapp:phone-numbers", "verify",
			"--api-key", "string",
			"--phone-number", "phone_number",
		)
	})
}

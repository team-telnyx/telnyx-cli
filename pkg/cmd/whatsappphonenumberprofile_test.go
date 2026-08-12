// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappPhoneNumbersProfileRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:profile", "retrieve",
			"--phone-number", "phone_number",
		)
	})
}

func TestWhatsappPhoneNumbersProfileUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:profile", "update",
			"--phone-number", "phone_number",
			"--about", "string",
			"--address", "string",
			"--category", "string",
			"--description", "string",
			"--display-name", "string",
			"--email", "string",
			"--profile-id", "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			"--website", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"about: string\n" +
			"address: string\n" +
			"category: string\n" +
			"description: string\n" +
			"display_name: string\n" +
			"email: string\n" +
			"profile_id: 3fa85f64-5717-4562-b3fc-2c963f66afa6\n" +
			"website: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"whatsapp:phone-numbers:profile", "update",
			"--phone-number", "phone_number",
		)
	})
}

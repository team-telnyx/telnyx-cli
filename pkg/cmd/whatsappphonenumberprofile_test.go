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
			t, "whatsapp:phone-numbers:profile", "retrieve",
			"--api-key", "string",
			"--phone-number", "phone_number",
		)
	})
}

func TestWhatsappPhoneNumbersProfileUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp:phone-numbers:profile", "update",
			"--api-key", "string",
			"--phone-number", "phone_number",
			"--about", "about",
			"--address", "address",
			"--category", "category",
			"--description", "description",
			"--display-name", "display_name",
			"--email", "email",
			"--website", "website",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"about: about\n" +
			"address: address\n" +
			"category: category\n" +
			"description: description\n" +
			"display_name: display_name\n" +
			"email: email\n" +
			"website: website\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "whatsapp:phone-numbers:profile", "update",
			"--api-key", "string",
			"--phone-number", "phone_number",
		)
	})
}

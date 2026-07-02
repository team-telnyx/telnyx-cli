// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappPhoneNumbersCallingSettingsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:calling-settings", "retrieve",
			"--phone-number", "phone_number",
		)
	})
}

func TestWhatsappPhoneNumbersCallingSettingsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:calling-settings", "update",
			"--phone-number", "phone_number",
			"--enabled=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("enabled: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"whatsapp:phone-numbers:calling-settings", "update",
			"--phone-number", "phone_number",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestWhatsappPhoneNumbersConversationalComponentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:conversational-components", "list",
			"--phone-number", "phone_number",
		)
	})
}

func TestWhatsappPhoneNumbersConversationalComponentsPatchAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:conversational-components", "patch-all",
			"--phone-number", "phone_number",
			"--command", "{command: command, description: description}",
			"--ice-breaker", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(whatsappPhoneNumbersConversationalComponentsPatchAll)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:conversational-components", "patch-all",
			"--phone-number", "phone_number",
			"--command.command", "command",
			"--command.description", "description",
			"--ice-breaker", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"commands:\n" +
			"  - command: command\n" +
			"    description: description\n" +
			"ice_breakers:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"whatsapp:phone-numbers:conversational-components", "patch-all",
			"--phone-number", "phone_number",
		)
	})
}

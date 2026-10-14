// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappPhoneNumbersCallingRoutingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:calling-routing", "list",
			"--id", "+13125550100",
		)
	})
}

func TestWhatsappPhoneNumbersCallingRoutingPatchAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:calling-routing", "patch-all",
			"--id", "+13125550100",
			"--connection-id", "'1234567890'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("connection_id: '1234567890'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"whatsapp:phone-numbers:calling-routing", "patch-all",
			"--id", "+13125550100",
		)
	})
}

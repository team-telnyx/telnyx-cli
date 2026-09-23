// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlCallsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:calls", "create",
			"--connection-id", "1234567890",
			"--from", "+13120001234",
			"--to", "+13121230000",
			"--method", "POST",
			"--texml", "<Response><Say>Hello</Say></Response>",
			"--url", "https://example.com/instructions.xml",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"From: '+13120001234'\n" +
			"To: '+13121230000'\n" +
			"Method: POST\n" +
			"Texml: <Response><Say>Hello</Say></Response>\n" +
			"Url: https://example.com/instructions.xml\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml:calls", "create",
			"--connection-id", "1234567890",
		)
	})
}

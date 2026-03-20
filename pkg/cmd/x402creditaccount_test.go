// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestX402CreditAccountCreateQuote(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"x402:credit-account", "create-quote",
			"--amount-usd", "50.00",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("amount_usd: '50.00'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"x402:credit-account", "create-quote",
		)
	})
}

func TestX402CreditAccountSettle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"x402:credit-account", "settle",
			"--id", "quote_abc123",
			"--payment-signature", "0xabc123...",
			"--payment-signature", "PAYMENT-SIGNATURE",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: quote_abc123\n" +
			"payment_signature: 0xabc123...\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"x402:credit-account", "settle",
			"--payment-signature", "PAYMENT-SIGNATURE",
		)
	})
}

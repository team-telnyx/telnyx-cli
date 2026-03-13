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
			t, "x402:credit-account", "create-quote",
			"--api-key", "string",
			"--amount-usd", "50.00",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("amount_usd: '50.00'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "x402:credit-account", "create-quote",
			"--api-key", "string",
		)
	})
}

func TestX402CreditAccountSettle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "x402:credit-account", "settle",
			"--api-key", "string",
			"--id", "quote_abc123",
			"--payment-signature", "0xabc123...",
			"--payment-signature-header", "PAYMENT-SIGNATURE",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: quote_abc123\n" +
			"payment_signature: 0xabc123...\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "x402:credit-account", "settle",
			"--api-key", "string",
			"--payment-signature-header", "PAYMENT-SIGNATURE",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPaymentCreateStoredPaymentTransaction(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "payment", "create-stored-payment-transaction",
			"--api-key", "string",
			"--amount", "120.00",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("amount: '120.00'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "payment", "create-stored-payment-transaction",
			"--api-key", "string",
		)
	})
}

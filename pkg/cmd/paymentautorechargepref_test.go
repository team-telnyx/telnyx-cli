// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPaymentAutoRechargePrefsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"payment:auto-recharge-prefs", "update",
			"--enabled=true",
			"--invoice-enabled=true",
			"--preference", "credit_paypal",
			"--recharge-amount", "104.00",
			"--threshold-amount", "104.00",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enabled: true\n" +
			"invoice_enabled: true\n" +
			"preference: credit_paypal\n" +
			"recharge_amount: '104.00'\n" +
			"threshold_amount: '104.00'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"payment:auto-recharge-prefs", "update",
		)
	})
}

func TestPaymentAutoRechargePrefsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"payment:auto-recharge-prefs", "list",
		)
	})
}

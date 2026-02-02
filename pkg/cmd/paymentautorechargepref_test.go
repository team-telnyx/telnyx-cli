// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestPaymentAutoRechargePrefsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"payment:auto-recharge-prefs", "update",
		"--enabled=true",
		"--invoice-enabled=true",
		"--preference", "credit_paypal",
		"--recharge-amount", "104.00",
		"--threshold-amount", "104.00",
	)
}

func TestPaymentAutoRechargePrefsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"payment:auto-recharge-prefs", "list",
	)
}

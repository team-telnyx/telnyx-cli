// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsCallsSiprecSiprecSidJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls:siprec", "siprec-sid-json",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
			"--siprec-sid", "siprec_sid",
			"--status", "stopped",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("Status: stopped")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml:accounts:calls:siprec", "siprec-sid-json",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
			"--siprec-sid", "siprec_sid",
		)
	})
}

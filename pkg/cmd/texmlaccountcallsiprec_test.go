// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsCallsSiprecSiprecSidJson(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:calls:siprec", "siprec-sid-json",
		"--account-sid", "account_sid",
		"--call-sid", "call_sid",
		"--siprec-sid", "siprec_sid",
		"--status", "stopped",
	)
}

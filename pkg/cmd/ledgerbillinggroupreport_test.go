// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLedgerBillingGroupReportsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ledger-billing-group-reports", "create",
		"--month", "10",
		"--year", "2019",
	)
}

func TestLedgerBillingGroupReportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ledger-billing-group-reports", "retrieve",
		"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
	)
}

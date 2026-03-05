// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLedgerBillingGroupReportsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ledger-billing-group-reports", "create",
			"--api-key", "string",
			"--month", "10",
			"--year", "2019",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"month: 10\n" +
			"year: 2019\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ledger-billing-group-reports", "create",
			"--api-key", "string",
		)
	})
}

func TestLedgerBillingGroupReportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ledger-billing-group-reports", "retrieve",
			"--api-key", "string",
			"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
		)
	})
}

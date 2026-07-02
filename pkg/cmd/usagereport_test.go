// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestUsageReportsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage-reports", "list",
			"--max-items", "10",
			"--dimension", "string",
			"--metric", "string",
			"--product", "product",
			"--date-range", "date_range",
			"--end-date", "end_date",
			"--filter", "filter",
			"--format", "csv",
			"--managed-accounts=true",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "string",
			"--start-date", "start_date",
			"--authorization-bearer", "authorization_bearer",
		)
	})
}

func TestUsageReportsGetOptions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage-reports", "get-options",
			"--product", "product",
			"--authorization-bearer", "authorization_bearer",
		)
	})
}

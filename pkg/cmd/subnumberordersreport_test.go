// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSubNumberOrdersReportCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sub-number-orders-report", "create",
			"--api-key", "string",
			"--country-code", "US",
			"--created-at-gt", "'2023-04-05T10:22:08.230549Z'",
			"--created-at-lt", "'2025-06-05T10:22:08.230549Z'",
			"--customer-reference", "STRING",
			"--order-request-id", "12ade33a-21c0-473b-b055-b3c836e1c293",
			"--status", "success",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"country_code: US\n" +
			"created_at_gt: '2023-04-05T10:22:08.230549Z'\n" +
			"created_at_lt: '2025-06-05T10:22:08.230549Z'\n" +
			"customer_reference: STRING\n" +
			"order_request_id: 12ade33a-21c0-473b-b055-b3c836e1c293\n" +
			"status: success\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "sub-number-orders-report", "create",
			"--api-key", "string",
		)
	})
}

func TestSubNumberOrdersReportRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sub-number-orders-report", "retrieve",
			"--api-key", "string",
			"--report-id", "12ade33a-21c0-473b-b055-b3c836e1c293",
		)
	})
}

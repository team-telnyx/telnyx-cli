// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortoutsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestPortoutsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts", "list",
		"--filter", "{carrier_name: carrier_name, country_code: US, country_code_in: [CA, US], foc_date: '2024-09-04T00:00:00.000Z', inserted_at: {gte: '2024-09-04T00:00:00.000Z', lte: '2024-09-04T00:00:00.000Z'}, phone_number: '+13035551212', pon: pon, ported_out_at: {gte: '2024-09-04T00:00:00.000Z', lte: '2024-09-04T00:00:00.000Z'}, spid: spid, status: pending, status_in: [pending], support_key: PO_abc123}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portoutsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts", "list",
		"--filter.carrier-name", "carrier_name",
		"--filter.country-code", "US",
		"--filter.country-code-in", "[CA, US]",
		"--filter.foc-date", "2024-09-04T00:00:00.000Z",
		"--filter.inserted-at", "{gte: '2024-09-04T00:00:00.000Z', lte: '2024-09-04T00:00:00.000Z'}",
		"--filter.phone-number", "+13035551212",
		"--filter.pon", "pon",
		"--filter.ported-out-at", "{gte: '2024-09-04T00:00:00.000Z', lte: '2024-09-04T00:00:00.000Z'}",
		"--filter.spid", "spid",
		"--filter.status", "pending",
		"--filter.status-in", "[pending]",
		"--filter.support-key", "PO_abc123",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestPortoutsListRejectionCodes(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts", "list-rejection-codes",
		"--portout-id", "329d6658-8f93-405d-862f-648776e8afd7",
		"--filter", "{code: 1002}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portoutsListRejectionCodes)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts", "list-rejection-codes",
		"--portout-id", "329d6658-8f93-405d-862f-648776e8afd7",
		"--filter.code", "1002",
	)
}

func TestPortoutsUpdateStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts", "update-status",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--status", "authorized",
		"--reason", "I do not recognize this transaction",
		"--host-messaging=false",
	)
}

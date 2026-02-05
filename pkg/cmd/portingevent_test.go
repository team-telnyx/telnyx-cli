// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingEventsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:events", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestPortingEventsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:events", "list",
		"--filter", "{created_at: {gte: '2021-01-01T00:00:00Z', lte: '2021-01-01T00:00:00Z'}, porting_order_id: 34dc46a9-53ed-4e01-9454-26227ea13326, type: porting_order.deleted}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingEventsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:events", "list",
		"--filter.created-at", "{gte: '2021-01-01T00:00:00Z', lte: '2021-01-01T00:00:00Z'}",
		"--filter.porting-order-id", "34dc46a9-53ed-4e01-9454-26227ea13326",
		"--filter.type", "porting_order.deleted",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestPortingEventsRepublish(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:events", "republish",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

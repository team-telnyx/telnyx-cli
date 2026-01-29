// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestPortoutsEventsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:events", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestPortoutsEventsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:events", "list",
		"--filter", "{created_at: {gte: '2021-01-01T00:00:00Z', lte: '2021-01-01T00:00:00Z'}, event_type: portout.status_changed, portout_id: 34dc46a9-53ed-4e01-9454-26227ea13326}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portoutsEventsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:events", "list",
		"--filter.created-at", "{gte: '2021-01-01T00:00:00Z', lte: '2021-01-01T00:00:00Z'}",
		"--filter.event-type", "portout.status_changed",
		"--filter.portout-id", "34dc46a9-53ed-4e01-9454-26227ea13326",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestPortoutsEventsRepublish(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:events", "republish",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

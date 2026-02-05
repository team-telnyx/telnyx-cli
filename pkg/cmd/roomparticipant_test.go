// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRoomParticipantsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-participants", "retrieve",
		"--room-participant-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
	)
}

func TestRoomParticipantsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-participants", "list",
		"--filter", "{context: Alice, date_joined_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_left_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_updated_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, session_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(roomParticipantsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-participants", "list",
		"--filter.context", "Alice",
		"--filter.date-joined-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-left-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-updated-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--page-number", "0",
		"--page-size", "0",
	)
}

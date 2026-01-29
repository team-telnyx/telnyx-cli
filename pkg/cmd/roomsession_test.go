// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestRoomsSessionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms:sessions", "retrieve",
		"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--include-participants=true",
	)
}

func TestRoomsSessionsList0(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms:sessions", "list-0",
		"--filter", "{active: true, date_created_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_ended_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_updated_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, room_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0}",
		"--include-participants=true",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(roomsSessionsList0)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms:sessions", "list-0",
		"--filter.active=true",
		"--filter.date-created-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-ended-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-updated-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--include-participants=true",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestRoomsSessionsList1(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms:sessions", "list-1",
		"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter", "{active: true, date_created_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_ended_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_updated_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}}",
		"--include-participants=true",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(roomsSessionsList1)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms:sessions", "list-1",
		"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter.active=true",
		"--filter.date-created-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-ended-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-updated-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--include-participants=true",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestRoomsSessionsRetrieveParticipants(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms:sessions", "retrieve-participants",
		"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter", "{context: Alice, date_joined_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_left_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_updated_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(roomsSessionsRetrieveParticipants)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms:sessions", "retrieve-participants",
		"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter.context", "Alice",
		"--filter.date-joined-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-left-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-updated-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--page.number", "1",
		"--page.size", "1",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRoomRecordingsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-recordings", "retrieve",
		"--room-recording-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
	)
}

func TestRoomRecordingsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-recordings", "list",
		"--filter", "{date_ended_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_started_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, duration_secs: 20, participant_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, room_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, session_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, status: completed, type: audio}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(roomRecordingsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-recordings", "list",
		"--filter.date-ended-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-started-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.duration-secs", "20",
		"--filter.participant-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter.room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter.session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter.status", "completed",
		"--filter.type", "audio",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestRoomRecordingsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-recordings", "delete",
		"--room-recording-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
	)
}

func TestRoomRecordingsDeleteBulk(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-recordings", "delete-bulk",
		"--filter", "{date_ended_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_started_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, duration_secs: 20, participant_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, room_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, session_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, status: completed, type: audio}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(roomRecordingsDeleteBulk)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-recordings", "delete-bulk",
		"--filter.date-ended-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-started-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.duration-secs", "20",
		"--filter.participant-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter.room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter.session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--filter.status", "completed",
		"--filter.type", "audio",
		"--page-number", "0",
		"--page-size", "0",
	)
}

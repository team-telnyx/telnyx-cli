// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRecordingsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recordings", "retrieve",
		"--recording-id", "recording_id",
	)
}

func TestRecordingsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recordings", "list",
		"--filter", "{call_leg_id: 428c31b6-7af4-4bcb-b7f5-5013ef9657c1, call_session_id: 428c31b6-7af4-4bcb-b7f5-5013ef9657c1, conference_id: 428c31b6-7af4-4bcb-b7f5-5013ef9657c1, connection_id: '175237942907135762', created_at: {gte: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}, from: '1234567890', sip_call_id: 428c31b6-7af4-4bcb-b7f5-5013ef9657c1, to: '1234567890'}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(recordingsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"recordings", "list",
		"--filter.call-leg-id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c1",
		"--filter.call-session-id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c1",
		"--filter.conference-id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c1",
		"--filter.connection-id", "175237942907135762",
		"--filter.created-at", "{gte: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}",
		"--filter.from", "1234567890",
		"--filter.sip-call-id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c1",
		"--filter.to", "1234567890",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestRecordingsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recordings", "delete",
		"--recording-id", "recording_id",
	)
}

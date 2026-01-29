// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestRoomCompositionsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-compositions", "create",
		"--format", "mp4",
		"--resolution", "800x600",
		"--session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777b0",
		"--video-layout", "{foo: {height: 360, max_columns: 3, max_rows: 3, video_sources: [7b61621f-62e0-4aad-ab11-9fd19e272e73], width: 480, x_pos: 100, y_pos: 100, z_pos: 1}}",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)
}

func TestRoomCompositionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-compositions", "retrieve",
		"--room-composition-id", "5219b3af-87c6-4c08-9b58-5a533d893e21",
	)
}

func TestRoomCompositionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-compositions", "list",
		"--filter", "{date_created_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, session_id: 92e7d459-bcc5-4386-9f5f-6dd14a82588d, status: completed}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(roomCompositionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-compositions", "list",
		"--filter.date-created-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.session-id", "92e7d459-bcc5-4386-9f5f-6dd14a82588d",
		"--filter.status", "completed",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestRoomCompositionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"room-compositions", "delete",
		"--room-composition-id", "5219b3af-87c6-4c08-9b58-5a533d893e21",
	)
}

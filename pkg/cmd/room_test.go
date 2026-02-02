// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestRoomsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms", "create",
		"--enable-recording=true",
		"--max-participants", "10",
		"--unique-name", "My room",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)
}

func TestRoomsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms", "retrieve",
		"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--include-sessions=true",
	)
}

func TestRoomsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms", "update",
		"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--enable-recording=true",
		"--max-participants", "10",
		"--unique-name", "My room",
		"--webhook-event-failover-url", "https://failover.example.com",
		"--webhook-event-url", "https://example.com",
		"--webhook-timeout-secs", "25",
	)
}

func TestRoomsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms", "list",
		"--filter", "{date_created_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_updated_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, unique_name: my_video_room}",
		"--include-sessions=true",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(roomsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms", "list",
		"--filter.date-created-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.date-updated-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
		"--filter.unique-name", "my_video_room",
		"--include-sessions=true",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestRoomsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rooms", "delete",
		"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
	)
}

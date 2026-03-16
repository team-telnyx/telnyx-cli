// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRoomsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms", "create",
			"--enable-recording=true",
			"--max-participants", "10",
			"--unique-name", "My room",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-event-url", "https://example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enable_recording: true\n" +
			"max_participants: 10\n" +
			"unique_name: My room\n" +
			"webhook_event_failover_url: https://failover.example.com\n" +
			"webhook_event_url: https://example.com\n" +
			"webhook_timeout_secs: 25\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rooms", "create",
		)
	})
}

func TestRoomsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms", "retrieve",
			"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
			"--include-sessions=true",
		)
	})
}

func TestRoomsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms", "update",
			"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
			"--enable-recording=true",
			"--max-participants", "10",
			"--unique-name", "My room",
			"--webhook-event-failover-url", "https://failover.example.com",
			"--webhook-event-url", "https://example.com",
			"--webhook-timeout-secs", "25",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enable_recording: true\n" +
			"max_participants: 10\n" +
			"unique_name: My room\n" +
			"webhook_event_failover_url: https://failover.example.com\n" +
			"webhook_event_url: https://example.com\n" +
			"webhook_timeout_secs: 25\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rooms", "update",
			"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

func TestRoomsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms", "list",
			"--max-items", "10",
			"--filter", "{date_created_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, date_updated_at: {eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}, unique_name: my_video_room}",
			"--include-sessions=true",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(roomsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms", "list",
			"--max-items", "10",
			"--filter.date-created-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
			"--filter.date-updated-at", "{eq: '2021-04-25', gte: '2021-04-25', lte: '2021-04-25'}",
			"--filter.unique-name", "my_video_room",
			"--include-sessions=true",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestRoomsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms", "delete",
			"--room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

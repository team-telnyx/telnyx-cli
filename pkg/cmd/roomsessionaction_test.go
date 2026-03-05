// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestRoomsSessionsActionsEnd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "rooms:sessions:actions", "end",
			"--api-key", "string",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

func TestRoomsSessionsActionsKick(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "rooms:sessions:actions", "kick",
			"--api-key", "string",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
			"--exclude", "7b61621f-62e0-4aad-ab11-9fd19e272e73",
			"--participants", "all",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"exclude:\n" +
			"  - 7b61621f-62e0-4aad-ab11-9fd19e272e73\n" +
			"participants: all\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "rooms:sessions:actions", "kick",
			"--api-key", "string",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

func TestRoomsSessionsActionsMute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "rooms:sessions:actions", "mute",
			"--api-key", "string",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
			"--exclude", "7b61621f-62e0-4aad-ab11-9fd19e272e73",
			"--participants", "all",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"exclude:\n" +
			"  - 7b61621f-62e0-4aad-ab11-9fd19e272e73\n" +
			"participants: all\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "rooms:sessions:actions", "mute",
			"--api-key", "string",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

func TestRoomsSessionsActionsUnmute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "rooms:sessions:actions", "unmute",
			"--api-key", "string",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
			"--exclude", "7b61621f-62e0-4aad-ab11-9fd19e272e73",
			"--participants", "all",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"exclude:\n" +
			"  - 7b61621f-62e0-4aad-ab11-9fd19e272e73\n" +
			"participants: all\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "rooms:sessions:actions", "unmute",
			"--api-key", "string",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

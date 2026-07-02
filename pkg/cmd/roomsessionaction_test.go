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
			t,
			"--api-key", "string",
			"rooms:sessions:actions", "end",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

func TestRoomsSessionsActionsKick(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms:sessions:actions", "kick",
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
			t, pipeData,
			"--api-key", "string",
			"rooms:sessions:actions", "kick",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

func TestRoomsSessionsActionsMute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms:sessions:actions", "mute",
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
			t, pipeData,
			"--api-key", "string",
			"rooms:sessions:actions", "mute",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

func TestRoomsSessionsActionsUnmute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rooms:sessions:actions", "unmute",
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
			t, pipeData,
			"--api-key", "string",
			"rooms:sessions:actions", "unmute",
			"--room-session-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})
}

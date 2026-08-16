// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMeetingSessionsActionsSendChat(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions:actions", "send-chat",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			"--text", "I will send the summary after this call.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("text: I will send the summary after this call.")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"meeting-sessions:actions", "send-chat",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

func TestMeetingSessionsActionsSpeak(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions:actions", "speak",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			"--text", "Here are the three decisions from this call.",
			"--interrupt=false",
			"--voice", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"text: Here are the three decisions from this call.\n" +
			"interrupt: false\n" +
			"voice: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"meeting-sessions:actions", "speak",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

func TestMeetingSessionsActionsStopSpeaking(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions:actions", "stop-speaking",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

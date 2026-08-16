// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMeetingSessionsArtifactsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions:artifacts", "create",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			"--type", "summary",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("type: summary")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"meeting-sessions:artifacts", "create",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

func TestMeetingSessionsArtifactsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions:artifacts", "retrieve",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			"--artifact-id", "mtgart_b2c3d4e5-f6a7-8901-bcde-f23456789012",
		)
	})
}

func TestMeetingSessionsArtifactsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions:artifacts", "list",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestCalendarConnectionsMeetingsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calendar-connections:meetings", "update",
			"--id", "id",
			"--meeting-id", "meeting_id",
			"--bot-will-join=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("bot_will_join: false")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calendar-connections:meetings", "update",
			"--id", "id",
			"--meeting-id", "meeting_id",
		)
	})
}

func TestCalendarConnectionsMeetingsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calendar-connections:meetings", "list",
			"--id", "id",
			"--from", "'2019-12-27T18:11:19.117Z'",
			"--to", "'2019-12-27T18:11:19.117Z'",
		)
	})
}

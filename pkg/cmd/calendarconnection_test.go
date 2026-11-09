// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestCalendarConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calendar-connections", "create",
			"--platform", "google_calendar",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("platform: google_calendar")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calendar-connections", "create",
		)
	})
}

func TestCalendarConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calendar-connections", "retrieve",
			"--id", "id",
		)
	})
}

func TestCalendarConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calendar-connections", "update",
			"--id", "id",
			"--auto-join-policy", "all_meetings",
			"--bot-name", "x",
			"--default-assistant", "{id: x, audio_gate: half_duplex, dynamic_variables: {foo: string}, leave_on_end: true}",
			"--default-barge-in=true",
			"--default-recording=true",
			"--default-recording-export", "{bucket: meeting-recordings, prefix: prefix}",
			"--default-summarize-on-end=true",
			"--default-voice", "x",
			"--default-webhook-url", "https://example.com/meeting-events",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(calendarConnectionsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calendar-connections", "update",
			"--id", "id",
			"--auto-join-policy", "all_meetings",
			"--bot-name", "x",
			"--default-assistant.id", "x",
			"--default-assistant.audio-gate", "half_duplex",
			"--default-assistant.dynamic-variables", "{foo: string}",
			"--default-assistant.leave-on-end=true",
			"--default-barge-in=true",
			"--default-recording=true",
			"--default-recording-export.bucket", "meeting-recordings",
			"--default-recording-export.prefix", "prefix",
			"--default-summarize-on-end=true",
			"--default-voice", "x",
			"--default-webhook-url", "https://example.com/meeting-events",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auto_join_policy: all_meetings\n" +
			"bot_name: x\n" +
			"default_assistant:\n" +
			"  id: x\n" +
			"  audio_gate: half_duplex\n" +
			"  dynamic_variables:\n" +
			"    foo: string\n" +
			"  leave_on_end: true\n" +
			"default_barge_in: true\n" +
			"default_recording: true\n" +
			"default_recording_export:\n" +
			"  bucket: meeting-recordings\n" +
			"  prefix: prefix\n" +
			"default_summarize_on_end: true\n" +
			"default_voice: x\n" +
			"default_webhook_url: https://example.com/meeting-events\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calendar-connections", "update",
			"--id", "id",
		)
	})
}

func TestCalendarConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calendar-connections", "list",
		)
	})
}

func TestCalendarConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calendar-connections", "delete",
			"--id", "id",
		)
	})
}

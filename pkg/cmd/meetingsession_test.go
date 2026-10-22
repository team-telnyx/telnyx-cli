// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMeetingSessionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "create",
			"--meeting-url", "https://zoom.us/j/1234567890",
			"--assistant", "{id: asst_fake-uuid-1234, audio_gate: full_duplex, avatar: true, dynamic_variables: {foo: string}, leave_on_end: true}",
			"--barge-in=true",
			"--bot-name", "Notetaker",
			"--camera-image", "{base64_data: /9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/wAALCAACAAIBAREA/8QAFAABAAAAAAAAAAAAAAAAAAAACP/EAB4QAAAEBwAAAAAAAAAAAAAAAAAEBgcCFic1RVNi/9oACAEBAAA/AH8hGJbWR09TxKW4vhC2qHgf/9k=, format: jpeg}",
			"--chat-on-enter", "x",
			"--idempotency-key", "x",
			"--join-at", "'2019-12-27T18:11:19.117Z'",
			"--metadata", "{foo: bar}",
			"--speak-on-enter", "x",
			"--summarize-on-end=true",
			"--voice", "x",
			"--webhook-url", "https://example.com",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(meetingSessionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "create",
			"--meeting-url", "https://zoom.us/j/1234567890",
			"--assistant.id", "asst_fake-uuid-1234",
			"--assistant.audio-gate", "full_duplex",
			"--assistant.avatar=true",
			"--assistant.dynamic-variables", "{foo: string}",
			"--assistant.leave-on-end=true",
			"--barge-in=true",
			"--bot-name", "Notetaker",
			"--camera-image", "{base64_data: /9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/wAALCAACAAIBAREA/8QAFAABAAAAAAAAAAAAAAAAAAAACP/EAB4QAAAEBwAAAAAAAAAAAAAAAAAEBgcCFic1RVNi/9oACAEBAAA/AH8hGJbWR09TxKW4vhC2qHgf/9k=, format: jpeg}",
			"--chat-on-enter", "x",
			"--idempotency-key", "x",
			"--join-at", "'2019-12-27T18:11:19.117Z'",
			"--metadata", "{foo: bar}",
			"--speak-on-enter", "x",
			"--summarize-on-end=true",
			"--voice", "x",
			"--webhook-url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"meeting_url: https://zoom.us/j/1234567890\n" +
			"assistant:\n" +
			"  id: asst_fake-uuid-1234\n" +
			"  audio_gate: full_duplex\n" +
			"  avatar: true\n" +
			"  dynamic_variables:\n" +
			"    foo: string\n" +
			"  leave_on_end: true\n" +
			"barge_in: true\n" +
			"bot_name: Notetaker\n" +
			"camera_image:\n" +
			"  base64_data: >-\n" +
			"    /9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/wAALCAACAAIBAREA/8QAFAABAAAAAAAAAAAAAAAAAAAACP/EAB4QAAAEBwAAAAAAAAAAAAAAAAAEBgcCFic1RVNi/9oACAEBAAA/AH8hGJbWR09TxKW4vhC2qHgf/9k=\n" +
			"  format: jpeg\n" +
			"chat_on_enter: x\n" +
			"idempotency_key: x\n" +
			"join_at: '2019-12-27T18:11:19.117Z'\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"speak_on_enter: x\n" +
			"summarize_on_end: true\n" +
			"voice: x\n" +
			"webhook_url: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"meeting-sessions", "create",
		)
	})
}

func TestMeetingSessionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "retrieve",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

func TestMeetingSessionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "update",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			"--bot-name", "x",
			"--join-at", "'2026-08-05T17:00:00Z'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bot_name: x\n" +
			"join_at: '2026-08-05T17:00:00Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"meeting-sessions", "update",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

func TestMeetingSessionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "list",
			"--status", "scheduled",
		)
	})
}

func TestMeetingSessionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "delete",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

func TestMeetingSessionsDeleteRecordingMedia(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "delete-recording-media",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		)
	})
}

func TestMeetingSessionsRetrieveEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "retrieve-events",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			"--after", "0",
			"--limit", "1",
		)
	})
}

func TestMeetingSessionsRetrieveTranscript(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting-sessions", "retrieve-transcript",
			"--id", "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			"--after", "0",
			"--limit", "1",
			"--wait-seconds", "0",
		)
	})
}

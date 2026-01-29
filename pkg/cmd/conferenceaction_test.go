// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestConferencesActionsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "update",
		"--id", "id",
		"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		"--supervisor-role", "whisper",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--region", "US",
		"--whisper-call-control-id", "v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ",
		"--whisper-call-control-id", "v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw",
	)
}

func TestConferencesActionsHold(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "hold",
		"--id", "id",
		"--audio-url", "http://example.com/message.wav",
		"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		"--media-name", "my_media_uploaded_to_media_storage_api",
		"--region", "US",
	)
}

func TestConferencesActionsJoin(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "join",
		"--id", "id",
		"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		"--beep-enabled", "always",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--end-conference-on-exit=true",
		"--hold=true",
		"--hold-audio-url", "http://www.example.com/audio.wav",
		"--hold-media-name", "my_media_uploaded_to_media_storage_api",
		"--mute=true",
		"--region", "US",
		"--soft-end-conference-on-exit=true",
		"--start-conference-on-enter=true",
		"--supervisor-role", "whisper",
		"--whisper-call-control-id", "v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ",
		"--whisper-call-control-id", "v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw",
	)
}

func TestConferencesActionsLeave(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "leave",
		"--id", "id",
		"--call-control-id", "c46e06d7-b78f-4b13-96b6-c576af9640ff",
		"--beep-enabled", "never",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--region", "US",
	)
}

func TestConferencesActionsMute(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "mute",
		"--id", "id",
		"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		"--region", "US",
	)
}

func TestConferencesActionsPlay(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "play",
		"--id", "id",
		"--audio-url", "http://www.example.com/sounds/greeting.wav",
		"--call-control-id", "string",
		"--loop", "infinity",
		"--media-name", "my_media_uploaded_to_media_storage_api",
		"--region", "US",
	)
}

func TestConferencesActionsRecordPause(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "record-pause",
		"--id", "id",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--recording-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--region", "US",
	)
}

func TestConferencesActionsRecordResume(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "record-resume",
		"--id", "id",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--recording-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--region", "US",
	)
}

func TestConferencesActionsRecordStart(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "record-start",
		"--id", "id",
		"--format", "wav",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--custom-file-name", "my_recording_file_name",
		"--play-beep=true",
		"--region", "US",
		"--trim", "trim-silence",
	)
}

func TestConferencesActionsRecordStop(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "record-stop",
		"--id", "id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--recording-id", "6e00ab49-9487-4364-8ad6-23965965afb2",
		"--region", "US",
	)
}

func TestConferencesActionsSpeak(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "speak",
		"--id", "id",
		"--payload", "Say this to participants",
		"--voice", "female",
		"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--language", "en-US",
		"--payload-type", "text",
		"--region", "US",
		"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
	)
}

func TestConferencesActionsStop(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "stop",
		"--id", "id",
		"--call-control-id", "string",
		"--region", "US",
	)
}

func TestConferencesActionsUnhold(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "unhold",
		"--id", "id",
		"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		"--region", "US",
	)
}

func TestConferencesActionsUnmute(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conferences:actions", "unmute",
		"--id", "id",
		"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		"--region", "US",
	)
}

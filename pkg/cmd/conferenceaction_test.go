// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestConferencesActionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "update",
			"--id", "id",
			"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--supervisor-role", "whisper",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--region", "US",
			"--whisper-call-control-id", "v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ",
			"--whisper-call-control-id", "v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_id: v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"supervisor_role: whisper\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"region: US\n" +
			"whisper_call_control_ids:\n" +
			"  - v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ\n" +
			"  - v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "update",
			"--id", "id",
		)
	})
}

func TestConferencesActionsEndConference(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "end-conference",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("command_id: 891510ac-f3e4-11e8-af5b-de00688a4901")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "end-conference",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestConferencesActionsGatherDtmfAudio(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "gather-dtmf-audio",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--audio-url", "http://example.com/gather_prompt.wav",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--gather-id", "gather_id",
			"--initial-timeout-millis", "10000",
			"--inter-digit-timeout-millis", "3000",
			"--invalid-audio-url", "invalid_audio_url",
			"--invalid-media-name", "invalid_media_name",
			"--maximum-digits", "4",
			"--maximum-tries", "3",
			"--media-name", "media_name",
			"--minimum-digits", "1",
			"--stop-playback-on-dtmf=true",
			"--terminating-digit", "#",
			"--timeout-millis", "30000",
			"--valid-digits", "0123456789",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_id: v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"audio_url: http://example.com/gather_prompt.wav\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"gather_id: gather_id\n" +
			"initial_timeout_millis: 10000\n" +
			"inter_digit_timeout_millis: 3000\n" +
			"invalid_audio_url: invalid_audio_url\n" +
			"invalid_media_name: invalid_media_name\n" +
			"maximum_digits: 4\n" +
			"maximum_tries: 3\n" +
			"media_name: media_name\n" +
			"minimum_digits: 1\n" +
			"stop_playback_on_dtmf: true\n" +
			"terminating_digit: '#'\n" +
			"timeout_millis: 30000\n" +
			"valid_digits: '0123456789'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "gather-dtmf-audio",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestConferencesActionsHold(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "hold",
			"--id", "id",
			"--audio-url", "http://example.com/message.wav",
			"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--media-name", "my_media_uploaded_to_media_storage_api",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"audio_url: http://example.com/message.wav\n" +
			"call_control_ids:\n" +
			"  - v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"media_name: my_media_uploaded_to_media_storage_api\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "hold",
			"--id", "id",
		)
	})
}

func TestConferencesActionsJoin(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_id: v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"beep_enabled: always\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"end_conference_on_exit: true\n" +
			"hold: true\n" +
			"hold_audio_url: http://www.example.com/audio.wav\n" +
			"hold_media_name: my_media_uploaded_to_media_storage_api\n" +
			"mute: true\n" +
			"region: US\n" +
			"soft_end_conference_on_exit: true\n" +
			"start_conference_on_enter: true\n" +
			"supervisor_role: whisper\n" +
			"whisper_call_control_ids:\n" +
			"  - v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ\n" +
			"  - v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "join",
			"--id", "id",
		)
	})
}

func TestConferencesActionsLeave(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "leave",
			"--id", "id",
			"--call-control-id", "c46e06d7-b78f-4b13-96b6-c576af9640ff",
			"--beep-enabled", "never",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_id: c46e06d7-b78f-4b13-96b6-c576af9640ff\n" +
			"beep_enabled: never\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "leave",
			"--id", "id",
		)
	})
}

func TestConferencesActionsMute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "mute",
			"--id", "id",
			"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_ids:\n" +
			"  - v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "mute",
			"--id", "id",
		)
	})
}

func TestConferencesActionsPlay(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "play",
			"--id", "id",
			"--audio-url", "http://www.example.com/sounds/greeting.wav",
			"--call-control-id", "string",
			"--loop", "infinity",
			"--media-name", "my_media_uploaded_to_media_storage_api",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"audio_url: http://www.example.com/sounds/greeting.wav\n" +
			"call_control_ids:\n" +
			"  - string\n" +
			"loop: infinity\n" +
			"media_name: my_media_uploaded_to_media_storage_api\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "play",
			"--id", "id",
		)
	})
}

func TestConferencesActionsRecordPause(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "record-pause",
			"--id", "id",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--recording-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"recording_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "record-pause",
			"--id", "id",
		)
	})
}

func TestConferencesActionsRecordResume(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "record-resume",
			"--id", "id",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--recording-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"recording_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "record-resume",
			"--id", "id",
		)
	})
}

func TestConferencesActionsRecordStart(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "record-start",
			"--id", "id",
			"--format", "wav",
			"--channels", "dual",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--custom-file-name", "my_recording_file_name",
			"--play-beep=true",
			"--region", "US",
			"--trim", "trim-silence",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"format: wav\n" +
			"channels: dual\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"custom_file_name: my_recording_file_name\n" +
			"play_beep: true\n" +
			"region: US\n" +
			"trim: trim-silence\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "record-start",
			"--id", "id",
		)
	})
}

func TestConferencesActionsRecordStop(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "record-stop",
			"--id", "id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--recording-id", "6e00ab49-9487-4364-8ad6-23965965afb2",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"recording_id: 6e00ab49-9487-4364-8ad6-23965965afb2\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "record-stop",
			"--id", "id",
		)
	})
}

func TestConferencesActionsSendDtmf(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "send-dtmf",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--digits", "1234#",
			"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--duration-millis", "250",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"digits: 1234#\n" +
			"call_control_ids:\n" +
			"  - v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"duration_millis: 250\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "send-dtmf",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestConferencesActionsSpeak(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"payload: Say this to participants\n" +
			"voice: female\n" +
			"call_control_ids:\n" +
			"  - v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"language: en-US\n" +
			"payload_type: text\n" +
			"region: US\n" +
			"voice_settings:\n" +
			"  type: elevenlabs\n" +
			"  api_key_ref: my_elevenlabs_api_key\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "speak",
			"--id", "id",
		)
	})
}

func TestConferencesActionsStop(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "stop",
			"--id", "id",
			"--call-control-id", "string",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_ids:\n" +
			"  - string\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "stop",
			"--id", "id",
		)
	})
}

func TestConferencesActionsUnhold(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "unhold",
			"--id", "id",
			"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_ids:\n" +
			"  - v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "unhold",
			"--id", "id",
		)
	})
}

func TestConferencesActionsUnmute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conferences:actions", "unmute",
			"--id", "id",
			"--call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--region", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_ids:\n" +
			"  - v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"region: US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conferences:actions", "unmute",
			"--id", "id",
		)
	})
}

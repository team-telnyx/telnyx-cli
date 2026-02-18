// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestCallsActionsAddAIAssistantMessages(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "add-ai-assistant-messages",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--message", "{content: Get the user's favorite color, role: system, metadata: {foo: bar}}",
	)
}

func TestCallsActionsAnswer(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "answer",
		"--call-control-id", "call_control_id",
		"--billing-group-id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--custom-header", "{name: head_1, value: val_1}",
		"--custom-header", "{name: head_2, value: val_2}",
		"--preferred-codecs", "G722,PCMU,PCMA,G729,OPUS,VP8,H264",
		"--record", "record-from-answer",
		"--record-channels", "single",
		"--record-custom-file-name", "my_recording_file_name",
		"--record-format", "wav",
		"--record-max-length", "1000",
		"--record-timeout-secs", "100",
		"--record-track", "outbound",
		"--record-trim", "trim-silence",
		"--send-silence-when-idle=true",
		"--sip-header", "{name: User-to-User, value: value}",
		"--sound-modifications", "{octaves: 0.1, pitch: 0.8, semitone: -2, track: both}",
		"--stream-bidirectional-codec", "G722",
		"--stream-bidirectional-mode", "rtp",
		"--stream-bidirectional-target-legs", "both",
		"--stream-codec", "PCMA",
		"--stream-track", "both_tracks",
		"--stream-url", "wss://www.example.com/websocket",
		"--transcription=true",
		"--transcription-config", "{client_state: aGF2ZSBhIG5pY2UgZGF5ID1d, command_id: 891510ac-f3e4-11e8-af5b-de00688a4901, transcription_engine: Google, transcription_engine_config: {enable_speaker_diarization: true, hints: [string], interim_results: true, language: en, max_speaker_count: 4, min_speaker_count: 4, model: latest_long, profanity_filter: true, speech_context: [{boost: 1, phrases: [string]}], transcription_engine: Google, use_enhanced: true}, transcription_tracks: both}",
		"--webhook-url", "https://www.example.com/server-b/",
		"--webhook-url-method", "POST",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callsActionsAnswer)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "answer",
		"--call-control-id", "call_control_id",
		"--billing-group-id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--custom-header.name", "head_1",
		"--custom-header.value", "val_1",
		"--custom-header.name", "head_2",
		"--custom-header.value", "val_2",
		"--preferred-codecs", "G722,PCMU,PCMA,G729,OPUS,VP8,H264",
		"--record", "record-from-answer",
		"--record-channels", "single",
		"--record-custom-file-name", "my_recording_file_name",
		"--record-format", "wav",
		"--record-max-length", "1000",
		"--record-timeout-secs", "100",
		"--record-track", "outbound",
		"--record-trim", "trim-silence",
		"--send-silence-when-idle=true",
		"--sip-header.name", "User-to-User",
		"--sip-header.value", "value",
		"--sound-modifications.octaves", "0.1",
		"--sound-modifications.pitch", "0.8",
		"--sound-modifications.semitone", "-2",
		"--sound-modifications.track", "both",
		"--stream-bidirectional-codec", "G722",
		"--stream-bidirectional-mode", "rtp",
		"--stream-bidirectional-target-legs", "both",
		"--stream-codec", "PCMA",
		"--stream-track", "both_tracks",
		"--stream-url", "wss://www.example.com/websocket",
		"--transcription=true",
		"--transcription-config.client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--transcription-config.command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--transcription-config.transcription-engine", "Google",
		"--transcription-config.transcription-engine-config", "{enable_speaker_diarization: true, hints: [string], interim_results: true, language: en, max_speaker_count: 4, min_speaker_count: 4, model: latest_long, profanity_filter: true, speech_context: [{boost: 1, phrases: [string]}], transcription_engine: Google, use_enhanced: true}",
		"--transcription-config.transcription-tracks", "both",
		"--webhook-url", "https://www.example.com/server-b/",
		"--webhook-url-method", "POST",
	)
}

func TestCallsActionsBridge(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "bridge",
		"--call-control-id-to-bridge", "call_control_id",
		"--call-control-id-to-bridge-with", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--mute-dtmf", "opposite",
		"--park-after-unbridge", "self",
		"--play-ringtone=true",
		"--queue", "support",
		"--record", "record-from-answer",
		"--record-channels", "single",
		"--record-custom-file-name", "my_recording_file_name",
		"--record-format", "wav",
		"--record-max-length", "1000",
		"--record-timeout-secs", "100",
		"--record-track", "outbound",
		"--record-trim", "trim-silence",
		"--ringtone", "pl",
		"--video-room-context", "Alice",
		"--video-room-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
	)
}

func TestCallsActionsEnqueue(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "enqueue",
		"--call-control-id", "call_control_id",
		"--queue-name", "support",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--keep-after-hangup=true",
		"--max-size", "20",
		"--max-wait-time-secs", "600",
	)
}

func TestCallsActionsGather(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "gather",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--gather-id", "my_gather_id",
		"--initial-timeout-millis", "10000",
		"--inter-digit-timeout-millis", "10000",
		"--maximum-digits", "10",
		"--minimum-digits", "1",
		"--terminating-digit", "#",
		"--timeout-millis", "60000",
		"--valid-digits", "123",
	)
}

func TestCallsActionsGatherUsingAI(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "gather-using-ai",
		"--call-control-id", "call_control_id",
		"--parameters", "{properties: bar, required: bar, type: bar}",
		"--assistant", "{instructions: You are a friendly voice assistant., model: Qwen/Qwen3-235B-A22B, openai_api_key_ref: my_openai_api_key, tools: [{book_appointment: {api_key_ref: my_calcom_api_key, event_type_id: 0, attendee_name: attendee_name, attendee_timezone: attendee_timezone}, type: book_appointment}]}",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--greeting", "Hello, can you tell me your age and where you live?",
		"--interruption-settings", "{enable: true}",
		"--language", "en",
		"--message-history", "{content: 'Hello, what''s your name?', role: assistant}",
		"--message-history", "{content: 'Hello, I''m John.', role: user}",
		"--send-message-history-updates=true",
		"--send-partial-results=true",
		"--transcription", "{model: distil-whisper/distil-large-v2}",
		"--user-response-timeout-ms", "5000",
		"--voice", "Telnyx.KokoroTTS.af",
		"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callsActionsGatherUsingAI)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "gather-using-ai",
		"--call-control-id", "call_control_id",
		"--parameters", "{properties: bar, required: bar, type: bar}",
		"--assistant.instructions", "You are a friendly voice assistant.",
		"--assistant.model", "Qwen/Qwen3-235B-A22B",
		"--assistant.openai-api-key-ref", "my_openai_api_key",
		"--assistant.tools", "[{book_appointment: {api_key_ref: my_calcom_api_key, event_type_id: 0, attendee_name: attendee_name, attendee_timezone: attendee_timezone}, type: book_appointment}]",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--greeting", "Hello, can you tell me your age and where you live?",
		"--interruption-settings.enable=true",
		"--language", "en",
		"--message-history.content", "Hello, what's your name?",
		"--message-history.role", "assistant",
		"--message-history.content", "Hello, I'm John.",
		"--message-history.role", "user",
		"--send-message-history-updates=true",
		"--send-partial-results=true",
		"--transcription.model", "distil-whisper/distil-large-v2",
		"--user-response-timeout-ms", "5000",
		"--voice", "Telnyx.KokoroTTS.af",
		"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
	)
}

func TestCallsActionsGatherUsingAudio(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "gather-using-audio",
		"--call-control-id", "call_control_id",
		"--audio-url", "http://example.com/message.wav",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--inter-digit-timeout-millis", "10000",
		"--invalid-audio-url", "http://example.com/message.wav",
		"--invalid-media-name", "my_media_uploaded_to_media_storage_api",
		"--maximum-digits", "10",
		"--maximum-tries", "3",
		"--media-name", "my_media_uploaded_to_media_storage_api",
		"--minimum-digits", "1",
		"--terminating-digit", "#",
		"--timeout-millis", "10000",
		"--valid-digits", "123",
	)
}

func TestCallsActionsGatherUsingSpeak(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "gather-using-speak",
		"--call-control-id", "call_control_id",
		"--payload", "say this on call",
		"--voice", "male",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--inter-digit-timeout-millis", "10000",
		"--invalid-payload", "say this on call",
		"--language", "arb",
		"--maximum-digits", "10",
		"--maximum-tries", "3",
		"--minimum-digits", "1",
		"--payload-type", "text",
		"--service-level", "premium",
		"--terminating-digit", "#",
		"--timeout-millis", "60000",
		"--valid-digits", "123",
		"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
	)
}

func TestCallsActionsHangup(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "hangup",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsLeaveQueue(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "leave-queue",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsPauseRecording(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "pause-recording",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--recording-id", "6e00ab49-9487-4364-8ad6-23965965afb2",
	)
}

func TestCallsActionsRefer(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "refer",
		"--call-control-id", "call_control_id",
		"--sip-address", "sip:username@sip.non-telnyx-address.com",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--custom-header", "{name: head_1, value: val_1}",
		"--custom-header", "{name: head_2, value: val_2}",
		"--sip-auth-password", "sip_auth_password",
		"--sip-auth-username", "sip_auth_username",
		"--sip-header", "{name: User-to-User, value: value}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callsActionsRefer)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "refer",
		"--call-control-id", "call_control_id",
		"--sip-address", "sip:username@sip.non-telnyx-address.com",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--custom-header.name", "head_1",
		"--custom-header.value", "val_1",
		"--custom-header.name", "head_2",
		"--custom-header.value", "val_2",
		"--sip-auth-password", "sip_auth_password",
		"--sip-auth-username", "sip_auth_username",
		"--sip-header.name", "User-to-User",
		"--sip-header.value", "value",
	)
}

func TestCallsActionsReject(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "reject",
		"--call-control-id", "call_control_id",
		"--cause", "USER_BUSY",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsResumeRecording(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "resume-recording",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--recording-id", "6e00ab49-9487-4364-8ad6-23965965afb2",
	)
}

func TestCallsActionsSendDtmf(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "send-dtmf",
		"--call-control-id", "call_control_id",
		"--digits", "1www2WABCDw9",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--duration-millis", "500",
	)
}

func TestCallsActionsSendSipInfo(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "send-sip-info",
		"--call-control-id", "call_control_id",
		"--body", `{"key": "value", "numValue": 100}`,
		"--content-type", "application/json",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsSpeak(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "speak",
		"--call-control-id", "call_control_id",
		"--payload", "Say this on the call",
		"--voice", "female",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--language", "arb",
		"--payload-type", "text",
		"--service-level", "basic",
		"--stop", "current",
		"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
	)
}

func TestCallsActionsStartAIAssistant(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-ai-assistant",
		"--call-control-id", "call_control_id",
		"--assistant", "{id: id, instructions: You are a friendly voice assistant., openai_api_key_ref: openai_api_key_ref}",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--greeting", "Hello, can you tell me your age and where you live?",
		"--interruption-settings", "{enable: true}",
		"--transcription", "{model: distil-whisper/distil-large-v2}",
		"--voice", "Telnyx.KokoroTTS.af",
		"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callsActionsStartAIAssistant)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-ai-assistant",
		"--call-control-id", "call_control_id",
		"--assistant.id", "id",
		"--assistant.instructions", "You are a friendly voice assistant.",
		"--assistant.openai-api-key-ref", "openai_api_key_ref",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--greeting", "Hello, can you tell me your age and where you live?",
		"--interruption-settings.enable=true",
		"--transcription.model", "distil-whisper/distil-large-v2",
		"--voice", "Telnyx.KokoroTTS.af",
		"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
	)
}

func TestCallsActionsStartForking(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-forking",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--rx", "udp:192.0.2.1:9000",
		"--stream-type", "decrypted",
		"--tx", "udp:192.0.2.1:9001",
	)
}

func TestCallsActionsStartNoiseSuppression(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-noise-suppression",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--direction", "both",
		"--noise-suppression-engine", "DeepFilterNet",
		"--noise-suppression-engine-config", "{attenuation_limit: 100}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callsActionsStartNoiseSuppression)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-noise-suppression",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--direction", "both",
		"--noise-suppression-engine", "DeepFilterNet",
		"--noise-suppression-engine-config.attenuation-limit", "100",
	)
}

func TestCallsActionsStartPlayback(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-playback",
		"--call-control-id", "call_control_id",
		"--audio-type", "wav",
		"--audio-url", "http://www.example.com/sounds/greeting.wav",
		"--cache-audio=true",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--loop", "infinity",
		"--media-name", "my_media_uploaded_to_media_storage_api",
		"--overlay=true",
		"--playback-content", "SUQzAwAAAAADf1...",
		"--stop", "current",
		"--target-legs", "self",
	)
}

func TestCallsActionsStartRecording(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-recording",
		"--call-control-id", "call_control_id",
		"--channels", "single",
		"--format", "wav",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--custom-file-name", "my_recording_file_name",
		"--max-length", "0",
		"--play-beep=true",
		"--recording-track", "outbound",
		"--timeout-secs", "0",
		"--transcription=true",
		"--transcription-engine", "B",
		"--transcription-language", "en",
		"--transcription-max-speaker-count", "4",
		"--transcription-min-speaker-count", "4",
		"--transcription-profanity-filter=true",
		"--transcription-speaker-diarization=true",
		"--trim", "trim-silence",
	)
}

func TestCallsActionsStartSiprec(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-siprec",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--connector-name", "my-siprec-connector",
		"--include-metadata-custom-headers=true",
		"--secure=true",
		"--session-timeout-secs", "900",
		"--sip-transport", "tcp",
		"--siprec-track", "both_tracks",
	)
}

func TestCallsActionsStartStreaming(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-streaming",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--dialogflow-config", "{analyze_sentiment: false, partial_automated_agent_reply: false}",
		"--enable-dialogflow=false",
		"--stream-bidirectional-codec", "G722",
		"--stream-bidirectional-mode", "rtp",
		"--stream-bidirectional-sampling-rate", "16000",
		"--stream-bidirectional-target-legs", "both",
		"--stream-codec", "PCMA",
		"--stream-track", "both_tracks",
		"--stream-url", "wss://www.example.com/websocket",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callsActionsStartStreaming)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-streaming",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--dialogflow-config.analyze-sentiment=false",
		"--dialogflow-config.partial-automated-agent-reply=false",
		"--enable-dialogflow=false",
		"--stream-bidirectional-codec", "G722",
		"--stream-bidirectional-mode", "rtp",
		"--stream-bidirectional-sampling-rate", "16000",
		"--stream-bidirectional-target-legs", "both",
		"--stream-codec", "PCMA",
		"--stream-track", "both_tracks",
		"--stream-url", "wss://www.example.com/websocket",
	)
}

func TestCallsActionsStartTranscription(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "start-transcription",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--transcription-engine", "Google",
		"--transcription-engine-config", "{enable_speaker_diarization: true, hints: [string], interim_results: true, language: en, max_speaker_count: 4, min_speaker_count: 4, model: latest_long, profanity_filter: true, speech_context: [{boost: 1, phrases: [string]}], transcription_engine: Google, use_enhanced: true}",
		"--transcription-tracks", "both",
	)
}

func TestCallsActionsStopAIAssistant(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-ai-assistant",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsStopForking(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-forking",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--stream-type", "decrypted",
	)
}

func TestCallsActionsStopGather(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-gather",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsStopNoiseSuppression(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-noise-suppression",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsStopPlayback(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-playback",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--overlay=false",
		"--stop", "all",
	)
}

func TestCallsActionsStopRecording(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-recording",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--recording-id", "6e00ab49-9487-4364-8ad6-23965965afb2",
	)
}

func TestCallsActionsStopSiprec(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-siprec",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsStopStreaming(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-streaming",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--stream-id", "1edb94f9-7ef0-4150-b502-e0ebadfd9491",
	)
}

func TestCallsActionsStopTranscription(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "stop-transcription",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
	)
}

func TestCallsActionsSwitchSupervisorRole(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "switch-supervisor-role",
		"--call-control-id", "call_control_id",
		"--role", "barge",
	)
}

func TestCallsActionsTransfer(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "transfer",
		"--call-control-id", "call_control_id",
		"--to", "+18005550100 or sip:username@sip.telnyx.com",
		"--answering-machine-detection", "detect",
		"--answering-machine-detection-config", "{after_greeting_silence_millis: 1000, between_words_silence_millis: 1000, greeting_duration_millis: 1000, greeting_silence_duration_millis: 2000, greeting_total_analysis_time_millis: 50000, initial_silence_millis: 1000, maximum_number_of_words: 1000, maximum_word_length_millis: 2000, silence_threshold: 512, total_analysis_time_millis: 5000}",
		"--audio-url", "http://www.example.com/sounds/greeting.wav",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--custom-header", "{name: head_1, value: val_1}",
		"--custom-header", "{name: head_2, value: val_2}",
		"--early-media=true",
		"--from", "+18005550101",
		"--from-display-name", "Company Name",
		"--media-encryption", "SRTP",
		"--media-name", "my_media_uploaded_to_media_storage_api",
		"--mute-dtmf", "opposite",
		"--park-after-unbridge", "self",
		"--record", "record-from-answer",
		"--record-channels", "single",
		"--record-custom-file-name", "my_recording_file_name",
		"--record-format", "wav",
		"--record-max-length", "1000",
		"--record-timeout-secs", "100",
		"--record-track", "outbound",
		"--record-trim", "trim-silence",
		"--sip-auth-password", "password",
		"--sip-auth-username", "username",
		"--sip-header", "{name: User-to-User, value: value}",
		"--sip-region", "Canada",
		"--sip-transport-protocol", "TLS",
		"--sound-modifications", "{octaves: 0.1, pitch: 0.8, semitone: -2, track: both}",
		"--target-leg-client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--time-limit-secs", "60",
		"--timeout-secs", "60",
		"--webhook-url", "https://www.example.com/server-b/",
		"--webhook-url-method", "POST",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callsActionsTransfer)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "transfer",
		"--call-control-id", "call_control_id",
		"--to", "+18005550100 or sip:username@sip.telnyx.com",
		"--answering-machine-detection", "detect",
		"--answering-machine-detection-config.after-greeting-silence-millis", "1000",
		"--answering-machine-detection-config.between-words-silence-millis", "1000",
		"--answering-machine-detection-config.greeting-duration-millis", "1000",
		"--answering-machine-detection-config.greeting-silence-duration-millis", "2000",
		"--answering-machine-detection-config.greeting-total-analysis-time-millis", "50000",
		"--answering-machine-detection-config.initial-silence-millis", "1000",
		"--answering-machine-detection-config.maximum-number-of-words", "1000",
		"--answering-machine-detection-config.maximum-word-length-millis", "2000",
		"--answering-machine-detection-config.silence-threshold", "512",
		"--answering-machine-detection-config.total-analysis-time-millis", "5000",
		"--audio-url", "http://www.example.com/sounds/greeting.wav",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		"--custom-header.name", "head_1",
		"--custom-header.value", "val_1",
		"--custom-header.name", "head_2",
		"--custom-header.value", "val_2",
		"--early-media=true",
		"--from", "+18005550101",
		"--from-display-name", "Company Name",
		"--media-encryption", "SRTP",
		"--media-name", "my_media_uploaded_to_media_storage_api",
		"--mute-dtmf", "opposite",
		"--park-after-unbridge", "self",
		"--record", "record-from-answer",
		"--record-channels", "single",
		"--record-custom-file-name", "my_recording_file_name",
		"--record-format", "wav",
		"--record-max-length", "1000",
		"--record-timeout-secs", "100",
		"--record-track", "outbound",
		"--record-trim", "trim-silence",
		"--sip-auth-password", "password",
		"--sip-auth-username", "username",
		"--sip-header.name", "User-to-User",
		"--sip-header.value", "value",
		"--sip-region", "Canada",
		"--sip-transport-protocol", "TLS",
		"--sound-modifications.octaves", "0.1",
		"--sound-modifications.pitch", "0.8",
		"--sound-modifications.semitone", "-2",
		"--sound-modifications.track", "both",
		"--target-leg-client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--time-limit-secs", "60",
		"--timeout-secs", "60",
		"--webhook-url", "https://www.example.com/server-b/",
		"--webhook-url-method", "POST",
	)
}

func TestCallsActionsUpdateClientState(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"calls:actions", "update-client-state",
		"--call-control-id", "call_control_id",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
	)
}

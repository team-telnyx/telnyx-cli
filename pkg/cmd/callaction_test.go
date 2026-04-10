// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestCallsActionsAddAIAssistantMessages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "add-ai-assistant-messages",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--message", "{content: Get the user's favorite color, role: system, metadata: {foo: bar}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"messages:\n" +
			"  - content: Get the user's favorite color\n" +
			"    role: system\n" +
			"    metadata:\n" +
			"      foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "add-ai-assistant-messages",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsAnswer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "answer",
			"--call-control-id", "call_control_id",
			"--assistant", "{id: asst_123, dynamic_variables: {customer_name: John, account_id: ACC-12345}, external_llm: {foo: bar}, fallback_config: {foo: bar}, greeting: 'Hi, I''m your assistant. How can I help?', instructions: You are a friendly voice assistant., llm_api_key_ref: my_llm_api_key, mcp_servers: [{foo: bar}], model: gpt-4o, name: name, observability_settings: {foo: bar}, openai_api_key_ref: my_openai_api_key, tools: [{hangup: {description: description}, type: hangup}]}",
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
			"--webhook-retries-policies", "{call.hangup: {retries_ms: [1000, 2000, 5000]}}",
			"--webhook-url", "https://www.example.com/server-b/",
			"--webhook-url-method", "POST",
			"--webhook-urls", "{call.hangup: https://www.example.com/webhooks/hangup, call.bridge: https://www.example.com/webhooks/bridge}",
			"--webhook-urls-method", "POST",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsAnswer)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "answer",
			"--call-control-id", "call_control_id",
			"--assistant.id", "asst_123",
			"--assistant.dynamic-variables", "{customer_name: John, account_id: ACC-12345}",
			"--assistant.external-llm", "{foo: bar}",
			"--assistant.fallback-config", "{foo: bar}",
			"--assistant.greeting", "Hi, I'm your assistant. How can I help?",
			"--assistant.instructions", "You are a friendly voice assistant.",
			"--assistant.llm-api-key-ref", "my_llm_api_key",
			"--assistant.mcp-servers", "[{foo: bar}]",
			"--assistant.model", "gpt-4o",
			"--assistant.name", "name",
			"--assistant.observability-settings", "{foo: bar}",
			"--assistant.openai-api-key-ref", "my_openai_api_key",
			"--assistant.tools", "[{hangup: {description: description}, type: hangup}]",
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
			"--webhook-retries-policies", "{call.hangup: {retries_ms: [1000, 2000, 5000]}}",
			"--webhook-url", "https://www.example.com/server-b/",
			"--webhook-url-method", "POST",
			"--webhook-urls", "{call.hangup: https://www.example.com/webhooks/hangup, call.bridge: https://www.example.com/webhooks/bridge}",
			"--webhook-urls-method", "POST",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"assistant:\n" +
			"  id: asst_123\n" +
			"  dynamic_variables:\n" +
			"    customer_name: John\n" +
			"    account_id: ACC-12345\n" +
			"  external_llm:\n" +
			"    foo: bar\n" +
			"  fallback_config:\n" +
			"    foo: bar\n" +
			"  greeting: Hi, I'm your assistant. How can I help?\n" +
			"  instructions: You are a friendly voice assistant.\n" +
			"  llm_api_key_ref: my_llm_api_key\n" +
			"  mcp_servers:\n" +
			"    - foo: bar\n" +
			"  model: gpt-4o\n" +
			"  name: name\n" +
			"  observability_settings:\n" +
			"    foo: bar\n" +
			"  openai_api_key_ref: my_openai_api_key\n" +
			"  tools:\n" +
			"    - hangup:\n" +
			"        description: description\n" +
			"      type: hangup\n" +
			"billing_group_id: f5586561-8ff0-4291-a0ac-84fe544797bd\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"custom_headers:\n" +
			"  - name: head_1\n" +
			"    value: val_1\n" +
			"  - name: head_2\n" +
			"    value: val_2\n" +
			"preferred_codecs: G722,PCMU,PCMA,G729,OPUS,VP8,H264\n" +
			"record: record-from-answer\n" +
			"record_channels: single\n" +
			"record_custom_file_name: my_recording_file_name\n" +
			"record_format: wav\n" +
			"record_max_length: 1000\n" +
			"record_timeout_secs: 100\n" +
			"record_track: outbound\n" +
			"record_trim: trim-silence\n" +
			"send_silence_when_idle: true\n" +
			"sip_headers:\n" +
			"  - name: User-to-User\n" +
			"    value: value\n" +
			"sound_modifications:\n" +
			"  octaves: 0.1\n" +
			"  pitch: 0.8\n" +
			"  semitone: -2\n" +
			"  track: both\n" +
			"stream_bidirectional_codec: G722\n" +
			"stream_bidirectional_mode: rtp\n" +
			"stream_bidirectional_target_legs: both\n" +
			"stream_codec: PCMA\n" +
			"stream_track: both_tracks\n" +
			"stream_url: wss://www.example.com/websocket\n" +
			"transcription: true\n" +
			"transcription_config:\n" +
			"  client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"  command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"  transcription_engine: Google\n" +
			"  transcription_engine_config:\n" +
			"    enable_speaker_diarization: true\n" +
			"    hints:\n" +
			"      - string\n" +
			"    interim_results: true\n" +
			"    language: en\n" +
			"    max_speaker_count: 4\n" +
			"    min_speaker_count: 4\n" +
			"    model: latest_long\n" +
			"    profanity_filter: true\n" +
			"    speech_context:\n" +
			"      - boost: 1\n" +
			"        phrases:\n" +
			"          - string\n" +
			"    transcription_engine: Google\n" +
			"    use_enhanced: true\n" +
			"  transcription_tracks: both\n" +
			"webhook_retries_policies:\n" +
			"  call.hangup:\n" +
			"    retries_ms:\n" +
			"      - 1000\n" +
			"      - 2000\n" +
			"      - 5000\n" +
			"webhook_url: https://www.example.com/server-b/\n" +
			"webhook_url_method: POST\n" +
			"webhook_urls:\n" +
			"  call.hangup: https://www.example.com/webhooks/hangup\n" +
			"  call.bridge: https://www.example.com/webhooks/bridge\n" +
			"webhook_urls_method: POST\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "answer",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsBridge(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "bridge",
			"--call-control-id-to-bridge", "call_control_id",
			"--call-control-id-to-bridge-with", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--hold-after-unbridge=true",
			"--mute-dtmf", "opposite",
			"--park-after-unbridge", "self",
			"--play-ringtone=true",
			"--prevent-double-bridge=true",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_control_id: v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"hold_after_unbridge: true\n" +
			"mute_dtmf: opposite\n" +
			"park_after_unbridge: self\n" +
			"play_ringtone: true\n" +
			"prevent_double_bridge: true\n" +
			"queue: support\n" +
			"record: record-from-answer\n" +
			"record_channels: single\n" +
			"record_custom_file_name: my_recording_file_name\n" +
			"record_format: wav\n" +
			"record_max_length: 1000\n" +
			"record_timeout_secs: 100\n" +
			"record_track: outbound\n" +
			"record_trim: trim-silence\n" +
			"ringtone: pl\n" +
			"video_room_context: Alice\n" +
			"video_room_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "bridge",
			"--call-control-id-to-bridge", "call_control_id",
		)
	})
}

func TestCallsActionsEnqueue(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "enqueue",
			"--call-control-id", "call_control_id",
			"--queue-name", "support",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--keep-after-hangup=true",
			"--max-size", "20",
			"--max-wait-time-secs", "600",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"queue_name: support\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"keep_after_hangup: true\n" +
			"max_size: 20\n" +
			"max_wait_time_secs: 600\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "enqueue",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsGather(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"gather_id: my_gather_id\n" +
			"initial_timeout_millis: 10000\n" +
			"inter_digit_timeout_millis: 10000\n" +
			"maximum_digits: 10\n" +
			"minimum_digits: 1\n" +
			"terminating_digit: '#'\n" +
			"timeout_millis: 60000\n" +
			"valid_digits: '123'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "gather",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsGatherUsingAI(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "gather-using-ai",
			"--call-control-id", "call_control_id",
			"--parameters", "{properties: bar, required: bar, type: bar}",
			"--assistant", "{instructions: You are a friendly voice assistant., model: Qwen/Qwen3-235B-A22B, openai_api_key_ref: my_openai_api_key, tools: [{book_appointment: {api_key_ref: my_calcom_api_key, event_type_id: 0, attendee_name: attendee_name, attendee_timezone: attendee_timezone}, type: book_appointment}]}",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--gather-ended-speech", "Thank you for providing the information.",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsGatherUsingAI)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "gather-using-ai",
			"--call-control-id", "call_control_id",
			"--parameters", "{properties: bar, required: bar, type: bar}",
			"--assistant.instructions", "You are a friendly voice assistant.",
			"--assistant.model", "Qwen/Qwen3-235B-A22B",
			"--assistant.openai-api-key-ref", "my_openai_api_key",
			"--assistant.tools", "[{book_appointment: {api_key_ref: my_calcom_api_key, event_type_id: 0, attendee_name: attendee_name, attendee_timezone: attendee_timezone}, type: book_appointment}]",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--gather-ended-speech", "Thank you for providing the information.",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"parameters:\n" +
			"  properties: bar\n" +
			"  required: bar\n" +
			"  type: bar\n" +
			"assistant:\n" +
			"  instructions: You are a friendly voice assistant.\n" +
			"  model: Qwen/Qwen3-235B-A22B\n" +
			"  openai_api_key_ref: my_openai_api_key\n" +
			"  tools:\n" +
			"    - book_appointment:\n" +
			"        api_key_ref: my_calcom_api_key\n" +
			"        event_type_id: 0\n" +
			"        attendee_name: attendee_name\n" +
			"        attendee_timezone: attendee_timezone\n" +
			"      type: book_appointment\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"gather_ended_speech: Thank you for providing the information.\n" +
			"greeting: Hello, can you tell me your age and where you live?\n" +
			"interruption_settings:\n" +
			"  enable: true\n" +
			"language: en\n" +
			"message_history:\n" +
			"  - content: Hello, what's your name?\n" +
			"    role: assistant\n" +
			"  - content: Hello, I'm John.\n" +
			"    role: user\n" +
			"send_message_history_updates: true\n" +
			"send_partial_results: true\n" +
			"transcription:\n" +
			"  model: distil-whisper/distil-large-v2\n" +
			"user_response_timeout_ms: 5000\n" +
			"voice: Telnyx.KokoroTTS.af\n" +
			"voice_settings:\n" +
			"  type: elevenlabs\n" +
			"  api_key_ref: my_elevenlabs_api_key\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "gather-using-ai",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsGatherUsingAudio(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"audio_url: http://example.com/message.wav\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"inter_digit_timeout_millis: 10000\n" +
			"invalid_audio_url: http://example.com/message.wav\n" +
			"invalid_media_name: my_media_uploaded_to_media_storage_api\n" +
			"maximum_digits: 10\n" +
			"maximum_tries: 3\n" +
			"media_name: my_media_uploaded_to_media_storage_api\n" +
			"minimum_digits: 1\n" +
			"terminating_digit: '#'\n" +
			"timeout_millis: 10000\n" +
			"valid_digits: '123'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "gather-using-audio",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsGatherUsingSpeak(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"payload: say this on call\n" +
			"voice: male\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"inter_digit_timeout_millis: 10000\n" +
			"invalid_payload: say this on call\n" +
			"language: arb\n" +
			"maximum_digits: 10\n" +
			"maximum_tries: 3\n" +
			"minimum_digits: 1\n" +
			"payload_type: text\n" +
			"service_level: premium\n" +
			"terminating_digit: '#'\n" +
			"timeout_millis: 60000\n" +
			"valid_digits: '123'\n" +
			"voice_settings:\n" +
			"  type: elevenlabs\n" +
			"  api_key_ref: my_elevenlabs_api_key\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "gather-using-speak",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsHangup(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "hangup",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--custom-header", "{name: head_1, value: val_1}",
			"--custom-header", "{name: head_2, value: val_2}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsHangup)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "hangup",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--custom-header.name", "head_1",
			"--custom-header.value", "val_1",
			"--custom-header.name", "head_2",
			"--custom-header.value", "val_2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"custom_headers:\n" +
			"  - name: head_1\n" +
			"    value: val_1\n" +
			"  - name: head_2\n" +
			"    value: val_2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "hangup",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsJoinAIAssistant(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "join-ai-assistant",
			"--call-control-id", "call_control_id",
			"--conversation-id", "v3:abc123",
			"--participant", "{id: v3:abc123def456, role: user, name: John Doe, on_hangup: continue_conversation}",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsJoinAIAssistant)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "join-ai-assistant",
			"--call-control-id", "call_control_id",
			"--conversation-id", "v3:abc123",
			"--participant.id", "v3:abc123def456",
			"--participant.role", "user",
			"--participant.name", "John Doe",
			"--participant.on-hangup", "continue_conversation",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"conversation_id: v3:abc123\n" +
			"participant:\n" +
			"  id: v3:abc123def456\n" +
			"  role: user\n" +
			"  name: John Doe\n" +
			"  on_hangup: continue_conversation\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "join-ai-assistant",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsLeaveQueue(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "leave-queue",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "leave-queue",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsPauseRecording(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "pause-recording",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--recording-id", "6e00ab49-9487-4364-8ad6-23965965afb2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"recording_id: 6e00ab49-9487-4364-8ad6-23965965afb2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "pause-recording",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsRefer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsRefer)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sip_address: sip:username@sip.non-telnyx-address.com\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"custom_headers:\n" +
			"  - name: head_1\n" +
			"    value: val_1\n" +
			"  - name: head_2\n" +
			"    value: val_2\n" +
			"sip_auth_password: sip_auth_password\n" +
			"sip_auth_username: sip_auth_username\n" +
			"sip_headers:\n" +
			"  - name: User-to-User\n" +
			"    value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "refer",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsReject(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "reject",
			"--call-control-id", "call_control_id",
			"--cause", "USER_BUSY",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cause: USER_BUSY\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "reject",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsResumeRecording(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "resume-recording",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--recording-id", "6e00ab49-9487-4364-8ad6-23965965afb2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"recording_id: 6e00ab49-9487-4364-8ad6-23965965afb2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "resume-recording",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsSendDtmf(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "send-dtmf",
			"--call-control-id", "call_control_id",
			"--digits", "1www2WABCDw9",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--duration-millis", "500",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"digits: 1www2WABCDw9\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"duration_millis: 500\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "send-dtmf",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsSendSipInfo(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "send-sip-info",
			"--call-control-id", "call_control_id",
			"--body", `{"key": "value", "numValue": 100}`,
			"--content-type", "application/json",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"body: '{\"key\": \"value\", \"numValue\": 100}'\n" +
			"content_type: application/json\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "send-sip-info",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsSpeak(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "speak",
			"--call-control-id", "call_control_id",
			"--payload", "Say this on the call",
			"--voice", "female",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--language", "arb",
			"--loop", "string",
			"--payload-type", "text",
			"--service-level", "basic",
			"--stop", "current",
			"--target-legs", "both",
			"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"payload: Say this on the call\n" +
			"voice: female\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"language: arb\n" +
			"loop: string\n" +
			"payload_type: text\n" +
			"service_level: basic\n" +
			"stop: current\n" +
			"target_legs: both\n" +
			"voice_settings:\n" +
			"  type: elevenlabs\n" +
			"  api_key_ref: my_elevenlabs_api_key\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "speak",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStartAIAssistant(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "start-ai-assistant",
			"--call-control-id", "call_control_id",
			"--assistant", "{id: id, dynamic_variables: {customer_name: John, account_id: ACC-12345}, external_llm: {foo: bar}, fallback_config: {foo: bar}, greeting: greeting, instructions: You are a friendly voice assistant., llm_api_key_ref: my_llm_api_key, mcp_servers: [{foo: bar}], model: gpt-4o, name: name, observability_settings: {foo: bar}, openai_api_key_ref: my_openai_api_key, tools: [{book_appointment: {api_key_ref: my_calcom_api_key, event_type_id: 0, attendee_name: attendee_name, attendee_timezone: attendee_timezone}, type: book_appointment}]}",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--greeting", "Hello, can you tell me your age and where you live?",
			"--interruption-settings", "{enable: true}",
			"--message-history", "{content: 'Hello, I would like some help.', role: user, metadata: {foo: bar}}",
			"--participant", "{id: v3:abc123def456, role: user, name: John Doe, on_hangup: continue_conversation}",
			"--send-message-history-updates=true",
			"--transcription", "{model: distil-whisper/distil-large-v2}",
			"--voice", "Telnyx.KokoroTTS.af",
			"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsStartAIAssistant)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "start-ai-assistant",
			"--call-control-id", "call_control_id",
			"--assistant.id", "id",
			"--assistant.dynamic-variables", "{customer_name: John, account_id: ACC-12345}",
			"--assistant.external-llm", "{foo: bar}",
			"--assistant.fallback-config", "{foo: bar}",
			"--assistant.greeting", "greeting",
			"--assistant.instructions", "You are a friendly voice assistant.",
			"--assistant.llm-api-key-ref", "my_llm_api_key",
			"--assistant.mcp-servers", "[{foo: bar}]",
			"--assistant.model", "gpt-4o",
			"--assistant.name", "name",
			"--assistant.observability-settings", "{foo: bar}",
			"--assistant.openai-api-key-ref", "my_openai_api_key",
			"--assistant.tools", "[{book_appointment: {api_key_ref: my_calcom_api_key, event_type_id: 0, attendee_name: attendee_name, attendee_timezone: attendee_timezone}, type: book_appointment}]",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--greeting", "Hello, can you tell me your age and where you live?",
			"--interruption-settings.enable=true",
			"--message-history", "{content: 'Hello, I would like some help.', role: user, metadata: {foo: bar}}",
			"--participant.id", "v3:abc123def456",
			"--participant.role", "user",
			"--participant.name", "John Doe",
			"--participant.on-hangup", "continue_conversation",
			"--send-message-history-updates=true",
			"--transcription.model", "distil-whisper/distil-large-v2",
			"--voice", "Telnyx.KokoroTTS.af",
			"--voice-settings", "{type: elevenlabs, api_key_ref: my_elevenlabs_api_key}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"assistant:\n" +
			"  id: id\n" +
			"  dynamic_variables:\n" +
			"    customer_name: John\n" +
			"    account_id: ACC-12345\n" +
			"  external_llm:\n" +
			"    foo: bar\n" +
			"  fallback_config:\n" +
			"    foo: bar\n" +
			"  greeting: greeting\n" +
			"  instructions: You are a friendly voice assistant.\n" +
			"  llm_api_key_ref: my_llm_api_key\n" +
			"  mcp_servers:\n" +
			"    - foo: bar\n" +
			"  model: gpt-4o\n" +
			"  name: name\n" +
			"  observability_settings:\n" +
			"    foo: bar\n" +
			"  openai_api_key_ref: my_openai_api_key\n" +
			"  tools:\n" +
			"    - book_appointment:\n" +
			"        api_key_ref: my_calcom_api_key\n" +
			"        event_type_id: 0\n" +
			"        attendee_name: attendee_name\n" +
			"        attendee_timezone: attendee_timezone\n" +
			"      type: book_appointment\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"greeting: Hello, can you tell me your age and where you live?\n" +
			"interruption_settings:\n" +
			"  enable: true\n" +
			"message_history:\n" +
			"  - content: Hello, I would like some help.\n" +
			"    role: user\n" +
			"    metadata:\n" +
			"      foo: bar\n" +
			"participants:\n" +
			"  - id: v3:abc123def456\n" +
			"    role: user\n" +
			"    name: John Doe\n" +
			"    on_hangup: continue_conversation\n" +
			"send_message_history_updates: true\n" +
			"transcription:\n" +
			"  model: distil-whisper/distil-large-v2\n" +
			"voice: Telnyx.KokoroTTS.af\n" +
			"voice_settings:\n" +
			"  type: elevenlabs\n" +
			"  api_key_ref: my_elevenlabs_api_key\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "start-ai-assistant",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStartForking(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "start-forking",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--rx", "udp:192.0.2.1:9000",
			"--stream-type", "decrypted",
			"--tx", "udp:192.0.2.1:9001",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"rx: udp:192.0.2.1:9000\n" +
			"stream_type: decrypted\n" +
			"tx: udp:192.0.2.1:9001\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "start-forking",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStartNoiseSuppression(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "start-noise-suppression",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--direction", "both",
			"--noise-suppression-engine", "DeepFilterNet",
			"--noise-suppression-engine-config", "{attenuation_limit: 100, enhancement_level: 0.5, family: sparrow, mode: standard, model: krisp-viva-tel-v2.kef, size: l, suppression_level: 50, voice_gain: 1}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsStartNoiseSuppression)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "start-noise-suppression",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--direction", "both",
			"--noise-suppression-engine", "DeepFilterNet",
			"--noise-suppression-engine-config.attenuation-limit", "100",
			"--noise-suppression-engine-config.enhancement-level", "0.5",
			"--noise-suppression-engine-config.family", "sparrow",
			"--noise-suppression-engine-config.mode", "standard",
			"--noise-suppression-engine-config.model", "krisp-viva-tel-v2.kef",
			"--noise-suppression-engine-config.size", "l",
			"--noise-suppression-engine-config.suppression-level", "50",
			"--noise-suppression-engine-config.voice-gain", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"direction: both\n" +
			"noise_suppression_engine: DeepFilterNet\n" +
			"noise_suppression_engine_config:\n" +
			"  attenuation_limit: 100\n" +
			"  enhancement_level: 0.5\n" +
			"  family: sparrow\n" +
			"  mode: standard\n" +
			"  model: krisp-viva-tel-v2.kef\n" +
			"  size: l\n" +
			"  suppression_level: 50\n" +
			"  voice_gain: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "start-noise-suppression",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStartPlayback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"audio_type: wav\n" +
			"audio_url: http://www.example.com/sounds/greeting.wav\n" +
			"cache_audio: true\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"loop: infinity\n" +
			"media_name: my_media_uploaded_to_media_storage_api\n" +
			"overlay: true\n" +
			"playback_content: SUQzAwAAAAADf1...\n" +
			"stop: current\n" +
			"target_legs: self\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "start-playback",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStartRecording(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"channels: single\n" +
			"format: wav\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"custom_file_name: my_recording_file_name\n" +
			"max_length: 0\n" +
			"play_beep: true\n" +
			"recording_track: outbound\n" +
			"timeout_secs: 0\n" +
			"transcription: true\n" +
			"transcription_engine: B\n" +
			"transcription_language: en\n" +
			"transcription_max_speaker_count: 4\n" +
			"transcription_min_speaker_count: 4\n" +
			"transcription_profanity_filter: true\n" +
			"transcription_speaker_diarization: true\n" +
			"trim: trim-silence\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "start-recording",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStartSiprec(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"connector_name: my-siprec-connector\n" +
			"include_metadata_custom_headers: true\n" +
			"secure: true\n" +
			"session_timeout_secs: 900\n" +
			"sip_transport: tcp\n" +
			"siprec_track: both_tracks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "start-siprec",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStartStreaming(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "start-streaming",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--custom-parameter", "{name: param1, value: value1}",
			"--custom-parameter", "{name: param2, value: value2}",
			"--dialogflow-config", "{analyze_sentiment: false, partial_automated_agent_reply: false}",
			"--enable-dialogflow=false",
			"--stream-auth-token", "your-auth-token",
			"--stream-bidirectional-codec", "G722",
			"--stream-bidirectional-mode", "rtp",
			"--stream-bidirectional-sampling-rate", "16000",
			"--stream-bidirectional-target-legs", "both",
			"--stream-codec", "PCMA",
			"--stream-track", "both_tracks",
			"--stream-url", "wss://www.example.com/websocket",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsStartStreaming)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "start-streaming",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--custom-parameter.name", "param1",
			"--custom-parameter.value", "value1",
			"--custom-parameter.name", "param2",
			"--custom-parameter.value", "value2",
			"--dialogflow-config.analyze-sentiment=false",
			"--dialogflow-config.partial-automated-agent-reply=false",
			"--enable-dialogflow=false",
			"--stream-auth-token", "your-auth-token",
			"--stream-bidirectional-codec", "G722",
			"--stream-bidirectional-mode", "rtp",
			"--stream-bidirectional-sampling-rate", "16000",
			"--stream-bidirectional-target-legs", "both",
			"--stream-codec", "PCMA",
			"--stream-track", "both_tracks",
			"--stream-url", "wss://www.example.com/websocket",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"custom_parameters:\n" +
			"  - name: param1\n" +
			"    value: value1\n" +
			"  - name: param2\n" +
			"    value: value2\n" +
			"dialogflow_config:\n" +
			"  analyze_sentiment: false\n" +
			"  partial_automated_agent_reply: false\n" +
			"enable_dialogflow: false\n" +
			"stream_auth_token: your-auth-token\n" +
			"stream_bidirectional_codec: G722\n" +
			"stream_bidirectional_mode: rtp\n" +
			"stream_bidirectional_sampling_rate: 16000\n" +
			"stream_bidirectional_target_legs: both\n" +
			"stream_codec: PCMA\n" +
			"stream_track: both_tracks\n" +
			"stream_url: wss://www.example.com/websocket\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "start-streaming",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStartTranscription(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "start-transcription",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--transcription-engine", "Google",
			"--transcription-engine-config", "{enable_speaker_diarization: true, hints: [string], interim_results: true, language: en, max_speaker_count: 4, min_speaker_count: 4, model: latest_long, profanity_filter: true, speech_context: [{boost: 1, phrases: [string]}], transcription_engine: Google, use_enhanced: true}",
			"--transcription-tracks", "both",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"transcription_engine: Google\n" +
			"transcription_engine_config:\n" +
			"  enable_speaker_diarization: true\n" +
			"  hints:\n" +
			"    - string\n" +
			"  interim_results: true\n" +
			"  language: en\n" +
			"  max_speaker_count: 4\n" +
			"  min_speaker_count: 4\n" +
			"  model: latest_long\n" +
			"  profanity_filter: true\n" +
			"  speech_context:\n" +
			"    - boost: 1\n" +
			"      phrases:\n" +
			"        - string\n" +
			"  transcription_engine: Google\n" +
			"  use_enhanced: true\n" +
			"transcription_tracks: both\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "start-transcription",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopAIAssistant(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-ai-assistant",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-ai-assistant",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopForking(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-forking",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--stream-type", "decrypted",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"stream_type: decrypted\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-forking",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopGather(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-gather",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-gather",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopNoiseSuppression(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-noise-suppression",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-noise-suppression",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopPlayback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-playback",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--overlay=false",
			"--stop", "all",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"overlay: false\n" +
			"stop: all\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-playback",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopRecording(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-recording",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--recording-id", "6e00ab49-9487-4364-8ad6-23965965afb2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"recording_id: 6e00ab49-9487-4364-8ad6-23965965afb2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-recording",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopSiprec(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-siprec",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-siprec",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopStreaming(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-streaming",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--stream-id", "1edb94f9-7ef0-4150-b502-e0ebadfd9491",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"stream_id: 1edb94f9-7ef0-4150-b502-e0ebadfd9491\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-streaming",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsStopTranscription(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "stop-transcription",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "stop-transcription",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsSwitchSupervisorRole(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "switch-supervisor-role",
			"--call-control-id", "call_control_id",
			"--role", "barge",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("role: barge")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "switch-supervisor-role",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsTransfer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
			"--preferred-codecs", "G722,PCMU,PCMA,G729,OPUS,VP8,H264",
			"--privacy", "id",
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
			"--webhook-retries-policies", "{call.answered: {retries_ms: [1000, 2000, 5000]}}",
			"--webhook-url", "https://www.example.com/server-b/",
			"--webhook-url-method", "POST",
			"--webhook-urls", "{call.answered: https://www.example.com/webhooks/answered, call.hangup: https://www.example.com/webhooks/hangup}",
			"--webhook-urls-method", "POST",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsActionsTransfer)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
			"--preferred-codecs", "G722,PCMU,PCMA,G729,OPUS,VP8,H264",
			"--privacy", "id",
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
			"--webhook-retries-policies", "{call.answered: {retries_ms: [1000, 2000, 5000]}}",
			"--webhook-url", "https://www.example.com/server-b/",
			"--webhook-url-method", "POST",
			"--webhook-urls", "{call.answered: https://www.example.com/webhooks/answered, call.hangup: https://www.example.com/webhooks/hangup}",
			"--webhook-urls-method", "POST",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"to: +18005550100 or sip:username@sip.telnyx.com\n" +
			"answering_machine_detection: detect\n" +
			"answering_machine_detection_config:\n" +
			"  after_greeting_silence_millis: 1000\n" +
			"  between_words_silence_millis: 1000\n" +
			"  greeting_duration_millis: 1000\n" +
			"  greeting_silence_duration_millis: 2000\n" +
			"  greeting_total_analysis_time_millis: 50000\n" +
			"  initial_silence_millis: 1000\n" +
			"  maximum_number_of_words: 1000\n" +
			"  maximum_word_length_millis: 2000\n" +
			"  silence_threshold: 512\n" +
			"  total_analysis_time_millis: 5000\n" +
			"audio_url: http://www.example.com/sounds/greeting.wav\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"custom_headers:\n" +
			"  - name: head_1\n" +
			"    value: val_1\n" +
			"  - name: head_2\n" +
			"    value: val_2\n" +
			"early_media: true\n" +
			"from: '+18005550101'\n" +
			"from_display_name: Company Name\n" +
			"media_encryption: SRTP\n" +
			"media_name: my_media_uploaded_to_media_storage_api\n" +
			"mute_dtmf: opposite\n" +
			"park_after_unbridge: self\n" +
			"preferred_codecs: G722,PCMU,PCMA,G729,OPUS,VP8,H264\n" +
			"privacy: id\n" +
			"record: record-from-answer\n" +
			"record_channels: single\n" +
			"record_custom_file_name: my_recording_file_name\n" +
			"record_format: wav\n" +
			"record_max_length: 1000\n" +
			"record_timeout_secs: 100\n" +
			"record_track: outbound\n" +
			"record_trim: trim-silence\n" +
			"sip_auth_password: password\n" +
			"sip_auth_username: username\n" +
			"sip_headers:\n" +
			"  - name: User-to-User\n" +
			"    value: value\n" +
			"sip_region: Canada\n" +
			"sip_transport_protocol: TLS\n" +
			"sound_modifications:\n" +
			"  octaves: 0.1\n" +
			"  pitch: 0.8\n" +
			"  semitone: -2\n" +
			"  track: both\n" +
			"target_leg_client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"time_limit_secs: 60\n" +
			"timeout_secs: 60\n" +
			"webhook_retries_policies:\n" +
			"  call.answered:\n" +
			"    retries_ms:\n" +
			"      - 1000\n" +
			"      - 2000\n" +
			"      - 5000\n" +
			"webhook_url: https://www.example.com/server-b/\n" +
			"webhook_url_method: POST\n" +
			"webhook_urls:\n" +
			"  call.answered: https://www.example.com/webhooks/answered\n" +
			"  call.hangup: https://www.example.com/webhooks/hangup\n" +
			"webhook_urls_method: POST\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "transfer",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestCallsActionsUpdateClientState(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls:actions", "update-client-state",
			"--call-control-id", "call_control_id",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("client_state: aGF2ZSBhIG5pY2UgZGF5ID1d")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls:actions", "update-client-state",
			"--call-control-id", "call_control_id",
		)
	})
}

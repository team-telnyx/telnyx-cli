// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestCallsDial(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "dial",
			"--connection-id", "7267xxxxxxxxxxxxxx",
			"--from", "+18005550101",
			"--to", "+18005550100 or sip:username@sip.telnyx.com",
			"--answering-machine-detection", "detect",
			"--answering-machine-detection-config", "{after_greeting_silence_millis: 1000, between_words_silence_millis: 1000, greeting_duration_millis: 1000, greeting_silence_duration_millis: 2000, greeting_total_analysis_time_millis: 50000, initial_silence_millis: 1000, maximum_number_of_words: 1000, maximum_word_length_millis: 2000, silence_threshold: 512, total_analysis_time_millis: 5000}",
			"--audio-url", "http://www.example.com/sounds/greeting.wav",
			"--billing-group-id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
			"--bridge-intent=true",
			"--bridge-on-answer=true",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--conference-config", "{id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, beep_enabled: on_exit, conference_name: telnyx-conference, early_media: false, end_conference_on_exit: true, hold: true, hold_audio_url: http://example.com/message.wav, hold_media_name: my_media_uploaded_to_media_storage_api, mute: true, soft_end_conference_on_exit: true, start_conference_on_create: false, start_conference_on_enter: true, supervisor_role: whisper, whisper_call_control_ids: [v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ, v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw]}",
			"--custom-header", "{name: head_1, value: val_1}",
			"--custom-header", "{name: head_2, value: val_2}",
			"--dialogflow-config", "{analyze_sentiment: false, partial_automated_agent_reply: false}",
			"--enable-dialogflow=false",
			"--from-display-name", "Company Name",
			"--link-to", "ilditnZK_eVysupV21KzmzN_sM29ygfauQojpm4BgFtfX5hXAcjotg==",
			"--media-encryption", "SRTP",
			"--media-name", "my_media_uploaded_to_media_storage_api",
			"--park-after-unbridge", "self",
			"--preferred-codecs", "G722,PCMU,PCMA,G729,OPUS,VP8,H264",
			"--prevent-double-bridge=true",
			"--privacy", "id",
			"--record", "record-from-answer",
			"--record-channels", "single",
			"--record-custom-file-name", "my_recording_file_name",
			"--record-format", "wav",
			"--record-max-length", "1000",
			"--record-timeout-secs", "100",
			"--record-track", "outbound",
			"--record-trim", "trim-silence",
			"--send-silence-when-idle=true",
			"--sip-auth-password", "password",
			"--sip-auth-username", "username",
			"--sip-header", "{name: User-to-User, value: '12345'}",
			"--sip-region", "Canada",
			"--sip-transport-protocol", "TLS",
			"--sound-modifications", "{octaves: 0.1, pitch: 0.8, semitone: -2, track: both}",
			"--stream-auth-token", "your-auth-token",
			"--stream-bidirectional-codec", "G722",
			"--stream-bidirectional-mode", "rtp",
			"--stream-bidirectional-sampling-rate", "16000",
			"--stream-bidirectional-target-legs", "both",
			"--stream-codec", "PCMA",
			"--stream-establish-before-call-originate=true",
			"--stream-track", "both_tracks",
			"--stream-url", "wss://www.example.com/websocket",
			"--supervise-call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--supervisor-role", "barge",
			"--time-limit-secs", "60",
			"--timeout-secs", "60",
			"--transcription=true",
			"--transcription-config", "{client_state: aGF2ZSBhIG5pY2UgZGF5ID1d, command_id: 891510ac-f3e4-11e8-af5b-de00688a4901, transcription_engine: Google, transcription_engine_config: {enable_speaker_diarization: true, hints: [string], interim_results: true, language: en, max_speaker_count: 4, min_speaker_count: 4, model: latest_long, profanity_filter: true, speech_context: [{boost: 1, phrases: [string]}], transcription_engine: Google, use_enhanced: true}, transcription_tracks: both}",
			"--webhook-url", "https://www.example.com/server-b/",
			"--webhook-url-method", "POST",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(callsDial)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "dial",
			"--connection-id", "7267xxxxxxxxxxxxxx",
			"--from", "+18005550101",
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
			"--billing-group-id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
			"--bridge-intent=true",
			"--bridge-on-answer=true",
			"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--conference-config.id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
			"--conference-config.beep-enabled", "on_exit",
			"--conference-config.conference-name", "telnyx-conference",
			"--conference-config.early-media=false",
			"--conference-config.end-conference-on-exit=true",
			"--conference-config.hold=true",
			"--conference-config.hold-audio-url", "http://example.com/message.wav",
			"--conference-config.hold-media-name", "my_media_uploaded_to_media_storage_api",
			"--conference-config.mute=true",
			"--conference-config.soft-end-conference-on-exit=true",
			"--conference-config.start-conference-on-create=false",
			"--conference-config.start-conference-on-enter=true",
			"--conference-config.supervisor-role", "whisper",
			"--conference-config.whisper-call-control-ids", "[v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ, v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw]",
			"--custom-header.name", "head_1",
			"--custom-header.value", "val_1",
			"--custom-header.name", "head_2",
			"--custom-header.value", "val_2",
			"--dialogflow-config.analyze-sentiment=false",
			"--dialogflow-config.partial-automated-agent-reply=false",
			"--enable-dialogflow=false",
			"--from-display-name", "Company Name",
			"--link-to", "ilditnZK_eVysupV21KzmzN_sM29ygfauQojpm4BgFtfX5hXAcjotg==",
			"--media-encryption", "SRTP",
			"--media-name", "my_media_uploaded_to_media_storage_api",
			"--park-after-unbridge", "self",
			"--preferred-codecs", "G722,PCMU,PCMA,G729,OPUS,VP8,H264",
			"--prevent-double-bridge=true",
			"--privacy", "id",
			"--record", "record-from-answer",
			"--record-channels", "single",
			"--record-custom-file-name", "my_recording_file_name",
			"--record-format", "wav",
			"--record-max-length", "1000",
			"--record-timeout-secs", "100",
			"--record-track", "outbound",
			"--record-trim", "trim-silence",
			"--send-silence-when-idle=true",
			"--sip-auth-password", "password",
			"--sip-auth-username", "username",
			"--sip-header.name", "User-to-User",
			"--sip-header.value", "12345",
			"--sip-region", "Canada",
			"--sip-transport-protocol", "TLS",
			"--sound-modifications.octaves", "0.1",
			"--sound-modifications.pitch", "0.8",
			"--sound-modifications.semitone", "-2",
			"--sound-modifications.track", "both",
			"--stream-auth-token", "your-auth-token",
			"--stream-bidirectional-codec", "G722",
			"--stream-bidirectional-mode", "rtp",
			"--stream-bidirectional-sampling-rate", "16000",
			"--stream-bidirectional-target-legs", "both",
			"--stream-codec", "PCMA",
			"--stream-establish-before-call-originate=true",
			"--stream-track", "both_tracks",
			"--stream-url", "wss://www.example.com/websocket",
			"--supervise-call-control-id", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--supervisor-role", "barge",
			"--time-limit-secs", "60",
			"--timeout-secs", "60",
			"--transcription=true",
			"--transcription-config.client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
			"--transcription-config.command-id", "891510ac-f3e4-11e8-af5b-de00688a4901",
			"--transcription-config.transcription-engine", "Google",
			"--transcription-config.transcription-engine-config", "{enable_speaker_diarization: true, hints: [string], interim_results: true, language: en, max_speaker_count: 4, min_speaker_count: 4, model: latest_long, profanity_filter: true, speech_context: [{boost: 1, phrases: [string]}], transcription_engine: Google, use_enhanced: true}",
			"--transcription-config.transcription-tracks", "both",
			"--webhook-url", "https://www.example.com/server-b/",
			"--webhook-url-method", "POST",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connection_id: 7267xxxxxxxxxxxxxx\n" +
			"from: '+18005550101'\n" +
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
			"billing_group_id: f5586561-8ff0-4291-a0ac-84fe544797bd\n" +
			"bridge_intent: true\n" +
			"bridge_on_answer: true\n" +
			"client_state: aGF2ZSBhIG5pY2UgZGF5ID1d\n" +
			"command_id: 891510ac-f3e4-11e8-af5b-de00688a4901\n" +
			"conference_config:\n" +
			"  id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0\n" +
			"  beep_enabled: on_exit\n" +
			"  conference_name: telnyx-conference\n" +
			"  early_media: false\n" +
			"  end_conference_on_exit: true\n" +
			"  hold: true\n" +
			"  hold_audio_url: http://example.com/message.wav\n" +
			"  hold_media_name: my_media_uploaded_to_media_storage_api\n" +
			"  mute: true\n" +
			"  soft_end_conference_on_exit: true\n" +
			"  start_conference_on_create: false\n" +
			"  start_conference_on_enter: true\n" +
			"  supervisor_role: whisper\n" +
			"  whisper_call_control_ids:\n" +
			"    - v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ\n" +
			"    - v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw\n" +
			"custom_headers:\n" +
			"  - name: head_1\n" +
			"    value: val_1\n" +
			"  - name: head_2\n" +
			"    value: val_2\n" +
			"dialogflow_config:\n" +
			"  analyze_sentiment: false\n" +
			"  partial_automated_agent_reply: false\n" +
			"enable_dialogflow: false\n" +
			"from_display_name: Company Name\n" +
			"link_to: ilditnZK_eVysupV21KzmzN_sM29ygfauQojpm4BgFtfX5hXAcjotg==\n" +
			"media_encryption: SRTP\n" +
			"media_name: my_media_uploaded_to_media_storage_api\n" +
			"park_after_unbridge: self\n" +
			"preferred_codecs: G722,PCMU,PCMA,G729,OPUS,VP8,H264\n" +
			"prevent_double_bridge: true\n" +
			"privacy: id\n" +
			"record: record-from-answer\n" +
			"record_channels: single\n" +
			"record_custom_file_name: my_recording_file_name\n" +
			"record_format: wav\n" +
			"record_max_length: 1000\n" +
			"record_timeout_secs: 100\n" +
			"record_track: outbound\n" +
			"record_trim: trim-silence\n" +
			"send_silence_when_idle: true\n" +
			"sip_auth_password: password\n" +
			"sip_auth_username: username\n" +
			"sip_headers:\n" +
			"  - name: User-to-User\n" +
			"    value: '12345'\n" +
			"sip_region: Canada\n" +
			"sip_transport_protocol: TLS\n" +
			"sound_modifications:\n" +
			"  octaves: 0.1\n" +
			"  pitch: 0.8\n" +
			"  semitone: -2\n" +
			"  track: both\n" +
			"stream_auth_token: your-auth-token\n" +
			"stream_bidirectional_codec: G722\n" +
			"stream_bidirectional_mode: rtp\n" +
			"stream_bidirectional_sampling_rate: 16000\n" +
			"stream_bidirectional_target_legs: both\n" +
			"stream_codec: PCMA\n" +
			"stream_establish_before_call_originate: true\n" +
			"stream_track: both_tracks\n" +
			"stream_url: wss://www.example.com/websocket\n" +
			"supervise_call_control_id: v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"supervisor_role: barge\n" +
			"time_limit_secs: 60\n" +
			"timeout_secs: 60\n" +
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
			"webhook_url: https://www.example.com/server-b/\n" +
			"webhook_url_method: POST\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"calls", "dial",
		)
	})
}

func TestCallsRetrieveStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "retrieve-status",
			"--call-control-id", "call_control_id",
		)
	})
}

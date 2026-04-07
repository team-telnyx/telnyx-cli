// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIAssistantsVersionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:versions", "retrieve",
			"--assistant-id", "assistant_id",
			"--version-id", "version_id",
			"--include-mcp-servers=true",
		)
	})
}

func TestAIAssistantsVersionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:versions", "update",
			"--assistant-id", "assistant_id",
			"--version-id", "version_id",
			"--description", "description",
			"--dynamic-variables", "{foo: bar}",
			"--dynamic-variables-webhook-url", "dynamic_variables_webhook_url",
			"--enabled-feature", "telephony",
			"--greeting", "greeting",
			"--insight-settings", "{insight_group_id: insight_group_id}",
			"--instructions", "instructions",
			"--llm-api-key-ref", "llm_api_key_ref",
			"--messaging-settings", "{conversation_inactivity_minutes: 1, default_messaging_profile_id: default_messaging_profile_id, delivery_status_webhook_url: delivery_status_webhook_url}",
			"--model", "model",
			"--name", "name",
			"--observability-settings", "{host: host, public_key_ref: public_key_ref, secret_key_ref: secret_key_ref, status: enabled}",
			"--privacy-settings", "{data_retention: true}",
			"--telephony-settings", "{default_texml_app_id: default_texml_app_id, noise_suppression: krisp, noise_suppression_config: {attenuation_limit: 0, mode: advanced}, recording_settings: {channels: single, enabled: true, format: wav}, supports_unauthenticated_web_calls: true, time_limit_secs: 30, user_idle_timeout_secs: 30, voicemail_detection: {on_voicemail_detected: {action: stop_assistant, voicemail_message: {message: message, prompt: prompt, type: prompt}}}}",
			"--tool-id", "string",
			"--tool", "{type: webhook, webhook: {description: description, name: name, url: https://example.com/api/v1/function, async: true, body_parameters: {properties: {age: bar, location: bar}, required: [age, location], type: object}, headers: [{name: name, value: value}], method: GET, path_parameters: {properties: {id: bar}, required: [id], type: object}, query_parameters: {properties: {page: bar}, required: [page], type: object}, store_fields_as_variables: [{name: x, value_path: x}], timeout_ms: 500}}",
			"--transcription", "{language: language, model: deepgram/flux, region: region, settings: {eager_eot_threshold: 0.3, eot_threshold: 0, eot_timeout_ms: 0, numerals: true, smart_format: true}}",
			"--voice-settings", "{voice: voice, api_key_ref: api_key_ref, background_audio: {type: predefined_media, value: silence}, expressive_mode: true, language_boost: auto, similarity_boost: 0, speed: 0, style: 0, temperature: 0, use_speaker_boost: true, voice_speed: 0}",
			"--widget-settings", "{agent_thinking_text: agent_thinking_text, audio_visualizer_config: {color: verdant, preset: preset}, default_state: expanded, give_feedback_url: give_feedback_url, logo_icon_url: logo_icon_url, position: fixed, report_issue_url: report_issue_url, speak_to_interrupt_text: speak_to_interrupt_text, start_call_text: start_call_text, theme: light, view_history_url: view_history_url}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiAssistantsVersionsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:versions", "update",
			"--assistant-id", "assistant_id",
			"--version-id", "version_id",
			"--description", "description",
			"--dynamic-variables", "{foo: bar}",
			"--dynamic-variables-webhook-url", "dynamic_variables_webhook_url",
			"--enabled-feature", "telephony",
			"--greeting", "greeting",
			"--insight-settings.insight-group-id", "insight_group_id",
			"--instructions", "instructions",
			"--llm-api-key-ref", "llm_api_key_ref",
			"--messaging-settings.conversation-inactivity-minutes", "1",
			"--messaging-settings.default-messaging-profile-id", "default_messaging_profile_id",
			"--messaging-settings.delivery-status-webhook-url", "delivery_status_webhook_url",
			"--model", "model",
			"--name", "name",
			"--observability-settings.host", "host",
			"--observability-settings.public-key-ref", "public_key_ref",
			"--observability-settings.secret-key-ref", "secret_key_ref",
			"--observability-settings.status", "enabled",
			"--privacy-settings.data-retention=true",
			"--telephony-settings.default-texml-app-id", "default_texml_app_id",
			"--telephony-settings.noise-suppression", "krisp",
			"--telephony-settings.noise-suppression-config", "{attenuation_limit: 0, mode: advanced}",
			"--telephony-settings.recording-settings", "{channels: single, enabled: true, format: wav}",
			"--telephony-settings.supports-unauthenticated-web-calls=true",
			"--telephony-settings.time-limit-secs", "30",
			"--telephony-settings.user-idle-timeout-secs", "30",
			"--telephony-settings.voicemail-detection", "{on_voicemail_detected: {action: stop_assistant, voicemail_message: {message: message, prompt: prompt, type: prompt}}}",
			"--tool-id", "string",
			"--tool", "{type: webhook, webhook: {description: description, name: name, url: https://example.com/api/v1/function, async: true, body_parameters: {properties: {age: bar, location: bar}, required: [age, location], type: object}, headers: [{name: name, value: value}], method: GET, path_parameters: {properties: {id: bar}, required: [id], type: object}, query_parameters: {properties: {page: bar}, required: [page], type: object}, store_fields_as_variables: [{name: x, value_path: x}], timeout_ms: 500}}",
			"--transcription.language", "language",
			"--transcription.model", "deepgram/flux",
			"--transcription.region", "region",
			"--transcription.settings", "{eager_eot_threshold: 0.3, eot_threshold: 0, eot_timeout_ms: 0, numerals: true, smart_format: true}",
			"--voice-settings.voice", "voice",
			"--voice-settings.api-key-ref", "api_key_ref",
			"--voice-settings.background-audio", "{type: predefined_media, value: silence}",
			"--voice-settings.expressive-mode=true",
			"--voice-settings.language-boost", "auto",
			"--voice-settings.similarity-boost", "0",
			"--voice-settings.speed", "0",
			"--voice-settings.style", "0",
			"--voice-settings.temperature", "0",
			"--voice-settings.use-speaker-boost=true",
			"--voice-settings.voice-speed", "0",
			"--widget-settings.agent-thinking-text", "agent_thinking_text",
			"--widget-settings.audio-visualizer-config", "{color: verdant, preset: preset}",
			"--widget-settings.default-state", "expanded",
			"--widget-settings.give-feedback-url", "give_feedback_url",
			"--widget-settings.logo-icon-url", "logo_icon_url",
			"--widget-settings.position", "fixed",
			"--widget-settings.report-issue-url", "report_issue_url",
			"--widget-settings.speak-to-interrupt-text", "speak_to_interrupt_text",
			"--widget-settings.start-call-text", "start_call_text",
			"--widget-settings.theme", "light",
			"--widget-settings.view-history-url", "view_history_url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"dynamic_variables:\n" +
			"  foo: bar\n" +
			"dynamic_variables_webhook_url: dynamic_variables_webhook_url\n" +
			"enabled_features:\n" +
			"  - telephony\n" +
			"greeting: greeting\n" +
			"insight_settings:\n" +
			"  insight_group_id: insight_group_id\n" +
			"instructions: instructions\n" +
			"llm_api_key_ref: llm_api_key_ref\n" +
			"messaging_settings:\n" +
			"  conversation_inactivity_minutes: 1\n" +
			"  default_messaging_profile_id: default_messaging_profile_id\n" +
			"  delivery_status_webhook_url: delivery_status_webhook_url\n" +
			"model: model\n" +
			"name: name\n" +
			"observability_settings:\n" +
			"  host: host\n" +
			"  public_key_ref: public_key_ref\n" +
			"  secret_key_ref: secret_key_ref\n" +
			"  status: enabled\n" +
			"privacy_settings:\n" +
			"  data_retention: true\n" +
			"telephony_settings:\n" +
			"  default_texml_app_id: default_texml_app_id\n" +
			"  noise_suppression: krisp\n" +
			"  noise_suppression_config:\n" +
			"    attenuation_limit: 0\n" +
			"    mode: advanced\n" +
			"  recording_settings:\n" +
			"    channels: single\n" +
			"    enabled: true\n" +
			"    format: wav\n" +
			"  supports_unauthenticated_web_calls: true\n" +
			"  time_limit_secs: 30\n" +
			"  user_idle_timeout_secs: 30\n" +
			"  voicemail_detection:\n" +
			"    on_voicemail_detected:\n" +
			"      action: stop_assistant\n" +
			"      voicemail_message:\n" +
			"        message: message\n" +
			"        prompt: prompt\n" +
			"        type: prompt\n" +
			"tool_ids:\n" +
			"  - string\n" +
			"tools:\n" +
			"  - type: webhook\n" +
			"    webhook:\n" +
			"      description: description\n" +
			"      name: name\n" +
			"      url: https://example.com/api/v1/function\n" +
			"      async: true\n" +
			"      body_parameters:\n" +
			"        properties:\n" +
			"          age: bar\n" +
			"          location: bar\n" +
			"        required:\n" +
			"          - age\n" +
			"          - location\n" +
			"        type: object\n" +
			"      headers:\n" +
			"        - name: name\n" +
			"          value: value\n" +
			"      method: GET\n" +
			"      path_parameters:\n" +
			"        properties:\n" +
			"          id: bar\n" +
			"        required:\n" +
			"          - id\n" +
			"        type: object\n" +
			"      query_parameters:\n" +
			"        properties:\n" +
			"          page: bar\n" +
			"        required:\n" +
			"          - page\n" +
			"        type: object\n" +
			"      store_fields_as_variables:\n" +
			"        - name: x\n" +
			"          value_path: x\n" +
			"      timeout_ms: 500\n" +
			"transcription:\n" +
			"  language: language\n" +
			"  model: deepgram/flux\n" +
			"  region: region\n" +
			"  settings:\n" +
			"    eager_eot_threshold: 0.3\n" +
			"    eot_threshold: 0\n" +
			"    eot_timeout_ms: 0\n" +
			"    numerals: true\n" +
			"    smart_format: true\n" +
			"voice_settings:\n" +
			"  voice: voice\n" +
			"  api_key_ref: api_key_ref\n" +
			"  background_audio:\n" +
			"    type: predefined_media\n" +
			"    value: silence\n" +
			"  expressive_mode: true\n" +
			"  language_boost: auto\n" +
			"  similarity_boost: 0\n" +
			"  speed: 0\n" +
			"  style: 0\n" +
			"  temperature: 0\n" +
			"  use_speaker_boost: true\n" +
			"  voice_speed: 0\n" +
			"widget_settings:\n" +
			"  agent_thinking_text: agent_thinking_text\n" +
			"  audio_visualizer_config:\n" +
			"    color: verdant\n" +
			"    preset: preset\n" +
			"  default_state: expanded\n" +
			"  give_feedback_url: give_feedback_url\n" +
			"  logo_icon_url: logo_icon_url\n" +
			"  position: fixed\n" +
			"  report_issue_url: report_issue_url\n" +
			"  speak_to_interrupt_text: speak_to_interrupt_text\n" +
			"  start_call_text: start_call_text\n" +
			"  theme: light\n" +
			"  view_history_url: view_history_url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:assistants:versions", "update",
			"--assistant-id", "assistant_id",
			"--version-id", "version_id",
		)
	})
}

func TestAIAssistantsVersionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:versions", "list",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAIAssistantsVersionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:versions", "delete",
			"--assistant-id", "assistant_id",
			"--version-id", "version_id",
		)
	})
}

func TestAIAssistantsVersionsPromote(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:versions", "promote",
			"--assistant-id", "assistant_id",
			"--version-id", "version_id",
		)
	})
}

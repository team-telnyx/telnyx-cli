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
			"--dynamic-variables-webhook-timeout-ms", "1",
			"--dynamic-variables-webhook-url", "dynamic_variables_webhook_url",
			"--enabled-feature", "telephony",
			"--external-llm", "{base_url: base_url, model: model, authentication_method: token, certificate_ref: certificate_ref, forward_metadata: true, llm_api_key_ref: llm_api_key_ref, token_retrieval_url: token_retrieval_url}",
			"--fallback-config", "{external_llm: {base_url: base_url, model: model, authentication_method: token, certificate_ref: certificate_ref, forward_metadata: true, llm_api_key_ref: llm_api_key_ref, token_retrieval_url: token_retrieval_url}, llm_api_key_ref: llm_api_key_ref, model: model}",
			"--greeting", "greeting",
			"--insight-settings", "{insight_group_id: insight_group_id}",
			"--instructions", "instructions",
			"--integration", "{integration_id: integration_id, allowed_list: [string]}",
			"--interruption-settings", "{enable: true, start_speaking_plan: {transcription_endpointing_plan: {on_no_punctuation_seconds: 0, on_number_seconds: 0, on_punctuation_seconds: 0}, wait_seconds: 0}}",
			"--llm-api-key-ref", "llm_api_key_ref",
			"--mcp-server", "{id: id, allowed_tools: [string]}",
			"--messaging-settings", "{conversation_inactivity_minutes: 1, default_messaging_profile_id: default_messaging_profile_id, delivery_status_webhook_url: delivery_status_webhook_url}",
			"--model", "model",
			"--name", "name",
			"--observability-settings", "{host: host, prompt_label: prompt_label, prompt_name: prompt_name, prompt_sync: enabled, prompt_version: 1, public_key_ref: public_key_ref, secret_key_ref: secret_key_ref, status: enabled}",
			"--post-conversation-settings", "{enabled: true}",
			"--privacy-settings", "{data_retention: true}",
			"--tag", "string",
			"--telephony-settings", "{default_texml_app_id: default_texml_app_id, noise_suppression: krisp, noise_suppression_config: {attenuation_limit: 0, mode: advanced}, recording_settings: {channels: single, enabled: true, format: wav}, supports_unauthenticated_web_calls: true, time_limit_secs: 30, user_idle_reply_secs: 0, user_idle_timeout_secs: 10, voicemail_detection: {on_voicemail_detected: {action: stop_assistant, voicemail_message: {message: message, prompt: prompt, type: prompt}}}}",
			"--tool-id", "string",
			"--tool", "{type: webhook, webhook: {description: description, name: name, url: https://example.com/api/v1/function, async: true, body_parameters: {properties: {age: bar, location: bar}, required: [age, location], type: object}, headers: [{name: name, value: value}], method: GET, path_parameters: {properties: {id: bar}, required: [id], type: object}, query_parameters: {properties: {page: bar}, required: [page], type: object}, store_fields_as_variables: [{name: x, value_path: x}], timeout_ms: 500}}",
			"--transcription", "{api_key_ref: api_key_ref, language: language, model: deepgram/flux, region: region, settings: {eager_eot_threshold: 0.3, end_of_turn_confidence_threshold: 0, eot_threshold: 0.5, eot_timeout_ms: 500, keyterm: keyterm, max_turn_silence: 100, min_turn_silence: 100, numerals: true, smart_format: true}}",
			"--version-name", "version_name",
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
			"--dynamic-variables-webhook-timeout-ms", "1",
			"--dynamic-variables-webhook-url", "dynamic_variables_webhook_url",
			"--enabled-feature", "telephony",
			"--external-llm.base-url", "base_url",
			"--external-llm.model", "model",
			"--external-llm.authentication-method", "token",
			"--external-llm.certificate-ref", "certificate_ref",
			"--external-llm.forward-metadata=true",
			"--external-llm.llm-api-key-ref", "llm_api_key_ref",
			"--external-llm.token-retrieval-url", "token_retrieval_url",
			"--fallback-config.external-llm", "{base_url: base_url, model: model, authentication_method: token, certificate_ref: certificate_ref, forward_metadata: true, llm_api_key_ref: llm_api_key_ref, token_retrieval_url: token_retrieval_url}",
			"--fallback-config.llm-api-key-ref", "llm_api_key_ref",
			"--fallback-config.model", "model",
			"--greeting", "greeting",
			"--insight-settings.insight-group-id", "insight_group_id",
			"--instructions", "instructions",
			"--integration.integration-id", "integration_id",
			"--integration.allowed-list", "[string]",
			"--interruption-settings.enable=true",
			"--interruption-settings.start-speaking-plan", "{transcription_endpointing_plan: {on_no_punctuation_seconds: 0, on_number_seconds: 0, on_punctuation_seconds: 0}, wait_seconds: 0}",
			"--llm-api-key-ref", "llm_api_key_ref",
			"--mcp-server.id", "id",
			"--mcp-server.allowed-tools", "[string]",
			"--messaging-settings.conversation-inactivity-minutes", "1",
			"--messaging-settings.default-messaging-profile-id", "default_messaging_profile_id",
			"--messaging-settings.delivery-status-webhook-url", "delivery_status_webhook_url",
			"--model", "model",
			"--name", "name",
			"--observability-settings.host", "host",
			"--observability-settings.prompt-label", "prompt_label",
			"--observability-settings.prompt-name", "prompt_name",
			"--observability-settings.prompt-sync", "enabled",
			"--observability-settings.prompt-version", "1",
			"--observability-settings.public-key-ref", "public_key_ref",
			"--observability-settings.secret-key-ref", "secret_key_ref",
			"--observability-settings.status", "enabled",
			"--post-conversation-settings.enabled=true",
			"--privacy-settings.data-retention=true",
			"--tag", "string",
			"--telephony-settings.default-texml-app-id", "default_texml_app_id",
			"--telephony-settings.noise-suppression", "krisp",
			"--telephony-settings.noise-suppression-config", "{attenuation_limit: 0, mode: advanced}",
			"--telephony-settings.recording-settings", "{channels: single, enabled: true, format: wav}",
			"--telephony-settings.supports-unauthenticated-web-calls=true",
			"--telephony-settings.time-limit-secs", "30",
			"--telephony-settings.user-idle-reply-secs", "0",
			"--telephony-settings.user-idle-timeout-secs", "10",
			"--telephony-settings.voicemail-detection", "{on_voicemail_detected: {action: stop_assistant, voicemail_message: {message: message, prompt: prompt, type: prompt}}}",
			"--tool-id", "string",
			"--tool", "{type: webhook, webhook: {description: description, name: name, url: https://example.com/api/v1/function, async: true, body_parameters: {properties: {age: bar, location: bar}, required: [age, location], type: object}, headers: [{name: name, value: value}], method: GET, path_parameters: {properties: {id: bar}, required: [id], type: object}, query_parameters: {properties: {page: bar}, required: [page], type: object}, store_fields_as_variables: [{name: x, value_path: x}], timeout_ms: 500}}",
			"--transcription.api-key-ref", "api_key_ref",
			"--transcription.language", "language",
			"--transcription.model", "deepgram/flux",
			"--transcription.region", "region",
			"--transcription.settings", "{eager_eot_threshold: 0.3, end_of_turn_confidence_threshold: 0, eot_threshold: 0.5, eot_timeout_ms: 500, keyterm: keyterm, max_turn_silence: 100, min_turn_silence: 100, numerals: true, smart_format: true}",
			"--version-name", "version_name",
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
			"dynamic_variables_webhook_timeout_ms: 1\n" +
			"dynamic_variables_webhook_url: dynamic_variables_webhook_url\n" +
			"enabled_features:\n" +
			"  - telephony\n" +
			"external_llm:\n" +
			"  base_url: base_url\n" +
			"  model: model\n" +
			"  authentication_method: token\n" +
			"  certificate_ref: certificate_ref\n" +
			"  forward_metadata: true\n" +
			"  llm_api_key_ref: llm_api_key_ref\n" +
			"  token_retrieval_url: token_retrieval_url\n" +
			"fallback_config:\n" +
			"  external_llm:\n" +
			"    base_url: base_url\n" +
			"    model: model\n" +
			"    authentication_method: token\n" +
			"    certificate_ref: certificate_ref\n" +
			"    forward_metadata: true\n" +
			"    llm_api_key_ref: llm_api_key_ref\n" +
			"    token_retrieval_url: token_retrieval_url\n" +
			"  llm_api_key_ref: llm_api_key_ref\n" +
			"  model: model\n" +
			"greeting: greeting\n" +
			"insight_settings:\n" +
			"  insight_group_id: insight_group_id\n" +
			"instructions: instructions\n" +
			"integrations:\n" +
			"  - integration_id: integration_id\n" +
			"    allowed_list:\n" +
			"      - string\n" +
			"interruption_settings:\n" +
			"  enable: true\n" +
			"  start_speaking_plan:\n" +
			"    transcription_endpointing_plan:\n" +
			"      on_no_punctuation_seconds: 0\n" +
			"      on_number_seconds: 0\n" +
			"      on_punctuation_seconds: 0\n" +
			"    wait_seconds: 0\n" +
			"llm_api_key_ref: llm_api_key_ref\n" +
			"mcp_servers:\n" +
			"  - id: id\n" +
			"    allowed_tools:\n" +
			"      - string\n" +
			"messaging_settings:\n" +
			"  conversation_inactivity_minutes: 1\n" +
			"  default_messaging_profile_id: default_messaging_profile_id\n" +
			"  delivery_status_webhook_url: delivery_status_webhook_url\n" +
			"model: model\n" +
			"name: name\n" +
			"observability_settings:\n" +
			"  host: host\n" +
			"  prompt_label: prompt_label\n" +
			"  prompt_name: prompt_name\n" +
			"  prompt_sync: enabled\n" +
			"  prompt_version: 1\n" +
			"  public_key_ref: public_key_ref\n" +
			"  secret_key_ref: secret_key_ref\n" +
			"  status: enabled\n" +
			"post_conversation_settings:\n" +
			"  enabled: true\n" +
			"privacy_settings:\n" +
			"  data_retention: true\n" +
			"tags:\n" +
			"  - string\n" +
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
			"  user_idle_reply_secs: 0\n" +
			"  user_idle_timeout_secs: 10\n" +
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
			"  api_key_ref: api_key_ref\n" +
			"  language: language\n" +
			"  model: deepgram/flux\n" +
			"  region: region\n" +
			"  settings:\n" +
			"    eager_eot_threshold: 0.3\n" +
			"    end_of_turn_confidence_threshold: 0\n" +
			"    eot_threshold: 0.5\n" +
			"    eot_timeout_ms: 500\n" +
			"    keyterm: keyterm\n" +
			"    max_turn_silence: 100\n" +
			"    min_turn_silence: 100\n" +
			"    numerals: true\n" +
			"    smart_format: true\n" +
			"version_name: version_name\n" +
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

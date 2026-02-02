// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestAIAssistantsVersionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:versions", "retrieve",
		"--assistant-id", "assistant_id",
		"--version-id", "version_id",
		"--include-mcp-servers=true",
	)
}

func TestAIAssistantsVersionsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
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
		"--privacy-settings", "{data_retention: true}",
		"--telephony-settings", "{default_texml_app_id: default_texml_app_id, noise_suppression: krisp, noise_suppression_config: {attenuation_limit: 0, mode: advanced}, supports_unauthenticated_web_calls: true, time_limit_secs: 30, user_idle_timeout_secs: 30, voicemail_detection: {on_voicemail_detected: {action: stop_assistant, voicemail_message: {message: message, prompt: prompt, type: prompt}}}}",
		"--tool", "{type: webhook, webhook: {description: description, name: name, url: https://example.com/api/v1/function, async: true, body_parameters: {properties: {age: bar, location: bar}, required: [age, location], type: object}, headers: [{name: name, value: value}], method: GET, path_parameters: {properties: {id: bar}, required: [id], type: object}, query_parameters: {properties: {page: bar}, required: [page], type: object}, timeout_ms: 500}}",
		"--transcription", "{language: language, model: deepgram/flux, region: region, settings: {eager_eot_threshold: 0.3, eot_threshold: 0, eot_timeout_ms: 0, numerals: true, smart_format: true}}",
		"--voice-settings", "{voice: voice, api_key_ref: api_key_ref, background_audio: {type: predefined_media, value: silence}, similarity_boost: 0, speed: 0, style: 0, temperature: 0, use_speaker_boost: true, voice_speed: 0}",
		"--widget-settings", "{agent_thinking_text: agent_thinking_text, audio_visualizer_config: {color: verdant, preset: preset}, default_state: expanded, give_feedback_url: give_feedback_url, logo_icon_url: logo_icon_url, position: fixed, report_issue_url: report_issue_url, speak_to_interrupt_text: speak_to_interrupt_text, start_call_text: start_call_text, theme: light, view_history_url: view_history_url}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(aiAssistantsVersionsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
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
		"--privacy-settings.data-retention=true",
		"--telephony-settings.default-texml-app-id", "default_texml_app_id",
		"--telephony-settings.noise-suppression", "krisp",
		"--telephony-settings.noise-suppression-config", "{attenuation_limit: 0, mode: advanced}",
		"--telephony-settings.supports-unauthenticated-web-calls=true",
		"--telephony-settings.time-limit-secs", "30",
		"--telephony-settings.user-idle-timeout-secs", "30",
		"--telephony-settings.voicemail-detection", "{on_voicemail_detected: {action: stop_assistant, voicemail_message: {message: message, prompt: prompt, type: prompt}}}",
		"--tool", "{type: webhook, webhook: {description: description, name: name, url: https://example.com/api/v1/function, async: true, body_parameters: {properties: {age: bar, location: bar}, required: [age, location], type: object}, headers: [{name: name, value: value}], method: GET, path_parameters: {properties: {id: bar}, required: [id], type: object}, query_parameters: {properties: {page: bar}, required: [page], type: object}, timeout_ms: 500}}",
		"--transcription.language", "language",
		"--transcription.model", "deepgram/flux",
		"--transcription.region", "region",
		"--transcription.settings", "{eager_eot_threshold: 0.3, eot_threshold: 0, eot_timeout_ms: 0, numerals: true, smart_format: true}",
		"--voice-settings.voice", "voice",
		"--voice-settings.api-key-ref", "api_key_ref",
		"--voice-settings.background-audio", "{type: predefined_media, value: silence}",
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
}

func TestAIAssistantsVersionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:versions", "list",
		"--assistant-id", "assistant_id",
	)
}

func TestAIAssistantsVersionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:versions", "delete",
		"--assistant-id", "assistant_id",
		"--version-id", "version_id",
	)
}

func TestAIAssistantsVersionsPromote(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:versions", "promote",
		"--assistant-id", "assistant_id",
		"--version-id", "version_id",
	)
}

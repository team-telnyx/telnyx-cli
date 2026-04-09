// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTexmlInitiateAICall(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml", "initiate-ai-call",
			"--connection-id", "connection_id",
			"--ai-assistant-id", "ai-assistant-id-123",
			"--from", "+13120001234",
			"--to", "+13121230000",
			"--ai-assistant-dynamic-variables", "{customer_name: John Doe, account_id: '12345'}",
			"--ai-assistant-version", "AIAssistantVersion",
			"--async-amd=true",
			"--async-amd-status-callback", "https://www.example.com/callback",
			"--async-amd-status-callback-method", "GET",
			"--caller-id", "Info",
			"--conversation-callback", "https://www.example.com/conversation-callback",
			"--conversation-callback-method", "GET",
			"--conversation-callback", "https://www.example.com/conversation-callback1",
			"--conversation-callback", "https://www.example.com/conversation-callback2",
			"--custom-header", "{name: X-Custom-Header, value: custom-value}",
			"--detection-mode", "Premium",
			"--machine-detection", "Enable",
			"--machine-detection-silence-timeout", "2000",
			"--machine-detection-speech-end-threshold", "2000",
			"--machine-detection-speech-threshold", "2000",
			"--machine-detection-timeout", "5000",
			"--passports", "Passports",
			"--preferred-codecs", "PCMA,PCMU",
			"--record=false",
			"--recording-channels", "dual",
			"--recording-status-callback", "https://example.com/recording_status_callback",
			"--recording-status-callback-event", "in-progress completed absent",
			"--recording-status-callback-method", "GET",
			"--recording-timeout", "5",
			"--recording-track", "inbound",
			"--send-recording-url=false",
			"--sip-auth-password", "1234",
			"--sip-auth-username", "user",
			"--sip-region", "Canada",
			"--status-callback", "https://www.example.com/callback",
			"--status-callback-event", "initiated answered",
			"--status-callback-method", "GET",
			"--status-callback", "https://www.example.com/callback1",
			"--status-callback", "https://www.example.com/callback2",
			"--time-limit", "3600",
			"--timeout-seconds", "60",
			"--trim", "trim-silence",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(texmlInitiateAICall)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml", "initiate-ai-call",
			"--connection-id", "connection_id",
			"--ai-assistant-id", "ai-assistant-id-123",
			"--from", "+13120001234",
			"--to", "+13121230000",
			"--ai-assistant-dynamic-variables", "{customer_name: John Doe, account_id: '12345'}",
			"--ai-assistant-version", "AIAssistantVersion",
			"--async-amd=true",
			"--async-amd-status-callback", "https://www.example.com/callback",
			"--async-amd-status-callback-method", "GET",
			"--caller-id", "Info",
			"--conversation-callback", "https://www.example.com/conversation-callback",
			"--conversation-callback-method", "GET",
			"--conversation-callback", "https://www.example.com/conversation-callback1",
			"--conversation-callback", "https://www.example.com/conversation-callback2",
			"--custom-header.name", "X-Custom-Header",
			"--custom-header.value", "custom-value",
			"--detection-mode", "Premium",
			"--machine-detection", "Enable",
			"--machine-detection-silence-timeout", "2000",
			"--machine-detection-speech-end-threshold", "2000",
			"--machine-detection-speech-threshold", "2000",
			"--machine-detection-timeout", "5000",
			"--passports", "Passports",
			"--preferred-codecs", "PCMA,PCMU",
			"--record=false",
			"--recording-channels", "dual",
			"--recording-status-callback", "https://example.com/recording_status_callback",
			"--recording-status-callback-event", "in-progress completed absent",
			"--recording-status-callback-method", "GET",
			"--recording-timeout", "5",
			"--recording-track", "inbound",
			"--send-recording-url=false",
			"--sip-auth-password", "1234",
			"--sip-auth-username", "user",
			"--sip-region", "Canada",
			"--status-callback", "https://www.example.com/callback",
			"--status-callback-event", "initiated answered",
			"--status-callback-method", "GET",
			"--status-callback", "https://www.example.com/callback1",
			"--status-callback", "https://www.example.com/callback2",
			"--time-limit", "3600",
			"--timeout-seconds", "60",
			"--trim", "trim-silence",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"AIAssistantId: ai-assistant-id-123\n" +
			"From: '+13120001234'\n" +
			"To: '+13121230000'\n" +
			"AIAssistantDynamicVariables:\n" +
			"  customer_name: John Doe\n" +
			"  account_id: '12345'\n" +
			"AIAssistantVersion: AIAssistantVersion\n" +
			"AsyncAmd: true\n" +
			"AsyncAmdStatusCallback: https://www.example.com/callback\n" +
			"AsyncAmdStatusCallbackMethod: GET\n" +
			"CallerId: Info\n" +
			"ConversationCallback: https://www.example.com/conversation-callback\n" +
			"ConversationCallbackMethod: GET\n" +
			"ConversationCallbacks:\n" +
			"  - https://www.example.com/conversation-callback1\n" +
			"  - https://www.example.com/conversation-callback2\n" +
			"CustomHeaders:\n" +
			"  - name: X-Custom-Header\n" +
			"    value: custom-value\n" +
			"DetectionMode: Premium\n" +
			"MachineDetection: Enable\n" +
			"MachineDetectionSilenceTimeout: 2000\n" +
			"MachineDetectionSpeechEndThreshold: 2000\n" +
			"MachineDetectionSpeechThreshold: 2000\n" +
			"MachineDetectionTimeout: 5000\n" +
			"Passports: Passports\n" +
			"PreferredCodecs: PCMA,PCMU\n" +
			"Record: false\n" +
			"RecordingChannels: dual\n" +
			"RecordingStatusCallback: https://example.com/recording_status_callback\n" +
			"RecordingStatusCallbackEvent: in-progress completed absent\n" +
			"RecordingStatusCallbackMethod: GET\n" +
			"RecordingTimeout: 5\n" +
			"RecordingTrack: inbound\n" +
			"SendRecordingUrl: false\n" +
			"SipAuthPassword: '1234'\n" +
			"SipAuthUsername: user\n" +
			"SipRegion: Canada\n" +
			"StatusCallback: https://www.example.com/callback\n" +
			"StatusCallbackEvent: initiated answered\n" +
			"StatusCallbackMethod: GET\n" +
			"StatusCallbacks:\n" +
			"  - https://www.example.com/callback1\n" +
			"  - https://www.example.com/callback2\n" +
			"TimeLimit: 3600\n" +
			"Timeout: 60\n" +
			"Trim: trim-silence\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml", "initiate-ai-call",
			"--connection-id", "connection_id",
		)
	})
}

func TestTexmlSecrets(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml", "secrets",
			"--name", "My Secret Name",
			"--value", "My Secret Value",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Secret Name\n" +
			"value: My Secret Value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml", "secrets",
		)
	})
}

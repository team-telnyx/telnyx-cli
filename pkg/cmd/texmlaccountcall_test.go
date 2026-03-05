// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTexmlAccountsCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls", "retrieve",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

func TestTexmlAccountsCallsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls", "update",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
			"--fallback-method", "GET",
			"--fallback-url", "https://www.example.com/intruction-c.xml",
			"--method", "GET",
			"--status", "completed",
			"--status-callback", "https://www.example.com/callback",
			"--status-callback-method", "GET",
			"--texml", `<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello</Say></Response>`,
			"--url", "https://www.example.com/intruction-b.xml",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"FallbackMethod: GET\n" +
			"FallbackUrl: https://www.example.com/intruction-c.xml\n" +
			"Method: GET\n" +
			"Status: completed\n" +
			"StatusCallback: https://www.example.com/callback\n" +
			"StatusCallbackMethod: GET\n" +
			"Texml: <?xml version=\"1.0\" encoding=\"UTF-8\"?><Response><Say>Hello</Say></Response>\n" +
			"Url: https://www.example.com/intruction-b.xml\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml:accounts:calls", "update",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

func TestTexmlAccountsCallsCalls(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls", "calls",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--application-sid", "example-app-sid",
			"--from", "+13120001234",
			"--to", "+13121230000",
			"--async-amd=true",
			"--async-amd-status-callback", "https://www.example.com/callback",
			"--async-amd-status-callback-method", "GET",
			"--caller-id", "Info",
			"--cancel-playback-on-detect-message-end=false",
			"--cancel-playback-on-machine-detection=false",
			"--custom-header", "{name: X-Custom-Header, value: custom-value}",
			"--detection-mode", "Premium",
			"--fallback-url", "https://www.example.com/instructions-fallback.xml",
			"--machine-detection", "Enable",
			"--machine-detection-silence-timeout", "2000",
			"--machine-detection-speech-end-threshold", "2000",
			"--machine-detection-speech-threshold", "2000",
			"--machine-detection-timeout", "5000",
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
			"--status-callback", "https://www.example.com/statuscallback-listener",
			"--status-callback-event", "initiated",
			"--status-callback-method", "GET",
			"--supervise-call-sid", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--supervising-role", "monitor",
			"--texml", `<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello</Say></Response>`,
			"--time-limit", "3600",
			"--timeout-seconds", "60",
			"--trim", "trim-silence",
			"--url", "https://www.example.com/texml.xml",
			"--url-method", "GET",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(texmlAccountsCallsCalls)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls", "calls",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--application-sid", "example-app-sid",
			"--from", "+13120001234",
			"--to", "+13121230000",
			"--async-amd=true",
			"--async-amd-status-callback", "https://www.example.com/callback",
			"--async-amd-status-callback-method", "GET",
			"--caller-id", "Info",
			"--cancel-playback-on-detect-message-end=false",
			"--cancel-playback-on-machine-detection=false",
			"--custom-header.name", "X-Custom-Header",
			"--custom-header.value", "custom-value",
			"--detection-mode", "Premium",
			"--fallback-url", "https://www.example.com/instructions-fallback.xml",
			"--machine-detection", "Enable",
			"--machine-detection-silence-timeout", "2000",
			"--machine-detection-speech-end-threshold", "2000",
			"--machine-detection-speech-threshold", "2000",
			"--machine-detection-timeout", "5000",
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
			"--status-callback", "https://www.example.com/statuscallback-listener",
			"--status-callback-event", "initiated",
			"--status-callback-method", "GET",
			"--supervise-call-sid", "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			"--supervising-role", "monitor",
			"--texml", `<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello</Say></Response>`,
			"--time-limit", "3600",
			"--timeout-seconds", "60",
			"--trim", "trim-silence",
			"--url", "https://www.example.com/texml.xml",
			"--url-method", "GET",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ApplicationSid: example-app-sid\n" +
			"From: '+13120001234'\n" +
			"To: '+13121230000'\n" +
			"AsyncAmd: true\n" +
			"AsyncAmdStatusCallback: https://www.example.com/callback\n" +
			"AsyncAmdStatusCallbackMethod: GET\n" +
			"CallerId: Info\n" +
			"CancelPlaybackOnDetectMessageEnd: false\n" +
			"CancelPlaybackOnMachineDetection: false\n" +
			"CustomHeaders:\n" +
			"  - name: X-Custom-Header\n" +
			"    value: custom-value\n" +
			"DetectionMode: Premium\n" +
			"FallbackUrl: https://www.example.com/instructions-fallback.xml\n" +
			"MachineDetection: Enable\n" +
			"MachineDetectionSilenceTimeout: 2000\n" +
			"MachineDetectionSpeechEndThreshold: 2000\n" +
			"MachineDetectionSpeechThreshold: 2000\n" +
			"MachineDetectionTimeout: 5000\n" +
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
			"StatusCallback: https://www.example.com/statuscallback-listener\n" +
			"StatusCallbackEvent: initiated\n" +
			"StatusCallbackMethod: GET\n" +
			"SuperviseCallSid: v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg\n" +
			"SupervisingRole: monitor\n" +
			"Texml: <?xml version=\"1.0\" encoding=\"UTF-8\"?><Response><Say>Hello</Say></Response>\n" +
			"TimeLimit: 3600\n" +
			"Timeout: 60\n" +
			"Trim: trim-silence\n" +
			"Url: https://www.example.com/texml.xml\n" +
			"UrlMethod: GET\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml:accounts:calls", "calls",
			"--api-key", "string",
			"--account-sid", "account_sid",
		)
	})
}

func TestTexmlAccountsCallsRetrieveCalls(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls", "retrieve-calls",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--end-time", "EndTime",
			"--end-time-gt", "EndTime_gt",
			"--end-time-lt", "EndTime_lt",
			"--from", "From",
			"--page", "0",
			"--page-size", "0",
			"--page-token", "PageToken",
			"--start-time", "StartTime",
			"--start-time-gt", "StartTime_gt",
			"--start-time-lt", "StartTime_lt",
			"--status", "canceled",
			"--to", "To",
		)
	})
}

func TestTexmlAccountsCallsSiprecJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls", "siprec-json",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
			"--connector-name", "my_connector",
			"--include-metadata-custom-headers=true",
			"--name", "my_siprec_session",
			"--secure=true",
			"--session-timeout-secs", "900",
			"--sip-transport", "tcp",
			"--status-callback", "https://www.example.com/callback",
			"--status-callback-method", "GET",
			"--track", "both_tracks",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ConnectorName: my_connector\n" +
			"IncludeMetadataCustomHeaders: true\n" +
			"Name: my_siprec_session\n" +
			"Secure: true\n" +
			"SessionTimeoutSecs: 900\n" +
			"SipTransport: tcp\n" +
			"StatusCallback: https://www.example.com/callback\n" +
			"StatusCallbackMethod: GET\n" +
			"Track: both_tracks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml:accounts:calls", "siprec-json",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

func TestTexmlAccountsCallsStreamsJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:calls", "streams-json",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
			"--bidirectional-codec", "G722",
			"--bidirectional-mode", "rtp",
			"--name", "My stream",
			"--status-callback", "http://webhook.com/callback",
			"--status-callback-method", "GET",
			"--track", "both_tracks",
			"--url", "wss://www.example.com/websocket",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"BidirectionalCodec: G722\n" +
			"BidirectionalMode: rtp\n" +
			"Name: My stream\n" +
			"StatusCallback: http://webhook.com/callback\n" +
			"StatusCallbackMethod: GET\n" +
			"Track: both_tracks\n" +
			"Url: wss://www.example.com/websocket\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml:accounts:calls", "streams-json",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

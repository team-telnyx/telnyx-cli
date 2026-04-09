// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:calls", "retrieve",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

func TestTexmlAccountsCallsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:calls", "update",
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
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:calls", "update",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

func TestTexmlAccountsCallsCalls(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:calls", "calls",
			"--account-sid", "account_sid",
			"--params", "{Url: https://www.example.com/texml.xml, ApplicationSid: example-app-sid, AsyncAmd: true, AsyncAmdStatusCallback: https://www.example.com/callback, AsyncAmdStatusCallbackMethod: GET, CallerId: Info, CancelPlaybackOnDetectMessageEnd: false, CancelPlaybackOnMachineDetection: false, CustomHeaders: [{name: X-Custom-Header, value: custom-value}], DetectionMode: Premium, FallbackUrl: https://www.example.com/instructions-fallback.xml, From: '+13120001234', MachineDetection: Enable, MachineDetectionSilenceTimeout: 2000, MachineDetectionSpeechEndThreshold: 2000, MachineDetectionSpeechThreshold: 2000, MachineDetectionTimeout: 5000, MediaEncryption: disabled, PreferredCodecs: 'PCMA,PCMU', Record: false, RecordingChannels: dual, RecordingStatusCallback: https://example.com/recording_status_callback, RecordingStatusCallbackEvent: in-progress completed absent, RecordingStatusCallbackMethod: GET, RecordingTimeout: 5, RecordingTrack: inbound, SendRecordingUrl: false, SipAuthPassword: '1234', SipAuthUsername: user, SipRegion: Canada, StatusCallback: https://www.example.com/statuscallback-listener, StatusCallbackEvent: initiated, StatusCallbackMethod: GET, SuperviseCallSid: v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg, SupervisingRole: monitor, Texml: Texml, TimeLimit: 3600, Timeout: 60, To: '+13121230000', Trim: trim-silence, UrlMethod: GET}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"Url: https://www.example.com/texml.xml\n" +
			"ApplicationSid: example-app-sid\n" +
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
			"From: '+13120001234'\n" +
			"MachineDetection: Enable\n" +
			"MachineDetectionSilenceTimeout: 2000\n" +
			"MachineDetectionSpeechEndThreshold: 2000\n" +
			"MachineDetectionSpeechThreshold: 2000\n" +
			"MachineDetectionTimeout: 5000\n" +
			"MediaEncryption: disabled\n" +
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
			"Texml: Texml\n" +
			"TimeLimit: 3600\n" +
			"Timeout: 60\n" +
			"To: '+13121230000'\n" +
			"Trim: trim-silence\n" +
			"UrlMethod: GET\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:calls", "calls",
			"--account-sid", "account_sid",
		)
	})
}

func TestTexmlAccountsCallsRetrieveCalls(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:calls", "retrieve-calls",
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
			t,
			"--api-key", "string",
			"texml:accounts:calls", "siprec-json",
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
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:calls", "siprec-json",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

func TestTexmlAccountsCallsStreamsJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:calls", "streams-json",
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
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:calls", "streams-json",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTexmlAccountsCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:calls", "retrieve",
		"--account-sid", "account_sid",
		"--call-sid", "call_sid",
	)
}

func TestTexmlAccountsCallsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

func TestTexmlAccountsCallsCalls(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:calls", "calls",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(texmlAccountsCallsCalls)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:calls", "calls",
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
}

func TestTexmlAccountsCallsRetrieveCalls(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

func TestTexmlAccountsCallsSiprecJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

func TestTexmlAccountsCallsStreamsJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

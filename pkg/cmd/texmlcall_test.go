// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestTexmlCallsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:calls", "update",
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

func TestTexmlCallsInitiate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:calls", "initiate",
		"--application-id", "application_id",
		"--from", "+13120001234",
		"--to", "+13121230000",
		"--async-amd=true",
		"--async-amd-status-callback", "https://www.example.com/callback",
		"--async-amd-status-callback-method", "GET",
		"--caller-id", "Info",
		"--cancel-playback-on-detect-message-end=false",
		"--cancel-playback-on-machine-detection=false",
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
		"--sip-auth-password", "1234",
		"--sip-auth-username", "user",
		"--status-callback", "https://www.example.com/statuscallback-listener",
		"--status-callback-event", "initiated",
		"--status-callback-method", "GET",
		"--trim", "trim-silence",
		"--url", "https://www.example.com/texml.xml",
		"--url-method", "GET",
	)
}

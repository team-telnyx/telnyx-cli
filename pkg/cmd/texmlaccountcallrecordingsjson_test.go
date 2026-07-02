// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsCallsRecordingsJsonRecordingsJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:calls:recordings-json", "recordings-json",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
			"--play-beep=false",
			"--recording-channels", "single",
			"--recording-status-callback", "http://webhook.com/callback",
			"--recording-status-callback-event", "in-progress completed absent",
			"--recording-status-callback-method", "GET",
			"--recording-track", "inbound",
			"--send-recording-url=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"PlayBeep: false\n" +
			"RecordingChannels: single\n" +
			"RecordingStatusCallback: http://webhook.com/callback\n" +
			"RecordingStatusCallbackEvent: in-progress completed absent\n" +
			"RecordingStatusCallbackMethod: GET\n" +
			"RecordingTrack: inbound\n" +
			"SendRecordingUrl: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:calls:recordings-json", "recordings-json",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

func TestTexmlAccountsCallsRecordingsJsonRetrieveRecordingsJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:calls:recordings-json", "retrieve-recordings-json",
			"--account-sid", "account_sid",
			"--call-sid", "call_sid",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTexmlAccountsConferencesParticipantsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences:participants", "retrieve",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
			"--call-sid-or-participant-label", "call_sid_or_participant_label",
		)
	})
}

func TestTexmlAccountsConferencesParticipantsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences:participants", "update",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
			"--call-sid-or-participant-label", "call_sid_or_participant_label",
			"--announce-method", "GET",
			"--announce-url", "https://www.example.com/announce.xml",
			"--beep-on-exit=false",
			"--call-sid-to-coach", "v3:9X2vxPDFY2RHSJ1EdMS0RHRksMTg7ldNxdjWbVr9zBjbGjGsSe-aiQ",
			"--coaching=false",
			"--end-conference-on-exit=false",
			"--hold=true",
			"--hold-method", "POST",
			"--hold-url", "https://www.example.com/hold-music.xml",
			"--muted=true",
			"--wait-url", "https://www.example.com/wait_music.mp3",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"AnnounceMethod: GET\n" +
			"AnnounceUrl: https://www.example.com/announce.xml\n" +
			"BeepOnExit: false\n" +
			"CallSidToCoach: v3:9X2vxPDFY2RHSJ1EdMS0RHRksMTg7ldNxdjWbVr9zBjbGjGsSe-aiQ\n" +
			"Coaching: false\n" +
			"EndConferenceOnExit: false\n" +
			"Hold: true\n" +
			"HoldMethod: POST\n" +
			"HoldUrl: https://www.example.com/hold-music.xml\n" +
			"Muted: true\n" +
			"WaitUrl: https://www.example.com/wait_music.mp3\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:conferences:participants", "update",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
			"--call-sid-or-participant-label", "call_sid_or_participant_label",
		)
	})
}

func TestTexmlAccountsConferencesParticipantsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences:participants", "delete",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
			"--call-sid-or-participant-label", "call_sid_or_participant_label",
		)
	})
}

func TestTexmlAccountsConferencesParticipantsParticipants(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences:participants", "participants",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
			"--amd-status-callback", "https://www.example.com/amd_result",
			"--amd-status-callback-method", "GET",
			"--application-sid", "1846572522338780702",
			"--beep", "onExit",
			"--caller-id", "Info",
			"--call-sid-to-coach", "v3:9X2vxPDFY2RHSJ1EdMS0RHRksMTg7ldNxdjWbVr9zBjbGjGsSe-aiQ",
			"--cancel-playback-on-detect-message-end=false",
			"--cancel-playback-on-machine-detection=false",
			"--coaching=false",
			"--conference-record", "record-from-start",
			"--conference-recording-status-callback", "https://example.com/conference_recording_status_callback",
			"--conference-recording-status-callback-event", "in-progress completed failed absent",
			"--conference-recording-status-callback-method", "GET",
			"--conference-recording-timeout", "5",
			"--conference-status-callback", "https://example.com/conference_status_callback",
			"--conference-status-callback-event", "start end join leave",
			"--conference-status-callback-method", "GET",
			"--conference-trim", "trim-silence",
			"--custom-header", "{name: X-Custom-Header, value: custom-value}",
			"--early-media=true",
			"--end-conference-on-exit=true",
			"--from", "+12065550200",
			"--label", "customer",
			"--machine-detection", "Enable",
			"--machine-detection-silence-timeout", "2000",
			"--machine-detection-speech-end-threshold", "2000",
			"--machine-detection-speech-threshold", "2000",
			"--machine-detection-timeout", "1000",
			"--max-participants", "30",
			"--muted=true",
			"--preferred-codecs", "PCMA,PCMU",
			"--record=false",
			"--recording-channels", "dual",
			"--recording-status-callback", "https://example.com/recording_status_callback",
			"--recording-status-callback-event", "in-progress completed absent",
			"--recording-status-callback-method", "GET",
			"--recording-track", "inbound",
			"--sip-auth-password", "1234",
			"--sip-auth-username", "user",
			"--start-conference-on-enter=false",
			"--status-callback", "https://www.example.com/callback",
			"--status-callback-event", "answered completed",
			"--status-callback-method", "GET",
			"--time-limit", "30",
			"--timeout-seconds", "30",
			"--to", "+12065550100",
			"--trim", "trim-silence",
			"--wait-url", "https://www.example.com/wait_music.mp3",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(texmlAccountsConferencesParticipantsParticipants)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences:participants", "participants",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
			"--amd-status-callback", "https://www.example.com/amd_result",
			"--amd-status-callback-method", "GET",
			"--application-sid", "1846572522338780702",
			"--beep", "onExit",
			"--caller-id", "Info",
			"--call-sid-to-coach", "v3:9X2vxPDFY2RHSJ1EdMS0RHRksMTg7ldNxdjWbVr9zBjbGjGsSe-aiQ",
			"--cancel-playback-on-detect-message-end=false",
			"--cancel-playback-on-machine-detection=false",
			"--coaching=false",
			"--conference-record", "record-from-start",
			"--conference-recording-status-callback", "https://example.com/conference_recording_status_callback",
			"--conference-recording-status-callback-event", "in-progress completed failed absent",
			"--conference-recording-status-callback-method", "GET",
			"--conference-recording-timeout", "5",
			"--conference-status-callback", "https://example.com/conference_status_callback",
			"--conference-status-callback-event", "start end join leave",
			"--conference-status-callback-method", "GET",
			"--conference-trim", "trim-silence",
			"--custom-header.name", "X-Custom-Header",
			"--custom-header.value", "custom-value",
			"--early-media=true",
			"--end-conference-on-exit=true",
			"--from", "+12065550200",
			"--label", "customer",
			"--machine-detection", "Enable",
			"--machine-detection-silence-timeout", "2000",
			"--machine-detection-speech-end-threshold", "2000",
			"--machine-detection-speech-threshold", "2000",
			"--machine-detection-timeout", "1000",
			"--max-participants", "30",
			"--muted=true",
			"--preferred-codecs", "PCMA,PCMU",
			"--record=false",
			"--recording-channels", "dual",
			"--recording-status-callback", "https://example.com/recording_status_callback",
			"--recording-status-callback-event", "in-progress completed absent",
			"--recording-status-callback-method", "GET",
			"--recording-track", "inbound",
			"--sip-auth-password", "1234",
			"--sip-auth-username", "user",
			"--start-conference-on-enter=false",
			"--status-callback", "https://www.example.com/callback",
			"--status-callback-event", "answered completed",
			"--status-callback-method", "GET",
			"--time-limit", "30",
			"--timeout-seconds", "30",
			"--to", "+12065550100",
			"--trim", "trim-silence",
			"--wait-url", "https://www.example.com/wait_music.mp3",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"AmdStatusCallback: https://www.example.com/amd_result\n" +
			"AmdStatusCallbackMethod: GET\n" +
			"ApplicationSid: '1846572522338780702'\n" +
			"Beep: onExit\n" +
			"CallerId: Info\n" +
			"CallSidToCoach: v3:9X2vxPDFY2RHSJ1EdMS0RHRksMTg7ldNxdjWbVr9zBjbGjGsSe-aiQ\n" +
			"CancelPlaybackOnDetectMessageEnd: false\n" +
			"CancelPlaybackOnMachineDetection: false\n" +
			"Coaching: false\n" +
			"ConferenceRecord: record-from-start\n" +
			"ConferenceRecordingStatusCallback: https://example.com/conference_recording_status_callback\n" +
			"ConferenceRecordingStatusCallbackEvent: in-progress completed failed absent\n" +
			"ConferenceRecordingStatusCallbackMethod: GET\n" +
			"ConferenceRecordingTimeout: 5\n" +
			"ConferenceStatusCallback: https://example.com/conference_status_callback\n" +
			"ConferenceStatusCallbackEvent: start end join leave\n" +
			"ConferenceStatusCallbackMethod: GET\n" +
			"ConferenceTrim: trim-silence\n" +
			"CustomHeaders:\n" +
			"  - name: X-Custom-Header\n" +
			"    value: custom-value\n" +
			"EarlyMedia: true\n" +
			"EndConferenceOnExit: true\n" +
			"From: '+12065550200'\n" +
			"Label: customer\n" +
			"MachineDetection: Enable\n" +
			"MachineDetectionSilenceTimeout: 2000\n" +
			"MachineDetectionSpeechEndThreshold: 2000\n" +
			"MachineDetectionSpeechThreshold: 2000\n" +
			"MachineDetectionTimeout: 1000\n" +
			"MaxParticipants: 30\n" +
			"Muted: true\n" +
			"PreferredCodecs: PCMA,PCMU\n" +
			"Record: false\n" +
			"RecordingChannels: dual\n" +
			"RecordingStatusCallback: https://example.com/recording_status_callback\n" +
			"RecordingStatusCallbackEvent: in-progress completed absent\n" +
			"RecordingStatusCallbackMethod: GET\n" +
			"RecordingTrack: inbound\n" +
			"SipAuthPassword: '1234'\n" +
			"SipAuthUsername: user\n" +
			"StartConferenceOnEnter: false\n" +
			"StatusCallback: https://www.example.com/callback\n" +
			"StatusCallbackEvent: answered completed\n" +
			"StatusCallbackMethod: GET\n" +
			"TimeLimit: 30\n" +
			"Timeout: 30\n" +
			"To: '+12065550100'\n" +
			"Trim: trim-silence\n" +
			"WaitUrl: https://www.example.com/wait_music.mp3\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"texml:accounts:conferences:participants", "participants",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
		)
	})
}

func TestTexmlAccountsConferencesParticipantsRetrieveParticipants(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"texml:accounts:conferences:participants", "retrieve-participants",
			"--account-sid", "account_sid",
			"--conference-sid", "conference_sid",
		)
	})
}

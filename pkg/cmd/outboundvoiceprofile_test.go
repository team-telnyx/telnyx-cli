// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestOutboundVoiceProfilesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"outbound-voice-profiles", "create",
		"--name", "office",
		"--billing-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--call-recording", "{call_recording_caller_phone_numbers: ['+19705555098'], call_recording_channels: dual, call_recording_format: mp3, call_recording_type: by_caller_phone_number}",
		"--calling-window", "{calls_per_cld: 5, end_time: 20:00:00.00Z, start_time: 08:00:00.00Z}",
		"--concurrent-call-limit", "10",
		"--daily-spend-limit", "100.00",
		"--daily-spend-limit-enabled=true",
		"--enabled=true",
		"--max-destination-rate", "10",
		"--service-plan", "global",
		"--tag", "office-profile",
		"--traffic-type", "conversational",
		"--usage-payment-method", "rate-deck",
		"--whitelisted-destination", "US",
		"--whitelisted-destination", "BR",
		"--whitelisted-destination", "AU",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(outboundVoiceProfilesCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"outbound-voice-profiles", "create",
		"--name", "office",
		"--billing-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--call-recording.call-recording-caller-phone-numbers", "['+19705555098']",
		"--call-recording.call-recording-channels", "dual",
		"--call-recording.call-recording-format", "mp3",
		"--call-recording.call-recording-type", "by_caller_phone_number",
		"--calling-window.calls-per-cld", "5",
		"--calling-window.end-time", "20:00:00.00Z",
		"--calling-window.start-time", "08:00:00.00Z",
		"--concurrent-call-limit", "10",
		"--daily-spend-limit", "100.00",
		"--daily-spend-limit-enabled=true",
		"--enabled=true",
		"--max-destination-rate", "10",
		"--service-plan", "global",
		"--tag", "office-profile",
		"--traffic-type", "conversational",
		"--usage-payment-method", "rate-deck",
		"--whitelisted-destination", "US",
		"--whitelisted-destination", "BR",
		"--whitelisted-destination", "AU",
	)
}

func TestOutboundVoiceProfilesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"outbound-voice-profiles", "retrieve",
		"--id", "1293384261075731499",
	)
}

func TestOutboundVoiceProfilesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"outbound-voice-profiles", "update",
		"--id", "1293384261075731499",
		"--name", "office",
		"--billing-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--call-recording", "{call_recording_caller_phone_numbers: ['+19705555098'], call_recording_channels: dual, call_recording_format: mp3, call_recording_type: by_caller_phone_number}",
		"--calling-window", "{calls_per_cld: 5, end_time: 20:00:00.00Z, start_time: 08:00:00.00Z}",
		"--concurrent-call-limit", "10",
		"--daily-spend-limit", "100.00",
		"--daily-spend-limit-enabled=true",
		"--enabled=true",
		"--max-destination-rate", "10",
		"--service-plan", "global",
		"--tag", "office-profile",
		"--traffic-type", "conversational",
		"--usage-payment-method", "rate-deck",
		"--whitelisted-destination", "US",
		"--whitelisted-destination", "BR",
		"--whitelisted-destination", "AU",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(outboundVoiceProfilesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"outbound-voice-profiles", "update",
		"--id", "1293384261075731499",
		"--name", "office",
		"--billing-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--call-recording.call-recording-caller-phone-numbers", "['+19705555098']",
		"--call-recording.call-recording-channels", "dual",
		"--call-recording.call-recording-format", "mp3",
		"--call-recording.call-recording-type", "by_caller_phone_number",
		"--calling-window.calls-per-cld", "5",
		"--calling-window.end-time", "20:00:00.00Z",
		"--calling-window.start-time", "08:00:00.00Z",
		"--concurrent-call-limit", "10",
		"--daily-spend-limit", "100.00",
		"--daily-spend-limit-enabled=true",
		"--enabled=true",
		"--max-destination-rate", "10",
		"--service-plan", "global",
		"--tag", "office-profile",
		"--traffic-type", "conversational",
		"--usage-payment-method", "rate-deck",
		"--whitelisted-destination", "US",
		"--whitelisted-destination", "BR",
		"--whitelisted-destination", "AU",
	)
}

func TestOutboundVoiceProfilesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"outbound-voice-profiles", "list",
		"--filter", "{name: {contains: office-profile}}",
		"--page", "{number: 1, size: 1}",
		"--sort", "name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(outboundVoiceProfilesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"outbound-voice-profiles", "list",
		"--filter.name", "{contains: office-profile}",
		"--page.number", "1",
		"--page.size", "1",
		"--sort", "name",
	)
}

func TestOutboundVoiceProfilesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"outbound-voice-profiles", "delete",
		"--id", "1293384261075731499",
	)
}

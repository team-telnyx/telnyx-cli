// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestOutboundVoiceProfilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(outboundVoiceProfilesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: office\n" +
			"billing_group_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"call_recording:\n" +
			"  call_recording_caller_phone_numbers:\n" +
			"    - '+19705555098'\n" +
			"  call_recording_channels: dual\n" +
			"  call_recording_format: mp3\n" +
			"  call_recording_type: by_caller_phone_number\n" +
			"calling_window:\n" +
			"  calls_per_cld: 5\n" +
			"  end_time: 20:00:00.00Z\n" +
			"  start_time: 08:00:00.00Z\n" +
			"concurrent_call_limit: 10\n" +
			"daily_spend_limit: '100.00'\n" +
			"daily_spend_limit_enabled: true\n" +
			"enabled: true\n" +
			"max_destination_rate: 10\n" +
			"service_plan: global\n" +
			"tags:\n" +
			"  - office-profile\n" +
			"traffic_type: conversational\n" +
			"usage_payment_method: rate-deck\n" +
			"whitelisted_destinations:\n" +
			"  - US\n" +
			"  - BR\n" +
			"  - AU\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"outbound-voice-profiles", "create",
		)
	})
}

func TestOutboundVoiceProfilesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"outbound-voice-profiles", "retrieve",
			"--id", "1293384261075731499",
		)
	})
}

func TestOutboundVoiceProfilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(outboundVoiceProfilesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: office\n" +
			"billing_group_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"call_recording:\n" +
			"  call_recording_caller_phone_numbers:\n" +
			"    - '+19705555098'\n" +
			"  call_recording_channels: dual\n" +
			"  call_recording_format: mp3\n" +
			"  call_recording_type: by_caller_phone_number\n" +
			"calling_window:\n" +
			"  calls_per_cld: 5\n" +
			"  end_time: 20:00:00.00Z\n" +
			"  start_time: 08:00:00.00Z\n" +
			"concurrent_call_limit: 10\n" +
			"daily_spend_limit: '100.00'\n" +
			"daily_spend_limit_enabled: true\n" +
			"enabled: true\n" +
			"max_destination_rate: 10\n" +
			"service_plan: global\n" +
			"tags:\n" +
			"  - office-profile\n" +
			"traffic_type: conversational\n" +
			"usage_payment_method: rate-deck\n" +
			"whitelisted_destinations:\n" +
			"  - US\n" +
			"  - BR\n" +
			"  - AU\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"outbound-voice-profiles", "update",
			"--id", "1293384261075731499",
		)
	})
}

func TestOutboundVoiceProfilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"outbound-voice-profiles", "list",
			"--max-items", "10",
			"--filter", "{name: {contains: office-profile}}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(outboundVoiceProfilesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"outbound-voice-profiles", "list",
			"--max-items", "10",
			"--filter.name", "{contains: office-profile}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "name",
		)
	})
}

func TestOutboundVoiceProfilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"outbound-voice-profiles", "delete",
			"--id", "1293384261075731499",
		)
	})
}

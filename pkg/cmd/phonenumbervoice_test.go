// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPhoneNumbersVoiceRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:voice", "retrieve",
		"--id", "1293384261075731499",
	)
}

func TestPhoneNumbersVoiceUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:voice", "update",
		"--id", "1293384261075731499",
		"--call-forwarding", "{call_forwarding_enabled: true, forwarding_type: always, forwards_to: '+13035559123'}",
		"--call-recording", "{inbound_call_recording_channels: single, inbound_call_recording_enabled: true, inbound_call_recording_format: wav}",
		"--caller-id-name-enabled=true",
		"--cnam-listing", "{cnam_listing_details: example, cnam_listing_enabled: true}",
		"--inbound-call-screening", "disabled",
		"--media-features", "{accept_any_rtp_packets_enabled: true, rtp_auto_adjust_enabled: true, t38_fax_gateway_enabled: true}",
		"--tech-prefix-enabled=true",
		"--translated-number", "+13035559999",
		"--usage-payment-method", "pay-per-minute",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumbersVoiceUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:voice", "update",
		"--id", "1293384261075731499",
		"--call-forwarding.call-forwarding-enabled=true",
		"--call-forwarding.forwarding-type", "always",
		"--call-forwarding.forwards-to", "+13035559123",
		"--call-recording.inbound-call-recording-channels", "single",
		"--call-recording.inbound-call-recording-enabled=true",
		"--call-recording.inbound-call-recording-format", "wav",
		"--caller-id-name-enabled=true",
		"--cnam-listing.cnam-listing-details", "example",
		"--cnam-listing.cnam-listing-enabled=true",
		"--inbound-call-screening", "disabled",
		"--media-features.accept-any-rtp-packets-enabled=true",
		"--media-features.rtp-auto-adjust-enabled=true",
		"--media-features.t38-fax-gateway-enabled=true",
		"--tech-prefix-enabled=true",
		"--translated-number", "+13035559999",
		"--usage-payment-method", "pay-per-minute",
	)
}

func TestPhoneNumbersVoiceList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:voice", "list",
		"--filter", "{connection_name: {contains: test}, customer_reference: customer_reference, phone_number: phone_number, voice.usage_payment_method: channel}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "connection_name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumbersVoiceList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:voice", "list",
		"--filter.connection-name", "{contains: test}",
		"--filter.customer-reference", "customer_reference",
		"--filter.phone-number", "phone_number",
		"--filter.voice-usage-payment-method", "channel",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "connection_name",
	)
}

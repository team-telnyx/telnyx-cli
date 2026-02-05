// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPhoneNumbersJobsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:jobs", "retrieve",
		"--id", "id",
	)
}

func TestPhoneNumbersJobsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:jobs", "list",
		"--filter", "{type: update_emergency_settings}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "created_at",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumbersJobsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:jobs", "list",
		"--filter.type", "update_emergency_settings",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "created_at",
	)
}

func TestPhoneNumbersJobsDeleteBatch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:jobs", "delete-batch",
		"--phone-number", "+19705555098",
		"--phone-number", "+19715555098",
		"--phone-number", "32873127836",
	)
}

func TestPhoneNumbersJobsUpdateBatch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:jobs", "update-batch",
		"--phone-number", "1583466971586889004",
		"--phone-number", "+13127367254",
		"--filter", "{billing_group_id: 62e4bf2e-c278-4282-b524-488d9c9c43b2, connection_id: '1521916448077776306', customer_reference: customer_reference, emergency_address_id: '9102160989215728032', has_bundle: has_bundle, phone_number: phone_number, status: active, tag: tag, voice.connection_name: {contains: test, ends_with: test, eq: test, starts_with: test}, voice.usage_payment_method: channel}",
		"--billing-group-id", "dc8e4d67-33a0-4cbb-af74-7b58f05bd494",
		"--connection-id", "dc8e4d67-33a0-4cbb-af74-7b58f05bd494",
		"--customer-reference", "customer-reference",
		"--deletion-lock-enabled=true",
		"--external-pin", "123456",
		"--hd-voice-enabled=true",
		"--tag", "tag",
		"--voice", "{call_forwarding: {call_forwarding_enabled: true, forwarding_type: always, forwards_to: '+13035559123'}, call_recording: {inbound_call_recording_channels: single, inbound_call_recording_enabled: true, inbound_call_recording_format: wav}, caller_id_name_enabled: true, cnam_listing: {cnam_listing_details: example, cnam_listing_enabled: true}, inbound_call_screening: disabled, media_features: {accept_any_rtp_packets_enabled: true, rtp_auto_adjust_enabled: true, t38_fax_gateway_enabled: true}, tech_prefix_enabled: true, translated_number: '+13035559999', usage_payment_method: pay-per-minute}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumbersJobsUpdateBatch)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:jobs", "update-batch",
		"--phone-number", "1583466971586889004",
		"--phone-number", "+13127367254",
		"--filter.billing-group-id", "62e4bf2e-c278-4282-b524-488d9c9c43b2",
		"--filter.connection-id", "1521916448077776306",
		"--filter.customer-reference", "customer_reference",
		"--filter.emergency-address-id", "9102160989215728032",
		"--filter.has-bundle", "has_bundle",
		"--filter.phone-number", "phone_number",
		"--filter.status", "active",
		"--filter.tag", "tag",
		"--filter.voice-connection-name", "{contains: test, ends_with: test, eq: test, starts_with: test}",
		"--filter.voice-usage-payment-method", "channel",
		"--billing-group-id", "dc8e4d67-33a0-4cbb-af74-7b58f05bd494",
		"--connection-id", "dc8e4d67-33a0-4cbb-af74-7b58f05bd494",
		"--customer-reference", "customer-reference",
		"--deletion-lock-enabled=true",
		"--external-pin", "123456",
		"--hd-voice-enabled=true",
		"--tag", "tag",
		"--voice.call-forwarding", "{call_forwarding_enabled: true, forwarding_type: always, forwards_to: '+13035559123'}",
		"--voice.call-recording", "{inbound_call_recording_channels: single, inbound_call_recording_enabled: true, inbound_call_recording_format: wav}",
		"--voice.caller-id-name-enabled=true",
		"--voice.cnam-listing", "{cnam_listing_details: example, cnam_listing_enabled: true}",
		"--voice.inbound-call-screening", "disabled",
		"--voice.media-features", "{accept_any_rtp_packets_enabled: true, rtp_auto_adjust_enabled: true, t38_fax_gateway_enabled: true}",
		"--voice.tech-prefix-enabled=true",
		"--voice.translated-number", "+13035559999",
		"--voice.usage-payment-method", "pay-per-minute",
	)
}

func TestPhoneNumbersJobsUpdateEmergencySettingsBatch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:jobs", "update-emergency-settings-batch",
		"--emergency-enabled=true",
		"--phone-number", "+19705555098",
		"--phone-number", "+19715555098",
		"--phone-number", "32873127836",
		"--emergency-address-id", "53829456729313",
	)
}

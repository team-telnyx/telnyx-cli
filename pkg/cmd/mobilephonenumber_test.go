// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMobilePhoneNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-phone-numbers", "retrieve",
			"--id", "id",
		)
	})
}

func TestMobilePhoneNumbersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-phone-numbers", "update",
			"--id", "id",
			"--call-forwarding", "{call_forwarding_enabled: true, forwarding_type: always, forwards_to: forwards_to}",
			"--call-recording", "{inbound_call_recording_channels: single, inbound_call_recording_enabled: true, inbound_call_recording_format: wav}",
			"--caller-id-name-enabled=true",
			"--cnam-listing", "{cnam_listing_details: cnam_listing_details, cnam_listing_enabled: true}",
			"--connection-id", "connection_id",
			"--customer-reference", "customer_reference",
			"--inbound", "{interception_app_id: interception_app_id}",
			"--inbound-call-screening", "disabled",
			"--noise-suppression=true",
			"--outbound", "{interception_app_id: interception_app_id}",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mobilePhoneNumbersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-phone-numbers", "update",
			"--id", "id",
			"--call-forwarding.call-forwarding-enabled=true",
			"--call-forwarding.forwarding-type", "always",
			"--call-forwarding.forwards-to", "forwards_to",
			"--call-recording.inbound-call-recording-channels", "single",
			"--call-recording.inbound-call-recording-enabled=true",
			"--call-recording.inbound-call-recording-format", "wav",
			"--caller-id-name-enabled=true",
			"--cnam-listing.cnam-listing-details", "cnam_listing_details",
			"--cnam-listing.cnam-listing-enabled=true",
			"--connection-id", "connection_id",
			"--customer-reference", "customer_reference",
			"--inbound.interception-app-id", "interception_app_id",
			"--inbound-call-screening", "disabled",
			"--noise-suppression=true",
			"--outbound.interception-app-id", "interception_app_id",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_forwarding:\n" +
			"  call_forwarding_enabled: true\n" +
			"  forwarding_type: always\n" +
			"  forwards_to: forwards_to\n" +
			"call_recording:\n" +
			"  inbound_call_recording_channels: single\n" +
			"  inbound_call_recording_enabled: true\n" +
			"  inbound_call_recording_format: wav\n" +
			"caller_id_name_enabled: true\n" +
			"cnam_listing:\n" +
			"  cnam_listing_details: cnam_listing_details\n" +
			"  cnam_listing_enabled: true\n" +
			"connection_id: connection_id\n" +
			"customer_reference: customer_reference\n" +
			"inbound:\n" +
			"  interception_app_id: interception_app_id\n" +
			"inbound_call_screening: disabled\n" +
			"noise_suppression: true\n" +
			"outbound:\n" +
			"  interception_app_id: interception_app_id\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"mobile-phone-numbers", "update",
			"--id", "id",
		)
	})
}

func TestMobilePhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-phone-numbers", "list",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

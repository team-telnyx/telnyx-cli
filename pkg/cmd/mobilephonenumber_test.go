// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMobilePhoneNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-phone-numbers", "retrieve",
		"--id", "id",
	)
}

func TestMobilePhoneNumbersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(mobilePhoneNumbersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

func TestMobilePhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-phone-numbers", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

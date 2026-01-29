// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestPhoneNumbersCsvDownloadsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:csv-downloads", "create",
		"--csv-format", "V2",
		"--filter", "{billing_group_id: 62e4bf2e-c278-4282-b524-488d9c9c43b2, connection_id: '1521916448077776306', customer_reference: customer_reference, emergency_address_id: '9102160989215728032', has_bundle: has_bundle, phone_number: phone_number, status: active, tag: tag, voice.connection_name: {contains: test, ends_with: test, eq: test, starts_with: test}, voice.usage_payment_method: channel}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumbersCsvDownloadsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:csv-downloads", "create",
		"--csv-format", "V2",
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
	)
}

func TestPhoneNumbersCsvDownloadsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:csv-downloads", "retrieve",
		"--id", "id",
	)
}

func TestPhoneNumbersCsvDownloadsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:csv-downloads", "list",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumbersCsvDownloadsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:csv-downloads", "list",
		"--page.number", "1",
		"--page.size", "1",
	)
}

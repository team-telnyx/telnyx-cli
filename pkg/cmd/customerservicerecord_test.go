// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestCustomerServiceRecordsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"customer-service-records", "create",
		"--phone-number", "+13035553000",
		"--additional-data", "{account_number: '123456789', address_line_1: 123 Main St, authorized_person_name: John Doe, billing_phone_number: '+12065551212', city: New York, customer_code: '123456789', name: Entity Inc., pin: '1234', state: NY, zip_code: '10001'}",
		"--webhook-url", "https://example.com/webhook",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(customerServiceRecordsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"customer-service-records", "create",
		"--phone-number", "+13035553000",
		"--additional-data.account-number", "123456789",
		"--additional-data.address-line-1", "123 Main St",
		"--additional-data.authorized-person-name", "John Doe",
		"--additional-data.billing-phone-number", "+12065551212",
		"--additional-data.city", "New York",
		"--additional-data.customer-code", "123456789",
		"--additional-data.name", "Entity Inc.",
		"--additional-data.pin", "1234",
		"--additional-data.state", "NY",
		"--additional-data.zip-code", "10001",
		"--webhook-url", "https://example.com/webhook",
	)
}

func TestCustomerServiceRecordsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"customer-service-records", "retrieve",
		"--customer-service-record-id", "customer_service_record_id",
	)
}

func TestCustomerServiceRecordsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"customer-service-records", "list",
		"--filter", "{created_at: {gt: '2020-01-01T00:00:00Z', lt: '2020-01-01T00:00:00Z'}, phone_number: {eq: '+12441239999', in: ['+12441239999']}, status: {eq: pending, in: [pending]}}",
		"--page", "{number: 1, size: 1}",
		"--sort", "{value: created_at}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(customerServiceRecordsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"customer-service-records", "list",
		"--filter.created-at", "{gt: '2020-01-01T00:00:00Z', lt: '2020-01-01T00:00:00Z'}",
		"--filter.phone-number", "{eq: '+12441239999', in: ['+12441239999']}",
		"--filter.status", "{eq: pending, in: [pending]}",
		"--page.number", "1",
		"--page.size", "1",
		"--sort.value", "created_at",
	)
}

func TestCustomerServiceRecordsVerifyPhoneNumberCoverage(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"customer-service-records", "verify-phone-number-coverage",
		"--phone-number", "+13035553000",
	)
}

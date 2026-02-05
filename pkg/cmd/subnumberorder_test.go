// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSubNumberOrdersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders", "retrieve",
		"--sub-number-order-id", "sub_number_order_id",
		"--filter", "{include_phone_numbers: true}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(subNumberOrdersRetrieve)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders", "retrieve",
		"--sub-number-order-id", "sub_number_order_id",
		"--filter.include-phone-numbers=true",
	)
}

func TestSubNumberOrdersUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders", "update",
		"--sub-number-order-id", "sub_number_order_id",
		"--regulatory-requirement", "{field_value: 45f45a04-b4be-4592-95b1-9306b9db2b21, requirement_id: 8ffb3622-7c6b-4ccc-b65f-7a3dc0099576}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(subNumberOrdersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders", "update",
		"--sub-number-order-id", "sub_number_order_id",
		"--regulatory-requirement.field-value", "45f45a04-b4be-4592-95b1-9306b9db2b21",
		"--regulatory-requirement.requirement-id", "8ffb3622-7c6b-4ccc-b65f-7a3dc0099576",
	)
}

func TestSubNumberOrdersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders", "list",
		"--filter", "{country_code: US, order_request_id: 12ade33a-21c0-473b-b055-b3c836e1c293, phone_number_type: local, phone_numbers_count: 1, status: status}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(subNumberOrdersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders", "list",
		"--filter.country-code", "US",
		"--filter.order-request-id", "12ade33a-21c0-473b-b055-b3c836e1c293",
		"--filter.phone-number-type", "local",
		"--filter.phone-numbers-count", "1",
		"--filter.status", "status",
	)
}

func TestSubNumberOrdersCancel(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders", "cancel",
		"--sub-number-order-id", "sub_number_order_id",
	)
}

func TestSubNumberOrdersUpdateRequirementGroup(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sub-number-orders", "update-requirement-group",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--requirement-group-id", "a4b201f9-8646-4e54-a7d2-b2e403eeaf8c",
	)
}

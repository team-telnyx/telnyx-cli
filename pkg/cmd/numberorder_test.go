// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNumberOrdersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-orders", "create",
		"--billing-group-id", "abc85f64-5717-4562-b3fc-2c9600",
		"--connection-id", "346789098765567",
		"--customer-reference", "MY REF 001",
		"--messaging-profile-id", "abc85f64-5717-4562-b3fc-2c9600",
		"--phone-number", "{phone_number: '+19705555098', bundle_id: bc8e4d67-33a0-4cbb-af74-7b58f05bd494, requirement_group_id: dbbd94ee-9079-488f-80ba-f566b247fd7}",
		"--phone-number", "{phone_number: '+492111609539', bundle_id: bc8e4d67-33a0-4cbb-af74-7b58f05bd494, requirement_group_id: dbbd94ee-9079-488f-80ba-f566b247fd79}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(numberOrdersCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-orders", "create",
		"--billing-group-id", "abc85f64-5717-4562-b3fc-2c9600",
		"--connection-id", "346789098765567",
		"--customer-reference", "MY REF 001",
		"--messaging-profile-id", "abc85f64-5717-4562-b3fc-2c9600",
		"--phone-number.phone-number", "+19705555098",
		"--phone-number.bundle-id", "bc8e4d67-33a0-4cbb-af74-7b58f05bd494",
		"--phone-number.requirement-group-id", "dbbd94ee-9079-488f-80ba-f566b247fd7",
		"--phone-number.phone-number", "+492111609539",
		"--phone-number.bundle-id", "bc8e4d67-33a0-4cbb-af74-7b58f05bd494",
		"--phone-number.requirement-group-id", "dbbd94ee-9079-488f-80ba-f566b247fd79",
	)
}

func TestNumberOrdersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-orders", "retrieve",
		"--number-order-id", "number_order_id",
	)
}

func TestNumberOrdersUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-orders", "update",
		"--number-order-id", "number_order_id",
		"--customer-reference", "MY REF 001",
		"--regulatory-requirement", "{field_value: 45f45a04-b4be-4592-95b1-9306b9db2b21, requirement_id: 8ffb3622-7c6b-4ccc-b65f-7a3dc0099576}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(numberOrdersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-orders", "update",
		"--number-order-id", "number_order_id",
		"--customer-reference", "MY REF 001",
		"--regulatory-requirement.field-value", "45f45a04-b4be-4592-95b1-9306b9db2b21",
		"--regulatory-requirement.requirement-id", "8ffb3622-7c6b-4ccc-b65f-7a3dc0099576",
	)
}

func TestNumberOrdersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-orders", "list",
		"--filter", "{created_at: {gt: gt, lt: lt}, customer_reference: customer_reference, phone_numbers_count: phone_numbers_count, requirements_met: true, status: status}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(numberOrdersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-orders", "list",
		"--filter.created-at", "{gt: gt, lt: lt}",
		"--filter.customer-reference", "customer_reference",
		"--filter.phone-numbers-count", "phone_numbers_count",
		"--filter.requirements-met=true",
		"--filter.status", "status",
		"--page-number", "0",
		"--page-size", "0",
	)
}

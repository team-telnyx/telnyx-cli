// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNumberOrderPhoneNumbersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-order-phone-numbers", "retrieve",
		"--number-order-phone-number-id", "number_order_phone_number_id",
	)
}

func TestNumberOrderPhoneNumbersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-order-phone-numbers", "list",
		"--filter", "{country_code: US}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(numberOrderPhoneNumbersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-order-phone-numbers", "list",
		"--filter.country-code", "US",
	)
}

func TestNumberOrderPhoneNumbersUpdateRequirementGroup(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-order-phone-numbers", "update-requirement-group",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--requirement-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestNumberOrderPhoneNumbersUpdateRequirements(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-order-phone-numbers", "update-requirements",
		"--number-order-phone-number-id", "number_order_phone_number_id",
		"--regulatory-requirement", "{field_value: 45f45a04-b4be-4592-95b1-9306b9db2b21, requirement_id: 8ffb3622-7c6b-4ccc-b65f-7a3dc0099576}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(numberOrderPhoneNumbersUpdateRequirements)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"number-order-phone-numbers", "update-requirements",
		"--number-order-phone-number-id", "number_order_phone_number_id",
		"--regulatory-requirement.field-value", "45f45a04-b4be-4592-95b1-9306b9db2b21",
		"--regulatory-requirement.requirement-id", "8ffb3622-7c6b-4ccc-b65f-7a3dc0099576",
	)
}

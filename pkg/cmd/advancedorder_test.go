// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAdvancedOrdersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"advanced-orders", "create",
		"--area-code", "xxx",
		"--comments", "comments",
		"--country-code", "xx",
		"--customer-reference", "customer_reference",
		"--feature", "sms",
		"--phone-number-type", "local",
		"--quantity", "1",
		"--requirement-group-id", "3fa85f64-5717-4562-b3fc-2c963f66afa6",
	)
}

func TestAdvancedOrdersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"advanced-orders", "retrieve",
		"--order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAdvancedOrdersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"advanced-orders", "list",
	)
}

func TestAdvancedOrdersUpdateRequirementGroup(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"advanced-orders", "update-requirement-group",
		"--advanced-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--area-code", "xxx",
		"--comments", "comments",
		"--country-code", "xx",
		"--customer-reference", "customer_reference",
		"--feature", "sms",
		"--phone-number-type", "local",
		"--quantity", "1",
		"--requirement-group-id", "3fa85f64-5717-4562-b3fc-2c963f66afa6",
	)
}

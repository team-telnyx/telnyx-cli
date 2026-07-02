// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAdvancedOrdersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"area_code: xxx\n" +
			"comments: comments\n" +
			"country_code: xx\n" +
			"customer_reference: customer_reference\n" +
			"features:\n" +
			"  - sms\n" +
			"phone_number_type: local\n" +
			"quantity: 1\n" +
			"requirement_group_id: 3fa85f64-5717-4562-b3fc-2c963f66afa6\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"advanced-orders", "create",
		)
	})
}

func TestAdvancedOrdersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"advanced-orders", "retrieve",
			"--order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAdvancedOrdersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"advanced-orders", "list",
		)
	})
}

func TestAdvancedOrdersUpdateRequirementGroup(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"area_code: xxx\n" +
			"comments: comments\n" +
			"country_code: xx\n" +
			"customer_reference: customer_reference\n" +
			"features:\n" +
			"  - sms\n" +
			"phone_number_type: local\n" +
			"quantity: 1\n" +
			"requirement_group_id: 3fa85f64-5717-4562-b3fc-2c963f66afa6\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"advanced-orders", "update-requirement-group",
			"--advanced-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

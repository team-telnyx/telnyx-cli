// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRequirementGroupsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "create",
			"--api-key", "string",
			"--action", "ordering",
			"--country-code", "US",
			"--phone-number-type", "local",
			"--customer-reference", "My Requirement Group",
			"--regulatory-requirement", "{field_value: field_value, requirement_id: requirement_id}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(requirementGroupsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "create",
			"--api-key", "string",
			"--action", "ordering",
			"--country-code", "US",
			"--phone-number-type", "local",
			"--customer-reference", "My Requirement Group",
			"--regulatory-requirement.field-value", "field_value",
			"--regulatory-requirement.requirement-id", "requirement_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"action: ordering\n" +
			"country_code: US\n" +
			"phone_number_type: local\n" +
			"customer_reference: My Requirement Group\n" +
			"regulatory_requirements:\n" +
			"  - field_value: field_value\n" +
			"    requirement_id: requirement_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "requirement-groups", "create",
			"--api-key", "string",
		)
	})
}

func TestRequirementGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestRequirementGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "update",
			"--api-key", "string",
			"--id", "id",
			"--customer-reference", "0002",
			"--regulatory-requirement", "{field_value: new requirement value, requirement_id: 53970723-fbff-4f46-a975-f62be6c1a585}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(requirementGroupsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "update",
			"--api-key", "string",
			"--id", "id",
			"--customer-reference", "0002",
			"--regulatory-requirement.field-value", "new requirement value",
			"--regulatory-requirement.requirement-id", "53970723-fbff-4f46-a975-f62be6c1a585",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customer_reference: '0002'\n" +
			"regulatory_requirements:\n" +
			"  - field_value: new requirement value\n" +
			"    requirement_id: 53970723-fbff-4f46-a975-f62be6c1a585\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "requirement-groups", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestRequirementGroupsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "list",
			"--api-key", "string",
			"--filter", "{action: ordering, country_code: country_code, customer_reference: customer_reference, phone_number_type: local, status: approved}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(requirementGroupsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "list",
			"--api-key", "string",
			"--filter.action", "ordering",
			"--filter.country-code", "country_code",
			"--filter.customer-reference", "customer_reference",
			"--filter.phone-number-type", "local",
			"--filter.status", "approved",
		)
	})
}

func TestRequirementGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestRequirementGroupsSubmitForApproval(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "requirement-groups", "submit-for-approval",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

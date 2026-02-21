// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRequirementGroupsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "create",
		"--action", "ordering",
		"--country-code", "US",
		"--phone-number-type", "local",
		"--customer-reference", "My Requirement Group",
		"--regulatory-requirement", "{field_value: field_value, requirement_id: requirement_id}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(requirementGroupsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "create",
		"--action", "ordering",
		"--country-code", "US",
		"--phone-number-type", "local",
		"--customer-reference", "My Requirement Group",
		"--regulatory-requirement.field-value", "field_value",
		"--regulatory-requirement.requirement-id", "requirement_id",
	)
}

func TestRequirementGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "retrieve",
		"--id", "id",
	)
}

func TestRequirementGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "update",
		"--id", "id",
		"--customer-reference", "0002",
		"--regulatory-requirement", "{field_value: new requirement value, requirement_id: 53970723-fbff-4f46-a975-f62be6c1a585}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(requirementGroupsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "update",
		"--id", "id",
		"--customer-reference", "0002",
		"--regulatory-requirement.field-value", "new requirement value",
		"--regulatory-requirement.requirement-id", "53970723-fbff-4f46-a975-f62be6c1a585",
	)
}

func TestRequirementGroupsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "list",
		"--filter", "{action: ordering, country_code: country_code, customer_reference: customer_reference, phone_number_type: local, status: approved}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(requirementGroupsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "list",
		"--filter.action", "ordering",
		"--filter.country-code", "country_code",
		"--filter.customer-reference", "customer_reference",
		"--filter.phone-number-type", "local",
		"--filter.status", "approved",
	)
}

func TestRequirementGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "delete",
		"--id", "id",
	)
}

func TestRequirementGroupsSubmitForApproval(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-groups", "submit-for-approval",
		"--id", "id",
	)
}

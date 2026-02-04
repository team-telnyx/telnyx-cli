// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersActionRequirementsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:action-requirements", "list",
		"--porting-order-id", "porting_order_id",
		"--filter", "{id: [182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e], action_type: au_id_verification, requirement_type_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, status: created}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "{value: created_at}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersActionRequirementsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:action-requirements", "list",
		"--porting-order-id", "porting_order_id",
		"--filter.id", "[182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e]",
		"--filter.action-type", "au_id_verification",
		"--filter.requirement-type-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.status", "created",
		"--page-number", "0",
		"--page-size", "0",
		"--sort.value", "created_at",
	)
}

func TestPortingOrdersActionRequirementsInitiate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:action-requirements", "initiate",
		"--porting-order-id", "porting_order_id",
		"--id", "id",
		"--params", "{first_name: John, last_name: Doe}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersActionRequirementsInitiate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:action-requirements", "initiate",
		"--porting-order-id", "porting_order_id",
		"--id", "id",
		"--params.first-name", "John",
		"--params.last-name", "Doe",
	)
}

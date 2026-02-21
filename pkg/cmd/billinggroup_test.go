// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestBillingGroupsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"billing-groups", "create",
		"--name", "string",
	)
}

func TestBillingGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"billing-groups", "retrieve",
		"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
	)
}

func TestBillingGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"billing-groups", "update",
		"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
		"--name", "string",
	)
}

func TestBillingGroupsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"billing-groups", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestBillingGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"billing-groups", "delete",
		"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
	)
}

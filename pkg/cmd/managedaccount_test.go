// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestManagedAccountsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts", "create",
		"--business-name", "Larry's Cat Food Inc",
		"--email", "larry_cat_food@customer.org",
		"--managed-account-allow-custom-pricing=false",
		"--password", "3jVjLq!tMuWKyWx4NN*CvhnB",
		"--rollup-billing=false",
	)
}

func TestManagedAccountsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts", "retrieve",
		"--id", "id",
	)
}

func TestManagedAccountsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts", "update",
		"--id", "id",
		"--managed-account-allow-custom-pricing=true",
	)
}

func TestManagedAccountsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts", "list",
		"--filter", "{email: {contains: john, eq: eq}, organization_name: {contains: contains, eq: Example Company LLC}}",
		"--include-cancelled-accounts=true",
		"--page", "{number: 1, size: 1}",
		"--sort", "email",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(managedAccountsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts", "list",
		"--filter.email", "{contains: john, eq: eq}",
		"--filter.organization-name", "{contains: contains, eq: Example Company LLC}",
		"--include-cancelled-accounts=true",
		"--page.number", "1",
		"--page.size", "1",
		"--sort", "email",
	)
}

func TestManagedAccountsGetAllocatableGlobalOutboundChannels(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts", "get-allocatable-global-outbound-channels",
	)
}

func TestManagedAccountsUpdateGlobalChannelLimit(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"managed-accounts", "update-global-channel-limit",
		"--id", "id",
		"--channel-limit", "30",
	)
}

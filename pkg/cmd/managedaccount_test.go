// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestManagedAccountsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts", "create",
			"--api-key", "string",
			"--business-name", "Larry's Cat Food Inc",
			"--email", "larry_cat_food@customer.org",
			"--managed-account-allow-custom-pricing=false",
			"--password", "3jVjLq!tMuWKyWx4NN*CvhnB",
			"--rollup-billing=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"business_name: Larry's Cat Food Inc\n" +
			"email: larry_cat_food@customer.org\n" +
			"managed_account_allow_custom_pricing: false\n" +
			"password: 3jVjLq!tMuWKyWx4NN*CvhnB\n" +
			"rollup_billing: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "managed-accounts", "create",
			"--api-key", "string",
		)
	})
}

func TestManagedAccountsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestManagedAccountsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts", "update",
			"--api-key", "string",
			"--id", "id",
			"--managed-account-allow-custom-pricing=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("managed_account_allow_custom_pricing: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "managed-accounts", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestManagedAccountsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{email: {contains: john, eq: eq}, organization_name: {contains: contains, eq: Example Company LLC}}",
			"--include-cancelled-accounts=true",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "email",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(managedAccountsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.email", "{contains: john, eq: eq}",
			"--filter.organization-name", "{contains: contains, eq: Example Company LLC}",
			"--include-cancelled-accounts=true",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "email",
		)
	})
}

func TestManagedAccountsGetAllocatableGlobalOutboundChannels(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts", "get-allocatable-global-outbound-channels",
			"--api-key", "string",
		)
	})
}

func TestManagedAccountsUpdateGlobalChannelLimit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "managed-accounts", "update-global-channel-limit",
			"--api-key", "string",
			"--id", "id",
			"--channel-limit", "30",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("channel_limit: 30")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "managed-accounts", "update-global-channel-limit",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

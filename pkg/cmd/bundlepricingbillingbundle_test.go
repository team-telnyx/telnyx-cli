// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestBundlePricingBillingBundlesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bundle-pricing:billing-bundles", "retrieve",
			"--bundle-id", "8661948c-a386-4385-837f-af00f40f111a",
			"--authorization-bearer", "authorization_bearer",
		)
	})
}

func TestBundlePricingBillingBundlesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bundle-pricing:billing-bundles", "list",
			"--max-items", "10",
			"--filter", "{country_iso: [US], resource: ['+15617819942']}",
			"--page-number", "0",
			"--page-size", "0",
			"--authorization-bearer", "authorization_bearer",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(bundlePricingBillingBundlesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bundle-pricing:billing-bundles", "list",
			"--max-items", "10",
			"--filter.country-iso", "[US]",
			"--filter.resource", "['+15617819942']",
			"--page-number", "0",
			"--page-size", "0",
			"--authorization-bearer", "authorization_bearer",
		)
	})
}

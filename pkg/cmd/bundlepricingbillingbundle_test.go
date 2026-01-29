// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestBundlePricingBillingBundlesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:billing-bundles", "retrieve",
		"--bundle-id", "8661948c-a386-4385-837f-af00f40f111a",
		"--authorization-bearer", "authorization_bearer",
	)
}

func TestBundlePricingBillingBundlesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:billing-bundles", "list",
		"--filter", "{country_iso: [US], resource: ['+15617819942']}",
		"--page", "{number: 1, size: 1}",
		"--authorization-bearer", "authorization_bearer",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(bundlePricingBillingBundlesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:billing-bundles", "list",
		"--filter.country-iso", "[US]",
		"--filter.resource", "['+15617819942']",
		"--page.number", "1",
		"--page.size", "1",
		"--authorization-bearer", "authorization_bearer",
	)
}

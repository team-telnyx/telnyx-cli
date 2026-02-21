// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestBundlePricingUserBundlesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "create",
		"--idempotency-key", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--item", "{billing_bundle_id: 12ade33a-21c0-473b-b055-b3c836e1c292, quantity: 0}",
		"--authorization-bearer", "authorization_bearer",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(bundlePricingUserBundlesCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "create",
		"--idempotency-key", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--item.billing-bundle-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--item.quantity", "0",
		"--authorization-bearer", "authorization_bearer",
	)
}

func TestBundlePricingUserBundlesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "retrieve",
		"--user-bundle-id", "ca1d2263-d1f1-43ac-ba53-248e7a4bb26a",
		"--authorization-bearer", "authorization_bearer",
	)
}

func TestBundlePricingUserBundlesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "list",
		"--filter", "{country_iso: [US], resource: ['+15617819942']}",
		"--page-number", "0",
		"--page-size", "0",
		"--authorization-bearer", "authorization_bearer",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(bundlePricingUserBundlesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "list",
		"--filter.country-iso", "[US]",
		"--filter.resource", "['+15617819942']",
		"--page-number", "0",
		"--page-size", "0",
		"--authorization-bearer", "authorization_bearer",
	)
}

func TestBundlePricingUserBundlesDeactivate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "deactivate",
		"--user-bundle-id", "ca1d2263-d1f1-43ac-ba53-248e7a4bb26a",
		"--authorization-bearer", "authorization_bearer",
	)
}

func TestBundlePricingUserBundlesListResources(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "list-resources",
		"--user-bundle-id", "ca1d2263-d1f1-43ac-ba53-248e7a4bb26a",
		"--authorization-bearer", "authorization_bearer",
	)
}

func TestBundlePricingUserBundlesListUnused(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "list-unused",
		"--filter", "{country_iso: [US], resource: ['+15617819942']}",
		"--authorization-bearer", "authorization_bearer",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(bundlePricingUserBundlesListUnused)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"bundle-pricing:user-bundles", "list-unused",
		"--filter.country-iso", "[US]",
		"--filter.resource", "['+15617819942']",
		"--authorization-bearer", "authorization_bearer",
	)
}

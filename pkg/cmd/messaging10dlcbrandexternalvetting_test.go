// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcBrandExternalVettingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand:external-vetting", "list",
		"--api-key", "string",
		"--brand-id", "brandId",
	)
}

func TestMessaging10dlcBrandExternalVettingImports(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand:external-vetting", "imports",
		"--api-key", "string",
		"--brand-id", "brandId",
		"--evp-id", "evpId",
		"--vetting-id", "vettingId",
		"--vetting-token", "vettingToken",
	)
}

func TestMessaging10dlcBrandExternalVettingOrder(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand:external-vetting", "order",
		"--api-key", "string",
		"--brand-id", "brandId",
		"--evp-id", "evpId",
		"--vetting-class", "vettingClass",
	)
}

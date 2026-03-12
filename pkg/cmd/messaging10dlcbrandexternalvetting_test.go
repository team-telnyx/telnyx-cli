// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcBrandExternalVettingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand:external-vetting", "list",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandExternalVettingImports(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand:external-vetting", "imports",
			"--api-key", "string",
			"--brand-id", "brandId",
			"--evp-id", "evpId",
			"--vetting-id", "vettingId",
			"--vetting-token", "vettingToken",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"evpId: evpId\n" +
			"vettingId: vettingId\n" +
			"vettingToken: vettingToken\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-10dlc:brand:external-vetting", "imports",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandExternalVettingOrder(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand:external-vetting", "order",
			"--api-key", "string",
			"--brand-id", "brandId",
			"--evp-id", "evpId",
			"--vetting-class", "vettingClass",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"evpId: evpId\n" +
			"vettingClass: vettingClass\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-10dlc:brand:external-vetting", "order",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

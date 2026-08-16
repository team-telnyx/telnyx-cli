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
			t,
			"--api-key", "string",
			"messaging-10dlc:brand:external-vetting", "list",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandExternalVettingImports(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:brand:external-vetting", "imports",
			"--brand-id", "brandId",
			"--evp-id", "Evpid",
			"--vetting-id", "Vettingid",
			"--vetting-token", "Vettingtoken",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"evpId: Evpid\n" +
			"vettingId: Vettingid\n" +
			"vettingToken: Vettingtoken\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-10dlc:brand:external-vetting", "imports",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandExternalVettingOrder(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:brand:external-vetting", "order",
			"--brand-id", "brandId",
			"--evp-id", "Evpid",
			"--vetting-class", "Vettingclass",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"evpId: Evpid\n" +
			"vettingClass: Vettingclass\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-10dlc:brand:external-vetting", "order",
			"--brand-id", "brandId",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestDynamicEmergencyEndpointsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "dynamic-emergency-endpoints", "create",
			"--api-key", "string",
			"--callback-number", "+13125550000",
			"--caller-name", "Jane Doe Desk Phone",
			"--dynamic-emergency-address-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"callback_number: '+13125550000'\n" +
			"caller_name: Jane Doe Desk Phone\n" +
			"dynamic_emergency_address_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "dynamic-emergency-endpoints", "create",
			"--api-key", "string",
		)
	})
}

func TestDynamicEmergencyEndpointsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "dynamic-emergency-endpoints", "retrieve",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestDynamicEmergencyEndpointsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "dynamic-emergency-endpoints", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{country_code: country_code, status: pending}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(dynamicEmergencyEndpointsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "dynamic-emergency-endpoints", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.country-code", "country_code",
			"--filter.status", "pending",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestDynamicEmergencyEndpointsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "dynamic-emergency-endpoints", "delete",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

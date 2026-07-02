// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPhoneNumbersActionsChangeBundleStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers:actions", "change-bundle-status",
			"--id", "1293384261075731499",
			"--bundle-id", "5194d8fc-87e6-4188-baa9-1c434bbe861b",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("bundle_id: 5194d8fc-87e6-4188-baa9-1c434bbe861b")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"phone-numbers:actions", "change-bundle-status",
			"--id", "1293384261075731499",
		)
	})
}

func TestPhoneNumbersActionsEnableEmergency(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers:actions", "enable-emergency",
			"--id", "1293384261075731499",
			"--emergency-address-id", "53829456729313",
			"--emergency-enabled=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"emergency_address_id: '53829456729313'\n" +
			"emergency_enabled: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"phone-numbers:actions", "enable-emergency",
			"--id", "1293384261075731499",
		)
	})
}

func TestPhoneNumbersActionsVerifyOwnership(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers:actions", "verify-ownership",
			"--phone-number", "+15551234567",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - '+15551234567'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"phone-numbers:actions", "verify-ownership",
		)
	})
}

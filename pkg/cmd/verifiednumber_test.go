// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVerifiedNumbersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verified-numbers", "create",
			"--api-key", "string",
			"--phone-number", "+15551234567",
			"--verification-method", "sms",
			"--extension", "ww243w1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_number: '+15551234567'\n" +
			"verification_method: sms\n" +
			"extension: ww243w1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "verified-numbers", "create",
			"--api-key", "string",
		)
	})
}

func TestVerifiedNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verified-numbers", "retrieve",
			"--api-key", "string",
			"--phone-number", "+15551234567",
		)
	})
}

func TestVerifiedNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verified-numbers", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestVerifiedNumbersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verified-numbers", "delete",
			"--api-key", "string",
			"--phone-number", "+15551234567",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEnterprisesReputationNumbersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "create",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--phone-number", "+16035551234",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - '+16035551234'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:reputation:numbers", "create",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestEnterprisesReputationNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "retrieve",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--phone-number", "+16035551234",
			"--fresh=true",
		)
	})
}

func TestEnterprisesReputationNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "list",
			"--max-items", "10",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--page-number", "1",
			"--page-size", "1",
			"--phone-number", "+16035551234",
		)
	})
}

func TestEnterprisesReputationNumbersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "delete",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--phone-number", "+16035551234",
		)
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestReputationNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"reputation:numbers", "retrieve",
			"--phone-number", "+19493253498",
			"--fresh=true",
		)
	})
}

func TestReputationNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"reputation:numbers", "list",
			"--max-items", "10",
			"--filter-enterprise-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-phone-number-contains", "+16035551234",
			"--filter-phone-number-eq", "+16035551234",
			"--page-number", "1",
			"--page-size", "20",
		)
	})
}

func TestReputationNumbersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"reputation:numbers", "delete",
			"--phone-number", "+19493253498",
		)
	})
}

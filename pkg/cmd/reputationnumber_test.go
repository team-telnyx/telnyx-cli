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
			"--phone-number", "+16035551234",
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
			"--page-number", "1",
			"--page-size", "1",
			"--phone-number", "+16035551234",
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
			"--phone-number", "+16035551234",
		)
	})
}

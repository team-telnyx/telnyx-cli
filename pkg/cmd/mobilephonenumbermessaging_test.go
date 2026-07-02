// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMobilePhoneNumbersMessagingRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-phone-numbers:messaging", "retrieve",
			"--id", "id",
		)
	})
}

func TestMobilePhoneNumbersMessagingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-phone-numbers:messaging", "list",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

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
			t, "mobile-phone-numbers:messaging", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestMobilePhoneNumbersMessagingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "mobile-phone-numbers:messaging", "list",
			"--api-key", "string",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

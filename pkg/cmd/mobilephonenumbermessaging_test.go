// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMobilePhoneNumbersMessagingRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-phone-numbers:messaging", "retrieve",
		"--id", "id",
	)
}

func TestMobilePhoneNumbersMessagingList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-phone-numbers:messaging", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

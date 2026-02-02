// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestPhoneNumbersMessagingRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:messaging", "retrieve",
		"--id", "id",
	)
}

func TestPhoneNumbersMessagingUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:messaging", "update",
		"--id", "id",
		"--messaging-product", "P2P",
		"--messaging-profile-id", "dd50eba1-a0c0-4563-9925-b25e842a7cb6",
	)
}

func TestPhoneNumbersMessagingList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:messaging", "list",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumbersMessagingList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:messaging", "list",
		"--page.number", "1",
		"--page.size", "1",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPhoneNumbersMessagingRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:messaging", "retrieve",
		"--id", "id",
	)
}

func TestPhoneNumbersMessagingUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:messaging", "update",
		"--id", "id",
		"--messaging-product", "P2P",
		"--messaging-profile-id", "dd50eba1-a0c0-4563-9925-b25e842a7cb6",
		"--tag", "string",
	)
}

func TestPhoneNumbersMessagingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:messaging", "list",
		"--filter-messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter-phone-number", "filter[phone_number]",
		"--filter-phone-number-contains", "filter[phone_number][contains]",
		"--filter-type", "tollfree",
		"--page-number", "0",
		"--page-size", "0",
		"--sort-phone-number", "asc",
	)
}

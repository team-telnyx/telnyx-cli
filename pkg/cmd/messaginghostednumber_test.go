// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessagingHostedNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-numbers", "retrieve",
		"--id", "id",
	)
}

func TestMessagingHostedNumbersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-numbers", "update",
		"--id", "id",
		"--messaging-product", "P2P",
		"--messaging-profile-id", "dd50eba1-a0c0-4563-9925-b25e842a7cb6",
		"--tag", "string",
	)
}

func TestMessagingHostedNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-numbers", "list",
		"--filter-messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter-phone-number", "filter[phone_number]",
		"--filter-phone-number-contains", "filter[phone_number][contains]",
		"--page-number", "0",
		"--page-size", "0",
		"--sort-phone-number", "asc",
	)
}

func TestMessagingHostedNumbersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-numbers", "delete",
		"--id", "id",
	)
}

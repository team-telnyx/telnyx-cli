// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAlphanumericSenderIDsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"alphanumeric-sender-ids", "create",
		"--alphanumeric-sender-id", "MyCompany",
		"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--us-long-code-fallback", "+15551234567",
	)
}

func TestAlphanumericSenderIDsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"alphanumeric-sender-ids", "retrieve",
		"--id", "id",
	)
}

func TestAlphanumericSenderIDsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"alphanumeric-sender-ids", "list",
		"--filter-messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestAlphanumericSenderIDsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"alphanumeric-sender-ids", "delete",
		"--id", "id",
	)
}

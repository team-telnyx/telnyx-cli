// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestNotificationProfilesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-profiles", "create",
		"--name", "name",
	)
}

func TestNotificationProfilesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-profiles", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestNotificationProfilesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-profiles", "update",
		"--notification-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--name", "name",
	)
}

func TestNotificationProfilesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-profiles", "list",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(notificationProfilesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-profiles", "list",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestNotificationProfilesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-profiles", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

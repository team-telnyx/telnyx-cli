// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestOrganizationsUsersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"organizations:users", "retrieve",
		"--id", "id",
		"--include-groups=true",
	)
}

func TestOrganizationsUsersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"organizations:users", "list",
		"--filter-email", "filter[email]",
		"--filter-user-status", "enabled",
		"--include-groups=true",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestOrganizationsUsersGetGroupsReport(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"organizations:users", "get-groups-report",
		"--accept", "application/json",
	)
}

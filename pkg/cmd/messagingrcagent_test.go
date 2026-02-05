// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessagingRcsAgentsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging:rcs:agents", "retrieve",
		"--id", "id",
	)
}

func TestMessagingRcsAgentsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging:rcs:agents", "update",
		"--id", "id",
		"--profile-id", "4001932a-b8a3-42fc-9389-021be6388909",
		"--webhook-failover-url", "http://example.com",
		"--webhook-url", "http://example.com",
	)
}

func TestMessagingRcsAgentsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging:rcs:agents", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

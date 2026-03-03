// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessagingRcsInviteTestNumber(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging:rcs", "invite-test-number",
		"--api-key", "string",
		"--id", "id",
		"--phone-number", "phone_number",
	)
}

func TestMessagingRcsListBulkCapabilities(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging:rcs", "list-bulk-capabilities",
		"--api-key", "string",
		"--agent-id", "TestAgent",
		"--phone-number", "+13125551234",
	)
}

func TestMessagingRcsRetrieveCapabilities(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging:rcs", "retrieve-capabilities",
		"--api-key", "string",
		"--agent-id", "agent_id",
		"--phone-number", "phone_number",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessagingRcsInviteTestNumber(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging:rcs", "invite-test-number",
			"--id", "id",
			"--phone-number", "phone_number",
		)
	})
}

func TestMessagingRcsListBulkCapabilities(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging:rcs", "list-bulk-capabilities",
			"--agent-id", "TestAgent",
			"--phone-number", "+13125551234",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_id: TestAgent\n" +
			"phone_numbers:\n" +
			"  - '+13125551234'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging:rcs", "list-bulk-capabilities",
		)
	})
}

func TestMessagingRcsRetrieveCapabilities(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging:rcs", "retrieve-capabilities",
			"--agent-id", "agent_id",
			"--phone-number", "phone_number",
		)
	})
}

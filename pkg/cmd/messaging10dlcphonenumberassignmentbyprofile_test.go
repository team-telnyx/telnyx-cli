// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcPhoneNumberAssignmentByProfileAssign(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-assignment-by-profile", "assign",
			"--messaging-profile-id", "4001767e-ce0f-4cae-9d5f-0d5e636e7809",
			"--campaign-id", "4b300178-131c-d902-d54e-72d90ba1620j",
			"--tcr-campaign-id", "CWZTFH1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"messagingProfileId: 4001767e-ce0f-4cae-9d5f-0d5e636e7809\n" +
			"campaignId: 4b300178-131c-d902-d54e-72d90ba1620j\n" +
			"tcrCampaignId: CWZTFH1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-10dlc:phone-number-assignment-by-profile", "assign",
		)
	})
}

func TestMessaging10dlcPhoneNumberAssignmentByProfileListPhoneNumberStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-assignment-by-profile", "list-phone-number-status",
			"--task-id", "taskId",
			"--page", "0",
			"--records-per-page", "0",
		)
	})
}

func TestMessaging10dlcPhoneNumberAssignmentByProfileRetrievePhoneNumberStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-assignment-by-profile", "retrieve-phone-number-status",
			"--task-id", "taskId",
			"--page", "0",
			"--records-per-page", "0",
		)
	})
}

func TestMessaging10dlcPhoneNumberAssignmentByProfileRetrieveStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-10dlc:phone-number-assignment-by-profile", "retrieve-status",
			"--task-id", "taskId",
		)
	})
}

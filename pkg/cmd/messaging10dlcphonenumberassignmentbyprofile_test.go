// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcPhoneNumberAssignmentByProfileAssign(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:phone-number-assignment-by-profile", "assign",
		"--messaging-profile-id", "4001767e-ce0f-4cae-9d5f-0d5e636e7809",
		"--campaign-id", "4b300178-131c-d902-d54e-72d90ba1620j",
		"--tcr-campaign-id", "CWZTFH1",
	)
}

func TestMessaging10dlcPhoneNumberAssignmentByProfileListPhoneNumberStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:phone-number-assignment-by-profile", "list-phone-number-status",
		"--task-id", "taskId",
		"--page", "0",
		"--records-per-page", "0",
	)
}

func TestMessaging10dlcPhoneNumberAssignmentByProfileRetrievePhoneNumberStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:phone-number-assignment-by-profile", "retrieve-phone-number-status",
		"--task-id", "taskId",
		"--page", "0",
		"--records-per-page", "0",
	)
}

func TestMessaging10dlcPhoneNumberAssignmentByProfileRetrieveStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:phone-number-assignment-by-profile", "retrieve-status",
		"--task-id", "taskId",
	)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPhoneNumbersVoicemailCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:voicemail", "create",
		"--phone-number-id", "123455678900",
		"--enabled=true",
		"--pin", "1234",
	)
}

func TestPhoneNumbersVoicemailRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:voicemail", "retrieve",
		"--phone-number-id", "123455678900",
	)
}

func TestPhoneNumbersVoicemailUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-numbers:voicemail", "update",
		"--phone-number-id", "123455678900",
		"--enabled=true",
		"--pin", "1234",
	)
}

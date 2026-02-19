// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMessagingOptoutsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-optouts", "list",
		"--created-at", "{gte: '2019-12-27T18:11:19.117Z', lte: '2019-12-27T18:11:19.117Z'}",
		"--filter", "{from: from, messaging_profile_id: messaging_profile_id}",
		"--page-number", "0",
		"--page-size", "0",
		"--redaction-enabled", "redaction_enabled",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingOptoutsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-optouts", "list",
		"--created-at.gte", "2019-12-27T18:11:19.117Z",
		"--created-at.lte", "2019-12-27T18:11:19.117Z",
		"--filter.from", "from",
		"--filter.messaging-profile-id", "messaging_profile_id",
		"--page-number", "0",
		"--page-size", "0",
		"--redaction-enabled", "redaction_enabled",
	)
}

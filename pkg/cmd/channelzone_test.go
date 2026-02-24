// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestChannelZonesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"channel-zones", "update",
		"--channel-zone-id", "channel_zone_id",
		"--channels", "0",
	)
}

func TestChannelZonesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"channel-zones", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

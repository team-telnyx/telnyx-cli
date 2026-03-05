// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestChannelZonesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "channel-zones", "update",
			"--api-key", "string",
			"--channel-zone-id", "channel_zone_id",
			"--channels", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("channels: 0")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "channel-zones", "update",
			"--api-key", "string",
			"--channel-zone-id", "channel_zone_id",
		)
	})
}

func TestChannelZonesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "channel-zones", "list",
			"--api-key", "string",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

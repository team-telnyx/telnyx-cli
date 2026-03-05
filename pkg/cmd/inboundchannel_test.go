// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestInboundChannelsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "inbound-channels", "update",
			"--api-key", "string",
			"--channels", "7",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("channels: 7")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "inbound-channels", "update",
			"--api-key", "string",
		)
	})
}

func TestInboundChannelsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "inbound-channels", "list",
			"--api-key", "string",
		)
	})
}

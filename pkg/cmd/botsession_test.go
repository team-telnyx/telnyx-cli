// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestBotSessionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bot-sessions", "list",
			"--email", "agent-owner@example.com",
			"--portal-redirect-token", "01890a7e-e2f7-7c3d-8dbb-9a2c5f3d1e0b",
		)
	})
}

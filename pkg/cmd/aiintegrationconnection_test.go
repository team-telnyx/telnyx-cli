// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIIntegrationsConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:integrations:connections", "retrieve",
			"--api-key", "string",
			"--user-connection-id", "user_connection_id",
		)
	})
}

func TestAIIntegrationsConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:integrations:connections", "list",
			"--api-key", "string",
		)
	})
}

func TestAIIntegrationsConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:integrations:connections", "delete",
			"--api-key", "string",
			"--user-connection-id", "user_connection_id",
		)
	})
}

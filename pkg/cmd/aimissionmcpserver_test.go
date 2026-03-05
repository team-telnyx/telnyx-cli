// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsMcpServersCreateMcpServer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:mcp-servers", "create-mcp-server",
			"--api-key", "string",
			"--mission-id", "mission_id",
		)
	})
}

func TestAIMissionsMcpServersDeleteMcpServer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:mcp-servers", "delete-mcp-server",
			"--api-key", "string",
			"--mission-id", "mission_id",
			"--mcp-server-id", "mcp_server_id",
		)
	})
}

func TestAIMissionsMcpServersGetMcpServer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:mcp-servers", "get-mcp-server",
			"--api-key", "string",
			"--mission-id", "mission_id",
			"--mcp-server-id", "mcp_server_id",
		)
	})
}

func TestAIMissionsMcpServersListMcpServers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:mcp-servers", "list-mcp-servers",
			"--api-key", "string",
			"--mission-id", "mission_id",
		)
	})
}

func TestAIMissionsMcpServersUpdateMcpServer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:mcp-servers", "update-mcp-server",
			"--api-key", "string",
			"--mission-id", "mission_id",
			"--mcp-server-id", "mcp_server_id",
		)
	})
}

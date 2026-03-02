// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMcpServersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:mcp-servers", "create",
		"--api-key", "string",
		"--name", "name",
		"--type", "type",
		"--url", "url",
		"--allowed-tool", "[string]",
		"--api-key-ref", "api_key_ref",
	)
}

func TestAIMcpServersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:mcp-servers", "retrieve",
		"--api-key", "string",
		"--mcp-server-id", "mcp_server_id",
	)
}

func TestAIMcpServersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:mcp-servers", "update",
		"--api-key", "string",
		"--mcp-server-id", "mcp_server_id",
		"--id", "id",
		"--allowed-tool", "[string]",
		"--api-key-ref", "api_key_ref",
		"--created-at", "'2019-12-27T18:11:19.117Z'",
		"--name", "name",
		"--type", "type",
		"--url", "url",
	)
}

func TestAIMcpServersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:mcp-servers", "list",
		"--api-key", "string",
		"--page-number", "1",
		"--page-size", "1",
		"--type", "type",
		"--url", "url",
	)
}

func TestAIMcpServersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:mcp-servers", "delete",
		"--api-key", "string",
		"--mcp-server-id", "mcp_server_id",
	)
}

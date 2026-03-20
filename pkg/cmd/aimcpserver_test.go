// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMcpServersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:mcp-servers", "create",
			"--name", "name",
			"--type", "type",
			"--url", "url",
			"--allowed-tool", "[string]",
			"--api-key-ref", "api_key_ref",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"type: type\n" +
			"url: url\n" +
			"allowed_tools:\n" +
			"  - string\n" +
			"api_key_ref: api_key_ref\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:mcp-servers", "create",
		)
	})
}

func TestAIMcpServersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:mcp-servers", "retrieve",
			"--mcp-server-id", "mcp_server_id",
		)
	})
}

func TestAIMcpServersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:mcp-servers", "update",
			"--mcp-server-id", "mcp_server_id",
			"--id", "id",
			"--allowed-tool", "[string]",
			"--api-key-ref", "api_key_ref",
			"--created-at", "'2019-12-27T18:11:19.117Z'",
			"--name", "name",
			"--type", "type",
			"--url", "url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"allowed_tools:\n" +
			"  - string\n" +
			"api_key_ref: api_key_ref\n" +
			"created_at: '2019-12-27T18:11:19.117Z'\n" +
			"name: name\n" +
			"type: type\n" +
			"url: url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:mcp-servers", "update",
			"--mcp-server-id", "mcp_server_id",
		)
	})
}

func TestAIMcpServersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:mcp-servers", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
			"--type", "type",
			"--url", "url",
		)
	})
}

func TestAIMcpServersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:mcp-servers", "delete",
			"--mcp-server-id", "mcp_server_id",
		)
	})
}

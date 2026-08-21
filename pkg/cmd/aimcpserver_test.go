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
			"--name", "Name",
			"--type", "Type",
			"--url", "Url",
			"--allowed-tool", "[string]",
			"--api-key-ref", "api_key_ref",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Name\n" +
			"type: Type\n" +
			"url: Url\n" +
			"allowed_tools:\n" +
			"  - string\n" +
			"api_key_ref: api_key_ref\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:mcp-servers", "create",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
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
			"--id", "Id",
			"--allowed-tool", "[string]",
			"--api-key-ref", "api_key_ref",
			"--created-at", "'2024-01-23T18:10:02.574Z'",
			"--name", "Name",
			"--type", "Type",
			"--url", "Url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: Id\n" +
			"allowed_tools:\n" +
			"  - string\n" +
			"api_key_ref: api_key_ref\n" +
			"created_at: '2024-01-23T18:10:02.574Z'\n" +
			"name: Name\n" +
			"type: Type\n" +
			"url: Url\n")
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

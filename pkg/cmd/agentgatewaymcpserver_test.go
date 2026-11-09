// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAgentGatewaysMcpServersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:mcp-servers", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "orders",
			"--url", "http://10.40.3.7:9000/mcp",
			"--auth", "{type: api_key, value: example-orders-key, header: header}",
			"--forwarded-header", "X-Request-Id",
			"--open-to-every-key=true",
			"--static-headers", "{X-Env: prod}",
			"--tool", "string",
			"--transport", "http",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: orders\n" +
			"url: http://10.40.3.7:9000/mcp\n" +
			"auth:\n" +
			"  type: api_key\n" +
			"  value: example-orders-key\n" +
			"  header: header\n" +
			"forwarded_headers:\n" +
			"  - X-Request-Id\n" +
			"open_to_every_key: true\n" +
			"static_headers:\n" +
			"  X-Env: prod\n" +
			"tools:\n" +
			"  - string\n" +
			"transport: http\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:mcp-servers", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysMcpServersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:mcp-servers", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--server-id", "88d5a97b-4080-4855-8bac-a39811b1dc24",
			"--auth", "{type: none}",
			"--forwarded-header", "string",
			"--name", "name",
			"--open-to-every-key=true",
			"--static-headers", "{foo: string}",
			"--tool", "get_order",
			"--tool", "create_ticket",
			"--transport", "http",
			"--url", "url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auth:\n" +
			"  type: none\n" +
			"forwarded_headers:\n" +
			"  - string\n" +
			"name: name\n" +
			"open_to_every_key: true\n" +
			"static_headers:\n" +
			"  foo: string\n" +
			"tools:\n" +
			"  - get_order\n" +
			"  - create_ticket\n" +
			"transport: http\n" +
			"url: url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:mcp-servers", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--server-id", "88d5a97b-4080-4855-8bac-a39811b1dc24",
		)
	})
}

func TestAgentGatewaysMcpServersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:mcp-servers", "list",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysMcpServersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:mcp-servers", "delete",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--server-id", "88d5a97b-4080-4855-8bac-a39811b1dc24",
		)
	})
}

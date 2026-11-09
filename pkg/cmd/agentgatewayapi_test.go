// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAgentGatewaysAPIsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:apis", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "petstore",
			"--spec-url", "https://petstore3.swagger.io/api/v3/openapi.json",
			"--auth", "{type: none}",
			"--base-url", "https://petstore3.swagger.io/api/v3",
			"--forwarded-header", "string",
			"--open-to-every-key=true",
			"--static-headers", "{foo: string}",
			"--tool", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: petstore\n" +
			"spec_url: https://petstore3.swagger.io/api/v3/openapi.json\n" +
			"auth:\n" +
			"  type: none\n" +
			"base_url: https://petstore3.swagger.io/api/v3\n" +
			"forwarded_headers:\n" +
			"  - string\n" +
			"open_to_every_key: true\n" +
			"static_headers:\n" +
			"  foo: string\n" +
			"tools:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:apis", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysAPIsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:apis", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--api-id", "dbe53cc0-bf7e-4f7a-8d84-752d19890eb3",
			"--auth", "{type: none}",
			"--base-url", "base_url",
			"--forwarded-header", "string",
			"--name", "name",
			"--open-to-every-key=true",
			"--spec-url", "spec_url",
			"--static-headers", "{foo: string}",
			"--tool", "findpetsbystatus",
			"--tool", "getpetbyid",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auth:\n" +
			"  type: none\n" +
			"base_url: base_url\n" +
			"forwarded_headers:\n" +
			"  - string\n" +
			"name: name\n" +
			"open_to_every_key: true\n" +
			"spec_url: spec_url\n" +
			"static_headers:\n" +
			"  foo: string\n" +
			"tools:\n" +
			"  - findpetsbystatus\n" +
			"  - getpetbyid\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:apis", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--api-id", "dbe53cc0-bf7e-4f7a-8d84-752d19890eb3",
		)
	})
}

func TestAgentGatewaysAPIsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:apis", "list",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysAPIsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:apis", "delete",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--api-id", "dbe53cc0-bf7e-4f7a-8d84-752d19890eb3",
		)
	})
}
